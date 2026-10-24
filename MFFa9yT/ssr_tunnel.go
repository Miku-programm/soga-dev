package MFFa9yT

import (
	"context"
	"io"
	"net"
	P "recovered.local/soga/CpHCdepV"
	D "recovered.local/soga/D_zCag1qTV"
	H "recovered.local/soga/HxKViEp2a"
	J "recovered.local/soga/J1P4KNvfI_J"
	S "recovered.local/soga/UQ7OErS0s"
	C "recovered.local/soga/cNT24rG4F"
	T "recovered.local/soga/g0MlUOhvO"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
	"reflect"
	"time"
)

// z0MUgmyQ @1d7f0a0，defer1e710c0。SSR 在 GetProxyType 成功后才增加连接状态。
func (v *Ze0fOaAko) z0MUgmyQ(ctx context.Context, conn net.Conn, u *U.LJJkaD) {
	defer func() { H.SlG_Y_mg("handle tunnel conn"); v.bj10P7X(); _ = conn.Close() }()
	var tunnel *T.Ql1qrX
	for current := conn; current != nil; {
		if t, ok := current.(*T.Ql1qrX); ok {
			tunnel = t
			break
		}
		if t, ok := current.(interface{ GetRawConn() net.Conn }); ok {
			current = t.GetRawConn()
		} else {
			break
		}
	}
	if tunnel == nil {
		Log.TRTbrquD(8, "unknown conn type:", reflect.TypeOf(conn))
		_ = conn.Close()
		return
	}
	kind, err := tunnel.GetProxyType()
	if err != nil {
		Log.TRTbrquD(-4, "read tunnel err:", err)
		return
	}
	v.sbbdX3aT6E()
	switch kind {
	case 1:
		v.nlJ8YjzjZ(ctx, conn, u)
	case 2:
		v.mZ4KusAhuze(ctx, P.NewTunnelPacketFromStream(conn), u)
	case 3:
		Log.TRTbrquD(4, "not support mux proxy yet")
	default:
		Log.TRTbrquD(4, "unknown tunnel proxy type:", kind)
	}
}

// mZ4KusAhuze @1d7f780 / defer1e71160。原 EOF 使用直接相等而非 errors.Is。
func (v *Ze0fOaAko) mZ4KusAhuze(ctx context.Context, pc net.PacketConn, u *U.LJJkaD) {
	stream, err := S.INOdx8(pc, v.ssrServerInfo(u))
	if err != nil {
		Log.TRTbrquD(-4, "new tunnel ssr udp conn err:", err)
		return
	}
	defer func() { H.SlG_Y_mg("handle tunnel udp conn"); H.PpnNLr_pf7He(stream) }()
	timeout := time.Duration(C.UDPTimeoutMinutes()) * time.Minute
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if err := stream.SetReadDeadline(time.Now().Add(timeout)); err != nil {
			return
		}
		buf := D.PoolAt50961D8.Get(4096)
		n, user, client, target, err := stream.ReadFromProxy(buf)
		if err != nil {
			if err != io.EOF {
				Log.TRTbrquD(-4, "read ssr udp err:", err)
			}
			D.PoolAt50961D8.Put(buf)
			return
		}
		if n == 0 {
			D.PoolAt50961D8.Put(buf)
			continue
		}
		if user == nil {
			user = u
		}
		v.dR8c1ufH6k(ctx, J.AFui66gdB{Y8YsJ7QnEI3h: stream, AIO3Qb: client, PNliwRWkrU: target, S5eGpRGjp5oo: user, AGHsUHa: buf[:n]})
	}
}
