package encmux

import (
	"crypto/tls"
	"errors"
	"fmt"
	"golang.org/x/crypto/sha3"
	D "recovered.local/soga/crypto/kdf"
	A "recovered.local/soga/transport/utlsconfig"
	I "recovered.local/soga/transport/realitytypes"
	"strings"
)

// InboundConfiguration 是恢复工程对 FZzfvtn 第一个输入已读取字段的源码命名。
// 原输入的类型/字段名没有留在 RTTI 中；下列注释是机器码读取的原偏移，
// 不是声称此 Go 结构与原 408 字节输入布局相同。未消费字段不伪造填充。
type InboundConfiguration struct {
	Protocol      string      // +120
	Method        D.RQvD2v3oZ // +136
	WebSocketPath string      // +200
	Password      string      // +216
	Certificate   *I.ZK5THfX  // +240
	ReceiveProxy  int         // +272；原判断为 >0
	ProxyMode     bool        // +384
	Mux           bool        // +385
	MuxNum        int         // +392
}

// OutboundConfiguration 同理按第二输入实际消费字段命名，非原类型原文。
type OutboundConfiguration struct {
	Protocol      string      // +64
	Method        D.RQvD2v3oZ // +80
	Host          string      // +104
	ServerName    string      // +160；非空时优先于 Host
	Insecure      bool        // +176
	WebSocketPath string      // +184
	Password      string      // +200
	ProxyMode     bool        // +216
	ProxyHost     string      // +224
	ProxyPort     int         // +240
	Mux           bool        // +338
	MuxNum        int         // +344
}

// FZzfvtn 对应 0x11da840。保留原先解析入站证书、再覆盖出站参数的顺序。
// 类型命名为恢复辅助；逻辑逐分支来自原机器码，包含原证书/密码失败路径。
func FZzfvtn(in *InboundConfiguration, out *OutboundConfiguration) (*VTl94pzz2IkB, error) {
	var protocol, password, path, serverName string
	var method D.RQvD2v3oZ
	var cert *tls.Certificate
	var receiveProxy, proxy, mux, insecure bool
	var muxNum int
	var proxyAddr *A.FGMp1YT7
	if in != nil {
		protocol, method, path, password = in.Protocol, in.Method, in.WebSocketPath, in.Password
		switch protocol {
		case "tls", "tls-tunnel", "quic-tunnel", "ws-tls-tunnel":
			if in.Certificate == nil {
				return nil, fmt.Errorf("tls tunnel but not config cert")
			}
			parsed, err := tls.X509KeyPair([]byte(in.Certificate.Y1mdfY), []byte(in.Certificate.SzI2MbNsk))
			if err != nil {
				return nil, err
			}
			cert = &parsed
		}
		receiveProxy = in.ReceiveProxy > 0
		proxy, mux, muxNum = in.ProxyMode, in.Mux, in.MuxNum
	}
	if out != nil {
		protocol, method, path, password = out.Protocol, out.Method, out.WebSocketPath, out.Password
		serverName = out.ServerName
		if serverName == "" {
			serverName = out.Host
		}
		insecure, proxy, mux, muxNum = out.Insecure, out.ProxyMode, out.Mux, out.MuxNum
		if proxy {
			proxyAddr = A.Fqn4mE5(out.ProxyHost, out.ProxyPort)
		}
	}
	if password == "" {
		return nil, errors.New("please set password for tunnel type")
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	suite, err := D.MakT9YbIpyNJ(method)
	if err != nil {
		return nil, err
	}
	key := D.BMQB8vpnQXXB(sha3.New512, []byte(password), suite.KeyLen())
	token := D.ENKzXYE1a(sha3.New512, key, derivationInfoAt5099930, []byte(password), 16)
	v := &VTl94pzz2IkB{
		ufiMqR6x:     protocol == "ws-tunnel" || protocol == "ws-tls-tunnel",
		hZORuZstvS07: protocol == "quic-tunnel",
		rj1DhSBHZ7O:  protocol == "tls" || protocol == "tls-tunnel" || protocol == "quic-tunnel" || protocol == "ws-tls-tunnel",
		e2TiiBKaBt:   mux && protocol != "quic-tunnel", pyauaz: muxNum,
		p1Wtyln2GJ6: serverName, psdXJ2Mi: insecure, pJLNnm: password, _rKF9rrYaf8a: cert,
		lyXfKPaz4m23: path, jrF7WDTK: receiveProxy, hrXHbLocdr: proxy, dlAoUM: proxyAddr, zkn35IN: suite, oWvGY4I9cA: key,
	}
	copy(v.x5Lor8hDyn[:], token)
	return v, nil
}
