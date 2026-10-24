package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestTLSCipherSelectionOriginalRules(t *testing.T) {
	old := configAt509EBE0
	defer func() { configAt509EBE0 = old }()
	configAt509EBE0.KVniDErTtx = " TLS_AES_128_GCM_SHA256 :0x1302::0X1303:TLS_RSA_WITH_RC4_128_SHA:0x0005:0x10000:0xzz:0x1301 "
	want := []uint16{0x1301, 0x1302, 0x1301}
	if got := B02ZlILA8i(); !reflect.DeepEqual(got, want) {
		t.Fatalf("%x != %x", got, want)
	}
	configAt509EBE0.KVniDErTtx = " : "
	if B02ZlILA8i() != nil {
		t.Fatal("empty cipher result must remain nil")
	}
}
func TestTLSConfigurationVersionAndAliasing(t *testing.T) {
	old := configAt509EBE0
	defer func() { configAt509EBE0 = old }()
	configAt509EBE0 = jkGZ0JHJw{RS1kBzpBezso: []string{"h2", "http/1.1"}, JOyIs3s9WAcA: true}
	cert := tls.Certificate{Certificate: [][]byte{{1, 2, 3}}}
	for _, tc := range []struct {
		s string
		v uint16
	}{{"1.0", 769}, {"1.1", 770}, {"1.2", 771}, {"1.3", 772}, {"TLS1.3", 0}, {" 1.3", 0}, {"", 0}} {
		configAt509EBE0.CJri5kB26 = tc.s
		configAt509EBE0.NH7oB3o7 = tc.s
		c := c2dSGQJ4q(cert)
		if c.MinVersion != tc.v || c.MaxVersion != tc.v || !c.PreferServerCipherSuites {
			t.Fatal(tc, c)
		}
		if &c.NextProtos[0] != &configAt509EBE0.RS1kBzpBezso[0] || &c.Certificates[0].Certificate[0][0] != &cert.Certificate[0][0] {
			t.Fatal("original slice alias lost")
		}
	}
	configAt509EBE0.RS1kBzpBezso = make([]string, 0, 3)
	if c2dSGQJ4q(cert).NextProtos != nil {
		t.Fatal("empty ALPN must not be assigned")
	}
}
func TestTLSPEMAndFileFallback(t *testing.T) {
	oldTLS, oldCert := configAt509EBE0, configAt509F9C0
	defer func() { configAt509EBE0 = oldTLS; configAt509F9C0 = oldCert }()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "localhost"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), DNSNames: []string{"localhost"}, KeyUsage: x509.KeyUsageDigitalSignature, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	kd, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	cp := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	kp := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: kd})
	dir := t.TempDir()
	configAt509F9C0.C2BNS6_QJC = filepath.Join(dir, "cert.pem")
	configAt509F9C0.JISdid54S5t = filepath.Join(dir, "key.pem")
	for path, data := range map[string][]byte{configAt509F9C0.C2BNS6_QJC: cp, configAt509F9C0.JISdid54S5t: kp} {
		if err := os.WriteFile(path, data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	for _, pair := range [][2]string{{string(cp), string(kp)}, {"", ""}, {"invalid", ""}, {"", "invalid"}} {
		c, err := XmxN0NDuhcA(pair[0], pair[1])
		if err != nil || len(c.Certificates) != 1 {
			t.Fatal(err)
		}
	}
	if c, err := XmxN0NDuhcA("invalid", "invalid"); err == nil || c != nil {
		t.Fatal("two supplied invalid PEMs must not fall back")
	}
	if err := os.Remove(configAt509F9C0.C2BNS6_QJC); err != nil {
		t.Fatal(err)
	}
	if c, err := Ks6Wk5LnO(); err == nil || c != nil {
		t.Fatal("missing file")
	}
}
