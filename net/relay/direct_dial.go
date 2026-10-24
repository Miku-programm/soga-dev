package relay

import (
	"context"
	"net"
	"recovered.local/soga/net/sockopt"
	"time"
)

// DE9CHPNLbK @ 0x15b2d80。原函数使用 Background，不接收外部 context。
// Control 函数表 0x37757a8/0x37757a0 分别为 reuse 和 no-port-bind。
func DE9CHPNLbK(local *net.TCPAddr, address string) (net.Conn, error) {
	dialer := net.Dialer{Timeout: 30 * time.Second, KeepAlive: 60 * time.Second, LocalAddr: local, Control: sockopt.CyQmxg(sockopt.UA1QxZBYiPM, sockopt.ICQr8wkaYl)}
	return dialer.DialContext(context.Background(), "tcp", address)
}

// Y0B53iB @ 0x15b3940。nil 和非 nil 空 IP 按原指针判定区分。
func Y0B53iB(local, remote net.IP) bool {
	if local == nil || local.IsUnspecified() {
		return true
	}
	if remote == nil {
		return false
	}
	return (local.To4() != nil) == (remote.To4() != nil)
}
