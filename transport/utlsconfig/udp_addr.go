package utlsconfig

import (
	"fmt"
	"net"
	"strconv"
)

// Q_ZqnPkUlb9 @0x1171500. A UDP address is returned by identity. All other
// address implementations are parsed as literal IP:port; no DNS or port-range
// validation is added. In particular this does not preserve a TCP IPv6 zone.
func Q_ZqnPkUlb9(addr net.Addr) (*net.UDPAddr, error) {
	if a, ok := addr.(*net.UDPAddr); ok {
		return a, nil
	}
	host, port, err := net.SplitHostPort(addr.String())
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return nil, fmt.Errorf("parse ip failed: %v", host)
	}
	return &net.UDPAddr{IP: ip, Port: n}, nil
}
