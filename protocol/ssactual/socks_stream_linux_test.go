//go:build linux

package ssactual

import (
	"bytes"
	"context"
	"io"
	"net"
	E "recovered.local/soga/net/outbound"
	"testing"
	"time"
)

func TestSocksDialAgainstLocalServer(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	done := make(chan error, 1)
	go func() {
		c, err := listener.Accept()
		if err != nil {
			done <- err
			return
		}
		defer c.Close()
		c.SetDeadline(time.Now().Add(time.Second))
		hello := make([]byte, 3)
		_, err = io.ReadFull(c, hello)
		if err != nil {
			done <- err
			return
		}
		if !bytes.Equal(hello, []byte{5, 1, 0}) {
			done <- io.ErrUnexpectedEOF
			return
		}
		c.Write([]byte{5, 0})
		request := make([]byte, 10)
		_, err = io.ReadFull(c, request)
		if err != nil {
			done <- err
			return
		}
		if !bytes.Equal(request, []byte{5, 1, 0, 1, 192, 0, 2, 1, 0, 80}) {
			done <- io.ErrUnexpectedEOF
			return
		}
		c.Write([]byte{5, 0, 0, 1, 127, 0, 0, 1, 0, 9})
		_, err = io.Copy(c, c)
		done <- err
	}()
	a := Auqn90u2NBQ(SocksOptions{Name: "socks", Server: "127.0.0.1", Port: listener.Addr().(*net.TCPAddr).Port})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c, err := a.DialContext(ctx, &E.D0mLgQF{ZouEpmjw8: net.IPv4(192, 0, 2, 1), JRXMVvg: 80})
	if err != nil {
		t.Fatal(err)
	}
	c.SetDeadline(time.Now().Add(time.Second))
	c.Write([]byte("echo"))
	b := make([]byte, 4)
	if _, err = io.ReadFull(c, b); err != nil || string(b) != "echo" || c.Chains()[0] != "socks" {
		t.Fatal(string(b), err, c.Chains())
	}
	c.Close()
	if err = <-done; err != nil {
		t.Fatal(err)
	}
}
