package proxyproto

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"strings"
	"testing"
	"time"
	"unsafe"

	"recovered.local/soga/buf/pool"
)

func parseFixture(t *testing.T, data []byte) (*URdA97qPFc, error) {
	t.Helper()
	r := pool.C8FauP3Auy(bytes.NewReader(data), 4096)
	defer r.FreeBuf()
	return LavOMmUY2(r)
}
func TestOriginalProxyV1Parsing(t *testing.T) {
	for _, tc := range []struct {
		line, source string
		want         error
	}{
		{"PROXY TCP4 192.0.2.1 198.51.100.4 1234 443\r\n", "192.0.2.1:1234", nil},
		{"PROXY TCP6 ::1 ::2 1 2\r\n", "[::1]:1", nil},
		{"PROXY TCP4 ::ffff:192.0.2.1 198.51.100.4 1 2\r\n", "", ErrorAt5097ED0},
		{"PROXY TCP6 ::ffff:192.0.2.1 ::2 1 2\r\n", "192.0.2.1:1", nil},
		{"PROXY TCP4 1.2.3.4 2.3.4.5 -1 2\r\n", "", ErrorAt5097EE0},
		{"PROXY TCP4 1.2.3.4 2.3.4.5 nope 2\r\n", "", ErrorAt5097EE0},
		{"PROXY TCP4 1.2.3.4 2.3.4.5 65536 2\r\n", "", ErrorAt5097EE0},
		{"PROXY TCP4 1.2.3.4 2.3.4.5 1 2\n", "", ErrorAt5097E20},
		{"PROXY  TCP4 1.2.3.4 2.3.4.5 1 2\r\n", "", ErrorAt5097E40},
	} {
		h, err := parseFixture(t, []byte(tc.line))
		if err != tc.want {
			t.Fatalf("%q: %v", tc.line, err)
		}
		if err == nil && h.Dafxa_BOdZdw.String() != tc.source {
			t.Fatal(h.Dafxa_BOdZdw)
		}
	}
	h, err := parseFixture(t, []byte("PROXY UNKNOWN anything\r\n"))
	if err != nil || h.LgI8Cr1feIYb != 0x20 || h.Dafxa_BOdZdw != nil {
		t.Fatal(h, err)
	}
	for _, data := range []string{"hello", "PRO", "\r\n"} {
		if _, err := parseFixture(t, []byte(data)); err != ErrorAt5097E80 {
			t.Fatal(data, err)
		}
	}
	if _, err := parseFixture(t, []byte("PROXY "+strings.Repeat("x", 102))); err != ErrorAt5097E10 {
		t.Fatal(err)
	}
}

func v2Frame(command, protocol byte, payload []byte) []byte {
	b := append([]byte{}, signatureAt5099D50...)
	b = append(b, command, protocol, byte(len(payload)>>8), byte(len(payload)))
	return append(b, payload...)
}
func TestOriginalProxyV2FamiliesAndBounds(t *testing.T) {
	for _, p := range []byte{0x11, 0x12, 0x21, 0x22, 0x31, 0x32} {
		var payload []byte
		switch p & 0xf0 {
		case 0x10:
			payload = []byte{192, 0, 2, 1, 198, 51, 100, 1, 4, 210, 1, 187}
		case 0x20:
			payload = make([]byte, 36)
			payload[15] = 1
			payload[31] = 2
			binary.BigEndian.PutUint16(payload[32:], 1234)
			binary.BigEndian.PutUint16(payload[34:], 443)
		case 0x30:
			payload = make([]byte, 216)
			copy(payload, "/tmp/src")
			copy(payload[108:], "/tmp/dst")
		}
		payload = append(payload, 1, 0, 1, 42)
		h, err := parseFixture(t, v2Frame(0x21, p, payload))
		if err != nil {
			t.Fatalf("%x %v", p, err)
		}
		if h.Dafxa_BOdZdw == nil || h.KhtRcTD89 == nil || !bytes.Equal(h.ttF0mSAdaYG, []byte{1, 0, 1, 42}) {
			t.Fatal(h)
		}
	}
	if _, err := parseFixture(t, v2Frame(0x21, 0, nil)); err != ErrorAt5097EB0 {
		t.Fatal(err)
	}
	if h, err := parseFixture(t, v2Frame(0x20, 0, nil)); err != nil || h.ttF0mSAdaYG != nil {
		t.Fatal(h, err)
	}
	if _, err := parseFixture(t, v2Frame(0x22, 0, nil)); err != ErrorAt5097EA0 {
		t.Fatal(err)
	}
	if _, err := parseFixture(t, v2Frame(0x21, 0x11, make([]byte, 11))); err != ErrorAt5097EC0 {
		t.Fatal(err)
	}
	if _, err := parseFixture(t, v2Frame(0x20, 0, make([]byte, 4097))); err != ErrorAt5097EC0 {
		t.Fatal(err)
	}
	// 原解析允许未定义的传输低半字节；IPv4地址读取后两地址均为nil。
	if h, err := parseFixture(t, v2Frame(0x21, 0x13, make([]byte, 12))); err != nil || h.Dafxa_BOdZdw != nil || h.KhtRcTD89 != nil {
		t.Fatal(h, err)
	}
	if mLIKJLb([]byte{0, 'x'}) != "" {
		t.Fatal("original does not preserve abstract Unix names")
	}
}

type proxyMemoryConn struct {
	*bytes.Reader
	writes        bytes.Buffer
	closed        int
	deadlines     []time.Time
	deadlineError error
}

func (c *proxyMemoryConn) Write(p []byte) (int, error) { return c.writes.Write(p) }
func (c *proxyMemoryConn) Close() error                { c.closed++; return nil }
func (c *proxyMemoryConn) LocalAddr() net.Addr {
	return &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 443}
}
func (c *proxyMemoryConn) RemoteAddr() net.Addr {
	return &net.TCPAddr{IP: net.IPv4(127, 0, 0, 2), Port: 444}
}
func (c *proxyMemoryConn) SetDeadline(t time.Time) error { return c.SetReadDeadline(t) }
func (c *proxyMemoryConn) SetReadDeadline(t time.Time) error {
	c.deadlines = append(c.deadlines, t)
	return c.deadlineError
}
func (c *proxyMemoryConn) SetWriteDeadline(time.Time) error { return nil }

func TestOriginalProxyConnectionPoliciesAndBuffer(t *testing.T) {
	for _, policy := range []XTAkygU{0, 1, 2, 3, 4} {
		base := &proxyMemoryConn{Reader: bytes.NewReader([]byte("PROXY TCP4 192.0.2.5 198.51.100.1 123 443\r\npayload"))}
		c := GXOJV4B0a4D(base, WTqa9yT(policy))
		defer c.Free()
		if unsafe.Sizeof(*c) != 104 || unsafe.Offsetof(c.y7HtLBfne) != 48 || unsafe.Offsetof(c.aVUuZorW_) != 72 {
			t.Fatal("layout")
		}
		buf := make([]byte, 100)
		n, err := c.Read(buf)
		if policy == 2 {
			if err != ErrorAt5097EF0 || n != 0 || base.closed != 0 {
				t.Fatal("reject does not close", n, err, base.closed)
			}
			continue
		}
		if err != nil || string(buf[:n]) != "payload" {
			t.Fatal(n, err, string(buf[:n]))
		}
		if c.aVUuZorW_ == nil {
			t.Fatal("original releases buffer on following Read")
		}
		if n, err := c.Read(buf); n != 0 || err != io.EOF || c.aVUuZorW_ != nil {
			t.Fatal(n, err)
		}
		want := "127.0.0.2:444"
		if policy == 0 || policy == 3 {
			want = "192.0.2.5:123"
		}
		if c.RemoteAddr().String() != want || c.LocalAddr().String() != "127.0.0.1:443" {
			t.Fatal(c.RemoteAddr())
		}
	}
	base := &proxyMemoryConn{Reader: bytes.NewReader([]byte("hello"))}
	c := GXOJV4B0a4D(base, WTqa9yT(3))
	defer c.Free()
	if _, err := c.Read(make([]byte, 8)); err != ErrorAt5097E80 {
		t.Fatal(err)
	}
	base = &proxyMemoryConn{Reader: bytes.NewReader([]byte("PROXY UNKNOWN\r\n"))}
	sentinel := errors.New("validation")
	c = GXOJV4B0a4D(base, O8oMdhQ(func(*URdA97qPFc) error { return sentinel }))
	defer c.Free()
	if _, err := c.Read(nil); err != sentinel || c.ProxyHeader() != nil {
		t.Fatal(err)
	}
}

func TestProxyDeadlineRestorationAndTruePipe(t *testing.T) {
	base := &proxyMemoryConn{Reader: bytes.NewReader([]byte("plain"))}
	c := GXOJV4B0a4D(base)
	defer c.Free()
	old := time.Now().Add(time.Minute)
	c.SetReadDeadline(old)
	c.gQaLqCURuN1 = time.Second
	if _, err := c.Read(make([]byte, 5)); err != nil {
		t.Fatal(err)
	}
	if len(base.deadlines) != 3 || base.deadlines[2] != old {
		t.Fatal(base.deadlines)
	}
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	c = GXOJV4B0a4D(a)
	defer c.Free()
	c.SetReadDeadline(time.Now().Add(time.Second))
	done := make(chan error, 1)
	go func() { _, err := b.Write([]byte("PROXY TCP4 192.0.2.9 198.51.100.1 123 443\r\nhello")); done <- err }()
	buf := make([]byte, 5)
	if _, err := io.ReadFull(c, buf); err != nil || string(buf) != "hello" {
		t.Fatal(err, string(buf))
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	if c.RemoteAddr().String() != "192.0.2.9:123" {
		t.Fatal(c.RemoteAddr())
	}
}
