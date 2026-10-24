package listener

import (
	"crypto/tls"
	"fmt"
	"net"
	H2 "recovered.local/soga/transport/h2"
	G "recovered.local/soga/net/relay"
	P "recovered.local/soga/transport/proxyproto"
	GRPC "recovered.local/soga/transport/grpc"
	X "recovered.local/soga/transport/xhttp"
	T "recovered.local/soga/transport/encmux"
	"recovered.local/soga/reality"
)

// NewListener @0x18b9340 preserves original wrapper order and cleanup. The
// defer at 0x18c2800 reads the local setup error, not the returned error: the
// explicit invalid-network/Reality branches therefore do not close the raw
// listener. This original distinction is intentionally retained.
func (o *Ap9uD2) NewListener() (net.Listener, error) {
	listener, setupError := G.D9_HjRw(&net.TCPAddr{IP: net.ParseIP(o.Host), Port: o.Port})
	if setupError != nil {
		return nil, setupError
	}
	defer func() {
		if setupError != nil {
			_ = listener.Close()
		}
	}()
	if o.ProxyProtocol {
		p := &P.KprDQw4aZEOg{I9Hb4SAu: listener}
		if o.ForceProxyProtocol {
			p.Uw20KUu1 = func(net.Addr) (P.XTAkygU, error) { return 3, nil }
		} //0x18c4080
		listener = p
	}
	if o.Tunnel != nil {
		var wrapped net.Listener
		wrapped, setupError = T.HNbZee047(listener, o.Tunnel)
		if setupError != nil {
			return nil, setupError
		}
		listener = wrapped
	}
	if o.Network == "" {
		o.Network = "tcp"
	}
	switch o.Network {
	case "tcp", "xhttp":
		if o.TLS != nil {
			listener = tls.NewListener(listener, o.TLS)
		} else if o.Reality != nil {
			listener = reality.NewListener(listener, o.Reality)
		}
	case "ws":
		if o.TLS != nil {
			listener = tls.NewListener(listener, o.TLS)
		} else if o.Reality != nil {
			return nil, fmt.Errorf("reality not support websocket")
		}
	case "h2", "grpc":
		if o.Reality != nil {
			return nil, fmt.Errorf("reality not support h2 or grpc")
		}
	default:
		return nil, fmt.Errorf("unknown stream type: <%v>", o.Network)
	}
	switch o.Network {
	case "h2":
		listener = H2.Qi8msD8m74(listener, o.Path, o.TLS)
	case "ws":
		listener = H7N2ayu(listener, o.Path)
	case "grpc":
		listener = GRPC.SNiVo8(listener, o.ServiceName, o.TLS)
	case "xhttp":
		listener = X.DWB3oG(listener, o.XHTTP)
	}
	return listener, nil
}
