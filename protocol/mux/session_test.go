package mux

import (
	"context"
	"encoding/binary"
	"io"
	"net"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/user"
	"reflect"
	"testing"
	"time"
)

type muxInput struct{ net.Conn }

func (*muxInput) GetUser() user.PgaNd3014U       { return nil }
func (*muxInput) GetDestAddr() *relay.AHU2MRmjzQu { return nil }
func (*muxInput) NetworkType() sess.ZES6F5    { return 2 }
func newWire(id uint16, network byte, payload string) []byte {
	meta := binary.BigEndian.AppendUint16(nil, id)
	meta = append(meta, 1, 1, network, 0, 53, 1, 127, 0, 0, 1)
	return frameBytes(meta, []byte(payload), true)
}
func TestSessionRealMultiplexing(t *testing.T) {
	a, b := net.Pipe()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	b.SetDeadline(time.Now().Add(5 * time.Second))
	p := KJwDQE2y3(ctx, &muxInput{a}, nil)
	defer func() { b.Close(); p.Close() }()
	for _, v := range []struct {
		id      uint16
		network byte
		payload string
	}{{11, 1, "first"}, {12, 2, "udp"}} {
		if _, err := b.Write(newWire(v.id, v.network, v.payload)); err != nil {
			t.Fatal(err)
		}
		c, ok := p.Accept()
		if !ok {
			t.Fatal("accept")
		}
		if c.NetworkType() != sess.ZES6F5(v.network-1) {
			t.Fatal(c.NetworkType())
		}
		buf := make([]byte, 10)
		if udp, ok := c.(*WSnkIcJKU); ok {
			n, addr, err := udp.ReadFromProxy(buf)
			if err != nil || string(buf[:n]) != v.payload || addr.String() != "127.0.0.1:53" {
				t.Fatal(n, addr, err)
			}
		} else {
			n, err := c.Read(buf)
			if err != nil || string(buf[:n]) != v.payload {
				t.Fatal(n, err)
			}
		}
		done := make(chan error, 1)
		go func() { _, err := c.Write([]byte("response")); done <- err }()
		f, err := HlVrAnC111(b)
		if err != nil {
			t.Fatal(err)
		}
		if f.ljtkJw != v.id || string(f.c8AtRj7hb.Bytes()) != "response" {
			t.Fatal(f)
		}
		f.FreeData()
		if err = <-done; err != nil {
			t.Fatal(err)
		}
		// End causes the original stream Close to acknowledge End on the same connection.
		meta := binary.BigEndian.AppendUint16(nil, v.id)
		meta = append(meta, 3, 0)
		if _, err = b.Write(frameBytes(meta, nil, false)); err != nil {
			t.Fatal(err)
		}
		f, err = HlVrAnC111(b)
		if err != nil || f.ljtkJw != v.id || f.xvTw9VQDQdY != 3 {
			t.Fatal(f, err)
		}
		if _, err = c.Read(buf); err != io.EOF {
			t.Fatal(err)
		}
	}
	b.Close()
	select {
	case <-p.aojobe.Done():
	case <-time.After(time.Second):
		t.Fatal("session did not cancel on EOF")
	}
	if c, ok := p.Accept(); ok || c != nil {
		t.Fatal(c, ok)
	}
}
func TestSessionHandlersRejectAndRelease(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	raw := &muxCapture{}
	p := &Ps5tpGhQ{aojobe: ctx, bYFyxREdBX: cancel, mECvR4TiwY: raw, zbRz7R: &XtJioO{}, i1xUuKWJ: make(chan sess.Ki8urSHTx44, 8)}
	bad := dataFrame("bad")
	bad.ljtkJw = 7
	bad.ePmcTrBC = 9
	if err := p.caYqS8sJ(bad); err == nil || err.Error() != "unknown mux net type: 9" || bad.c8AtRj7hb != nil {
		t.Fatal(err, bad)
	}
	f := dataFrame("a")
	f.ljtkJw = 8
	f.ePmcTrBC = 1
	if err := p.caYqS8sJ(f); err != nil {
		t.Fatal(err)
	}
	c, _ := p.Accept()
	defer c.(*Ldrr25M).Free()
	duplicate := dataFrame("x")
	duplicate.ljtkJw = 8
	if err := p.caYqS8sJ(duplicate); err == nil || err.Error() != "repeat mux id: 8" || duplicate.c8AtRj7hb != nil {
		t.Fatal(err)
	}
	missing := dataFrame("drop")
	missing.ljtkJw = 99
	p.l4vHJHVTLmEy(missing)
	if missing.c8AtRj7hb != nil {
		t.Fatal(missing)
	}
	end := dataFrame("end")
	end.ljtkJw = 8
	end.m3vf0n8 = 3
	if err := p.ie6rc5Jas(end); err == nil || err.Error() != "session has error" || end.c8AtRj7hb != nil {
		t.Fatal(err)
	}
	if p.zbRz7R.f56bH6.Len() != 0 {
		t.Fatal("end not removed")
	}
}
func TestSessionRTTI(t *testing.T) {
	typ := reflect.TypeFor[Ps5tpGhQ]()
	if typ.Size() != 80 {
		t.Fatal(typ.Size())
	}
	for i, off := range []uintptr{0, 16, 24, 40, 48, 64, 72} {
		if typ.Field(i).Offset != off {
			t.Fatal(typ.Field(i))
		}
	}
}
