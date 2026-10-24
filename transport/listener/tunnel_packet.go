package listener

import (
	"net"
	"time"
)

type WCApOca6IQW7 = net.PacketConn

// Hzw0dY3ZZL preserves RTTI 0x35ec1e0: a nonembedded net.Conn field.
type Hzw0dY3ZZL struct{ j_Ra2Blepbh net.Conn }

// NSuBse preserves RTTI 0x36dcdb8: embedded PacketConn and fixed source Addr.
type NSuBse struct {
	WCApOca6IQW7
	lxRU7JIrLmL net.Addr
}

func (c *Hzw0dY3ZZL) ReadFrom(p []byte) (int, net.Addr, error) {
	n, e := c.j_Ra2Blepbh.Read(p)
	return n, c.j_Ra2Blepbh.RemoteAddr(), e
}                                                               //0x18bcfe0
func (c *Hzw0dY3ZZL) WriteTo(p []byte, _ net.Addr) (int, error) { return c.j_Ra2Blepbh.Write(p) }            //0x18bd080
func (c *Hzw0dY3ZZL) Close() error                              { return c.j_Ra2Blepbh.Close() }             //0x18bd100
func (c *Hzw0dY3ZZL) LocalAddr() net.Addr                       { return c.j_Ra2Blepbh.LocalAddr() }         //0x18bd140
func (c *Hzw0dY3ZZL) SetDeadline(t time.Time) error             { return c.j_Ra2Blepbh.SetDeadline(t) }      //0x18bd180
func (c *Hzw0dY3ZZL) SetReadDeadline(t time.Time) error         { return c.j_Ra2Blepbh.SetReadDeadline(t) }  //0x18bd1e0
func (c *Hzw0dY3ZZL) SetWriteDeadline(t time.Time) error        { return c.j_Ra2Blepbh.SetWriteDeadline(t) } //0x18bd240
func (c *NSuBse) ReadFrom(p []byte) (int, net.Addr, error) {
	n, _, e := c.WCApOca6IQW7.ReadFrom(p)
	return n, c.lxRU7JIrLmL, e
}                                                  //0x18bcf40
func (c *NSuBse) GetRawPacketConn() net.PacketConn { return c.WCApOca6IQW7 } //0x18bcfc0
// NewTunnelPacketFromStream is an explicitly named source bridge for the
// inlined cross-package construction at 0x1d7034c—0x1d70433. It snapshots
// RemoteAddr before creating the two wrappers, as in the original caller.
func NewTunnelPacketFromStream(c net.Conn) *NSuBse {
	a := c.RemoteAddr()
	return &NSuBse{WCApOca6IQW7: &Hzw0dY3ZZL{j_Ra2Blepbh: c}, lxRU7JIrLmL: a}
}
