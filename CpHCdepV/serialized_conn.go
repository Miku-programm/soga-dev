package CpHCdepV

import (
	"net"
	"sync"
)

type DrpXao = net.Conn

// MhXtuhyPkR5 为原24字节结构，两个嵌入字段偏移0/16。
type MhXtuhyPkR5 struct {
	DrpXao
	sync.Mutex
}

// Write 对应 0x18bd2a0。原代码显式解锁，没有 defer；底层 panic 时锁仍锁定。
func (c *MhXtuhyPkR5) Write(p []byte) (int, error) {
	c.Lock()
	n, err := c.DrpXao.Write(p)
	c.Unlock()
	return n, err
}
func (c *MhXtuhyPkR5) GetRawConn() net.Conn { return c.DrpXao } //0x18bd3a0
