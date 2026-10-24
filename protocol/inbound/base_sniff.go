package inbound

import (
	"fmt"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/user"
)

// Sniff 对应0x1d3f200。探测未成功时返回空且无错误；只有已识别的其他
// 类型进入包检查，HTTP/TLS/QUIC直接返回原探测器主机结果。
func (b *JWN154qJ4) Sniff(data []byte, user user.PgaNd3014U, network relay.CjhtWb) (string, error) {
	result, err := b.cEUo2gYKI.Sniff(data, network)
	if err != nil {
		return "", nil
	}
	switch result.Type() {
	case "bt":
		if config.IsInspectionExempt(user.GetId()) {
			return "", nil
		}
		return "", fmt.Errorf("user %v access bt protocol", user.GetId())
	case "tls", "http", "quic":
		host := result.Value()
		logx.JucPcHhJ("detect %v host: %v", result.Type(), host)
		return host, nil
	}
	if b.TestPacket(data, user) {
		return "", fmt.Errorf("user %v access invalid protocol", user.GetId())
	}
	return "", nil
}
