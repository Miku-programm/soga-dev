package aE7TCWlgkqc

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"net"
	buffer "recovered.local/soga/Hx04IbKk"
	pool "recovered.local/soga/mzi1d3gDO7u"
	"time"
)

type MSUT6Drr struct {
	net.Conn
	aR9AEa      string
	cEZ8tF04gpy int
	xQi_DMa8h_  bool
	iwYJhx      bool
}

// NewTLSObfs reproduces the private-field literal in 0x1b5bce0 under a
// recovery-only exported name. This is TLS-shaped obfuscation, not TLS crypto.
func NewTLSObfs(c net.Conn, host string) *MSUT6Drr {
	return &MSUT6Drr{Conn: c, aR9AEa: host, xQi_DMa8h_: true, iwYJhx: true}
}

// bX6Lk56RcUCZ: 0x19578e0. A length-field read error is swallowed as (0,nil).
func (c *MSUT6Drr) bX6Lk56RcUCZ(p []byte, skip int) (int, error) {
	b := pool.SharedPoolAt5096978.Get(skip)
	_, e := io.ReadFull(c.Conn, b)
	pool.SharedPoolAt5096978.Put(b)
	if e != nil {
		return 0, e
	}
	var size [2]byte
	if _, e = io.ReadFull(c.Conn, size[:]); e != nil {
		return 0, nil
	}
	length := int(binary.BigEndian.Uint16(size[:]))
	if len(p) < length {
		n, e := c.Conn.Read(p)
		if e == nil {
			c.cEZ8tF04gpy = length - n
		}
		return n, e
	}
	return io.ReadFull(c.Conn, p[:length])
}

// Read: 0x1957c20.
func (c *MSUT6Drr) Read(p []byte) (int, error) {
	if c.cEZ8tF04gpy > 0 {
		n, e := io.ReadFull(c.Conn, p[:min(c.cEZ8tF04gpy, len(p))])
		c.cEZ8tF04gpy -= n
		return n, e
	}
	if c.iwYJhx {
		c.iwYJhx = false
		return c.bX6Lk56RcUCZ(p, 105)
	}
	return c.bX6Lk56RcUCZ(p, 3)
}

// Write: 0x1957d60. An error returns the failing chunk's count only.
func (c *MSUT6Drr) Write(p []byte) (int, error) {
	for i := 0; i < len(p); i += 16384 {
		n, e := c.pZ2hHQPXCy1(p[i:min(i+16384, len(p))])
		if e != nil {
			return n, e
		}
	}
	return len(p), nil
}

// pZ2hHQPXCy1: 0x1957e40. First-write state changes even on write failure.
func (c *MSUT6Drr) pZ2hHQPXCy1(p []byte) (int, error) {
	if c.xQi_DMa8h_ {
		b := oRzazA8u(p, c.aR9AEa)
		_, e := c.Conn.Write(b)
		c.xQi_DMa8h_ = false
		return len(p), e
	}
	b := pool.AppendPoolAt5044260.Get().(*buffer.BaDekw7qe)
	defer pool.Wr41CoUmtL0(b)
	b.PutSlice([]byte{23, 3, 3, byte(len(p) >> 8), byte(len(p))})
	b.PutSlice(p)
	_, e := c.Conn.Write(b.Bytes())
	return len(p), e
}

// oRzazA8u: 0x1958120. ClientHello constants below are exact decoded bytes,
// and all integer fields retain the sample's uint16/uint32 truncation.
func oRzazA8u(p []byte, host string) []byte {
	b := buffer.BaDekw7qe(make([]byte, 0))
	put16 := func(n int) { b.PutSlice([]byte{byte(n >> 8), byte(n)}) }
	b.PutSlice([]byte{22, 3, 1})
	put16(len(host) + len(p) + 212)
	b.PutSlice([]byte{1, 0})
	put16(len(host) + len(p) + 208)
	b.PutSlice([]byte{3, 3})
	var stamp [4]byte
	binary.BigEndian.PutUint32(stamp[:], uint32(time.Now().Unix()))
	b.PutSlice(stamp[:])
	b.ReadFull(rand.Reader, 28)
	b.PutSlice([]byte{32})
	b.ReadFull(rand.Reader, 32)
	b.PutSlice([]byte{0, 56})
	b.PutSlice([]byte("\xc0\x2c\xc0\x30\x00\x9f\xcc\xa9\xcc\xa8\xcc\xaa\xc0\x2b\xc0\x2f\x00\x9e\xc0\x24\xc0\x28\x00\x6b\xc0\x23\xc0\x27\x00\x67\xc0\x0a\xc0\x14\x00\x39\xc0\x09\xc0\x13\x00\x33\x00\x9d\x00\x9c\x00\x3d\x00\x3c\x00\x35\x00\x2f\x00\xff"))
	b.PutSlice([]byte{1, 0})
	put16(len(host) + len(p) + 79)
	b.PutSlice([]byte{0, 35})
	put16(len(p))
	b.PutSlice(p)
	b.PutSlice([]byte{0, 0})
	put16(len(host) + 5)
	put16(len(host) + 3)
	b.PutSlice([]byte{0})
	put16(len(host))
	b.PutSlice([]byte(host))
	b.PutSlice([]byte("\x00\x0b\x00\x04\x03\x01\x00\x02"))
	b.PutSlice([]byte("\x00\x0a\x00\x0a\x00\x08\x00\x1d\x00\x17\x00\x19\x00\x18"))
	b.PutSlice([]byte("\x00\x0d\x00\x20\x00\x1e\x06\x01\x06\x02\x06\x03\x05\x01\x05\x02\x05\x03\x04\x01\x04\x02\x04\x03\x03\x01\x03\x02\x03\x03\x02\x01\x02\x02\x02\x03"))
	b.PutSlice([]byte{0, 22, 0, 0, 0, 23, 0, 0})
	return b.Bytes()
}
