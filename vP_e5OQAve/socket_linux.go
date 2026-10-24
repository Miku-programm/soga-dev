//go:build linux

package vP_e5OQAve

import (
	"context"
	"errors"
	"go.uber.org/atomic"
	"net"
	"recovered.local/soga/vHUSxG"
	"strconv"
	"strings"
)

// Static/initial state: 0x5095ec8 (init 0x1922b40), 0x5029410 -> 0x50c78e0,
// and nil option slice 0x5099b50. Original setters are not yet recovered.
var interfaceAt5095EC8 = atomic.NewString("")
var markAt5029410 = atomic.NewInt32(0)
var optionsAt5099B50 []PNSpXRb

func collectOptions(options []PNSpXRb) *_yd5wQwv90W {
	o := &_yd5wQwv90W{tUTZ5PkMp: interfaceAt5095EC8.Load(), hPRXKD4ALV: int(markAt5029410.Load())}
	for _, f := range optionsAt5099B50 {
		f(o)
	}
	for _, f := range options {
		f(o)
	}
	return o
}

// eZJf4hKbr: 0x1926b20. The port is the uint64 returned by ParseUint in callers.
func eZJf4hKbr(device, network string, destination net.IP, port uint64) (net.Addr, error) {
	iface, err := vHUSxG.DINaz4(device)
	if err != nil {
		return nil, err
	}
	var selected *net.IPNet
	switch network {
	case "tcp4", "udp4":
		selected, err = iface.PickIPv4Addr(destination)
	case "tcp6", "udp6":
		selected, err = iface.PickIPv6Addr(destination)
	default:
		if destination == nil || destination.To4() != nil {
			selected, err = iface.PickIPv4Addr(destination)
		} else {
			selected, err = iface.PickIPv6Addr(destination)
		}
	}
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(network, "tcp") {
		return &net.TCPAddr{IP: selected.IP, Port: int(port)}, nil
	}
	if strings.HasPrefix(network, "udp") {
		return &net.UDPAddr{IP: selected.IP, Port: int(port)}, nil
	}
	return nil, vHUSxG.AddrNotFoundAt5098660
}

// sFayrZY2WQ: 0x1926e40.
func sFayrZY2WQ(device string, dialer *net.Dialer, network string, destination net.IP) error {
	if !destination.IsGlobalUnicast() {
		return nil
	}
	var port uint64
	if dialer.LocalAddr != nil {
		_, p, err := net.SplitHostPort(dialer.LocalAddr.String())
		if err == nil {
			port, _ = strconv.ParseUint(p, 10, 16)
		}
	}
	local, err := eZJf4hKbr(device, network, destination, port)
	if err != nil {
		return err
	}
	dialer.LocalAddr = local
	return nil
}

// i2ypu3M0u: 0x1926fc0. The ListenConfig argument is present but unused.
func i2ypu3M0u(device string, _ *net.ListenConfig, network, address string) (string, error) {
	_, port, err := net.SplitHostPort(address)
	if err != nil {
		port = "0"
	}
	number, _ := strconv.ParseUint(port, 10, 16)
	local, err := eZJf4hKbr(device, network, nil, number)
	if err != nil {
		return "", err
	}
	return local.String(), nil
}

// rJE25S: 0x1923540. It dials the already-resolved IP using the actual option,
// interface binding, routing mark and source-IP family checks in the sample.
func rJE25S(ctx context.Context, network string, ip net.IP, port string, options ...PNSpXRb) (net.Conn, error) {
	o := collectOptions(options)
	var d net.Dialer
	if o.tUTZ5PkMp != "" {
		if o.oyFnOf5D {
			if err := sFayrZY2WQ(o.tUTZ5PkMp, &d, network, ip); err != nil {
				return nil, err
			}
		} else {
			d.Control = bindDeviceControl(d.Control, o.tUTZ5PkMp)
		}
	}
	if o.hPRXKD4ALV != 0 {
		d.Control = markControl(d.Control, o.hPRXKD4ALV)
	}
	if o.kxe61DaBVTFX != nil {
		if (o.kxe61DaBVTFX.To4() != nil) != (ip.To4() != nil) {
			return nil, errors.New("local addr ip family mismatch with destination")
		}
		switch network {
		case "tcp4", "tcp6":
			d.LocalAddr = &net.TCPAddr{IP: o.kxe61DaBVTFX}
		case "udp4", "udp6":
			d.LocalAddr = &net.UDPAddr{IP: o.kxe61DaBVTFX}
		}
	}
	return d.DialContext(ctx, network, net.JoinHostPort(ip.String(), port))
}

// OpEY1B7hjf: 0x1922e60. Caller-supplied nonempty hosts override the local-IP
// option; a malformed address is replaced by the selected IP and port zero.
func OpEY1B7hjf(ctx context.Context, network, address string, options ...PNSpXRb) (net.PacketConn, error) {
	o := collectOptions(options)
	if o.kxe61DaBVTFX != nil {
		host, port, err := net.SplitHostPort(address)
		if err != nil {
			address = net.JoinHostPort(o.kxe61DaBVTFX.String(), "0")
		} else if host == "" {
			address = net.JoinHostPort(o.kxe61DaBVTFX.String(), port)
		}
	}
	var config net.ListenConfig
	if o.tUTZ5PkMp != "" {
		if o.oyFnOf5D {
			var err error
			address, err = i2ypu3M0u(o.tUTZ5PkMp, &config, network, address)
			if err != nil {
				return nil, err
			}
		} else {
			config.Control = bindDeviceControl(config.Control, o.tUTZ5PkMp)
		}
	}
	if o.zgyOoox {
		config.Control = reuseControl(config.Control)
	}
	if o.hPRXKD4ALV != 0 {
		config.Control = markControl(config.Control, o.hPRXKD4ALV)
	}
	return config.ListenPacket(ctx, network, address)
}
