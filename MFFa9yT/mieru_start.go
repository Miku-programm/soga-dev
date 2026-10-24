package MFFa9yT

import (
	"fmt"
	"net"
	P "recovered.local/soga/CpHCdepV"
	L "recovered.local/soga/LFWTTXF7JR"
	N "recovered.local/soga/NtxdAE"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	"strings"
)

func tzgQP2vl(arg *DV_XPJ4z0) (Kz_8ajU7, error) { //0x1d5f2a0
	base, err := n2s6Gf9ey(arg)
	if err != nil {
		return nil, err
	}
	v := &ONGT2EHM40r{JWN154qJ4: base}
	v.xVMggjOZQW = &N.O0UldNL4Eq{RCaEGDlm: base.zb8f3cHWW8v, Cebit_sS: int64(arg.HG5k0365.H47AzZsYHB1N.EMlXEc), RacA3he0TX: base.jc_QPN6}
	v.UpdateAPIData(arg.HG5k0365)
	Log.TRTbrquD(0, "create mieru server")
	return v, nil
}
func (v *ONGT2EHM40r) Start() error { //0x1d5f460
	p := &v.wiy6P3.Load().ABA4hq
	network := strings.ToUpper(p.GGSLzD)
	if network != "" && network != "TCP" {
		return fmt.Errorf("unsupported mieru transport: %v (only TCP is supported)", p.GGSLzD)
	}
	listeners, err := v.p7SNWA(p)
	if err != nil {
		return err
	}
	v.d7Fav9U0ykgg = listeners
	for _, l := range listeners {
		go v.kn9Lups0qhav(l)
	}
	return nil
}
func (v *ONGT2EHM40r) p7SNWA(p *L.Xh4IiOa0) (listeners []net.Listener, err error) { //0x1d5f6e0
	tunnel, err := C.R3DjPWSnt()
	if err != nil {
		return nil, err
	}
	if tunnel != nil {
		C.QitK3yi()
	}
	var opened []net.Listener
	success := false
	defer func() {
		if !success {
			for _, l := range opened {
				l.Close()
			}
		}
	}()
	for _, host := range C.EdlWFAt() {
		proxy, force := C.ListenerProxyProtocolSettings()
		options := &P.Ap9uD2{Host: host, Port: p.MQE3K4mq, Network: "tcp", ProxyProtocol: proxy, ForceProxyProtocol: force, Tunnel: tunnel}
		l, e := options.NewListener()
		if e != nil {
			return nil, e
		}
		opened = append(opened, l)
		Log.V8GN5Stk(0, "mieru listen %v", options.String())
	}
	success = true
	return opened, nil
}
