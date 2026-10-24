package reality

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
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	stdtls "crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"net"
	"strings"
	"testing"
	"time"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/crypto/hkdf"
)

// 实际 loopback TCP + TLS 1.3 目标站 + uTLS ClientHello；不启动原始样本。
func TestRecoveredRealityAuthenticatedTCP(t *testing.T) {
	for _, tc := range []struct {
		name   string
		curve  stdtls.CurveID
		reject string
	}{{"X25519", stdtls.X25519, ""}, {"X25519MLKEM768", stdtls.X25519MLKEM768, ""}, {"wrong-short-id", stdtls.X25519, "short"}, {"stale-time", stdtls.X25519, "time"}, {"client-version-too-low", stdtls.X25519, "min"}, {"client-version-too-high", stdtls.X25519, "max"}, {"unknown-sni", stdtls.X25519, "sni"}} {
		t.Run(tc.name, func(t *testing.T) {
			curve := tc.curve
			private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			if err != nil {
				t.Fatal(err)
			}
			public := &private.PublicKey
			certT := &x509.Certificate{SerialNumber: big.NewInt(1), DNSNames: []string{"fixture.invalid"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}}
			cert, err := x509.CreateCertificate(rand.Reader, certT, certT, public, private)
			if err != nil {
				t.Fatal(err)
			}
			target, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				t.Fatal(err)
			}
			defer target.Close()
			targetDone := make(chan struct{})
			go func() {
				defer close(targetDone)
				c, e := target.Accept()
				if e != nil {
					return
				}
				defer c.Close()
				c.SetDeadline(time.Now().Add(5 * time.Second))
				s := stdtls.Server(c, &stdtls.Config{Certificates: []stdtls.Certificate{{Certificate: [][]byte{cert}, PrivateKey: private}}, SessionTicketsDisabled: true, MinVersion: stdtls.VersionTLS13, CurvePreferences: []stdtls.CurveID{curve}, NextProtos: []string{"h2", "http/1.1"}})
				_ = s.Handshake()
			}()
			listen, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				t.Fatal(err)
			}
			defer listen.Close()
			key, err := ecdh.X25519().GenerateKey(rand.Reader)
			if err != nil {
				t.Fatal(err)
			}
			shortID := [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
			version := [3]byte{1, 2, 3}
			cfg := &Config{SessionTicketsDisabled: true, DialContext: (&net.Dialer{Timeout: time.Second}).DialContext, Type: "tcp", Dest: target.Addr().String(), ServerNames: map[string]bool{"fixture.invalid": true}, PrivateKey: key.Bytes(), ShortIds: map[[8]byte]bool{shortID: true}, MinClientVer: []byte{1}, MaxTimeDiff: time.Minute}
			switch tc.reject {
			case "short":
				cfg.ShortIds = map[[8]byte]bool{}
			case "time":
				cfg.MaxTimeDiff = time.Nanosecond
			case "min":
				cfg.MinClientVer = []byte{2, 0, 0}
			case "max":
				cfg.MaxClientVer = []byte{1, 0, 0}
			case "sni":
				cfg.ServerNames = map[string]bool{}
			}
			cacheKey := cfg.Dest + " fixture.invalid 2"
			GlobalPostHandshakeRecordsLens.Store(cacheKey, []int{})
			defer GlobalPostHandshakeRecordsLens.Delete(cacheKey)
			payload := bytes.Repeat([]byte("original-fork-TLS-record/"), 1000)
			reply := bytes.Repeat([]byte("return-record/"), 1700)
			serverDone := make(chan error, 1)
			go func() {
				c, e := listen.Accept()
				if e != nil {
					serverDone <- e
					return
				}
				defer c.Close()
				c.SetDeadline(time.Now().Add(5 * time.Second))
				s, e := Server(context.Background(), c, cfg)
				if e != nil {
					serverDone <- e
					return
				}
				if s.ClientVer != version || s.ClientShortId != shortID || !s.ConnectionState().HandshakeComplete {
					serverDone <- fmt.Errorf("authenticated state mismatch")
					return
				}
				b := make([]byte, len(payload))
				if _, e = io.ReadFull(s, b); e == nil && !bytes.Equal(b, payload) {
					e = fmt.Errorf("payload mismatch")
				}
				if e == nil {
					_, e = s.Write(reply)
				}
				serverDone <- e
			}()
			raw, err := net.Dial("tcp", listen.Addr().String())
			if err != nil {
				t.Fatal(err)
			}
			defer raw.Close()
			raw.SetDeadline(time.Now().Add(5 * time.Second))
			var authKey []byte
			client := utls.UClient(raw, &utls.Config{ServerName: "fixture.invalid", InsecureSkipVerify: true, VerifyPeerCertificate: func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
				if len(rawCerts) != 1 {
					return fmt.Errorf("cert count")
				}
				c, e := x509.ParseCertificate(rawCerts[0])
				if e != nil {
					return e
				}
				pk, ok := c.PublicKey.(ed25519.PublicKey)
				if !ok {
					return fmt.Errorf("not REALITY Ed25519 certificate")
				}
				mac := hmac.New(sha512.New, authKey)
				mac.Write(pk)
				if !hmac.Equal(c.Signature, mac.Sum(nil)) {
					return fmt.Errorf("REALITY certificate HMAC mismatch")
				}
				return nil
			}}, utls.HelloChrome_133)
			if err = client.BuildHandshakeState(); err != nil {
				t.Fatal(err)
			}
			hello := client.HandshakeState.Hello
			share := client.HandshakeState.State13.KeyShareKeys.Ecdhe
			if share == nil {
				t.Fatal("missing X25519 share")
			}
			authKey, err = share.ECDH(key.PublicKey())
			if err != nil {
				t.Fatal(err)
			}
			if _, err = io.ReadFull(hkdf.New(sha256.New, authKey, hello.Random[:20], []byte("REALITY")), authKey); err != nil {
				t.Fatal(err)
			}
			clear(hello.SessionId)
			clear(hello.Raw[39:71])
			plain := make([]byte, 16)
			copy(plain, version[:])
			binary.BigEndian.PutUint32(plain[4:8], uint32(time.Now().Unix()))
			copy(plain[8:], shortID[:])
			block, _ := aes.NewCipher(authKey)
			aead, _ := cipher.NewGCM(block)
			sealed := aead.Seal(nil, hello.Random[20:], plain, hello.Raw)
			copy(hello.SessionId, sealed)
			copy(hello.Raw[39:71], sealed)
			err = client.Handshake()
			if tc.reject != "" {
				if err == nil {
					t.Fatal("unauthorized client accepted")
				}
				raw.Close()
				e := <-serverDone
				if e == nil || !strings.Contains(e.Error(), "REALITY: processed invalid connection") {
					t.Fatalf("server rejection: %v", e)
				}
				<-targetDone
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if _, err = client.Write(payload); err != nil {
				t.Fatalf("write: %v", err)
			}
			got := make([]byte, len(reply))
			if _, err = io.ReadFull(client, got); err != nil {
				t.Fatalf("read: %v", err)
			}
			if !bytes.Equal(got, reply) {
				t.Fatal("reply mismatch")
			}
			if err = <-serverDone; err != nil {
				t.Fatal(err)
			}
			<-targetDone
		})
	}
}
