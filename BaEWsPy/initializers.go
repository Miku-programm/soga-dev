package BaEWsPy

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/EiDD1sTNvpQ"
)

// 静态指针 0x5029cb0 -> 0x5060cd0；Control 由 init @ 0x1585620 赋值。
var dialerAt5029CB0 = &net.Dialer{Timeout: 10 * time.Second, KeepAlive: 60 * time.Second, Control: EiDD1sTNvpQ.CyQmxg(EiDD1sTNvpQ.UA1QxZBYiPM)}
var defaultResolverAt5096320 *D23y2888a_
var ipv4AvailableAt50C6D60 atomic.Int32
var ipv6AvailableAt50C6D64 atomic.Int32
var availabilityMutexAt50C7910 sync.Mutex
var ipv4FailuresAt50C7918 int
var ipv6FailuresAt50C7920 int

func init() {
	// init.0 @ 0x15899a0。20 个原始字符串按原顺序受限解码恢复。
	defaultResolverAt5096320 = DAlA5L(context.Background(), &PQL1aLrtju{
		Zfr5Q8ybETV: []string{
			"tcp://8.8.8.8:53", "tcp://8.8.4.4:53", "tcp://1.1.1.1:53", "tcp://208.67.222.222:53", "tcp://208.67.220.220:53",
			"8.8.8.8:53", "8.8.4.4:53", "1.1.1.1:53", "208.67.222.222:53", "208.67.220.220:53",
			"tcp-tls://8.8.8.8:853", "tcp-tls://8.8.4.4:853", "tcp-tls://1.1.1.1:853", "tcp-tls://208.67.222.222:853", "tcp-tls://208.67.220.220:853",
			"https://8.8.8.8/dns-query", "https://8.8.4.4/dns-query", "https://1.1.1.1/dns-query", "https://208.67.222.222/dns-query", "https://208.67.220.220/dns-query",
		}, X6T7XOmp: "ipv4_first", C_TYygjEa0: 30 * time.Second, DMa5Vg24rVR: false, ZBMqaLb: true,
	})
	// init.1 @ 0x1592040。UDP Dial/Close 检查本机路由可用性，不发送 DNS 查询。
	ipv4AvailableAt50C6D60.Store(1)
	ipv6AvailableAt50C6D64.Store(1)
	ipv4FailuresAt50C7918, ipv6FailuresAt50C7920 = 2, 2
	fBTPTo()
	CAodbhZf.SwUrlM3M(context.Background(), fBTPTo, time.Minute)
}

// fBTPTo @ 0x15920c0。保持 IPv4、IPv6 两次探测的先后顺序。
func fBTPTo() {
	ipv4 := qYgUF5Z("udp4", "8.8.8.8:53")
	ipv6 := qYgUF5Z("udp6", "[2001:4860:4860::8888]:53")
	egKhwF(ipv4, ipv6)
}

// egKhwF @ 0x1592180。两种地址族同时失败时不改变已有状态。
func egKhwF(ipv4, ipv6 bool) {
	if !ipv4 && !ipv6 {
		return
	}
	availabilityMutexAt50C7910.Lock()
	defer availabilityMutexAt50C7910.Unlock()
	if ipv4 {
		ipv4FailuresAt50C7918 = 0
	} else {
		ipv4FailuresAt50C7918++
		if ipv4FailuresAt50C7918 > 3 {
			ipv4FailuresAt50C7918 = 3
		}
	}
	if ipv6 {
		ipv6FailuresAt50C7920 = 0
	} else {
		ipv6FailuresAt50C7920++
		if ipv6FailuresAt50C7920 > 3 {
			ipv6FailuresAt50C7920 = 3
		}
	}
	if ipv4FailuresAt50C7918 < 3 {
		ipv4AvailableAt50C6D60.Store(1)
	} else {
		ipv4AvailableAt50C6D60.Store(0)
	}
	if ipv6FailuresAt50C7920 < 3 {
		ipv6AvailableAt50C6D64.Store(1)
	} else {
		ipv6AvailableAt50C6D64.Store(0)
	}
}

// qYgUF5Z @ 0x15922e0，连接后立即 Close，关闭错误忽略。
func qYgUF5Z(network, address string) bool {
	connection, err := net.Dial(network, address)
	if err != nil {
		return false
	}
	_ = connection.Close()
	return true
}
