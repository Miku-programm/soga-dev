package Bx4XCl

import (
	"encoding/binary"
	"errors"
	"net"
	"net/netip"
	H "recovered.local/soga/Hx04IbKk"
	"strconv"
)

// L4JRaRodkw restores 0x192c7e0. Invalid host/port syntax, oversized domain
// names and non-uint16 ports return nil, without a separate error.
func L4JRaRodkw(address string) []byte {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil
	}
	b := H.BaDekw7qe{}
	ip := net.ParseIP(host)
	if ip == nil {
		if len(host) > 255 {
			return nil
		}
		iByte := b.Grow(1)
		b[iByte] = 3
		iByte = b.Grow(1)
		b[iByte] = byte(len(host))
		i := b.Grow(len(host))
		copy(b[i:], host)
	} else if v4 := ip.To4(); v4 != nil {
		iByte := b.Grow(1)
		b[iByte] = 1
		i := b.Grow(len(v4))
		copy(b[i:], v4)
	} else {
		iByte := b.Grow(1)
		b[iByte] = 4
		i := b.Grow(len(ip))
		copy(b[i:], ip)
	}
	p, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return nil
	}
	i := b.Grow(2)
	binary.BigEndian.PutUint16(b[i:], uint16(p))
	return b
}

// KvyKMZ: 0x192cc00. The original checks DefaultMask, not merely To4.
// Non-nil malformed IP slices are retained in the type-4 branch.
func KvyKMZ(addr net.Addr) []byte {
	var ip net.IP
	var port int
	switch a := addr.(type) {
	case *net.UDPAddr:
		ip, port = a.IP, a.Port
	case *net.TCPAddr:
		ip, port = a.IP, a.Port
	}
	if ip == nil {
		return L4JRaRodkw(addr.String())
	}
	var b H.BaDekw7qe
	if v4 := ip.To4(); v4.DefaultMask() != nil {
		b = make([]byte, 0, 7)
		iByte := b.Grow(1)
		b[iByte] = 1
		i := b.Grow(len(v4))
		copy(b[i:], v4)
	} else {
		b = make([]byte, 0, 19)
		iByte := b.Grow(1)
		b[iByte] = 4
		i := b.Grow(len(ip))
		copy(b[i:], ip)
	}
	i := b.Grow(2)
	binary.BigEndian.PutUint16(b[i:], uint16(port))
	return b
}

// Va09RL21: 0x192cf60. Invalid and IPv4-mapped IPv6 addresses use type 4.
func Va09RL21(addr netip.Addr, port uint16) []byte {
	if addr.Is4() {
		b := make([]byte, 7)
		b[0] = 1
		ip := addr.As4()
		copy(b[1:5], ip[:])
		binary.BigEndian.PutUint16(b[5:], port)
		return b
	}
	b := make([]byte, 19)
	b[0] = 4
	copy(b[1:17], addr.AsSlice())
	binary.BigEndian.PutUint16(b[17:], port)
	return b
}

// CkUmXXHwOEH: 0x192d100. Returned slices alias packet. On an invalid
// address the remainder after the three-byte prefix is still returned.
func CkUmXXHwOEH(packet []byte) (address, payload []byte, err error) {
	if len(packet) < 5 {
		return nil, nil, errors.New("insufficient length of packet")
	}
	if packet[0] != 0 || packet[1] != 0 {
		return nil, nil, errors.New("reserved fields should be zero")
	}
	if packet[2] != 0 {
		return nil, nil, errors.New("discarding fragmented payload")
	}
	packet = packet[3:]
	n := 0
	switch packet[0] {
	case 1:
		n = 7
	case 4:
		n = 19
	case 3:
		if len(packet) >= 2 {
			n = int(packet[1]) + 4
		}
	}
	if n == 0 || len(packet) < n {
		return nil, packet, errors.New("failed to read UDP header")
	}
	return packet[:n], packet[n:], nil
}

// XFq85XFER: 0x192d580. A non-nil empty address is accepted.
func XFq85XFER(address, payload []byte) ([]byte, error) {
	if address == nil {
		return nil, errors.New("address is invalid")
	}
	b := H.BaDekw7qe{}
	i := b.Grow(3)
	copy(b[i:], []byte{0, 0, 0})
	i = b.Grow(len(address))
	copy(b[i:], address)
	i = b.Grow(len(payload))
	copy(b[i:], payload)
	return b, nil
}
