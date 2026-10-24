//go:build linux

package ssactual

import (
	"context"
	E "recovered.local/soga/net/outbound"
	O "recovered.local/soga/net/socket"
)

// NewDirectAdapter expresses the original inlined construction in
// routing.(*GyyO51).o7v2ON9D at 0x1b7c4c0. Helper name is new.
func NewDirectAdapter() *W8KUYH {
	return &W8KUYH{&B8EHsVY7a{m50I26ZDL: "DIRECT", iOKxDl: 0, ntseIWz: true}}
}

// DialContext: 0x1b56260. It resolves/dials through the restored original
// outbound socket package, adjusts concrete TCP keepalive, then records its chain.
func (d *W8KUYH) DialContext(ctx context.Context, m *E.D0mLgQF, options ...O.PNSpXRb) (E.PnyceLR, error) {
	remote := m.RemoteAddress()
	c, err := O.I_y9nlK1q8(ctx, "tcp", remote, d.DialOptions(options)...)
	if err != nil {
		return nil, err
	}
	mS0F0Ti(c)
	return &lGta_1QkvQ{Conn: c, bnOtGuS: E.GYe3KD{d.Name()}}, nil
}

var _ E.C74CUtQ9FqH = (*W8KUYH)(nil)
