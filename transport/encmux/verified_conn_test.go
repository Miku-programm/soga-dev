package encmux

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"testing"
	"time"
	"unsafe"
)

type wireFixture struct {
	*bytes.Reader
	out        bytes.Buffer
	max        int
	readErr    error
	zero       bool
	shortWrite bool
	writeErr   error
	freed      bool
}

func (c *wireFixture) Read(p []byte) (int, error) {
	if c.zero {
		c.zero = false
		return 0, nil
	}
	if c.max > 0 && len(p) > c.max {
		p = p[:c.max]
	}
	n, e := c.Reader.Read(p)
	if c.readErr != nil {
		return n, c.readErr
	}
	return n, e
}
func (c *wireFixture) Write(p []byte) (int, error) {
	if c.shortWrite {
		return 1, c.writeErr
	}
	return c.out.Write(p)
}
func (c *wireFixture) Close() error                     { return nil }
func (c *wireFixture) LocalAddr() net.Addr              { return &net.TCPAddr{Port: 1} }
func (c *wireFixture) RemoteAddr() net.Addr             { return &net.TCPAddr{Port: 2} }
func (c *wireFixture) SetDeadline(time.Time) error      { return nil }
func (c *wireFixture) SetReadDeadline(time.Time) error  { return nil }
func (c *wireFixture) SetWriteDeadline(time.Time) error { return nil }
func (c *wireFixture) Free()                            { c.freed = true }
func makeVerifiedFixture(data []byte, proxy bool) (*Ql1qrX, *wireFixture) {
	raw := &wireFixture{Reader: bytes.NewReader(data)}
	return &Ql1qrX{DrpXao: raw, fR9Xy3: &VTl94pzz2IkB{hrXHbLocdr: proxy}}, raw
}
func wireToken(cmd ...byte) []byte { return append(make([]byte, 16), cmd...) }
func TestTunnelOriginalTypeLayout(t *testing.T) {
	var c Ql1qrX
	var v VTl94pzz2IkB
	var l Ppx7OU
	if unsafe.Sizeof(c) != 56 || unsafe.Offsetof(c.fR9Xy3) != 16 || unsafe.Offsetof(c.cv7owVg) != 24 || unsafe.Offsetof(c.quUa_I) != 32 || unsafe.Offsetof(c.mPP4038Qh) != 40 || unsafe.Offsetof(c.vfKdlLJ) != 41 || unsafe.Offsetof(c.z5K6vhFUp) != 48 || unsafe.Sizeof(v) != 144 || unsafe.Offsetof(v.pyauaz) != 8 || unsafe.Offsetof(v.p1Wtyln2GJ6) != 16 || unsafe.Offsetof(v.psdXJ2Mi) != 32 || unsafe.Offsetof(v.pJLNnm) != 40 || unsafe.Offsetof(v._rKF9rrYaf8a) != 56 || unsafe.Offsetof(v.lyXfKPaz4m23) != 64 || unsafe.Offsetof(v.jrF7WDTK) != 80 || unsafe.Offsetof(v.hrXHbLocdr) != 81 || unsafe.Offsetof(v.dlAoUM) != 88 || unsafe.Offsetof(v.zkn35IN) != 96 || unsafe.Offsetof(v.oWvGY4I9cA) != 104 || unsafe.Offsetof(v.x5Lor8hDyn) != 128 || unsafe.Sizeof(l) != 24 {
		t.Fatal("original RTTI mismatch")
	}
}
func TestTunnelAuthenticationStrictSingleReadAndCommands(t *testing.T) {
	for _, tc := range []struct {
		name  string
		data  []byte
		limit int
		err   error
		want  string
	}{{"short", wireToken(1), 8, nil, "tunnel verify failed"}, {"mismatch", append([]byte{1}, make([]byte, 16)...), 0, nil, "tunnel verify failed"}, {"eof", nil, 0, nil, "EOF"}, {"data_and_error", wireToken(1), 0, io.EOF, "EOF"}, {"bad_command", wireToken(9), 0, nil, "tunnel wrong cmd: 9"}, {"proxy_disabled", wireToken(3), 0, nil, "not proxy mode"}} {
		t.Run(tc.name, func(t *testing.T) {
			c, raw := makeVerifiedFixture(tc.data, false)
			raw.max = tc.limit
			raw.readErr = tc.err
			if _, e := c.GetProxyType(); e == nil || e.Error() != tc.want || c.cv7owVg {
				t.Fatal(e, c.cv7owVg)
			}
		})
	}
	c, raw := makeVerifiedFixture(wireToken(1, 'x'), false)
	raw.zero = true
	if kind, e := c.GetProxyType(); e != nil || kind != 1 || !c.cv7owVg {
		t.Fatal(kind, e)
	}
	var one [1]byte
	if n, e := c.Read(one[:]); e != nil || n != 1 || one[0] != 'x' {
		t.Fatal(n, e, one)
	}
	if yes, e := c.IsProxy(); yes || e != nil {
		t.Fatal(yes, e)
	}
	c.Free()
	if !raw.freed || c.GetRawConn() != raw {
		t.Fatal("raw/free delegation")
	}
}
func TestTunnelProxyAddressAndPartialState(t *testing.T) {
	data := wireToken(3, 4, 3, 4, 'h', 'o', 's', 't', 1, 187)
	c, _ := makeVerifiedFixture(data, true)
	if kind, e := c.GetProxyType(); e != nil || kind != 4 {
		t.Fatal(kind, e)
	}
	a := c.GetProxyAddr()
	if a.GetHost() != "host" || a.GetPort() != 443 {
		t.Fatal(a)
	}
	if yes, e := c.IsProxy(); !yes || e != nil {
		t.Fatal(yes, e)
	}
	c, _ = makeVerifiedFixture(wireToken(3, 7), true)
	if _, e := c.GetProxyType(); e == nil || e.Error() != "tunnel wrong cmd: 7" || !c.vfKdlLJ || c.cv7owVg {
		t.Fatal("partial proxy flag", e)
	}
	c, _ = makeVerifiedFixture(wireToken(3, 2, 1, 127), true)
	if _, e := c.GetProxyType(); !errors.Is(e, io.ErrUnexpectedEOF) || c.mPP4038Qh != 2 || !c.vfKdlLJ || c.cv7owVg {
		t.Fatal("partial type state", e)
	}
}
func TestTunnelDatagramDrainWriteAndErrors(t *testing.T) {
	c, raw := makeVerifiedFixture(append(wireToken(2), []byte{0, 5, 'a', 'b', 'c', 'd', 'e', 0, 2, 'x', 'y'}...), false)
	p := make([]byte, 2)
	if n, e := c.Read(p); n != 2 || e != nil || string(p) != "ab" {
		t.Fatal(n, e, string(p))
	}
	if n, a, e := c.ReadFrom(p); n != 2 || e != nil || string(p) != "xy" || a.String() != raw.RemoteAddr().String() {
		t.Fatal(n, a, e, string(p))
	}
	if n, e := c.Write([]byte("abc")); n != 3 || e != nil || !bytes.Equal(raw.out.Bytes(), []byte{0, 3, 'a', 'b', 'c'}) {
		t.Fatal(n, e, raw.out.Bytes())
	}
	raw.out.Reset()
	large := make([]byte, 65536)
	if n, e := c.WriteTo(large, nil); n != len(large) || e != nil || raw.out.Len() != len(large)+2 || binary.BigEndian.Uint16(raw.out.Bytes()) != 0 {
		t.Fatal("wrapped length", n, e)
	}
	raw.shortWrite = true
	if n, e := c.Write([]byte("payload")); n != 7 || e != nil {
		t.Fatal("short nil write", n, e)
	}
	raw.writeErr = io.ErrClosedPipe
	if n, e := c.Write([]byte("payload")); n != 0 || e != io.ErrClosedPipe {
		t.Fatal(n, e)
	}
	c, _ = makeVerifiedFixture(append(wireToken(2), 0, 5, 'a'), false)
	if n, a, e := c.ReadFrom(p); n != 0 || a != nil || e == nil {
		t.Fatal("ReadFrom discards partial count", n, a, e)
	}
	c, _ = makeVerifiedFixture(append(wireToken(2), 0, 5, 'a'), false)
	if n, e := c.Read(p); n != 1 || e == nil {
		t.Fatal("Read retains partial count", n, e)
	}
}
func TestTunnelListenerActualTCPAuthentication(t *testing.T) {
	listener, e := net.Listen("tcp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer listener.Close()
	config := &VTl94pzz2IkB{}
	copy(config.x5Lor8hDyn[:], []byte("0123456789abcdef"))
	wrapped := &Ppx7OU{O6B2qZ0p0: listener, tS5gsM72s: config}
	done := make(chan error, 1)
	go func() {
		conn, e := wrapped.Accept()
		if e != nil {
			done <- e
			return
		}
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		q := conn.(*Ql1qrX)
		if q.cv7owVg {
			done <- errors.New("eager auth")
			return
		}
		p := make([]byte, 4)
		_, e = io.ReadFull(q, p)
		if e == nil {
			_, e = q.Write(p)
		}
		done <- e
	}()
	client, e := net.Dial("tcp", listener.Addr().String())
	if e != nil {
		t.Fatal(e)
	}
	defer client.Close()
	client.SetDeadline(time.Now().Add(5 * time.Second))
	if _, e = client.Write(append(append(config.x5Lor8hDyn[:], 1), 'p', 'i', 'n', 'g')); e != nil {
		t.Fatal(e)
	}
	var reply [4]byte
	if _, e = io.ReadFull(client, reply[:]); e != nil || string(reply[:]) != "ping" {
		t.Fatal(e, string(reply[:]))
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
}
