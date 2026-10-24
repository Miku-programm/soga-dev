package hysteria

import (
	"context"
	quic "github.com/apernet/quic-go"
	"github.com/apernet/quic-go/http3"
	"net"
	"net/http"
	B "recovered.local/soga/net/dns"
	G "recovered.local/soga/net/relay"
	J "recovered.local/soga/net/sess"
	S "recovered.local/soga/transport/quicstream"
	W "recovered.local/soga/crypto/misc"
	Log "recovered.local/soga/util/logx"
	Congestion "recovered.local/soga/transport/quicpacing"
	U "recovered.local/soga/user"
	"sync"
)

type cYzxTBfcV struct {
	szJp1HM      context.Context
	ff2_VHDMO_e  *InfdqOmZK
	lsVoia1      quic.Connection
	chDy7ZwdT    bool
	gjUa1J3Pmhm  sync.Mutex
	tZIs5a3Jb    U.PgaNd3014U
	ehiG0t4uYJad *B.D23y2888a_
	dXJZalu06    chan J.Ki8urSHTx44
}

func (h *cYzxTBfcV) ServeHTTP(w http.ResponseWriter, r *http.Request) { //0x1c30240
	if r.Method != "POST" || r.Host != "hysteria" || r.URL.Path != "/auth" {
		h._Oplj7Ks(w, r)
		return
	}
	h.gjUa1J3Pmhm.Lock()
	defer h.gjUa1J3Pmhm.Unlock()
	c := h.ff2_VHDMO_e
	if h.chDy7ZwdT {
		W.CIvpFjH270ca(w.Header(), !c.QlDN7UZH, c.FYaeQJDXC.OfLAiHnwZo, c.TDNLdWhC)
		w.WriteHeader(233)
		return
	}
	password, rx := W.GZDAjvxfC(r.Header)
	u, ok := c.IDZiaXEBjVIW.Authenticate(h.lsVoia1.RemoteAddr(), password, rx)
	if !ok {
		h._Oplj7Ks(w, r)
		return
	}
	h.chDy7ZwdT = true
	h.tZIs5a3Jb = u
	h.szJp1HM = u.GetCtx()
	stop := context.AfterFunc(h.szJp1HM, func() { h.lsVoia1.CloseWithError(0x100, "") })
	context.AfterFunc(h.lsVoia1.Context(), func() { stop() })
	tx := rx
	if c.FYaeQJDXC.GSKeMid != 0 && tx > c.FYaeQJDXC.GSKeMid {
		tx = c.FYaeQJDXC.GSKeMid
	}
	if !c.TDNLdWhC && tx != 0 {
		Congestion.XpQZo1W7(h.lsVoia1, tx)
	} else {
		Congestion.AW6N_Jc(h.lsVoia1)
	}
	W.CIvpFjH270ca(w.Header(), !c.QlDN7UZH, c.FYaeQJDXC.OfLAiHnwZo, c.TDNLdWhC)
	w.WriteHeader(233)
	if !c.QlDN7UZH {
		go func() {
			v := &chFbgN{yoncNMg: h.szJp1HM, eAEJVc: h.ehiG0t4uYJad, ykxy0e: h.lsVoia1, ndlUQyGetN: h.dXJZalu06, iuT4giLUja: u}
			go v.aMFjq2VlDD()
		}()
	}
}

func (h *cYzxTBfcV) _Oplj7Ks(w http.ResponseWriter, r *http.Request) { //0x1c30e00
	if h.ff2_VHDMO_e.DGm6Idekj != nil {
		h.ff2_VHDMO_e.DGm6Idekj.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (h *cYzxTBfcV) ProxyStreamHijacker(kind http3.FrameType, _ quic.ConnectionTracingID, stream quic.Stream, err error) (bool, error) { //0x1c30840
	if err != nil || !h.chDy7ZwdT {
		return false, nil
	}
	select {
	case <-h.szJp1HM.Done():
		return false, nil
	default:
	}
	s := &S.U6qURXBS2qu{Fjf4om: stream}
	if kind != 0x401 {
		return false, nil
	}
	go h.xnuZavu4X(s)
	return true, nil
}

func (h *cYzxTBfcV) xnuZavu4X(stream quic.Stream) { //0x1c309e0
	dest, err := W.T7aXhP(stream)
	if err != nil {
		Log.TRTbrquD(-4, "read hysteria request err:", err)
		stream.Close()
		return
	}
	addr, err := G.U80FDt(dest)
	if err != nil {
		Log.TRTbrquD(-4, "hysteria invalid address:", dest)
		stream.Close()
		return
	}
	if err = W.C3FDNGT2zI(stream, true, ""); err != nil {
		Log.TRTbrquD(-4, err)
		stream.Close()
		return
	}
	h.dXJZalu06 <- vEaGVT9SbgBM(stream, h.lsVoia1.LocalAddr(), h.lsVoia1.RemoteAddr(), h.tZIs5a3Jb, addr)
}

type BqTbYn4 struct {
	quic.Stream
	wp78c97cOv net.Addr
	kWYaml     net.Addr
}

func (c *BqTbYn4) LocalAddr() net.Addr  { return c.wp78c97cOv } //0x1c34600
func (c *BqTbYn4) RemoteAddr() net.Addr { return c.kWYaml }     //0x1c34620
type DrpXao = net.Conn
type StUVueP6a struct {
	DrpXao
	bQt7p94rWfG U.PgaNd3014U
	z8ptDc4     *G.AHU2MRmjzQu
}

func (c *StUVueP6a) GetUser() U.PgaNd3014U       { return c.bQt7p94rWfG } //0x1c301e0
func (c *StUVueP6a) GetDestAddr() *G.AHU2MRmjzQu { return c.z8ptDc4 }     //0x1c30200
func (c *StUVueP6a) NetworkType() J.ZES6F5       { return 0 }             //0x1c30220
func vEaGVT9SbgBM(stream quic.Stream, local, remote net.Addr, user U.PgaNd3014U, dest *G.AHU2MRmjzQu) J.Ki8urSHTx44 { //0x1c2ffc0
	counter := &J.Mdn09F6aMrpa{DrpXao: &BqTbYn4{stream, local, remote}}
	if user != nil {
		counter.SetUser(user)
	}
	return &StUVueP6a{counter, user, dest}
}
