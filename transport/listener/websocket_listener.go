package listener

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"io"
	"net"
	"net/http"
	ws "recovered.local/soga/transport/websocket"
	"recovered.local/soga/util/errutil"
	util "recovered.local/soga/transport/wsutil"
	"recovered.local/soga/util/logx"
	"strings"
	"sync"
	"time"
)

type O6B2qZ0p0 = net.Listener

// 原 RTTI 显示前缀 eucMacxiQPzD，但 uncommon_package_path 为 CpHCdepV。
// 两个结构分别为48/88字节；以下为机器码源码表达，不是作者文件原文。
type RETaJc91 struct {
	O6B2qZ0p0
	kmDrwd       chan net.Conn
	jPywQp       string
	z35DsaN5f7ru bool
}
type E1sAyQm struct {
	DrpXao
	rm3mq9jnM_   io.Reader
	v5a_08       util.KTXwG8eroaXF
	qx_a3m7JIk6v net.IP
	ruhm89I      sync.Once
	hkUJ_SLSaJ_5 func()
}

// 原5095FD0由base64.init@72a7c0构造：20ed3ba字母表末尾为-_，
// 72aa43设置padding=-1，72aa6c发布。不是根据协议文档推断。
var websocketBase64At5095FD0 = base64.RawURLEncoding

// H7N2ayu 0x18bd3c0、accept循环0x18c2aa0、处理器0x18c2c20。
func H7N2ayu(listener net.Listener, path string) net.Listener {
	path = strings.Split(path, "?")[0]
	if len(path) == 0 {
		path = "/"
	} else if path[0] != '/' {
		path = "/" + path
	}
	v := &RETaJc91{O6B2qZ0p0: listener, kmDrwd: make(chan net.Conn, 1024), jPywQp: path}
	go func() {
		for {
			c, err := listener.Accept()
			if err != nil {
				return
			}
			go func() {
				var clientIP net.IP
				var early io.Reader
				upgrader := ws.L4gwp9B{
					OnRequest: func(uri []byte) error { //0x18c3960
						if string(uri) == path {
							return nil
						}
						return fmt.Errorf("path mismatch: expect: %v, actual: %v", path, string(uri))
					},
					ProtocolCustom: func(value []byte) (string, bool) { //0x18c32c0
						b := make([]byte, websocketBase64At5095FD0.DecodedLen(len(value)))
						if n, e := websocketBase64At5095FD0.Decode(b, value); e == nil {
							early = bytes.NewReader(b[:n])
						}
						return "", true
					},
					OnHeader: func(key, value []byte) error { //0x18c3460
						switch strings.ToLower(string(key)) {
						case "x-real-ip":
							if clientIP == nil {
								clientIP = net.ParseIP(string(value))
							}
						case "x-forwarded-for":
							if clientIP == nil {
								clientIP = net.ParseIP(strings.TrimSpace(strings.Split(string(value), ",")[0]))
							}
						case "sec-websocket-protocol":
							b := make([]byte, websocketBase64At5095FD0.DecodedLen(len(value)))
							if n, e := websocketBase64At5095FD0.Decode(b, value); e == nil {
								early = bytes.NewReader(b[:n])
							}
						}
						return nil
					},
				}
				_, err = upgrader.Upgrade(c)
				if err != nil {
					logx.V8GN5Stk(-4, "ws handshake failed: %v", err)
					c.Close()
					return
				}
				c = &E1sAyQm{DrpXao: c, rm3mq9jnM_: early, v5a_08: util.L9DiQX(c, ws.StateServerSide), qx_a3m7JIk6v: clientIP, hkUJ_SLSaJ_5: func() {}} //原18c4060直接RET
				select {
				case v.kmDrwd <- c:
				default:
					c.Close()
				}
			}()
		}
	}()
	return v
}

// Accept 0x18bd5c0：关闭后异步清理队列。原闭包捕获复用变量c，照原保留。
func (v *RETaJc91) Accept() (net.Conn, error) {
	if v.z35DsaN5f7ru {
		return nil, errors.New("ws listener closed")
	}
	c := <-v.kmDrwd
	if v.z35DsaN5f7ru {
		if c != nil {
			c.Close()
		}
	} else if c != nil {
		return c, nil
	}
	for c = range v.kmDrwd {
		go func() { c.Close() }()
	}
	return nil, errors.New("ws listener closed")
}
func (v *RETaJc91) Close() error { //0x18bd880
	if v.z35DsaN5f7ru {
		return nil
	}
	v.z35DsaN5f7ru = true
	close(v.kmDrwd)
	return v.O6B2qZ0p0.Close()
}

// fa_PWrTv 0x18bd8e0；HTTP接口路径采用固定优先级，和raw回调的首个有效IP不同。
func (v *RETaJc91) fa_PWrTv(h http.Header) net.IP {
	s := h.Get("X-Real-Ip")
	if s == "" {
		if forwarded := h.Get("X-Forwarded-For"); forwarded != "" {
			s = strings.TrimSpace(strings.Split(forwarded, ",")[0])
		}
	}
	if s == "" {
		return nil
	}
	return net.ParseIP(s)
}

// ServeHTTP 0x18bdaa0。原实现先Upgrade，再读取early-data请求头；
// 此处创建的回显header map没有参与已完成的响应，不能擅自修成有效回显。
func (v *RETaJc91) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() { errutil.SlG_Y_mg("serve ws"); r.Body.Close() }()
	if v.z35DsaN5f7ru {
		return
	}
	if r.URL.Path != v.jPywQp {
		w.WriteHeader(404)
		return
	}
	c, _, _, e := ws.DefaultHTTPUpgrader.Upgrade(r, w)
	if e != nil {
		return
	}
	headers := make(http.Header)
	var early io.Reader
	if s := r.Header.Get("Sec-WebSocket-Protocol"); s != "" {
		if b, e := websocketBase64At5095FD0.DecodeString(s); e == nil && len(b) > 0 {
			early = bytes.NewReader(b)
			headers.Set("Sec-WebSocket-Protocol", s)
		}
	}
	ip := v.fa_PWrTv(r.Header)
	ctx, cancel := context.WithCancel(context.Background())
	c = &E1sAyQm{DrpXao: c, rm3mq9jnM_: early, v5a_08: util.L9DiQX(c, ws.StateServerSide), qx_a3m7JIk6v: ip, hkUJ_SLSaJ_5: func() { cancel() }}
	v.kmDrwd <- c
	<-ctx.Done()
}

// gw309ChmEWia 0x18c2260。
func (v *E1sAyQm) gw309ChmEWia() (io.Reader, error) {
	if v.rm3mq9jnM_ != nil {
		return v.rm3mq9jnM_, nil
	}
	r := &util.HLKO3Cdqw{GX04GFKh1Gz: v.DrpXao, WXUh5psqY7: ws.StateServerSide, EbVEHBvY4lh: v.v5a_08}
	for {
		h, e := r.NextFrame()
		if e != nil {
			return nil, e
		}
		if h.IsHe8C&8 == 0 {
			break
		}
		if e = v.v5a_08(h, r); e != nil {
			return nil, e
		}
	}
	v.rm3mq9jnM_ = r
	return r, nil
}
func (v *E1sAyQm) Read(p []byte) (int, error) { //0x18c2440
	for {
		r, e := v.gw309ChmEWia()
		if e != nil {
			return 0, e
		}
		n, e := r.Read(p)
		if pkgerrors.Cause(e) != io.EOF {
			return n, e
		}
		v.rm3mq9jnM_ = nil
		if n > 0 {
			return n, nil
		}
	}
}
func (v *E1sAyQm) Write(p []byte) (int, error) { //0x18c2560
	e := util.WriteFrameRecovered(v.DrpXao, ws.StateServerSide, ws.OpBinary, true, p)
	if e != nil {
		return 0, e
	}
	return len(p), nil
}
func (v *E1sAyQm) RemoteAddr() net.Addr { //0x18c26a0
	if v.qx_a3m7JIk6v != nil {
		return &net.TCPAddr{IP: v.qx_a3m7JIk6v}
	}
	return v.DrpXao.RemoteAddr()
}
func (v *E1sAyQm) Close() error { //0x18c2740/18c3f60
	v.ruhm89I.Do(func() {
		v.DrpXao.SetWriteDeadline(time.Now().Add(time.Second))
		util.WriteFrameRecovered(v.DrpXao, ws.StateServerSide, ws.OpClose, true, nil)
	})
	e := v.DrpXao.Close()
	v.hkUJ_SLSaJ_5()
	return e
}
func (v *E1sAyQm) GetRawConn() net.Conn { return v.DrpXao } //0x18c27e0
