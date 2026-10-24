package RoaZP1oJy

import (
	"io"
	"net"
	T "recovered.local/soga/TaFdeUcnzQ"
	S "recovered.local/soga/ad5KEbw_EzWk"
)

type WCApOca6IQW7 = net.PacketConn

// Vzuv_s9ZS5 is the exact fifteen-method native protocol interface,
// recovered from RTTI 0x36d1930. Free is not an interface method.
type Vzuv_s9ZS5 interface {
	io.ReadWriter
	net.PacketConn
	SetServerInfo(*S.L7E7H2aW8KVs)
	ReadFromUID([]byte) (int, int64, net.Addr, error)
	WriteToUID([]byte, net.Addr, int64) (int, error)
	CurUID() int64
	CID() int64
	GetOverhead() int
}
type ojt4CWtzUu func(*T.KrOWYtzUYft) Vzuv_s9ZS5

// RTTI 0x36f6f80, 32 bytes. Origin passes through to the original cipher.
type f_E7c41 struct {
	io.ReadWriter
	WCApOca6IQW7
}

func jyvuea4xsHn(c *T.KrOWYtzUYft) Vzuv_s9ZS5  { return &f_E7c41{c, c} } //0x1cdd2e0
func (*f_E7c41) SetServerInfo(*S.L7E7H2aW8KVs) {}                        //0x1cdd380
func (c *f_E7c41) ReadFromUID(p []byte) (int, int64, net.Addr, error) {
	n, a, e := c.WCApOca6IQW7.ReadFrom(p)
	return n, 0, a, e
} //0x1cdd3a0
func (c *f_E7c41) WriteToUID(p []byte, a net.Addr, _ int64) (int, error) {
	return c.WCApOca6IQW7.WriteTo(p, a)
}                                 //0x1cdd420
func (*f_E7c41) CurUID() int64    { return 0 } //0x1cdd4a0
func (*f_E7c41) CID() int64       { return 0 } //0x1cdd4c0
func (*f_E7c41) GetOverhead() int { return 0 } //0x1cdd4e0
