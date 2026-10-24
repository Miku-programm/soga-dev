package AKP5mbedNKj

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"golang.org/x/net/http2"
)

func localTLS(t *testing.T) (*tls.Config, *tls.Config) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	cert := &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "soga-loopback"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, IPAddresses: []net.IP{net.ParseIP("127.0.0.1")}}
	der, err := x509.CreateCertificate(rand.Reader, cert, cert, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	private, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	pair, err := tls.X509KeyPair(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}), pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: private}))
	if err != nil {
		t.Fatal(err)
	}
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
	return &tls.Config{Certificates: []tls.Certificate{pair}, NextProtos: []string{"discard-me"}}, &tls.Config{RootCAs: roots}
}

func TestH2ListenerRealConcurrentStreams(t *testing.T) {
	for _, encrypted := range []bool{false, true} {
		t.Run(fmt.Sprintf("tls=%t", encrypted), func(t *testing.T) {
			raw, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				t.Fatal(err)
			}
			var serverTLS, clientTLS *tls.Config
			if encrypted {
				serverTLS, clientTLS = localTLS(t)
			}
			listener := Qi8msD8m74(raw, "/configured-but-not-filtered", serverTLS)
			defer listener.Close()
			internal := listener.(*SNvD5V)
			if cap(internal.vTk17Wmts) != 16384 || internal.f4RBDAi.ReadHeaderTimeout != 4*time.Second {
				t.Fatal("original constructor settings changed")
			}
			if encrypted && (len(serverTLS.NextProtos) != 1 || serverTLS.NextProtos[0] != "h2" || internal.f4RBDAi.TLSConfig != serverTLS) {
				t.Fatal("TLS pointer/ALPN mutation changed")
			}
			tr := &http2.Transport{AllowHTTP: !encrypted, TLSClientConfig: clientTLS}
			if !encrypted {
				tr.DialTLSContext = func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
					return (&net.Dialer{}).DialContext(ctx, network, addr)
				}
			}
			defer tr.CloseIdleConnections()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			methods := []string{"POST", "GET", "CUSTOM"}
			accepted := make(chan net.Conn, len(methods))
			serveErrors := make(chan error, len(methods))
			go func() {
				for range methods {
					c, e := listener.Accept()
					if e != nil {
						serveErrors <- e
						return
					}
					accepted <- c
					go func(c net.Conn) {
						defer c.Close()
						payload, e := io.ReadAll(c)
						if e == nil {
							_, e = c.Write(append([]byte("echo:"), payload...))
						}
						serveErrors <- e
					}(c)
				}
			}()
			var wg sync.WaitGroup
			for i, method := range methods {
				wg.Add(1)
				go func(i int, method string) {
					defer wg.Done()
					scheme := "http"
					if encrypted {
						scheme = "https"
					}
					payload := fmt.Sprintf("stream-%d", i)
					req, e := http.NewRequestWithContext(ctx, method, scheme+"://"+raw.Addr().String()+"/different-path", strings.NewReader(payload))
					if e != nil {
						t.Error(e)
						return
					}
					response, e := tr.RoundTrip(req)
					if e != nil {
						t.Error(e)
						return
					}
					defer response.Body.Close()
					if response.ProtoMajor != 2 || response.StatusCode != 200 || response.Header.Get("Cache-Control") != "no-store" {
						t.Errorf("unexpected response: %s %d %v", response.Proto, response.StatusCode, response.Header)
					}
					if encrypted && (response.TLS == nil || response.TLS.NegotiatedProtocol != "h2") {
						t.Error("TLS did not negotiate h2")
					}
					got, e := io.ReadAll(response.Body)
					if e != nil || !bytes.Equal(got, []byte("echo:"+payload)) {
						t.Errorf("body %q: %v", got, e)
					}
				}(i, method)
			}
			wg.Wait()
			for range methods {
				select {
				case e := <-serveErrors:
					if e != nil {
						t.Error(e)
					}
				case <-ctx.Done():
					t.Fatal("server stream did not finish")
				}
			}
			for range methods {
				c := <-accepted
				if c.LocalAddr().String() != raw.Addr().String() {
					t.Error("local address changed")
				}
				if _, e := c.Read(make([]byte, 1)); e != io.ErrClosedPipe {
					t.Errorf("closed read = %v", e)
				}
				if _, e := c.Write([]byte("x")); e != io.ErrClosedPipe {
					t.Errorf("closed write = %v", e)
				}
			}
		})
	}
}

func TestListenerCloseReleasesAccept(t *testing.T) {
	raw, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	l := Qi8msD8m74(raw, "/", nil)
	done := make(chan error, 1)
	go func() { _, err := l.Accept(); done <- err }()
	l.Close()
	select {
	case err := <-done:
		if err == nil || err.Error() != "h2 listener closed" {
			t.Fatalf("Accept error = %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Accept blocked after Close")
	}
}

// Exercises the separate HTTP/1.1 Upgrade route, not prior-knowledge h2c.
func TestH2CUpgradeRealFrames(t *testing.T) {
	raw, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	l := Qi8msD8m74(raw, "/unused", nil)
	defer l.Close()
	serverDone := make(chan error, 1)
	go func() {
		c, e := l.Accept()
		if e != nil {
			serverDone <- e
			return
		}
		defer c.Close()
		payload, e := io.ReadAll(c)
		if e == nil && string(payload) != "up" {
			e = fmt.Errorf("upgrade body %q", payload)
		}
		if e == nil {
			_, e = c.Write([]byte("upgrade-echo"))
		}
		serverDone <- e
	}()
	c, err := net.Dial("tcp", raw.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(3 * time.Second))
	io.WriteString(c, "POST /arbitrary HTTP/1.1\r\nHost: example.test\r\nConnection: Upgrade, HTTP2-Settings\r\nUpgrade: h2c\r\nHTTP2-Settings: \r\nContent-Length: 2\r\n\r\nup")
	reader := bufio.NewReader(c)
	status, err := reader.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(status, "101 Switching Protocols") {
		t.Fatalf("upgrade status %q", status)
	}
	for {
		line, e := reader.ReadString('\n')
		if e != nil {
			t.Fatal(e)
		}
		if line == "\r\n" {
			break
		}
	}
	io.WriteString(c, http2.ClientPreface)
	framer := http2.NewFramer(c, reader)
	if err := framer.WriteSettings(); err != nil {
		t.Fatal(err)
	}
	var body []byte
	done := false
	for !done {
		frame, e := framer.ReadFrame()
		if e != nil {
			t.Fatal(e)
		}
		switch frame := frame.(type) {
		case *http2.SettingsFrame:
			if !frame.IsAck() {
				if e := framer.WriteSettingsAck(); e != nil {
					t.Fatal(e)
				}
			}
		case *http2.DataFrame:
			if frame.StreamID == 1 {
				body = append(body, frame.Data()...)
				done = frame.StreamEnded()
			}
		case *http2.HeadersFrame:
			if frame.StreamID == 1 && frame.StreamEnded() {
				done = true
			}
		case *http2.RSTStreamFrame:
			t.Fatalf("stream reset: %v", frame.ErrCode)
		case *http2.GoAwayFrame:
			t.Fatalf("goaway: %v %s", frame.ErrCode, frame.DebugData())
		}
	}
	if string(body) != "upgrade-echo" {
		t.Fatalf("upgrade response %q", body)
	}
	select {
	case e := <-serverDone:
		if e != nil {
			t.Fatal(e)
		}
	case <-time.After(time.Second):
		t.Fatal("upgrade stream did not finish")
	}
}
