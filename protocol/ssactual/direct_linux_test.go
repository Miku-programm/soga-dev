//go:build linux

package ssactual

import (
	"context"
	"io"
	"net"
	E "recovered.local/soga/net/outbound"
	O "recovered.local/soga/net/socket"
	"testing"
	"time"
)

func TestDirectActualTCPAndUDP(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	done := make(chan error, 1)
	go func() {
		c, e := listener.Accept()
		if e == nil {
			defer c.Close()
			_, e = io.Copy(c, c)
		}
		done <- e
	}()
	d := NewDirectAdapter()
	target := &E.D0mLgQF{ZouEpmjw8: net.ParseIP("127.0.0.1"), JRXMVvg: E.AN18DDo(listener.Addr().(*net.TCPAddr).Port)}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c, err := d.DialContext(ctx, target)
	if err != nil {
		t.Fatal(err)
	}
	c.SetDeadline(time.Now().Add(time.Second))
	c.Write([]byte("ok"))
	data := make([]byte, 2)
	if _, err = io.ReadFull(c, data); err != nil || string(data) != "ok" {
		t.Fatal(data, err)
	}
	if len(c.Chains()) != 1 || c.Chains()[0] != "DIRECT" {
		t.Fatal(c.Chains())
	}
	c.Close()
	if err = <-done; err != nil {
		t.Fatal(err)
	}
	p, err := d.ListenPacketContext(ctx, nil, O.Cv8Pwd5lLu(net.ParseIP("127.0.0.1")))
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
	if p.Chains()[0] != "DIRECT" || !p.LocalAddr().(*net.UDPAddr).IP.IsLoopback() {
		t.Fatal(p.Chains(), p.LocalAddr())
	}
}
