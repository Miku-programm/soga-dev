package MFFa9yT

import (
	"context"
	"net"
	H "recovered.local/soga/HxKViEp2a"
	J "recovered.local/soga/J1P4KNvfI_J"
	S "recovered.local/soga/UQ7OErS0s"
	A "recovered.local/soga/ad5KEbw_EzWk"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
	"strconv"
	"time"
)

// 新增源码组织辅助：原 1d7fe40/1d7f780/1d8b500 三处相同内联赋值。
// 不计为原样本保留函数。
func (v *Ze0fOaAko) ssrServerInfo(u *U.LJJkaD) *A.L7E7H2aW8KVs {
	info := &A.L7E7H2aW8KVs{NxYA16xMfF: u.GetSinglePortType(), MY3lagqY: u.GetPassword(), L6BrlCXd: u.GetMethod(), C8Qsa8G8XMw0: u.GetProtocol(), UNRudV1377x9: u.GetObfs(), ZlsyKgz4EI: v.tpQEL_qKogl, N3pMXuwX85c: v.jTyGfNkAw, RLXBf_luGj: v.ktET5DFha}
	if !v.g4fBrm0Bxl {
		info.Uklg6401 = u.GetId()
	}
	return info
}

// nlJ8YjzjZ @1d7fe40，cleanup 1e71260/1e71220，device defer 1e711e0。
// 原 context 参数不参与该函数；共用转发使用入站自身 context。
func (v *Ze0fOaAko) nlJ8YjzjZ(_ context.Context, conn net.Conn, u *U.LJJkaD) {
	defer func() {
		H.SlG_Y_mg("handle ssr tcp")
		_ = conn.Close()
		v.bj10P7X()
		Log.TRTbrquD(-4, "handle ssr tcp conn end")
	}()
	Log.TRTbrquD(-4, "start handle ssr tcp conn, client:", conn.RemoteAddr())
	if u == nil {
		Log.TRTbrquD(8, "单端口承载用户不存在")
		return
	}
	info := v.ssrServerInfo(u)
	_ = conn.SetReadDeadline(time.Now().Add(120 * time.Second))
	stream, err := S.RfxaTYLqLUJ(conn, info)
	_ = conn.SetReadDeadline(time.Time{})
	if err != nil {
		Log.TRTbrquD(-4, "new ssr conn err:", err)
		return
	}
	defer func() { H.Li3c4CJ1_D(stream) }()
	actual := stream.GetUser().(*U.LJJkaD)
	session := J.MHLL7QOie(stream)
	if C.SSRCIDLimitEnabled() && (info.NxYA16xMfF == 0 || info.NxYA16xMfF == 2) {
		cid := stream.CID()
		if !actual.AddDevice(cid) {
			Log.V8GN5Stk(-4, "user %v limit device num", actual.GetId())
			return
		}
		defer actual.DelDevice(cid)
		session.C66t8X = strconv.FormatInt(cid, 10)
	}
	v.sbbdX3aT6E()
	v.xR518G_hBbKv(session, stream)
}

// kOoHOawg4u @1d7ed80；真正循环 1e70620。模式读取发生在 goroutine 内。
func (v *Ze0fOaAko) kOoHOawg4u(ctx context.Context, l net.Listener, u *U.LJJkaD) {
	go func() {
		enabled, _ := C.TunnelSettings()
		tunnel := enabled && !C.Q0SDIN()
		defer l.Close()
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			conn, err := l.Accept()
			if err != nil {
				Log.TRTbrquD(-4, "accept ssr tcp err:", err)
				return
			}
			if tunnel {
				go v.z0MUgmyQ(ctx, conn, u)
			} else {
				go v.nlJ8YjzjZ(ctx, conn, u)
			}
		}
	}()
}
