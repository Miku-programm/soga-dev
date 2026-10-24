package vless

import (
	"crypto/cipher"
	"net"
)

// DrpXao is the original embedded net.Conn interface name.
type DrpXao = net.Conn

// LUKUT5W0R6 preserves the 112-byte TLS-header stream-cipher connection.
type LUKUT5W0R6 struct {
	DrpXao
	D1_SMs0WWKz  cipher.Stream
	EO5NXI6ex6   cipher.Stream
	IdEXNQK5b1k  int
	DhRIe1pBzWdS []byte
	LMy55MYg_l   int
	TfVMo2pTBnt  []byte
}

// Write recovers 0x1d012a0. It mutates the caller's header bytes, ignores
// header-validation errors and the underlying short-write count, as in the ELF.
func (c *LUKUT5W0R6) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	remaining := p
	for {
		if len(remaining) <= c.IdEXNQK5b1k {
			c.IdEXNQK5b1k -= len(remaining)
			break
		}
		remaining = remaining[c.IdEXNQK5b1k:]
		c.IdEXNQK5b1k = 0
		need := 5 - len(c.DhRIe1pBzWdS)
		if need > len(remaining) {
			c.DhRIe1pBzWdS = append(c.DhRIe1pBzWdS, remaining...)
			c.D1_SMs0WWKz.XORKeyStream(remaining, remaining)
			break
		}
		header := append(c.DhRIe1pBzWdS, remaining[:need]...)
		c.IdEXNQK5b1k, _ = QkDUgv(header)
		c.DhRIe1pBzWdS = c.DhRIe1pBzWdS[:0]
		c.D1_SMs0WWKz.XORKeyStream(remaining[:need], remaining[:need])
		remaining = remaining[need:]
	}
	_, err := c.DrpXao.Write(p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// Read recovers 0x1d015c0, processing returned bytes even alongside an error.
func (c *LUKUT5W0R6) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	n, err := c.DrpXao.Read(p)
	remaining := p[:n]
	for {
		if len(remaining) <= c.LMy55MYg_l {
			c.LMy55MYg_l -= len(remaining)
			break
		}
		remaining = remaining[c.LMy55MYg_l:]
		c.LMy55MYg_l = 0
		need := 5 - len(c.TfVMo2pTBnt)
		if need > len(remaining) {
			c.EO5NXI6ex6.XORKeyStream(remaining, remaining)
			c.TfVMo2pTBnt = append(c.TfVMo2pTBnt, remaining...)
			break
		}
		c.EO5NXI6ex6.XORKeyStream(remaining[:need], remaining[:need])
		header := append(c.TfVMo2pTBnt, remaining[:need]...)
		c.LMy55MYg_l, _ = QkDUgv(header)
		c.TfVMo2pTBnt = c.TfVMo2pTBnt[:0]
		remaining = remaining[need:]
	}
	return n, err
}
