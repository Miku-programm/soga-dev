package inbound

import (
	"crypto/tls"
	"net"
	P "recovered.local/soga/transport/listener"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
)

// Start @0x1da1360及错误关闭闭包0x1e737a0。
func (v *K4avFiYOPi) Start() (err error) {
	if v.l_ZoZQDxG7 != nil {
		loaded, e := v.l_ZoZQDxG7.Load(v._7le3Hqii(), v.jc_QPN6)
		if e != nil {
			Log.V8GN5Stk(4, "load ip user cache err: %v", e)
		} else if loaded {
			Log.TRTbrquD(0, "load ip user cache success")
		}
	}
	var tlsConfig *tls.Config
	if v.wiy6P3.Load().SwxJNgPpg.UVeahqDBT {
		tlsConfig, err = v.kaVZkL()
		if err != nil {
			return err
		}
		if err = v.dINraX_uOx5T(tlsConfig); err != nil {
			return err
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
			for _, listener := range listeners {
				listener.Close()
			}
		}
	}()
	for _, host := range C.EdlWFAt() {
		node := v.wiy6P3.Load()
		proxy, force := C.ListenerProxyProtocolSettings()
		options := &P.Ap9uD2{Host: host, Port: node.SwxJNgPpg.CKsOpayd, Network: node.SwxJNgPpg.F2GXyZZ, TLS: tlsConfig, Path: node.SwxJNgPpg.ChJnngzb, ServiceName: node.SwxJNgPpg.KWffzRDF, ProxyProtocol: proxy, ForceProxyProtocol: force, Tunnel: tunnel}
		Log.TRTbrquD(0, "VMess listen", options.String())
		var listener net.Listener
		listener, err = options.NewListener()
		if err != nil {
			return err
		}
		listeners = append(listeners, listener)
	}
	v.l8Uq4M = listeners
	v.yER8qM3e(v.zb8f3cHWW8v)
	return nil
}
