package G6bHD5JA

import (
	"errors"
	"fmt"
	"net"
	"testing"
	"testing/synctest"
	"time"
)

type retryListener struct {
	net.Listener
	errors []error
	at     []time.Time
}

func (l *retryListener) Accept() (net.Conn, error) {
	l.at = append(l.at, time.Now())
	if len(l.errors) == 0 {
		return nil, nil
	}
	e := l.errors[0]
	l.errors = l.errors[1:]
	return nil, e
}
func (l *retryListener) Addr() net.Addr { return &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1} }
func TestOriginalAcceptRetriesNonTemporaryErrors(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		raw := &retryListener{}
		for i := 0; i < 10; i++ {
			raw.errors = append(raw.errors, errors.New("non-temporary failure"))
		}
		l := AKJHYc5zd{O6B2qZ0p0: raw}
		if c, err := l.Accept(); c != nil || err != nil {
			t.Fatal(c, err)
		}
		for i, want := range []time.Duration{5, 10, 20, 40, 80, 160, 320, 640, 1000, 1000} {
			if delta := raw.at[i+1].Sub(raw.at[i]); delta != want*time.Millisecond {
				t.Fatal(i, delta, want)
			}
		}
		closed := fmt.Errorf("wrapped: %w", net.ErrClosed)
		raw.errors = []error{closed}
		start := time.Now()
		if c, err := l.Accept(); c != nil || err != closed {
			t.Fatal(c, err)
		}
		if !time.Now().Equal(start) {
			t.Fatal("closed error delayed")
		}
	})
}
func TestOriginalTCPListenerRealSocket(t *testing.T) {
	l, err := D9_HjRw(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	done := make(chan error, 1)
	go func() {
		c, err := l.Accept()
		if err == nil {
			err = c.Close()
		}
		done <- err
	}()
	c, err := net.DialTimeout("tcp", l.Addr().String(), time.Second)
	if err != nil {
		t.Fatal(err)
	}
	c.Close()
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	l.Close()
	if _, err = l.Accept(); !errors.Is(err, net.ErrClosed) {
		t.Fatal(err)
	}
}
