package MFFa9yT

import (
	"context"
	"net"
	H "recovered.local/soga/HxKViEp2a"
	J "recovered.local/soga/J1P4KNvfI_J"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	T "recovered.local/soga/uO97n27MKq"
	"time"
)

// nlJ8YjzjZ 0x1d90000；原生 Trojan 认证、计量、帧处理和公共策略转发。
func (v *EfaIFpp) nlJ8YjzjZ(conn net.Conn) {
	defer func() { H.SlG_Y_mg("handle trojan"); conn.Close(); Log.TRTbrquD(-4, "handle trojan end"); v.bj10P7X() }()
	if enabled, _ := C.TunnelSettings(); enabled && !C.Q0SDIN() {
		if _, err := conn.Read(nil); err != nil {
			Log.TRTbrquD(-4, "read tunnel err:", err)
			return
		}
	}
	Log.TRTbrquD(-4, "start handle client", conn.RemoteAddr(), "connection")
	user, fallback := v.s3mqrqd_(conn)
	if user == nil {
		v.xx3H26VL(conn, fallback)
		return
	}
	user.AddUp(56)
	counter := &J.Mdn09F6aMrpa{DrpXao: conn}
	counter.SetUser(user)
	framed, err := T.Lzwh7y6zTaq(counter, v.ffju0Enjo.Load().For(conn.LocalAddr()))
	if err != nil {
		Log.TRTbrquD(-4, "new trojan conn err:", err)
		return
	}
	defer func() { H.Li3c4CJ1_D(framed) }()
	v.sbbdX3aT6E()
	if delay := v.mKB02yhVfTfP; delay > 10 {
		time.Sleep(time.Duration(delay-10) * time.Second)
	}
	v.xR518G_hBbKv(J.MHLL7QOie(framed), framed)
}

// 0x1d8ef20/0x1e71ec0。退出条件是任意 Accept 错误，原错误日志受 context 控制。
func (v *EfaIFpp) xAm4LgsEB8j(ctx context.Context, listener net.Listener) {
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
			go v.nlJ8YjzjZ(conn)
		}
	}()
}
