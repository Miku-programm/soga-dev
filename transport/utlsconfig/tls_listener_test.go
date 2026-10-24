package utlsconfig

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTunnelTLSActualALPNAndVersions(t *testing.T) {
	fixture := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	cert := fixture.TLS.Certificates[0]
	fixture.Close()
	for _, h2 := range []bool{false, true} {
		raw, e := net.Listen("tcp", "127.0.0.1:0")
		if e != nil {
			t.Fatal(e)
		}
		l := MYA8Jh(raw, &cert, h2)
		done := make(chan error, 1)
		go func() {
			c, e := l.Accept()
			if e == nil {
				defer c.Close()
				c.SetDeadline(time.Now().Add(5 * time.Second))
				_, e = io.CopyN(c, c, 4)
			}
			done <- e
		}()
		c, e := tls.Dial("tcp", l.Addr().String(), &tls.Config{InsecureSkipVerify: true, NextProtos: []string{"h2", "http/1.1"}, MinVersion: tls.VersionTLS12, MaxVersion: tls.VersionTLS12})
		if e != nil {
			l.Close()
			t.Fatal(e)
		}
		want := "http/1.1"
		if h2 {
			want = "h2"
		}
		if c.ConnectionState().NegotiatedProtocol != want || c.ConnectionState().Version != tls.VersionTLS12 {
			t.Fatal(c.ConnectionState())
		}
		c.Write([]byte("ping"))
		p := make([]byte, 4)
		if _, e = io.ReadFull(c, p); e != nil || string(p) != "ping" {
			t.Fatal(e)
		}
		c.Close()
		l.Close()
		if e = <-done; e != nil {
			t.Fatal(e)
		}
	}
}
