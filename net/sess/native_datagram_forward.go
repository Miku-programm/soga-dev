package sess

import (
	"context"
	"net"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/util/errutil"
	C "recovered.local/soga/config"
	"recovered.local/soga/user"
	"time"
)

// Jv9wkYGl @0x18874a0 and cleanup closure @0x1888940 forward one
// outbound socket's responses to the authenticated native datagram transport.
func Jv9wkYGl(ctx context.Context, remote net.PacketConn, source net.Addr, local KsKg3yls, user user.PgaNd3014U) error {
	buf := pool.PoolAt50961D8.Get(4096)
	defer func() { pool.PoolAt50961D8.Put(buf); errutil.SlG_Y_mg("handle remote udp panic") }()
	for {
		select {
		case <-ctx.Done():
			return errutil.ErrorAt5097800
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
