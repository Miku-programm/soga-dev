// Package E3F12urwswL restores the outbound contracts recorded in RTTI.
package E3F12urwswL

import (
	"context"
	"net"
	"net/netip"
	"recovered.local/soga/vP_e5OQAve"
	"time"
)

type ASFyBDZD int
type AN18DDo uint16
type FRahPmPZ1B02 int
type INWQbDhB7i6 int
type CpPqAoGa_ int
type GYe3KD []string

// D0mLgQF: RTTI 0x370a350, size 160. COgCoHHQaf.IxcIKdAZLNX
// is net/netip.AddrPort: its Addr contains uint128 and unique.Handle[addrDetail].
type D0mLgQF struct {
	AcXafZ1      FRahPmPZ1B02   `json:"network"`
	IDNFAHr      INWQbDhB7i6    `json:"type"`
	Vp2FNOapN    net.IP         `json:"sourceIP"`
	ZouEpmjw8    net.IP         `json:"destinationIP"`
	KR6XV2CSa    AN18DDo        `json:"sourcePort"`
	JRXMVvg      AN18DDo        `json:"destinationPort"`
	WJ95LyOL     string         `json:"host"`
	ZPRYp07EvgW  CpPqAoGa_      `json:"dnsMode"`
	AmdGy4j      string         `json:"processPath"`
	DQoWWz_a     string         `json:"specialProxy"`
	LYMIRODDIrtj netip.AddrPort `json:"-"`
}

// Kna5ouY0MG9d: RTTI 0x3654b98, size 32.
type Kna5ouY0MG9d struct {
	YjBRQo63weIK time.Time `json:"time"`
	MJGtIA       uint16    `json:"delay"`
	Fw5W_ZvL     uint16    `json:"meanDelay"`
}

// Interface method sets are read from RTTI 0x3664cb0 and 0x3654af0.
type PnyceLR interface {
	net.Conn
	AppendToChains(C74CUtQ9FqH)
	Chains() GYe3KD
}
type C08JTZ2l interface {
	net.PacketConn
	AppendToChains(C74CUtQ9FqH)
	Chains() GYe3KD
}

// C74CUtQ9FqH: RTTI 0x3654c40, all nine methods.
type C74CUtQ9FqH interface {
	Addr() string
	DialContext(context.Context, *D0mLgQF, ...vP_e5OQAve.PNSpXRb) (PnyceLR, error)
	ListenPacketContext(context.Context, *D0mLgQF, ...vP_e5OQAve.PNSpXRb) (C08JTZ2l, error)
	MarshalJSON() ([]byte, error)
	Name() string
	StreamConn(net.Conn, *D0mLgQF) (net.Conn, error)
	SupportUDP() bool
	Type() ASFyBDZD
	Unwrap(*D0mLgQF) ECww8avx7eC
}

// ECww8avx7eC: RTTI 0x36c9d28, all fifteen methods.
type ECww8avx7eC interface {
	C74CUtQ9FqH
	Alive() bool
	DelayHistory() []Kna5ouY0MG9d
	Dial(*D0mLgQF) (PnyceLR, error)
	DialUDP(*D0mLgQF) (C08JTZ2l, error)
	LastDelay() uint16
	URLTest(context.Context, string) (uint16, uint16, error)
}
