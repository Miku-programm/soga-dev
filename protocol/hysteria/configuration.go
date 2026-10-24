package hysteria

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"math"
	"net"
	"net/http"
	U "recovered.local/soga/user"
	"time"
)

// 以下类型、字段和布局来自样本的 Go 类型元数据。
type BQrW66oHM interface {
	Authenticate(net.Addr, string, uint64) (U.PgaNd3014U, bool)
}

type DxvcQLaW struct {
	GinGbR     []tls.Certificate
	EG8WQwgS   func(*tls.ClientHelloInfo) (*tls.Certificate, error)
	FQaEaYFoFF []tls.EncryptedClientHelloKey
}

type VhfHmar5qsKt struct {
	UIFPM249f2   uint64
	ET3EbVcYR    uint64
	UCjF7o6elmz  uint64
	CtEnIL0gS    uint64
	YldcvpLEKua1 time.Duration
	VAZK_diUd    int64
	EJDv3cXV8    bool
}

type PfJbE3f struct {
	GSKeMid    uint64
	OfLAiHnwZo uint64
}

type InfdqOmZK struct {
	DGraqdoRv9ot DxvcQLaW
	I3PHCWIdg8EH VhfHmar5qsKt
	FYaeQJDXC    PfJbE3f
	TDNLdWhC     bool
	QlDN7UZH     bool
	Tnw69F       time.Duration
	IDZiaXEBjVIW BQrW66oHM
	DGm6Idekj    http.Handler
}

// LnSg9qSRSAT 恢复自 0x1c2eb00。原实现要求已有至少一张证书，
// 并就地缓存第一张证书的 Leaf；单独设置 GetCertificate 不满足此条件。
func LnSg9qSRSAT(config *tls.Config, auth BQrW66oHM, tx, rx float64) (*InfdqOmZK, error) {
	if len(config.Certificates) == 0 {
		return nil, errors.New("must set at least one certificate")
	}
	if config.Certificates[0].Leaf == nil {
		leaf, err := x509.ParseCertificate(config.Certificates[0].Certificate[0])
		if err != nil {
			return nil, err
		}
		config.Certificates[0].Leaf = leaf
	}
	return &InfdqOmZK{
		DGraqdoRv9ot: DxvcQLaW{config.Certificates, config.GetCertificate, config.EncryptedClientHelloKeys},
		I3PHCWIdg8EH: VhfHmar5qsKt{8 << 20, 8 << 20, 20 << 20, 20 << 20, 30 * time.Second, 1024, false},
		FYaeQJDXC:    PfJbE3f{uint64(math.Max(tx, 0) * 1024 * 1024 * 0.125), uint64(math.Max(rx, 0) * 1024 * 1024 * 0.125)},
		Tnw69F:       time.Minute,
		IDZiaXEBjVIW: auth,
	}, nil
}
