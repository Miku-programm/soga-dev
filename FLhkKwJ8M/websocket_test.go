package FLhkKwJ8M

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"github.com/gorilla/websocket"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestWebsocketLayouts(t *testing.T) {
	if unsafe.Sizeof(ELIQWabl31{}) != 96 || unsafe.Sizeof(ki0pcfT{}) != 56 || unsafe.Sizeof(mWynszCaKB{}) != 80 {
		t.Fatal("wrapper RTTI layout")
	}
	if unsafe.Sizeof(websocket.Dialer{}) != 128 || unsafe.Sizeof(websocket.Conn{}) != 336 {
		t.Fatal("dependency RTTI layout")
	}
}

func TestWebsocketActualEchoAndEarlyData(t *testing.T) {
	for _, mode := range []string{"plain", "tls", "query-early", "path-early"} {
		t.Run(mode, func(t *testing.T) {
			done := make(chan error, 1)
			server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var early string
				if mode == "query-early" {
					early = r.Header.Get("Sec-WebSocket-Protocol")
					if r.URL.RawQuery != "x=1" {
						done <- errors.New("ed query not removed")
						http.Error(w, "query", 400)
						return
					}
				}
				if mode == "path-early" {
					early = strings.TrimPrefix(r.URL.Path, "/base/")
				}
				if r.Header.Get("X-Test") != "one" || len(r.Header.Values("X-Test")) != 1 {
					done <- errors.New("header copy")
					http.Error(w, "header", 400)
					return
				}
				upgrader := websocket.Upgrader{}
				c, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					done <- err
					return
				}
				defer c.Close()
				c.SetReadDeadline(time.Now().Add(3 * time.Second))
				if early != "" {
					data, e := base64.RawURLEncoding.DecodeString(early)
					if e != nil {
						done <- e
						return
					}
					if e = c.WriteMessage(websocket.TextMessage, data); e != nil {
						done <- e
						return
					}
				}
				for {
					_, data, err := c.ReadMessage()
					if err != nil {
						if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
							err = nil
						}
						done <- err
						return
					}
					if err = c.WriteMessage(websocket.TextMessage, data); err != nil {
						done <- err
						return
					}
				}
			}))
			if mode == "tls" {
				server.StartTLS()
			} else {
				server.Start()
			}
			defer server.Close()
			host, port, _ := net.SplitHostPort(server.Listener.Addr().String())
			raw, err := net.Dial("tcp", server.Listener.Addr().String())
			if err != nil {
				t.Fatal(err)
			}
			defer raw.Close()
			config := &ELIQWabl31{OhyeDUGTCGH: host, RTIB8Nkai: port, WOT5UbYYto: "/base/", RlZC_oDxj: http.Header{"X-Test": {"one", "two"}}, Xtp2zWRK: mode == "tls", VbB82ra: &tls.Config{InsecureSkipVerify: true}}
			if mode == "query-early" {
				config.WOT5UbYYto = "/base/?x=1&ed=3"
			}
			if mode == "path-early" {
				config.A3jvUyhB = 3
			}
			c, err := E6YB0lYJp(raw, config)
			if err != nil {
				t.Fatal(err)
			}
			if n, err := c.Write([]byte("hello!")); err != nil || n != 6 {
				t.Fatal(n, err)
			}
			c.SetDeadline(time.Now().Add(3 * time.Second))
			b := make([]byte, 6)
			if _, err = io.ReadFull(c, b); err != nil || string(b) != "hello!" {
				t.Fatal(string(b), err)
			}
			if c.RemoteAddr().String() != raw.RemoteAddr().String() {
				t.Fatal("remote address changed")
			}
			if err = c.Close(); err != nil {
				t.Fatal(err)
			}
			if err = <-done; err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestEarlyCloseRetainsRawConnAndNoOpDeadlines(t *testing.T) {
	raw, peer := net.Pipe()
	defer raw.Close()
	defer peer.Close()
	out, err := uTCy4V(raw, &ELIQWabl31{A3jvUyhB: 3})
	if err != nil {
		t.Fatal(err)
	}
	c := out.(*mWynszCaKB)
	if err = c.SetDeadline(time.Now().Add(-time.Second)); err != nil {
		t.Fatal(err)
	}
	// Cancellation has its own error while the closed flag remains false.
	c.dW68su()
	if _, err = c.Read(make([]byte, 1)); err != io.ErrUnexpectedEOF {
		t.Fatal(err)
	}
	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err = c.Read(make([]byte, 1)); err != io.ErrClosedPipe {
		t.Fatal(err)
	}
	if _, err = c.Write(nil); err != io.ErrClosedPipe {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { _, err := peer.Write([]byte{7}); done <- err }()
	raw.SetReadDeadline(time.Now().Add(time.Second))
	b := make([]byte, 1)
	if _, err = io.ReadFull(raw, b); err != nil || b[0] != 7 {
		t.Fatal(b, err)
	}
	if err = <-done; err != nil {
		t.Fatal(err)
	}
}
