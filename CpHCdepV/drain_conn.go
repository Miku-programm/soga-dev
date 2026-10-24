package CpHCdepV

import (
	"net"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/uLofn6j"
)

// QBM9_LSBqsu retains the original 24-byte wrapper and random drain budget.
type QBM9_LSBqsu struct {
	DrpXao
	ndn9Wlx int
}

func MlZo77Me(conn net.Conn) net.Conn { //0x18b8fe0
	a := uLofn6j.Co6D_zfls(3266)
	n := uLofn6j.Co6D_zfls(64) + 1
	b := 0
	if n != 1 {
		b = uLofn6j.Co6D_zfls(n)
	}
	return &QBM9_LSBqsu{DrpXao: conn, ndn9Wlx: a + b + 54}
}
func (c *QBM9_LSBqsu) Read(p []byte) (int, error) { //0x18b90c0
	n, err := c.DrpXao.Read(p)
	if err == nil && n > 0 && c.ndn9Wlx > 0 {
		c.ndn9Wlx -= n
	}
	return n, err
}
func (c *QBM9_LSBqsu) Free() { //0x18b9160
	if c.ndn9Wlx > 0 {
		_ = G6bHD5JA.KlTAqrF42mR(c.DrpXao, c.ndn9Wlx)
	}
	HxKViEp2a.Li3c4CJ1_D(c.DrpXao)
}
