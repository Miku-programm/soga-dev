package inbound

import (
	"context"
	"encoding/binary"
	"io"
	"net"
	G "recovered.local/soga/net/relay"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	N "recovered.local/soga/protocol/anytlsconn"
	C "recovered.local/soga/config"
	Tunnel "recovered.local/soga/transport/encmux"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/user"
	UOT "recovered.local/soga/net/uot"
	Session "recovered.local/soga/protocol/anytls"
	"time"
)

// tFdBaee6T @0x1d34ee0。保留原认证失败消费、30秒读超时与填充分支。
func (v *R_n5ucb) tFdBaee6T(conn net.Conn) {
	defer func() { H.SlG_Y_mg("handle anytls"); conn.Close(); Log.TRTbrquD(-4, "handle anytls end"); v.bj10P7X() }()
	Log.TRTbrquD(-4, "start handle client", conn.RemoteAddr(), "connection")
	if enabled, _ := C.TunnelSettings(); enabled && !C.Q0SDIN() {
		if tunnel, ok := conn.(*Tunnel.Ql1qrX); ok {
			kind, err := tunnel.GetProxyType()
			if err != nil {
				Log.TRTbrquD(-4, "read tunnel err:", err)
				return
			}
			if kind != 1 {
				return
			}
		}
	}
	v.sbbdX3aT6E()
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	var auth [32]byte
	if _, err := io.ReadFull(conn, auth[:]); err != nil {
		Log.V8GN5Stk(-4, "read auth err: %v", err)
		G.Ad2VIxY(conn)
		return
	}
	user, ok := v.lLSuDoo.Get(auth)
	if !ok {
		Log.V8GN5Stk(-4, "no auth: %v", auth)
		if len(v.dZfoV0aXHW) > 0 && v.lAmttzTO8b > 0 {
			conn.SetReadDeadline(time.Time{})
			v.xx3H26VL(conn, auth[:])
		} else {
			G.Ad2VIxY(conn)
		}
		return
	}
	stop := context.AfterFunc(user.GetCtx(), func() { conn.Close() })
	defer stop()
	var length [2]byte
	if _, err := io.ReadFull(conn, length[:]); err != nil {
		Log.V8GN5Stk(-4, "read padding len err: %v", err)
		return
	}
	if size := int(binary.BigEndian.Uint16(length[:])); size != 0 {
		if err := G.KlTAqrF42mR(conn, size); err != nil {
			Log.V8GN5Stk(-4, "discard padding err: %v", err)
			return
		}
	}
	conn.SetReadDeadline(time.Time{})
	session := Session.NewServerFromInline(conn, func(stream *Session.CCXLCurYiT7) { v.handleAnyTLSStream(user, stream) }, &v.yO3uyY)
	defer session.Close()
	session.Run()
}

// handleAnyTLSStream 是原 func3 @0x1e69780 的具名恢复。
func (v *R_n5ucb) handleAnyTLSStream(user *U.HaB1NEnqcph, stream *Session.CCXLCurYiT7) {
	defer stream.Close()
	if err := stream.HandshakeSuccess(); err != nil {
		Log.V8GN5Stk(-4, "handshake err: %v", err)
		return
	}
	counter := &J.Mdn09F6aMrpa{DrpXao: stream}
	counter.SetUser(user)
	dest, err := G.SzhW4tOzUZg(counter)
	if err != nil {
		Log.V8GN5Stk(-4, "read addr err: %v", err)
		return
	}
	if UOT.Vd_ooUdqfJTg(dest.GetHost()) {
		framed, err := UOT.SACWxFiDLx22(counter, dest.GetHost(), user, v.ffju0Enjo.Load().For(counter.LocalAddr()))
		if err != nil {
			Log.V8GN5Stk(-4, "new packet conn err: %v", err)
			return
		}
		v.xR518G_hBbKv(J.MHLL7QOie(framed), framed)
	} else {
		framed := N.FromAnyTLSInline(counter, user, dest)
		v.xR518G_hBbKv(J.MHLL7QOie(framed), framed)
	}
}

// 原 accept goroutine @0x1e69200 任意 Accept 错误均结束。
func (v *R_n5ucb) xAm4LgsEB8j(ctx context.Context, listener net.Listener) {
	go func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				select {
				case <-ctx.Done():
				default:
					Log.TRTbrquD(4, "accept err:", err)
				}
				return
			}
			go v.tFdBaee6T(conn)
		}
	}()
}
