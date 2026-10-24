//go:build linux

package ssactual

import (
	"bytes"
	"net"
	"testing"
	"unsafe"
)

func TestShadowsocksPacketOriginalBoundaries(t *testing.T) {
	if unsafe.Sizeof(dTAcwl_gA{}) != 32 {
		t.Fatal("RTTI")
	}
	p := &packetFixture{}
	c := &dTAcwl_gA{PacketConn: p, zbHzsqAUlM: &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 9}}
	n, err := c.WriteTo([]byte("hi"), &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 80})
	if n != 9 || err != nil || !bytes.Equal(p.packet, []byte{1, 192, 0, 2, 1, 0, 80, 'h', 'i'}) {
		t.Fatal(n, err, p.packet)
	}
	p.reported = 9
	b := bytes.Repeat([]byte{'z'}, 13)
	n, addr, err := c.ReadFrom(b)
	if n != 2 || err != nil || addr.String() != "192.0.2.1:80" || string(b[:6]) != "hizzzz" || cap(addr.(*net.UDPAddr).IP) != 4 {
		t.Fatal(n, addr, err, b)
	}
	// Unlike the SOCKS wrapper, malformed short packets cannot use trailing bytes.
	p.reported = 2
	if n, a, err := c.ReadFrom(make([]byte, 40)); n != 0 || a != nil || err == nil {
		t.Fatal(n, a, err)
	}
	p.packet = []byte{3, 1, 'a', 0, 80, 'x'}
	p.reported = 6
	if _, _, err := c.ReadFrom(make([]byte, 40)); err == nil {
		t.Fatal("domain unexpectedly accepted")
	}
}
