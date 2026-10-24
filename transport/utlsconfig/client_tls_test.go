package utlsconfig

import (
	"crypto/tls"
	utls "github.com/refraction-networking/utls"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestOriginalClientTLSConfigurationAndActualHandshake(t *testing.T) {
	saved := clientTLSSettingsAt509D928
	defer func() { clientTLSSettingsAt509D928 = saved }()
	if saved.enabled || saved.useUTLS || saved.nextProtos == nil || saved.cipherSuites == nil || reflect.TypeOf(saved).Size() != 72 {
		t.Fatal("original global initialization")
	}
	fixture := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	cert := fixture.TLS.Certificates[0]
	fixture.Close()
	for _, mode := range []string{"default", "configured", "randomized-utls"} {
		t.Run(mode, func(t *testing.T) {
			clientTLSSettingsAt509D928 = saved
			if mode != "default" {
				clientTLSSettingsAt509D928 = clientTLSSettings{enabled: true, nextProtos: []string{"native"}, minVersion: tls.VersionTLS12, maxVersion: tls.VersionTLS12, cipherSuites: []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256}, sessionCache: true, useUTLS: mode == "randomized-utls"}
			}
			listener, e := net.Listen("tcp", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer listener.Close()
			a, e := net.Dial("tcp", listener.Addr().String())
			if e != nil {
				t.Fatal(e)
			}
			b, e := listener.Accept()
			if e != nil {
				t.Fatal(e)
			}
			defer a.Close()
			defer b.Close()
			a.SetDeadline(time.Now().Add(3 * time.Second))
			b.SetDeadline(time.Now().Add(3 * time.Second))
			rawClient := JZejsfshGB(a, true, "native.test")
			if mode == "randomized-utls" {
				c, ok := rawClient.(*utls.UConn)
				if !ok || c.ClientHelloID.Client != "Randomized" || c.ClientHelloID.Version != "0" {
					t.Fatal(rawClient)
				}
			} else {
				if _, ok := rawClient.(*tls.Conn); !ok {
					t.Fatal(rawClient)
				}
			}
			protocols := []string{"h2", "http/1.1"}
			if mode == "configured" {
				protocols = []string{"native"}
			}
			// 固定服务端为经典曲线，排除 X25519MLKEM768。utls v1.8.1 的随机指纹
			// (Randomized) 会宣告 MLKEM768 却不一定发送对应 key share，Go 1.24+ 标准库
			// 服务端默认优先 MLKEM768 并据此发起 HelloRetryRequest，而 utls 的 HRR 路径
			// 不支持该组，握手约 30% 概率失败。随机指纹在 TLS 1.3 下必然包含 P-256/P-384/
			// X25519，服务端限定经典曲线后总能命中双方共有组，握手稳定。这是测试侧约束，
			// 不改动生产的 JZejsfshGB 指纹行为。
			server := tls.Server(b, &tls.Config{Certificates: []tls.Certificate{cert}, NextProtos: protocols, CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256, tls.CurveP384, tls.CurveP521}})
			done := make(chan error, 1)
			go func() {
				p := make([]byte, 4)
				_, e := io.ReadFull(server, p)
				if e == nil {
					_, e = server.Write(p)
				}
				done <- e
			}()
			if _, e := rawClient.Write([]byte("test")); e != nil {
				t.Fatal(e)
			}
			p := make([]byte, 4)
			if _, e := io.ReadFull(rawClient, p); e != nil || string(p) != "test" {
				t.Fatal(p, e)
			}
			if e := <-done; e != nil {
				t.Fatal(e)
			}
			state := server.ConnectionState()
			if state.ServerName != "native.test" {
				t.Fatal(state)
			}
			if mode == "configured" && (state.Version != tls.VersionTLS12 || state.NegotiatedProtocol != "native" || state.CipherSuite != tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256) {
				t.Fatal(state)
			}
		})
	}
}
