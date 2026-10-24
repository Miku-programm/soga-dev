package routing

import (
	"fmt"
	"net"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/config"
)

// vhKZu6vCmJ_ restores 0x1b7df80: only literal server IPs yield a candidate.
func (out *GyyO51) vhKZu6vCmJ_() []net.Addr {
	if ip := net.ParseIP(out.AHTlIEynXL); ip != nil {
		return []net.Addr{&net.TCPAddr{IP: ip}}
	}
	return nil
}

// w4foeAH7FyI restores 0x1b7e0c0. It chooses a compatible local/remote
// pair, falling back to the first local IP if no pair has matching families.
func (out *GyyO51) w4foeAH7FyI(local net.Addr, remotes []net.Addr) (net.IP, net.Addr, error) {
	var candidates []net.IP
	if out.iCySnjtd != nil {
		if out.xk9pAT != nil {
			if out.JaH5hv == "ipv6_first" {
				candidates = []net.IP{out.xk9pAT, out.iCySnjtd}
			} else {
				candidates = []net.IP{out.iCySnjtd, out.xk9pAT}
			}
		} else {
			candidates = []net.IP{out.iCySnjtd}
		}
	} else if out.xk9pAT != nil {
		candidates = []net.IP{out.xk9pAT}
	}
	if len(candidates) == 0 && !config.Q0SDIN() && config.AutoOutIPEnabled() {
		switch addr := local.(type) {
		case *net.TCPAddr:
			candidates = []net.IP{addr.IP}
		case *net.UDPAddr:
			candidates = []net.IP{addr.IP}
		default:
			return nil, nil, fmt.Errorf("unknown local addr type: %v", local)
		}
	}
	filtered := make([]net.IP, 0, len(candidates))
	for _, ip := range candidates {
		if !ip.IsLoopback() {
			filtered = append(filtered, ip)
		}
	}
	if len(filtered) == 0 {
		if len(remotes) > 0 {
			return nil, remotes[0], nil
		}
		return nil, nil, nil
	}
	for _, localIP := range filtered {
		for _, remote := range remotes {
			var remoteIP net.IP
			switch addr := remote.(type) {
			case *net.TCPAddr:
				remoteIP = addr.IP
			case *net.UDPAddr:
				remoteIP = addr.IP
			default:
				continue
			}
			if relay.Y0B53iB(localIP, remoteIP) {
				return localIP, remote, nil
			}
		}
	}
	return filtered[0], nil, nil
}
