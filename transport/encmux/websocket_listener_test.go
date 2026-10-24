package encmux

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	ws "recovered.local/soga/transport/websocket"
	util "recovered.local/soga/transport/wsutil"
	"reflect"
	"testing"
	"time"
)

func TestOriginalWebsocketListenerTCP(t *testing.T) {
	raw, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	listener := GHrg5b6zAf(raw, "/tunnel")
	defer listener.Close()
	c, e := net.Dial("tcp", raw.Addr().String())
	if e != nil {
		t.Fatal(e)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(4 * time.Second))
	io.WriteString(c, "GET /tunnel HTTP/1.1\r\nHost: native\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Version: 13\r\nSec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==\r\n\r\n")
	response, e := http.ReadResponse(bufio.NewReader(c), nil)
	if e != nil || response.StatusCode != 101 {
		t.Fatal(response, e)
	}
	s, e := listener.Accept()
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	s.SetDeadline(time.Now().Add(4 * time.Second))
	data := bytes.Repeat([]byte{0, 255, 3}, 20000)
	done := make(chan error, 1)
	go func() {
		if e := util.WriteFrameRecovered(c, ws.StateClientSide, ws.OpBinary, true, data[:17000]); e != nil {
			done <- e
			return
		}
		done <- util.WriteFrameRecovered(c, ws.StateClientSide, ws.OpBinary, true, data[17000:])
	}()
	b := make([]byte, len(data))
	if _, e = io.ReadFull(s, b); e != nil || !bytes.Equal(data, b) {
		t.Fatal(e)
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
	go func() { _, e := s.Write(data); done <- e }()
	r := &util.HLKO3Cdqw{GX04GFKh1Gz: c, WXUh5psqY7: ws.StateClientSide}
	h, e := r.NextFrame()
	if e != nil || h.IsHe8C != ws.OpBinary {
		t.Fatal(h, e)
	}
	b, e = io.ReadAll(r)
	if e != nil || !bytes.Equal(data, b) {
		t.Fatal(e)
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
	listener.Close()
	if _, e = listener.Accept(); e == nil || e.Error() != "ws listener closed" {
		t.Fatal(e)
	}
}

func TestOriginalWebsocketLayoutsAndWrapperOrder(t *testing.T) {
	for _, tc := range []struct {
		v       any
		size    uintptr
		offsets []uintptr
	}{{YE2KB1I{}, 48, []uintptr{0, 16, 24, 40}}, {V72bzf8{}, 64, []uintptr{0, 16, 32, 40, 56}}} {
		typ := reflect.TypeOf(tc.v)
		if typ.Size() != tc.size {
			t.Fatal(typ)
		}
		for i, x := range tc.offsets {
			if typ.Field(i).Offset != x {
				t.Fatal(typ.Field(i))
			}
		}
	}
	raw, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer raw.Close()
	cfg := &VTl94pzz2IkB{hZORuZstvS07: true, rj1DhSBHZ7O: true}
	if l, e := cfg.WrapListener(raw); e != nil || l != raw {
		t.Fatal(l, e)
	}
	cfg.hZORuZstvS07 = false
	if l, e := cfg.WrapListener(raw); l != nil || e == nil || e.Error() != "tunnel config certificate is nil" {
		t.Fatal(l, e)
	}
	cfg.rj1DhSBHZ7O = false
	l, e := HNbZee047(raw, cfg)
	if e != nil {
		t.Fatal(e)
	}
	outer, ok := l.(*Ppx7OU)
	if !ok {
		t.Fatalf("%T", l)
	}
	if encrypted, ok := outer.O6B2qZ0p0.(*T4k2R7cJ); !ok || encrypted.O6B2qZ0p0 != raw {
		t.Fatal("original encrypted-before-verification wrapper order")
	}
}
