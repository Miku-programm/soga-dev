//go:build linux

package vP_e5OQAve

import (
	"context"
	"errors"
	"net"
	R "recovered.local/soga/LpKMLQuaAsk"
)

// I_y9nlK1q8: 0x1922ba0.
func I_y9nlK1q8(ctx context.Context, network, address string, options ...PNSpXRb) (net.Conn, error) {
	switch network {
	case "tcp", "udp":
		return uns9dRYJYBP(ctx, network, address, options...)
	case "tcp4", "tcp6", "udp4", "udp6":
		host, port, err := net.SplitHostPort(address)
		if err != nil {
			return nil, err
		}
		var ip net.IP
		if network[3] == '4' {
			ip, err = R.Adyk1_7N(host)
		} else {
			ip, err = R.K4HgV0sXi(host)
		}
		if err != nil {
			return nil, err
		}
		return rJE25S(ctx, network, ip, port, options...)
	default:
		return nil, errors.New("network invalid")
	}
}

// bW961oCfn3t retains original 40-byte RTTI layout. Flags mean DNS succeeded,
// IPv6 attempt, and attempt completed, respectively.
type bW961oCfn3t struct {
	net.Conn
	error
	a40PR18CzQ bool
	xGIbOv5    bool
	fqRpuN1iq  bool
}

// uns9dRYJYBP: 0x1923d40; attempt closure: 0x19279e0; deferred delivery/close:
// 0x1927c40. Both families start immediately. Success wins; with two failures,
// a dialing error outranks a DNS error, preferring IPv4 when both reached Dial.
func uns9dRYJYBP(ctx context.Context, network, address string, options ...PNSpXRb) (net.Conn, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}
	done := make(chan struct{})
	defer close(done)
	results := make(chan bW961oCfn3t)
	attempt := func(network string, ipv6 bool) {
		result := bW961oCfn3t{xGIbOv5: ipv6, fqRpuN1iq: true}
		defer func() {
			select {
			case results <- result:
			case <-done:
				if result.Conn != nil {
					_ = result.Conn.Close()
				}
			}
		}()
		var ip net.IP
		if ipv6 {
			ip, result.error = R.K4HgV0sXi(host)
		} else {
			ip, result.error = R.Adyk1_7N(host)
		}
		if result.error != nil {
			return
		}
		result.a40PR18CzQ = true
		result.Conn, result.error = rJE25S(ctx, network, ip, port, options...)
	}
	go attempt(network+"4", false)
	go attempt(network+"6", true)
	var v4, v6 bW961oCfn3t
	for result := range results {
		if result.error == nil {
			return result.Conn, nil
		}
		if result.xGIbOv5 {
			v6 = result
		} else {
			v4 = result
		}
		if v4.fqRpuN1iq && v6.fqRpuN1iq {
			if v4.a40PR18CzQ {
				return nil, v4.error
			}
			if v6.a40PR18CzQ {
				return nil, v6.error
			}
			return nil, v4.error
		}
	}
	return nil, errors.New("never touched")
}
