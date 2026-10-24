package ssrprotocol

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	Clock "recovered.local/soga/util/rtx"
	C "recovered.local/soga/transport/listener"
	P "recovered.local/soga/buf/pool"
	H "recovered.local/soga/util/ssrhelper"
	T "recovered.local/soga/protocol/ssrcipher"
	S "recovered.local/soga/protocol/ssrshared"
	"strconv"
)

type d9oPDg_ func([]byte, ...[]byte) []byte
type nw012W7AIY func([]byte) []byte

// Original RTTI 0x37273b8, 216 bytes. Both AES variants use this type.
type lKlYpdomLrhv struct {
	io.ReadWriter
	WCApOca6IQW7
	ggoDIOgz8    *S.L7E7H2aW8KVs
	h6LJlye      []byte
	aVq78HP      string
	v9eMMxHHk    d9oPDg_
	idj1Fac8     nw012W7AIY
	eOnqFdRmAaYM uint32
	jJzR7oD6     []byte
	u7Ndn2SMlS   uint32
	a2AxZU       []byte
	m74sJGwc     bool
	qQSF_JOQQp   *P.X_aXX5
	oQSVsjCkhD8  int64
	wZMIk2Ur     int64
	C.JQi6241
}

func jKGRC86Lx9g(c *T.KrOWYtzUYft) Vzuv_s9ZS5 {
	return &lKlYpdomLrhv{ReadWriter: c, WCApOca6IQW7: c, aVq78HP: "auth_aes128_md5", v9eMMxHHk: H.GhU1IKb4, idj1Fac8: H.KWa8h0, eOnqFdRmAaYM: 1, u7Ndn2SMlS: 1, qQSF_JOQQp: P.S4OQXWIZ(8192)}
} //0x1cbfd40
func dffBLS(c *T.KrOWYtzUYft) Vzuv_s9ZS5 {
	b := P.S4OQXWIZ(8192)
	b.Reset()
	return &lKlYpdomLrhv{ReadWriter: c, WCApOca6IQW7: c, aVq78HP: "auth_aes128_sha1", v9eMMxHHk: H.L01ck7znS, idj1Fac8: H.EyTkbUUkjjn, eOnqFdRmAaYM: 1, u7Ndn2SMlS: 1, qQSF_JOQQp: b}
}                                                       //0x1ccbce0
func (c *lKlYpdomLrhv) SetServerInfo(i *S.L7E7H2aW8KVs) { c.ggoDIOgz8 = i }                       //0x1cbfe80
func (c *lKlYpdomLrhv) CurUID() int64                   { return c.oQSVsjCkhD8 }                  //0x1cc2700
func (c *lKlYpdomLrhv) CID() int64                      { return c.wZMIk2Ur }                     //0x1cc2720
func (*lKlYpdomLrhv) GetOverhead() int                  { return 9 }                              //0x1cc2740
func (c *lKlYpdomLrhv) Free()                           { c.JQi6241.Free(); c.qQSF_JOQQp.Free() } //0x1cc2760
func (c *lKlYpdomLrhv) bQ9xnsqWMS65(p []byte) ([]byte, error) { //0x1cbfee0
	b, e := c.GetWriteBuf(len(p) + 9)
	if e != nil {
		return nil, e
	}
	binary.LittleEndian.PutUint16(b, uint16(len(p)+9))
	if c.a2AxZU == nil {
		c.a2AxZU = make([]byte, len(c.h6LJlye)+4)
		copy(c.a2AxZU, c.h6LJlye)
	}
	binary.LittleEndian.PutUint32(c.a2AxZU[len(c.a2AxZU)-4:], c.u7Ndn2SMlS)
	c.u7Ndn2SMlS++
	copy(b[2:], c.v9eMMxHHk(c.a2AxZU, b[:2])[:2])
	b[4] = 1
	copy(b[5:], p)
	copy(b[len(p)+5:], c.v9eMMxHHk(c.a2AxZU, b[:len(p)+5])[:4])
	return b[:len(p)+9], nil
}
func (c *lKlYpdomLrhv) rTiayU() ([]byte, error) { //0x1cc0300
	var header [31]byte
	if _, e := io.ReadFull(c.ReadWriter, header[:]); e != nil {
		return nil, e
	}
	var key []byte
	switch c.ggoDIOgz8.NxYA16xMfF {
	case 0:
		c.oQSVsjCkhD8 = c.ggoDIOgz8.Uklg6401
		key = c.ggoDIOgz8.IawCjSkuTM4
	case 1, 2:
		c.oQSVsjCkhD8 = int64(binary.LittleEndian.Uint32(header[7:]))
		password, ok := c.ggoDIOgz8.ZlsyKgz4EI.Get(c.oQSVsjCkhD8)
		if !ok {
			return nil, errors.New("can not find uid: " + strconv.FormatInt(c.oQSVsjCkhD8, 10))
		}
		key = c.idj1Fac8([]byte(password))
	default:
		return nil, fmt.Errorf("auth_aes not support single port type: %v", c.ggoDIOgz8.NxYA16xMfF)
	}
	c.h6LJlye = make([]byte, len(key))
	copy(c.h6LJlye, key)
	block, e := aes.NewCipher(H.Is8IqhjVE6(base64.StdEncoding.EncodeToString(key)+c.aVq78HP, 16))
	if e != nil {
		return nil, errors.New("creating aes cipher failed: " + e.Error())
	}
	var decoded [16]byte
	cipher.NewCBCDecrypter(block, make([]byte, 16)).CryptBlocks(decoded[:], header[11:27])
	stamp := binary.LittleEndian.Uint32(decoded[:])
	c.wZMIk2Ur = int64(binary.LittleEndian.Uint32(decoded[4:]))
	length := binary.LittleEndian.Uint16(decoded[12:])
	pad := int(binary.LittleEndian.Uint16(decoded[14:]))
	// Original 0x1cc06d8/0x1cc06e9 use 32-bit LEA and unsigned comparisons.
	now := uint32(Clock.C5Jff5vs0())
	if stamp < now-86400 || stamp > now+86400 {
		return nil, errors.New("utc invalid")
	}
	info := c.ggoDIOgz8
	authkey := make([]byte, info.SP2KC4lkkl3C+info.DvdB7JZX)
	copy(authkey, info.Ds7eTY7Qp66N)
	copy(authkey[info.DvdB7JZX:], info.IawCjSkuTM4)
	if !bytes.Equal(header[1:7], c.v9eMMxHHk(authkey, header[:1])[:6]) {
		return nil, errors.New("hmac1 not equal")
	}
	if !bytes.Equal(header[27:], c.v9eMMxHHk(authkey, header[7:27])[:4]) {
		return nil, errors.New("hmac2 not equal")
	}
	data, e := c.GetReadBuf(int(length - 31))
	if e != nil {
		return nil, e
	}
	if _, e = io.ReadFull(c.ReadWriter, data); e != nil {
		return nil, e
	}
	end := len(data) - 4
	if !bytes.Equal(data[end:], c.v9eMMxHHk(c.h6LJlye, header[:7], header[7:], data[:end])[:4]) {
		return nil, errors.New("hmac3 not equal")
	}
	return data[pad:end], nil
}
func (c *lKlYpdomLrhv) ibH0hkF() ([]byte, error) { //0x1cc10e0
	var header [4]byte
	if _, e := io.ReadFull(c.ReadWriter, header[:]); e != nil {
		return nil, e
	}
	length := int(binary.LittleEndian.Uint16(header[:]))
	if c.jJzR7oD6 == nil {
		c.jJzR7oD6 = make([]byte, len(c.h6LJlye)+4)
		copy(c.jJzR7oD6, c.h6LJlye)
	}
	binary.LittleEndian.PutUint32(c.jJzR7oD6[len(c.jJzR7oD6)-4:], c.eOnqFdRmAaYM)
	c.eOnqFdRmAaYM++
	if !bytes.Equal(header[2:], c.v9eMMxHHk(c.jJzR7oD6, header[:2])[:2]) {
		return nil, errors.New("hmac1 error")
	}
	if length < 7 {
		return nil, fmt.Errorf("packet length too small: %v", length)
	}
	if length > 32768 {
		return nil, fmt.Errorf("packet length too large: %v", length)
	}
	data, e := c.GetReadBuf(length - 4)
	if e != nil {
		return nil, e
	}
	if _, e = io.ReadFull(c.ReadWriter, data); e != nil {
		return nil, e
	}
	pad := int(data[0])
	if pad == 255 {
		pad = int(binary.LittleEndian.Uint16(data[1:]))
	}
	end := len(data) - 4
	if !bytes.Equal(data[end:], c.v9eMMxHHk(c.jJzR7oD6, header[:], data[:end])[:4]) {
		return nil, errors.New("hmac2 error")
	}
	if pad == end {
		_, e = c.Write(nil)
		return nil, e
	}
	return data[pad:end], nil
}
func (c *lKlYpdomLrhv) Read(p []byte) (int, error) { //0x1cc1a20
	if !c.m74sJGwc {
		data, e := c.rTiayU()
		if e != nil {
			return 0, e
		}
		c.m74sJGwc = true
		n := copy(p, data)
		if n < len(data) {
			c.qQSF_JOQQp.Write(data[n:])
		}
		return n, nil
	}
	if c.qQSF_JOQQp.Len() > 0 {
		return c.qQSF_JOQQp.Read(p)
	}
	data, e := c.ibH0hkF()
	n := copy(p, data)
	if e != nil {
		return n, e
	}
	c.qQSF_JOQQp.Write(data[n:])
	return n, nil
}
func (c *lKlYpdomLrhv) Write(p []byte) (int, error) { //0x1cc1d00
	if len(p) > 8000 {
		n, e := c.Write(p[:8000])
		if e != nil {
			return n, e
		}
		m, e := c.Write(p[8000:])
		return n + m, e
	}
	data, e := c.bQ9xnsqWMS65(p)
	if e != nil {
		return 0, e
	}
	n, e := c.ReadWriter.Write(data)
	if e != nil {
		return min(n, len(p)), e
	}
	return len(p), nil
}
func (c *lKlYpdomLrhv) WriteTo(p []byte, addr net.Addr) (int, error) { //0x1cc1e40
	size := len(p)
	if cap(p)-len(p) < 4 {
		b, e := c.GetWriteBuf(size + 4)
		if e != nil {
			return 0, e
		}
		copy(b, p)
		p = b[:size]
	}
	tag := c.v9eMMxHHk(c.ggoDIOgz8.IawCjSkuTM4, p[:size])
	p = p[:size+4]
	copy(p[size:], tag[:4])
	_, e := c.WCApOca6IQW7.WriteTo(p, addr)
	if e != nil {
		return 0, e
	}
	return size, nil
}
func (c *lKlYpdomLrhv) WriteToUID(p []byte, a net.Addr, _ int64) (int, error) { return c.WriteTo(p, a) } //0x1cc2060
func (c *lKlYpdomLrhv) ReadFromUID(p []byte) (int, int64, net.Addr, error) { //0x1cc20e0
	n, a, e := c.WCApOca6IQW7.ReadFrom(p)
	if e != nil {
		return 0, 0, nil, e
	}
	if n < 8 {
		return 0, 0, nil, errors.New("auth_aes128_* udp packet len small than 8")
	}
	p = p[:n]
	var uid int64
	var key []byte
	switch c.ggoDIOgz8.NxYA16xMfF {
	case 0:
		uid = c.ggoDIOgz8.Uklg6401
		key = c.ggoDIOgz8.IawCjSkuTM4
	case 1, 2:
		uid = int64(binary.LittleEndian.Uint32(p[n-8:]))
		password, ok := c.ggoDIOgz8.ZlsyKgz4EI.Get(uid)
		if !ok {
			return 0, 0, nil, errors.New("can not find uid: " + strconv.FormatInt(uid, 10))
		}
		key = c.idj1Fac8([]byte(password))
	default:
		return 0, 0, nil, fmt.Errorf("auth_aes not support single port type: %v", c.ggoDIOgz8.NxYA16xMfF)
	}
	c.h6LJlye = key
	if !bytes.Equal(p[n-4:], c.v9eMMxHHk(key, p[:n-4])[:4]) {
		return 0, 0, nil, errors.New("auth_aes128_* udp packet hmac err")
	}
	return n - 8, uid, a, nil
}
