package encmux

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	ws "recovered.local/soga/transport/websocket"
	util "recovered.local/soga/transport/wsutil"
	"reflect"
	"testing"
	"time"
)

type wsClientFixture struct {
	net.Conn
	bytes.Buffer
	closeCount int
	closeErr   error
}

func (c *wsClientFixture) Read(p []byte) (int, error)  { return c.Buffer.Read(p) }
func (c *wsClientFixture) Close() error                { c.closeCount++; return c.closeErr }
func (c *wsClientFixture) Write(p []byte) (int, error) { return c.Buffer.Write(p) }

type clientCauseReader struct{}

func (clientCauseReader) Read(p []byte) (int, error) {
	copy(p, "x")
	return 1, pkgerrors.Wrap(io.EOF, "original")
}
func TestOriginalClientFrameMaskCloseAndCause(t *testing.T) {
	raw := &wsClientFixture{closeErr: errors.New("close error")}
	c := &F2Eb1b{DrpXao: raw}
	typ := reflect.TypeOf(*c)
	if typ.Size() != 56 {
		t.Fatal(typ.Size())
	}
	for i, off := range []uintptr{0, 16, 32, 40} {
		if typ.Field(i).Offset != off {
			t.Fatal(typ.Field(i))
		}
	}
	if c.GetRawConn() != raw {
		t.Fatal("raw")
	}
	payload := bytes.Repeat([]byte("payload"), 1000)
	if n, e := c.Write(payload); e != nil || n != len(payload) {
		t.Fatal(n, e)
	}
	r := &util.HLKO3Cdqw{GX04GFKh1Gz: bytes.NewReader(raw.Bytes()), WXUh5psqY7: ws.StateServerSide}
	h, e := r.NextFrame()
	if e != nil || h.IsHe8C != ws.OpBinary {
		t.Fatal(h, e)
	}
	p, e := io.ReadAll(r)
	if e != nil || !bytes.Equal(p, payload) {
		t.Fatal(e)
	}
	raw.Reset()
	if c.Close() != raw.closeErr || c.Close() != nil || raw.closeCount != 1 {
		t.Fatal("close-once behavior")
	}
	if wire := raw.Bytes(); len(wire) != 6 || wire[0] != 0x88 || wire[1] != 0x80 {
		t.Fatalf("close frame %x", wire)
	}
	c.hESODkc8p = clientCauseReader{}
	p = make([]byte, 1)
	if n, e := c.Read(p); n != 1 || e != nil || c.hESODkc8p != nil {
		t.Fatal(n, e)
	}
}

func TestOriginalWrapClientConnActualTransports(t *testing.T) {
	fixture := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	cert := fixture.TLS.Certificates[0]
	fixture.Close()
	for _, useTLS := range []bool{false, true} {
		for _, useWS := range []bool{false, true} {
			name := "plain"
			if useTLS {
				name += "-tls"
			}
			if useWS {
				name += "-ws"
			}
			t.Run(name, func(t *testing.T) {
				cfg := testTunnelConfig(t)
				cfg.rj1DhSBHZ7O = useTLS
				cfg.ufiMqR6x = useWS
				cfg._rKF9rrYaf8a = &cert
				cfg.psdXJ2Mi = true
				cfg.p1Wtyln2GJ6 = "ignored-for-websocket.test"
				cfg.lyXfKPaz4m23 = "/native"
				raw, e := net.Listen("tcp", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer raw.Close()
				l, e := HNbZee047(raw, cfg)
				if e != nil {
					t.Fatal(e)
				}
				defer l.Close()
				data := bytes.Repeat([]byte("original tunnel client"), 2000)
				done := make(chan error, 1)
				go func() {
					c, e := l.Accept()
					if e != nil {
						done <- e
						return
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(4 * time.Second))
					p := make([]byte, len(data))
					_, e = io.ReadFull(c, p)
					if e == nil && !bytes.Equal(p, data) {
						e = io.ErrNoProgress
					}
					if e == nil {
						_, e = c.Write(p)
					}
					done <- e
				}()
				conn, e := net.Dial("tcp", raw.Addr().String())
				if e != nil {
					t.Fatal(e)
				}
				defer conn.Close()
				conn.SetDeadline(time.Now().Add(4 * time.Second))
				c, e := cfg.WrapClientConn(context.Background(), conn)
				if e != nil {
					t.Fatal(e)
				}
				defer c.Close()
				encrypted, ok := c.(*Cb7vj3KQA8)
				if !ok || encrypted.npFKkBd4j != nil {
					t.Fatal("client encryption wrapper")
				}
				defer encrypted.Free()
				if useWS {
					if _, ok = encrypted.DrpXao.(*F2Eb1b); !ok {
						t.Fatal("WS before encryption")
					}
				}
				// WrapClientConn only applies transport/encryption. The higher tunnel
				// protocol caller must still supply the original 16-byte token and command.
				request := append(append(append([]byte{}, cfg.x5Lor8hDyn[:]...), byte(1)), data...)
				if _, e = c.Write(request); e != nil {
					t.Fatal(e)
				}
				p := make([]byte, len(data))
				if _, e = io.ReadFull(c, p); e != nil || !bytes.Equal(p, data) {
					t.Fatal(e)
				}
				if e = <-done; e != nil {
					t.Fatal(e)
				}
			})
		}
	}
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	cfg := &VTl94pzz2IkB{hZORuZstvS07: true, rj1DhSBHZ7O: true, ufiMqR6x: true}
	c, e := cfg.WrapClientConn(context.Background(), a)
	if c != a || e != nil {
		t.Fatal("raw QUIC bypass", c, e)
	}
}

type clientFrameDuplex struct {
	net.Conn
	input  *bytes.Reader
	output bytes.Buffer
}

func (c *clientFrameDuplex) Read(p []byte) (int, error)  { return c.input.Read(p) }
func (c *clientFrameDuplex) Write(p []byte) (int, error) { return c.output.Write(p) }
func TestOriginalClientControlAndFragmentedRead(t *testing.T) {
	var input bytes.Buffer
	for _, op := range []ws.Ao3txMm8{ws.OpPing, ws.OpBinary, ws.OpBinary} {
		if e := util.WriteFrameRecovered(&input, ws.StateServerSide, op, true, []byte("abc")); e != nil {
			t.Fatal(e)
		}
	}
	raw := &clientFrameDuplex{input: bytes.NewReader(input.Bytes())}
	c := &F2Eb1b{DrpXao: raw, jsKotLrECn: util.L9DiQX(raw, ws.StateClientSide)}
	p := make([]byte, 6)
	if _, e := io.ReadFull(c, p); e != nil || string(p) != "abcabc" {
		t.Fatal(string(p), e)
	}
	reader := &util.HLKO3Cdqw{GX04GFKh1Gz: bytes.NewReader(raw.output.Bytes()), WXUh5psqY7: ws.StateServerSide}
	h, e := reader.NextFrame()
	if e != nil || h.IsHe8C != ws.OpPong {
		t.Fatal(h, e)
	}
	p, e = io.ReadAll(reader)
	if e != nil || string(p) != "abc" {
		t.Fatal(string(p), e)
	}
}

// This fixture returns the complete response and one frame from a single Read,
// making the original discarded-buffer behavior deterministic.
type clientHandshakeFixture struct {
	wireFixture
	request    bytes.Buffer
	responded  bool
	host, path string
}

func (c *clientHandshakeFixture) RemoteAddr() net.Addr {
	return &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 9000}
}
func (c *clientHandshakeFixture) Write(p []byte) (int, error) { return c.request.Write(p) }
func (c *clientHandshakeFixture) Read(p []byte) (int, error) {
	if c.responded {
		return 0, io.EOF
	}
	c.responded = true
	req, e := http.ReadRequest(bufio.NewReader(bytes.NewReader(c.request.Bytes())))
	if e != nil {
		return 0, e
	}
	c.host = req.Host
	c.path = req.URL.RequestURI()
	digest := sha1.Sum([]byte(req.Header.Get("Sec-WebSocket-Key") + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	response := fmt.Sprintf("HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n", base64.StdEncoding.EncodeToString(digest[:]))
	wire := append([]byte(response), 0x82, 3, 'a', 'b', 'c')
	return copy(p, wire), nil
}
func TestOriginalClientHostUnusedAndBufferedFrameDiscarded(t *testing.T) {
	raw := &clientHandshakeFixture{}
	c, e := SYSb8haLOir(context.Background(), raw, "unused.host", "/native?query=1")
	if e != nil {
		t.Fatal(e)
	}
	if raw.host != "127.0.0.1:9000" || raw.path != "/native?query=1" {
		t.Fatal(raw.host, raw.path)
	}
	p := make([]byte, 3)
	if n, e := c.Read(p); n != 0 || e != io.EOF {
		t.Fatal("original loses Dial buffered bytes", n, e, p)
	}
}
