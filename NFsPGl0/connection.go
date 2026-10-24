package NFsPGl0

import (
	"net"
	G "recovered.local/soga/G6bHD5JA"
	J "recovered.local/soga/J1P4KNvfI_J"
	U "recovered.local/soga/nzATvajt"
)

// EBBW09RjabE 对应原 AnyTLS TCP 连接的 40 字节包装。
type EBBW09RjabE struct {
	net.Conn
	cKQ1zlT U.PgaNd3014U
	wyzfVC  *G.AHU2MRmjzQu
}

func (c *EBBW09RjabE) GetUser() U.PgaNd3014U       { return c.cKQ1zlT } // 0x188a540
func (c *EBBW09RjabE) GetDestAddr() *G.AHU2MRmjzQu { return c.wyzfVC }  // 0x188a560
func (c *EBBW09RjabE) NetworkType() J.ZES6F5       { return 0 }         // 0x188a580

// FromAnyTLSInline 提取 MFF 回调 @0x1e69b73 的原结构构造。
func FromAnyTLSInline(c net.Conn, u U.PgaNd3014U, a *G.AHU2MRmjzQu) *EBBW09RjabE {
	return &EBBW09RjabE{Conn: c, cKQ1zlT: u, wyzfVC: a}
}
