// 原 soga-go/lib/reality 分支：按二进制证据修复上游源码差异；详见 provenance.go。
package reality

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"
	"net"
	"strconv"
	"sync"
	"time"

	utls "github.com/refraction-networking/utls"
	proxyproto "recovered.local/soga/transport/proxyproto"
)

var GlobalPostHandshakeRecordsLens sync.Map

func DetectPostHandshakeRecordsLens(config *Config) {
	for sni := range config.ServerNames {
		for alpn := range 3 { // 0, 1, 2
			key := config.Dest + " " + sni + " " + strconv.Itoa(alpn)
			if _, loaded := GlobalPostHandshakeRecordsLens.LoadOrStore(key, false); !loaded {
				go func() {
					defer func() {
						val, _ := GlobalPostHandshakeRecordsLens.Load(key)
						if _, ok := val.(bool); ok {
							GlobalPostHandshakeRecordsLens.Store(key, []int{})
						}
					}()
					var target net.Conn
					var err error
					if config.DialContext != nil {
						target, err = config.DialContext(context.Background(), config.Type, config.Dest)
					} else {
						target, err = net.Dial("tcp", config.Dest)
					}
					if err != nil {
						return
					}
					defer target.Close()
					if config.Xver == 1 || config.Xver == 2 {
						if _, err = proxyproto.X_WkJPwas4L(config.Xver, target.LocalAddr(), target.RemoteAddr()).WriteTo(target); err != nil {
							return
						}
					}
					detectConn := &DetectConn{
						DrpXao: target,
						Key:    key,
					}
					fingerprint := utls.HelloChrome_Auto
					nextProtos := []string{"h2", "http/1.1"}
					if alpn != 2 {
						fingerprint = utls.HelloGolang
					}
					if alpn == 1 {
						nextProtos = []string{"http/1.1"}
					}
					if alpn == 0 {
						nextProtos = nil
					}
					uConn := utls.UClient(detectConn, &utls.Config{
						ServerName: sni, // needs new loopvar behaviour
						NextProtos: nextProtos,
					}, fingerprint)
					if err = uConn.Handshake(); err != nil {
						return
					}
					io.Copy(io.Discard, uConn)
				}()
			}
		}
	}
}

type DetectConn struct {
	DrpXao
	Key     string
	CcsSent bool
}

func (c *DetectConn) Write(b []byte) (n int, err error) {
	if len(b) >= 3 && bytes.Equal(b[:3], []byte{20, 3, 3}) {
		c.CcsSent = true
	}
	return c.DrpXao.Write(b)
}

func (c *DetectConn) Read(b []byte) (n int, err error) {
	if !c.CcsSent {
		return c.DrpXao.Read(b)
	}
	c.DrpXao.SetReadDeadline(time.Now().Add(5 * time.Second))
	data, _ := io.ReadAll(c.DrpXao)
	var postHandshakeRecordsLens []int
	for {
		if len(data) >= 5 && bytes.Equal(data[:3], []byte{23, 3, 3}) {
			length := int(binary.BigEndian.Uint16(data[3:5])) + 5
			if len(data) < length {
				break
			}
			postHandshakeRecordsLens = append(postHandshakeRecordsLens, length)
			data = data[length:]
		} else {
			break
		}
	}
	GlobalPostHandshakeRecordsLens.Store(c.Key, postHandshakeRecordsLens)
	return 0, io.EOF
}
