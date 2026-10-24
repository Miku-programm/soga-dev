package CpHCdepV

import (
	"errors"
	"github.com/jellydator/ttlcache/v3"
	"net"
	"testing"
	"time"
	"unsafe"
)

func proxyPacketFixture(raw net.PacketConn, force bool) *EuI1uJ2zx {
	return &EuI1uJ2zx{raw, force, ttlcache.New[string, net.Addr](ttlcache.WithTTL[string, net.Addr](time.Minute)), ttlcache.New[string, net.Addr](ttlcache.WithTTL[string, net.Addr](time.Minute))}
}
func TestOriginalUDPProxyRealSocket(t *testing.T) {
	var v EuI1uJ2zx
	if unsafe.Sizeof(v) != 40 || unsafe.Offsetof(v.akXo7g4lc) != 16 || unsafe.Offsetof(v.sUKRhMwzm) != 24 || unsafe.Offsetof(v.aLeZsGA) != 32 {
		t.Fatal("RTTI")
	}
	for _, force := range []bool{false, true} {
		server, e := net.ListenPacket("udp", "127.0.0.1:0")
		if e != nil {
			t.Fatal(e)
		}
		defer server.Close()
		client, e := net.ListenPacket("udp", "127.0.0.1:0")
		if e != nil {
			t.Fatal(e)
		}
		defer client.Close()
		server.SetDeadline(time.Now().Add(3 * time.Second))
		client.SetDeadline(time.Now().Add(3 * time.Second))
		wrapper := proxyPacketFixture(server, force)
		if wrapper.GetRawPacketConn() != server {
			t.Fatal("raw")
		}
		if force {
			client.WriteTo([]byte("drop-before-header"), server.LocalAddr())
		}
		client.WriteTo([]byte("PROXY UNKNOWN\r\ndrop-unknown"), server.LocalAddr())
		client.WriteTo([]byte("PROXY TCP4 192.0.2.8 198.51.100.4 3210 443\r\nfirst"), server.LocalAddr())
		buf := make([]byte, 2048)
		n, a, e := wrapper.ReadFrom(buf)
		if e != nil || string(buf[:n]) != "first" || a.String() != "192.0.2.8:3210" {
			t.Fatal(n, a, e, string(buf[:n]))
		}
		client.WriteTo([]byte("second"), server.LocalAddr())
		n, a, e = wrapper.ReadFrom(buf)
		if e != nil || string(buf[:n]) != "second" || a.String() != "192.0.2.8:3210" {
			t.Fatal(n, a, e)
		}
		if n, e = wrapper.WriteTo([]byte("reply"), a); n != 5 || e != nil {
			t.Fatal(n, e)
		}
		n, _, e = client.ReadFrom(buf)
		if e != nil || string(buf[:n]) != "reply" {
			t.Fatal(n, e)
		}
		wrapper.sUKRhMwzm.DeleteAll()
		wrapper.aLeZsGA.DeleteAll()
		if !force {
			client.WriteTo([]byte("plain"), server.LocalAddr())
			n, a, e = wrapper.ReadFrom(buf)
			if e != nil || string(buf[:n]) != "plain" || a.String() != client.LocalAddr().String() {
				t.Fatal(n, a, e)
			}
		}
	}
}

type proxyRawFixture struct {
	net.PacketConn
	addr    net.Addr
	err     error
	written net.Addr
}

func (f *proxyRawFixture) ReadFrom(p []byte) (int, net.Addr, error) {
	copy(p, "PROXY")
	return 5, f.addr, f.err
}
func (f *proxyRawFixture) WriteTo(p []byte, a net.Addr) (int, error) { f.written = a; return 1, f.err }
func TestOriginalUDPProxyPartialErrors(t *testing.T) {
	a := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1}
	err := errors.New("partial")
	raw := &proxyRawFixture{addr: a, err: err}
	c := proxyPacketFixture(raw, true)
	buf := make([]byte, 20)
	n, got, e := c.ReadFrom(buf)
	if n != 5 || got != a || e != err {
		t.Fatal(n, got, e)
	}
	n, e = c.WriteTo(buf, a)
	if n != 1 || e != err || raw.written != a {
		t.Fatal(n, e)
	}
	b := &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 2}
	c.pEL3YD3aSM2r(a, b)
	n, e = c.WriteTo(buf, b)
	if n != 1 || e != err || raw.written != a {
		t.Fatal(n, e)
	}
}
