package inbound

import (
	"net"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	M "recovered.local/soga/protocol/mieru"
	Log "recovered.local/soga/util/logx"
)

func (v *ONGT2EHM40r) kn9Lups0qhav(listener net.Listener) { //0x1d5fb80, closure1eb87a0
	defer func() { H.SlG_Y_mg("mieru accept loop") }()
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-v.zb8f3cHWW8v.Done():
			default:
				Log.TRTbrquD(-4, "mieru tcp accept err:", err)
			}
			return
		}
		go v.nlJ8YjzjZ(conn)
	}
}
func (v *ONGT2EHM40r) nlJ8YjzjZ(conn net.Conn) { //0x1d5fd60, closure1e6da60
	defer func() { H.SlG_Y_mg("handle mieru conn"); conn.Close() }()
	session, err := M.M_03miDCRcqa(v.zb8f3cHWW8v, conn, v.xVMggjOZQW)
	if err != nil {
		Log.TRTbrquD(-4, "mieru new session err:", err)
		return
	}
	defer session.Close()
	for {
		acceptor, err := session.Next()
		if err != nil {
			select {
			case <-v.zb8f3cHWW8v.Done():
			default:
				Log.TRTbrquD(-4, "mieru session next err:", err)
			}
			return
		}
		go v.handleMieruAcceptor(acceptor)
	}
}

// handleMieruAcceptor names original closure1e6d720. Its deferred release is
// Free, not Close; the routing chain owns its own stream-close lifecycle.
func (v *ONGT2EHM40r) handleMieruAcceptor(acceptor *M.TZ3YCqW) {
	defer func() { H.SlG_Y_mg("handle mieru proxy conn"); v.bj10P7X() }()
	conn, err := acceptor.Accept()
	if err != nil {
		Log.TRTbrquD(-4, "mieru acceptor accept err:", err)
		return
	}
	defer H.Li3c4CJ1_D(conn)
	v.sbbdX3aT6E()
	v.xR518G_hBbKv(J.MHLL7QOie(conn), conn)
}
