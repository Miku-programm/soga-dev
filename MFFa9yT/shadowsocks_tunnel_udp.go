package MFFa9yT

import (
	"context"
	"net"
	D "recovered.local/soga/D_zCag1qTV"
	H "recovered.local/soga/HxKViEp2a"
	J "recovered.local/soga/J1P4KNvfI_J"
	F "recovered.local/soga/JX5NGDvCQwty"
	"time"
)

// mZ4KusAhuze @0x1d706c0 handles decrypted tunnel datagrams with the same
// native SS factory. Cleanup @0x1e6fb40 frees the wrapper but does not Close it.
func (v *Lp07FzBLx9V) mZ4KusAhuze(ctx context.Context, raw net.PacketConn) {
	packet := F.CXui61CuSLa(F.PacketOptions{Conn: raw, Cipher: v.wiy6P3.Load().FnMUny.HMkc4BgR, Users: v.jc_QPN6, Candidates: v.wZhn2Y, Invalid: v.gtKatlWm0, Manager: v.mCYkzsUd})
	defer func() { H.SlG_Y_mg("handle tunnel udp conn"); H.U9tk4Xosh9(packet) }()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if err := packet.SetReadDeadline(time.Now().Add(5 * time.Minute)); err != nil {
			return
		}
		buf := D.PoolAt50961D8.Get(4096)
		n, user, source, target, err := packet.ReadFromProxy(buf)
		if err != nil {
			D.PoolAt50961D8.Put(buf)
			return
		}
		if n == 0 {
			D.PoolAt50961D8.Put(buf)
			continue
		}
		v.dR8c1ufH6k(ctx, J.AFui66gdB{Y8YsJ7QnEI3h: packet, AIO3Qb: source, PNliwRWkrU: target, S5eGpRGjp5oo: user, AGHsUHa: buf[:n]})
	}
}
