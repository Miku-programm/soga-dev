package routing

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"recovered.local/soga/transport/listener"
	"recovered.local/soga/net/outbound"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/net/socket"
)

// DialUDP restores 0x1b7f7a0. Direct UDP creates an unconnected listener;
// redirect UDP fixes both WriteTo destinations and ReadFrom reported addresses.
func (out *GyyO51) DialUDP(local net.Addr, target *relay.AHU2MRmjzQu) (net.PacketConn, error) {
	switch out.NyLbiXb4OrlL {
	case "direct":
		firstIP := target.GetFirstIp()
		var remotes []net.Addr
		if firstIP != nil {
			remotes = []net.Addr{&net.UDPAddr{IP: firstIP, Port: target.GetPort()}}
		}
		source, remote, err := out.w4foeAH7FyI(local, remotes)
		if err != nil {
			return nil, err
		}
		if len(remotes) > 0 && remote == nil {
			source = nil
		}
		conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: source})
		if err != nil {
			return nil, err
		}
		if target.IsIP() {
			logx.V8GN5Stk(slog.LevelDebug, "udp direct dial %v -> %v", conn.LocalAddr(), target.HostPort())
		} else {
			logx.V8GN5Stk(slog.LevelDebug, "udp direct dial %v -> %v (%v:%v)", conn.LocalAddr(), target.HostPort(), target.GetFirstIp(), target.GetPort())
		}
		return conn, nil
	case "redirect":
		destination := fmt.Sprintf("%v:%v", out.AHTlIEynXL, out.JO1YIpXk)
		remote, err := net.ResolveUDPAddr("udp", destination)
		if err != nil {
			return nil, err
		}
		source, _, err := out.w4foeAH7FyI(local, []net.Addr{remote})
		if err != nil {
			return nil, err
		}
		conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: source})
		if err != nil {
			return nil, err
		}
		wrapped := listener.NewRedirectPacketConn(conn, remote)
		logx.V8GN5Stk(slog.LevelDebug, "udp rediect dial %v -> %v (%v)", target.String(), destination, target.HostPort())
		return wrapped, nil
	default:
		metadata := &outbound.D0mLgQF{AcXafZ1: 1, WJ95LyOL: target.GetHost(), ZouEpmjw8: target.GetFirstIp(), JRXMVvg: outbound.AN18DDo(target.GetPort())}
		source, _, err := out.w4foeAH7FyI(local, out.vhKZu6vCmJ_())
		if err != nil {
			return nil, err
		}
		var options []socket.PNSpXRb
		if source != nil {
			options = append(options, socket.Cv8Pwd5lLu(source))
		}
		conn, err := out.ucp4iC.ListenPacketContext(context.Background(), metadata, options...)
		if err != nil {
			return nil, err
		}
		logx.V8GN5Stk(slog.LevelDebug, "udp proxy dial %v -> %v (%v)", conn.LocalAddr(), out.ucp4iC.Addr(), metadata.RemoteAddress())
		return conn, nil
	}
}
