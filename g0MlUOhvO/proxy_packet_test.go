package g0MlUOhvO

import (
	"bytes"
	"context"
	"io"
	"net"
	"recovered.local/soga/JcmlVIK0T0i"
	"testing"
	"testing/synctest"
	"time"
	"unsafe"
)

func TestProxyPacketRealUDPAddressRewrite(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	client, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	server.SetDeadline(time.Now().Add(3 * time.Second))
	client.SetDeadline(time.Now().Add(3 * time.Second))
	p := SimI1bq(ctx, server, true).(*ARZKOqmN5e)
	if unsafe.Sizeof(*p) != 40 || p.GetRawPacketConn() != server {
		t.Fatal("layout/raw")
	}
	// A strict-mode naked datagram is skipped, then the real PROXY v2 datagram
	// establishes both directions. Subsequent packets no longer need a header.
	client.WriteTo([]byte("skip"), server.LocalAddr())
	source := &net.UDPAddr{IP: net.IPv4(192, 0, 2, 7), Port: 1234}
	header, err := JcmlVIK0T0i.X_WkJPwas4L(2, source, server.LocalAddr()).Format()
	if err != nil {
		t.Fatal(err)
	}
	client.WriteTo(append(header, []byte("first")...), server.LocalAddr())
	buf := make([]byte, 2048)
	n, addr, err := p.ReadFrom(buf)
	if err != nil || string(buf[:n]) != "first" || addr.String() != source.String() {
		t.Fatal(n, addr, err, string(buf[:n]))
	}
	client.WriteTo([]byte("second"), server.LocalAddr())
	n, addr, err = p.ReadFrom(buf)
	if err != nil || string(buf[:n]) != "second" || addr.String() != source.String() {
		t.Fatal(n, addr, err, string(buf[:n]))
	}
	if _, err = p.WriteTo([]byte("reply"), addr); err != nil {
		t.Fatal(err)
	}
	n, _, err = client.ReadFrom(buf)
	if err != nil || string(buf[:n]) != "reply" {
		t.Fatal(n, err, string(buf[:n]))
	}
}

type proxyDatagram struct {
	data []byte
	addr net.Addr
	err  error
}
type packetFixture struct {
	net.PacketConn
	packets []proxyDatagram
	wrote   net.Addr
}

func (c *packetFixture) ReadFrom(p []byte) (int, net.Addr, error) {
	if len(c.packets) == 0 {
		return 0, nil, io.EOF
	}
	x := c.packets[0]
	c.packets = c.packets[1:]
	return copy(p, x.data), x.addr, x.err
}
func (c *packetFixture) WriteTo(p []byte, a net.Addr) (int, error) { c.wrote = a; return len(p), nil }
func TestProxyPacketLocalAndMalformedHeaders(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	from := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 9}
	local := append([]byte("\r\n\r\n\x00\r\nQUIT\n"), 0x20, 0, 0, 0)
	local = append(local, []byte("discard")...)
	c := &packetFixture{packets: []proxyDatagram{{local, from, nil}, {[]byte("PROXY broken\r\nraw"), from, nil}}}
	p := SimI1bq(ctx, c, false).(*ARZKOqmN5e)
	b := make([]byte, 128)
	n, a, err := p.ReadFrom(b)
	if err != nil || a != from || string(b[:n]) != "PROXY broken\r\nraw" {
		t.Fatal(n, a, err, string(b[:n]))
	}
	// Underlying n+error is returned without parsing or cache mutation.
	c.packets = []proxyDatagram{{[]byte("partial"), from, io.ErrUnexpectedEOF}}
	n, a, err = p.ReadFrom(b)
	if n != 7 || a != from || err != io.ErrUnexpectedEOF {
		t.Fatal(n, a, err)
	}
	if _, err = p.WriteTo([]byte("x"), from); err != nil || c.wrote != from {
		t.Fatal("uncached response", err)
	}
}
func TestProxyPacketBidirectionalTTL(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		p := SimI1bq(ctx, &packetFixture{}, false).(*ARZKOqmN5e)
		from := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1}
		to := &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 2}
		p.h5MasPA(from, to)
		synctest.Wait()
		time.Sleep(4 * time.Minute)
		if a, ok := p.ssa9qlNCrrBM(from); !ok || a != to {
			t.Fatal("mapping")
		}
		time.Sleep(4 * time.Minute)
		if a, ok := p.eQq7GgK(to); !ok || a != from {
			t.Fatal("both directions were not refreshed")
		}
		time.Sleep(6 * time.Minute)
		synctest.Wait()
		if _, ok := p.ssa9qlNCrrBM(from); ok {
			t.Fatal("forward expired")
		}
		if _, ok := p.eQq7GgK(to); ok {
			t.Fatal("reverse expired")
		}
	})
}
func TestProxyV1StreamAddressConvertedToUDP(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	from := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 9}
	c := &packetFixture{packets: []proxyDatagram{{[]byte("PROXY TCP4 192.0.2.1 192.0.2.2 456 789\r\ncontent"), from, nil}}}
	p := SimI1bq(ctx, c, false)
	b := make([]byte, 128)
	n, a, err := p.ReadFrom(b)
	if err != nil || !bytes.Equal(b[:n], []byte("content")) || a.String() != "192.0.2.1:456" {
		t.Fatal(n, a, err)
	}
	if _, ok := a.(*net.UDPAddr); !ok {
		t.Fatal("v1 TCP source must become UDP")
	}
}
