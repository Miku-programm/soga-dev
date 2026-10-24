package VwwneWVFFv

import (
	"crypto/cipher"
	"crypto/rand"
	"io"
	"net"
	pool "recovered.local/soga/mzi1d3gDO7u"
)

type SBrL2pUbjeTr struct {
	io.Writer
	cipher.Stream
	boO05T [2048]byte
}
type Z0mJsCE struct {
	io.Reader
	cipher.Stream
	hazikx8LhJA [2048]byte
}
type HtnJMv4Lsng struct {
	net.PacketConn
	ZDHlm7vHo
}
type AriwPpia struct {
	net.Conn
	ZDHlm7vHo
	ctE9MncrAhJ *Z0mJsCE
	htrPBreLEcY *SBrL2pUbjeTr
	l1CVSA      []byte
	ipDTi3E     []byte
}

// LTkaLW3KDj: 0x194ad60; SP6JM3LOTFRv: 0x194af60.
func LTkaLW3KDj(dst, plain []byte, c ZDHlm7vHo) ([]byte, error) {
	if len(dst) < len(plain)+c.IVSize() {
		return nil, io.ErrShortBuffer
	}
	iv := dst[:c.IVSize()]
	if _, e := rand.Read(iv); e != nil {
		return nil, e
	}
	c.Encrypter(iv).XORKeyStream(dst[len(iv):], plain)
	return dst[:len(iv)+len(plain)], nil
}
func SP6JM3LOTFRv(dst, packet []byte, c ZDHlm7vHo) ([]byte, error) {
	if len(packet) < c.IVSize() {
		return nil, errorAt5098550
	}
	if len(dst) < len(packet)-c.IVSize() {
		return nil, io.ErrShortBuffer
	}
	iv := packet[:c.IVSize()]
	c.Decrypter(iv).XORKeyStream(dst, packet[len(iv):])
	return dst[:len(packet)-len(iv)], nil
}

// WriteTo: 0x194b160; ReadFrom: 0x194b360.
func (c *HtnJMv4Lsng) WriteTo(p []byte, addr net.Addr) (int, error) {
	b := pool.SharedPoolAt5096978.Get(65536)
	defer pool.SharedPoolAt5096978.Put(b)
	packet, e := LTkaLW3KDj(b, p, c.ZDHlm7vHo)
	if e != nil {
		return 0, e
	}
	_, e = c.PacketConn.WriteTo(packet, addr)
	return len(p), e
}
func (c *HtnJMv4Lsng) ReadFrom(p []byte) (int, net.Addr, error) {
	n, addr, e := c.PacketConn.ReadFrom(p)
	if e != nil {
		return n, addr, e
	}
	out, e := SP6JM3LOTFRv(p[c.IVSize():], p[:n], c.ZDHlm7vHo)
	if e != nil {
		return n, addr, e
	}
	copy(p, out)
	return len(out), addr, nil
}

// Write: 0x194b6c0. Advance plaintext by the underlying write count, even
// though the stream has consumed the whole attempted chunk on a short write.
func (w *SBrL2pUbjeTr) Write(p []byte) (n int, err error) {
	for n < len(p) && err == nil {
		size := min(2048, len(p)-n)
		w.XORKeyStream(w.boO05T[:], p[n:n+size])
		var k int
		k, err = w.Writer.Write(w.boO05T[:size])
		n += k
	}
	return
}

// ReadFrom: 0x194b820. Count input bytes, ignore short output counts, and write
// an empty buffer for a zero-byte terminal read.
func (w *SBrL2pUbjeTr) ReadFrom(r io.Reader) (total int64, err error) {
	for {
		n, re := r.Read(w.boO05T[:])
		b := w.boO05T[:n]
		total += int64(n)
		w.XORKeyStream(b, b)
		_, err = w.Writer.Write(b)
		if err != nil {
			return
		}
		if re != nil {
			if re != io.EOF {
				err = re
			}
			return
		}
	}
}

// Read: 0x194b9c0. Any input error discards even a positive input count.
func (r *Z0mJsCE) Read(p []byte) (int, error) {
	n, e := r.Reader.Read(p)
	if e != nil {
		return 0, e
	}
	r.XORKeyStream(p, p[:n])
	return n, nil
}

// WriteTo: 0x194baa0 differs from Read: bytes with an input error are decoded
// and written first, and only the actual output count is accumulated.
func (r *Z0mJsCE) WriteTo(w io.Writer) (total int64, err error) {
	for {
		n, re := r.Reader.Read(r.hazikx8LhJA[:])
		if n > 0 {
			r.XORKeyStream(r.hazikx8LhJA[:], r.hazikx8LhJA[:n])
			k, e := w.Write(r.hazikx8LhJA[:n])
			total += int64(k)
			if e != nil {
				return total, e
			}
		}
		if re != nil {
			if re != io.EOF {
				err = re
			}
			return
		}
	}
}

// ObtainWriteIV: 0x194c200. The returned IV aliases the cached slice.
func (c *AriwPpia) ObtainWriteIV() ([]byte, error) {
	if len(c.ipDTi3E) == c.IVSize() {
		return c.ipDTi3E, nil
	}
	iv := make([]byte, c.IVSize())
	if _, e := rand.Read(iv); e != nil {
		return nil, e
	}
	c.ipDTi3E = iv
	return iv, nil
}

// ObtainReadIV: 0x194c320. Partial read failure leaves the cache untouched.
func (c *AriwPpia) ObtainReadIV() ([]byte, error) {
	if len(c.l1CVSA) == c.IVSize() {
		return c.l1CVSA, nil
	}
	iv := make([]byte, c.IVSize())
	if _, e := io.ReadFull(c.Conn, iv); e != nil {
		return nil, e
	}
	c.l1CVSA = iv
	return iv, nil
}

// yoZaA7MSKhPg: 0x194bc40; _T6g1HQkOdT: 0x194bf00.
func (c *AriwPpia) yoZaA7MSKhPg() error {
	if c.ctE9MncrAhJ == nil {
		iv, e := c.ObtainReadIV()
		if e != nil {
			return e
		}
		c.ctE9MncrAhJ = &Z0mJsCE{Reader: c.Conn, Stream: c.Decrypter(iv)}
	}
	return nil
}
func (c *AriwPpia) _T6g1HQkOdT() error {
	if c.htrPBreLEcY == nil {
		iv, e := c.ObtainWriteIV()
		if e != nil {
			return e
		}
		if _, e = c.Conn.Write(iv); e != nil {
			return e
		}
		c.htrPBreLEcY = &SBrL2pUbjeTr{Writer: c.Conn, Stream: c.Encrypter(iv)}
	}
	return nil
}

// Read/WriteTo/Write/ReadFrom: 0x194bdc0/194be60/194c0c0/194c160.
func (c *AriwPpia) Read(p []byte) (int, error) {
	if e := c.yoZaA7MSKhPg(); e != nil {
		return 0, e
	}
	return c.ctE9MncrAhJ.Read(p)
}
func (c *AriwPpia) WriteTo(w io.Writer) (int64, error) {
	if e := c.yoZaA7MSKhPg(); e != nil {
		return 0, e
	}
	return c.ctE9MncrAhJ.WriteTo(w)
}
func (c *AriwPpia) Write(p []byte) (int, error) {
	if e := c._T6g1HQkOdT(); e != nil {
		return 0, e
	}
	return c.htrPBreLEcY.Write(p)
}
func (c *AriwPpia) ReadFrom(r io.Reader) (int64, error) {
	if e := c._T6g1HQkOdT(); e != nil {
		return 0, e
	}
	return c.htrPBreLEcY.ReadFrom(r)
}
