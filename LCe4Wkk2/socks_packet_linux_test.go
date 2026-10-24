//go:build linux

package LCe4Wkk2

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	S "recovered.local/soga/Bx4XCl"
	"testing"
	"time"
	"unsafe"
)

type packetFixture struct {
	net.PacketConn
	packet   []byte
	reported int
	closed   bool
}

func (p *packetFixture) ReadFrom(b []byte) (int, net.Addr, error) {
	copy(b, p.packet)
	return p.reported, nil, nil
}
func (p *packetFixture) WriteTo(b []byte, _ net.Addr) (int, error) {
	p.packet = append([]byte(nil), b...)
	return len(b), nil
}
func (p *packetFixture) Close() error { p.closed = true; return io.ErrClosedPipe }

type closeFixture struct {
	net.Conn
	closed bool
}

func (c *closeFixture) Close() error { c.closed = true; return io.ErrUnexpectedEOF }

func TestSocksUDPOriginalBufferAndCountBehavior(t *testing.T) {
	if unsafe.Sizeof(hNLZlR1I{}) != 48 {
		t.Fatal("RTTI layout")
	}
	p := &packetFixture{}
	control := &closeFixture{}
	c := &hNLZlR1I{PacketConn: p, f67_L17asEA: control}
	n, err := c.WriteTo([]byte("hi"), &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 9})
	if err != nil || n != 12 {
		t.Fatal(n, err)
	}
	p.reported = 12
	b := bytes.Repeat([]byte{'z'}, 16)
	n, addr, err := c.ReadFrom(b)
	if err != nil || n != 2 || addr.String() != "127.0.0.1:9" || string(b[:6]) != "hizzzz" {
		t.Fatal(n, addr, err, b)
	}
	if cap(addr.(*net.UDPAddr).IP) != 4 {
		t.Fatal("original IP allocation size")
	}
	// Full-buffer parsing can consume bytes beyond the reported datagram count.
	p.reported = 2
	n, _, err = c.ReadFrom(make([]byte, 16))
	if err != nil || n != -8 {
		t.Fatal(n, err)
	}
	p.packet, _ = S.XFq85XFER(S.L4JRaRodkw("host:1"), nil)
	if _, _, err = c.ReadFrom(make([]byte, 32)); err == nil || err.Error() != "parse udp addr error" {
		t.Fatal(err)
	}
	if err = c.Close(); err != io.ErrClosedPipe || !p.closed || !control.closed {
		t.Fatal(err)
	}
}

func TestSocksUDPAssociateLocalRelay(t *testing.T) {
	tcp, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer tcp.Close()
	udp, err := net.ListenPacket("udp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer udp.Close()
	udp.SetDeadline(time.Now().Add(3 * time.Second))
	tcpDone := make(chan error, 1)
	udpDone := make(chan error, 1)
	go func() {
		conn, err := tcp.Accept()
		if err != nil {
			tcpDone <- err
			return
		}
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(3 * time.Second))
		hello := make([]byte, 3)
		if _, err = io.ReadFull(conn, hello); err != nil {
			tcpDone <- err
			return
		}
		if !bytes.Equal(hello, []byte{5, 1, 0}) {
			tcpDone <- errors.New("greeting")
			return
		}
		if _, err = conn.Write([]byte{5, 0}); err != nil {
			tcpDone <- err
			return
		}
		request := make([]byte, 10)
		if _, err = io.ReadFull(conn, request); err != nil {
			tcpDone <- err
			return
		}
		if !bytes.Equal(request, []byte{5, 3, 0, 1, 0, 0, 0, 0, 0, 0}) {
			tcpDone <- errors.New("associate")
			return
		}
		// An unspecified relay IP must be replaced by the proxy server IP;
		// the independently allocated UDP port must remain intact.
		bound := S.KvyKMZ(&net.UDPAddr{IP: net.IPv4zero, Port: udp.LocalAddr().(*net.UDPAddr).Port})
		if _, err = conn.Write(append([]byte{5, 0, 0}, bound...)); err != nil {
			tcpDone <- err
			return
		}
		_, err = io.Copy(io.Discard, conn)
		tcpDone <- err
	}()
	go func() {
		b := make([]byte, 128)
		n, peer, err := udp.ReadFrom(b)
		if err == nil {
			var addr, payload []byte
			addr, payload, err = S.CkUmXXHwOEH(b[:n])
			if err == nil && (!bytes.Equal(addr, S.L4JRaRodkw("192.0.2.1:53")) || string(payload) != "echo") {
				err = errors.New("relay datagram")
			}
			if err == nil {
				_, err = udp.WriteTo(b[:n], peer)
			}
		}
		udpDone <- err
	}()
	a := Auqn90u2NBQ(SocksOptions{Name: "udp-socks", Server: "127.0.0.1", Port: tcp.Addr().(*net.TCPAddr).Port, UDP: true})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	c, err := a.ListenPacketContext(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(3 * time.Second))
	if n, err := c.WriteTo([]byte("echo"), &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 53}); err != nil || n != 14 {
		t.Fatal(n, err)
	}
	b := make([]byte, 128)
	n, addr, err := c.ReadFrom(b)
	if err != nil || string(b[:n]) != "echo" || addr.String() != "192.0.2.1:53" || c.Chains()[0] != "udp-socks" {
		t.Fatal(n, addr, err, c.Chains())
	}
	c.Close()
	if err := <-udpDone; err != nil {
		t.Fatal(err)
	}
	if err := <-tcpDone; err != nil {
		t.Fatal(err)
	}
}
