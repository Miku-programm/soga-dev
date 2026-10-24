package EP3M1uC

import (
	"context"
	"errors"
	"io"
	"net"
	"recovered.local/soga/E3F12urwswL"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/vP_e5OQAve"
	"testing"
	"time"
)

func TestDirectAndRedirectTCPActualConnection(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	target, err := G6bHD5JA.RNFJeaRY(listener.Addr())
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 2)
	go func() {
		for i := 0; i < 2; i++ {
			c, e := listener.Accept()
			if e != nil {
				done <- e
				return
			}
			c.SetDeadline(time.Now().Add(2 * time.Second))
			_, e = io.CopyN(c, c, 4)
			c.Close()
			done <- e
		}
	}()
	for _, kind := range []Fdaw35mqIz{"direct", "redirect"} {
		out := &GyyO51{NyLbiXb4OrlL: kind, AHTlIEynXL: "127.0.0.1", iCySnjtd: net.IPv4(127, 0, 0, 1)}
		// Redirect's zero configured port must fall back to the target port.
		c, e := out.DialTCP(nil, target)
		if e != nil {
			t.Fatal(e)
		}
		c.SetDeadline(time.Now().Add(2 * time.Second))
		if _, e = c.Write([]byte("ping")); e != nil {
			t.Fatal(e)
		}
		var b [4]byte
		_, e = io.ReadFull(c, b[:])
		c.Close()
		if e != nil || string(b[:]) != "ping" {
			t.Fatalf("%s: %q %v", kind, b, e)
		}
		if e = <-done; e != nil {
			t.Fatal(e)
		}
	}
}
func TestDirectAndRedirectUDPActualPacket(t *testing.T) {
	server, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	server.SetDeadline(time.Now().Add(3 * time.Second))
	target, err := G6bHD5JA.RNFJeaRY(server.LocalAddr())
	if err != nil {
		t.Fatal(err)
	}
	for _, kind := range []Fdaw35mqIz{"direct", "redirect"} {
		out := &GyyO51{NyLbiXb4OrlL: kind, AHTlIEynXL: "127.0.0.1", JO1YIpXk: target.GetPort(), iCySnjtd: net.IPv4(127, 0, 0, 1)}
		c, e := out.DialUDP(nil, target)
		if e != nil {
			t.Fatal(e)
		}
		c.SetDeadline(time.Now().Add(2 * time.Second))
		destination := server.LocalAddr()
		if kind == "redirect" {
			destination = &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1}
		}
		if _, e = c.WriteTo([]byte("hello"), destination); e != nil {
			t.Fatal(e)
		}
		var b [16]byte
		n, src, e := server.ReadFrom(b[:])
		if e != nil || string(b[:n]) != "hello" {
			t.Fatalf("%s write: %q %v", kind, b[:n], e)
		}
		if _, e = server.WriteTo([]byte("reply"), src); e != nil {
			t.Fatal(e)
		}
		n, remote, e := c.ReadFrom(b[:])
		c.Close()
		if e != nil || string(b[:n]) != "reply" || remote.String() != server.LocalAddr().String() {
			t.Fatalf("%s read: %q %v %v", kind, b[:n], remote, e)
		}
	}
}

// captureAdapter implements the complete original adapter contract for testing
// metadata passed to an already-installed adapter; it is not production code.
type captureAdapter struct {
	meta    *E3F12urwswL.D0mLgQF
	options int
	err     error
	udp     bool
}

func (a *captureAdapter) Addr() string                                        { return "test" }
func (a *captureAdapter) Name() string                                        { return "test" }
func (a *captureAdapter) Type() E3F12urwswL.ASFyBDZD                          { return 0 }
func (a *captureAdapter) SupportUDP() bool                                    { return true }
func (a *captureAdapter) MarshalJSON() ([]byte, error)                        { return nil, nil }
func (a *captureAdapter) Unwrap(*E3F12urwswL.D0mLgQF) E3F12urwswL.ECww8avx7eC { return nil }
func (a *captureAdapter) StreamConn(c net.Conn, _ *E3F12urwswL.D0mLgQF) (net.Conn, error) {
	return c, nil
}
func (a *captureAdapter) DialContext(_ context.Context, m *E3F12urwswL.D0mLgQF, o ...vP_e5OQAve.PNSpXRb) (E3F12urwswL.PnyceLR, error) {
	a.meta = m
	a.options = len(o)
	return nil, a.err
}
func (a *captureAdapter) ListenPacketContext(_ context.Context, m *E3F12urwswL.D0mLgQF, o ...vP_e5OQAve.PNSpXRb) (E3F12urwswL.C08JTZ2l, error) {
	a.meta = m
	a.options = len(o)
	a.udp = true
	return nil, a.err
}
func TestAdapterMetadataAndErrorPropagation(t *testing.T) {
	sentinel := errors.New("adapter failed")
	adapter := &captureAdapter{err: sentinel}
	out := &GyyO51{NyLbiXb4OrlL: "vmess", ucp4iC: adapter, iCySnjtd: net.ParseIP("192.0.2.1")}
	target, err := G6bHD5JA.U80FDt("example.com:443")
	if err != nil {
		t.Fatal(err)
	}
	if c, e := out.DialTCP(nil, target); c != nil || e != sentinel {
		t.Fatal("TCP adapter error changed")
	}
	if adapter.meta.WJ95LyOL != "example.com" || adapter.meta.JRXMVvg != 443 || adapter.meta.AcXafZ1 != 0 || adapter.options != 1 {
		t.Fatalf("TCP metadata/options: %+v/%d", adapter.meta, adapter.options)
	}
	if c, e := out.DialUDP(nil, target); c != nil || e != sentinel {
		t.Fatal("UDP adapter error changed")
	}
	if !adapter.udp || adapter.meta.AcXafZ1 != 1 || adapter.meta.ZouEpmjw8 != nil {
		t.Fatal("UDP metadata differs")
	}
}
