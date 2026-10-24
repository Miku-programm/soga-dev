//go:build linux

package ssactual

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	E "recovered.local/soga/net/outbound"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
)

type httpFixture struct {
	*bytes.Reader
	writes bytes.Buffer
}

func (c *httpFixture) Write(p []byte) (int, error) { return c.writes.Write(p) }

func TestHTTPConstructorAndResponseRules(t *testing.T) {
	if unsafe.Sizeof(DlqIvv{}) != 56 || unsafe.Sizeof(HTTPOptions{}) != 136 {
		t.Fatal("RTTI layout")
	}
	h := CJW16StU3I(HTTPOptions{Name: "http", Server: "::1", Port: 443, TLS: true, SNI: "proxy.example", SkipCertVerify: true, Username: "u", Password: "p", Headers: map[string]string{"x-custom": "v", "proxy-connection": "custom"}})
	if h.Addr() != "[::1]:443" || h.savFEH.ServerName != "proxy.example" || !h.savFEH.InsecureSkipVerify || h.SupportUDP() || h.Type() != 6 {
		t.Fatal(h)
	}
	m := &E.D0mLgQF{ZouEpmjw8: net.IPv4(192, 0, 2, 1), JRXMVvg: 80}
	for _, test := range []struct{ status, want string }{
		{"200 Connection established", ""},
		{"407 Auth", "HTTP need auth"},
		{"405 Not Allowed", "CONNECT method not allowed by proxy"},
		{"403 Forbidden", "can not connect remote err code: 403"},
		{"502 Bad Gateway", "502 Bad Gateway"},
	} {
		c := &httpFixture{Reader: bytes.NewReader([]byte("HTTP/1.1 " + test.status + "\r\nContent-Length: 0\r\n\r\n"))}
		err := h.g9TEl8UN(m, c)
		if test.want == "" {
			if err != nil {
				t.Fatal(err)
			}
		} else if err == nil || err.Error() != test.want {
			t.Fatal(err, test.want)
		}
		r, err := http.ReadRequest(bufio.NewReader(&c.writes))
		if err != nil {
			t.Fatal(err)
		}
		if r.Method != "CONNECT" || r.Host != "192.0.2.1:80" || r.Header.Get("Proxy-Authorization") != "Basic dTpw" || r.Header.Get("X-Custom") != "v" {
			t.Fatal(r)
		}
		if got := r.Header.Values("Proxy-Connection"); len(got) != 2 || got[0] != "custom" || got[1] != "Keep-Alive" {
			t.Fatal(got)
		}
		if len(h.DvfjZBhXc.Values("Proxy-Connection")) != 1 || h.DvfjZBhXc.Get("Proxy-Authorization") != "" {
			t.Fatal("mutated base headers")
		}
	}
	// One empty credential suppresses authentication, and response-prefetched
	// bytes are not returned through a buffered wrapper by the original.
	h.bqekXVE = ""
	c := &httpFixture{Reader: bytes.NewReader([]byte("HTTP/1.1 200 OK\r\n\r\nearly-data"))}
	if err := h.g9TEl8UN(m, c); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(c.writes.String(), "Proxy-Authorization:") || c.Len() != 0 {
		t.Fatal(c.writes.String(), c.Len())
	}
}

func TestHTTPConnectAndTLSActualLocalEcho(t *testing.T) {
	for _, secure := range []bool{false, true} {
		t.Run(fmt.Sprint(secure), func(t *testing.T) {
			done := make(chan error, 1)
			server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "CONNECT" || r.Host != "192.0.2.1:80" || r.Header.Get("Proxy-Authorization") != "Basic dTpw" {
					http.Error(w, "bad request", 400)
					done <- fmt.Errorf("request %v", r)
					return
				}
				conn, buffer, err := w.(http.Hijacker).Hijack()
				if err != nil {
					done <- err
					return
				}
				defer conn.Close()
				conn.SetDeadline(time.Now().Add(3 * time.Second))
				_, err = buffer.WriteString("HTTP/1.1 200 Connection established\r\n\r\n")
				if err == nil {
					err = buffer.Flush()
				}
				if err == nil {
					_, err = io.Copy(conn, conn)
				}
				done <- err
			}))
			if secure {
				server.StartTLS()
			} else {
				server.Start()
			}
			defer server.Close()
			host, port, _ := net.SplitHostPort(server.Listener.Addr().String())
			portNumber, _ := strconv.Atoi(port)
			h := CJW16StU3I(HTTPOptions{Name: "http", Server: host, Port: portNumber, TLS: secure, SkipCertVerify: true, Username: "u", Password: "p"})
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			c, err := h.DialContext(ctx, &E.D0mLgQF{ZouEpmjw8: net.IPv4(192, 0, 2, 1), JRXMVvg: 80})
			if err != nil {
				t.Fatal(err)
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(3 * time.Second))
			if _, err = c.Write([]byte("echo")); err != nil {
				t.Fatal(err)
			}
			b := make([]byte, 4)
			if _, err = io.ReadFull(c, b); err != nil || string(b) != "echo" || c.Chains()[0] != "http" {
				t.Fatal(string(b), err, c.Chains())
			}
			c.Close()
			if err = <-done; err != nil {
				t.Fatal(err)
			}
		})
	}
}
