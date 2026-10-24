package mieru

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	D "recovered.local/soga/buf/pool"
	G "recovered.local/soga/net/relay"
	J "recovered.local/soga/net/sess"
)

type h7SOFexGyw struct{ *xjTVP8 }

func (*h7SOFexGyw) NetworkType() J.ZES6F5 { return 1 } //0x1c41da0
func yT5ySB_MIS(addr net.Addr) (net.IP, int, error) { //0x1c42c40
	switch a := addr.(type) {
	case *net.TCPAddr:
		return a.IP, a.Port, nil
	case *net.UDPAddr:
		return a.IP, a.Port, nil
	default:
		return nil, 0, fmt.Errorf("unsupported addr type %T", addr)
	}
}
func rNsudCbv(b []byte) (*G.AHU2MRmjzQu, []byte, error) { //0x1c41dc0
	if len(b) < 6 {
		return nil, nil, fmt.Errorf("socks5 udp datagram too short: %d", len(b))
	}
	if b[0] != 0 || b[1] != 0 {
		return nil, nil, fmt.Errorf("socks5 udp datagram RSV not zero: %x %x", b[0], b[1])
	}
	if b[2] != 0 {
		return nil, nil, fmt.Errorf("socks5 udp fragmentation not supported (FRAG=%d)", b[2])
	}
	var ip net.IP
	var host string
	var end int
	switch b[3] {
	case 1:
		if len(b) < 10 {
			return nil, nil, fmt.Errorf("socks5 udp ipv4 short: %d", len(b))
		}
		ip = append(net.IP(nil), b[4:8]...)
		end = 10
	case 4:
		if len(b) < 22 {
			return nil, nil, fmt.Errorf("socks5 udp ipv6 short: %d", len(b))
		}
		ip = append(net.IP(nil), b[4:20]...)
		end = 22
	case 3:
		end = int(b[4]) + 7
		if len(b) < end {
			return nil, nil, fmt.Errorf("socks5 udp fqdn body short: need %d, have %d", end, len(b))
		}
		host = string(b[5 : end-2])
		ip = net.ParseIP(host)
	default:
		return nil, nil, fmt.Errorf("socks5 udp unsupported ATYP %d", b[3])
	}
	port := int(binary.BigEndian.Uint16(b[end-2 : end]))
	var addr *G.AHU2MRmjzQu
	if b[3] == 3 && ip == nil {
		addr = G.NewDomainAddress(host, port)
	} else {
		addr, _ = G.RNFJeaRY(&net.UDPAddr{IP: ip, Port: port})
	}
	return addr, b[end:], nil
}
func (h *h7SOFexGyw) ReadFromProxy(b []byte) (int, net.Addr, error) { //0x1c41120
	var prefix [3]byte
	if _, err := io.ReadFull(h.xjTVP8, prefix[:]); err != nil {
		return 0, nil, err
	}
	if prefix[0] != 0 {
		return 0, nil, fmt.Errorf("packet prefix 0x%x is not 0x00", prefix[0])
	}
	n := int(binary.BigEndian.Uint16(prefix[1:]))
	if n > len(b) {
		return 0, nil, fmt.Errorf("packet length %d exceeds buffer %d", n, len(b))
	}
	if _, err := io.ReadFull(h.xjTVP8, b[:n]); err != nil {
		return 0, nil, err
	}
	var suffix [1]byte
	if _, err := io.ReadFull(h.xjTVP8, suffix[:]); err != nil {
		return 0, nil, err
	}
	if suffix[0] != 255 {
		return 0, nil, fmt.Errorf("packet suffix 0x%x is not 0xFF", suffix[0])
	}
	a, payload, err := rNsudCbv(b[:n])
	if err != nil {
		return 0, nil, err
	}
	addr, err := a.ToUDPAddr()
	if err != nil {
		return 0, nil, err
	}
	return copy(b, payload), addr, nil
}
func (h *h7SOFexGyw) WriteToProxy(payload []byte, addr net.Addr) (int, error) { //0x1c41680
	ip, port, err := yT5ySB_MIS(addr)
	if err != nil {
		return 0, err
	}
	v4 := ip.To4()
	headerLen := 22
	if v4 != nil {
		headerLen = 10
	}
	length := headerLen + len(payload)
	b := D.PoolAt50961D8.Get(length + 4)
	defer D.PoolAt50961D8.Put(b)
	b[0] = 0
	binary.BigEndian.PutUint16(b[1:3], uint16(length))
	inner := b[3:]
	inner[0], inner[1], inner[2] = 0, 0, 0
	if v4 != nil {
		inner[3] = 1
		copy(inner[4:8], v4)
		binary.BigEndian.PutUint16(inner[8:10], uint16(port))
	} else {
		inner[3] = 4
		copy(inner[4:20], ip.To16())
		binary.BigEndian.PutUint16(inner[20:22], uint16(port))
	}
	copy(inner[headerLen:], payload)
	b[length+3] = 255
	if _, err = h.xjTVP8.Write(b); err != nil {
		return 0, err
	}
	return len(payload), nil
}
