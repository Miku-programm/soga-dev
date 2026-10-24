//go:build linux

package LCe4Wkk2

import (
	"encoding/binary"
	"errors"
	"net"
	B "recovered.local/soga/Bx4XCl"
)

// dTAcwl_gA is the original 32-byte decrypted Shadowsocks packet wrapper.
type dTAcwl_gA struct {
	net.PacketConn
	zbHzsqAUlM net.Addr
}

// WriteTo: 0x1b5da20. Drop the three-byte SOCKS prefix; retain encoded count.
func (c *dTAcwl_gA) WriteTo(p []byte, address net.Addr) (int, error) {
	packet, err := B.XFq85XFER(B.KvyKMZ(address), p)
	if err != nil {
		return 0, err
	}
	return c.PacketConn.WriteTo(packet[3:], c.zbHzsqAUlM)
}

// ReadFrom: 0x1b5db20. Parse only the reported datagram, then shift the
// entire caller buffer. Domain headers are structurally parsed but rejected.
func (c *dTAcwl_gA) ReadFrom(p []byte) (int, net.Addr, error) {
	n, _, err := c.PacketConn.ReadFrom(p)
	if err != nil {
		return 0, nil, err
	}
	packet := p[:n]
	length := 0
	if len(packet) > 0 {
		switch packet[0] {
		case 1:
			length = 7
		case 4:
			length = 19
		case 3:
			if len(packet) > 1 {
				length = int(packet[1]) + 4
			}
		}
	}
	if length == 0 || len(packet) < length {
		return 0, nil, errors.New("parse addr error")
	}
	var ip net.IP
	switch packet[0] {
	case 1:
		ip = make(net.IP, 4)
		copy(ip, packet[1:5])
	case 4:
		ip = make(net.IP, 16)
		copy(ip, packet[1:17])
	default:
		return 0, nil, errors.New("parse addr error")
	}
	addr := &net.UDPAddr{IP: ip, Port: int(binary.BigEndian.Uint16(packet[length-2 : length]))}
	copy(p, p[length:])
	return n - length, addr, nil
}
