package inbound

import (
	"net"
	P "recovered.local/soga/transport/listener"
	S22 "recovered.local/soga/protocol/ss2022"
	J "recovered.local/soga/net/sess"
	F "recovered.local/soga/protocol/ssudp"
	N "recovered.local/soga/net/limit"
	C "recovered.local/soga/config"
	T "recovered.local/soga/transport/encmux"
	Log "recovered.local/soga/util/logx"
)

// aXSbxcDLdeEY @0x1d6d140 is the original SS inbound constructor, with
// SS2022 session management or legacy IP-user candidates selected by cipher.
func aXSbxcDLdeEY(arg *DV_XPJ4z0) (Kz_8ajU7, error) {
	b, err := n2s6Gf9ey(arg)
	if err != nil {
		return nil, err
	}
	v := &Lp07FzBLx9V{JWN154qJ4: b, gtKatlWm0: N.Azb3Zr(b.zb8f3cHWW8v, C.SSInvalidAccessSettings())}
	node := arg.HG5k0365.H47AzZsYHB1N
	switch node.FnMUny.HMkc4BgR {
	case "2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm":
		v.mCYkzsUd, err = S22.DkAuKfGU(b.zb8f3cHWW8v, &node.FnMUny)
		if err != nil {
			return nil, err
		}
	default:
		v.wZhn2Y = J.YCcacra(b.zb8f3cHWW8v)
	}
	v.UpdateAPIData(arg.HG5k0365)
	Log.TRTbrquD(0, "create shadowsocks server")
	return v, nil
}

// Start @0x1d6d5c0 and cleanup @0x1e6e980. Original TCP is always tcp;
// obfs framing is applied after accepting. UDP binds the configured port,
// even when port zero gives TCP a different ephemeral port.
func (v *Lp07FzBLx9V) Start() error {
	if v.wZhn2Y != nil {
		loaded, err := v.wZhn2Y.Load(v._7le3Hqii(), v.jc_QPN6)
		if err != nil {
			Log.V8GN5Stk(4, "load ip user cache err: %v", err)
		} else if loaded {
			Log.TRTbrquD(0, "load ip user cache success")
		}
	}
	tunnel, err := C.R3DjPWSnt()
	if err != nil {
		return err
	}
	var listeners []net.Listener
	var packets []J.KsKg3yls
	failed := false
	defer func() {
		if failed {
			for _, l := range listeners {
				_ = l.Close()
			}
			for _, p := range packets {
				_ = p.Close()
			}
		}
	}()
	for _, host := range C.EdlWFAt() {
		proxy, force := C.ListenerProxyProtocolSettings()
		options := P.Ap9uD2{Host: host, Port: v.wiy6P3.Load().FnMUny.Am3QoeF, Network: "tcp", ProxyProtocol: proxy, ForceProxyProtocol: force, Tunnel: tunnel}
		listener, err := options.NewListener()
		if err != nil {
			failed = true
			return err
		}
		listeners = append(listeners, listener)
		enabled, _ := C.TunnelSettings()
		if !enabled || C.Q0SDIN() {
			raw, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(host), Port: v.wiy6P3.Load().FnMUny.Am3QoeF})
			if err != nil {
				failed = true
				return err
			}
			var packet net.PacketConn = raw
			if C.UDPProxyProtocolEnabled() {
				packet = T.SimI1bq(v.zb8f3cHWW8v, raw, force)
			}
			packetWrapper := F.CXui61CuSLa(F.PacketOptions{Conn: packet, Cipher: v.wiy6P3.Load().FnMUny.HMkc4BgR, Users: v.jc_QPN6, Candidates: v.wZhn2Y, Invalid: v.gtKatlWm0, Manager: v.mCYkzsUd})
			packets = append(packets, packetWrapper)
		}
		if tunnel != nil {
			C.QitK3yi()
		}
		node := v.wiy6P3.Load()
		switch node.FnMUny.ErKysFf2I {
		case "tls":
			Log.V8GN5Stk(0, "Shadowsocks listen %v (obfs: %v, host: %v, mode: %v)", options.String(), node.FnMUny.ErKysFf2I, node.FnMUny.DURh5VYO, C.SSObfsMode())
		case "http":
			Log.V8GN5Stk(0, "Shadowsocks listen %v (obfs: %v, path: %v, host: %v, mode: %v)", options.String(), node.FnMUny.ErKysFf2I, node.FnMUny.PzwWMqEpPguJ, node.FnMUny.DURh5VYO, C.SSObfsMode())
		default:
			Log.V8GN5Stk(0, "Shadowsocks listen %v", options.String())
		}
		switch node.FnMUny.HMkc4BgR {
		case "2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm":
			Log.TRTbrquD(0, "Shadowsocks node password:", node.FnMUny.VuhWqou)
		}
	}
	v.vLRjbS7NDeZZ = listeners
	v.sSqGyaLW = packets
	v.yER8qM3e()
	return nil
}
