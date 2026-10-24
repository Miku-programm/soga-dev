package httpobfs

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
	"unsafe"
)

type memoryConn struct {
	reader   io.Reader
	writes   [][]byte
	writeErr error
	failAt   int
}

func (c *memoryConn) Read(p []byte) (int, error) { return c.reader.Read(p) }
func (c *memoryConn) Write(p []byte) (int, error) {
	c.writes = append(c.writes, append([]byte(nil), p...))
	if c.failAt == 0 || len(c.writes) == c.failAt {
		return len(p), c.writeErr
	}
	return len(p), nil
}
func (*memoryConn) Close() error                     { return nil }
func (*memoryConn) LocalAddr() net.Addr              { return &net.TCPAddr{} }
func (*memoryConn) RemoteAddr() net.Addr             { return &net.TCPAddr{} }
func (*memoryConn) SetDeadline(time.Time) error      { return nil }
func (*memoryConn) SetReadDeadline(time.Time) error  { return nil }
func (*memoryConn) SetWriteDeadline(time.Time) error { return nil }
func TestHTTPFirstWriteAndRawSubsequent(t *testing.T) {
	raw := &memoryConn{}
	c := NewHTTPObfs(raw, "example.test", "8080")
	p := []byte("payload")
	if n, e := c.Write(p); n != len(p) || e != nil {
		t.Fatal(n, e)
	}
	req, e := http.ReadRequest(bufio.NewReader(bytes.NewReader(bytes.Join(raw.writes, nil))))
	if e != nil {
		t.Fatal(e)
	}
	body, e := io.ReadAll(req.Body)
	if e != nil || !bytes.Equal(body, p) || req.Method != "GET" || req.URL.Path != "/" || req.Host != "example.test:8080" || req.Header.Get("Connection") != "Upgrade" || req.Header.Get("Upgrade") != "websocket" {
		t.Fatal(req, e)
	}
	if key, e := base64.URLEncoding.DecodeString(req.Header.Get("Sec-WebSocket-Key")); e != nil || len(key) != 16 {
		t.Fatal(e)
	}
	if !strings.HasPrefix(req.UserAgent(), "curl/7.") {
		t.Fatal(req.UserAgent())
	}
	raw.writes = nil
	if _, e = c.Write(p); e != nil || !bytes.Equal(raw.writes[0], p) {
		t.Fatal(e)
	}
}
func TestHTTPReadHeaderBufferAndFailure(t *testing.T) {
	raw := &memoryConn{reader: bytes.NewReader([]byte("arbitrary\r\n\r\nabcdef"))}
	c := NewHTTPObfs(raw, "x", "80")
	p := make([]byte, 2)
	var got []byte
	for i := 0; i < 3; i++ {
		n, e := c.Read(p)
		if e != nil {
			t.Fatal(e)
		}
		got = append(got, p[:n]...)
	}
	if string(got) != "abcdef" || c.u9AzvkFozu != nil {
		t.Fatal(got)
	}
	c = NewHTTPObfs(&memoryConn{reader: strings.NewReader("partial header")}, "x", "80")
	if n, e := c.Read(p); n != 0 || e != io.EOF || !c.iO_umx {
		t.Fatal(n, e)
	}
}
func TestTLSHelloStructure(t *testing.T) {
	p := []byte("secret payload")
	host := "example.test"
	before := time.Now().Unix()
	b := oRzazA8u(p, host)
	if len(b) != 217+len(p)+len(host) || int(binary.BigEndian.Uint16(b[3:5])) != len(b)-5 || int(binary.BigEndian.Uint16(b[7:9])) != len(b)-9 || !bytes.Equal(b[:3], []byte{22, 3, 1}) {
		t.Fatalf("length %d", len(b))
	}
	stamp := int64(binary.BigEndian.Uint32(b[11:15]))
	if stamp < before || stamp > time.Now().Unix() {
		t.Fatal(stamp)
	}
	if b[43] != 32 || binary.BigEndian.Uint16(b[76:78]) != 56 || !bytes.Equal(b[134:136], []byte{1, 0}) {
		t.Fatal("session/cipher")
	}
	ext := b[138:]
	seen := map[uint16][]byte{}
	for len(ext) > 0 {
		kind := binary.BigEndian.Uint16(ext)
		size := int(binary.BigEndian.Uint16(ext[2:]))
		seen[kind] = ext[4 : 4+size]
		ext = ext[4+size:]
	}
	if !bytes.Equal(seen[35], p) || string(seen[0][5:]) != host || !bytes.Equal(seen[11], []byte{3, 1, 0, 2}) || len(seen) != 7 {
		t.Fatal(seen)
	}
}
func TestTLSChunkWriteFailureAndZero(t *testing.T) {
	sentinel := errors.New("write")
	raw := &memoryConn{writeErr: sentinel, failAt: 2}
	c := NewTLSObfs(raw, "x")
	if n, e := c.Write(make([]byte, 17000)); n != 616 || e != sentinel {
		t.Fatal(n, e)
	}
	if c.xQi_DMa8h_ || len(raw.writes) != 2 || !bytes.Equal(raw.writes[1][:5], []byte{23, 3, 3, 2, 104}) {
		t.Fatal("frame")
	}
	raw = &memoryConn{}
	c = NewTLSObfs(raw, "x")
	if n, e := c.Write(nil); n != 0 || e != nil || len(raw.writes) != 0 || !c.xQi_DMa8h_ {
		t.Fatal(n, e)
	}
}
func TestTLSReadLengthsRemainderAndSwallowedError(t *testing.T) {
	wire := append(make([]byte, 105), 0, 6)
	wire = append(wire, []byte("abcdef")...)
	wire = append(wire, 23, 3, 3, 0, 3, 'x', 'y', 'z')
	c := NewTLSObfs(&memoryConn{reader: bytes.NewReader(wire)}, "x")
	p := make([]byte, 2)
	var got []byte
	for i := 0; i < 5; i++ {
		n, e := c.Read(p)
		if e != nil {
			t.Fatal(e)
		}
		got = append(got, p[:n]...)
	}
	if string(got) != "abcdefxyz" {
		t.Fatal(string(got))
	}
	c = NewTLSObfs(&memoryConn{reader: bytes.NewReader(make([]byte, 106))}, "x")
	if n, e := c.Read(p); n != 0 || e != nil || c.iwYJhx {
		t.Fatal(n, e)
	}
}
func TestObfsLayouts(t *testing.T) {
	if unsafe.Sizeof(WEp9oDv_ru{}) != 88 || unsafe.Sizeof(MSUT6Drr{}) != 48 {
		t.Fatal("RTTI")
	}
}
