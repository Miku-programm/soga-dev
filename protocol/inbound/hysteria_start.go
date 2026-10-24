package inbound

import (
	"context"
	"fmt"
	"net"
	B "recovered.local/soga/net/dns"
	Native "recovered.local/soga/protocol/hysteria"
	P "recovered.local/soga/transport/listener"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
)

func hpdxYbrDSPs(arg *DV_XPJ4z0) (Kz_8ajU7, error) { //0x1d56a00
	b, err := n2s6Gf9ey(arg)
	if err != nil {
		return nil, err
	}
	v := &Aq5i0Bwez8{JWN154qJ4: b}
	v.UpdateAPIData(arg.HG5k0365)
	Log.TRTbrquD(0, "create hysteria server")
	return v, nil
}

func (v *Aq5i0Bwez8) Start() (err error) { //0x1d56ba0
	node := v.wiy6P3.Load()
	cfg := &node.AYf5kQSQW
	version := cfg.Li6nvK1
	if version == 0 {
		version = 2
	}
	if version != 2 {
		return fmt.Errorf("not support hysteria version %v", version)
	}
	tlsConfig, err := v.kaVZkL()
	if err != nil {
		return err
	}
	if err = v.dINraX_uOx5T(tlsConfig); err != nil {
		return err
	}
	var listeners []*Native.P3_qJL5v
	defer func() {
		if err != nil {
			for _, l := range listeners {
				l.Close()
			}
		}
	}()
	config, err := Native.LnSg9qSRSAT(tlsConfig, v, cfg.W4ggEz, cfg.HnHarW)
	if err != nil {
		return err
	}
	for _, host := range C.EdlWFAt() {
		raw, e := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(host), Port: cfg.PIvpxN6YK})
		if e != nil {
			return e
		}
		var conn net.PacketConn = raw
		if C.UDPProxyProtocolEnabled() {
			_, force := C.ListenerProxyProtocolSettings()
			conn = P.KlRWVpX39J(v.zb8f3cHWW8v, raw, force)
		}
		l, e := Native.VOhu4P9d2b(v.zb8f3cHWW8v, conn, string(cfg.B30Jt2kO), cfg.UjgtNuouuKVl, config, func() *B.D23y2888a_ { return v.ffju0Enjo.Load() })
		if e != nil {
			conn.Close()
			return e
		}
		listeners = append(listeners, l)
		if cfg.B30Jt2kO == "" || cfg.B30Jt2kO == "plain" {
			Log.V8GN5Stk(0, "hysteria%v listen udp: %v, obfs: plain", version, l.Addr())
		} else {
			Log.V8GN5Stk(0, "hysteria%v listen udp: %v, obfs: %v, obfs-password: %v", version, l.Addr(), cfg.B30Jt2kO, cfg.UjgtNuouuKVl)
		}
	}
	v.uqbfp_xRjia = listeners
	for _, l := range listeners {
		v.xAm4LgsEB8j(l)
	}
	return nil
}

func (v *Aq5i0Bwez8) xAm4LgsEB8j(l *Native.P3_qJL5v) { //0x1d57540及0x1e6ce80
	tunnel, _ := C.TunnelSettings()
	unsupported := tunnel && !C.Q0SDIN()
	go func() {
		defer func() { Log.TRTbrquD(-4, "hysteria listener close:", l.Addr()); l.Close() }()
		for {
			select {
			case <-v.zb8f3cHWW8v.Done():
				return
			default:
			}
			conn, err := l.AcceptProxy()
			if err != nil {
				return
			}
			if unsupported {
				go func() { Log.TRTbrquD(8, "not support hysteria tunnel") }()
				continue
			}
			go v.nlJ8YjzjZ(v.zb8f3cHWW8v, conn)
		}
	}()
}

func (v *Aq5i0Bwez8) nlJ8YjzjZ(_ context.Context, conn J.Ki8urSHTx44) { //0x1d58260
	defer func() {
		H.SlG_Y_mg("handle hysteria")
		conn.Close()
		Log.TRTbrquD(-4, "handle hysteria end")
		v.bj10P7X()
	}()
	Log.TRTbrquD(-4, "start handle client", conn.RemoteAddr(), "connection")
	v.sbbdX3aT6E()
	v.xR518G_hBbKv(J.MHLL7QOie(conn), conn)
}
