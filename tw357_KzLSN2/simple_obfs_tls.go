package tw357_KzLSN2

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	R "recovered.local/soga/A818JHb"
	T "recovered.local/soga/CAodbhZf"
	C "recovered.local/soga/CpHCdepV"
	P "recovered.local/soga/D_zCag1qTV"
	S "recovered.local/soga/ad5KEbw_EzWk"
	"time"
)

// cqVDRTx preserves the original 104-byte RTTI, including the session ID
// alias into the subsequently released read buffer.
type cqVDRTx struct {
	io.ReadWriter
	bg3qmAv     int
	a1yhtxpPH   []byte
	pDziAaUrdOb *P.X_aXX5
	sBRQD6eVrE  int
	jCBNufN     int64
	bi41cm802Z  *S.L7E7H2aW8KVs
	C.JQi6241
}

func crpBB1_mZ(rw io.ReadWriter) Jakq78yIQ         { return &cqVDRTx{ReadWriter: rw} } // 0x1cb0000
func (c *cqVDRTx) SetServerInfo(i *S.L7E7H2aW8KVs) { c.bi41cm802Z = i }                // 0x1cb1aa0
func (*cqVDRTx) GetOverhead() int                  { return 5 }                        // 0x1cb1b00
// a8xWy1CIZ3n restores 0x1cb0080. Original unchecked short-field slices
// panic; it is not a general-purpose TLS parser or a cryptographic TLS session.
func (c *cqVDRTx) a8xWy1CIZ3n() error {
	header, e := c.GetReadBuf(5)
	if e != nil {
		return e
	}
	if _, e = io.ReadFull(c.ReadWriter, header); e != nil {
		return e
	}
	if !bytes.HasPrefix(header, []byte{22, 3, 1}) {
		return errors.New("tls handshake failed by prefix: \\x16\\x03\\x01")
	}
	record, e := c.GetReadBuf(int(binary.BigEndian.Uint16(header[3:])))
	if e != nil {
		return e
	}
	if _, e = io.ReadFull(c.ReadWriter, record); e != nil {
		return e
	}
	if !bytes.HasPrefix(record, []byte{1, 0}) {
		return errors.New("tls_auth not client hello message")
	}
	if int(binary.BigEndian.Uint16(record[2:])) != len(record)-4 {
		return errors.New("tls_auth wrong message size")
	}
	rest := record[4:]
	if !bytes.HasPrefix(rest, []byte{3, 3}) {
		return errors.New("tls_auth wrong tls version")
	}
	rest = rest[2:]
	random := rest[:32]
	rest = rest[32:]
	sidlen := int(rest[0])
	if sidlen < 32 {
		return errors.New("tls_auth wrong session id length")
	}
	c.a1yhtxpPH = rest[1 : 1+sidlen]
	rest = rest[1+sidlen:]
	timestamp := int64(binary.BigEndian.Uint32(random))
	now := time.Now().Unix()
	if now-timestamp > 86400 || timestamp-now > 86400 {
		return errors.New("tls auth wrong time")
	}
	if !R.HFlnXWty(random[:22]) {
		return errors.New("tls auth repeat attack detect")
	}
	rest = rest[62:]
	if !bytes.HasPrefix(rest, []byte{0, 35}) {
		return errors.New("ext header error")
	}
	rest = rest[2:]
	size := int(binary.BigEndian.Uint16(rest))
	rest = rest[2:]
	if size != 0 {
		c.pDziAaUrdOb = P.S4OQXWIZ(size)
		c.pDziAaUrdOb.Write(rest[:size])
		c.bg3qmAv = size
	}
	rest = rest[size+7:]
	hostlen := int(binary.BigEndian.Uint16(rest))
	c.bi41cm802Z.EFLMTH = string(rest[2 : 2+hostlen])
	c.FreeReadBuffer()
	return nil
}

// lHS3ppzaAv restores 0x1cb0da0, returning at most one record fragment.
func (c *cqVDRTx) lHS3ppzaAv(p []byte) (int, error) {
	if c.bg3qmAv > 0 {
		if c.pDziAaUrdOb != nil {
			n := copy(p, c.pDziAaUrdOb.Bytes())
			c.pDziAaUrdOb.Skip(n)
			if c.pDziAaUrdOb.IsEmpty() {
				c.pDziAaUrdOb.Free()
				c.pDziAaUrdOb = nil
			}
			c.bg3qmAv -= n
			return n, nil
		}
		n, e := c.ReadWriter.Read(p[:min(len(p), c.bg3qmAv)])
		c.bg3qmAv -= n
		return n, e
	}
	header, e := c.GetReadBuf(5)
	if e != nil {
		return 0, e
	}
	if _, e = io.ReadFull(c.ReadWriter, header); e != nil {
		return 0, e
	}
	if !bytes.HasPrefix(header, []byte{23, 3, 3}) {
		return 0, errors.New("data prefix wrong:")
	}
	size := int(binary.BigEndian.Uint16(header[3:]))
	n, e := c.ReadWriter.Read(p[:min(len(p), size)])
	c.bg3qmAv = size - n
	return n, e
}
func (c *cqVDRTx) Read(p []byte) (int, error) { // 0x1cb11c0
	if c.jCBNufN == 1 {
		return c.lHS3ppzaAv(p)
	}
	if e := c.a8xWy1CIZ3n(); e != nil {
		return 0, e
	}
	c.jCBNufN = 1
	return c.Read(p)
}
func (c *cqVDRTx) aF2gRvPD(p []byte) (int, error) { // 0x1cb1280
	timestamp := uint32(T.C5Jff5vs0())
	b, e := c.GetWriteBuffer(8192)
	if e != nil {
		return 0, e
	}
	var scratch [32]byte
	b.Write([]byte{22, 3, 1, 0, 91, 2, 0, 0, 87, 3, 3})
	binary.BigEndian.PutUint32(scratch[:], timestamp)
	b.Write(scratch[:4])
	rand.Read(scratch[:28])
	b.Write(scratch[:28])
	b.WriteByte(32)
	if len(c.a1yhtxpPH) >= 32 {
		b.Write(c.a1yhtxpPH[:32])
	} else {
		rand.Read(scratch[:])
		b.Write(scratch[:])
	}
	// Decoder at 0x1cb1475 returns this exact 20-byte binary string.
	b.Write([]byte{0xcc, 0xa8, 0, 0, 0, 0xff, 1, 0, 1, 0, 0, 0x17, 0, 0, 0, 0x0b, 0, 2, 1, 0})
	b.Write([]byte{20, 3, 3, 0, 1, 1})
	var random [4]byte
	rand.Read(random[:])
	size := min(len(p), int(binary.BigEndian.Uint32(random[:])&4095)+100)
	b.Write([]byte{22, 3, 3})
	binary.BigEndian.PutUint16(scratch[:], uint16(size))
	b.Write(scratch[:2])
	b.Write(p[:size])
	c.sBRQD6eVrE++
	if _, e = c.ReadWriter.Write(b.Bytes()); e != nil {
		return 0, e
	}
	return size, nil
}
func (c *cqVDRTx) fv8JhHP0(p []byte) (int, error) { // 0x1cb16c0
	b, e := c.GetWriteBuffer(len(p) + 100)
	if e != nil {
		return 0, e
	}
	original := len(p)
	var random [4]byte
	var sizeBytes [2]byte
	emit := func(data []byte) {
		b.Write([]byte{23, 3, 3})
		binary.BigEndian.PutUint16(sizeBytes[:], uint16(len(data)))
		b.Write(sizeBytes[:])
		b.Write(data)
	}
	for len(p) > 2048 {
		rand.Read(random[:])
		size := min(len(p), int(uint16(binary.BigEndian.Uint32(random[:]))))
		emit(p[:size])
		p = p[size:]
	}
	if len(p) != 0 {
		emit(p)
	}
	if _, e = c.ReadWriter.Write(b.Bytes()); e != nil {
		return 0, e
	}
	return original, nil
}
func (c *cqVDRTx) Write(p []byte) (int, error) { //0x1cb19c0
	if c.sBRQD6eVrE == 1 {
		return c.fv8JhHP0(p)
	}
	n, e := c.aF2gRvPD(p)
	if e != nil {
		return 0, e
	}
	if n >= len(p) {
		return n, nil
	}
	m, e := c.Write(p[n:])
	return n + m, e
}
