package encmux

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"golang.org/x/crypto/sha3"
	"net"
	U "recovered.local/soga/net/connunion"
	D "recovered.local/soga/buf/pool"
	H "recovered.local/soga/util/errutil"
)

// 下面三类型来自原RTTI，保持136/24/24字节及实际混淆字段。
type Cb7vj3KQA8 struct {
	DrpXao
	hp3cHCehD2pe *VTl94pzz2IkB
	e8aahi_R3S   cipher.AEAD
	rOfYjOzCdA1L []byte
	bUWvbqWzlH   cipher.AEAD
	cdwwdRLQG    []byte
	w1vg8Dd      *ErrfcUUB
	r6A494       *ErrfcUUB
	npFKkBd4j    U.DwBIUXLAa
}
type ErrfcUUB struct {
	pzOWW_       sha3.ShakeHash
	dyp_KBECm6Dj [2]byte
}
type T4k2R7cJ struct {
	O6B2qZ0p0
	tS5gsM72s *VTl94pzz2IkB
}

// ReadUint16 对应0x11da6a0，原忽略SHAKE读取错误。
func (p *ErrfcUUB) ReadUint16() uint16 {
	p.pzOWW_.Read(p.dyp_KBECm6Dj[:])
	return binary.BigEndian.Uint16(p.dyp_KBECm6Dj[:])
}

// cDhSkRzGR 对应0x11d4d40。认证失败仍推进小端nonce。
func (c *Cb7vj3KQA8) cDhSkRzGR(b U.DwBIUXLAa) error {
	_, err := b.Open(c.e8aahi_R3S, c.rOfYjOzCdA1L, nil)
	for i := range c.rOfYjOzCdA1L {
		c.rOfYjOzCdA1L[i]++
		if c.rOfYjOzCdA1L[i] != 0 {
			break
		}
	}
	return err
}

// vudTI7UEB 对应0x11d4e00，无返回值。
func (c *Cb7vj3KQA8) vudTI7UEB(b U.DwBIUXLAa, data []byte) {
	b.SealWrite(c.bUWvbqWzlH, c.cdwwdRLQG, data, nil)
	for i := range c.cdwwdRLQG {
		c.cdwwdRLQG[i]++
		if c.cdwwdRLQG[i] != 0 {
			break
		}
	}
}

// bpzXJ_7IMXKd 对应0x11d4f00，两个SHAKE值顺序为padding、长度掩码。
func (c *Cb7vj3KQA8) bpzXJ_7IMXKd() (int, int, error) {
	size := c.e8aahi_R3S.Overhead() + 2
	b := D.S4OQXWIZ(size)
	defer b.Free()
	if _, err := b.ReadFullFrom(c.DrpXao, size); err != nil {
		return 0, 0, err
	}
	if err := c.cDhSkRzGR(b); err != nil {
		return 0, 0, err
	}
	length, err := b.ReadUint16()
	if err != nil {
		return 0, 0, err
	}
	padding := int(c.w1vg8Dd.ReadUint16() & 127)
	return int(c.ekKzRm(length)), padding, nil
}

// ekKzRm 对应0x11d52c0。
func (c *Cb7vj3KQA8) ekKzRm(n uint16) uint16 { return n ^ c.w1vg8Dd.ReadUint16() }

// qTbnGwltuBsb 对应0x11d5340。
func (c *Cb7vj3KQA8) qTbnGwltuBsb(n uint16) uint16 { return n ^ c.r6A494.ReadUint16() }

// e21qHx 对应0x11d53c0。原载荷认证错误返回(0,nil)，头认证错误仍返回错误。
func (c *Cb7vj3KQA8) e21qHx(p []byte) (int, error) {
	if c.npFKkBd4j != nil {
		n, _ := c.npFKkBd4j.Read(p)
		if c.npFKkBd4j.Len() == 0 {
			c.npFKkBd4j.Free()
			c.npFKkBd4j = nil
		}
		return n, nil
	}
	length, padding, err := c.bpzXJ_7IMXKd()
	if err != nil {
		return 0, err
	}
	size := length + c.e8aahi_R3S.Overhead()
	b := D.S4OQXWIZ(size)
	free := true
	defer func() {
		if free {
			b.Free()
		}
	}()
	if _, err := b.ReadFullFrom(c.DrpXao, size); err != nil {
		return 0, err
	}
	if err := c.cDhSkRzGR(b); err != nil {
		return 0, nil
	}
	b.ReverseEnd(padding)
	n := 0
	if !b.IsEmpty() {
		n, _ = b.Read(p)
	}
	if b.Len() > 0 {
		free = false
		c.npFKkBd4j = b
	}
	return n, nil
}

// evEIO6oon0 对应0x11d5820。
func (c *Cb7vj3KQA8) evEIO6oon0(b U.DwBIUXLAa, length uint16) {
	header := make([]byte, 2)
	binary.BigEndian.PutUint16(header, c.qTbnGwltuBsb(length))
	c.vudTI7UEB(b, header)
}

// oH5wsFE 对应0x11d58e0。超过65535递归切块；填充长度仍按uint16截断。
// 原实现允许写入调用者切片的剩余容量，底层短写且无错误视作成功。
func (c *Cb7vj3KQA8) oH5wsFE(prefix, data []byte) (int, error) {
	if len(data) > 65535 {
		n, err := c.oH5wsFE(prefix, data[:65535])
		if err != nil {
			return n, err
		}
		more, err := c.oH5wsFE(nil, data[65535:])
		return n + more, err
	}
	base := len(data) + c.bUWvbqWzlH.Overhead() + len(prefix)
	padding := int(c.r6A494.ReadUint16() & 127)
	b := D.S4OQXWIZ(base + padding + c.bUWvbqWzlH.Overhead() + 2)
	defer b.Free()
	if len(prefix) != 0 {
		b.Write(prefix)
	}
	length := len(data) + padding
	c.evEIO6oon0(b, uint16(length))
	var padded []byte
	if cap(data) < length {
		offset := b.Len()
		b.Write(data)
		b.WriteRandom(padding)
		padded = b.BytesFrom(offset)
		b.ReverseEnd(length)
	} else {
		padded = data[:length]
		rand.Read(padded[len(data):])
	}
	c.vudTI7UEB(b, padded)
	if _, err := c.DrpXao.Write(b.Bytes()); err != nil {
		return 0, err
	}
	return len(data), nil
}

// GetRawConn 对应0x11d5f40。
func (c *Cb7vj3KQA8) GetRawConn() net.Conn { return c.DrpXao }

// Free 对应0x11d5f60，不清缓存接口或关闭连接。
func (c *Cb7vj3KQA8) Free() {
	if c.npFKkBd4j != nil {
		c.npFKkBd4j.Free()
	}
	H.Li3c4CJ1_D(c.DrpXao)
}
