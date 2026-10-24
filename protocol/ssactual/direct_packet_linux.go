//go:build linux

package ssactual

import (
	"context"
	E "recovered.local/soga/net/outbound"
	O "recovered.local/soga/net/socket"
)

// W8KUYH retains the original one-pointer direct-adapter RTTI layout.
type W8KUYH struct{ *B8EHsVY7a }

// ListenPacketContext restores 0x1b56680: the target metadata is not read;
// an unconnected UDP socket is opened, wrapped and marked with this adapter name.
func (d *W8KUYH) ListenPacketContext(ctx context.Context, _ *E.D0mLgQF, options ...O.PNSpXRb) (E.C08JTZ2l, error) {
	c, err := O.OpEY1B7hjf(ctx, "udp", "", d.DialOptions(options)...)
	if err != nil {
		return nil, err
	}
	return &uMtOz6o{PacketConn: &fS848iPPRMrU{PacketConn: c}, f7i_T5n5JXyM: E.GYe3KD{d.Name()}}, nil
}
