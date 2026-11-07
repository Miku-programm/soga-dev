package mieru

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	G "recovered.local/soga/net/relay"
	J "recovered.local/soga/net/sess"
)

// dOUoamsBibCa @0x1c58520. The reserved request byte is intentionally not
// checked. A zero-length FQDN follows the sample's empty-IP construction.
func dOUoamsBibCa(r io.Reader) (byte, *G.AHU2MRmjzQu, error) {
	var h [4]byte
	if _, err := io.ReadFull(r, h[:]); err != nil {
		return 0, nil, fmt.Errorf("read socks5 head: %w", err)
	}
	if h[0] != 5 {
		return 0, nil, fmt.Errorf("unsupported socks version: %d", h[0])
	}
	var ip net.IP
	var host string
	switch h[3] {
	case 1, 4:
		n, label := 4, "ipv4"
		if h[3] == 4 {
			n, label = 16, "ipv6"
		}
		ip = make(net.IP, n)
		if _, err := io.ReadFull(r, ip); err != nil {
			return 0, nil, fmt.Errorf("read %s: %w", label, err)
		}
	case 3:
		var n [1]byte
		if _, err := io.ReadFull(r, n[:]); err != nil {
			return 0, nil, fmt.Errorf("read fqdn len: %w", err)
		}
		b := make([]byte, int(n[0]))
		if _, err := io.ReadFull(r, b); err != nil {
			return 0, nil, fmt.Errorf("read fqdn: %w", err)
		}
		host = string(b)
	default:
		return h[1], nil, fmt.Errorf("unsupported atyp: %d", h[3])
	}
	var p [2]byte
	if _, err := io.ReadFull(r, p[:]); err != nil {
		return 0, nil, fmt.Errorf("read port: %w", err)
	}
	port := int(binary.BigEndian.Uint16(p[:]))
	if len(host) > 0 {
		if ip = net.ParseIP(host); ip == nil {
			return h[1], G.NewDomainAddress(host, port), nil
		}
	}
	a, _ := G.RNFJeaRY(&net.UDPAddr{IP: ip, Port: port})
	return h[1], a, nil
}
func (t *TZ3YCqW) Accept() (conn J.Ki8urSHTx44, err error) { //0x1c364c0; cleanup0x1c5c320
	defer func() {
		if err != nil {
			t.j_Ra2Blepbh.Close()
			t.j_Ra2Blepbh.Free()
		}
	}()
	cmd, addr, err := dOUoamsBibCa(t.j_Ra2Blepbh)
	if err != nil {
		return nil, fmt.Errorf("read socks5 request: %w", err)
	}
	switch cmd {
	case 1:
		t.j_Ra2Blepbh.wsyay7SHLui = addr
		t.j_Ra2Blepbh.wP6Ql2PLz = 0
		if err = t.hTInzJrFcBf(0); err != nil {
			return nil, err
		}
		return t.j_Ra2Blepbh, nil
	case 3:
		t.j_Ra2Blepbh.wsyay7SHLui = addr
		t.j_Ra2Blepbh.wP6Ql2PLz = 1
		if err = t.hTInzJrFcBf(0); err != nil {
			return nil, err
		}
		return &h7SOFexGyw{t.j_Ra2Blepbh}, nil
	default:
		t.hTInzJrFcBf(7)
		return nil, fmt.Errorf("unsupported socks5 cmd: %d", cmd)
	}
}
func (t *TZ3YCqW) hTInzJrFcBf(status byte) error { //0x1c369c0
	b := [10]byte{5, status, 0, 1}
	s := t.j_Ra2Blepbh
	if err := s.waitOpened(); err != nil {
		return err
	}
	return s.kmcW3r.WriteDataSegment(&IqB1foHWg{Rf6N7kY6GVCs: 7, HNaaDh: s.xIwb7TaM, Tt1OfazA: s.qYZ2Uu.Add(1) - 1, UlMT6MeaV2: 4096}, b[:])
}
