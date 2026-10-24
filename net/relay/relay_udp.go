package relay

import (
	"context"
	"net"
	D "recovered.local/soga/buf/pool"
	H "recovered.local/soga/util/errutil"
	C "recovered.local/soga/config"
	"time"
)

// ZcvMMHXas8k: 0x15af840. The read deadline is set on the destination
// packet socket, even though the next Read operates on the source connection.
func ZcvMMHXas8k(ctx context.Context, src net.Conn, dst net.PacketConn, addr *net.UDPAddr, buf []byte) error {
	defer func() { H.SlG_Y_mg("pipeline tcp to udp") }()
	timeout := time.Duration(C.UDPTimeoutMinutes()) * time.Minute
	for {
		select {
		case <-ctx.Done():
			return H.ErrorAt5097800
		default:
		}
		dst.SetReadDeadline(time.Now().Add(timeout))
		n, readErr := src.Read(buf)
		if n > 0 {
			dst.SetWriteDeadline(time.Now().Add(time.Minute))
			if _, writeErr := dst.WriteTo(buf[:n], addr); writeErr != nil {
				return H.T2i1ItMaa(readErr, writeErr)
			}
		}
		if readErr != nil {
			return readErr
		}
	}
}

// DjO2Gp: 0x15afc40. A ReadFrom error discards accompanying bytes; otherwise
// even a zero-length packet is written to the destination connection.
func DjO2Gp(ctx context.Context, src net.PacketConn, dst net.Conn) error {
	buf := D.PoolAt50961D8.Get(16384)
	defer func() { D.PoolAt50961D8.Put(buf); H.SlG_Y_mg("pipeline udp to tcp") }()
	timeout := time.Duration(C.UDPTimeoutMinutes()) * time.Minute
	for {
		select {
		case <-ctx.Done():
			return H.ErrorAt5097800
		default:
		}
		src.SetReadDeadline(time.Now().Add(timeout))
		n, _, err := src.ReadFrom(buf)
		if err != nil {
			return err
		}
		dst.SetWriteDeadline(time.Now().Add(10 * time.Minute))
		if _, err = dst.Write(buf[:n]); err != nil {
			return H.T2i1ItMaa(err)
		}
	}
}
