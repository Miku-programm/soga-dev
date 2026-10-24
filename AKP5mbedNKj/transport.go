package AKP5mbedNKj

import (
	"context"
	"crypto/tls"
	"errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"io"
	"log"
	"net"
	"net/http"
	G "recovered.local/soga/G6bHD5JA"
	Log "recovered.local/soga/muc279Mpa6S4"
	Q "recovered.local/soga/qcqMSUWwIx40"
	"strconv"
	"sync"
	"time"
)

// 业务逻辑依据原包 AKP5mbedNKj 的 C/ASM 恢复。HTTP/1 写出需通过 buildhttp
// 构建，以保留原 CONNECT/Method 分支；h2c 与 TLS 的 HTTP/2 调用链分别核对。
type O6B2qZ0p0 = net.Listener
type SNvD5V struct {
	O6B2qZ0p0
	u7tJdDFk5  context.Context
	byhCVWt    context.CancelFunc
	f4RBDAi    *http.Server
	a4Hcy37uDv string
	vTk17Wmts  chan net.Conn
}
type IwDQb32fnlH9 struct {
	sync.RWMutex
	oC5aEPzrec9w http.ResponseWriter
	px3ua9JUg    *http.Request
	fRm48w2      *Q.Q5KYzU
	mot9aQsb62ts net.Addr
	zjpabCK      net.Addr
	fMRTh247ym2l bool
	g2MwZxLFm_YD func()
}

func (c *IwDQb32fnlH9) Read(p []byte) (int, error) { //0x18b4f60
	c.RLock()
	closed := c.fMRTh247ym2l
	c.RUnlock()
	if closed {
		return 0, io.ErrClosedPipe
	}
	return c.fRm48w2.Read(p)
}
func (c *IwDQb32fnlH9) Write(p []byte) (int, error) { //0x18b5060
	c.RLock()
	closed := c.fMRTh247ym2l
	c.RUnlock()
	if closed {
		return 0, io.ErrClosedPipe
	}
	n, err := c.oC5aEPzrec9w.Write(p)
	if flusher, ok := c.oC5aEPzrec9w.(http.Flusher); ok {
		flusher.Flush()
	}
	return n, err
}
func (c *IwDQb32fnlH9) Close() error { //0x18b5220
	c.Lock()
	defer c.Unlock()
	if c.fMRTh247ym2l {
		return nil
	}
	c.fMRTh247ym2l = true
	if c.g2MwZxLFm_YD != nil {
		c.g2MwZxLFm_YD()
	}
	if flusher, ok := c.oC5aEPzrec9w.(http.Flusher); ok {
		flusher.Flush()
	}
	return c.px3ua9JUg.Body.Close()
}
func (c *IwDQb32fnlH9) LocalAddr() net.Addr               { return c.mot9aQsb62ts }               //0x18b53a0
func (c *IwDQb32fnlH9) RemoteAddr() net.Addr              { return c.zjpabCK }                    //0x18b53c0
func (c *IwDQb32fnlH9) SetDeadline(t time.Time) error     { return c.fRm48w2.SetReadDeadline(t) } //0x18b53e0
func (c *IwDQb32fnlH9) SetReadDeadline(t time.Time) error { return c.fRm48w2.SetReadDeadline(t) } //0x18b5440
func (c *IwDQb32fnlH9) SetWriteDeadline(time.Time) error  { return nil }                          //0x18b54a0

// Qi8msD8m74 @0x18b54c0，TLS分支直接覆盖传入配置的NextProtos。
func Qi8msD8m74(listener net.Listener, path string, tlsConfig *tls.Config) net.Listener {
	ctx, cancel := context.WithCancel(context.Background())
	l := &SNvD5V{O6B2qZ0p0: listener, u7tJdDFk5: ctx, byhCVWt: cancel, a4Hcy37uDv: path, vTk17Wmts: make(chan net.Conn, 16384)}
	handler := http.HandlerFunc(l.dPjtzcKR)
	server := &http.Server{ErrorLog: log.New(io.Discard, "", 1), ReadHeaderTimeout: 4 * time.Second}
	if tlsConfig != nil {
		tlsConfig.NextProtos = []string{"h2"}
		server.TLSConfig = tlsConfig
		server.Handler = handler
	} else {
		server.Handler = h2c.NewHandler(handler, &http2.Server{})
	}
	l.f4RBDAi = server
	go func() {
		if tlsConfig != nil {
			server.ServeTLS(listener, "", "")
		} else {
			server.Serve(listener)
		}
	}()
	return l
}

// dPjtzcKR @0x18b5980；原handler不检查path/method，字段仅在构造中保存。
func (l *SNvD5V) dPjtzcKR(w http.ResponseWriter, r *http.Request) {
	addr := G.SOmphOG1Ty(r.Header)
	if addr == nil {
		host, port, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			if ip := net.ParseIP(host); ip != nil {
				number, _ := strconv.Atoi(port)
				addr = &net.TCPAddr{IP: ip, Port: number}
			}
		}
	}
	if addr == nil {
		Log.TRTbrquD(4, "can not decode h2 remote addr:", r.RemoteAddr)
		w.WriteHeader(404)
		return
	}
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(200)
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
	ctx, cancel := context.WithCancel(l.u7tJdDFk5)
	c := &IwDQb32fnlH9{oC5aEPzrec9w: w, px3ua9JUg: r, fRm48w2: Q.NewReader(r.Body), mot9aQsb62ts: l.O6B2qZ0p0.Addr(), zjpabCK: addr, g2MwZxLFm_YD: cancel}
	// 原发送无select，保留满队列时的阻塞行为。
	l.vTk17Wmts <- c
	<-ctx.Done()
}
func (l *SNvD5V) Accept() (net.Conn, error) { //0x18b6100
	select {
	case <-l.u7tJdDFk5.Done():
		return nil, errors.New("h2 listener closed")
	case c := <-l.vTk17Wmts:
		select {
		case <-l.u7tJdDFk5.Done():
			if c != nil {
				c.Close()
			}
			return nil, errors.New("h2 listener closed")
		default:
			return c, nil
		}
	}
}
func (l *SNvD5V) Close() error { l.byhCVWt(); return l.f4RBDAi.Close() } //0x18b6320
