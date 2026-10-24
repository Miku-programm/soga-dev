package G6bHD5JA

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"recovered.local/soga/BaEWsPy"
	"recovered.local/soga/EiDD1sTNvpQ"
	"recovered.local/soga/HxKViEp2a"
)

// 静态指针 0x5029428 -> 0x5045180：仅 Timeout 非零，等于 60 秒。
var clientAt5029428 = &http.Client{Timeout: 60 * time.Second}
var clientOnceAt50C7B70 sync.Once

// WuvamjL5 @ 0x15a40e0，原调用方通过 sync.Once 执行。
func WuvamjL5() {
	resolver := BaEWsPy.DAlA5L(context.Background(), &BaEWsPy.PQL1aLrtju{
		Zfr5Q8ybETV: []string{
			"https://8.8.8.8/dns-query", "https://1.1.1.1/dns-query", "https://208.67.222.222/dns-query",
			"tcp-tls://8.8.8.8:853", "tcp-tls://1.1.1.1:853", "tcp-tls://208.67.222.222:853",
			"tcp://8.8.8.8:53", "tcp://1.1.1.1:53", "tcp://208.67.222.222:53",
			"8.8.8.8:53", "1.1.1.1:53", "208.67.222.222:53",
			"2001:4860:4860::8888", "2400:3200::1", "2400:3200:baba::1", "local",
		}, X6T7XOmp: "ipv4_first", C_TYygjEa0: time.Minute,
	})
	dialer := &net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second, Control: EiDD1sTNvpQ.CyQmxg(EiDD1sTNvpQ.UA1QxZBYiPM)}
	clientAt5029428.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			// func1 @ 0x15b4940：顺序尝试每个解析结果，首个成功即返回。
			host, port, err := net.SplitHostPort(address)
			if err != nil {
				return nil, err
			}
			ips, err := resolver.Resolve(host)
			if err != nil {
				return nil, err
			}
			var failures []error
			for _, ip := range ips {
				connection, err := dialer.DialContext(ctx, network, net.JoinHostPort(ip.String(), port))
				if err == nil {
					return connection, nil
				}
				failures = append(failures, err)
			}
			return nil, HxKViEp2a.T2i1ItMaa(failures...)
		},
		DisableKeepAlives: true, MaxIdleConns: 5,
		IdleConnTimeout: 10 * time.Second, ResponseHeaderTimeout: 30 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second, ExpectContinueTimeout: time.Second,
	}
}
