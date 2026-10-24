//go:build linux

package LCe4Wkk2

import (
	"context"
	"errors"
	"fmt"
	"net"
	E "recovered.local/soga/E3F12urwswL"
	obfs "recovered.local/soga/GtTssd5z"
	stream "recovered.local/soga/VwwneWVFFv"
	protocol "recovered.local/soga/low0s_10x"
	cipher "recovered.local/soga/m3Cw8P3_Le"
	aead "recovered.local/soga/u3UOJmirww7"
	O "recovered.local/soga/vP_e5OQAve"
	"strconv"
)

// ShadowsocksROptions is an analysis name for the original 168-byte argument
// to CR6bC5 (0x1b62020). Offsets are verified against its caller and callee.
type ShadowsocksROptions struct {
	FFrN8gdu
	Name          string
	Server        string
	Port          int
	Password      string
	Cipher        string
	Obfs          string
	ObfsParam     string
	Protocol      string
	ProtocolParam string
	UDP           bool
}

// RpNX5wlEsF preserves the target's 56-byte RTTI layout.
type RpNX5wlEsF struct {
	*B8EHsVY7a
	dJQXBb      cipher.HjDHFt_
	dGQqHhbAecJ obfs.AGNBYcIkk
	nioHey      protocol.Shdj3DH1
}

// CR6bC5: 0x1b62020. Only exact lowercase "none" is rewritten to "dummy".
// The latter supplies a 16-byte password-derived authentication key with no IV.
func CR6bC5(o ShadowsocksROptions) (*RpNX5wlEsF, error) {
	if o.Cipher == "none" {
		o.Cipher = "dummy"
	}
	addr := net.JoinHostPort(o.Server, strconv.Itoa(o.Port))
	c, e := cipher.B93MVPtD(o.Cipher, nil, o.Password)
	if e != nil {
		return nil, fmt.Errorf("ssr %s initialize error: %w", addr, e)
	}
	var key []byte
	var ivSize int
	if o.Cipher == "dummy" {
		key = cipher.ArNrGS_EV(o.Password, 16)
	} else if s, ok := c.(*cipher.E3HaL6A8); ok {
		key = s.Ez2IY3fv
		ivSize = s.IVSize()
	} else {
		return nil, fmt.Errorf("%s is not none or a supported stream cipher in ssr", o.Cipher)
	}
	ob, overhead, e := obfs.IRT7O1m(o.Obfs, &obfs.Ja7FBino7HL{O1k0EzteK: o.Server, P2XQpgsRo9: o.Port, JtkCoK71gtjm: key, UpKlHDjL: ivSize, NhWlHt412H8S: o.ObfsParam})
	if e != nil {
		return nil, fmt.Errorf("ssr %s initialize obfs error: %w", addr, e)
	}
	pr, e := protocol.HBLvlfQr0d(o.Protocol, &protocol.PZfVxx4UGRb{Z2UHwhi2h4: key, UqLugVVWCKJ: overhead, BOJR3xVshAb: o.ProtocolParam})
	if e != nil {
		return nil, fmt.Errorf("ssr %s initialize protocol error: %w", addr, e)
	}
	return &RpNX5wlEsF{B8EHsVY7a: &B8EHsVY7a{m50I26ZDL: o.Name, x3HUaU8Ln1S1: addr, xDan8hA5LV: o.PiTAGr6U5KDG, iOKxDl: 3, ntseIWz: o.UDP, bEbqkM8: o.VsBRTR}, dJQXBb: c, dGQqHhbAecJ: ob, nioHey: pr}, nil
}

// StreamConn: 0x1b611e0. ObtainWriteIV is called before the protocol header.
func (s *RpNX5wlEsF) StreamConn(c net.Conn, m *E.D0mLgQF) (net.Conn, error) {
	c = s.dGQqHhbAecJ.StreamConn(c)
	c = s.dJQXBb.StreamConn(c)
	var iv []byte
	switch v := c.(type) {
	case *stream.AriwPpia:
		var e error
		iv, e = v.ObtainWriteIV()
		if e != nil {
			return nil, e
		}
	case *aead.AZVSCxZmZQ:
		return nil, errors.New("invalid connection type")
	}
	c = s.nioHey.StreamConn(c, iv)
	_, e := c.Write(o2HGjDHIVO(m))
	return c, e
}

// DialContext: 0x1b61420, with original raw-socket cleanup on any failure.
func (s *RpNX5wlEsF) DialContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (out E.PnyceLR, err error) {
	c, err := O.I_y9nlK1q8(ctx, "tcp", s.Addr(), s.DialOptions(options)...)
	if err != nil {
		return nil, fmt.Errorf("%s connect error: %w", s.Addr(), err)
	}
	mS0F0Ti(c)
	defer func() {
		if err != nil {
			c.Close()
		}
	}()
	wrapped, err := s.StreamConn(c, m)
	if err != nil {
		return nil, err
	}
	return &lGta_1QkvQ{Conn: wrapped, bnOtGuS: E.GYe3KD{s.Name()}}, nil
}

// ListenPacketContext: 0x1b61ae0. UDP bypasses obfs and wraps cipher then protocol.
func (s *RpNX5wlEsF) ListenPacketContext(ctx context.Context, _ *E.D0mLgQF, options ...O.PNSpXRb) (E.C08JTZ2l, error) {
	c, e := O.OpEY1B7hjf(ctx, "udp", "", s.DialOptions(options)...)
	if e != nil {
		return nil, e
	}
	addr, e := g0YPF3("udp", s.Addr())
	if e != nil {
		c.Close()
		return nil, e
	}
	c = s.nioHey.PacketConn(s.dJQXBb.PacketConn(c))
	return &uMtOz6o{PacketConn: &dTAcwl_gA{PacketConn: c, zbHzsqAUlM: addr}, f7i_T5n5JXyM: E.GYe3KD{s.Name()}}, nil
}

var _ E.C74CUtQ9FqH = (*RpNX5wlEsF)(nil)
