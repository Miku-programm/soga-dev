// Package rqTdre6AyaDw restores the original native UDP-over-TCP adapter.
package rqTdre6AyaDw

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	B "recovered.local/soga/BaEWsPy"
	D "recovered.local/soga/D_zCag1qTV"
	G "recovered.local/soga/G6bHD5JA"
	J "recovered.local/soga/J1P4KNvfI_J"
	C "recovered.local/soga/cNT24rG4F"
	N "recovered.local/soga/nzATvajt"
	"time"
)

type DrpXao = net.Conn

// JAYwg0sLs preserves RTTI 0x370cc38, 64 bytes.
type JAYwg0sLs struct {
	DrpXao
	q9XNwb3     N.PgaNd3014U
	pAILOi      *B.D23y2888a_
	eStVCYS     bool
	bPg7GR8DCGh *G.AHU2MRmjzQu
	il9Qrj      bool
}

// MYA15vm4nVrn preserves RTTI 0x36314f0, 16 bytes.
type MYA15vm4nVrn struct {
	Hja8Qj bool
	N9uzgg *G.AHU2MRmjzQu
}

func Vd_ooUdqfJTg(host string) bool { //0x18ac640
	return host == "sp.v2.udp-over-tcp.arpa" || host == "sp.udp-over-tcp.arpa"
}
func Lfd9DDhfr(r io.Reader) (*MYA15vm4nVrn, error) { //0x18ac740
	result := new(MYA15vm4nVrn)
	if err := binary.Read(r, binary.BigEndian, &result.Hja8Qj); err != nil {
		return nil, err
	}
	var err error
	result.N9uzgg, err = G.SzhW4tOzUZg(r)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func SACWxFiDLx22(conn net.Conn, magic string, user N.PgaNd3014U, resolver *B.D23y2888a_) (J.Ki8urSHTx44, error) { //0x18a9660
	c := &JAYwg0sLs{DrpXao: conn, q9XNwb3: user, pAILOi: resolver}
	_ = conn.SetReadDeadline(time.Now().Add(time.Duration(C.UDPTimeoutMinutes()) * time.Minute))
	defer conn.SetReadDeadline(time.Time{})
	switch magic {
	case "sp.v2.udp-over-tcp.arpa":
		request, err := Lfd9DDhfr(conn)
		if err != nil {
			return nil, fmt.Errorf("error reading uot request: %v", err)
		}
		c.eStVCYS = request.Hja8Qj
		c.bPg7GR8DCGh = request.N9uzgg
	case "sp.udp-over-tcp.arpa":
		first, err := G.O2xtlb(conn)
		if err != nil {
			return nil, fmt.Errorf("error reading uot first packet addr: %v", err)
		}
		c.bPg7GR8DCGh = first
		c.il9Qrj = true
	default:
		return nil, fmt.Errorf("unknown uot magic address: %v", magic)
	}
	return c, nil
}
func (c *JAYwg0sLs) GetUser() N.PgaNd3014U       { return c.q9XNwb3 }     //0x18aa520
func (c *JAYwg0sLs) GetDestAddr() *G.AHU2MRmjzQu { return c.bPg7GR8DCGh } //0x18aa540
func (c *JAYwg0sLs) NetworkType() J.ZES6F5       { return 1 }             //0x18aa560

func (c *JAYwg0sLs) ReadFromProxy(buf []byte) (int, net.Addr, error) { //0x18a9e00
	target := c.bPg7GR8DCGh
	if !c.eStVCYS && !c.il9Qrj {
		var err error
		target, err = G.O2xtlb(c.DrpXao)
		if err != nil {
			return 0, nil, err
		}
	}
	if err := target.ResolveHost(c.pAILOi); err != nil {
		return 0, nil, err
	}
	addr, err := target.ToUDPAddr()
	if err != nil {
		return 0, nil, err
	}
	var length uint16
	if err = binary.Read(c.DrpXao, binary.BigEndian, &length); err != nil {
		return 0, nil, err
	}
	c.il9Qrj = false
	if len(buf) < int(length) {
		if _, err = io.ReadFull(c.DrpXao, buf); err != nil {
			return 0, nil, err
		}
		// The original intentionally discards the drain error after a full prefix.
		_ = G.KlTAqrF42mR(c.DrpXao, int(length)-len(buf))
		return len(buf), addr, nil
	}
	if _, err = io.ReadFull(c.DrpXao, buf[:length]); err != nil {
		return 0, nil, err
	}
	return int(length), addr, nil
}
func (c *JAYwg0sLs) WriteToProxy(payload []byte, addr net.Addr) (int, error) { //0x18aa220
	size := len(payload) + 2
	var target *G.AHU2MRmjzQu
	if !c.eStVCYS {
		var err error
		target, err = G.RNFJeaRY(addr)
		if err != nil {
			return 0, err
		}
		size += target.Size()
	}
	buf := D.S4OQXWIZ(size)
	defer buf.Free()
	if target != nil {
		if _, err := G.VCQ8zrwYebS(buf, target); err != nil {
			return 0, err
		}
	}
	buf.WriteUint16(uint16(len(payload)))
	_, _ = buf.Write(payload)
	if _, err := c.DrpXao.Write(buf.Bytes()); err != nil {
		return 0, err
	}
	// Both uint16 truncation for large payloads and nil-error short writes are
	// retained from the original, rather than imposing additional validation.
	return len(payload), nil
}
