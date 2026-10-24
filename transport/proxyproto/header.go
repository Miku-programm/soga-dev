package proxyproto

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
	"strconv"
)

// X_WkJPwas4L @ 0x1175840。原程序优先检查任一端 To16；IPv4 也能成功 To16。
func X_WkJPwas4L(version byte, source, destination net.Addr) *URdA97qPFc {
	if version < 1 || version > 2 {
		version = 2
	}
	h := &URdA97qPFc{TmZyJAm: version, LgI8Cr1feIYb: 0x20}
	switch src := source.(type) {
	case *net.TCPAddr:
		if dst, ok := destination.(*net.TCPAddr); ok {
			if len(src.IP.To16()) == 16 || len(dst.IP.To16()) == 16 {
				h.UCYj4Sj65ui = 0x21
			} else if len(src.IP.To4()) == 4 {
				h.UCYj4Sj65ui = 0x11
			} else if len(src.IP) == 16 {
				h.UCYj4Sj65ui = 0x21
			}
		}
	case *net.UDPAddr:
		if dst, ok := destination.(*net.UDPAddr); ok {
			if len(src.IP.To16()) == 16 || len(dst.IP.To16()) == 16 {
				h.UCYj4Sj65ui = 0x22
			} else if len(src.IP.To4()) == 4 {
				h.UCYj4Sj65ui = 0x12
			} else if len(src.IP) == 16 {
				h.UCYj4Sj65ui = 0x22
			}
		}
	case *net.UnixAddr:
		if _, ok := destination.(*net.UnixAddr); ok {
			switch src.Net {
			case "unix":
				h.UCYj4Sj65ui = 0x31
			case "unixgram":
				h.UCYj4Sj65ui = 0x32
			}
		}
	}
	if h.UCYj4Sj65ui != 0 {
		h.LgI8Cr1feIYb = 0x21
		h.Dafxa_BOdZdw = source
		h.KhtRcTD89 = destination
	}
	return h
}
func (h *URdA97qPFc) TCPAddrs() (*net.TCPAddr, *net.TCPAddr, bool) { // 0x1175ee0
	if !h.UCYj4Sj65ui.IsStream() {
		return nil, nil, false
	}
	a, ok := h.Dafxa_BOdZdw.(*net.TCPAddr)
	b, ok2 := h.KhtRcTD89.(*net.TCPAddr)
	return a, b, ok && ok2
}
func (h *URdA97qPFc) UDPAddrs() (*net.UDPAddr, *net.UDPAddr, bool) { // 0x1175f40
	if !h.UCYj4Sj65ui.IsDatagram() {
		return nil, nil, false
	}
	a, ok := h.Dafxa_BOdZdw.(*net.UDPAddr)
	b, ok2 := h.KhtRcTD89.(*net.UDPAddr)
	return a, b, ok && ok2
}
func (h *URdA97qPFc) UnixAddrs() (*net.UnixAddr, *net.UnixAddr, bool) { // 0x1175fa0
	if !h.UCYj4Sj65ui.IsUnix() {
		return nil, nil, false
	}
	a, ok := h.Dafxa_BOdZdw.(*net.UnixAddr)
	b, ok2 := h.KhtRcTD89.(*net.UnixAddr)
	return a, b, ok && ok2
}
func (h *URdA97qPFc) IPs() (net.IP, net.IP, bool) { // 0x1176000
	if a, b, ok := h.TCPAddrs(); ok {
		return a.IP, b.IP, true
	}
	if a, b, ok := h.UDPAddrs(); ok {
		return a.IP, b.IP, true
	}
	return nil, nil, false
}
func (h *URdA97qPFc) Ports() (int, int, bool) { // 0x11760e0
	if a, b, ok := h.TCPAddrs(); ok {
		return a.Port, b.Port, true
	}
	if a, b, ok := h.UDPAddrs(); ok {
		return a.Port, b.Port, true
	}
	return 0, 0, false
}
func (h *URdA97qPFc) EqualTo(other *URdA97qPFc) bool { return h.EqualsTo(other) } // 0x11761a0
func (h *URdA97qPFc) EqualsTo(other *URdA97qPFc) bool { // 0x11761e0
	if other == nil {
		return false
	}
	if h.TmZyJAm == 2 && !bytes.Equal(h.ttF0mSAdaYG, other.ttF0mSAdaYG) {
		return false
	}
	if h.TmZyJAm != other.TmZyJAm || h.LgI8Cr1feIYb != other.LgI8Cr1feIYb || h.UCYj4Sj65ui != other.UCYj4Sj65ui {
		return false
	}
	if h.LgI8Cr1feIYb.IsLocal() {
		return true
	}
	return h.Dafxa_BOdZdw.String() == other.Dafxa_BOdZdw.String() && h.KhtRcTD89.String() == other.KhtRcTD89.String()
}
func (h *URdA97qPFc) WriteTo(w io.Writer) (int64, error) { // 0x1176340
	b, err := h.Format()
	if err != nil {
		return 0, err
	}
	return bytes.NewBuffer(b).WriteTo(w)
}
func (h *URdA97qPFc) Format() ([]byte, error) { // 0x11763e0
	switch h.TmZyJAm {
	case 1:
		return h.bLutKOp()
	case 2:
		return h.b3Incwjd()
	default:
		return nil, ErrorAt5097E90
	}
}
func (h *URdA97qPFc) bLutKOp() ([]byte, error) { // 0x1188080
	var protocol string
	switch h.UCYj4Sj65ui {
	case 0x11:
		protocol = "TCP4"
	case 0x21:
		protocol = "TCP6"
	default:
		return []byte("PROXY UNKNOWN\r\n"), nil
	}
	a, aok := h.Dafxa_BOdZdw.(*net.TCPAddr)
	b, bok := h.KhtRcTD89.(*net.TCPAddr)
	if !aok || !bok {
		return nil, ErrorAt5097ED0
	}
	src, dst := a.IP, b.IP
	if h.UCYj4Sj65ui == 0x11 {
		src = src.To4()
		dst = dst.To4()
	} else {
		src = src.To16()
		dst = dst.To16()
	}
	if src == nil || dst == nil {
		return nil, ErrorAt5097ED0
	}
	buf := bytes.NewBuffer(make([]byte, 0, 108))
	buf.Write(signatureAt5034690)
	buf.WriteByte(' ')
	buf.WriteString(protocol)
	buf.WriteByte(' ')
	buf.WriteString(src.String())
	buf.WriteByte(' ')
	buf.WriteString(dst.String())
	buf.WriteByte(' ')
	buf.WriteString(strconv.Itoa(a.Port))
	buf.WriteByte(' ')
	buf.WriteString(strconv.Itoa(b.Port))
	buf.WriteString("\r\n")
	return buf.Bytes(), nil
}
func (h *URdA97qPFc) b3Incwjd() ([]byte, error) { // 0x1189720
	var buf bytes.Buffer
	buf.Write(signatureAt5099D50)
	command := byte(0x20)
	if h.LgI8Cr1feIYb.IsProxy() {
		command = 0x21
	}
	buf.WriteByte(command)
	buf.WriteByte(h.UCYj4Sj65ui.uGOaaar())
	writeLength := func(base int) error {
		n := base + len(h.ttF0mSAdaYG)
		if n >= 65536 {
			return ErrorAt5097F30
		}
		buf.Write([]byte{byte(n >> 8), byte(n)})
		return nil
	}
	p := h.UCYj4Sj65ui
	if p.IsUnspec() {
		if err := writeLength(0); err != nil {
			return nil, err
		}
	} else {
		var src, dst []byte
		switch {
		case p.IsIPv4():
			if err := writeLength(12); err != nil {
				return nil, err
			}
			a, b, _ := h.IPs()
			src = a.To4()
			dst = b.To4()
		case p.IsIPv6():
			if err := writeLength(36); err != nil {
				return nil, err
			}
			a, b, _ := h.IPs()
			src = a.To16()
			dst = b.To16()
		case p.IsUnix():
			buf.Write([]byte{0, 216})
			a, b, ok := h.UnixAddrs()
			if !ok {
				return nil, ErrorAt5097ED0
			}
			src = make([]byte, 108)
			dst = make([]byte, 108)
			copy(src, a.Name)
			copy(dst, b.Name)
		}
		if src == nil || dst == nil {
			return nil, ErrorAt5097ED0
		}
		buf.Write(src)
		buf.Write(dst)
		if a, b, ok := h.Ports(); ok {
			var raw [2]byte
			binary.BigEndian.PutUint16(raw[:], uint16(a))
			buf.Write(raw[:])
			binary.BigEndian.PutUint16(raw[:], uint16(b))
			buf.Write(raw[:])
		}
	}
	if len(h.ttF0mSAdaYG) > 0 {
		buf.Write(h.ttF0mSAdaYG)
	}
	return buf.Bytes(), nil
}
