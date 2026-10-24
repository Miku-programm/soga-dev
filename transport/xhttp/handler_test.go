package xhttp

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	Q "recovered.local/soga/util/deadline2"
	"recovered.local/soga/xray/done"
	"recovered.local/soga/xray/internet"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func xhttpFixture(t *testing.T, cfg *R0J19auWof) (*KcahyAO, string) {
	t.Helper()
	raw, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	addr := raw.Addr().String()
	l := DWB3oG(raw, cfg).(*KcahyAO)
	t.Cleanup(func() { l.uSJi0d.Close(); l.Addr() })
	return l, "http://" + addr
}
func paddedRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	t.Helper()
	r, e := http.NewRequest(method, url, body)
	if e != nil {
		t.Fatal(e)
	}
	r.Header.Set("Referer", "http://test/?x_padding="+strings.Repeat("X", 100))
	return r
}
func xhttpAccept(t *testing.T, l *KcahyAO) net.Conn {
	t.Helper()
	ready := make(chan net.Conn, 1)
	go func() { c, _ := l.Accept(); ready <- c }()
	select {
	case c := <-ready:
		if c == nil {
			t.Fatal("closed queue")
		}
		return c
	case <-time.After(3 * time.Second):
		t.Fatal("Accept timeout")
		return nil
	}
}
func TestNativeXHTTPPacketUploadRealHTTP1(t *testing.T) {
	for _, place := range []string{"body", "header", "cookie", "auto"} {
		t.Run(place, func(t *testing.T) {
			l, base := xhttpFixture(t, &R0J19auWof{HoBz4V: "/t", MfWFUwd1: "packet-up", Aonc19Hp33E: place, HFEkUe4a9Sas: "P"})
			client := &http.Client{Timeout: 3 * time.Second}
			var expected []byte
			payloads := [][]byte{bytes.Repeat([]byte("first\x00"), 900), bytes.Repeat([]byte("second\xff"), 700)}
			for _, p := range payloads {
				expected = append(expected, p...)
			}
			// 上传乱序 1、0，下载仍必须输出 0、1。
			for _, seq := range []int{1, 0} {
				data := payloads[seq]
				var body []byte
				if place == "body" || place == "auto" {
					body = data
				}
				r := paddedRequest(t, "POST", fmt.Sprintf("%s/t/id/%d", base, seq), bytes.NewReader(body))
				if place == "header" {
					r.Header.Set("P-0", base64.RawURLEncoding.EncodeToString(data))
				}
				if place == "cookie" {
					r.AddCookie(&http.Cookie{Name: "P_0", Value: base64.RawURLEncoding.EncodeToString(data)})
				}
				if place == "auto" { // 原顺序 header、cookie、body。
					r.Header.Set("P-0", base64.RawURLEncoding.EncodeToString([]byte("H")))
					r.AddCookie(&http.Cookie{Name: "P_0", Value: base64.RawURLEncoding.EncodeToString([]byte("C"))})
				}
				response, e := client.Do(r)
				if e != nil {
					t.Fatal(e)
				}
				io.Copy(io.Discard, response.Body)
				response.Body.Close()
				if response.StatusCode != 200 {
					t.Fatalf("upload %s", response.Status)
				}
			}
			if place == "auto" {
				expected = append([]byte("HC"), payloads[0]...)
				expected = append(expected, []byte("HC")...)
				expected = append(expected, payloads[1]...)
			}
			r := paddedRequest(t, "GET", base+"/t/id", nil)
			response, e := client.Do(r)
			if e != nil {
				t.Fatal(e)
			}
			defer response.Body.Close()
			if response.Header.Get("X-Accel-Buffering") != "no" || response.Header.Get("Content-Type") != "text/event-stream" {
				t.Fatal(response.Header)
			}
			c := xhttpAccept(t, l)
			defer c.Close()
			if c.LocalAddr().String() != strings.TrimPrefix(base, "http://") {
				t.Fatalf("wrong local addr %s", c.LocalAddr())
			}
			c.SetReadDeadline(time.Now().Add(2 * time.Second))
			got := make([]byte, len(expected))
			if _, e = io.ReadFull(c, got); e != nil || !bytes.Equal(got, expected) {
				t.Fatalf("upload payload mismatch %v", e)
			}
			reply := bytes.Repeat([]byte("response"), 8192)
			writeDone := make(chan error, 1)
			go func() { _, e := c.Write(reply); c.Close(); writeDone <- e }()
			result, e := io.ReadAll(response.Body)
			if e != nil || !bytes.Equal(result, reply) {
				t.Fatalf("download mismatch %d %v", len(result), e)
			}
			if e = <-writeDone; e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestNativeXHTTPSingleAndSplitStreamsHTTP2(t *testing.T) {
	for _, split := range []bool{false, true} {
		t.Run(fmt.Sprint(split), func(t *testing.T) {
			cfg := &R0J19auWof{HoBz4V: "/duplex", MfWFUwd1: "auto", Iv950ouQ5AW: &MUVoWl0{AS9zG362: -1, HHTPwX0N: -1}}
			l, base := xhttpFixture(t, cfg)
			tr := &http2.Transport{AllowHTTP: true, DialTLSContext: func(ctx context.Context, n, a string, _ *tls.Config) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(ctx, n, a)
			}}
			defer tr.CloseIdleConnections()
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			defer cancel()
			upload := bytes.Repeat([]byte("native-stream\x00"), 9000)
			path := "/duplex/"
			if split {
				path += "session"
			}
			up := paddedRequest(t, "POST", base+path, bytes.NewReader(upload)).WithContext(ctx)
			uplink := make(chan *http.Response, 1)
			errs := make(chan error, 1)
			go func() {
				r, e := tr.RoundTrip(up)
				if e != nil {
					errs <- e
					return
				}
				uplink <- r
			}()
			var down *http.Response
			if split {
				r := paddedRequest(t, "GET", base+path, nil).WithContext(ctx)
				var e error
				down, e = tr.RoundTrip(r)
				if e != nil {
					t.Fatal(e)
				}
			} else {
				select {
				case down = <-uplink:
				case e := <-errs:
					t.Fatal(e)
				case <-ctx.Done():
					t.Fatal(ctx.Err())
				}
			}
			defer down.Body.Close()
			if down.ProtoMajor != 2 || down.StatusCode != 200 {
				t.Fatalf("%s", down.Status)
			}
			c := xhttpAccept(t, l)
			defer c.Close()
			c.SetReadDeadline(time.Now().Add(3 * time.Second))
			got := make([]byte, len(upload))
			if _, e := io.ReadFull(c, got); e != nil || !bytes.Equal(got, upload) {
				t.Fatalf("stream read: %v", e)
			}
			reply := bytes.Repeat([]byte("down"), 14000)
			ch := make(chan error, 1)
			go func() { _, e := c.Write(reply); c.Close(); ch <- e }()
			got, e := io.ReadAll(down.Body)
			if e != nil || !bytes.Equal(got, reply) {
				t.Fatalf("stream write: %v", e)
			}
			if e := <-ch; e != nil {
				t.Fatal(e)
			}
			if split {
				select {
				case resp := <-uplink:
					resp.Body.Close()
				case e := <-errs:
					t.Fatal(e)
				case <-ctx.Done():
					t.Fatal(ctx.Err())
				}
			}
		})
	}
}

func TestNativeXHTTPRejectsAndAddressRules(t *testing.T) {
	cfg := &R0J19auWof{HoBz4V: "/x", TzkkcVTusJD: "example.test", MfWFUwd1: "packet-up", Aonc19Hp33E: "body", ULnN_iDWLIXN: &MUVoWl0{AS9zG362: 8, HHTPwX0N: 8}}
	l, base := xhttpFixture(t, cfg)
	h := l.uSJi0d.Handler.(*e8Q7HURSMg)
	tests := []struct {
		name, method, path, mode, place, seq, body string
		length                                     int64
		status                                     int
	}{
		{"missing-session", "POST", "/x/", "packet-up", "body", "", "a", 1, 400},
		{"wrong-path", "POST", "/wrong", "auto", "body", "", "a", 1, 404},
		{"options", "OPTIONS", "/x/", "packet-up", "body", "", "", 0, 200},
		{"wrong-mode", "POST", "/x/id/0", "stream-up", "body", "", "a", 1, 400},
		{"invalid-sequence", "POST", "/x/id/-1", "packet-up", "body", "", "a", 1, 500},
		{"too-large", "POST", "/x/id/0", "packet-up", "body", "", "123456789", 9, 413},
		{"truncated", "POST", "/x/id/0", "packet-up", "body", "", "abc", 4, 400},
		{"unknown-length-eof", "POST", "/x/id/0", "packet-up", "body", "", "abc", -1, 400},
		{"unknown-length-overflow", "POST", "/x/id/0", "packet-up", "body", "", "123456789", -1, 413},
		{"empty-body", "POST", "/x/id/0", "packet-up", "body", "", "", 0, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.MfWFUwd1 = tt.mode
			cfg.Aonc19Hp33E = tt.place
			r := paddedRequest(t, tt.method, base+tt.path, strings.NewReader(tt.body))
			r.Host = "EXAMPLE.TEST:80"
			r.ContentLength = tt.length
			w := httptest.NewRecorder()
			h.ServeHTTP(w, r)
			if w.Code != tt.status {
				t.Fatalf("got %d want %d", w.Code, tt.status)
			}
		})
	}
	r := paddedRequest(t, "POST", base+"/x/id/0", strings.NewReader("a"))
	r.Host = "bad.test"
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != 404 || w.Header().Get("X-Padding") != "" {
		t.Fatal("host must reject before padding")
	}
	r.RemoteAddr = "bad-address"
	r.ProtoMajor = 3
	if a := h.requestRemoteAddr(r); a.Network() != "udp" || a.String() != "0.0.0.0:0" {
		t.Fatal(a)
	}
	r.Header.Set("X-Forwarded-For", "example.test, 192.0.2.1")
	if a := h.requestRemoteAddr(r); a.Network() != "udp" {
		t.Fatal("must not skip domain first item")
	}
	r.Header.Set("X-Forwarded-For", " 192.0.2.1 , example.test")
	if a := h.requestRemoteAddr(r); a.Network() != "tcp" || a.String() != "192.0.2.1:0" {
		t.Fatal(a)
	}
	h.epe3GXboMP1 = &internet.SocketConfig{TrustedXForwardedFor: []string{"Trusted"}}
	if a := h.requestRemoteAddr(r); a.Network() != "udp" {
		t.Fatal("trust gate ignored")
	}
	r.Header["Trusted"] = []string{""}
	if a := h.requestRemoteAddr(r); a.Network() != "tcp" {
		t.Fatal("gate checks values count, not value content")
	}
}

func TestNativeXHTTPAddrAndLayouts(t *testing.T) {
	if unsafe.Sizeof(KcahyAO{}) != 352 || unsafe.Sizeof(e8Q7HURSMg{}) != 136 || unsafe.Sizeof(aaktBBMB{}) != 48 || unsafe.Sizeof(_0_9jH{}) != 80 {
		t.Fatal("original RTTI layouts changed")
	}
	l, _ := xhttpFixture(t, nil)
	l.Addr()
	l.Addr()
	if c, e := l.Accept(); c != nil || e == nil || e.Error() != "listener closed" {
		t.Fatalf("Addr close side effect missing %v", e)
	}
}

type closeProbe struct {
	calls *[]string
	name  string
	err   error
}

func (p *closeProbe) Read(b []byte) (int, error)  { return 0, io.EOF }
func (p *closeProbe) Write(b []byte) (int, error) { return len(b), nil }
func (p *closeProbe) Close() error                { *p.calls = append(*p.calls, p.name); return p.err }
func TestOriginalConnectionCloseOrderAndDeadline(t *testing.T) {
	var calls []string
	writerErr := errors.New("writer")
	w := &closeProbe{&calls, "writer", writerErr}
	r := &closeProbe{&calls, "reader", errors.New("reader")}
	c := &_0_9jH{scAGARMF0XsA: w, xxabacttFOn: r, pcYQQ0: Q.NewReader(r), gvvJIQwQyA: func() { calls = append(calls, "callback") }}
	if e := c.Close(); e != writerErr {
		t.Fatal(e)
	}
	if strings.Join(calls, ",") != "callback,writer,reader" {
		t.Fatal(calls)
	}
	c.SetWriteDeadline(time.Now().Add(-time.Second))
	if n, e := c.Write([]byte("x")); n != 1 || e != nil {
		t.Fatal(n, e)
	}
	c.SetDeadline(time.Now().Add(-time.Second))
	if _, e := c.Read(make([]byte, 1)); e == nil {
		t.Fatal("read deadline ignored")
	}
}

type failFlushWriter struct {
	header  http.Header
	err     error
	flushes int
}

func (w *failFlushWriter) Header() http.Header         { return w.header }
func (w *failFlushWriter) WriteHeader(int)             {}
func (w *failFlushWriter) Write(p []byte) (int, error) { return len(p), w.err }
func (w *failFlushWriter) Flush()                      { w.flushes++ }
func TestOriginalResponseWriterState(t *testing.T) {
	w := &failFlushWriter{header: make(http.Header), err: errors.New("write")}
	s := &aaktBBMB{Instance: done.New(), Reader: strings.NewReader("x"), Eu1uXZE: w}
	if _, e := s.Write([]byte("a")); e == nil || w.flushes != 0 {
		t.Fatal("failed writes must not flush")
	}
	w.err = nil
	s.Write([]byte("a"))
	if w.flushes != 1 {
		t.Fatal("successful write not flushed")
	}
	s.Close()
	s.Close()
	if _, e := s.Write([]byte("a")); e != io.ErrClosedPipe {
		t.Fatal(e)
	}
	if p, e := io.ReadAll(s); e != nil || string(p) != "x" {
		t.Fatal("Close incorrectly closed the reader")
	}
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)
		go func() { defer wg.Done(); s.Close() }()
	}
	wg.Wait()
}
