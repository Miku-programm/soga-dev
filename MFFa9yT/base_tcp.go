package MFFa9yT

import (
	"context"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/EP3M1uC"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"time"
)

// lbGar_ 对应0x1d40f40。保留原500ms首次读取、忽略首次读写错误和两种转发分支。
func (b *JWN154qJ4) lbGar_(ctx context.Context, session *J1P4KNvfI_J.ZzkV7_5l, conn J1P4KNvfI_J.Ki8urSHTx44) {
	dest := session.Iayb4V
	buffered, _ := conn.(G6bHD5JA.EHIIYP)
	state := G6bHD5JA.AsVABr{C8FYQVoBTa: ctx, GdPSmtGs_1: conn}
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
		state.C93uQCY9TN = D_zCag1qTV.PoolAt50961D8.Get(4096)
		defer state.FreeBuf()
		_ = conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, _ = conn.Read(state.C93uQCY9TN)
		_ = conn.SetReadDeadline(time.Time{})
	}
	if n > 0 {
		host, err := b.Sniff(state.C93uQCY9TN[:n], session.PayAlbdv6, 0)
		if err != nil {
			muc279Mpa6S4.TRTbrquD(-4, err)
			return
		}
		previousHost := dest.GetHost()
		if host != "" && host != previousHost {
			dest.SetHost(host)
			if cNT24rG4F.SniffRedirect() {
				dest = G6bHD5JA.NewDomainAddress(host, dest.GetPort())
				if err = dest.ResolveHost(b.ffju0Enjo.Load().For(conn.RemoteAddr())); err != nil {
					muc279Mpa6S4.TRTbrquD(-4, err)
					return
				}
				session.Iayb4V = dest
				muc279Mpa6S4.V8GN5Stk(-4, "sniff redirect: %v -> %v, resolved: %v", previousHost, host, dest.GetIPs())
			}
			if b.TestAddr(dest, session.PayAlbdv6) {
				muc279Mpa6S4.V8GN5Stk(-4, "block address: %v", dest)
				return
			}
		}
	}
	route, err := EP3M1uC.CurrentRoutes().SelectOut(&EP3M1uC.AbtxqK3iZa2{BqOvrz: session.Iayb4V, YnL8s4_1IH: session.IcoyFqRBU, Kbj7uES: session.KeHzUwuXWh0, DmaA6p: session.PayAlbdv6.GetId(), WBmwgcAHw: "tcp"})
	if err != nil {
		muc279Mpa6S4.TRTbrquD(-4, "select out err:", err)
		return
	}
	remote, err := route.DialTCP(conn.RemoteAddr(), dest)
	if err != nil {
		muc279Mpa6S4.TRTbrquD(-4, "dial remote", dest, "err:", err)
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
	muc279Mpa6S4.V8GN5Stk(-4, "start tcp relay: %v", session)
	defer func() { muc279Mpa6S4.V8GN5Stk(-4, "end tcp relay: %v", session) }()
	if buffered != nil {
		G6bHD5JA.BKn06KsCTP(ctx, buffered, remote)
	} else {
		state.FreeBuf()
		G6bHD5JA.FvtgLyQrWX(state)
	}
}
