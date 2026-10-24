package MFFa9yT

import (
	"context"
	"net"
	H "recovered.local/soga/HxKViEp2a"
	J "recovered.local/soga/J1P4KNvfI_J"
	C "recovered.local/soga/cNT24rG4F"
	V "recovered.local/soga/gk1z9v"
	Log "recovered.local/soga/muc279Mpa6S4"
	M "recovered.local/soga/nXZ99wahk"
	"time"
)

// tFdBaee6T @0x1da35e0及闭包0x1e743a0—0x1e745a0。
// 原入口为VMess自有握手和数据帧处理，随后调用公共用户策略与转发层。
func (v *K4avFiYOPi) tFdBaee6T(conn net.Conn) {
	defer func() {
		H.SlG_Y_mg("handle vmess")
		_ = conn.Close()
		Log.TRTbrquD(-4, "handle vmess end")
		v.bj10P7X()
	}()
	if enabled, _ := C.TunnelSettings(); enabled && !C.Q0SDIN() {
		if _, err := conn.Read(nil); err != nil {
			Log.TRTbrquD(-4, "read tunnel err:", err)
			return
		}
	}
	Log.TRTbrquD(-4, "start handle client", conn.RemoteAddr(), "connection")
	auth, fallback := v.s3mqrqd_(conn)
	if auth == nil {
		if fallback != nil {
			v.xx3H26VL(conn, fallback)
		}
		return
	}
	auth.Xwwpafq.AddUp(16)
	counter := &J.Mdn09F6aMrpa{DrpXao: conn}
	counter.SetUser(auth.Xwwpafq)
	conn = counter
	framed, err := V.EAQVlR(conn, auth)
	if err != nil {
		Log.TRTbrquD(-4, "new vmess conn err:", err)
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

// xAm4LgsEB8j @0x1da1a60/0x1e73820。原实现遇到Accept错误即退出。
func (v *K4avFiYOPi) xAm4LgsEB8j(ctx context.Context, listener net.Listener) {
	go func() {
		defer func() { _ = listener.Close() }()
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

// yER8qM3e @0x1da1b60，缓存持久化开关只影响保存任务。
func (v *K4avFiYOPi) yER8qM3e(ctx context.Context) {
	for _, listener := range v.l8Uq4M {
		v.xAm4LgsEB8j(ctx, listener)
	}
	if v.l_ZoZQDxG7 != nil && C.IPUserCacheSaveEnabled() {
		go v.saveVmessIPCache()
	}
}
