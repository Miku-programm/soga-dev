package inbound

import (
	"context"
	"net"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/routing"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/util/errutil"
	"recovered.local/soga/net/sess"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
	"time"
)

// dR8c1ufH6k @0x1d40620, closures 0x1e6a5c0—0x1e6bb20. The context
// argument is unused in the original: workers and responses use b's context.
func (b *JWN154qJ4) dR8c1ufH6k(_ context.Context, packet sess.AFui66gdB) {
	owned := true
	defer func() {
		// Preserve the original nested recover call rather than broadening recovery.
		errutil.SlG_Y_mg("handle proxy pure udp")
		if owned {
			pool.PoolAt50961D8.Put(packet.AGHsUHa)
		}
	}()
	key := packet.AIO3Qb.String()
	dest := packet.PNliwRWkrU
	if len(dest.GetIPs()) == 0 {
		if err := dest.ResolveHost(b.ffju0Enjo.Load().For(packet.Y8YsJ7QnEI3h.LocalAddr())); err != nil {
			Log.V8GN5Stk(-4, "resolve udp domain %v failed: %v", dest.GetHost(), err)
			return
		}
	}
	if b.TestAddr(dest, packet.S5eGpRGjp5oo) {
		Log.V8GN5Stk(-4, "block udp address: %v", dest)
		return
	}
	udpAddr, err := dest.ToUDPAddr()
	if err != nil {
		Log.TRTbrquD(4, err)
		return
	}
	// func2 @0x1e6b9a0. udpAddr is fixed before sniffing; dest is captured
	// by reference for the error log, as in the original closure.
	write := func(remote net.PacketConn) {
		if _, err := remote.WriteTo(packet.AGHsUHa, udpAddr); err != nil {
			Log.V8GN5Stk(-4, "udp write to addr <%v> error: %v", dest.String(), err)
		} else {
			_ = remote.SetReadDeadline(time.Now().Add(time.Duration(C.UDPTimeoutMinutes()) * time.Minute))
		}
	}
	entry, loaded := b.hS8xXFfcBr.GetOrCreate(key)
	if loaded {
		select {
		case <-entry.Ready():
			if conn := entry.Conn(); conn != nil {
				write(conn)
			}
		default:
			owned = false
			go func() { // func3 @0x1e6b800; cleanup @0x1e6b920.
				defer func() { errutil.SlG_Y_mg("wait udp nat entry"); pool.PoolAt50961D8.Put(packet.AGHsUHa) }()
				select {
				case <-b.zb8f3cHWW8v.Done():
				case <-entry.Ready():
					if conn := entry.Conn(); conn != nil {
						write(conn)
					}
				}
			}()
		}
		return
	}
	owned = false
	go func() { // func4 @0x1e6a5c0.
		complete := false
		defer func() { // func4.1 @0x1e6b700.
			errutil.SlG_Y_mg("dial udp")
			if !complete {
				b.hS8xXFfcBr.Delete(key)
				entry.Complete(nil)
			}
			pool.PoolAt50961D8.Put(packet.AGHsUHa)
		}()
		session := sess.TgFd4ag1(packet)
		session.IcoyFqRBU = b.wiy6P3.Load().EMlXEc
		// The original sniff helper returns no status. Its early returns do not
		// terminate this goroutine, including a post-sniff block result.
		b.nativeUDPsniff(packet, &dest, session)
		route, err := routing.CurrentRoutes().SelectOut(&routing.AbtxqK3iZa2{BqOvrz: session.Iayb4V, YnL8s4_1IH: session.IcoyFqRBU, Kbj7uES: session.KeHzUwuXWh0, DmaA6p: session.PayAlbdv6.GetId(), WBmwgcAHw: "udp"})
		if err != nil {
			Log.TRTbrquD(-4, "select out err:", err)
			return
		}
		remote, err := route.DialUDP(packet.Y8YsJ7QnEI3h.LocalAddr(), packet.PNliwRWkrU)
		if err != nil {
			Log.TRTbrquD(-4, "new remote udp err:", err)
			return
		}
		session.SetOutboundLocal(remote.LocalAddr())
		session.RReCtt("udp", session.Iayb4V, session.IksfOg)
		Log.V8GN5Stk(-4, "start user %v udp relay: %v -> %v", packet.S5eGpRGjp5oo.GetId(), key, dest)
		go func() { // func4.3 @0x1e6ad80.
			ctx, cancel := context.WithCancel(b.zb8f3cHWW8v)
			stop := context.AfterFunc(packet.S5eGpRGjp5oo.GetCtx(), func() { cancel(); _ = remote.Close() })
			defer func() { stop(); cancel(); b.hS8xXFfcBr.Delete(key); _ = remote.Close() }()
			_ = sess.Jv9wkYGl(ctx, remote, packet.AIO3Qb, packet.Y8YsJ7QnEI3h, packet.S5eGpRGjp5oo)
			Log.V8GN5Stk(-4, "end user %v udp relay: %v -> %v", packet.S5eGpRGjp5oo.GetId(), key, dest)
		}()
		write(remote)
		complete = true
		entry.Complete(remote)
	}()
}

// nativeUDPsniff is the source name for 0x1e6b200 (func4.2).
func (b *JWN154qJ4) nativeUDPsniff(packet sess.AFui66gdB, dest **relay.AHU2MRmjzQu, session *sess.ZzkV7_5l) {
	host, err := b.Sniff(packet.AGHsUHa, packet.S5eGpRGjp5oo, 1)
	if err != nil {
		Log.TRTbrquD(-4, err)
		return
	}
	old := (*dest).GetHost()
	if host == "" || host == old {
		return
	}
	(*dest).SetHost(host)
	if C.SniffRedirect() {
		*dest = relay.NewDomainAddress(host, (*dest).GetPort())
		if err = (*dest).ResolveHost(b.ffju0Enjo.Load().For(packet.Y8YsJ7QnEI3h.LocalAddr())); err != nil {
			Log.TRTbrquD(-4, err)
			return
		}
		session.Iayb4V = *dest
		Log.V8GN5Stk(-4, "sniff redirect: %v -> %v, resolved: %v", old, host, (*dest).GetIPs())
	}
	if b.TestAddr(*dest, packet.S5eGpRGjp5oo) {
		Log.V8GN5Stk(-4, "block udp address: %v", *dest)
		return
	}
}
