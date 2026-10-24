package inbound

import (
	"context"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/routing"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"time"
)

// lbGar_ 对应0x1d40f40。保留原500ms首次读取、忽略首次读写错误和两种转发分支。
func (b *JWN154qJ4) lbGar_(ctx context.Context, session *sess.ZzkV7_5l, conn sess.Ki8urSHTx44) {
	dest := session.Iayb4V
	buffered, _ := conn.(relay.EHIIYP)
	state := relay.AsVABr{C8FYQVoBTa: ctx, GdPSmtGs_1: conn}
	n := 0
	if buffered != nil {
		_ = conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		buf, _ := buffered.ReadBuffer()
		_ = conn.SetReadDeadline(time.Time{})
		if buf != nil {
			state.C93uQCY9TN = buf.Bytes()
			n = len(state.C93uQCY9TN)
		}
	} else {
		state.C93uQCY9TN = pool.PoolAt50961D8.Get(4096)
		defer state.FreeBuf()
		_ = conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, _ = conn.Read(state.C93uQCY9TN)
		_ = conn.SetReadDeadline(time.Time{})
	}
	if n > 0 {
		host, err := b.Sniff(state.C93uQCY9TN[:n], session.PayAlbdv6, 0)
		if err != nil {
			logx.TRTbrquD(-4, err)
			return
		}
		previousHost := dest.GetHost()
		if host != "" && host != previousHost {
			dest.SetHost(host)
			if config.SniffRedirect() {
				dest = relay.NewDomainAddress(host, dest.GetPort())
				if err = dest.ResolveHost(b.ffju0Enjo.Load().For(conn.RemoteAddr())); err != nil {
					logx.TRTbrquD(-4, err)
					return
				}
				session.Iayb4V = dest
				logx.V8GN5Stk(-4, "sniff redirect: %v -> %v, resolved: %v", previousHost, host, dest.GetIPs())
			}
			if b.TestAddr(dest, session.PayAlbdv6) {
				logx.V8GN5Stk(-4, "block address: %v", dest)
				return
			}
		}
	}
	route, err := routing.CurrentRoutes().SelectOut(&routing.AbtxqK3iZa2{BqOvrz: session.Iayb4V, YnL8s4_1IH: session.IcoyFqRBU, Kbj7uES: session.KeHzUwuXWh0, DmaA6p: session.PayAlbdv6.GetId(), WBmwgcAHw: "tcp"})
	if err != nil {
		logx.TRTbrquD(-4, "select out err:", err)
		return
	}
	remote, err := route.DialTCP(conn.RemoteAddr(), dest)
	if err != nil {
		logx.TRTbrquD(-4, "dial remote", dest, "err:", err)
		return
	}
	defer remote.Close()
	session.SetOutboundLocal(remote.LocalAddr())
	session.SetOutboundRemote(remote.RemoteAddr())
	session.RReCtt("tcp", session.Iayb4V, session.IksfOg)
	if n > 0 {
		_, _ = remote.Write(state.C93uQCY9TN[:n])
	}
	state.HEfD13ahR__b = remote
	logx.V8GN5Stk(-4, "start tcp relay: %v", session)
	defer func() { logx.V8GN5Stk(-4, "end tcp relay: %v", session) }()
	if buffered != nil {
		relay.BKn06KsCTP(ctx, buffered, remote)
	} else {
		state.FreeBuf()
		relay.FvtgLyQrWX(state)
	}
}
