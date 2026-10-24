// Restored against IDA 0x1b27ae0, 0x1b27b60, 0x1b27c20, 0x1b27d80.
// Source comparison: github.com/Ieooo/clash/transport/ssr/obfs (GPL-3.0).
// Original dependency revision is unproven; names/layout and deviations follow target evidence.
package hysteriaconn

import (
	"crypto/rand"
	"encoding/binary"
	"hash/crc32"
	mathRand "math/rand"
	"net"

	"recovered.local/soga/buf/syncpool"
)

func init() {
	register("random_head", hDXu4s, 0)
}

type qlJYtyKi struct {
	*Ja7FBino7HL
}

func hDXu4s(b *Ja7FBino7HL) AGNBYcIkk {
	return &qlJYtyKi{Ja7FBino7HL: b}
}

type nOAagf7gG3W struct {
	net.Conn
	*qlJYtyKi
	cqIn5z      bool
	bYxvLiDBdWa bool
	fku_KtEBE   bool
	vo4hXV      []byte
}

func (r *qlJYtyKi) StreamConn(c net.Conn) net.Conn {
	return &nOAagf7gG3W{Conn: c, qlJYtyKi: r}
}

func (c *nOAagf7gG3W) Read(b []byte) (int, error) {
	if c.fku_KtEBE {
		return c.Conn.Read(b)
	}
	vo4hXV := syncpool.SharedPoolAt5096978.Get(20480)
	defer syncpool.SharedPoolAt5096978.Put(vo4hXV)
	c.Conn.Read(vo4hXV)
	c.fku_KtEBE = true
	c.Write(nil)
	return 0, nil
}

func (c *nOAagf7gG3W) Write(b []byte) (int, error) {
	if c.bYxvLiDBdWa {
		return c.Conn.Write(b)
	}
	c.vo4hXV = append(c.vo4hXV, b...)
	if !c.cqIn5z {
		c.cqIn5z = true
		dataLength := mathRand.Intn(96) + 4
		vo4hXV := syncpool.SharedPoolAt5096978.Get(dataLength + 4)
		defer syncpool.SharedPoolAt5096978.Put(vo4hXV)
		rand.Read(vo4hXV[:dataLength])
		binary.LittleEndian.PutUint32(vo4hXV[dataLength:], 0xffffffff-crc32.ChecksumIEEE(vo4hXV[:dataLength]))
		_, err := c.Conn.Write(vo4hXV)
		return len(b), err
	}
	if c.fku_KtEBE {
		_, err := c.Conn.Write(c.vo4hXV)
		c.vo4hXV = nil
		c.bYxvLiDBdWa = true
		return len(b), err
	}
	return len(b), nil
}
