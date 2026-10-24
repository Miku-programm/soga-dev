package inbound

import (
	"context"
	"errors"
	"net"
	G "recovered.local/soga/net/relay"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	F "recovered.local/soga/protocol/ssudp"
	V "recovered.local/soga/transport/muxerr"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/net/uot"
	"time"
)

// nlJ8YjzjZ @0x1d6f5c0 with original cleanup closures 0x1e6f900/0x1e6f980.
// The context argument is unused in this entry; the shared relay uses the
// inbound's context. Native SS authentication precedes the common policy layer.
func (v *Lp07FzBLx9V) nlJ8YjzjZ(_ context.Context, conn net.Conn) {
	defer func() { H.SlG_Y_mg("handle ss"); _ = conn.Close(); v.bj10P7X(); Log.TRTbrquD(-4, "handle ss end") }()
	Log.TRTbrquD(-4, "start handle client", conn.RemoteAddr(), "connection")
	remote, err := G.JY6gu_Rqfz(conn.RemoteAddr())
	if err != nil {
		Log.TRTbrquD(-4, err)
		return
	}
	ip := remote.IP.String()
	if v.gtKatlWm0.IsForbidden(ip) {
		Log.V8GN5Stk(-4, "forbidden ip %v access", ip)
		return
	}
	_ = conn.SetReadDeadline(time.Now().Add(120 * time.Second))
	node := v.wiy6P3.Load()
	framed, err := F.FciVp4yFHL(F.StreamOptions{Conn: conn, Cipher: node.FnMUny.HMkc4BgR, Obfs: string(node.FnMUny.ErKysFf2I), Mode: C.SSObfsMode(), Path: node.FnMUny.PzwWMqEpPguJ, Host: node.FnMUny.DURh5VYO, Candidates: v.wZhn2Y, Invalid: v.gtKatlWm0, Users: v.jc_QPN6, Manager: v.mCYkzsUd})
	_ = conn.SetReadDeadline(time.Time{})
	if err != nil {
		if errors.Is(err, V.ErrorAt5098410()) {
			v.gtKatlWm0.Invalid(ip)
		}
		Log.TRTbrquD(-4, "new ss tcp conn err:", err)
		return
	}
	defer func() { H.Li3c4CJ1_D(framed) }()
	v.gtKatlWm0.Valid(ip)
	v.sbbdX3aT6E()
	if n := v.mKB02yhVfTfP; n > 10 {
		time.Sleep(time.Duration(n-10) * time.Second)
	}
	target := framed.GetDestAddr()
	stream := framed
	if U.Vd_ooUdqfJTg(target.GetHost()) {
		stream, err = U.SACWxFiDLx22(framed, target.GetHost(), framed.GetUser(), v.ffju0Enjo.Load().For(conn.LocalAddr()))
		if err != nil {
			Log.TRTbrquD(-4, "new ss uot conn err:", err)
			return
		}
	}
	v.xR518G_hBbKv(J.MHLL7QOie(stream), stream)
}
