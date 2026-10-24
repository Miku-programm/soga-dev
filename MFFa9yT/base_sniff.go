package MFFa9yT

import (
	"fmt"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/nzATvajt"
)

// Sniff 对应0x1d3f200。探测未成功时返回空且无错误；只有已识别的其他
// 类型进入包检查，HTTP/TLS/QUIC直接返回原探测器主机结果。
func (b *JWN154qJ4) Sniff(data []byte, user nzATvajt.PgaNd3014U, network G6bHD5JA.CjhtWb) (string, error) {
	result, err := b.cEUo2gYKI.Sniff(data, network)
	if err != nil {
		return "", nil
	}
	switch result.Type() {
	case "bt":
		if cNT24rG4F.IsInspectionExempt(user.GetId()) {
			return "", nil
		}
		return "", fmt.Errorf("user %v access bt protocol", user.GetId())
	case "tls", "http", "quic":
		host := result.Value()
		muc279Mpa6S4.JucPcHhJ("detect %v host: %v", result.Type(), host)
		return host, nil
	}
	if b.TestPacket(data, user) {
		return "", fmt.Errorf("user %v access invalid protocol", user.GetId())
	}
	return "", nil
}
