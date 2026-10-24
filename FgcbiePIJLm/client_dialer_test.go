package FgcbiePIJLm

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	head "recovered.local/soga/EsS8DL8PE9"
	pool "recovered.local/soga/gPJ6UJF"
	"reflect"
	"strings"
	"testing"
	"time"
)

type clientMemoryRW struct {
	io.Reader
	bytes.Buffer
}

func (m *clientMemoryRW) Read(p []byte) (int, error) { return m.Reader.Read(p) }
func TestOriginalClientStatusHeadersAndExtensions(t *testing.T) {
	for _, tc := range []struct {
		name, status, headers string
		want                  error
	}{
		{"status", "HTTP/1.1 403 Forbidden", "", XChC4Wxtj(403)},
		{"version", "HTTP/2.0 101 OK", "", ErrHandshakeBadProtocol},
		{"malformed", "HTTP/1.1 garbage OK", "", ErrMalformedResponse},
		{"missing", "HTTP/1.1 101 OK", "", ErrHandshakeBadUpgrade},
		{"connection-token-list", "HTTP/1.1 101 OK", "Upgrade: websocket\r\nConnection: keep-alive, Upgrade\r\n", ErrHandshakeBadConnection},
		{"bad-accept", "HTTP/1.1 101 OK", "Upgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: invalid\r\n", ErrHandshakeBadSecAccept},
		{"unrequested-protocol", "HTTP/1.1 101 OK", "Sec-WebSocket-Protocol: unrequested\r\n", ErrHandshakeBadSubProtocol},
		{"unrequested-extension", "HTTP/1.1 101 OK", "Sec-WebSocket-Extensions: unknown\r\n", ErrHandshakeBadExtensions},
		{"bad-header", "HTTP/1.1 101 OK", "invalid\r\n", ErrMalformedResponse},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wire := &clientMemoryRW{Reader: strings.NewReader(tc.status + "\r\n" + tc.headers + "\r\n")}
			u, _ := url.Parse("ws://native.test/path?q=yes")
			var status int
			var response string
			d := Z4oWaKU5v{OnStatusError: func(s int, _ []byte, r io.Reader) { status = s; b, _ := io.ReadAll(r); response = string(b) }}
			br, _, err := d.Upgrade(wire, u)
			if br != nil || err != tc.want {
				t.Fatalf("%v %v want %v", br, err, tc.want)
			}
			if status != 0 && (status != 403 || !strings.HasPrefix(response, tc.status+"\r\n")) {
				t.Fatal(status, response)
			}
			request, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(wire.Bytes())))
			if err != nil || request.URL.RequestURI() != "/path?q=yes" || request.Header.Get("Sec-WebSocket-Version") != "13" {
				t.Fatal(request, err)
			}
		})
	}
	wanted := []head.BIMZTmbYHb{{KXa0Kt2I: []byte("example")}}
	selected := []byte("example; key=val")
	result, err := n3LyWV(selected, wanted, nil)
	if err != nil || len(result) != 1 || result[0].Jd2j6aD7.Size() != 6 {
		t.Fatal(result, err)
	}
	selected[len(selected)-1] = 'z'
	var b bytes.Buffer
	head.KBgYhFImpuOF(&b, result)
	if b.String() != "example;key=val" {
		t.Fatal(b.String())
	}
	if _, err = n3LyWV([]byte("other"), wanted, nil); err != ErrHandshakeBadExtensions {
		t.Fatal(err)
	}
	if XChC4Wxtj(503).Error() != "unexpected HTTP response status: 503" || ErrHandshakeBadExtensions.Error() != `unexpected extensions in "Sec-WebSocket-Protocol" header` {
		t.Fatal("decoded messages changed")
	}
}

func TestOriginalClientDialActualTCPBufferedFrameAndCallbacks(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	done := make(chan error, 1)
	go func() {
		c, e := l.Accept()
		if e != nil {
			done <- e
			return
		}
		defer c.Close()
		c.SetDeadline(time.Now().Add(3 * time.Second))
		r, e := http.ReadRequest(bufio.NewReader(c))
		if e != nil {
			done <- e
			return
		}
		if r.Host != "chosen.host" || r.Header.Get("X-Client") != "native" || r.URL.RequestURI() != "/stream?a=b" {
			done <- fmt.Errorf("request %v", r)
			return
		}
		accept := make([]byte, 28)
		vhaUn9iUya1L(accept, []byte(r.Header.Get("Sec-WebSocket-Key")))
		_, e = fmt.Fprintf(c, "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\nSec-WebSocket-Protocol: native\r\nX-Server: present\r\n\r\nFRAME", accept)
		done <- e
	}()
	var gotHeader bool
	var wrapCount int
	d := Z4oWaKU5v{Timeout: time.Second, Host: "chosen.host", Protocols: []string{"native"}, Header: NCYj6EtaAgQ(http.Header{"X-Client": []string{"native"}}), OnHeader: func(k, v []byte) error { gotHeader = string(k) == "X-Server" && string(v) == "present"; return nil }, WrapConn: func(c net.Conn) net.Conn { wrapCount++; return c }}
	c, br, hs, err := d.Dial(context.Background(), "ws://"+l.Addr().String()+"/stream?a=b")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	if !gotHeader || wrapCount != 1 || hs.Protocol != "native" || br == nil {
		t.Fatal(gotHeader, wrapCount, hs, br)
	}
	defer pool.PutReader(br)
	b := make([]byte, 5)
	if _, err = io.ReadFull(br, b); err != nil || string(b) != "FRAME" {
		t.Fatal(string(b), err)
	}
	if err = <-done; err != nil {
		t.Fatal(err)
	}
}

func TestOriginalClientContextAndDialErrorBranches(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	ctx, cancel := context.WithCancel(context.Background())
	done := cVc8h7D3zW(ctx, a)
	cancel()
	// A pending I/O observes the real deadline set by the interruption goroutine.
	_, e := a.Read(make([]byte, 1))
	done(&e)
	if e != context.Canceled {
		t.Fatal(e)
	}
	target := errors.New("native failure")
	var wrapped bool
	d := Z4oWaKU5v{NetDial: func(context.Context, string, string) (net.Conn, error) { return a, target }, WrapConn: func(c net.Conn) net.Conn { wrapped = true; return c }}
	u, _ := url.Parse("ws://host/")
	c, e := d.xYiGvj(context.Background(), u)
	if c != a || e != target || !wrapped {
		t.Fatal(c, e, wrapped)
	}
	wrapped = false
	u.Scheme = "wss"
	c, e = d.xYiGvj(context.Background(), u)
	if c != nil || e != target || wrapped {
		t.Fatal(c, e, wrapped)
	}
	u.Scheme = "http"
	if _, e = d.xYiGvj(context.Background(), u); e == nil || e.Error() != `unexpected websocket scheme: "http"` {
		t.Fatal(e)
	}
	for _, tc := range []struct{ in, host, addr string }{{"host", "host", "host:80"}, {"host:9", "host", "host:9"}, {"[::1]", "[::1]", "[::1]:80"}, {"[::1]:9", "[::1]", "[::1]:9"}} {
		h, a := clTie97(tc.in, ":80")
		if h != tc.host || a != tc.addr {
			t.Fatal(tc, h, a)
		}
	}
	if reflect.TypeOf(Z4oWaKU5v{}).Size() != 152 {
		t.Fatal("original Dialer ABI")
	}
}

func TestOriginalClientWSSAndTLSConfigClone(t *testing.T) {
	fixture := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	cert := fixture.TLS.Certificates[0]
	fixture.Close()
	l, e := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: []tls.Certificate{cert}})
	if e != nil {
		t.Fatal(e)
	}
	defer l.Close()
	done := make(chan error, 1)
	go func() {
		c, e := l.Accept()
		if e != nil {
			done <- e
			return
		}
		defer c.Close()
		c.SetDeadline(time.Now().Add(3 * time.Second))
		r, e := http.ReadRequest(bufio.NewReader(c))
		if e != nil {
			done <- e
			return
		}
		a := make([]byte, 28)
		vhaUn9iUya1L(a, []byte(r.Header.Get("Sec-WebSocket-Key")))
		_, e = fmt.Fprintf(c, "HTTP/1.1 101 OK\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n", a)
		done <- e
	}()
	config := &tls.Config{InsecureSkipVerify: true, MaxVersion: tls.VersionTLS12}
	d := Z4oWaKU5v{TLSConfig: config, Timeout: time.Second}
	c, br, _, e := d.Dial(context.Background(), "wss://"+l.Addr().String()+"/native")
	if e != nil {
		t.Fatal(e)
	}
	defer c.Close()
	if br != nil {
		pool.PutReader(br)
	}
	if config.ServerName != "" || c.(*tls.Conn).ConnectionState().Version != tls.VersionTLS12 {
		t.Fatal("clone or negotiated version")
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
}
