package mux

import (
	"bytes"
	"errors"
	"io"
	"net"
	"os"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/util/deadline2"
	"reflect"
	"testing"
	"time"
)

type muxCapture struct {
	net.Conn
	bytes.Buffer
	err   error
	short bool
}

func (c *muxCapture) Read(p []byte) (int, error) { return c.Buffer.Read(p) }
func (c *muxCapture) Write(p []byte) (int, error) {
	if c.err != nil {
		return 0, c.err
	}
	if c.short {
		return 0, nil
	}
	return c.Buffer.Write(p)
}
func newStreamFixture(c net.Conn) *Ldrr25M {
	return &Ldrr25M{gwE4Or1X8JW: 0x1234, fj8yHOR: make(chan *CUwik6VP2WU0, 2), hiMx_c: c, s1w1RA: deadline2.NewDeadline()}
}
func dataFrame(p string) *CUwik6VP2WU0 {
	b := pool.S4OQXWIZ(len(p))
	b.WriteString(p)
	return &CUwik6VP2WU0{m3vf0n8: 1, c8AtRj7hb: b}
}
func TestStreamPartialDeadlineAndCloseOwnership(t *testing.T) {
	raw := &muxCapture{}
	c := newStreamFixture(raw)
	f := dataFrame("abcd")
	c.Feed(f)
	p := make([]byte, 2)
	n, e := c.Read(p)
	if n != 2 || e != nil || string(p) != "ab" {
		t.Fatal(n, e, p)
	}
	c.SetReadDeadline(time.Now().Add(-time.Second))
	n, e = c.Read(p)
	if n != 2 || e != nil || string(p) != "cd" || c.altuteX != nil {
		t.Fatal(n, e, p)
	}
	if _, e = c.Read(p); e != os.ErrDeadlineExceeded {
		t.Fatal(e)
	}
	c.SetReadDeadline(time.Time{})
	queued := dataFrame("pending")
	c.Feed(queued)
	if e = c.Close(); e != nil {
		t.Fatal(e)
	}
	c.Close()
	if queued.c8AtRj7hb != nil {
		t.Fatal("queued payload not freed")
	}
	if !bytes.Equal(raw.Bytes(), []byte{0, 4, 0x12, 0x34, 3, 0}) {
		t.Fatalf("%x", raw.Bytes())
	}
	if _, e = c.Read(p); e != io.EOF {
		t.Fatal(e)
	}
	late := dataFrame("late")
	c.Feed(late)
	if late.c8AtRj7hb != nil {
		t.Fatal("closed feed did not free")
	}
}
func TestStreamRealWireAndWriteSemantics(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	c := newStreamFixture(a)
	done := make(chan error, 1)
	go func() {
		n, e := c.Write([]byte("hello"))
		if n != 5 && e == nil {
			e = io.ErrShortWrite
		}
		done <- e
	}()
	f, e := HlVrAnC111(b)
	if e != nil {
		t.Fatal(e)
	}
	defer f.FreeData()
	if f.ljtkJw != 0x1234 || f.xvTw9VQDQdY != 2 || !f.HasData() || string(f.c8AtRj7hb.Bytes()) != "hello" {
		t.Fatal(f)
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
	boom := errors.New("write")
	raw := &muxCapture{err: boom}
	c.hiMx_c = raw
	if n, e := c.Write([]byte("a")); n != 0 || e != boom {
		t.Fatal(n, e)
	}
	raw.err = nil
	raw.short = true
	if n, e := c.Write([]byte("a")); n != 1 || e != nil {
		t.Fatal(n, e)
	}
}
func TestDatagramAddressSwitchAndWriteError(t *testing.T) {
	raw := &muxCapture{}
	base := newStreamFixture(raw)
	base.yFmjxzi, _ = relay.U80FDt("127.0.0.1:53")
	c := &WSnkIcJKU{Ldrr25M: base}
	f := dataFrame("abcd")
	c.Feed(f)
	p := make([]byte, 2)
	n, a, e := c.ReadFromProxy(p)
	if n != 2 || e != nil || a.String() != "127.0.0.1:53" || c.dBLfTvhGz || f.c8AtRj7hb != nil {
		t.Fatal(n, a, e)
	}
	f = dataFrame("xy")
	f.hUHpeGfV, _ = relay.U80FDt("[::1]:99")
	c.Feed(f)
	n, a, e = c.ReadFromProxy(p)
	if n != 2 || e != nil || a.String() != "[::1]:99" || !c.dBLfTvhGz {
		t.Fatal(n, a, e)
	}
	for _, host := range []string{"127.1.2.3:80", "[::1]:8080"} {
		raw.Reset()
		addr, _ := net.ResolveUDPAddr("udp", host)
		n, e = c.WriteToProxy([]byte("xyz"), addr)
		if n != 3 || e != nil {
			t.Fatal(n, e)
		}
		frame, e := HlVrAnC111(bytes.NewReader(raw.Bytes()))
		if e != nil {
			t.Fatal(e)
		}
		if frame.hUHpeGfV.Join() != host || frame.ePmcTrBC != 2 || frame.ljtkJw != 0x1234 || string(frame.c8AtRj7hb.Bytes()) != "xyz" {
			t.Fatal(frame)
		}
		frame.FreeData()
	}
	raw.err = errors.New("write failed")
	addr := &net.UDPAddr{IP: net.IPv4(1, 2, 3, 4), Port: 53}
	if n, e = c.WriteToProxy(p, addr); n != len(p) || e != nil {
		t.Fatal("original address branch ignores write error", n, e)
	}
	c.dBLfTvhGz = false
	if n, e = c.WriteToProxy(p, addr); n != 0 || e != raw.err {
		t.Fatal(n, e)
	}
}
func TestMuxConnectionLayouts(t *testing.T) {
	for _, v := range []struct {
		x       any
		size    uintptr
		offsets []uintptr
	}{{Ldrr25M{}, 88, []uintptr{0, 12, 16, 24, 40, 56, 64, 80}}, {XtJioO{}, 56, []uintptr{0}}, {WSnkIcJKU{}, 24, []uintptr{0, 8, 16}}} {
		typ := reflect.TypeOf(v.x)
		if typ.Size() != v.size {
			t.Fatalf("%s %d", typ, typ.Size())
		}
		for i, off := range v.offsets {
			if typ.Field(i).Offset != off {
				t.Fatalf("%s.%s", typ, typ.Field(i).Name)
			}
		}
	}
}
