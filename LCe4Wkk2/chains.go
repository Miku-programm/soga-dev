package LCe4Wkk2

import (
	"net"
	E "recovered.local/soga/E3F12urwswL"
)

// The embedded obfuscated standard-library interfaces map to net.Conn/PacketConn.
// Both original structs are 40 bytes; Chains returns the stored slice without copying.
type lGta_1QkvQ struct {
	net.Conn
	bnOtGuS E.GYe3KD
}
type uMtOz6o struct {
	net.PacketConn
	f7i_T5n5JXyM E.GYe3KD
}

// Chains: 0x1b55800.
func (c *lGta_1QkvQ) Chains() E.GYe3KD { return c.bnOtGuS }

// AppendToChains: 0x1b55820.
func (c *lGta_1QkvQ) AppendToChains(a E.C74CUtQ9FqH) { c.bnOtGuS = append(c.bnOtGuS, a.Name()) }

// Chains: 0x1b55920.
func (c *uMtOz6o) Chains() E.GYe3KD { return c.f7i_T5n5JXyM }

// AppendToChains: 0x1b55940.
func (c *uMtOz6o) AppendToChains(a E.C74CUtQ9FqH) { c.f7i_T5n5JXyM = append(c.f7i_T5n5JXyM, a.Name()) }

var _ E.PnyceLR = (*lGta_1QkvQ)(nil)
var _ E.C08JTZ2l = (*uMtOz6o)(nil)
