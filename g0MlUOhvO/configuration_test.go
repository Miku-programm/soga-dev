package g0MlUOhvO

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"math/big"
	I "recovered.local/soga/I1A6RWw0B3Mq"
	"testing"
	"time"
)

func tunnelCertificate(t *testing.T) *I.ZK5THfX {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{SerialNumber: big.NewInt(1), NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour)}
	cert, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	kd, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return &I.ZK5THfX{Y1mdfY: string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})), SzI2MbNsk: string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: kd}))}
}
func TestOriginalConfigurationDerivationAndOverrides(t *testing.T) {
	in := &InboundConfiguration{Protocol: "tls-tunnel", Method: "aes-256-gcm", Password: "inbound", Certificate: tunnelCertificate(t), ReceiveProxy: 2, ProxyMode: true, Mux: true, MuxNum: 3}
	out := &OutboundConfiguration{Protocol: "ws-tunnel", Method: "aes-128-gcm", Password: "test-tunnel-secret", Host: "fallback.example", ServerName: "override.example", WebSocketPath: "stream", Insecure: true, ProxyMode: true, ProxyHost: "::1", ProxyPort: 443, Mux: true, MuxNum: 7}
	v, e := FZzfvtn(in, out)
	if e != nil {
		t.Fatal(e)
	}
	if !v.ufiMqR6x || v.hZORuZstvS07 || v.rj1DhSBHZ7O || !v.IsMux() || v.MuxNum() != 7 || !v.TunnelRecvProxy() || !v.IsProxyMode() || v.ProxyAddr().Join() != "[::1]:443" {
		t.Fatalf("flags: %+v", v)
	}
	if v.p1Wtyln2GJ6 != "override.example" || !v.psdXJ2Mi || v.lyXfKPaz4m23 != "/stream" || v._rKF9rrYaf8a == nil {
		t.Fatal("overrides/retained certificate", v)
	}
	// 独立 Python hashlib.sha3_512 / HMAC-HKDF 计算的固定向量。
	if hex.EncodeToString(v.oWvGY4I9cA) != "4a6837ff5986fab660da2841772dc156" || hex.EncodeToString(v.x5Lor8hDyn[:]) != "9db29d0cad3a9fc2aea9cacd1135071c" {
		t.Fatalf("derivation %x %x", v.oWvGY4I9cA, v.x5Lor8hDyn)
	}
	out.ServerName = ""
	out.Protocol = "quic-tunnel"
	out.WebSocketPath = ""
	out.ProxyMode = false
	v, e = FZzfvtn(in, out)
	if e != nil {
		t.Fatal(e)
	}
	if v.IsMux() || !v.IsQuic() || !v.rj1DhSBHZ7O || v.ProxyAddr() != nil || v.p1Wtyln2GJ6 != "fallback.example" || v.lyXfKPaz4m23 != "/" {
		t.Fatal("quic/fallback", v)
	}
	// 第二输入不能掩盖先发生的入站证书失败。
	in.Certificate = nil
	if _, e = FZzfvtn(in, out); e == nil || e.Error() != "tls tunnel but not config cert" {
		t.Fatal(e)
	}
}
func TestOriginalConfigurationErrorOrderAndModes(t *testing.T) {
	if _, e := FZzfvtn(nil, nil); e == nil || e.Error() != "please set password for tunnel type" {
		t.Fatal(e)
	}
	for _, protocol := range []string{"tcp-tunnel", "tls", "tls-tunnel", "ws-tunnel", "ws-tls-tunnel", "quic-tunnel", "unknown"} {
		in := &InboundConfiguration{Protocol: protocol, Method: "aes-128-gcm", Password: "p", Mux: true, ReceiveProxy: -1, Certificate: tunnelCertificate(t)}
		v, e := FZzfvtn(in, nil)
		if e != nil {
			t.Fatal(protocol, e)
		}
		if v.IsMux() != (protocol != "quic-tunnel") || v.TunnelRecvProxy() {
			t.Fatal(protocol, v)
		}
		if v.ufiMqR6x != (protocol == "ws-tunnel" || protocol == "ws-tls-tunnel") {
			t.Fatal(protocol)
		}
		if v.rj1DhSBHZ7O != (protocol == "tls" || protocol == "tls-tunnel" || protocol == "quic-tunnel" || protocol == "ws-tls-tunnel") {
			t.Fatal(protocol)
		}
	}
	in := &InboundConfiguration{Protocol: "tcp-tunnel", Method: "unknown", Password: ""}
	if _, e := FZzfvtn(in, nil); e == nil || e.Error() != "please set password for tunnel type" {
		t.Fatal(e)
	}
	in.Password = "p"
	if _, e := FZzfvtn(in, nil); e == nil || e.Error() != "not support method: unknown" {
		t.Fatal(e)
	}
}
