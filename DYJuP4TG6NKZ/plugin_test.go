package DYJuP4TG6NKZ

import (
	"bytes"
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

type fakeConn struct {
	r      io.Reader
	writes [][]byte
	err    error
	closed bool
}

func (c *fakeConn) Read(p []byte) (int, error) { return c.r.Read(p) }
func (c *fakeConn) Write(p []byte) (int, error) {
	c.writes = append(c.writes, append([]byte(nil), p...))
	return len(p), c.err
}
func (c *fakeConn) Close() error                   { c.closed = true; return nil }
func (*fakeConn) LocalAddr() net.Addr              { return &net.TCPAddr{} }
func (*fakeConn) RemoteAddr() net.Addr             { return &net.TCPAddr{} }
func (*fakeConn) SetDeadline(time.Time) error      { return nil }
func (*fakeConn) SetReadDeadline(time.Time) error  { return nil }
func (*fakeConn) SetWriteDeadline(time.Time) error { return nil }
func TestMuxWireCountsAndClose(t *testing.T) {
	raw := &fakeConn{}
	c := VvqqszrNmTt(raw, [2]byte{0xab, 0xcd}, 443, "127.0.0.1", "udp")
	expected := []byte{0, 12, 0xab, 0xcd, 1, 0, 2, 1, 0xbb, 1, 127, 0, 0, 1}
	if !bytes.Equal(c.w9i2ksWrHUTk, expected) {
		t.Fatalf("%x", c.w9i2ksWrHUTk)
	}
	if n, e := c.Write([]byte("abc")); n != 11 || e != nil {
		t.Fatal(n, e)
	}
	if len(raw.writes) != 2 || !bytes.Equal(raw.writes[1], []byte{0, 4, 0xab, 0xcd, 2, 1, 0, 3, 'a', 'b', 'c'}) {
		t.Fatal(raw.writes)
	}
	if e := c.Close(); e != nil || !raw.closed || !bytes.Equal(raw.writes[2], []byte{0, 4, 0xab, 0xcd, 3, 0}) {
		t.Fatal(e)
	}
	raw = &fakeConn{err: errors.New("write")}
	c = VvqqszrNmTt(raw, [2]byte{}, 0, "host", "tcp")
	if !bytes.Equal(c.w9i2ksWrHUTk, []byte{0, 12, 0, 0, 1, 0, 1, 0, 0, 2, 'h', 'o', 's', 't'}) {
		t.Fatal(c.w9i2ksWrHUTk)
	}
	if n, e := c.Write(nil); n != 0 || e != raw.err || c.w9i2ksWrHUTk == nil {
		t.Fatal(n, e)
	}
	if e := c.Close(); e != raw.err || raw.closed {
		t.Fatal(e)
	}
}
func TestMuxReadControlAndRemainder(t *testing.T) {
	wire := []byte{0, 4, 0, 9, 4, 0, 0, 4, 0, 7, 2, 1, 0, 5, 'a', 'b', 'c', 'd', 'e'}
	c := &ClSQMRRJ{Conn: &fakeConn{r: bytes.NewReader(wire)}}
	p := make([]byte, 2)
	var got []byte
	for i := 0; i < 3; i++ {
		n, e := c.Read(p)
		if e != nil {
			t.Fatal(e)
		}
		got = append(got, p[:n]...)
	}
	if string(got) != "abcde" || c.heaj4Yoh != [2]byte{0, 7} || c.baWziQysg3CB != 0 {
		t.Fatal(string(got), c)
	}
	c = &ClSQMRRJ{Conn: &fakeConn{r: bytes.NewReader([]byte{2, 1})}}
	if _, e := c.Read(p); e == nil || e.Error() != "invalid metalen" {
		t.Fatal(e)
	}
}
func TestActualWebSocketPlugin(t *testing.T) {
	for _, secure := range []bool{false, true} {
		for _, mux := range []bool{false, true} {
			t.Run(strings.Join([]string{map[bool]string{true: "tls", false: "plain"}[secure], map[bool]string{true: "mux", false: "raw"}[mux]}, "/"), func(t *testing.T) {
				serverErr := make(chan error, 1)
				up := websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}
				handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != "/plugin" || r.Header.Get("X-Check") != "yes" {
						serverErr <- errors.New("request")
						return
					}
					c, e := up.Upgrade(w, r, nil)
					if e != nil {
						serverErr <- e
						return
					}
					defer c.Close()
					if mux {
						_, p, e := c.ReadMessage()
						if e != nil || !bytes.Equal(p, []byte{0, 12, 0, 0, 1, 0, 1, 0, 0, 1, 127, 0, 0, 1}) {
							serverErr <- errors.New("mux hello")
							return
						}
					}
					_, p, e := c.ReadMessage()
					if e == nil {
						if mux {
							if len(p) != 13 || !bytes.Equal(p[:8], []byte{0, 4, 0, 0, 2, 1, 0, 5}) {
								e = errors.New("mux frame")
							}
						}
						if e == nil {
							e = c.WriteMessage(websocket.BinaryMessage, p)
						}
					}
					serverErr <- e
				})
				srv := httptest.NewUnstartedServer(handler)
				if secure {
					srv.StartTLS()
				} else {
					srv.Start()
				}
				defer srv.Close()
				addr := srv.Listener.Addr().String()
				host, port, _ := net.SplitHostPort(addr)
				raw, e := net.Dial("tcp", addr)
				if e != nil {
					t.Fatal(e)
				}
				raw.SetDeadline(time.Now().Add(3 * time.Second))
				c, e := DO4qmnxxv4_(raw, &YgrMuLrWXz{PRkTz1l: host, D6eC4LPWJm0: port, G3GjlvRQW: "/plugin", BwX3DgZaf: map[string]string{"X-Check": "yes"}, PwG_Wda: secure, HjmCr8: true, AUavUP: mux})
				if e != nil {
					t.Fatal(e)
				}
				defer raw.Close()
				n, e := c.Write([]byte("hello"))
				want := 5
				if mux {
					want = 13
				}
				if n != want || e != nil {
					t.Fatal(n, e)
				}
				p := make([]byte, 5)
				if _, e = io.ReadFull(c, p); e != nil || string(p) != "hello" {
					t.Fatal(string(p), e)
				}
				if e = <-serverErr; e != nil {
					t.Fatal(e)
				}
			})
		}
	}
}
func TestLayouts(t *testing.T) {
	if unsafe.Sizeof(YgrMuLrWXz{}) != 64 || unsafe.Sizeof(ClSQMRRJ{}) != 80 {
		t.Fatal("RTTI")
	}
}
