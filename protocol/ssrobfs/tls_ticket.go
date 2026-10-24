package ssrobfs

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	R "recovered.local/soga/crypto/replaywindow"
	C "recovered.local/soga/transport/listener"
	H "recovered.local/soga/util/ssrhelper"
	S "recovered.local/soga/protocol/ssrshared"
	Q "recovered.local/soga/crypto/rand"
	"strconv"
	"time"
)

// Original 120-byte RTTI 0x36f6d40.
type g7X2glAwFd91 struct {
	io.ReadWriter
	kr85QF       int
	g3yMrTl9L    bool
	vZGBncOF67kg int
	lHSOwnACEPf  []byte
	sxpzXaFp32P  []byte
	g30R56fXh    *S.L7E7H2aW8KVs
	C.JQi6241
}

// Original init@1ca71a0 decoded string converted to mutable byte slice.
var ticketPrefixAt509A770 = []byte{20, 3, 3, 0, 1, 1, 22, 3, 3, 0}

func jRGJRCFN(rw io.ReadWriter) Jakq78yIQ               { return &g7X2glAwFd91{ReadWriter: rw} } // 0x1cb5de0
func (c *g7X2glAwFd91) SetServerInfo(i *S.L7E7H2aW8KVs) { c.g30R56fXh = i }                      //0x1cb5e60
func (*g7X2glAwFd91) GetOverhead() int                  { return 5 }                             //0x1cb7ee0
func (c *g7X2glAwFd91) hS8VwDp3P() error { //0x1cb5ec0
	b, e := c.GetReadBuffer(1024)
	if e != nil {
		return e
	}
	if _, e = b.ReadFullFrom(c.ReadWriter, 5); e != nil {
		return e
	}
	header := b.Bytes()
	if !bytes.HasPrefix(header, []byte{22, 3, 1}) {
		return errors.New("tls handshake failed by prefix: \\x16\\x03\\x01")
	}
	size := int(binary.BigEndian.Uint16(header[3:]))
	b.Reset()
	if _, e = b.ReadFullFrom(c.ReadWriter, size); e != nil {
		return e
	}
	data := b.Bytes()
	if !bytes.HasPrefix(data, []byte{1, 0}) {
		return errors.New("tls auth not client hello message")
	}
	declared := int(binary.BigEndian.Uint16(data[2:]))
	if declared != len(data)-4 {
		return errors.New("tls auth wrong message size: " + strconv.Itoa(declared))
	}
	rest := data[4:]
	if !bytes.HasPrefix(rest, []byte{3, 3}) {
		return errors.New("tls auth wrong tls version")
	}
	rest = rest[2:]
	random := rest[:32]
	rest = rest[32:]
	sidlen := int(rest[0])
	if sidlen < 32 {
		return errors.New("tls auth wrong session_id_len")
	}
	sid := rest[1 : 1+sidlen]
	rest = rest[1+sidlen:]
	c.sxpzXaFp32P = make([]byte, sidlen)
	copy(c.sxpzXaFp32P, sid)
	c.lHSOwnACEPf = make([]byte, c.g30R56fXh.SP2KC4lkkl3C+sidlen)
	copy(c.lHSOwnACEPf, c.g30R56fXh.IawCjSkuTM4)
	copy(c.lHSOwnACEPf[c.g30R56fXh.SP2KC4lkkl3C:], sid)
	now := time.Now().Unix()
	stamp := int64(binary.BigEndian.Uint32(random))
	if now-stamp > 86400 || stamp-now > 86400 {
		return errors.New("tls auth wrong time")
	}
	if !bytes.Equal(H.L01ck7znS(c.lHSOwnACEPf, random[:22]), random[22:32]) {
		return errors.New("tls auth wrong sha1")
	}
	if !R.HFlnXWty(random[:22]) {
		return errors.New("tls auth repeat attack detect")
	}
	c.kr85QF = 1
	rest = rest[48:]
	if index := bytes.Index(rest, []byte{0, 23, 0, 0}); index > 0 {
		c.g30R56fXh.EFLMTH = string(rest[:index])
	}
	return c.e1hcfL5l()
}
func (c *g7X2glAwFd91) pd3J1p3AGFK() []byte { //0x1cb6c80
	data := make([]byte, 32)
	binary.BigEndian.PutUint32(data, uint32(time.Now().Unix()))
	rand.Read(data[4:22])
	copy(data[22:], H.L01ck7znS(c.lHSOwnACEPf, data[:22]))
	return data
}
func (c *g7X2glAwFd91) e1hcfL5l() error { //0x1cb6da0
	random := c.pd3J1p3AGFK()
	// The original decoder returns these ASCII hexadecimal characters, not
	// their decoded bytes. Retain the resulting 20-byte field exactly.
	tail := []byte("c02f000005ff01000100")
	size := len(c.sxpzXaFp32P) + len(random) + 2 + len(tail) + 1
	b, e := c.GetWriteBuffer(1024)
	if e != nil {
		return e
	}
	var u16 [2]byte
	writeSize := func(n int) { binary.BigEndian.PutUint16(u16[:], uint16(n)); b.Write(u16[:]) }
	b.Write([]byte{22, 3, 3})
	writeSize(size + 4)
	b.Write([]byte{2, 0})
	writeSize(size)
	b.Write([]byte{3, 3})
	b.Write(random)
	b.WriteByte(32)
	b.Write(c.sxpzXaFp32P)
	b.Write(tail)
	if Q.Co6D_zfls(8) <= 0 {
		n := Q.Co6D_zfls(65535)%164*2 + 64
		b.Write([]byte{22, 3, 3})
		writeSize(n + 4)
		b.Write([]byte{4, 0})
		writeSize(n)
		b.WriteRandom(n)
	}
	b.Write([]byte{20, 3, 3, 0, 1, 1, 22, 3, 3})
	n := Q.Co6D_zfls(8) + 22
	writeSize(n + 10)
	b.WriteRandom(n)
	b.Write(H.L01ck7znS(c.lHSOwnACEPf, b.Bytes()))
	_, e = c.ReadWriter.Write(b.Bytes())
	return e
}
func (c *g7X2glAwFd91) gZrZyY() error { //0x1cb7360
	b, e := c.GetReadBuffer(1024)
	if e != nil {
		return e
	}
	if _, e = b.ReadFullFrom(c.ReadWriter, 11); e != nil {
		return e
	}
	if !bytes.HasPrefix(b.Bytes(), ticketPrefixAt509A770) {
		return errors.New("handshake2 first data not match")
	}
	size := int(binary.BigEndian.Uint16(b.BytesRange(9, 11)))
	if _, e = b.ReadFullFrom(c.ReadWriter, size); e != nil {
		return e
	}
	if !bytes.Equal(H.L01ck7znS(c.lHSOwnACEPf, b.BytesTo(size+1)), b.BytesRange(size+1, size+11)) {
		return errors.New("handshake2 verify failed")
	}
	c.kr85QF |= 4
	return nil
}
func (c *g7X2glAwFd91) edo12OH() error { //0x1cb7800
	if e := c.hS8VwDp3P(); e != nil {
		return e
	}
	if c.kr85QF&1 != 0 {
		if e := c.gZrZyY(); e != nil {
			return e
		}
	}
	c.FreeReadBuffer()
	c.g3yMrTl9L = true
	return nil
}
func (c *g7X2glAwFd91) lHS3ppzaAv(p []byte) (int, error) { //0x1cb78c0
	if c.vZGBncOF67kg > 0 {
		n, e := c.ReadWriter.Read(p[:min(len(p), c.vZGBncOF67kg)])
		c.vZGBncOF67kg -= n
		return n, e
	}
	var header [5]byte
	if _, e := io.ReadFull(c.ReadWriter, header[:]); e != nil {
		return 0, e
	}
	if header[0] != 23 || header[1] != 3 || header[2] != 3 {
		return 0, errors.New("handshake3 prefix wrong")
	}
	size := int(binary.BigEndian.Uint16(header[3:]))
	n, e := c.ReadWriter.Read(p[:min(len(p), size)])
	if e == nil {
		c.vZGBncOF67kg = size - n
	}
	return n, e
}
func (c *g7X2glAwFd91) Read(p []byte) (int, error) { //0x1cb7b40
	if !c.g3yMrTl9L {
		if e := c.edo12OH(); e != nil {
			return 0, e
		}
	}
	return c.lHS3ppzaAv(p)
}
func (c *g7X2glAwFd91) fv8JhHP0(p []byte) (int, error) { //0x1cb7be0
	b, e := c.GetWriteBuffer(len(p) + 50)
	if e != nil {
		return 0, e
	}
	total := len(p)
	var header = [5]byte{23, 3, 3, 0, 0}
	emit := func(part []byte) {
		binary.BigEndian.PutUint16(header[3:], uint16(len(part)))
		b.Write(header[:])
		b.Write(part)
	}
	for len(p) > 2048 {
		n := min(len(p), Q.Co6D_zfls(65535)%4096+100)
		emit(p[:n])
		p = p[n:]
	}
	if len(p) != 0 {
		emit(p)
	}
	if _, e = c.ReadWriter.Write(b.Bytes()); e != nil {
		return 0, e
	}
	return total, nil
}
func (c *g7X2glAwFd91) Write(p []byte) (int, error) { return c.fv8JhHP0(p) } //0x1cb7e80
