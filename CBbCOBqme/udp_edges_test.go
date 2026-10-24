package CBbCOBqme

import (
	"context"
	"errors"
	quic "github.com/apernet/quic-go"
	"net"
	B "recovered.local/soga/BaEWsPy"
	W "recovered.local/soga/mlbRsO44"
	"testing"
	"time"
)

type datagramFixture struct {
	quic.Connection
	ctx  context.Context
	sent [][]byte
	err  error
}

func (c *datagramFixture) Context() context.Context { return c.ctx }
func (c *datagramFixture) LocalAddr() net.Addr {
	return &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 101}
}
func (c *datagramFixture) RemoteAddr() net.Addr {
	return &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 102}
}
func (c *datagramFixture) SendDatagram(b []byte) error {
	c.sent = append(c.sent, append([]byte(nil), b...))
	return c.err
}
func TestOriginalUDPDeadlineAndClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	raw := &datagramFixture{ctx: ctx}
	packets := make(chan *W.ItLMumzBf, 8)
	closed := 0
	c := oaAwkKrUsl(raw, 5, packets, &B.D23y2888a_{}, nil, nil, func() { closed++ }).(*aK1tZAfWP2z)
	defer c.g54qtKk8h.Stop()
	defer c.Close()
	c.SetReadDeadline(time.Now().Add(time.Second))
	packets <- &W.ItLMumzBf{Kb9FLg7wuNE: 1, XrDcBXcD: "bad address", HmoFPk: []byte("x")}
	packets <- &W.ItLMumzBf{Kb9FLg7wuNE: 1, XrDcBXcD: "127.0.0.1:8", HmoFPk: make([]byte, 20)}
	packets <- &W.ItLMumzBf{Kb9FLg7wuNE: 1, XrDcBXcD: "127.0.0.1:8", HmoFPk: []byte("ok")}
	buf := make([]byte, 8)
	n, addr, e := c.ReadFromProxy(buf)
	if e != nil || string(buf[:n]) != "ok" || addr.String() != "127.0.0.1:8" {
		t.Fatal(n, addr, e)
	}
	// 原 Write 目的地址采用 QUIC 的 RemoteAddr，不是 GetDestAddr。
	if n, e = c.Write([]byte("reply")); n != 5 || e != nil {
		t.Fatal(n, e)
	}
	p, e := W.CY9jLhcY1f5(raw.sent[0])
	if e != nil || p.XrDcBXcD != "127.0.0.1:102" || p.D35ycLc5aB6r != 5 {
		t.Fatal(p, e)
	}
	raw.err = errors.New("send")
	if n, e = c.Write(buf); n != 0 || e != raw.err {
		t.Fatal(n, e)
	}
	c.SetDeadline(time.Now().Add(-time.Second))
	if n, _, e = c.ReadFromProxy(buf); n != 0 || e == nil || e.Error() != "timeout" {
		t.Fatal(n, e)
	}
	c.SetReadDeadline(time.Time{})
	c.Close()
	c.Close()
	if closed != 1 {
		t.Fatal("close callback", closed)
	}
	if n, _, e = c.ReadFromProxy(buf); n != 0 || e == nil || e.Error() != "closed" {
		t.Fatal(n, e)
	}
	if n, e = c.Write(buf); n != 0 || e == nil || e.Error() != "closed" {
		t.Fatal(n, e)
	}
	if c.SetWriteDeadline(time.Now()) != nil {
		t.Fatal("write deadline")
	}
}
