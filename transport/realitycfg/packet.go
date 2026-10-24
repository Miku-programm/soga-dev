// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package realitycfg

import (
	"bytes"
	"net"

	"recovered.local/soga/buf/syncpool"
)

// Preserve the original embedded field name, so the protocol PacketConn method
// remains promoted instead of colliding with the standard-library field name.
type WCApOca6IQW7 = net.PacketConn

type S_N0Ixh struct {
	WCApOca6IQW7
	Shdj3DH1
}

// (*S_N0Ixh).WriteTo: 0x1b3fe80.
func (c *S_N0Ixh) WriteTo(b []byte, addr net.Addr) (int, error) {
	buf := syncpool.BytesBufferPoolAt5044220.Get().(*bytes.Buffer)
	defer syncpool.LjZLCz9(buf)
	err := c.EncodePacket(buf, b)
	if err != nil {
		return 0, err
	}
	_, err = c.WCApOca6IQW7.WriteTo(buf.Bytes(), addr)
	return len(b), err
}

// (*S_N0Ixh).ReadFrom: 0x1b400a0.
func (c *S_N0Ixh) ReadFrom(b []byte) (int, net.Addr, error) {
	n, addr, err := c.WCApOca6IQW7.ReadFrom(b)
	if err != nil {
		return n, addr, err
	}
	wt_fjza, err := c.DecodePacket(b[:n])
	if err != nil {
		return n, addr, err
	}
	copy(b, wt_fjza)
	return len(wt_fjza), addr, nil
}
