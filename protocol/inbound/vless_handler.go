package inbound

import (
	"context"
	"github.com/gofrs/uuid"
	"net"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	C "recovered.local/soga/config"
	T "recovered.local/soga/transport/encmux"
	V "recovered.local/soga/protocol/vlessvision"
	Log "recovered.local/soga/util/logx"
	M "recovered.local/soga/protocol/mux"
	"time"
)

// 原1d98520及闭包1e72c40..1e72d00；不是替代协议实现。
func (v *KHonGaqbIy_6) tFdBaee6T(conn net.Conn) {
	defer func() { H.SlG_Y_mg("handle vless"); conn.Close(); Log.TRTbrquD(-4, "handle vless end"); v.bj10P7X() }()
	Log.TRTbrquD(-4, "start handle client", conn.RemoteAddr(), "connection")
	if enabled, _ := C.TunnelSettings(); enabled {
		if tunnel, ok := conn.(*T.Ql1qrX); ok {
			typ, err := tunnel.GetProxyType()
			if err != nil {
				Log.TRTbrquD(-4, "read tunnel err:", err)
				return
			}
			if typ != 1 {
				return
			}
		}
	}
	if v.i420mh4Yr != nil {
		c, err := v.i420mh4Yr.Handshake(conn, nil)
		if err != nil {
			Log.V8GN5Stk(-4, "vless decryption handshake err: %v", err)
			return
		}
		conn = c
	}
	var header [17]byte
	conn.SetReadDeadline(time.Now().Add(120 * time.Second))
	var n int
	var err error
	for {
		n, err = conn.Read(header[:])
		if n > 0 || err != nil {
			break
		}
	}
	conn.SetReadDeadline(time.Time{})
	if err != nil {
		return
	}
	if n != 17 {
		v.xx3H26VL(conn, header[:n])
		return
	}
	var id uuid.UUID
	copy(id[:], header[1:])
	user, ok := v.vU51MPl_at_.Get(id)
	if !ok {
		Log.TRTbrquD(-4, "vless auth failed:", [16]byte(id))
		v.xx3H26VL(conn, header[:])
		return
	}
	user.AddUp(17)
	counter := &J.Mdn09F6aMrpa{DrpXao: conn}
	counter.SetUser(user)
	framed, err := V.V7ZRKYAK(counter, v.wiy6P3.Load().Ck6DxjB.NMzMUQ2, header[0])
	if err != nil {
		Log.TRTbrquD(-4, "new vless conn err:", err)
		return
	}
	defer func() { H.Li3c4CJ1_D(framed) }()
	v.sbbdX3aT6E()
	if delay := v.mKB02yhVfTfP; delay > 10 {
		time.Sleep(time.Duration(delay-10) * time.Second)
	}
	if framed.NetworkType() == 2 {
		session := M.KJwDQE2y3(v.zb8f3cHWW8v, framed, v.ffju0Enjo.Load().For(framed.LocalAddr()))
		defer session.Close()
		for {
			sub, ok := session.Accept()
			if !ok {
				return
			}
			metadata := J.MHLL7QOie(sub)
			go v.xR518G_hBbKv(metadata, sub)
		}
	}
	v.xR518G_hBbKv(J.MHLL7QOie(framed), framed)
}
func (v *KHonGaqbIy_6) xAm4LgsEB8j(ctx context.Context, listener net.Listener) { //1d97620,1e72740
	go func() {
		defer func() { listener.Close() }()
		for {
			c, err := listener.Accept()
			if err != nil {
				select {
				case <-ctx.Done():
				default:
					Log.TRTbrquD(4, "accept err:", err)
				}
				return
			}
			go v.tFdBaee6T(c)
		}
	}()
}
