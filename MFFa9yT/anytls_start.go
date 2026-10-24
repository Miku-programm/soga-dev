package MFFa9yT

import (
	"crypto/tls"
	"net"
	P "recovered.local/soga/CpHCdepV"
	G "recovered.local/soga/G7A7csgrs62p"
	Padding "recovered.local/soga/HHGbXWqYaYt_"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
	"recovered.local/soga/reality"
)

// aboatPBa 0x1d33480；原工厂参数和返回接口。
func aboatPBa(arg *DV_XPJ4z0) (Kz_8ajU7, error) {
	base, err := n2s6Gf9ey(arg)
	if err != nil {
		return nil, err
	}
	v := &R_n5ucb{JWN154qJ4: base, lLSuDoo: &G.GTFAWF[[32]byte, *U.HaB1NEnqcph]{}}
	v.yO3uyY.Store(Padding.DefaultAt5098340.Load())
	v.UpdateAPIData(arg.HG5k0365)
	Log.TRTbrquD(0, "create anytls server")
	return v, nil
}

// Start 0x1d33620。none、reality、普通TLS三支不可合并成 ForceCloseSSL 判断。
func (v *R_n5ucb) Start() (err error) {
	node := v.wiy6P3.Load()
	var tlsConfig *tls.Config
	var realityConfig *reality.Config
	switch node.QaZjDzw95.J6bt8AOW8fL8 {
	case "none":
	case "reality":
		p := node.QaZjDzw95
		realityConfig, err = fwKvAvIIe(v.zb8f3cHWW8v, p.Jq7YxRdb18Xx, p.CTa7WN, p.F9P9aeOu23U, p.NWbT7wk)
		if err != nil {
			return err
		}
	default:
		if !C.ForceCloseSSL() {
			tlsConfig, err = v.kaVZkL()
			if err != nil {
				return err
			}
			if err = v.dINraX_uOx5T(tlsConfig); err != nil {
				return err
			}
		}
	}
	tunnel, err := C.R3DjPWSnt()
	if err != nil {
		return err
	}
	if tunnel != nil {
		C.QitK3yi()
	}
	var listeners []net.Listener
	defer func() {
		if err != nil {
			for _, l := range listeners {
				l.Close()
			}
		}
	}()
	for _, host := range C.EdlWFAt() {
		p := v.wiy6P3.Load().QaZjDzw95
		proxy, force := C.ListenerProxyProtocolSettings()
		opts := &P.Ap9uD2{Host: host, Port: p.CgE0xkW, Network: "tcp", TLS: tlsConfig, Reality: realityConfig, ProxyProtocol: proxy, ForceProxyProtocol: force, Tunnel: tunnel}
		Log.TRTbrquD(0, "Anytls listen", opts.String())
		var l net.Listener
		l, err = opts.NewListener()
		if err != nil {
			return err
		}
		listeners = append(listeners, l)
	}
	v.qjukoB = listeners
	for _, l := range listeners {
		v.xAm4LgsEB8j(v.zb8f3cHWW8v, l)
	}
	return nil
}
