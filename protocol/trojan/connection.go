package trojan

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	B "recovered.local/soga/net/dns"
	D "recovered.local/soga/buf/pool"
	G "recovered.local/soga/net/relay"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	U "recovered.local/soga/user"
)

type DrpXao = net.Conn

// DaFYwI 对应原Trojan协议连接，RTTI fr1Zn3.DaFYwI，64字节。
type DaFYwI struct {
	DrpXao
	zBXBLn    *B.D23y2888a_
	uszzyXZ5k *G.AHU2MRmjzQu
	aMMQHtw   bool
	dXeEnBG   U.PgaNd3014U
	nZSKiASba [2]byte
}

// 原0x5036cb0指向二字节0d0a，长度与容量均为2。
var delimiterAt5036CB0 = []byte{'\r', '\n'}

func Lzwh7y6zTaq(conn *J.Mdn09F6aMrpa, resolver *B.D23y2888a_) (*DaFYwI, error) { // 0x1ce8560
	c := &DaFYwI{DrpXao: conn, zBXBLn: resolver, dXeEnBG: conn.GetUser()}
	if err := c.uU9AiiEkD(); err != nil {
		H.Li3c4CJ1_D(c)
		return nil, err
	}
	return c, nil
}
func (c *DaFYwI) uU9AiiEkD() error { //0x1ce8680
	if err := c.tt2gS0tPbX(); err != nil {
		return err
	}
	command, err := G.TBUcsYQLu25q(c.DrpXao)
	if err != nil {
		return err
	}
	if command != 1 && command != 3 {
		return fmt.Errorf("unknown trojan cmd: %v", command)
	}
	addr, err := G.SzhW4tOzUZg(c.DrpXao)
	if err != nil {
		return err
	}
	if err = c.tt2gS0tPbX(); err != nil {
		return err
	}
	c.uszzyXZ5k = addr
	c.aMMQHtw = command == 1
	return nil
}
func (c *DaFYwI) tt2gS0tPbX() error { //0x1ce8940
	if _, err := io.ReadFull(c.DrpXao, c.nZSKiASba[:]); err != nil {
		return err
	}
	if string(c.nZSKiASba[:]) != string(delimiterAt5036CB0) {
		return fmt.Errorf("not crlf: %v", c.nZSKiASba)
	}
	return nil
}
func (c *DaFYwI) eanofxQKvNxV() (int, error) { //0x1ce8b40
	if _, err := io.ReadFull(c.DrpXao, c.nZSKiASba[:]); err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint16(c.nZSKiASba[:])), nil
}

// ReadFromProxy @0x1ce8c20：短缓冲先填满，再丢弃同一报文尾部以保持流边界。
// 正文不足时原实现丢弃已读取的n，返回0,nil,error。
func (c *DaFYwI) ReadFromProxy(p []byte) (int, net.Addr, error) {
	a, err := G.SzhW4tOzUZg(c.DrpXao)
	if err != nil {
		return 0, nil, err
	}
	if !a.HasIP() {
		if err = a.ResolveHost(c.zBXBLn); err != nil {
			return 0, nil, err
		}
	}
	addr, err := a.ToUDPAddr()
	if err != nil {
		return 0, nil, err
	}
	size, err := c.eanofxQKvNxV()
	if err != nil {
		return 0, nil, err
	}
	if err = c.tt2gS0tPbX(); err != nil {
		return 0, nil, err
	}
	if size > len(p) {
		n, err := io.ReadFull(c.DrpXao, p)
		if err != nil {
			return 0, nil, err
		}
		err = G.KlTAqrF42mR(c.DrpXao, size-n)
		return n, addr, err
	}
	if _, err = io.ReadFull(c.DrpXao, p[:size]); err != nil {
		return 0, nil, err
	}
	return size, addr, nil
}

// WriteToProxy @0x1ce9040：原长度字段按uint16截断；底层短写且无错误仍报告正文全长。
func (c *DaFYwI) WriteToProxy(p []byte, addr net.Addr) (int, error) {
	b := D.S4OQXWIZ(len(p) + 23)
	defer b.Free()
	a, err := G.RNFJeaRY(addr)
	if err != nil {
		return 0, err
	}
	if _, err = G.DdYroT1n(b, a); err != nil {
		return 0, err
	}
	b.WriteUint16(uint16(len(p)))
	b.Write(delimiterAt5036CB0)
	b.Write(p)
	if _, err = c.DrpXao.Write(b.Bytes()); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (c *DaFYwI) GetUser() U.PgaNd3014U       { return c.dXeEnBG }   //0x1ce9360
func (c *DaFYwI) GetDestAddr() *G.AHU2MRmjzQu { return c.uszzyXZ5k } //0x1ce9380
func (c *DaFYwI) NetworkType() J.ZES6F5 {
	if c.aMMQHtw {
		return 0
	}
	return 1
}                                      //0x1ce93a0
func (c *DaFYwI) IsMux() bool          { return false }    //0x1ce93c0
func (c *DaFYwI) GetRawConn() net.Conn { return c.DrpXao } //0x1ce93e0
