package MFFa9yT

import (
	"context"
	"net"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/EP3M1uC"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"sync"
	"time"
)

// nEFMCmrMs 对应0x1d41f00及func1—6。两种入站地址接口保留原分支。
func (b *JWN154qJ4) nEFMCmrMs(ctx context.Context, session *J1P4KNvfI_J.ZzkV7_5l, conn J1P4KNvfI_J.Ki8urSHTx44) {
	dest := session.Iayb4V
	buf := D_zCag1qTV.PoolAt50961D8.Get(4096)
	defer D_zCag1qTV.PoolAt50961D8.Put(buf)
	proxy, _ := conn.(J1P4KNvfI_J.NTe8CQ)
	var n int
	var firstAddr net.Addr
	if proxy != nil {
		_ = proxy.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, firstAddr, _ = proxy.ReadFromProxy(buf)
		_ = proxy.SetReadDeadline(time.Time{})
	} else {
		_ = conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, _ = conn.Read(buf)
		_ = conn.SetReadDeadline(time.Time{})
	}
	if n > 0 {
		host, err := b.Sniff(buf[:n], session.PayAlbdv6, 1)
		if err != nil {
			muc279Mpa6S4.TRTbrquD(-4, err)
			return
		}
		old := dest.GetHost()
		if host != "" && host != old {
			dest.SetHost(host)
			if cNT24rG4F.SniffRedirect() {
				dest = G6bHD5JA.NewDomainAddress(host, dest.GetPort())
				if err = dest.ResolveHost(b.ffju0Enjo.Load().For(conn.RemoteAddr())); err != nil {
					muc279Mpa6S4.TRTbrquD(-4, err)
					return
				}
				session.Iayb4V = dest
				muc279Mpa6S4.V8GN5Stk(-4, "sniff redirect: %v -> %v, resolved: %v", old, host, dest.GetIPs())
			}
			if b.TestAddr(dest, session.PayAlbdv6) {
				muc279Mpa6S4.V8GN5Stk(-4, "block address: %v", dest)
				return
			}
		}
	}
	route, err := EP3M1uC.CurrentRoutes().SelectOut(&EP3M1uC.AbtxqK3iZa2{BqOvrz: session.Iayb4V, YnL8s4_1IH: session.IcoyFqRBU, Kbj7uES: session.KeHzUwuXWh0, DmaA6p: session.PayAlbdv6.GetId(), WBmwgcAHw: "udp"})
	if err != nil {
		muc279Mpa6S4.TRTbrquD(-4, "select out err:", err)
		return
	}
	udpAddr, err := dest.ToUDPAddr()
	if err != nil {
		muc279Mpa6S4.TRTbrquD(-4, dest, "to udp addr err:", err)
		return
	}
	remote, err := route.DialUDP(conn.RemoteAddr(), dest)
	if err != nil {
		muc279Mpa6S4.TRTbrquD(-4, "dial udp err:", err)
		return
	}
	defer remote.Close()
	session.SetOutboundLocal(remote.LocalAddr())
	session.RReCtt("udp", session.Iayb4V, session.IksfOg)
	muc279Mpa6S4.V8GN5Stk(-4, "start udp(over tcp) relay: %v", session)
	defer func() { muc279Mpa6S4.V8GN5Stk(-4, "end udp(over tcp) relay: %v", session) }()
	if n > 0 {
		if proxy != nil {
			_, _ = remote.WriteTo(buf[:n], firstAddr)
		} else {
			_, _ = remote.WriteTo(buf[:n], udpAddr)
		}
	}
	var wg sync.WaitGroup
	wg.Add(2)
	if proxy != nil {
		go func() {
			_ = J1P4KNvfI_J.G3dXf_C86Ujq(ctx, proxy, remote, buf, session)
			_ = proxy.Close()
			_ = remote.Close()
			wg.Done()
		}()
		go func() { _ = J1P4KNvfI_J.Nw8wZD(ctx, remote, proxy); _ = remote.Close(); _ = proxy.Close(); wg.Done() }()
	} else {
		go func() {
			_ = G6bHD5JA.ZcvMMHXas8k(ctx, conn, remote, udpAddr, buf)
			_ = conn.Close()
			_ = remote.Close()
			wg.Done()
		}()
		go func() { _ = G6bHD5JA.DjO2Gp(ctx, remote, conn); _ = remote.Close(); _ = conn.Close(); wg.Done() }()
	}
	wg.Wait()
}
