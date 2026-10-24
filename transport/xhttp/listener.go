package xhttp

import (
	"errors"
	"net"
	"net/http"
	Log "recovered.local/soga/util/logx"
	"recovered.local/soga/xray/done"
	"recovered.local/soga/xray/internet"
	"sync"
	"time"
)

// 原 RTTI 352 字节及 136 字节；所有字段均来自样本元数据。
type KcahyAO struct {
	sync.Mutex
	uSJi0d      http.Server
	c_GCatXUp_G net.Listener
	fBC_MC1GOIw *R0J19auWof
	eSKdfEmK    chan net.Conn
	sJyUTeDl    sync.Once
}
type e8Q7HURSMg struct {
	jgg_3z8gKTCK *R0J19auWof
	taRd4Q3PP5   string
	jD05lxT      string
	mZcQrY       *KcahyAO
	cl3pW8JGwae  *sync.Mutex
	wLTHxb       sync.Map
	_7JXu_SO6K   net.Addr
	epe3GXboMP1  *internet.SocketConfig
	vat5bP_nCH   chan net.Conn
}
type f30vZAC4M2Oh struct {
	jg638CNSsPu *y0bchQKJ
	zYk4hazu    *done.Instance
}

// 0x125e5c0。HTTP/1 调用需使用 buildhttp 中经核对的构建覆盖。
func DWB3oG(raw net.Listener, cfg *R0J19auWof) net.Listener {
	if cfg == nil {
		cfg = &R0J19auWof{}
	}
	queue := make(chan net.Conn, 128)
	l := &KcahyAO{c_GCatXUp_G: raw, fBC_MC1GOIw: cfg, eSKdfEmK: queue}
	h := &e8Q7HURSMg{jgg_3z8gKTCK: cfg, taRd4Q3PP5: cfg.TzkkcVTusJD, jD05lxT: cfg.GetNormalizedPath(), mZcQrY: l, cl3pW8JGwae: new(sync.Mutex), _7JXu_SO6K: raw.Addr(), vat5bP_nCH: queue}
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	l.uSJi0d = http.Server{Handler: h, ReadHeaderTimeout: 4 * time.Second, MaxHeaderBytes: cfg.GetNormalizedServerMaxHeaderBytes(), Protocols: protocols}
	go func() {
		if err := l.uSJi0d.Serve(raw); err != nil {
			Log.V8GN5Stk(0, "xhttp server stopped: %v", err)
		}
	}()
	return l
}
func (l *KcahyAO) Accept() (net.Conn, error) { // 0x125e9a0
	c, ok := <-l.eSKdfEmK
	if !ok {
		return nil, errors.New("listener closed")
	}
	return c, nil
}

// 0x125ea40 / 0x12731e0：原 Addr 第一次调用确实关闭队列；保留此异常行为。
func (l *KcahyAO) Addr() net.Addr {
	l.sJyUTeDl.Do(func() { close(l.eSKdfEmK) })
	return l.c_GCatXUp_G.Addr()
}
func (l *KcahyAO) Close() error { return l.c_GCatXUp_G.Close() } //0x125eac0

func (h *e8Q7HURSMg) d39ITgdaw(id string) *f30vZAC4M2Oh { //0x125a360
	if v, ok := h.wLTHxb.Load(id); ok {
		return v.(*f30vZAC4M2Oh)
	}
	h.cl3pW8JGwae.Lock()
	defer h.cl3pW8JGwae.Unlock()
	if v, ok := h.wLTHxb.Load(id); ok {
		return v.(*f30vZAC4M2Oh)
	}
	limit := int(h.mZcQrY.fBC_MC1GOIw.GetNormalizedScMaxBufferedPosts())
	s := &f30vZAC4M2Oh{jg638CNSsPu: &y0bchQKJ{dL_ULZL: make(chan UmYbKhCpX, limit), vPdgI_lx: make(eMSUo20X, 0), fAEWeMURGLm0: limit}, zYk4hazu: done.New()}
	h.wLTHxb.Store(id, s)
	timeout := done.New()
	go func() { time.Sleep(30 * time.Second); timeout.Close() }() //0x1272e40
	go func() {                                                   //0x1272d40
		select {
		case <-s.zYk4hazu.Wait():
		case <-timeout.Wait():
			h.wLTHxb.Delete(id)
			s.jg638CNSsPu.Close()
		}
	}()
	return s
}
