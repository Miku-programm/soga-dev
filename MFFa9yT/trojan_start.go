package MFFa9yT

import (
	"crypto/tls"
	"net"
	P "recovered.local/soga/CpHCdepV"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/reality"
)

// f1mPzegGt 0x1d8e680；原工厂参数和返回接口。
func f1mPzegGt(arg *DV_XPJ4z0) (Kz_8ajU7, error) {
	base, err := n2s6Gf9ey(arg)
	if err != nil {
		return nil, err
	}
	v := &EfaIFpp{JWN154qJ4: base}
	v.UpdateAPIData(arg.HG5k0365)
	Log.TRTbrquD(0, "create trojan server")
	return v, nil
}

// Start 0x1d8e7a0。none、reality、普通TLS三支不可合并成 ForceCloseSSL 判断。
func (v *EfaIFpp) Start() (err error) {
	node := v.wiy6P3.Load()
	var tlsConfig *tls.Config
	var realityConfig *reality.Config
	switch node.JhIYwEukjd.AJjaqrBjj {
	case "none":
	case "reality":
		p := node.JhIYwEukjd
		realityConfig, err = fwKvAvIIe(v.zb8f3cHWW8v, p.HIlblJq1, p.SmInTc, p.ZxabzIRCNeV, p.TJ8rFSEH)
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
		p := v.wiy6P3.Load().JhIYwEukjd
		proxy, force := C.ListenerProxyProtocolSettings()
		opts := &P.Ap9uD2{Host: host, Port: p.GbD_2PU, Network: p.T4cuhG7MTbG, TLS: tlsConfig, Reality: realityConfig, Path: p.PBDb3Wc, ServiceName: p.QVlaSCZ4RCK, ProxyProtocol: proxy, ForceProxyProtocol: force, Tunnel: tunnel}
		Log.TRTbrquD(0, "Trojan listen", opts.String())
		var l net.Listener
		l, err = opts.NewListener()
		if err != nil {
			return err
		}
		listeners = append(listeners, l)
	}
	v.cDyEMW2 = listeners
	for _, l := range listeners {
		v.xAm4LgsEB8j(v.zb8f3cHWW8v, l)
	}
	return nil
}
