/*
Recovered dependency source matched to the supplied ELF: https://github.com/dgryski/go-camellia/tree/69a8a13fb23d261da63db7c927e21e532da9423a
Original embedded revision is not proven. Names/layout mapped from RTTI;
key schedules and encrypt/decrypt results checked against original machine code.

Copyright (c) 2013-2019 Damian Gryski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.


*/
// Package camellia is an implementation of the CAMELLIA encryption algorithm
/*

   This is an unoptimized version based on the description in RFC 3713.

   References:
   http://en.wikipedia.org/wiki/Camellia_%28cipher%29
   https://info.isl.ntt.co.jp/crypt/eng/camellia/
*/
package ubT9IZfVz0MQ

import (
	"crypto/cipher"
	"encoding/binary"
	"math/bits"
	"strconv"
)

const BlockSize = 16

type ZSU8iaw int

func (nad7N9io2tZx ZSU8iaw) Error() string {
	return "camellia: invalid key size " + strconv.Itoa(int(nad7N9io2tZx))
}

type gaC3UK struct {
	qwJx_C       [5]uint64
	nad7N9io2tZx [25]uint64
	__qexTX      [7]uint64
	ay02T8ZK     int
}

const (
	sigma1 = 0xA09E667F3BCC908B
	sigma2 = 0xB67AE8584CAA73B2
	sigma3 = 0xC6EF372FE94F82BE
	sigma4 = 0x54FF53A5F1D36F1C
	sigma5 = 0x10E527FADE682D1D
	sigma6 = 0xB05688C2B3E6C1FD
)

func init() {
	// initialize other sboxes
	for i := range sbox1 {
		sbox2[i] = bits.RotateLeft8(sbox1[i], 1)
		sbox3[i] = bits.RotateLeft8(sbox1[i], 7)
		sbox4[i] = sbox1[bits.RotateLeft8(uint8(i), 1)]
	}
}

func rotl128(nad7N9io2tZx [2]uint64, rot uint) (hi, lo uint64) {

	if rot > 64 {
		rot -= 64
		nad7N9io2tZx[0], nad7N9io2tZx[1] = nad7N9io2tZx[1], nad7N9io2tZx[0]
	}

	t := nad7N9io2tZx[0] >> (64 - rot)
	hi = (nad7N9io2tZx[0] << rot) | (nad7N9io2tZx[1] >> (64 - rot))
	lo = (nad7N9io2tZx[1] << rot) | t
	return hi, lo
}

// New creates and returns a new cipher.Block.
// The key argument should be 16, 24, or 32 bytes.
func RMycR7V(key []byte) (cipher.Block, error) {

	ay02T8ZK := len(key)
	switch ay02T8ZK {
	default:
		return nil, ZSU8iaw(ay02T8ZK)
	case 16, 24, 32:
		break
	}

	var d1, d2 uint64

	var kl [2]uint64
	var kr [2]uint64
	var ka [2]uint64
	var kb [2]uint64

	kl[0] = binary.BigEndian.Uint64(key[0:])
	kl[1] = binary.BigEndian.Uint64(key[8:])

	switch ay02T8ZK {
	case 24:
		kr[0] = binary.BigEndian.Uint64(key[16:])
		kr[1] = ^kr[0]
	case 32:
		kr[0] = binary.BigEndian.Uint64(key[16:])
		kr[1] = binary.BigEndian.Uint64(key[24:])

	}

	d1 = (kl[0] ^ kr[0])
	d2 = (kl[1] ^ kr[1])

	d2 = d2 ^ eSzEzGm(d1, sigma1)
	d1 = d1 ^ eSzEzGm(d2, sigma2)

	d1 = d1 ^ (kl[0])
	d2 = d2 ^ (kl[1])
	d2 = d2 ^ eSzEzGm(d1, sigma3)
	d1 = d1 ^ eSzEzGm(d2, sigma4)
	ka[0] = d1
	ka[1] = d2
	d1 = (ka[0] ^ kr[0])
	d2 = (ka[1] ^ kr[1])
	d2 = d2 ^ eSzEzGm(d1, sigma5)
	d1 = d1 ^ eSzEzGm(d2, sigma6)
	kb[0] = d1
	kb[1] = d2

	// here we generate our keys
	c := new(gaC3UK)

	c.ay02T8ZK = ay02T8ZK

	if ay02T8ZK == 16 {

		c.qwJx_C[1], c.qwJx_C[2] = rotl128(kl, 0)

		c.nad7N9io2tZx[1], c.nad7N9io2tZx[2] = rotl128(ka, 0)
		c.nad7N9io2tZx[3], c.nad7N9io2tZx[4] = rotl128(kl, 15)
		c.nad7N9io2tZx[5], c.nad7N9io2tZx[6] = rotl128(ka, 15)

		c.__qexTX[1], c.__qexTX[2] = rotl128(ka, 30)

		c.nad7N9io2tZx[7], c.nad7N9io2tZx[8] = rotl128(kl, 45)
		c.nad7N9io2tZx[9], _ = rotl128(ka, 45)
		_, c.nad7N9io2tZx[10] = rotl128(kl, 60)
		c.nad7N9io2tZx[11], c.nad7N9io2tZx[12] = rotl128(ka, 60)

		c.__qexTX[3], c.__qexTX[4] = rotl128(kl, 77)

		c.nad7N9io2tZx[13], c.nad7N9io2tZx[14] = rotl128(kl, 94)
		c.nad7N9io2tZx[15], c.nad7N9io2tZx[16] = rotl128(ka, 94)
		c.nad7N9io2tZx[17], c.nad7N9io2tZx[18] = rotl128(kl, 111)

		c.qwJx_C[3], c.qwJx_C[4] = rotl128(ka, 111)

	} else {
		// 24 or 32

		c.qwJx_C[1], c.qwJx_C[2] = rotl128(kl, 0)

		c.nad7N9io2tZx[1], c.nad7N9io2tZx[2] = rotl128(kb, 0)
		c.nad7N9io2tZx[3], c.nad7N9io2tZx[4] = rotl128(kr, 15)
		c.nad7N9io2tZx[5], c.nad7N9io2tZx[6] = rotl128(ka, 15)

		c.__qexTX[1], c.__qexTX[2] = rotl128(kr, 30)

		c.nad7N9io2tZx[7], c.nad7N9io2tZx[8] = rotl128(kb, 30)
		c.nad7N9io2tZx[9], c.nad7N9io2tZx[10] = rotl128(kl, 45)
		c.nad7N9io2tZx[11], c.nad7N9io2tZx[12] = rotl128(ka, 45)

		c.__qexTX[3], c.__qexTX[4] = rotl128(kl, 60)

		c.nad7N9io2tZx[13], c.nad7N9io2tZx[14] = rotl128(kr, 60)
		c.nad7N9io2tZx[15], c.nad7N9io2tZx[16] = rotl128(kb, 60)
		c.nad7N9io2tZx[17], c.nad7N9io2tZx[18] = rotl128(kl, 77)

		c.__qexTX[5], c.__qexTX[6] = rotl128(ka, 77)

		c.nad7N9io2tZx[19], c.nad7N9io2tZx[20] = rotl128(kr, 94)
		c.nad7N9io2tZx[21], c.nad7N9io2tZx[22] = rotl128(ka, 94)
		c.nad7N9io2tZx[23], c.nad7N9io2tZx[24] = rotl128(kl, 111)

		c.qwJx_C[3], c.qwJx_C[4] = rotl128(kb, 111)
	}

	return c, nil
}

func (c *gaC3UK) Encrypt(dst, src []byte) {

	d1 := binary.BigEndian.Uint64(src[0:])
	d2 := binary.BigEndian.Uint64(src[8:])

	d1 ^= c.qwJx_C[1]
	d2 ^= c.qwJx_C[2]

	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[1])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[2])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[3])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[4])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[5])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[6])

	d1 = fl(d1, c.__qexTX[1])
	d2 = flinv(d2, c.__qexTX[2])

	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[7])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[8])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[9])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[10])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[11])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[12])

	d1 = fl(d1, c.__qexTX[3])
	d2 = flinv(d2, c.__qexTX[4])

	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[13])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[14])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[15])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[16])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[17])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[18])

	if c.ay02T8ZK > 16 {
		// 24 or 32

		d1 = fl(d1, c.__qexTX[5])
		d2 = flinv(d2, c.__qexTX[6])

		d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[19])
		d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[20])
		d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[21])
		d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[22])
		d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[23])
		d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[24])
	}

	d2 = d2 ^ c.qwJx_C[3]
	d1 = d1 ^ c.qwJx_C[4]

	binary.BigEndian.PutUint64(dst[0:], d2)
	binary.BigEndian.PutUint64(dst[8:], d1)
}

func (c *gaC3UK) Decrypt(dst, src []byte) {

	d2 := binary.BigEndian.Uint64(src[0:])
	d1 := binary.BigEndian.Uint64(src[8:])

	d1 = d1 ^ c.qwJx_C[4]
	d2 = d2 ^ c.qwJx_C[3]

	if c.ay02T8ZK > 16 {
		// 24 or 32

		d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[24])
		d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[23])
		d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[22])
		d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[21])
		d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[20])
		d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[19])

		d2 = fl(d2, c.__qexTX[6])
		d1 = flinv(d1, c.__qexTX[5])
	}

	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[18])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[17])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[16])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[15])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[14])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[13])

	d2 = fl(d2, c.__qexTX[4])
	d1 = flinv(d1, c.__qexTX[3])

	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[12])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[11])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[10])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[9])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[8])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[7])

	d2 = fl(d2, c.__qexTX[2])
	d1 = flinv(d1, c.__qexTX[1])

	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[6])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[5])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[4])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[3])
	d1 = d1 ^ eSzEzGm(d2, c.nad7N9io2tZx[2])
	d2 = d2 ^ eSzEzGm(d1, c.nad7N9io2tZx[1])

	d2 ^= c.qwJx_C[2]
	d1 ^= c.qwJx_C[1]

	binary.BigEndian.PutUint64(dst[0:], d1)
	binary.BigEndian.PutUint64(dst[8:], d2)
}

func (c *gaC3UK) BlockSize() int {
	return BlockSize
}

func eSzEzGm(fin, __qexTX uint64) uint64 {
	var x uint64
	x = fin ^ __qexTX
	t1 := sbox1[uint8(x>>56)]
	t2 := sbox2[uint8(x>>48)]
	t3 := sbox3[uint8(x>>40)]
	t4 := sbox4[uint8(x>>32)]
	t5 := sbox2[uint8(x>>24)]
	t6 := sbox3[uint8(x>>16)]
	t7 := sbox4[uint8(x>>8)]
	t8 := sbox1[uint8(x)]
	y1 := t1 ^ t3 ^ t4 ^ t6 ^ t7 ^ t8
	y2 := t1 ^ t2 ^ t4 ^ t5 ^ t7 ^ t8
	y3 := t1 ^ t2 ^ t3 ^ t5 ^ t6 ^ t8
	y4 := t2 ^ t3 ^ t4 ^ t5 ^ t6 ^ t7
	y5 := t1 ^ t2 ^ t6 ^ t7 ^ t8
	y6 := t2 ^ t3 ^ t5 ^ t7 ^ t8
	y7 := t3 ^ t4 ^ t5 ^ t6 ^ t8
	y8 := t1 ^ t4 ^ t5 ^ t6 ^ t7
	return uint64(y1)<<56 | uint64(y2)<<48 | uint64(y3)<<40 | uint64(y4)<<32 | uint64(y5)<<24 | uint64(y6)<<16 | uint64(y7)<<8 | uint64(y8)
}

func fl(flin, __qexTX uint64) uint64 {
	x1 := uint32(flin >> 32)
	x2 := uint32(flin & 0xffffffff)
	k1 := uint32(__qexTX >> 32)
	k2 := uint32(__qexTX & 0xffffffff)
	x2 = x2 ^ bits.RotateLeft32(x1&k1, 1)
	x1 = x1 ^ (x2 | k2)
	return uint64(x1)<<32 | uint64(x2)
}

func flinv(flin, __qexTX uint64) uint64 {
	y1 := uint32(flin >> 32)
	y2 := uint32(flin & 0xffffffff)
	k1 := uint32(__qexTX >> 32)
	k2 := uint32(__qexTX & 0xffffffff)
	y1 = y1 ^ (y2 | k2)
	y2 = y2 ^ bits.RotateLeft32(y1&k1, 1)
	return uint64(y1)<<32 | uint64(y2)
}

var sbox1 = [...]byte{
	0x70, 0x82, 0x2c, 0xec, 0xb3, 0x27, 0xc0, 0xe5, 0xe4, 0x85, 0x57, 0x35, 0xea, 0x0c, 0xae, 0x41,
	0x23, 0xef, 0x6b, 0x93, 0x45, 0x19, 0xa5, 0x21, 0xed, 0x0e, 0x4f, 0x4e, 0x1d, 0x65, 0x92, 0xbd,
	0x86, 0xb8, 0xaf, 0x8f, 0x7c, 0xeb, 0x1f, 0xce, 0x3e, 0x30, 0xdc, 0x5f, 0x5e, 0xc5, 0x0b, 0x1a,
	0xa6, 0xe1, 0x39, 0xca, 0xd5, 0x47, 0x5d, 0x3d, 0xd9, 0x01, 0x5a, 0xd6, 0x51, 0x56, 0x6c, 0x4d,
	0x8b, 0x0d, 0x9a, 0x66, 0xfb, 0xcc, 0xb0, 0x2d, 0x74, 0x12, 0x2b, 0x20, 0xf0, 0xb1, 0x84, 0x99,
	0xdf, 0x4c, 0xcb, 0xc2, 0x34, 0x7e, 0x76, 0x05, 0x6d, 0xb7, 0xa9, 0x31, 0xd1, 0x17, 0x04, 0xd7,
	0x14, 0x58, 0x3a, 0x61, 0xde, 0x1b, 0x11, 0x1c, 0x32, 0x0f, 0x9c, 0x16, 0x53, 0x18, 0xf2, 0x22,
	0xfe, 0x44, 0xcf, 0xb2, 0xc3, 0xb5, 0x7a, 0x91, 0x24, 0x08, 0xe8, 0xa8, 0x60, 0xfc, 0x69, 0x50,
	0xaa, 0xd0, 0xa0, 0x7d, 0xa1, 0x89, 0x62, 0x97, 0x54, 0x5b, 0x1e, 0x95, 0xe0, 0xff, 0x64, 0xd2,
	0x10, 0xc4, 0x00, 0x48, 0xa3, 0xf7, 0x75, 0xdb, 0x8a, 0x03, 0xe6, 0xda, 0x09, 0x3f, 0xdd, 0x94,
	0x87, 0x5c, 0x83, 0x02, 0xcd, 0x4a, 0x90, 0x33, 0x73, 0x67, 0xf6, 0xf3, 0x9d, 0x7f, 0xbf, 0xe2,
	0x52, 0x9b, 0xd8, 0x26, 0xc8, 0x37, 0xc6, 0x3b, 0x81, 0x96, 0x6f, 0x4b, 0x13, 0xbe, 0x63, 0x2e,
	0xe9, 0x79, 0xa7, 0x8c, 0x9f, 0x6e, 0xbc, 0x8e, 0x29, 0xf5, 0xf9, 0xb6, 0x2f, 0xfd, 0xb4, 0x59,
	0x78, 0x98, 0x06, 0x6a, 0xe7, 0x46, 0x71, 0xba, 0xd4, 0x25, 0xab, 0x42, 0x88, 0xa2, 0x8d, 0xfa,
	0x72, 0x07, 0xb9, 0x55, 0xf8, 0xee, 0xac, 0x0a, 0x36, 0x49, 0x2a, 0x68, 0x3c, 0x38, 0xf1, 0xa4,
	0x40, 0x28, 0xd3, 0x7b, 0xbb, 0xc9, 0x43, 0xc1, 0x15, 0xe3, 0xad, 0xf4, 0x77, 0xc7, 0x80, 0x9e,
}

var sbox2 [256]byte
var sbox3 [256]byte
var sbox4 [256]byte

var _ cipher.Block = &gaC3UK{}
