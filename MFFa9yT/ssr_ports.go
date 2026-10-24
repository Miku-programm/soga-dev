package MFFa9yT

import (
	"context"
	"net"
	P "recovered.local/soga/CpHCdepV"
	S "recovered.local/soga/UQ7OErS0s"
	A "recovered.local/soga/ad5KEbw_EzWk"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
)

// v_Q8IiV@1d8b500；TCP 全建好后才建 UDP，失败清理已创建的监听。
// UDP 在 tunnel 配置非空时没有本地 listener，但仍输出原 enabled 日志。
func (m *xdWl88) v_Q8IiV(u *U.LJJkaD) error {
	v := m.spzpWtrEXC1G
	port := u.SSRPort()
	listeners := []net.Listener{}
	failed := false
	defer func() {
		if failed {
			for _, l := range listeners {
				_ = l.Close()
			}
		}
	}()
	tunnel, err := C.R3DjPWSnt()
	if err != nil {
		failed = true
		return err
	}
	for _, host := range C.EdlWFAt() {
		proxy, force := C.ListenerProxyProtocolSettings()
		o := P.Ap9uD2{Host: host, Port: port, Network: "tcp", ProxyProtocol: proxy, ForceProxyProtocol: force, Tunnel: tunnel}
		listener, err := o.NewListener()
		if err != nil {
			failed = true
			return err
		}
		listeners = append(listeners, listener)
	}
	var packets []net.PacketConn
	if u.GetSinglePortType() != 1 || C.SSRObfsUDPEnabled() {
		defer func() {
			if failed {
				for _, p := range packets {
					_ = p.Close()
				}
			}
		}()
		if tunnel == nil {
			for _, host := range C.EdlWFAt() {
				raw, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(host), Port: port})
				if err != nil {
					failed = true
					return err
				}
				var pc net.PacketConn = raw
				if C.UDPProxyProtocolEnabled() {
					_, force := C.ListenerProxyProtocolSettings()
					pc = P.KlRWVpX39J(v.zb8f3cHWW8v, pc, force)
				}
				info := &A.L7E7H2aW8KVs{NxYA16xMfF: u.GetSinglePortType(), MY3lagqY: u.GetPassword(), L6BrlCXd: u.GetMethod(), C8Qsa8G8XMw0: u.GetProtocol(), UNRudV1377x9: u.GetObfs(), ZlsyKgz4EI: v.tpQEL_qKogl, N3pMXuwX85c: v.jTyGfNkAw, RLXBf_luGj: v.ktET5DFha}
				if !v.g4fBrm0Bxl {
					info.Uklg6401 = u.GetId()
				}
				wrapped, err := S.INOdx8(pc, info)
				if err != nil {
					failed = true
					_ = pc.Close()
					return err
				}
				packets = append(packets, wrapped)
			}
		}
		Log.TRTbrquD(0, "ssr udp enabled")
	} else {
		Log.TRTbrquD(0, "ssr udp disabled")
	}
	ctx, cancel := context.WithCancel(v.zb8f3cHWW8v)
	for _, l := range listeners {
		v.kOoHOawg4u(ctx, l, u)
	}
	for _, pc := range packets {
		v.e5tsZB2Ru_2v(ctx, pc.(*S.AXWGpkYq), u)
	}
	state := &pxPRmFCpW6SU{ot1kuknWf35: ctx, iHn9sGij: cancel, yF2TzminVQ1M: port, gHBoVOaQY: u, woqAbWjQH_A: listeners, aaBpAQ44Zfs: packets}
	m.Lock()
	m.sSKJ79jFt[port] = state
	m.Unlock()
	Log.TRTbrquD(0, "password:", u.GetPassword())
	Log.TRTbrquD(0, "method:", u.GetMethod())
	Log.TRTbrquD(0, "protocol:", u.GetProtocol())
	Log.TRTbrquD(0, "obfs:", u.GetObfs())
	if tunnel != nil {
		C.QitK3yi()
	}
	Log.TRTbrquD(0, "start ssr port:", u.SSRPort())
	return nil
}
