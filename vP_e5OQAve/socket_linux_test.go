//go:build linux

package vP_e5OQAve

import (
	"context"
	"net"
	"testing"
	"time"
)

func TestResolvedDialSourceAndFamily(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	_, port, _ := net.SplitHostPort(listener.Addr().String())
	accepted := make(chan net.Conn, 1)
	go func() {
		c, err := listener.Accept()
		if err == nil {
			accepted <- c
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c, err := rJE25S(ctx, "tcp4", net.ParseIP("127.0.0.1"), port, Cv8Pwd5lLu(net.ParseIP("127.0.0.1")))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	other := <-accepted
	defer other.Close()
	if !c.LocalAddr().(*net.TCPAddr).IP.IsLoopback() {
		t.Fatal(c.LocalAddr())
	}
	c, err = rJE25S(ctx, "tcp4", net.ParseIP("127.0.0.1"), port, Cv8Pwd5lLu(net.ParseIP("::1")))
	if c != nil || err == nil || err.Error() != "local addr ip family mismatch with destination" {
		t.Fatal(c, err)
	}
}
func TestListenLocalIPSelection(t *testing.T) {
	for _, address := range []string{"", ":0", "invalid"} {
		c, err := OpEY1B7hjf(context.Background(), "udp4", address, Cv8Pwd5lLu(net.ParseIP("127.0.0.1")))
		if err != nil {
			t.Fatal(address, err)
		}
		if !c.LocalAddr().(*net.UDPAddr).IP.IsLoopback() {
			t.Fatal(c.LocalAddr())
		}
		c.Close()
	}
	c, err := OpEY1B7hjf(context.Background(), "udp4", "127.0.0.1:0", Cv8Pwd5lLu(net.ParseIP("192.0.2.1")))
	if err != nil {
		t.Fatal("explicit host must override option", err)
	}
	c.Close()
}
func TestInterfaceOptionSelection(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}
	var name string
	for _, iface := range interfaces {
		if iface.Flags&net.FlagLoopback != 0 {
			name = iface.Name
			break
		}
	}
	if name == "" {
		t.Skip("no loopback interface")
	}
	local, err := eZJf4hKbr(name, "udp4", nil, 7)
	if err != nil {
		t.Fatal(err)
	}
	if !local.(*net.UDPAddr).IP.IsLoopback() || local.(*net.UDPAddr).Port != 7 {
		t.Fatal(local)
	}
	address, err := i2ypu3M0u(name, new(net.ListenConfig), "udp4", "bad")
	if err != nil || address != "127.0.0.1:0" {
		t.Fatal(address, err)
	}
}
