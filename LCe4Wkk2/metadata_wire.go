package LCe4Wkk2

import (
	"encoding/binary"
	E "recovered.local/soga/E3F12urwswL"
	B "recovered.local/soga/Hx04IbKk"
)

// o2HGjDHIVO: 0x1b6c140. Domain length wraps to a byte but the original
// appends the entire hostname; it does not reject names longer than 255 bytes.
func o2HGjDHIVO(m *E.D0mLgQF) []byte {
	b := B.BaDekw7qe{}
	kind := m.AddrType()
	b[b.Grow(1)] = byte(kind)
	switch kind {
	case 1:
		ip := m.ZouEpmjw8.To4()
		at := b.Grow(len(ip))
		copy(b[at:], ip)
	case 3:
		b[b.Grow(1)] = byte(len(m.WJ95LyOL))
		at := b.Grow(len(m.WJ95LyOL))
		copy(b[at:], m.WJ95LyOL)
	case 4:
		ip := m.ZouEpmjw8.To16()
		at := b.Grow(len(ip))
		copy(b[at:], ip)
	}
	at := b.Grow(2)
	binary.BigEndian.PutUint16(b[at:], uint16(m.JRXMVvg))
	return b
}
