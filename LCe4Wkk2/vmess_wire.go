package LCe4Wkk2

import (
	"errors"
	"net"
	E "recovered.local/soga/E3F12urwswL"
	V "recovered.local/soga/FLhkKwJ8M"
)

// eEgDc4: 0x1b6f120. VMess address tags differ from SOCKS tags.
func eEgDc4(m *E.D0mLgQF) *V.IxhMZ66CqoH1 {
	var tag byte
	var address []byte
	switch m.AddrType() {
	case 1:
		tag = 1
		address = make([]byte, 4)
		copy(address, m.ZouEpmjw8.To4())
	case 3:
		tag = 2
		address = make([]byte, len(m.WJ95LyOL)+1)
		address[0] = byte(len(m.WJ95LyOL))
		copy(address[1:], m.WJ95LyOL)
	case 4:
		tag = 3
		address = make([]byte, 16)
		copy(address, m.ZouEpmjw8.To16())
	}
	return &V.IxhMZ66CqoH1{GV3P5rY: m.AcXafZ1 == 1, OyEjzi1v93: tag, TIBiNxg: address, GHP6ivev2Db: uint(m.JRXMVvg)}
}

// Original init 0x1b55160 global; the name expresses its binary address.
var udpDestinationErrorAt5097920 = errors.New("udp packet dropped due to mismatched remote address")

type keoydRNdlJ struct {
	net.Conn
	oDRJBoi net.Addr
}

// WriteTo: 0x1b6f400. Both concrete address assertions and nil dereferences are
// retained. Zone is not compared by the original; only IP and port matter.
func (c *keoydRNdlJ) WriteTo(p []byte, address net.Addr) (int, error) {
	remote := c.oDRJBoi.(*net.UDPAddr)
	other := address.(*net.UDPAddr)
	if remote.IP.Equal(other.IP) && remote.Port == other.Port {
		return c.Conn.Write(p)
	}
	return 0, udpDestinationErrorAt5097920
}

// ReadFrom: 0x1b6f560; preserves the fixed address even when Read fails.
func (c *keoydRNdlJ) ReadFrom(p []byte) (int, net.Addr, error) {
	n, err := c.Conn.Read(p)
	return n, c.oDRJBoi, err
}
