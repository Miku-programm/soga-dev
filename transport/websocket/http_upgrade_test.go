package websocket

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	head "recovered.local/soga/transport/httphead"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNativeHTTPUpgradeActualTCP(t *testing.T) {
	for _, tc := range []struct {
		name, method, version, key, upgrade, connection string
		status                                          int
	}{{"valid", "GET", "13", "dGhlIHNhbXBsZSBub25jZQ==", "websocket", "Upgrade", 101}, {"case_and_tokens", "GET", "13", "dGhlIHNhbXBsZSBub25jZQ==", "WebSoCkEt", "keep-alive, UpGraDe", 101}, {"only_length_checked", "GET", "13", strings.Repeat("!", 24), "websocket", "Upgrade", 101}, {"method", "POST", "13", "dGhlIHNhbXBsZSBub25jZQ==", "websocket", "Upgrade", 405}, {"version_missing", "GET", "", "dGhlIHNhbXBsZSBub25jZQ==", "websocket", "Upgrade", 400}, {"version_other", "GET", "12", "dGhlIHNhbXBsZSBub25jZQ==", "websocket", "Upgrade", 426}, {"key_length", "GET", "13", "short", "websocket", "Upgrade", 400}, {"upgrade_wrong", "GET", "13", "dGhlIHNhbXBsZSBub25jZQ==", "notws", "Upgrade", 400}, {"connection_wrong", "GET", "13", "dGhlIHNhbXBsZSBub25jZQ==", "websocket", "keep-alive", 400}} {
		t.Run(tc.name, func(t *testing.T) {
			done := make(chan error, 1)
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c, rw, _, e := (OKnCFm3eQ{Timeout: time.Second}).Upgrade(r, w)
				if c != nil {
					defer c.Close()
				}
				if e == nil {
					p := make([]byte, 3)
					_, e = io.ReadFull(rw, p)
					if e == nil && !bytes.Equal(p, []byte{0x82, 1, 'x'}) {
						e = errors.New("buffered bytes changed")
					}
				}
				done <- e
			}))
			defer srv.Close()
			c, e := net.Dial("tcp", strings.TrimPrefix(srv.URL, "http://"))
			if e != nil {
				t.Fatal(e)
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(3 * time.Second))
			request := tc.method + " /ws HTTP/1.1\r\nHost: example.test\r\nUpgrade: " + tc.upgrade + "\r\nConnection: " + tc.connection + "\r\nSec-WebSocket-Key: " + tc.key + "\r\nSec-WebSocket-Version: " + tc.version + "\r\n\r\n"
			if _, e = c.Write(append([]byte(request), 0x82, 1, 'x')); e != nil {
				t.Fatal(e)
			}
			res, e := http.ReadResponse(bufio.NewReader(c), nil)
			if e != nil {
				t.Fatal(e)
			}
			defer res.Body.Close()
			if res.StatusCode != tc.status {
				t.Fatal(res.Status)
			}
			if tc.name == "valid" && res.Header.Get("Sec-Websocket-Accept") != "s3pPLMBiTxaQ9kYGzzhZRbK+xOo=" {
				t.Fatal(res.Header)
			}
			if tc.status == 426 && res.Header.Get("Sec-Websocket-Version") != "13" {
				t.Fatal(res.Header)
			}
			if e = <-done; (e == nil) != (tc.status == 101) {
				t.Fatal(e, tc.status)
			}
		})
	}
}
func TestNativeHTTPUpgradeNegotiationAndCustomHeaders(t *testing.T) {
	done := make(chan error, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _, hs, e := (OKnCFm3eQ{Header: http.Header{"X-Native": {"yes"}}, Protocol: func(s string) bool { return s == "native" }, Negotiate: func(o head.BIMZTmbYHb) (head.BIMZTmbYHb, error) { return o.Copy(make([]byte, o.Size())), nil }}).Upgrade(r, w)
		if c != nil {
			defer c.Close()
		}
		if e == nil && (hs.Protocol != "native" || len(hs.Extensions) != 1) {
			e = errors.New("negotiation mismatch")
		}
		done <- e
	}))
	defer srv.Close()
	c, e := net.Dial("tcp", strings.TrimPrefix(srv.URL, "http://"))
	if e != nil {
		t.Fatal(e)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(3 * time.Second))
	io.WriteString(c, "GET / HTTP/1.1\r\nHost: x\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==\r\nSec-WebSocket-Version: 13\r\nSec-WebSocket-Protocol: ignored, native\r\nSec-WebSocket-Extensions: sample; foo=\"hello world\"\r\n\r\n")
	res, e := http.ReadResponse(bufio.NewReader(c), nil)
	if e != nil {
		t.Fatal(e)
	}
	defer res.Body.Close()
	if res.StatusCode != 101 || res.Header.Get("Sec-Websocket-Protocol") != "native" || res.Header.Get("X-Native") != "yes" || res.Header.Get("Sec-Websocket-Extensions") != "sample;foo=\"hello world\"" {
		t.Fatal(res.Header, res.Status)
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
}
func TestNativeNonHijackerAndTypedError(t *testing.T) {
	r := httptest.NewRequest("GET", "http://x/", nil)
	w := httptest.NewRecorder()
	c, rw, _, e := (OKnCFm3eQ{}).Upgrade(r, w)
	if c != nil || rw != nil || e != ErrNotHijacker || w.Code != 500 {
		t.Fatal(c, rw, e, w.Code)
	}
	if w.Body.String() != "given http.ResponseWriter is not a http.Hijacker" || w.Header().Get("Content-Length") != "48" {
		t.Fatal(w.Body.String(), w.Header())
	}
	e = RejectConnectionError(RejectionStatus(418), RejectionReason("reason"))
	var rej *Zbu9jHc
	if !errors.As(e, &rej) || rej.StatusCode() != 418 || rej.Error() != "reason" {
		t.Fatal(e)
	}
}
func TestNativeNonceExactAndLayout(t *testing.T) {
	p := make([]byte, 28)
	vhaUn9iUya1L(p, []byte("dGhlIHNhbXBsZSBub25jZQ=="))
	if string(p) != "s3pPLMBiTxaQ9kYGzzhZRbK+xOo=" {
		t.Fatal(string(p))
	}
	for _, tc := range []struct {
		a, n int
		want string
	}{{27, 24, "accept buffer is invalid"}, {28, 23, "nonce is invalid"}} {
		func() {
			defer func() {
				if e := recover(); e != tc.want {
					t.Fatal(e)
				}
			}()
			vhaUn9iUya1L(make([]byte, tc.a), make([]byte, tc.n))
		}()
	}
	if reflect.TypeOf(OKnCFm3eQ{}).Size() != 40 || reflect.TypeOf(Handshake{}).Size() != 40 {
		t.Fatal("original 40 byte value layout")
	}
}
