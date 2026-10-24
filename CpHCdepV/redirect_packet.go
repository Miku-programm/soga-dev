package CpHCdepV

import "net"

// CX1dpcZA6 is the RTTI 24-byte UDP redirect wrapper. PacketConn is the
// original embedded WCApOca6IQW7 type mapped back to net.PacketConn.
type CX1dpcZA6 struct {
	net.PacketConn
	lxRU7JIrLmL *net.UDPAddr
}

// NewRedirectPacketConn exposes the exact inline construction at 0x1b7fe0c.
// It is a source-organization helper, not an independently retained symbol.
func NewRedirectPacketConn(conn net.PacketConn, remote *net.UDPAddr) *CX1dpcZA6 {
	return &CX1dpcZA6{conn, remote}
}

// ReadFrom restores 0x18b9240: reported address is always the configured target,
// even when the underlying read returns an error.
func (c *CX1dpcZA6) ReadFrom(b []byte) (int, net.Addr, error) {
	n, _, err := c.PacketConn.ReadFrom(b)
	return n, c.lxRU7JIrLmL, err
}

// WriteTo restores 0x18b92c0, replacing the caller's destination.
func (c *CX1dpcZA6) WriteTo(b []byte, _ net.Addr) (int, error) {
	return c.PacketConn.WriteTo(b, c.lxRU7JIrLmL)
}
