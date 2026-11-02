package inbound

import (
	"context"
	"net"
	P "recovered.local/soga/transport/listener"
	H "recovered.local/soga/util/errutil"
	C "recovered.local/soga/config"
	T "recovered.local/soga/transport/encmux"
	Log "recovered.local/soga/util/logx"
	"reflect"
)

// z0MUgmyQ @0x1d70000 traverses only the original GetRawConn interface.
// The original received connection remains the stream passed to both relays.
func (v *Lp07FzBLx9V) z0MUgmyQ(ctx context.Context, conn net.Conn) {
	defer func() { H.SlG_Y_mg("handle tunnel conn"); v.bj10P7X(); _ = conn.Close() }()
	var tunnel *T.Ql1qrX
	current := conn
	for current != nil {
		if c, ok := current.(*T.Ql1qrX); ok {
			tunnel = c
			break
		}
		if c, ok := current.(interface{ GetRawConn() net.Conn }); ok {
			current = c.GetRawConn()
		} else {
			break
		}
	}
	if tunnel == nil {
		Log.TRTbrquD(8, "unknown conn type:", reflect.TypeOf(conn))
		_ = conn.Close()
		return
	}
	v.sbbdX3aT6E()
	kind, err := tunnel.GetProxyType()
	if err != nil {
		Log.TRTbrquD(-4, "read tunnel err:", err)
		return
	}
	switch kind {
	case 1:
		v.nlJ8YjzjZ(ctx, conn)
	case 2:
		v.mZ4KusAhuze(ctx, P.NewTunnelPacketFromStream(conn))
	case 3:
		Log.TRTbrquD(4, "not support mux proxy yet")
	default:
		Log.TRTbrquD(4, "unknown tunnel proxy type:", kind)
	}
}

// xAm4LgsEB8j @0x1d6e580 and loop @0x1e6ec60. Tunnel mode is captured
// once before the goroutine. On the original session-blocked branch the
// accepted connection is discarded without Close; that behavior is retained.
func (v *Lp07FzBLx9V) xAm4LgsEB8j(listener net.Listener) {
	enabled, _ := C.TunnelSettings()
	tunnel := enabled
	go func() {
		defer listener.Close()
		for {
			select {
			case <-v.zb8f3cHWW8v.Done():
				return
			default:
			}
			conn, err := listener.Accept()
			if err != nil {
				Log.TRTbrquD(-4, "accept ss tcp err:", err)
				return
			}
			if tunnel {
				go v.z0MUgmyQ(v.zb8f3cHWW8v, conn)
			} else {
				go v.nlJ8YjzjZ(v.zb8f3cHWW8v, conn)
			}
		}
	}()
}

// yER8qM3e @0x1d6e480 starts the configured native listeners and cache saver.
func (v *Lp07FzBLx9V) yER8qM3e() {
	for _, listener := range v.vLRjbS7NDeZZ {
		v.xAm4LgsEB8j(listener)
	}
	for _, packet := range v.sSqGyaLW {
		v.wXPmSs2(packet)
	}
	if v.wZhn2Y != nil && C.IPUserCacheSaveEnabled() {
		go v.saveShadowsocksIPCache()
	}
}
