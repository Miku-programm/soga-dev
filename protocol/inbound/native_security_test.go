package inbound

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	utls "github.com/refraction-networking/utls"
	"golang.org/x/crypto/hkdf"
	"io"
	"math/big"
	"net"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	Q "recovered.local/soga/panel/registry"
	R "recovered.local/soga/reality"
	"testing"
	"time"
)

// 独立 uTLS ClientHello + X25519/HKDF/AES-GCM 认证；验证服务端临时证书的
// Ed25519 公钥 HMAC，不以跳过验证当作 Reality 握手成功。
func nativeRealityClient(t *testing.T, addr string, key *ecdh.PrivateKey) net.Conn {
	t.Helper()
	raw, e := net.DialTimeout("tcp", addr, time.Second)
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { raw.Close() })
	raw.SetDeadline(time.Now().Add(6 * time.Second))
	return nativeRealityOnConn(t, raw, key)
}

func nativeRealityOnConn(t *testing.T, raw net.Conn, key *ecdh.PrivateKey) net.Conn {
	t.Helper()
	var e error
	var authKey []byte
	client := utls.UClient(raw, &utls.Config{ServerName: "hysteria", InsecureSkipVerify: true, VerifyPeerCertificate: func(certs [][]byte, _ [][]*x509.Certificate) error {
		if len(certs) != 1 {
			return fmt.Errorf("unexpected certificate count")
		}
		c, e := x509.ParseCertificate(certs[0])
		if e != nil {
			return e
		}
		pk, ok := c.PublicKey.(ed25519.PublicKey)
		if !ok {
			return fmt.Errorf("not Reality Ed25519 certificate")
		}
		mac := hmac.New(sha512.New, authKey)
		mac.Write(pk)
		if !hmac.Equal(c.Signature, mac.Sum(nil)) {
			return fmt.Errorf("Reality certificate HMAC mismatch")
		}
		return nil
	}}, utls.HelloChrome_133)
	if e = client.BuildHandshakeState(); e != nil {
		t.Fatal(e)
	}
	hello := client.HandshakeState.Hello
	share := client.HandshakeState.State13.KeyShareKeys.Ecdhe
	if share == nil {
		t.Fatal("missing X25519 share")
	}
	authKey, e = share.ECDH(key.PublicKey())
	if e != nil {
		t.Fatal(e)
	}
	if _, e = io.ReadFull(hkdf.New(sha256.New, authKey, hello.Random[:20], []byte("REALITY")), authKey); e != nil {
		t.Fatal(e)
	}
	clear(hello.SessionId)
	clear(hello.Raw[39:71])
	plain := make([]byte, 16)
	copy(plain, []byte{1, 2, 3})
	binary.BigEndian.PutUint32(plain[4:8], uint32(time.Now().Unix()))
	copy(plain[8:], []byte{1, 2, 3, 4, 5, 6, 7, 8})
	block, _ := aes.NewCipher(authKey)
	aead, _ := cipher.NewGCM(block)
	sealed := aead.Seal(nil, hello.Random[20:], plain, hello.Raw)
	copy(hello.SessionId, sealed)
	copy(hello.Raw[39:71], sealed)
	if e = client.Handshake(); e != nil {
		t.Fatal(e)
	}
	return client
}
func nativeRealityDestination(t *testing.T) string {
	t.Helper()
	// Chrome's offered signature algorithms require an ECDSA/RSA destination
	// certificate. The separate Go TLS fixtures may legitimately use Ed25519.
	private, e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if e != nil {
		t.Fatal(e)
	}
	template := &x509.Certificate{SerialNumber: big.NewInt(72), DNSNames: []string{"hysteria"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}}
	der, e := x509.CreateCertificate(rand.Reader, template, template, &private.PublicKey, private)
	if e != nil {
		t.Fatal(e)
	}
	pair := tls.Certificate{Certificate: [][]byte{der}, PrivateKey: private}
	l, e := net.Listen("tcp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { l.Close() })
	go func() {
		for {
			c, e := l.Accept()
			if e != nil {
				return
			}
			go func() {
				defer c.Close()
				c.SetDeadline(time.Now().Add(3 * time.Second))
				s := tls.Server(c, &tls.Config{Certificates: []tls.Certificate{pair}, SessionTicketsDisabled: true, MinVersion: tls.VersionTLS13, CurvePreferences: []tls.CurveID{tls.X25519}, NextProtos: []string{"h2", "http/1.1"}})
				_ = s.Handshake()
			}()
		}
	}()
	return l.Addr().String()
}
func anyTLSMatrixWire(udp bool, port int) (wire, want []byte) {
	auth := sha256.Sum256([]byte("native-anytls-password"))
	wire = append([]byte{}, auth[:]...)
	wire = append(wire, 0, 3, 0xaa, 0xbb, 0xcc)
	wire = append(wire, rawAnyTLSFrame(4, 0, []byte(fmt.Sprintf("v=2\nclient=independent-security-test\npadding-md5=%x", md5.Sum([]byte("stop=0")))))...)
	wire = append(wire, rawAnyTLSFrame(1, 1, nil)...)
	target := binary.BigEndian.AppendUint16([]byte{1, 127, 0, 0, 1}, uint16(port))
	request := append([]byte{}, target...)
	want = bytes.Repeat([]byte("native-transport\x00"), 91)
	if udp {
		magic := "sp.v2.udp-over-tcp.arpa"
		request = append([]byte{3, byte(len(magic))}, []byte(magic)...)
		request = append(request, 0, 0, 1)
		request = append(request, target...)
		want = append(binary.BigEndian.AppendUint16(nil, uint16(len(want))), want...)
	}
	request = append(request, want...)
	wire = append(wire, rawAnyTLSFrame(2, 1, request)...)
	return
}
func readAnyTLSMatrixReply(t *testing.T, c net.Conn, size int) []byte {
	t.Helper()
	var out []byte
	for len(out) < size {
		var h [7]byte
		if _, e := io.ReadFull(c, h[:]); e != nil {
			t.Fatal(e)
		}
		p := make([]byte, int(binary.BigEndian.Uint16(h[5:])))
		if _, e := io.ReadFull(c, p); e != nil {
			t.Fatal(e)
		}
		switch h[0] {
		case 2:
			if binary.BigEndian.Uint32(h[1:5]) != 1 {
				t.Fatal("wrong stream")
			}
			out = append(out, p...)
		case 0, 6, 7, 10:
		default:
			t.Fatalf("unexpected AnyTLS frame %d %q", h[0], p)
		}
	}
	return out
}
func TestNativeVLESSAnyTLSSecurityStartMatrix(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	ssrConfigFixture(t)
	cert, pemKey := hysteriaFixtureCertificate(t)
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM([]byte(cert)) {
		t.Fatal("certificate")
	}
	for _, protocol := range []string{"vless", "anytls"} {
		for _, security := range []string{"tls", "reality"} {
			if protocol == "vless" && security == "tls" {
				continue
			}
			for _, udp := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%s/udp=%t", protocol, security, udp), func(t *testing.T) {
					node := &L.Wr5SVP{EMlXEc: 32, I3k4JnNp7nB: cert, XATiVun: pemKey}
					node.Ck6DxjB = L.VkkFLH{XmIaPj3CzOF: "tcp", WAdJ4ftdZFk: L.D7FHIoaa8s4a(security)}
					node.QaZjDzw95 = L.Tt_Moslrup{J6bt8AOW8fL8: L.D7FHIoaa8s4a(security), UPnToXrM: "stop=0"}
					var private *ecdh.PrivateKey
					var dest string
					if security == "reality" {
						var e error
						private, e = ecdh.X25519().GenerateKey(rand.Reader)
						if e != nil {
							t.Fatal(e)
						}
						dest = nativeRealityDestination(t)
						encoded := base64.RawURLEncoding.EncodeToString(private.Bytes())
						node.Ck6DxjB.BANVIZ72wh = []string{"hysteria"}
						node.Ck6DxjB.CrS6C1 = encoded
						node.Ck6DxjB.IqSfNzea = dest
						node.Ck6DxjB.DxKJ7lSe3m8F = []string{"0102030405060708"}
						node.QaZjDzw95.Jq7YxRdb18Xx = []string{"hysteria"}
						node.QaZjDzw95.CTa7WN = encoded
						node.QaZjDzw95.F9P9aeOu23U = dest
						node.QaZjDzw95.NWbT7wk = []string{"0102030405060708"}
					}
					panel := panelVless(9, "11111111-1111-4111-8111-111111111111")
					if protocol == "anytls" {
						panel = &L.Aga91y{WZ89pTAOrHK4: 9, JAfXWwzYrW: &L.ZtMFlBTZXyD{CBWeRfrO1ahs: "native-anytls-password"}}
					}
					arg := &DV_XPJ4z0{Dmpl48sXM: context.Background(), HG5k0365: &Q.D5gTIZ5aVg{H47AzZsYHB1N: node, VTme1SuHL6i_: L.Sxcf3wu3L{panel}}}
					var core Kz_8ajU7
					var e error
					var addr string
					if protocol == "vless" {
						core, e = oHr1LeATmS(arg)
					} else {
						core, e = aboatPBa(arg)
					}
					if e != nil {
						t.Fatal(e)
					}
					defer core.Stop()
					switch v := core.(type) {
					case *KHonGaqbIy_6:
						v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
					case *R_n5ucb:
						v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
					}
					if e = core.Start(); e != nil {
						t.Fatal(e)
					}
					switch v := core.(type) {
					case *KHonGaqbIy_6:
						addr = v.kX5i8U[0].Addr().String()
					case *R_n5ucb:
						addr = v.qjukoB[0].Addr().String()
					}
					if security == "reality" {
						cacheKey := dest + " hysteria 2"
						defer R.GlobalPostHandshakeRecordsLens.Delete(cacheKey)
						deadline := time.Now().Add(4 * time.Second)
						for {
							if val, ok := R.GlobalPostHandshakeRecordsLens.Load(cacheKey); ok {
								if _, ok = val.([]int); ok {
									break
								}
							}
							if time.Now().After(deadline) {
								t.Fatal("real destination probe did not complete")
							}
							time.Sleep(time.Millisecond)
						}
					}
					targetPort, done := trojanMatrixTarget(t, udp)
					var wire, want []byte
					if protocol == "vless" {
						wire, want = vlessMatrixWire(udp, targetPort)
					} else {
						wire, want = anyTLSMatrixWire(udp, targetPort)
					}
					var c net.Conn
					if security == "reality" {
						c = nativeRealityClient(t, addr, private)
					} else {
						c, e = tls.DialWithDialer(&net.Dialer{Timeout: time.Second}, "tcp", addr, &tls.Config{RootCAs: roots, ServerName: "hysteria"})
						if e != nil {
							t.Fatal(e)
						}
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(5 * time.Second))
					if _, e = c.Write(wire); e != nil {
						t.Fatal(e)
					}
					var got []byte
					if protocol == "vless" {
						got = make([]byte, len(want))
						if _, e = io.ReadFull(c, got); e != nil {
							t.Fatal(e)
						}
					} else {
						got = readAnyTLSMatrixReply(t, c, len(want))
					}
					if !bytes.Equal(got, want) {
						t.Fatal("target payload differs")
					}
					select {
					case e = <-done:
						if e != nil {
							t.Fatal(e)
						}
					case <-time.After(5 * time.Second):
						t.Fatal("target not reached")
					}
				})
			}
		}
	}
}
