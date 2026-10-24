package proxyproto

import (
	"encoding/binary"
	"fmt"
)

type NMEUUf9 byte
type M46tt8Zs struct {
	Be8_QISO NMEUUf9
	J8xIVW   []byte
}

func FUahHJe3a(raw []byte) ([]M46tt8Zs, error) { // 0x1185e60
	var records []M46tt8Zs
	for i := 0; i < len(raw); {
		kind := NMEUUf9(raw[i])
		if len(raw)-i <= 2 {
			return nil, ErrorAt5097F00
		}
		n := int(binary.BigEndian.Uint16(raw[i+1 : i+3]))
		i += 3
		if i+n > len(raw) {
			return nil, ErrorAt5097F00
		}
		record := M46tt8Zs{Be8_QISO: kind}
		if kind != 4 {
			record.J8xIVW = make([]byte, n)
			copy(record.J8xIVW, raw[i:i+n])
		}
		i += n
		records = append(records, record)
	}
	return records, nil
}
func I_BfofJ(records []M46tt8Zs) ([]byte, error) { // 0x11861c0
	var raw []byte
	for _, record := range records {
		n := len(record.J8xIVW)
		if n > 65535 {
			return nil, fmt.Errorf("proxyproto: cannot format TLV %v with length %d", record.Be8_QISO, n)
		}
		raw = append(raw, byte(record.Be8_QISO), byte(n>>8), byte(n))
		raw = append(raw, record.J8xIVW...)
	}
	return raw, nil
}
func (h *URdA97qPFc) TLVs() ([]M46tt8Zs, error) { return FUahHJe3a(h.ttF0mSAdaYG) } // 0x1176460
func (h *URdA97qPFc) SetTLVs(records []M46tt8Zs) error { // 0x11764a0
	raw, err := I_BfofJ(records)
	if err != nil {
		return err
	}
	h.ttF0mSAdaYG = raw
	return nil
}
func (p NMEUUf9) Registered() bool { return p >= 1 && p <= 5 || p >= 0x20 && p <= 0x25 || p == 0x30 }   // 0x1186580
func (p NMEUUf9) App() bool        { return p >= 0xe0 && p <= 0xef }                                    // 0x11865a0
func (p NMEUUf9) Experiment() bool { return p >= 0xf0 && p <= 0xf7 }                                    // 0x11865c0
func (p NMEUUf9) Future() bool     { return p >= 0xf8 }                                                 // 0x11865e0
func (p NMEUUf9) Spec() bool       { return p.Registered() || p.App() || p.Experiment() || p.Future() } // 0x1186600
