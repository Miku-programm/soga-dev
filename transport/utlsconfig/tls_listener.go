package utlsconfig

import (
	"crypto/tls"
	"net"
)

// MYA8Jh 对应0x1171ec0，原隧道专用TLS1.2~1.3配置与九项套件顺序。
// http2开关只控制ALPN；不复制传入证书的内部切片。
func MYA8Jh(listener net.Listener, cert *tls.Certificate, http2 bool) net.Listener {
	protocols := []string{"http/1.1"}
	if http2 {
		protocols = []string{"h2", "http/1.1"}
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{*cert}, CipherSuites: []uint16{0x1302, 0x1303, 0x1301, 0xc02c, 0xcca9, 0xc02b, 0xc023, 0xc030, 0xc02f}, MinVersion: tls.VersionTLS12, MaxVersion: tls.VersionTLS13, NextProtos: protocols}
	return tls.NewListener(listener, cfg)
}
