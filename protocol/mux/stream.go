package mux

import (
	"io"
	"net"
	"os"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/buf/bufiface"
	"recovered.local/soga/user"
	"recovered.local/soga/util/deadline2"
	"sync"
	"time"
)

// Ldrr25M 按原 RTTI 的88字节结构恢复；源码表达非原作者文件原文。
type Ldrr25M struct {
	asvtPk      sync.Once
	gwE4Or1X8JW uint16
	fj8yHOR     chan *CUwik6VP2WU0
	hiMx_c      net.Conn
	iBwiz6Mebe  user.PgaNd3014U
	yFmjxzi     *relay.AHU2MRmjzQu
	altuteX     bufiface.Mf3phSOVqgf
	s1w1RA      *deadline2.TErFW7l9B2S
}
type JOfnIVIU_k interface {
	Close() error
	Feed(*CUwik6VP2WU0)
}
type XtJioO struct {
	f56bH6 genmap.GTFAWF[uint16, JOfnIVIU_k]
}

// Feed 对应 0x1cea640/0x1cf1c60；向关闭通道发送的 panic 被原闭包恢复并释放帧。
func (c *Ldrr25M) Feed(f *CUwik6VP2WU0) {
	defer func() {
		if recover() != nil {
			f.FreeData()
		}
	}()
	c.fj8yHOR <- f
}

// Read 对应 0x1cea6e0；已有缓冲优先于超时，并忽略底层缓冲 Read 的错误。
func (c *Ldrr25M) Read(p []byte) (int, error) {
	if c.altuteX != nil {
		n, _ := c.altuteX.Read(p)
		if c.altuteX.IsEmpty() {
			c.altuteX.Free()
			c.altuteX = nil
		}
		return n, nil
	}
	select {
	case <-c.s1w1RA.Wait():
		return 0, os.ErrDeadlineExceeded
	case f, ok := <-c.fj8yHOR:
		if !ok {
			return 0, io.EOF
		}
		b := f.c8AtRj7hb
		n, _ := b.Read(p)
		if b.IsEmpty() {
			b.Free()
		} else {
			c.altuteX = b
		}
		return n, nil
	}
}

// Write 对应 0x1cea940；长度写为 uint16，不引入原函数没有的分片或短写检查。
func (c *Ldrr25M) Write(p []byte) (int, error) {
	b := pool.LdHEmSs9(len(p) + 8)
	defer b.Free()
	b.WriteUint16(4)
	b.WriteUint16(c.gwE4Or1X8JW)
	b.WriteByte(2)
	b.WriteByte(1)
	b.WriteUint16(uint16(len(p)))
	b.Write(p)
	if _, err := c.hiMx_c.Write(b.Bytes()); err != nil {
		return 0, err
	}
	return len(p), nil
}

// Close 对应 0x1cead20/0x1cf1d20；清空排队帧，发送 End，写入错误被原函数忽略。
func (c *Ldrr25M) Close() error {
	c.asvtPk.Do(func() {
		close(c.fj8yHOR)
		for f := range c.fj8yHOR {
			f.FreeData()
		}
		b := pool.LdHEmSs9(6)
		defer b.Free()
		b.WriteUint16(4)
		b.WriteUint16(c.gwE4Or1X8JW)
		b.WriteByte(3)
		b.WriteByte(0)
		c.hiMx_c.Write(b.Bytes())
	})
	return nil
}
func (c *Ldrr25M) LocalAddr() net.Addr                { return c.hiMx_c.LocalAddr() }  //0x1cead80
func (c *Ldrr25M) RemoteAddr() net.Addr               { return c.hiMx_c.RemoteAddr() } //0x1ceadc0
func (c *Ldrr25M) SetDeadline(t time.Time) error      { c.s1w1RA.Set(t); return nil }  //0x1ceae00
func (c *Ldrr25M) SetReadDeadline(t time.Time) error  { c.s1w1RA.Set(t); return nil }  //0x1ceae60
func (c *Ldrr25M) SetWriteDeadline(time.Time) error   { return nil }                   //0x1ceaec0，原方法确为空操作
func (c *Ldrr25M) GetUser() user.PgaNd3014U       { return c.iBwiz6Mebe }          //0x1ceaee0
func (c *Ldrr25M) GetDestAddr() *relay.AHU2MRmjzQu { return c.yFmjxzi }             //0x1ceaf00
func (c *Ldrr25M) NetworkType() sess.ZES6F5    { return 0 }                     //0x1ceaf20
func (c *Ldrr25M) Free() {
	if c.altuteX != nil {
		c.altuteX.Free()
		c.altuteX = nil
	}
} //0x1ceaf40
// Clear 对应 0x1ceafc0/0x1cf22a0；原函数关闭所有流后把整个映射重置为零值。
func (m *XtJioO) Clear() {
	m.f56bH6.Range(func(_ uint16, c JOfnIVIU_k) bool { c.Close(); return true })
	m.f56bH6 = genmap.GTFAWF[uint16, JOfnIVIU_k]{}
}
