package inbound

import (
	"context"
	"net"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/routing"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"sync"
	"time"
)

// nEFMCmrMs 对应0x1d41f00及func1—6。两种入站地址接口保留原分支。
func (b *JWN154qJ4) nEFMCmrMs(ctx context.Context, session *sess.ZzkV7_5l, conn sess.Ki8urSHTx44) {
	dest := session.Iayb4V
	buf := pool.PoolAt50961D8.Get(4096)
	defer pool.PoolAt50961D8.Put(buf)
	proxy, _ := conn.(sess.NTe8CQ)
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
			logx.TRTbrquD(-4, err)
			return
		}
		old := dest.GetHost()
		if host != "" && host != old {
			dest.SetHost(host)
			if config.SniffRedirect() {
				access := dest.AccessHost()
				dest = relay.NewDomainAddress(host, dest.GetPort())
				dest.SetAccessHost(access)
				if err = dest.ResolveHost(b.ffju0Enjo.Load().For(conn.RemoteAddr())); err != nil {
					logx.TRTbrquD(-4, err)
					return
				}
				session.Iayb4V = dest
				logx.V8GN5Stk(-4, "sniff redirect: %v -> %v, resolved: %v", old, host, dest.GetIPs())
			}
			if b.TestAddr(dest, session.PayAlbdv6) {
				logx.V8GN5Stk(-4, "block address: %v", dest)
				return
			}
		}
	}
	route, err := routing.CurrentRoutes().SelectOut(&routing.AbtxqK3iZa2{BqOvrz: session.Iayb4V, YnL8s4_1IH: session.IcoyFqRBU, Kbj7uES: session.KeHzUwuXWh0, DmaA6p: session.PayAlbdv6.GetId(), WBmwgcAHw: "udp"})
	if err != nil {
		logx.TRTbrquD(-4, "select out err:", err)
		return
	}
	udpAddr, err := dest.ToUDPAddr()
	if err != nil {
		logx.TRTbrquD(-4, dest, "to udp addr err:", err)
		return
	}
	remote, err := route.DialUDP(conn.RemoteAddr(), dest)
	if err != nil {
		logx.TRTbrquD(-4, "dial udp err:", err)
		return
	}
	defer remote.Close()
	session.SetOutboundLocal(remote.LocalAddr())
	session.RReCtt("udp", session.Iayb4V, session.IksfOg)
	logx.V8GN5Stk(-4, "start udp(over tcp) relay: %v", session)
	defer func() { logx.V8GN5Stk(-4, "end udp(over tcp) relay: %v", session) }()
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
			_ = sess.G3dXf_C86Ujq(ctx, proxy, remote, buf, session)
			_ = proxy.Close()
			_ = remote.Close()
			wg.Done()
		}()
		go func() { _ = sess.Nw8wZD(ctx, remote, proxy); _ = remote.Close(); _ = proxy.Close(); wg.Done() }()
	} else {
		go func() {
			_ = relay.ZcvMMHXas8k(ctx, conn, remote, udpAddr, buf)
			_ = conn.Close()
			_ = remote.Close()
			wg.Done()
		}()
		go func() { _ = relay.DjO2Gp(ctx, remote, conn); _ = remote.Close(); _ = conn.Close(); wg.Done() }()
	}
	wg.Wait()
}
