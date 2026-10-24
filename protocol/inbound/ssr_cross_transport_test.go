//go:build linux

package inbound

import "testing"

// 补足已有按单维度验证遗漏的交叉项。独立 outbound 客户端经过原生
// SSR Start/UpdateAPIData、协议/加密/混淆组合，到真实 TCP 目标并回传。
func TestNativeSSRIndependentCipherProtocolObfsCrossProduct(t *testing.T) {
	if !liveProcess(t) {
		return
	}
	ssrConfigFixture(t)
	for _, method := range []string{"none", "aes-192-cfb", "aes-256-cfb", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr", "rc4-md5", "chacha20-ietf"} {
		for _, protocol := range []string{"origin", "auth_aes128_md5", "auth_aes128_sha1", "auth_chain_a", "auth_chain_b"} {
			for _, obfs := range []string{"http_simple", "http_post", "tls1.2_ticket_auth"} {
				t.Run(method+"/"+protocol+"/"+obfs, func(t *testing.T) {
					ssrNativeMatrixExchange(t, method, protocol, obfs, false)
				})
			}
		}
	}
}
