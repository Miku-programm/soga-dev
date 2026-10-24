package FLhkKwJ8M

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"hash/fnv"
	"io"
	mathrand "math/rand"
	"recovered.local/soga/Hx04IbKk"
	"time"
)

// nhGe6Mb_IeZ: 0x19aec20. The random-padding read and short nil-error write
// outcomes are ignored in the original. The request bytes still carry a FNV checksum.
func (c *HRzHVG1yP5O) nhGe6Mb_IeZ() error {
	now := time.Now()
	wire := Hx04IbKk.BaDekw7qe{}
	if !c.tu04N8lBj6I {
		h := hmac.New(md5.New, c.fqAsdD3hjU.AJxI2iBn[:])
		binary.Write(h, binary.BigEndian, uint64(now.Unix()))
		auth := h.Sum(nil)
		pos := wire.Grow(len(auth))
		copy(wire[pos:], auth)
	}
	header := Hx04IbKk.BaDekw7qe{}
	pos := header.Grow(1)
	header[pos] = 1
	for _, part := range [][]byte{c.c646iDyBoW, c._7ezur1gq_ZX} {
		pos = header.Grow(len(part))
		copy(header[pos:pos+len(part)], part)
	}
	pos = header.Grow(1)
	header[pos] = c.kONmi_w_yL3D
	pos = header.Grow(1)
	header[pos] = c.kljQIi6N
	padding := mathrand.Intn(16)
	pos = header.Grow(1)
	header[pos] = c.iGH6ra | byte(padding<<4)
	pos = header.Grow(1)
	header[pos] = 0
	pos = header.Grow(1)
	if c.mNrCvc.GV3P5rY {
		header[pos] = 2
	} else {
		header[pos] = 1
	}
	pos = header.Grow(2)
	binary.BigEndian.PutUint16(header[pos:pos+2], uint16(c.mNrCvc.GHP6ivev2Db))
	pos = header.Grow(1)
	header[pos] = c.mNrCvc.OyEjzi1v93
	pos = header.Grow(len(c.mNrCvc.TIBiNxg))
	copy(header[pos:], c.mNrCvc.TIBiNxg)
	if padding > 0 {
		header.ReadFull(rand.Reader, padding)
	}
	h := fnv.New32a()
	h.Write(header)
	sum := h.Sum32()
	pos = header.Grow(4)
	binary.BigEndian.PutUint32(header[pos:pos+4], sum)
	if c.tu04N8lBj6I {
		var key [16]byte
		copy(key[:], c.fqAsdD3hjU.P3zlkY2W1n)
		_, err := c.DrpXao.Write(s9ADKgL(key, header, now))
		return err
	}
	block, err := aes.NewCipher(c.fqAsdD3hjU.P3zlkY2W1n)
	if err != nil {
		return err
	}
	cipher.NewCFBEncrypter(block, iAK7_ZD8(now)).XORKeyStream(header, header)
	pos = wire.Grow(len(header))
	copy(wire[pos:], header)
	_, err = c.DrpXao.Write(wire)
	return err
}

// dCFgAbCz: 0x19af5e0. The encrypted response length adds 16 as uint16,
// including wraparound. Only auth byte and dynamic-command byte are validated.
func (c *HRzHVG1yP5O) dCFgAbCz() error {
	var response []byte
	if !c.tu04N8lBj6I {
		block, err := aes.NewCipher(c.z1TtPlbnfnl)
		if err != nil {
			return err
		}
		stream := cipher.NewCFBDecrypter(block, c.f9_fTlq)
		response = make([]byte, 4)
		if _, err = io.ReadFull(c.DrpXao, response); err != nil {
			return err
		}
		stream.XORKeyStream(response, response)
	} else {
		key := xXVvCU5cjuB(c.z1TtPlbnfnl, "AEAD Resp Header Len Key")[:16]
		iv := xXVvCU5cjuB(c.f9_fTlq, "AEAD Resp Header Len IV")[:12]
		block, _ := aes.NewCipher(key)
		aead, _ := cipher.NewGCM(block)
		encryptedLength := make([]byte, 18)
		if _, err := io.ReadFull(c.DrpXao, encryptedLength); err != nil {
			return err
		}
		length, err := aead.Open(nil, iv, encryptedLength, nil)
		if err != nil {
			return err
		}
		size := binary.BigEndian.Uint16(length)
		key = xXVvCU5cjuB(c.z1TtPlbnfnl, "AEAD Resp Header Key")[:16]
		iv = xXVvCU5cjuB(c.f9_fTlq, "AEAD Resp Header IV")[:12]
		block, _ = aes.NewCipher(key)
		aead, _ = cipher.NewGCM(block)
		encryptedPayload := make([]byte, size+uint16(16))
		if _, err = io.ReadFull(c.DrpXao, encryptedPayload); err != nil {
			return err
		}
		response, err = aead.Open(nil, iv, encryptedPayload, nil)
		if err != nil {
			return err
		}
		if len(response) < 4 {
			return errors.New("unexpected buffer length")
		}
	}
	if response[0] != c.kONmi_w_yL3D {
		return errors.New("unexpected response header")
	}
	if response[2] != 0 {
		return errors.New("dynamic port is not supported now")
	}
	return nil
}
