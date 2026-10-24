// Restored against IDA factory 0x1b24680 and init 0x19c5000.
// Source comparison: github.com/Ieooo/clash/transport/ssr/obfs (GPL-3.0).
// Original dependency revision is unproven; names/layout and deviations follow target evidence.
package GtTssd5z

import (
	"errors"
	"fmt"
	"net"
)

var (
	errorAt5098560 = errors.New("tls1.2_ticket_auth incorrect magic number")
	errorAt5098570 = errors.New("tls1.2_ticket_auth too short data")
	errorAt5098580 = errors.New("tls1.2_ticket_auth hmac verifying failed")
)

type sFaLAK struct {
	l57vr8H [32]byte
}

type AGNBYcIkk interface {
	StreamConn(net.Conn) net.Conn
}

type zQneYx9 func(b *Ja7FBino7HL) AGNBYcIkk

var registryAt50964A8 = make(map[string]struct {
	overhead int
	new      zQneYx9
})

func register(name string, c zQneYx9, o int) {
	registryAt50964A8[name] = struct {
		overhead int
		new      zQneYx9
	}{overhead: o, new: c}
}

func IRT7O1m(name string, b *Ja7FBino7HL) (AGNBYcIkk, int, error) {
	if choice, ok := registryAt50964A8[name]; ok {
		return choice.new(b), choice.overhead, nil
	}
	return nil, 0, fmt.Errorf("Obfs %s not supported", name)
}
