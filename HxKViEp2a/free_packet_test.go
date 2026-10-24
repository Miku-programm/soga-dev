package HxKViEp2a

import (
	"net"
	"testing"
)

type packetFreeFixture struct {
	net.PacketConn
	freed int
}

func (p *packetFreeFixture) Free() { p.freed++ }

type packetRawFixture struct {
	net.PacketConn
	raw net.PacketConn
}

func (p *packetRawFixture) GetRawPacketConn() net.PacketConn { return p.raw }
func TestOriginalPacketFreeTraversal(t *testing.T) {
	inner := &packetFreeFixture{}
	PpnNLr_pf7He(&packetRawFixture{raw: &packetRawFixture{raw: inner}})
	if inner.freed != 1 {
		t.Fatal("unwrap")
	}
	PpnNLr_pf7He(nil)
	PpnNLr_pf7He((*packetFreeFixture)(nil))
	PpnNLr_pf7He(&packetRawFixture{})
	if inner.freed != 1 {
		t.Fatal("unexpected free")
	}
}
