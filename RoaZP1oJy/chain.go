package RoaZP1oJy

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rc4"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	Clock "recovered.local/soga/CAodbhZf"
	P "recovered.local/soga/D_zCag1qTV"
	H "recovered.local/soga/JafvCs5"
	T "recovered.local/soga/TaFdeUcnzQ"
	S "recovered.local/soga/ad5KEbw_EzWk"
	L "recovered.local/soga/muc279Mpa6S4"
)

// Constructors 0x1ccc020, 0x1cd8da0, 0x1cd9ae0, 0x1cda6e0, 0x1cdb660,
// 0x1cdbc60. Each replaces the embedded padding strategy with its own type.
// Salts are independently decoded from the corresponding original call sites.
func oPPzNk(c *T.KrOWYtzUYft) Vzuv_s9ZS5 {
	v := &h2QD2cI1{ReadWriter: c, WCApOca6IQW7: c, qjWQaiA6: "auth_chain_a", aBfz5qx5HJ7: 1, aL42JWaXIu: 1}
	v.d_JgwRxIbD4o = v
	return v
}
func gYQ1Wtm(c *T.KrOWYtzUYft) Vzuv_s9ZS5 {
	v := &yeGSuIuC{h2QD2cI1: oPPzNk(c).(*h2QD2cI1)}
	v.d_JgwRxIbD4o = v
	v.qjWQaiA6 = "auth_chain_b"
	return v
}
func vHtqJ7xSQiSj(c *T.KrOWYtzUYft) Vzuv_s9ZS5 {
	v := &j8lPExkSUzBF{gYQ1Wtm(c).(*yeGSuIuC)}
	v.d_JgwRxIbD4o = v
	v.qjWQaiA6 = "auth_chain_c"
	return v
}
func daajdwCbv7(c *T.KrOWYtzUYft) Vzuv_s9ZS5 {
	v := &gdcp8ElhzO{gYQ1Wtm(c).(*yeGSuIuC)}
	v.d_JgwRxIbD4o = v
	v.qjWQaiA6 = "auth_chain_d"
	return v
}
func xrwCp6(c *T.KrOWYtzUYft) Vzuv_s9ZS5 {
	v := &yMywJL{daajdwCbv7(c).(*gdcp8ElhzO)}
	v.d_JgwRxIbD4o = v
	v.qjWQaiA6 = "auth_chain_e"
	return v
}
func hH3aFx(c *T.KrOWYtzUYft) Vzuv_s9ZS5 {
	v := &pMrxs3_Z{yMywJL: xrwCp6(c).(*yMywJL)}
	v.d_JgwRxIbD4o = v
	v.qjWQaiA6 = "auth_chain_f"
	return v
}
func (c *h2QD2cI1) SetServerInfo(info *S.L7E7H2aW8KVs) { c.oHHJQje = info }   //0x1ccc140
func (c *h2QD2cI1) CurUID() int64                      { return c.sKkAQIIlO } //0x1ccf880
func (c *h2QD2cI1) CID() int64                         { return c.ahwJUuyQ }  //0x1ccf8a0
func (*h2QD2cI1) GetOverhead() int                     { return 4 }           //0x1ccf8c0

// hAt43dMn3W restores the original 36-byte server authentication at 0x1ccc720.
// All three user modes, intermediate state writes, and error/log side effects
// are retained. The two stream ciphers are independent instances with one key.
func (c *h2QD2cI1) hAt43dMn3W() error {
	header := make([]byte, 36)
	if _, err := io.ReadFull(c.ReadWriter, header); err != nil {
		return err
	}
	info := c.oHHJQje
	authKey := make([]byte, info.SP2KC4lkkl3C+info.DvdB7JZX)
	copy(authKey, info.Ds7eTY7Qp66N)
	copy(authKey[info.DvdB7JZX:], info.IawCjSkuTM4)
	c.zOQ_6vA = c.hmcSfmiGu(authKey, header[:4])
	if !bytes.Equal(c.zOQ_6vA[:8], header[4:12]) {
		return errors.New("hmac1 error")
	}
	var userKey []byte
	switch info.NxYA16xMfF {
	case 0:
		c.sKkAQIIlO = info.Uklg6401
		userKey = info.IawCjSkuTM4
	case 1, 2:
		c.sKkAQIIlO = int64(binary.LittleEndian.Uint32(c.zOQ_6vA[8:]) ^ binary.LittleEndian.Uint32(header[12:]))
		password, ok := info.ZlsyKgz4EI.Get(c.sKkAQIIlO)
		if !ok {
			return fmt.Errorf("can not find uid: %v", c.sKkAQIIlO)
		}
		userKey = []byte(password)
	default:
		return fmt.Errorf("auth_chain_* not support single port type: %v", info.NxYA16xMfF)
	}
	c.bZq5ndT = make([]byte, len(userKey))
	copy(c.bZq5ndT, userKey)
	c.emo_yYrLZ0DL = c.hmcSfmiGu(c.bZq5ndT, header[12:32])
	if !bytes.Equal(c.emo_yYrLZ0DL[:4], header[32:]) {
		return errors.New("hmac2 error")
	}
	encoded := base64.StdEncoding.EncodeToString(c.bZq5ndT)
	block, err := aes.NewCipher(H.Is8IqhjVE6(encoded+c.qjWQaiA6, 16))
	if err != nil {
		return fmt.Errorf("creating aes cipher err: %v", err)
	}
	decoded := make([]byte, 16)
	cipher.NewCBCDecrypter(block, make([]byte, 16)).CryptBlocks(decoded, header[16:32])
	timestamp := binary.LittleEndian.Uint32(decoded)
	c.ahwJUuyQ = int64(binary.LittleEndian.Uint32(decoded[4:]))
	c.mFh_3tzSZh = int(binary.LittleEndian.Uint16(decoded[12:]))
	now := uint32(Clock.C5Jff5vs0())
	if timestamp < now-86400 || timestamp > now+86400 {
		return errors.New("utc invalid")
	}
	c.h9waQgLNy = make([]byte, len(c.bZq5ndT)+4)
	copy(c.h9waQgLNy, c.bZq5ndT)
	c.fOtNm3lE9H5 = make([]byte, len(c.bZq5ndT)+4)
	copy(c.fOtNm3lE9H5, c.bZq5ndT)
	key := H.Is8IqhjVE6(encoded+base64.StdEncoding.EncodeToString(c.zOQ_6vA), 16)
	enc, err := rc4.NewCipher(key)
	if err != nil {
		L.TRTbrquD(slog.LevelDebug, "new rc4 cipher err:", err)
		return err
	}
	dec, err := rc4.NewCipher(key)
	if err != nil {
		L.TRTbrquD(slog.LevelDebug, "new rc4 cipher err:", err)
		return err
	}
	c.pAvqa5e1EyuH = enc
	c.imP42RHIDKk = dec
	return nil
}

// ibH0hkF restores 0x1ccd5e0. Counter advancement occurs before checksum
// validation; buffered plaintext aliases the owned read buffer.
func (c *h2QD2cI1) ibH0hkF(p []byte) (int, error) {
	if c.gVeP6b != nil && c.gVeP6b.Len() > 0 {
		return c.gVeP6b.Read(p)
	}
	header := make([]byte, 2)
	if _, err := io.ReadFull(c.ReadWriter, header); err != nil {
		return 0, err
	}
	length := int(binary.LittleEndian.Uint16(c.zOQ_6vA[14:16]) ^ binary.LittleEndian.Uint16(header))
	padding := c.d_JgwRxIbD4o.z0O7jemcj(length, c.zOQ_6vA, &c.zg86iGzd)
	total := length + padding
	if total >= 4096 {
		L.TRTbrquD(slog.LevelDebug, "auth_chain_* over size")
		return 0, errors.New("auth_chain_* over size")
	}
	body, err := c.GetReadBuf(total + 2)
	if err != nil {
		return 0, err
	}
	if _, err = io.ReadFull(c.ReadWriter, body); err != nil {
		return 0, err
	}
	binary.LittleEndian.PutUint32(c.h9waQgLNy[len(c.h9waQgLNy)-4:], c.aBfz5qx5HJ7)
	c.aBfz5qx5HJ7++
	mac := c.hmcSfmiGu(c.h9waQgLNy, header, body[:total])
	if !bytes.Equal(mac[:2], body[total:]) {
		return 0, errors.New("hmac wrong")
	}
	offset := 0
	if length != 0 && padding > 0 {
		offset = int(c.zg86iGzd.Next() % 0x200000011 % uint64(padding))
	}
	plaintext := body[offset : offset+length]
	c.imP42RHIDKk.XORKeyStream(plaintext, plaintext)
	c.zOQ_6vA = mac
	n := copy(p, plaintext)
	if n < length {
		c.gVeP6b = bytes.NewReader(plaintext[n:])
	}
	return n, nil
}
func (c *h2QD2cI1) Read(p []byte) (int, error) { //0x1ccdee0
	if !c.rmErDQPkS31O {
		if err := c.hAt43dMn3W(); err != nil {
			return 0, err
		}
		c.rmErDQPkS31O = true
	}
	return c.ibH0hkF(p)
}

// xnIaLlgPST restores 0x1ccdfa0. Padding entropy errors are ignored, just
// as in the original; ciphertext length is masked by the previous MAC.
func (c *h2QD2cI1) xnIaLlgPST(p []byte) ([]byte, error) {
	padding := c.d_JgwRxIbD4o.z0O7jemcj(len(p), c.emo_yYrLZ0DL, &c.ydRwMfttAwX)
	b, err := c.GetWriteBuf(len(p) + padding + 4)
	if err != nil {
		return nil, err
	}
	if len(p) > 0 {
		offset := 0
		if padding > 0 {
			offset = int(c.ydRwMfttAwX.Next() % 0x200000011 % uint64(padding))
			rand.Read(b[2 : offset+2])
			rand.Read(b[offset+len(p)+2 : len(b)-2])
		}
		c.pAvqa5e1EyuH.XORKeyStream(b[offset+2:], p)
	} else {
		rand.Read(b[2 : len(b)-2])
	}
	binary.LittleEndian.PutUint16(b, binary.LittleEndian.Uint16(c.emo_yYrLZ0DL[14:])^uint16(len(p)))
	binary.LittleEndian.PutUint32(c.fOtNm3lE9H5[len(c.fOtNm3lE9H5)-4:], c.aL42JWaXIu)
	c.aL42JWaXIu++
	mac := c.hmcSfmiGu(c.fOtNm3lE9H5, b[:len(b)-2])
	copy(b[len(b)-2:], mac[:2])
	c.emo_yYrLZ0DL = mac
	return b, nil
}
func (c *h2QD2cI1) aYEomSvMN0RG(p []byte) (int, error) { //0x1cce4a0
	if len(p) > c.jLK6jCJ00azL {
		n, err := c.aYEomSvMN0RG(p[:c.jLK6jCJ00azL])
		if err != nil {
			return n, err
		}
		m, err := c.aYEomSvMN0RG(p[c.jLK6jCJ00azL:])
		return n + m, err
	}
	b, err := c.xnIaLlgPST(p)
	if err != nil {
		return 0, err
	}
	if _, err = c.ReadWriter.Write(b); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (c *h2QD2cI1) Write(p []byte) (int, error) { //0x1cce5e0
	if c.aL42JWaXIu != 1 {
		return c.aYEomSvMN0RG(p)
	}
	originalLen := len(p)
	c.jLK6jCJ00azL = 1460 - c.mFh_3tzSZh
	var b []byte
	if cap(p) < len(p)+2 {
		b = P.PoolAt50961D8.Get(len(p) + 2)
		defer P.PoolAt50961D8.Put(b)
		copy(b[2:], p)
	} else {
		b = p[:len(p)+2]
		for i := len(p) - 1; i >= 0; i-- {
			b[i+2] = b[i]
		}
	}
	binary.LittleEndian.PutUint16(b, 1460)
	if _, err := c.aYEomSvMN0RG(b); err != nil {
		return 0, err
	}
	return originalLen, nil
}

// WriteTo is deliberately the original panic rather than an alias to
// WriteToUID (0x1ccf120; decoded message at 0x1ccf152).
func (*h2QD2cI1) WriteTo([]byte, net.Addr) (int, error) { panic("please use WriteToUID()") }
func (c *h2QD2cI1) ReadFromUID(p []byte) (int, int64, net.Addr, error) { //0x1cce8a0
	n, addr, err := c.WCApOca6IQW7.ReadFrom(p)
	if err != nil {
		return 0, 0, nil, err
	}
	if n < 8 {
		return 0, 0, nil, errors.New("ssr udp packet size small than 8")
	}
	b := p[:n]
	mac := c.hmcSfmiGu(c.oHHJQje.IawCjSkuTM4, b[n-8:n-5])
	uid := int64(binary.LittleEndian.Uint32(mac) ^ binary.LittleEndian.Uint32(b[n-5:]))
	var key []byte
	switch c.oHHJQje.NxYA16xMfF {
	case 0:
		uid = c.oHHJQje.Uklg6401
		key = c.oHHJQje.IawCjSkuTM4
	case 1, 2:
		password, ok := c.oHHJQje.ZlsyKgz4EI.Get(uid)
		if !ok {
			return 0, 0, nil, fmt.Errorf("can not find uid: %v", uid)
		}
		key = []byte(password)
	default:
		return 0, 0, nil, fmt.Errorf("auth_chain_* not support single port type: %v", c.oHHJQje.NxYA16xMfF)
	}
	if c.hmcSfmiGu(key, b[:n-1])[0] != b[n-1] {
		return 0, 0, nil, errors.New("auth_chain_* udp packet hmac err")
	}
	padding := c.vuOMfc(mac, &c.zg86iGzd)
	stream, err := rc4.NewCipher(H.Is8IqhjVE6(base64.StdEncoding.EncodeToString(key)+base64.StdEncoding.EncodeToString(mac), 16))
	if err != nil {
		return 0, 0, nil, err
	}
	length := n - padding - 8
	if length < 0 {
		return 0, 0, nil, fmt.Errorf("rand length error")
	}
	stream.XORKeyStream(b, b[:length])
	return length, uid, addr, nil
}
func (c *h2QD2cI1) WriteToUID(p []byte, addr net.Addr, uid int64) (int, error) { //0x1ccf180
	var key []byte
	switch c.oHHJQje.NxYA16xMfF {
	case 0:
		key = c.oHHJQje.IawCjSkuTM4
	case 1, 2:
		password, ok := c.oHHJQje.ZlsyKgz4EI.Get(uid)
		if !ok {
			return 0, fmt.Errorf("can not find uid: %v", uid)
		}
		key = []byte(password)
	default:
		return 0, fmt.Errorf("auth_chain_* not support single port type: %v", c.oHHJQje.NxYA16xMfF)
	}
	var auth [7]byte
	rand.Read(auth[:])
	mac := c.hmcSfmiGu(c.oHHJQje.IawCjSkuTM4, auth[:])
	padding := c.vuOMfc(mac, &c.ydRwMfttAwX)
	stream, err := rc4.NewCipher(H.Is8IqhjVE6(base64.StdEncoding.EncodeToString(key)+base64.StdEncoding.EncodeToString(mac), 16))
	if err != nil {
		return 0, err
	}
	b, err := c.GetWriteBuf(len(p) + padding + 8)
	if err != nil {
		return 0, err
	}
	stream.XORKeyStream(b, p)
	rand.Read(b[len(p) : len(p)+padding])
	copy(b[len(p)+padding:], auth[:])
	copy(b[len(b)-1:], c.hmcSfmiGu(key, b[:len(b)-1]))
	if _, err = c.WCApOca6IQW7.WriteTo(b, addr); err != nil {
		return 0, err
	}
	return len(p), nil
}
