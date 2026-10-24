package httpconn

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"unsafe"

	"go.uber.org/atomic"
	"golang.org/x/net/http2"
)

func framed(p []byte) []byte {
	var h [11]byte
	h[0] = 10
	n := 1 + binary.PutUvarint(h[1:], uint64(len(p)))
	result := make([]byte, 5)
	binary.BigEndian.PutUint32(result[1:], uint32(n+len(p)))
	result = append(result, h[:n]...)
	return append(result, p...)
}
func preparedRead(data []byte) *Nma1KEcWah {
	c := &Nma1KEcWah{gy7Tt5cwR2i_: &http.Response{}, ivecq3: bufio.NewReader(bytes.NewReader(data))}
	c.tYx02bvA.Do(func() {})
	return c
}
func TestLayouts(t *testing.T) {
	var c Nma1KEcWah
	if unsafe.Sizeof(c) != 96 || unsafe.Offsetof(c.cN_qy4ypoit) != 56 || unsafe.Offsetof(c.mWaHrPjpuc) != 88 || unsafe.Sizeof(EgMrkYjAnvwk{}) != 32 {
		t.Fatal("RTTI mismatch")
	}
}
func TestFramesAndRemainder(t *testing.T) {
	wire := framed([]byte("abcdef"))
	// The original ignores compression, outer size, and protobuf field tag.
	copy(wire[:6], []byte{7, 255, 255, 255, 255, 99})
	c := preparedRead(wire)
	p := make([]byte, 2)
	for _, want := range []string{"ab", "cd", "ef"} {
		n, err := c.Read(p)
		if err != nil || n != 2 || string(p) != want {
			t.Fatalf("%d %q %v", n, p, err)
		}
	}
	if c.qaeg__ni6JWF != 0 {
		t.Fatal(c.qaeg__ni6JWF)
	}
	c = preparedRead(append(make([]byte, 6), 0x80))
	if n, err := c.Read(p); n != 0 || err != errorAt5098610 {
		t.Fatalf("varint %d %v", n, err)
	}
	c = preparedRead([]byte{1, 2})
	if n, err := c.Read(p); n != 0 || err != io.EOF {
		t.Fatalf("discard %d %v", n, err)
	}
	c = preparedRead([]byte("x"))
	c.qaeg__ni6JWF = 3
	if n, err := c.Read(make([]byte, 3)); n != 1 || err != io.ErrUnexpectedEOF || c.qaeg__ni6JWF != 2 {
		t.Fatalf("partial %d %v remaining=%d", n, err, c.qaeg__ni6JWF)
	}
	c = preparedRead(nil)
	c.gy7Tt5cwR2i_ = nil
	if _, err := c.Read(p); err != net.ErrClosed {
		t.Fatal(err)
	}
}
func TestWriteFramingAndFailureCount(t *testing.T) {
	for _, p := range [][]byte{nil, []byte("hello"), bytes.Repeat([]byte{42}, 130)} {
		r, w := io.Pipe()
		c := &Nma1KEcWah{h7EQWALPWw: w}
		result := make(chan []byte, 1)
		go func() { b := make([]byte, len(framed(p))); io.ReadFull(r, b); result <- b }()
		n, err := c.Write(p)
		if n != len(p) || err != nil {
			t.Fatalf("%d %v", n, err)
		}
		if b := <-result; !bytes.Equal(b, framed(p)) {
			t.Fatalf("frame %x", b)
		}
		r.Close()
		w.Close()
		c.cN_qy4ypoit = errors.New("upstream failed")
		n, err = c.Write(p)
		if n != len(p) || err != c.cN_qy4ypoit {
			t.Fatalf("error count %d %v", n, err)
		}
	}
}
func TestDeadlineZeroClosesPipe(t *testing.T) {
	r, w := io.Pipe()
	defer r.Close()
	c := &Nma1KEcWah{h7EQWALPWw: w, gFnUW3RHY: atomic.NewBool(false)}
	c.SetReadDeadline(time.Time{})
	done := make(chan error, 1)
	go func() { _, err := r.Read(make([]byte, 1)); done <- err }()
	select {
	case err := <-done:
		if err != io.EOF {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("zero deadline did not close")
	}
	if !c.gFnUW3RHY.Load() {
		t.Fatal("not marked closed")
	}
}
func TestHTTP2Echo(t *testing.T) {
	for _, service := range []string{"", "custom"} {
		t.Run(service, func(t *testing.T) {
			observed := make(chan *http.Request, 1)
			server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				observed <- r
				w.WriteHeader(503) // The original does not reject non-200 statuses.
				w.(http.Flusher).Flush()
				b := make([]byte, 4096)
				for {
					n, err := r.Body.Read(b)
					if n > 0 {
						w.Write(b[:n])
						w.(http.Flusher).Flush()
					}
					if err != nil {
						return
					}
				}
			}))
			server.EnableHTTP2 = true
			server.StartTLS()
			defer server.Close()
			raw, err := net.Dial("tcp", server.Listener.Addr().String())
			if err != nil {
				t.Fatal(err)
			}
			config := &tls.Config{InsecureSkipVerify: true, NextProtos: []string{"h2"}}
			conn, err := H6iElYbPlX4(raw, config, &EgMrkYjAnvwk{JbYoxUV: service, EbklgbD: server.Listener.Addr().String()})
			if err != nil {
				t.Fatal(err)
			}
			c := conn.(*Nma1KEcWah)
			defer c.jE67IuTLu.CloseIdleConnections()
			defer conn.Close()
			conn.SetDeadline(time.Now().Add(3 * time.Second))
			payload := bytes.Repeat([]byte("abc"), 200)
			if n, err := conn.Write(payload); n != len(payload) || err != nil {
				t.Fatalf("write %d %v", n, err)
			}
			reply := make([]byte, len(payload))
			if _, err := io.ReadFull(conn, reply); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(reply, payload) {
				t.Fatal("echo mismatch")
			}
			request := <-observed
			expected := service
			if expected == "" {
				expected = "GunService"
			}
			if request.URL.Path != "/"+expected+"/Tun" || request.Method != "POST" || request.ProtoMajor != 2 || request.Header.Get("User-Agent") != "grpc-go/1.36.0" || request.Header.Get("Content-Type") != "application/grpc" {
				t.Fatalf("request %#v", request)
			}
			if c.fZ6viLKTqLal.URL.Opaque != "//"+request.Host+"/"+expected+"/Tun" {
				t.Fatal(c.fZ6viLKTqLal.URL)
			}
			c.mWaHrPjpuc.Stop()
		})
	}
}
func TestALPNRejectionAndDialFailure(t *testing.T) {
	expected := errors.New("dial failed")
	fn := AljPVof590m(func(string, string) (net.Conn, error) { return nil, expected })
	if _, err := fn(context.Background(), "tcp", "example", nil); err != expected {
		t.Fatal(err)
	}
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer server.Close()
	fn = AljPVof590m(net.Dial)
	_, err := fn(context.Background(), "tcp", server.Listener.Addr().String(), &tls.Config{InsecureSkipVerify: true, NextProtos: []string{"http/1.1"}})
	if err == nil || !strings.Contains(err.Error(), "unexpected ALPN protocol http/1.1, want h2") {
		t.Fatal(err)
	}
}
func TestRoundTripErrorPropagates(t *testing.T) {
	expected := errors.New("specific dial failure")
	transport := &http2.Transport{DialTLSContext: func(context.Context, string, string, *tls.Config) (net.Conn, error) { return nil, expected }}
	conn, err := TwFKkx(transport, &EgMrkYjAnvwk{EbklgbD: "invalid.example"})
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if _, err = conn.Read(make([]byte, 1)); err != expected {
		t.Fatal(err)
	}
	if n, err := conn.Write([]byte("data")); n != 4 || err != expected {
		t.Fatalf("%d %v", n, err)
	}
}
