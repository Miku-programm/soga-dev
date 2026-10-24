//go:build linux

package LCe4Wkk2

import (
	"context"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/netip"
	S "recovered.local/soga/Bx4XCl"
	E "recovered.local/soga/E3F12urwswL"
	R "recovered.local/soga/LpKMLQuaAsk"
	O "recovered.local/soga/vP_e5OQAve"
	"time"
)

// hNLZlR1I retains the original 48-byte RTTI layout.
type hNLZlR1I struct {
	net.PacketConn
	iYgXoENCDQ5w net.Addr
	f67_L17asEA  net.Conn
}

// WriteTo: 0x1b66540. Returns the encoded datagram's write count, including
// its SOCKS header, as the original does.
func (c *hNLZlR1I) WriteTo(b []byte, addr net.Addr) (int, error) {
	packet, err := S.XFq85XFER(S.KvyKMZ(addr), b)
	if err != nil {
		return 0, err
	}
	return c.PacketConn.WriteTo(packet, c.iYgXoENCDQ5w)
}

// wireUDPAddr is a factored copy of the identical inlined conversions in
// 0x1b66620 and 0x1b64fc0. Domain addresses return nil; IP bytes are copied.
func wireUDPAddr(address []byte) *net.UDPAddr {
	if len(address) == 0 {
		return nil
	}
	switch address[0] {
	case 1:
		ip := make(net.IP, 4)
		copy(ip, address[1:5])
		return &net.UDPAddr{IP: ip, Port: int(binary.BigEndian.Uint16(address[5:7]))}
	case 4:
		ip := make(net.IP, 16)
		copy(ip, address[1:17])
		return &net.UDPAddr{IP: ip, Port: int(binary.BigEndian.Uint16(address[17:19]))}
	default:
		return nil
	}
}

// ReadFrom: 0x1b66620. Deliberately decodes the entire supplied buffer,
// rather than b[:n], and returns n minus the decoded header length.
func (c *hNLZlR1I) ReadFrom(b []byte) (int, net.Addr, error) {
	n, _, err := c.PacketConn.ReadFrom(b)
	if err != nil {
		return 0, nil, err
	}
	address, payload, err := S.CkUmXXHwOEH(b)
	if err != nil {
		return 0, nil, err
	}
	addr := wireUDPAddr(address)
	if addr == nil {
		return 0, nil, errors.New("parse udp addr error")
	}
	copy(b, payload)
	return n - len(address) - 3, addr, nil
}

// Close: 0x1b669a0. The TCP close error is discarded.
func (c *hNLZlR1I) Close() error {
	_ = c.f67_L17asEA.Close()
	return c.PacketConn.Close()
}

// g0YPF3: 0x1b6c4e0.
func g0YPF3(network, address string) (*net.UDPAddr, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}
	ip, err := R.TEx0aVr56JU(host)
	if err != nil {
		return nil, err
	}
	return net.ResolveUDPAddr(network, net.JoinHostPort(ip.String(), port))
}

// ListenPacketContext: 0x1b64fc0; cleanup goroutine: 0x1b73280. Metadata is
// unused. A failed TLS handshake is followed by SOCKS negotiation in this
// original method; its result replaces the earlier handshake error.
func (s *AwDE24aRWC) ListenPacketContext(ctx context.Context, _ *E.D0mLgQF, options ...O.PNSpXRb) (out E.C08JTZ2l, err error) {
	conn, err := O.I_y9nlK1q8(ctx, "tcp", s.Addr(), s.DialOptions(options)...)
	if err != nil {
		return nil, fmt.Errorf("%s connect error: %w", s.Addr(), err)
	}
	if s.m0JWRLb2CH {
		secured := tls.Client(conn, s.idTOPJOmPpB)
		tlsCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = secured.HandshakeContext(tlsCtx)
		conn = secured
	}
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()
	mS0F0Ti(conn)
	var auth *S.Auth
	if s.gmXALhCpyKF != "" {
		auth = &S.Auth{Username: s.gmXALhCpyKF, Password: s.eOL60dT1}
	}
	bound, err := S.AMwLSL6GewF(conn, S.Va09RL21(netip.IPv4Unspecified(), 0), 3, auth)
	if err != nil {
		return nil, fmt.Errorf("client hanshake error: %w", err)
	}
	packet, err := O.OpEY1B7hjf(ctx, "udp", "", s.DialOptions(options)...)
	if err != nil {
		return nil, err
	}
	go func() {
		_, _ = io.Copy(io.Discard, conn)
		_ = conn.Close()
		_ = packet.Close()
	}()
	relay := wireUDPAddr(bound)
	if relay == nil {
		return nil, errors.New("invalid UDP bind address")
	}
	if relay.IP.IsUnspecified() {
		resolved, resolveErr := g0YPF3("udp", s.Addr())
		if resolveErr != nil {
			return nil, resolveErr
		}
		relay.IP = resolved.IP
	}
	return &uMtOz6o{PacketConn: &hNLZlR1I{PacketConn: packet, iYgXoENCDQ5w: relay, f67_L17asEA: conn}, f7i_T5n5JXyM: E.GYe3KD{s.Name()}}, nil
}
