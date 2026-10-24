package routing

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"recovered.local/soga/net/outbound"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/net/socket"
)

// DialTCP restores 0x1b7ea40, including original address fallback and logs.
func (out *GyyO51) DialTCP(local net.Addr, target *relay.AHU2MRmjzQu) (net.Conn, error) {
	switch out.NyLbiXb4OrlL {
	case "direct":
		addresses, err := target.ToTCPAddrs()
		if err != nil {
			return nil, err
		}
		remotes := make([]net.Addr, 0, len(addresses))
		for _, addr := range addresses {
			remotes = append(remotes, addr)
		}
		source, remote, err := out.w4foeAH7FyI(local, remotes)
		if err != nil {
			return nil, err
		}
		if remote == nil {
			source = nil
			remote = remotes[0]
		}
		conn, err := relay.DE9CHPNLbK(&net.TCPAddr{IP: source}, remote.String())
		if err != nil {
			return nil, err
		}
		if target.IsIP() {
			logx.V8GN5Stk(slog.LevelDebug, "tcp direct dial %v -> %v", conn.LocalAddr(), conn.RemoteAddr())
		} else {
			logx.V8GN5Stk(slog.LevelDebug, "tcp direct dial %v -> %v (%v)", conn.LocalAddr(), target.HostPort(), conn.RemoteAddr())
		}
		return conn, nil
	case "redirect":
		source, _, err := out.w4foeAH7FyI(local, out.vhKZu6vCmJ_())
		if err != nil {
			return nil, err
		}
		port := out.JO1YIpXk
		if port == 0 {
			port = target.GetPort()
		}
		destination := fmt.Sprintf("%v:%v", out.AHTlIEynXL, port)
		conn, err := relay.DE9CHPNLbK(&net.TCPAddr{IP: source}, destination)
		if err != nil {
			return nil, err
		}
		logx.V8GN5Stk(slog.LevelDebug, "tcp rediect dial %v -> %v (%v)", conn.LocalAddr(), destination, target.HostPort())
		return conn, nil
	default:
		var destinationIP net.IP
		if ips := target.GetIPs(); len(ips) > 0 {
			destinationIP = ips[0]
		}
		metadata := &outbound.D0mLgQF{WJ95LyOL: target.GetHost(), ZouEpmjw8: destinationIP, JRXMVvg: outbound.AN18DDo(target.GetPort())}
		source, _, err := out.w4foeAH7FyI(local, out.vhKZu6vCmJ_())
		if err != nil {
			return nil, err
		}
		var options []socket.PNSpXRb
		if source != nil {
			options = append(options, socket.Cv8Pwd5lLu(source))
		}
		conn, err := out.ucp4iC.DialContext(context.Background(), metadata, options...)
		if err != nil {
			return nil, err
		}
		logx.V8GN5Stk(slog.LevelDebug, "tcp proxy dial %v -> %v (%v)", conn.LocalAddr(), out.ucp4iC.Addr(), metadata.RemoteAddress())
		return conn, nil
	}
}
