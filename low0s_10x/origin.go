// Restored using target IDA functions and RTTI; reference github.com/Ieooo/clash/transport/ssr/protocol.
// GPL-3.0; exact original dependency revision remains unproven.
package low0s_10x

import (
	"bytes"
	"net"
)

type fy1trpRM struct{}

// init: 0x1b338e0.
func init() { register("origin", p3Mh1oCO, 0) }

// p3Mh1oCO: 0x1b3fca0.
func p3Mh1oCO(b *PZfVxx4UGRb) Shdj3DH1 { return &fy1trpRM{} }

// (*fy1trpRM).StreamConn: 0x1b3fcc0.
func (o *fy1trpRM) StreamConn(c net.Conn, iv []byte) net.Conn { return c }

// (*fy1trpRM).PacketConn: 0x1b3fce0.
func (o *fy1trpRM) PacketConn(c net.PacketConn) net.PacketConn { return c }

// (*fy1trpRM).Decode: 0x1b3fd00.
func (o *fy1trpRM) Decode(dst, src *bytes.Buffer) error {
	dst.ReadFrom(src)
	return nil
}

// (*fy1trpRM).Encode: 0x1b3fd60.
func (o *fy1trpRM) Encode(buf *bytes.Buffer, b []byte) error {
	buf.Write(b)
	return nil
}

// (*fy1trpRM).DecodePacket: 0x1b3fde0.
func (o *fy1trpRM) DecodePacket(b []byte) ([]byte, error) { return b, nil }

// (*fy1trpRM).EncodePacket: 0x1b3fe00.
func (o *fy1trpRM) EncodePacket(buf *bytes.Buffer, b []byte) error {
	buf.Write(b)
	return nil
}
