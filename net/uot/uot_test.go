package uot

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	D "recovered.local/soga/buf/pool"
	G "recovered.local/soga/net/relay"
	"reflect"
	"strings"
	"testing"
	"time"
)

type uotMemoryConn struct {
	net.Conn
	*bytes.Reader
	output    []byte
	deadlines []time.Time
	short     bool
}

func (c *uotMemoryConn) Read(p []byte) (int, error) { return c.Reader.Read(p) }
func (c *uotMemoryConn) Write(p []byte) (int, error) {
	c.output = append(c.output, p...)
	if c.short {
		return 0, nil
	}
	return len(p), nil
}
func (c *uotMemoryConn) SetReadDeadline(d time.Time) error {
	c.deadlines = append(c.deadlines, d)
	return errors.New("ignored deadline error")
}
func frame(addr []byte, data string) []byte {
	b := append([]byte{}, addr...)
	b = binary.BigEndian.AppendUint16(b, uint16(len(data)))
	return append(b, data...)
}

var socksIP = []byte{1, 192, 0, 2, 1, 0, 53}
var uotIP = []byte{0, 192, 0, 2, 1, 0, 53}
var uotOther = []byte{0, 198, 51, 100, 2, 0, 54}

func TestOriginalUOTLayoutAndMagic(t *testing.T) {
	typ := reflect.TypeFor[JAYwg0sLs]()
	if typ.Size() != 64 {
		t.Fatal(typ)
	}
	for i, o := range []uintptr{0, 16, 32, 40, 48, 56} {
		if typ.Field(i).Offset != o {
			t.Fatal(typ.Field(i))
		}
	}
	if reflect.TypeFor[MYA15vm4nVrn]().Size() != 16 {
		t.Fatal("request layout")
	}
	for _, s := range []string{"sp.v2.udp-over-tcp.arpa", "sp.udp-over-tcp.arpa"} {
		if !Vd_ooUdqfJTg(s) {
			t.Fatal(s)
		}
	}
	for _, s := range []string{"SP.V2.UDP-OVER-TCP.ARPA", "sp.udp-over-tcp.arpa.", " sp.udp-over-tcp.arpa"} {
		if Vd_ooUdqfJTg(s) {
			t.Fatal(s)
		}
	}
	c := &uotMemoryConn{Reader: bytes.NewReader(nil)}
	if _, e := SACWxFiDLx22(c, "unknown", nil, nil); e == nil || e.Error() != "unknown uot magic address: unknown" {
		t.Fatal(e)
	}
	if len(c.deadlines) != 2 || !c.deadlines[1].IsZero() {
		t.Fatal(c.deadlines)
	}
}
func TestOriginalUOTVersionsFrameAlignmentAndTruncation(t *testing.T) {
	for _, mode := range []string{"v1", "v2-connect", "v2-associate"} {
		t.Run(mode, func(t *testing.T) {
			magic := "sp.v2.udp-over-tcp.arpa"
			var wire []byte
			switch mode {
			case "v1":
				magic = "sp.udp-over-tcp.arpa"
				wire = frame(uotIP, "abcdef")
				wire = append(wire, frame(uotOther, "next")...)
			case "v2-connect":
				wire = append([]byte{1}, socksIP...)
				wire = append(wire, frame(nil, "abcdef")...)
				wire = append(wire, frame(nil, "next")...)
			case "v2-associate":
				wire = append([]byte{0}, socksIP...)
				wire = append(wire, frame(uotOther, "abcdef")...)
				wire = append(wire, frame(uotIP, "next")...)
			}
			raw := &uotMemoryConn{Reader: bytes.NewReader(wire)}
			conn, e := SACWxFiDLx22(raw, magic, nil, nil)
			if e != nil {
				t.Fatal(e)
			}
			c := conn.(*JAYwg0sLs)
			if c.GetDestAddr().String() != "192.0.2.1:53" || c.GetUser() != nil || c.NetworkType() != 1 {
				t.Fatal(c.GetDestAddr())
			}
			p := make([]byte, 3)
			n, addr, e := c.ReadFromProxy(p)
			if e != nil || n != 3 || string(p) != "abc" {
				t.Fatal(n, addr, e, string(p))
			}
			want := "192.0.2.1:53"
			if mode == "v2-associate" {
				want = "198.51.100.2:54"
			}
			if addr.String() != want {
				t.Fatal(addr)
			}
			p = make([]byte, 8)
			n, addr, e = c.ReadFromProxy(p)
			if e != nil || string(p[:n]) != "next" {
				t.Fatal(n, addr, e)
			}
			want = "198.51.100.2:54"
			if mode != "v1" {
				want = "192.0.2.1:53"
			}
			if addr.String() != want {
				t.Fatal(addr)
			}
			n, e = c.WriteToProxy([]byte("reply"), &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 53})
			if e != nil || n != 5 {
				t.Fatal(n, e)
			}
			address := uotIP
			if mode == "v2-connect" {
				address = nil
			}
			if !bytes.Equal(raw.output, frame(address, "reply")) {
				t.Fatalf("%x", raw.output)
			}
		})
	}
}
func TestOriginalUOTReadErrorsAndWriteQuirks(t *testing.T) {
	c := &JAYwg0sLs{DrpXao: &uotMemoryConn{Reader: bytes.NewReader([]byte{0, 5, 'a', 'b'})}, eStVCYS: true, bPg7GR8DCGh: mustAddr(t)}
	p := make([]byte, 2)
	n, a, e := c.ReadFromProxy(p)
	if e != nil || n != 2 || a == nil || string(p) != "ab" {
		t.Fatal(n, a, e)
	} // drain EOF is discarded
	c.DrpXao = &uotMemoryConn{Reader: bytes.NewReader([]byte{0, 5, 'a'})}
	n, a, e = c.ReadFromProxy(p)
	if n != 0 || a != nil || e != io.ErrUnexpectedEOF {
		t.Fatal(n, a, e)
	}
	raw := &uotMemoryConn{Reader: bytes.NewReader(nil), short: true}
	c.DrpXao = raw
	payload := bytes.Repeat([]byte{7}, 65536)
	n, e = c.WriteToProxy(payload, nil)
	if n != len(payload) || e != nil || binary.BigEndian.Uint16(raw.output[:2]) != 0 || len(raw.output) != 65538 {
		t.Fatal(n, e, len(raw.output))
	}
	raw = &uotMemoryConn{Reader: bytes.NewReader([]byte{1})}
	_, e = SACWxFiDLx22(raw, "sp.v2.udp-over-tcp.arpa", nil, nil)
	if e == nil || !strings.HasPrefix(e.Error(), "error reading uot request: ") || errors.Is(e, io.EOF) {
		t.Fatal(e)
	}
	r, e := Lfd9DDhfr(bytes.NewReader(append([]byte{2}, socksIP...)))
	if e != nil || !r.Hja8Qj {
		t.Fatal(r, e)
	} // binary.Read accepts any nonzero bool
}
func mustAddr(t *testing.T) *G.AHU2MRmjzQu {
	t.Helper()
	a, e := G.U80FDt("192.0.2.1:53")
	if e != nil {
		t.Fatal(e)
	}
	return a
}

// Actual stream IO verifies framing without a replacement UDP-over-TCP stack.
func TestOriginalUOTActualTCP(t *testing.T) {
	ln, e := net.Listen("tcp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer ln.Close()
	done := make(chan error, 1)
	go func() {
		c, e := ln.Accept()
		if e != nil {
			done <- e
			return
		}
		defer c.Close()
		raw := &uotRealConn{Conn: c}
		wrapped, e := SACWxFiDLx22(raw, "sp.v2.udp-over-tcp.arpa", nil, nil)
		if e != nil {
			done <- e
			return
		}
		x := wrapped.(*JAYwg0sLs)
		buf := make([]byte, 8)
		n, a, e := x.ReadFromProxy(buf)
		if e == nil {
			_, e = x.WriteToProxy(buf[:n], a)
		}
		done <- e
	}()
	c, e := net.Dial("tcp4", ln.Addr().String())
	if e != nil {
		t.Fatal(e)
	}
	defer c.Close()
	_ = c.SetDeadline(time.Now().Add(3 * time.Second))
	wire := append([]byte{1}, socksIP...)
	wire = append(wire, frame(nil, "wire")...)
	if _, e = c.Write(wire); e != nil {
		t.Fatal(e)
	}
	got := make([]byte, 6)
	if _, e = io.ReadFull(c, got); e != nil || !bytes.Equal(got, frame(nil, "wire")) {
		t.Fatal(e, got)
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
}

type uotRealConn struct{ net.Conn }

func (c *uotRealConn) SetReadDeadline(d time.Time) error {
	if !d.IsZero() {
		d = time.Now().Add(3 * time.Second)
	}
	return c.Conn.SetReadDeadline(d)
}

func TestOriginalUOTAddressTagsAndErrors(t *testing.T) {
	for _, host := range []string{"192.0.2.1:53", "[2001:db8::1]:53", "domain.test:53"} {
		a, e := G.U80FDt(host)
		if e != nil {
			t.Fatal(e)
		}
		b := D.S4OQXWIZ(300)
		_, e = G.VCQ8zrwYebS(b, a)
		if e != nil {
			t.Fatal(e)
		}
		out, e := G.O2xtlb(bytes.NewReader(b.Bytes()))
		if e != nil || out.GetHost() != a.GetHost() || out.GetPort() != a.GetPort() {
			t.Fatal(out, e)
		}
		b.Free()
	}
	if _, e := G.O2xtlb(bytes.NewReader([]byte{3})); e == nil || e.Error() != "unknown addr type: 3" {
		t.Fatal(e)
	}
	b := D.S4OQXWIZ(300)
	defer b.Free()
	if _, e := G.VCQ8zrwYebS(b, G.NewDomainAddress(strings.Repeat("x", 256), 53)); e == nil || b.Len() != 0 {
		t.Fatal(e, b.Len())
	}
}
