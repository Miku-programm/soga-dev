package proxyproto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/netip"
	"strconv"
	"strings"

	"recovered.local/soga/buf/pool"
)

var signatureAt5034690 = []byte("PROXY")
var signatureAt5099D50 = []byte("\r\n\r\n\x00\r\nQUIT\n")

// LavOMmUY2 @ 0x1176540。检测失败时保留缓冲内的应用数据。
func LavOMmUY2(r *pool.NADyavA) (*URdA97qPFc, error) {
	first, err := r.Peek(1)
	if err != nil {
		if err == io.EOF {
			return nil, ErrorAt5097E80
		}
		return nil, err
	}
	if first[0] == signatureAt5034690[0] || first[0] == signatureAt5099D50[0] {
		v1, err := r.Peek(5)
		if err != nil {
			if err == io.EOF {
				return nil, ErrorAt5097E80
			}
			return nil, err
		}
		if bytes.Equal(v1[:5], signatureAt5034690) {
			return _irOIZr(r)
		}
		v2, err := r.Peek(12)
		if err != nil {
			if err == io.EOF {
				return nil, ErrorAt5097E80
			}
			return nil, err
		}
		if bytes.Equal(v2[:12], signatureAt5099D50) {
			return yQ5Eq7R_(r)
		}
	}
	return nil, ErrorAt5097E80
}

func _irOIZr(r *pool.NADyavA) (*URdA97qPFc, error) { // 0x11879e0
	buf := make([]byte, 0, 107)
	for {
		b, err := r.ReadByte()
		if err != nil {
			return nil, fmt.Errorf(ErrorAt5097E00.Error()+": %v", err)
		}
		buf = append(buf, b)
		if b == '\n' {
			break
		}
		if len(buf) == 107 {
			return nil, ErrorAt5097E10
		}
		if r.Buffered() == 0 {
			return nil, ErrorAt5097E00
		}
	}
	if len(buf) < 2 || buf[len(buf)-2] != '\r' {
		return nil, ErrorAt5097E20
	}
	tokens := strings.Split(string(buf[:len(buf)-2]), " ")
	if len(tokens) < 2 {
		return nil, ErrorAt5097E40
	}
	var p JvCeJK0s
	switch tokens[1] {
	case "TCP4":
		p = 0x11
	case "TCP6":
		p = 0x21
	case "UNKNOWN":
	default:
		return nil, ErrorAt5097E40
	}
	if p != 0 && len(tokens) < 6 {
		return nil, ErrorAt5097E40
	}
	h := &URdA97qPFc{TmZyJAm: 1, LgI8Cr1feIYb: 0x21, UCYj4Sj65ui: p}
	if p == 0 {
		h.LgI8Cr1feIYb = 0x20
		return h, nil
	}
	src, err := eERftdt549X(p, tokens[2])
	if err != nil {
		return nil, err
	}
	dst, err := eERftdt549X(p, tokens[3])
	if err != nil {
		return nil, err
	}
	sp, err := dU5ZQka(tokens[4])
	if err != nil {
		return nil, err
	}
	dp, err := dU5ZQka(tokens[5])
	if err != nil {
		return nil, err
	}
	h.Dafxa_BOdZdw = &net.TCPAddr{IP: src, Port: sp}
	h.KhtRcTD89 = &net.TCPAddr{IP: dst, Port: dp}
	return h, nil
}
func dU5ZQka(s string) (int, error) { // 0x1188640
	p, err := strconv.Atoi(s)
	if err != nil || p < 0 || p > 65535 {
		return 0, ErrorAt5097EE0
	}
	return p, nil
}
func eERftdt549X(p JvCeJK0s, s string) (net.IP, error) { // 0x11886c0
	a, err := netip.ParseAddr(s)
	if err != nil {
		return nil, ErrorAt5097ED0
	}
	if p == 0x11 && a.Is4() || p == 0x21 && (a.Is6() || a.Is4In6()) {
		return net.IP(a.AsSlice()), nil
	}
	return nil, ErrorAt5097ED0
}

type uaDQzIVULF struct {
	FAesY38     [4]byte
	ZwfrNJlamEa [4]byte
	HUYQaAKRw   uint16
	KNT5RQVQ1W  uint16
}
type hkugEWyyFNRP struct {
	IHhM0Y   uint16
	DTxky60f uint16
}
type ds1wekpD7D struct {
	TF4La_u8w [16]byte
	TkrEYa    [16]byte
	hkugEWyyFNRP
}
type tfN60roQHar1 struct {
	HCgc5i      [108]byte
	A0rhdeBStgZ [108]byte
}

// yQ5Eq7R_ @ 0x1188a60。原程序先 Peek 整个 payload，容量不足也报长度错误。
func yQ5Eq7R_(r *pool.NADyavA) (*URdA97qPFc, error) {
	for i := 0; i < 12; i++ {
		if _, err := r.ReadByte(); err != nil {
			return nil, ErrorAt5097E30
		}
	}
	h := &URdA97qPFc{TmZyJAm: 2}
	b, err := r.ReadByte()
	if err != nil {
		return nil, ErrorAt5097E30
	}
	h.LgI8Cr1feIYb = ZYgyOe(b)
	if h.LgI8Cr1feIYb != 0x20 && h.LgI8Cr1feIYb != 0x21 {
		return nil, ErrorAt5097EA0
	}
	b, err = r.ReadByte()
	if err != nil {
		return nil, ErrorAt5097E40
	}
	h.UCYj4Sj65ui = JvCeJK0s(b)
	if b == 0 && !h.LgI8Cr1feIYb.IsLocal() {
		return nil, ErrorAt5097EB0
	}
	var length uint16
	if err := binary.Read(io.LimitReader(r, 2), binary.BigEndian, &length); err != nil {
		return nil, ErrorAt5097E50
	}
	p := h.UCYj4Sj65ui
	minimum := uint16(0)
	switch {
	case p.IsIPv4():
		minimum = 12
	case p.IsIPv6():
		minimum = 36
	case p.IsUnix():
		minimum = 216
	case !p.IsUnspec():
		return nil, ErrorAt5097EC0
	}
	if length < minimum {
		return nil, ErrorAt5097EC0
	}
	if length == 0 {
		return h, nil
	}
	if _, err := r.Peek(int(length)); err != nil {
		return nil, ErrorAt5097EC0
	}
	payload := &io.LimitedReader{R: r, N: int64(length)}
	// 下面两个地址转换分支对应原函数中的内联 net.TCPAddr/net.UDPAddr 构造。
	addr := func(ip net.IP, port uint16) net.Addr {
		if p.IsStream() {
			return &net.TCPAddr{IP: ip, Port: int(port)}
		}
		if p.IsDatagram() {
			return &net.UDPAddr{IP: ip, Port: int(port)}
		}
		return nil
	}
	if p != 0 {
		switch {
		case p.IsIPv4():
			a := new(uaDQzIVULF)
			if binary.Read(payload, binary.BigEndian, a) != nil {
				return nil, ErrorAt5097ED0
			}
			h.Dafxa_BOdZdw = addr(a.FAesY38[:], a.HUYQaAKRw)
			h.KhtRcTD89 = addr(a.ZwfrNJlamEa[:], a.KNT5RQVQ1W)
		case p.IsIPv6():
			a := new(ds1wekpD7D)
			if binary.Read(payload, binary.BigEndian, a) != nil {
				return nil, ErrorAt5097ED0
			}
			h.Dafxa_BOdZdw = addr(a.TF4La_u8w[:], a.IHhM0Y)
			h.KhtRcTD89 = addr(a.TkrEYa[:], a.DTxky60f)
		case p.IsUnix():
			var a tfN60roQHar1
			if binary.Read(payload, binary.BigEndian, &a) != nil {
				return nil, ErrorAt5097ED0
			}
			network := "unix"
			if p.IsDatagram() {
				network = "unixgram"
			}
			h.Dafxa_BOdZdw = &net.UnixAddr{Net: network, Name: mLIKJLb(a.HCgc5i[:])}
			h.KhtRcTD89 = &net.UnixAddr{Net: network, Name: mLIKJLb(a.A0rhdeBStgZ[:])}
		}
	}
	h.ttF0mSAdaYG = make([]byte, payload.N)
	if _, err := io.ReadFull(payload, h.ttF0mSAdaYG); err != nil && err != io.EOF {
		return nil, err
	}
	return h, nil
}
func mLIKJLb(b []byte) string { // 0x118a180，首个 NUL 截断，包括开头 NUL。
	if i := bytes.IndexByte(b, 0); i >= 0 {
		return string(b[:i])
	}
	return string(b)
}
