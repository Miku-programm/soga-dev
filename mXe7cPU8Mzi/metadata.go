package mXe7cPU8Mzi

import (
	"encoding/binary"
	"fmt"
	"time"
)

// xofvu_nD97c and IqB1foHWg retain the original byte enum and 32-byte
// Go layout (RTTI 0x35e1cf0 and 0x370ad28). The wire layout is different.
type xofvu_nD97c uint8

type IqB1foHWg struct {
	Rf6N7kY6GVCs xofvu_nD97c
	WyK3yKmJKe1V uint32
	HNaaDh       uint32
	Tt1OfazA     uint32
	DGMUmJ7      uint8
	N01v7C       uint32
	UlMT6MeaV2   uint16
	VGN7gjQ4p    uint8
	B8QsA0N2yo   uint8
	HPUJVQ0S     uint16
	FO7MEJ       uint8
}

// Marshal is 0x1c3a1a0. Reserved bytes are zero. Unknown enum values still
// serialize the common prefix, exactly as the original writer does.
func (m *IqB1foHWg) Marshal() (b [32]byte) {
	b[0] = byte(m.Rf6N7kY6GVCs)
	binary.BigEndian.PutUint32(b[2:6], m.WyK3yKmJKe1V)
	binary.BigEndian.PutUint32(b[6:10], m.HNaaDh)
	binary.BigEndian.PutUint32(b[10:14], m.Tt1OfazA)
	switch m.Rf6N7kY6GVCs | 1 {
	case 3, 5:
		b[14] = m.DGMUmJ7
		binary.BigEndian.PutUint16(b[15:17], m.HPUJVQ0S)
		b[17] = m.FO7MEJ
	case 7, 9:
		binary.BigEndian.PutUint32(b[14:18], m.N01v7C)
		binary.BigEndian.PutUint16(b[18:20], m.UlMT6MeaV2)
		b[20], b[21] = m.VGN7gjQ4p, m.B8QsA0N2yo
		binary.BigEndian.PutUint16(b[22:24], m.HPUJVQ0S)
		b[24] = m.FO7MEJ
	}
	return
}

// Ox0UfI5kIfcR is 0x1c39bc0. Length and timestamp validation precede
// protocol validation. Time bounds and multiplication retain uint32 overflow.
func Ox0UfI5kIfcR(b []byte) (*IqB1foHWg, error) {
	if len(b) != 32 {
		return nil, fmt.Errorf("metadata length %d, want %d", len(b), 32)
	}
	return decodeMetadataAt(b, time.Now())
}

// The explicit clock is an analyst extraction of the original time.Now call;
// it permits deterministic checks of the original minute-window arithmetic.
func decodeMetadataAt(b []byte, now time.Time) (*IqB1foHWg, error) {
	if len(b) != 32 {
		return nil, fmt.Errorf("metadata length %d, want %d", len(b), 32)
	}
	stamp := binary.BigEndian.Uint32(b[2:6])
	minute := uint32(now.Unix() / 60)
	if stamp < minute-1 || stamp > minute+1 {
		return nil, fmt.Errorf("invalid timestamp %d", stamp*60)
	}
	m := &IqB1foHWg{Rf6N7kY6GVCs: xofvu_nD97c(b[0]), WyK3yKmJKe1V: stamp, HNaaDh: binary.BigEndian.Uint32(b[6:10]), Tt1OfazA: binary.BigEndian.Uint32(b[10:14])}
	switch m.Rf6N7kY6GVCs | 1 {
	case 3, 5:
		m.DGMUmJ7 = b[14]
		m.HPUJVQ0S = binary.BigEndian.Uint16(b[15:17])
		m.FO7MEJ = b[17]
		if m.HPUJVQ0S > 1024 {
			return nil, fmt.Errorf("session payload size %d exceeds maximum %d", m.HPUJVQ0S, 1024)
		}
	case 7, 9:
		m.N01v7C = binary.BigEndian.Uint32(b[14:18])
		m.UlMT6MeaV2 = binary.BigEndian.Uint16(b[18:20])
		m.VGN7gjQ4p = b[20]
		m.B8QsA0N2yo = b[21]
		m.HPUJVQ0S = binary.BigEndian.Uint16(b[22:24])
		m.FO7MEJ = b[24]
	default:
		return nil, fmt.Errorf("unknown protocol type %d", b[0])
	}
	return m, nil
}
