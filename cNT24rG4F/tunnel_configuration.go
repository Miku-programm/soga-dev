package cNT24rG4F

import (
	"fmt"
	"os"
	D "recovered.local/soga/DDYEn9Z_b"
	I "recovered.local/soga/I1A6RWw0B3Mq"
	T "recovered.local/soga/g0MlUOhvO"
	L "recovered.local/soga/muc279Mpa6S4"
)

// R3DjPWSnt 对应 0x1330080。原全局配置与 Q0SDIN 条件均保留。
func R3DjPWSnt() (*T.VTl94pzz2IkB, error) {
	if !configAt509E300.JHXrVGNjg31 || Q0SDIN() {
		return nil, nil
	}
	mode := configAt509E300.GbyJkc
	switch mode {
	case "tcp-tunnel", "tls-tunnel", "ws-tunnel", "ws-tls-tunnel", "quic-tunnel":
	default:
		return nil, fmt.Errorf("not support tunnel type: %v", mode)
	}
	var cert *I.ZK5THfX
	switch mode {
	case "tls-tunnel", "quic-tunnel", "ws-tls-tunnel":
		certPEM, err := os.ReadFile(configAt509F9C0.C2BNS6_QJC)
		if err != nil {
			return nil, err
		}
		keyPEM, err := os.ReadFile(configAt509F9C0.JISdid54S5t)
		if err != nil {
			return nil, err
		}
		cert = &I.ZK5THfX{Y1mdfY: string(certPEM), SzI2MbNsk: string(keyPEM)}
	}
	return T.FZzfvtn(&T.InboundConfiguration{
		Protocol: mode, Method: D.RQvD2v3oZ(configAt509E300.ZjL7PiTY),
		Password: configAt509E300.D38NXN3uV, WebSocketPath: configAt509E300.LYSz_mSWFTQp,
		Certificate: cert, ReceiveProxy: 1, Mux: configAt509E300.RtHB6a,
	}, nil)
}

// QitK3yi 对应 0x1330580。仅记录当前配置，原函数并不写入任何状态。
// 文本来自受限 decoder 仿真；保留原记录密码字段的行为。
func QitK3yi() {
	L.TRTbrquD(0, "enable tunnel")
	L.TRTbrquD(0, "tunnel type:", configAt509E300.GbyJkc)
	L.TRTbrquD(0, "tunnel mux:", configAt509E300.RtHB6a)
	L.TRTbrquD(0, "tunnel method:", D.RQvD2v3oZ(configAt509E300.ZjL7PiTY))
	L.TRTbrquD(0, "tunnel password:", configAt509E300.D38NXN3uV)
	if configAt509E300.GbyJkc == "ws-tunnel" || configAt509E300.GbyJkc == "ws-tls-tunnel" {
		L.TRTbrquD(0, "tunnel ws path:", configAt509E300.LYSz_mSWFTQp)
	}
}
