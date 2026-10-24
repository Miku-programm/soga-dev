package J1P4KNvfI_J

import (
	"net"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/nzATvajt"
	"time"
)

// 以下接口、判别类型与字段来自原样本 RTTI；这是源码表示，不是作者原文件。
// 原 bPyexLTWaQN 类型名与实际包路径 J1P4KNvfI_J 不同，采用元数据包路径。
type HpOTlmdRdFv uint8

type E1U4s5iM interface {
	net.Conn
	Free()
}
type NlU9CLU8MJPy interface {
	Close() error
	Free()
	LocalAddr() net.Addr
	ReadFrom([]byte) (int, *G6bHD5JA.AHU2MRmjzQu, error)
	WriteTo([]byte, *G6bHD5JA.AHU2MRmjzQu) (int, error)
	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
}
type Ki8urSHTx44 interface {
	net.Conn
	GetDestAddr() *G6bHD5JA.AHU2MRmjzQu
	GetUser() nzATvajt.PgaNd3014U
	NetworkType() ZES6F5
}

type DqwVny0NcXb struct {
	ccjikIE21 HpOTlmdRdFv
	e9aaCzCOm E1U4s5iM
	zOUdRe0   NlU9CLU8MJPy
}

// AsStream 对应 0x187c720。判别值为零时即使连接为 nil 也返回 true。
func (c *DqwVny0NcXb) AsStream() (E1U4s5iM, bool) {
	if c.ccjikIE21 == 0 {
		return c.e9aaCzCOm, true
	}
	return nil, false
}

// AsDatagram 对应 0x187c740。
func (c *DqwVny0NcXb) AsDatagram() (NlU9CLU8MJPy, bool) {
	if c.ccjikIE21 == 1 {
		return c.zOUdRe0, true
	}
	return nil, false
}

// Free 对应 0x187c760。释放底层缓冲资源，不额外关闭连接。
func (c *DqwVny0NcXb) Free() {
	switch c.ccjikIE21 {
	case 0:
		c.e9aaCzCOm.Free()
	case 1:
		c.zOUdRe0.Free()
	}
}

// Close 对应 0x187c7c0。未知判别值返回 nil；原来的空连接 panic 行为保留。
func (c *DqwVny0NcXb) Close() error {
	switch c.ccjikIE21 {
	case 0:
		return c.e9aaCzCOm.Close()
	case 1:
		return c.zOUdRe0.Close()
	}
	return nil
}
