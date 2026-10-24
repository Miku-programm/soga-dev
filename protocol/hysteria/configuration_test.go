package hysteria

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"math/big"
	"reflect"
	"testing"
	"unsafe"
)

func TestCertificateValidationAndAliasing(t *testing.T) {
	get := func(*tls.ClientHelloInfo) (*tls.Certificate, error) { t.Fatal("unexpected callback"); return nil, nil }
	if c, e := LnSg9qSRSAT(&tls.Config{GetCertificate: get}, nil, 0, 0); c != nil || e == nil || e.Error() != "must set at least one certificate" {
		t.Fatalf("missing: %v %v", c, e)
	}
	bad := &tls.Config{Certificates: []tls.Certificate{{Certificate: [][]byte{{0}}}}}
	if c, e := LnSg9qSRSAT(bad, nil, 0, 0); c != nil || e == nil || bad.Certificates[0].Leaf != nil {
		t.Fatal("invalid DER accepted")
	}
	pub, key, e := ed25519.GenerateKey(rand.Reader)
	if e != nil {
		t.Fatal(e)
	}
	template := &x509.Certificate{SerialNumber: big.NewInt(1)}
	der, e := x509.CreateCertificate(rand.Reader, template, template, pub, key)
	if e != nil {
		t.Fatal(e)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{{Certificate: [][]byte{der}}}, GetCertificate: get, EncryptedClientHelloKeys: []tls.EncryptedClientHelloKey{{Config: []byte{1}}}}
	c, e := LnSg9qSRSAT(tlsConfig, nil, 1, 2)
	if e != nil {
		t.Fatal(e)
	}
	if c.DGraqdoRv9ot.GinGbR[0].Leaf == nil || c.DGraqdoRv9ot.GinGbR[0].Leaf != tlsConfig.Certificates[0].Leaf {
		t.Fatal("Leaf not cached in original slice")
	}
	if &c.DGraqdoRv9ot.GinGbR[0] != &tlsConfig.Certificates[0] || &c.DGraqdoRv9ot.FQaEaYFoFF[0] != &tlsConfig.EncryptedClientHelloKeys[0] {
		t.Fatal("original slice alias lost")
	}
	if reflect.ValueOf(get).Pointer() != reflect.ValueOf(c.DGraqdoRv9ot.EG8WQwgS).Pointer() {
		t.Fatal("callback changed")
	}
	// 原实现直接索引第一张证书的 DER；空 DER 链仍会触发边界 panic。
	defer func() {
		if recover() == nil {
			t.Fatal("empty DER did not panic")
		}
	}()
	LnSg9qSRSAT(&tls.Config{Certificates: []tls.Certificate{{}}}, nil, 0, 0)
}

func TestOriginalConfigurationLayout(t *testing.T) {
	var c InfdqOmZK
	got := []uintptr{unsafe.Sizeof(c), unsafe.Sizeof(c.DGraqdoRv9ot), unsafe.Sizeof(c.I3PHCWIdg8EH), unsafe.Offsetof(c.FYaeQJDXC), unsafe.Offsetof(c.TDNLdWhC), unsafe.Offsetof(c.QlDN7UZH), unsafe.Offsetof(c.Tnw69F), unsafe.Offsetof(c.IDZiaXEBjVIW), unsafe.Offsetof(c.DGm6Idekj)}
	want := []uintptr{176, 56, 56, 112, 128, 129, 136, 144, 160}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("layout %v", got)
	}
}
