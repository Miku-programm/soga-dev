package l_5zgG8VLP

import (
	"io"
	"net"

	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/nzATvajt"
)

type DrpXao = net.Conn

// J3LOqQMzo保留原88字节结构及原字段名。
type J3LOqQMzo struct {
	DrpXao
	g9LlwZz4Rdy  io.Reader
	oWbWTrr2MUiU io.Writer
	d7otQpiN     nzATvajt.PgaNd3014U
	hC8tTwO      *ZfmpaHtmu
	pN7sLgFP     bool
	eJeaO7LW     *MxS0kZ
}

func (c *J3LOqQMzo) Close() error { // @0x1d03920
	if state := c.eJeaO7LW; state != nil && (state.SDw6IIM1D51U || state.Pd9MU7P6IX) {
		return CocBnud6(c.DrpXao.(*J1P4KNvfI_J.Mdn09F6aMrpa)).Close()
	}
	return c.DrpXao.Close()
}
func (c *J3LOqQMzo) Read(p []byte) (int, error)         { return c.g9LlwZz4Rdy.Read(p) }   //0x1d039c0
func (c *J3LOqQMzo) Write(p []byte) (int, error)        { return c.oWbWTrr2MUiU.Write(p) } //0x1d03a20
func (c *J3LOqQMzo) GetUser() nzATvajt.PgaNd3014U       { return c.d7otQpiN }              //0x1d03a80
func (c *J3LOqQMzo) GetDestAddr() *G6bHD5JA.AHU2MRmjzQu { return c.hC8tTwO.if2alzq2rt6 }   //0x1d03aa0
func (c *J3LOqQMzo) NetworkType() J1P4KNvfI_J.ZES6F5 { //0x1d03ac0
	switch c.hC8tTwO.aj8eMgsHQRY {
	case 1:
		return 0
	case 2:
		return 1
	case 3:
		return 2
	default:
		return 255
	}
}
func (c *J3LOqQMzo) GetHeader() *ZfmpaHtmu { return c.hC8tTwO } //0x1d03b00
func (c *J3LOqQMzo) GetRawConn() net.Conn  { return c.DrpXao }  //0x1d03b20
func (c *J3LOqQMzo) Free() { //0x1d03b40
	if free, ok := c.g9LlwZz4Rdy.(HxKViEp2a.JZgM1sMF); ok {
		HxKViEp2a.U9tk4Xosh9(free)
	}
	if free, ok := c.oWbWTrr2MUiU.(HxKViEp2a.JZgM1sMF); ok {
		HxKViEp2a.U9tk4Xosh9(free)
	}
	HxKViEp2a.Li3c4CJ1_D(c.DrpXao)
}
