package inbound

import (
	"net"
	"testing"
)

// 独立构造 PROXY v2 的 IPv4/UDP 头，验证原中转地址映射与两种混淆的装配顺序。
type hysteriaProxyClient struct{ net.PacketConn }

func (c *hysteriaProxyClient) WriteTo(b []byte, addr net.Addr) (int, error) {
	header := []byte{'\r', '\n', '\r', '\n', 0, '\r', '\n', 'Q', 'U', 'I', 'T', '\n', 0x21, 0x12, 0, 12, 192, 0, 2, 9, 198, 51, 100, 1, 0x04, 0xd2, 0x01, 0xbb}
	_, err := c.PacketConn.WriteTo(append(header, b...), addr)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func TestNativeHysteriaForcedUDPProxyTCPUDP(t *testing.T) {
	t.Setenv("udp_proxy_protocol", "true")
	t.Setenv("force_proxy_protocol", "true")
	TestNativeHysteriaConstructStartTCPUDPStop(t)
}
