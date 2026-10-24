package GtTssd5z

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
	"net"
	"strings"
	"testing"
	"time"
	"unsafe"
)

type wireConn struct {
	input  *bytes.Reader
	output bytes.Buffer
	err    error
	reads  int
}

func (c *wireConn) Read(p []byte) (int, error) {
	c.reads++
	if c.err != nil {
		return 0, c.err
	}
	if c.input == nil {
		return 0, io.EOF
	}
	return c.input.Read(p)
}
func (c *wireConn) Write(p []byte) (int, error) {
	if c.err != nil {
		return 0, c.err
	}
	return c.output.Write(p)
}
func (*wireConn) Close() error                     { return nil }
func (*wireConn) LocalAddr() net.Addr              { return nil }
func (*wireConn) RemoteAddr() net.Addr             { return nil }
func (*wireConn) SetDeadline(time.Time) error      { return nil }
func (*wireConn) SetReadDeadline(time.Time) error  { return nil }
func (*wireConn) SetWriteDeadline(time.Time) error { return nil }

func TestLayoutsAndFactory(t *testing.T) {
	for _, x := range []struct {
		n         string
		got, want uintptr
	}{{"base", unsafe.Sizeof(Ja7FBino7HL{}), 72}, {"http", unsafe.Sizeof(vkmx7_{}), 56}, {"random", unsafe.Sizeof(nOAagf7gG3W{}), 56}, {"tls", unsafe.Sizeof(cFCr3rhcmDX{}), 160}} {
		if x.got != x.want {
			t.Fatalf("%s %d", x.n, x.got)
		}
	}
	for _, name := range []string{"plain", "http_simple", "http_post", "random_head", "tls1.2_ticket_auth", "tls1.2_ticket_fastauth"} {
		o, n, e := IRT7O1m(name, &Ja7FBino7HL{})
		if e != nil || o == nil {
			t.Fatal(name, e)
		}
		want := 0
		if strings.HasPrefix(name, "tls") {
			want = 5
		}
		if n != want {
			t.Fatal(name, n)
		}
	}
	if o, _, e := IRT7O1m("PLAIN", &Ja7FBino7HL{}); o != nil || e == nil || e.Error() != "Obfs PLAIN not supported" {
		t.Fatal(o, e)
	}
	if len(userAgentsAt50A6328) != 250 {
		t.Fatal(len(userAgentsAt50A6328))
	}
}
func TestHTTPWireAndErrorBehavior(t *testing.T) {
	for _, name := range []string{"http_simple", "http_post"} {
		raw := &wireConn{}
		o, _, _ := IRT7O1m(name, &Ja7FBino7HL{O1k0EzteK: "example.test", P2XQpgsRo9: 8080, NhWlHt412H8S: "selected.test#X-One: 1\\nX-Two: 2"})
		c := o.StreamConn(raw)
		if n, e := c.Write([]byte{0, 0xff, 0x2f}); n != 3 || e != nil {
			t.Fatal(n, e)
		}
		method := "GET"
		if name == "http_post" {
			method = "POST"
		}
		want := method + " /%00%ff%2f HTTP/1.1\r\nHost: selected.test:8080\r\nX-One: 1\r\nX-Two: 2\r\n\r\n"
		if raw.output.String() != want {
			t.Fatalf("%q", raw.output.String())
		}
		c.Write([]byte("after"))
		if !strings.HasSuffix(raw.output.String(), "after") {
			t.Fatal("later write")
		}
		raw.input = bytes.NewReader([]byte("HTTP/1.1 200 OK\r\n\r\nabcdef"))
		b := make([]byte, 2)
		for _, want := range []string{"ab", "cd", "ef"} {
			n, e := c.Read(b)
			if n != 2 || e != nil || string(b) != want {
				t.Fatal(n, e, string(b))
			}
		}
		if raw.reads != 1 {
			t.Fatal(raw.reads)
		}
	}
	sentinel := errors.New("write failed")
	raw := &wireConn{err: sentinel}
	o, _, _ := IRT7O1m("http_simple", &Ja7FBino7HL{})
	c := o.StreamConn(raw).(*vkmx7_)
	if n, e := c.Write([]byte("x")); n != 0 || e != nil || c.smRoaiq_uhA6 {
		t.Fatal(n, e, c.smRoaiq_uhA6)
	}
}
func TestRandomHeadCRCAndBufferedTransition(t *testing.T) {
	raw := &wireConn{}
	o, _, _ := IRT7O1m("random_head", &Ja7FBino7HL{})
	c := o.StreamConn(raw).(*nOAagf7gG3W)
	if n, e := c.Write([]byte("first")); n != 5 || e != nil {
		t.Fatal(n, e)
	}
	header := append([]byte(nil), raw.output.Bytes()...)
	if len(header) < 8 || len(header) > 103 {
		t.Fatal(len(header))
	}
	if crc32.ChecksumIEEE(header) != 0xffffffff {
		t.Fatal("CRC residue")
	}
	c.Write([]byte("second"))
	if raw.output.Len() != len(header) {
		t.Fatal("premature flush")
	}
	n, e := c.Read(nil)
	if n != 0 || e != nil || !c.fku_KtEBE || !c.bYxvLiDBdWa {
		t.Fatal(n, e)
	}
	if string(raw.output.Bytes()[len(header):]) != "firstsecond" {
		t.Fatal(raw.output.String())
	}
	c.Write([]byte("third"))
	if !strings.HasSuffix(raw.output.String(), "third") {
		t.Fatal("passthrough")
	}
}
func TestTLSTicketHandshakeAndRecordRead(t *testing.T) {
	raw := &wireConn{}
	base := &Ja7FBino7HL{O1k0EzteK: "example.test", JtkCoK71gtjm: []byte("shared key")}
	o, _, _ := IRT7O1m("tls1.2_ticket_auth", base)
	c := o.StreamConn(raw).(*cFCr3rhcmDX)
	if n, e := c.Write([]byte("pending")); n != 7 || e != nil || c.vpScqu != 1 {
		t.Fatal(n, e)
	}
	hello := raw.output.Bytes()
	if len(hello) < 76 || !bytes.Equal(hello[:3], []byte{22, 3, 1}) || int(binary.BigEndian.Uint16(hello[3:5])) != len(hello)-5 {
		t.Fatalf("bad hello %x", hello)
	}
	hmac10 := func(p []byte) []byte {
		key := append(append([]byte(nil), base.JtkCoK71gtjm...), c.l57vr8H[:]...)
		h := hmac.New(sha1.New, key)
		h.Write(p)
		return h.Sum(nil)[:10]
	}
	if !hmac.Equal(hello[33:43], hmac10(hello[11:33])) {
		t.Fatal("client authentication")
	}
	if !bytes.Contains(hello, []byte("example.test")) {
		t.Fatal("SNI")
	}
	response := make([]byte, 76)
	copy(response[33:43], hmac10(response[11:33]))
	copy(response[66:], hmac10(response[:66]))
	raw.input = bytes.NewReader(response)
	before := raw.output.Len()
	if n, e := c.Read(nil); n != 0 || e != nil || c.vpScqu != 8 {
		t.Fatal(n, e)
	}
	tail := raw.output.Bytes()[before:]
	if !bytes.Equal(tail[:11], []byte{20, 3, 3, 0, 1, 1, 22, 3, 3, 0, 32}) || !hmac.Equal(tail[33:43], hmac10(tail[:33])) {
		t.Fatal("finish handshake")
	}
	if !bytes.Equal(tail[43:], []byte{23, 3, 3, 0, 7, 'p', 'e', 'n', 'd', 'i', 'n', 'g'}) {
		t.Fatalf("pending %x", tail[43:])
	}
	raw.input = bytes.NewReader([]byte{23, 3, 3, 0, 3, 'a', 'b', 'c', 23, 3, 3, 0, 2, 'd', 'e'})
	p := make([]byte, 3)
	if n, e := c.Read(p); n != 3 || e != nil || string(p) != "abc" {
		t.Fatal(n, e, string(p))
	}
	if n, e := c.Read(p); n != 2 || e != nil || string(p[:n]) != "de" {
		t.Fatal(n, e, string(p[:n]))
	}
	raw.input = bytes.NewReader([]byte{22, 3, 3, 0, 1, 0})
	if _, e := c.Read(p); e != errorAt5098560 {
		t.Fatal(e)
	}
}
