package LCe4Wkk2

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	E "recovered.local/soga/E3F12urwswL"
	O "recovered.local/soga/vP_e5OQAve"
)

// B8EHsVY7a retains the original 72-byte RTTI layout (aAuS0Xw.B8EHsVY7a).
type B8EHsVY7a struct {
	m50I26ZDL    string
	x3HUaU8Ln1S1 string
	xDan8hA5LV   string
	iOKxDl       E.ASFyBDZD
	ntseIWz      bool
	bEbqkM8      int
}

// Name: 0x1b55220.
func (b *B8EHsVY7a) Name() string { return b.m50I26ZDL }

// Type: 0x1b55240.
func (b *B8EHsVY7a) Type() E.ASFyBDZD { return b.iOKxDl }

// StreamConn: 0x1b55260. The original returns the input connection with the error.
func (b *B8EHsVY7a) StreamConn(c net.Conn, _ *E.D0mLgQF) (net.Conn, error) {
	return c, errors.New("no support")
}

// ListenPacketContext: 0x1b55360. This is the actual base implementation.
func (b *B8EHsVY7a) ListenPacketContext(context.Context, *E.D0mLgQF, ...O.PNSpXRb) (E.C08JTZ2l, error) {
	return nil, errors.New("no support")
}

// SupportUDP: 0x1b55420.
func (b *B8EHsVY7a) SupportUDP() bool { return b.ntseIWz }

// MarshalJSON: 0x1b55440. The original v2 call passes DefaultOptionsV1 at 0x505b4e0.
func (b *B8EHsVY7a) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"type": b.iOKxDl.String()})
}

// Addr: 0x1b55520.
func (b *B8EHsVY7a) Addr() string { return b.x3HUaU8Ln1S1 }

// Unwrap: 0x1b55540.
func (b *B8EHsVY7a) Unwrap(*E.D0mLgQF) E.ECww8avx7eC { return nil }

// DialOptions: 0x1b55560. Appended adapter options override earlier caller options.
func (b *B8EHsVY7a) DialOptions(options []O.PNSpXRb) []O.PNSpXRb {
	if b.xDan8hA5LV != "" {
		options = append(options, O.VcC742(b.xDan8hA5LV))
	}
	if b.bEbqkM8 != 0 {
		options = append(options, O.Fum1PEvMPct(b.bEbqkM8))
	}
	return options
}
