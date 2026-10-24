package relay

import (
	"fmt"
	"net"
	"strconv"
)

// JY6gu_Rqfz corresponds to 0x15b3200. TCP pointers are retained; UDP
// conversion shares the IP slice and discards the zone. Other Addr values
// must contain a numeric IP; this function does not perform DNS resolution.
func JY6gu_Rqfz(addr net.Addr) (*net.TCPAddr, error) {
	switch a := addr.(type) {
	case *net.TCPAddr:
		return a, nil
	case *net.UDPAddr:
		return &net.TCPAddr{IP: a.IP, Port: a.Port}, nil
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
	return &net.TCPAddr{IP: ip, Port: n}, nil
}

// LQYEqaxN corresponds to 0x15b2fa0; unlike JY6gu_Rqfz it only has the
// identical UDP-pointer fast path. Even TCPAddr goes through String parsing.
func LQYEqaxN(addr net.Addr) (*net.UDPAddr, error) {
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
