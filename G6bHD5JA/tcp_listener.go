package G6bHD5JA

import (
	"context"
	"errors"
	"net"
	"recovered.local/soga/EiDD1sTNvpQ"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"time"
)

type O6B2qZ0p0 = net.Listener
type AKJHYc5zd struct{ O6B2qZ0p0 }

func D9_HjRw(addr *net.TCPAddr) (net.Listener, error) { //0x15b2be0
	config := net.ListenConfig{KeepAlive: 60 * time.Second, Control: EiDD1sTNvpQ.CyQmxg(EiDD1sTNvpQ.UA1QxZBYiPM)}
	if cNT24rG4F.MPTCPEnabled() {
		config.SetMultipathTCP(true)
	}
	l, err := config.Listen(context.Background(), "tcp", addr.String())
	if err != nil {
		return nil, err
	}
	return &AKJHYc5zd{O6B2qZ0p0: l}, nil
}

// Accept @0x15b1f00 retries every error except net.ErrClosed. The original does
// not check Temporary/Timeout and does not close a connection returned with err.
func (l *AKJHYc5zd) Accept() (net.Conn, error) {
	delay := 5 * time.Millisecond
	for {
		c, err := l.O6B2qZ0p0.Accept()
		if err == nil {
			return c, nil
		}
		if errors.Is(err, net.ErrClosed) {
			return nil, err
		}
		muc279Mpa6S4.V8GN5Stk(-4, "accept %v err: %v, retry after %v", l.O6B2qZ0p0.Addr(), err, delay)
		time.Sleep(delay)
		delay = min(delay*2, time.Second)
	}
}
