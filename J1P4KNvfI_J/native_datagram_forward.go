package J1P4KNvfI_J

import (
	"context"
	"net"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/HxKViEp2a"
	C "recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/nzATvajt"
	"time"
)

// Jv9wkYGl @0x18874a0 and cleanup closure @0x1888940 forward one
// outbound socket's responses to the authenticated native datagram transport.
func Jv9wkYGl(ctx context.Context, remote net.PacketConn, source net.Addr, local KsKg3yls, user nzATvajt.PgaNd3014U) error {
	buf := D_zCag1qTV.PoolAt50961D8.Get(4096)
	defer func() { D_zCag1qTV.PoolAt50961D8.Put(buf); HxKViEp2a.SlG_Y_mg("handle remote udp panic") }()
	for {
		select {
		case <-ctx.Done():
			return HxKViEp2a.ErrorAt5097800
		default:
		}
		_ = remote.SetReadDeadline(time.Now().Add(time.Duration(C.UDPTimeoutMinutes()) * time.Minute))
		n, from, err := remote.ReadFrom(buf)
		if err != nil {
			return err
		}
		_ = local.SetWriteDeadline(time.Now().Add(time.Minute))
		if _, err = local.WriteToProxy(buf[:n], user, source, from); err != nil {
			return err
		}
	}
}
