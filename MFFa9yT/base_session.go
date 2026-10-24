package MFFa9yT

import (
	"context"
	H "recovered.local/soga/HxKViEp2a"
	J "recovered.local/soga/J1P4KNvfI_J"
	C "recovered.local/soga/cNT24rG4F"
	L "recovered.local/soga/muc279Mpa6S4"
)

// xR518G_hBbKv: 0x1d3f6e0，闭包0x1e6a2a0—0x1e6a540。
// 这是二进制控制流的源级表示，保留两次共享禁止状态判断与用户限制。
func (b *JWN154qJ4) xR518G_hBbKv(s *J.ZzkV7_5l, conn J.Ki8urSHTx44) {
	// func1 是普通闭包：helper 里的 recover 并非直接从本 defer 调用，保持层级。
	defer func() { H.SlG_Y_mg("handle proxy conn"); _ = conn.Close() }()
	s.IcoyFqRBU = b.wiy6P3.Load().EMlXEc
	if s.KeHzUwuXWh0 == nil {
		L.V8GN5Stk(4, "session %v invalid client ip", s.E3L0l8hZ)
		return
	}
	user := s.PayAlbdv6
	reportIP := s.Uqzvq94e
	if _, _, realIP := C.SessionIPConfiguration(); realIP {
		reportIP = s.KeHzUwuXWh0.String()
	}
	// 汇编1d3f7e0保留rdi/rsi中的reportIP，rbx/rcx是计数键；两字符串实参。
	if !user.AddAliveIp(s.Uqzvq94e, reportIP) {
		L.TRTbrquD(-4, "user", user.GetId(), "alive conn limit")
		return
	}
	defer user.DelAliveIp(s.Uqzvq94e, reportIP)
	if !user.AddTCP() {
		L.TRTbrquD(-4, "user", user.GetId(), "tcp limit")
		return
	}
	defer user.DelTCP()
	if s.YlxwftjCvz_ != 2 {
		dest := s.Iayb4V
		L.TRTbrquD(-4, "user", user.GetId(), "access address:", dest)
		if !dest.HasIP() {
			if err := dest.ResolveHost(b.ffju0Enjo.Load().For(conn.RemoteAddr())); err != nil {
				L.TRTbrquD(-4, "resolve domain", dest.GetHost(), "failed:", err)
				return
			}
			L.V8GN5Stk(-4, "domain %v resolved: %v", dest.GetHost(), dest.GetIPs())
		}
		if b.TestAddr(dest, user) {
			L.V8GN5Stk(-4, "block address: %v", dest)
			return
		}
		if SessionBlocked() {
			return
		}
	}
	ctx, cancel := context.WithCancel(b.zb8f3cHWW8v)
	defer cancel()
	stop := context.AfterFunc(user.GetCtx(), func() { cancel(); _ = conn.Close() })
	defer stop()
	if b.af8f_J != nil && user.GetConnLimit() > 0 {
		closeLimited := func() { cancel(); _ = conn.Close(); L.V8GN5Stk(-4, "redis user %v ip full", user.GetId()) }
		record := func() { user.RecordAliveIp(reportIP) }
		if !b.af8f_J.StartUserConn(ctx, closeLimited, record, user, s.C66t8X) {
			L.V8GN5Stk(-4, "user %v ip full", user.GetId())
			return
		}
		defer func() { b.af8f_J.EndUserConn(user, s.C66t8X) }()
	} else {
		user.RecordAliveIp(reportIP)
	}
	if !SessionBlocked() {
		switch s.YlxwftjCvz_ {
		case 0:
			b.lbGar_(ctx, s, conn)
		case 1:
			b.nEFMCmrMs(ctx, s, conn)
		default:
			L.TRTbrquD(-4, "unknown proxy conn:", conn)
		}
	}
}
