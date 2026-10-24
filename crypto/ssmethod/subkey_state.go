package ssmethod

import (
	"bytes"
	"crypto/sha1"
	"encoding"
	"hash"
	"lukechampine.com/blake3"
)

// MXXLkaPB retains all original cached SHA-1 states and buffers (352 bytes).
type MXXLkaPB struct {
	gRldz9mNt    hash.Hash
	xMuBKew5Sx   hash.Hash
	b8aGJ4aIvJE  []byte
	zW1JAGMGD    []byte
	vjV_TXt9     []byte
	sa6pd3D      [20]byte
	fO1kMPCd0gG  [20]byte
	rYY2XWXjYjq7 [20]byte
	gQWRfV0dgL   [32]byte
	fN3a3_       [64]byte
	aLAk_YZ9     [64]byte
	lGjGYnir7l_L []byte
}

func SaYhx5xog9() *MXXLkaPB { // 0x10a6be0
	return &MXXLkaPB{gRldz9mNt: sha1.New(), xMuBKew5Sx: sha1.New()}
}
func (s *MXXLkaPB) KeyBuf(n int) []byte { return s.gQWRfV0dgL[:n] } // 0x10a6ce0
func (s *MXXLkaPB) eO4Tdy(dst *[64]byte, key []byte, pad byte) { // 0x10a6d20
	clear(dst[:])
	if len(key) > 64 {
		s.gRldz9mNt.Reset()
		s.gRldz9mNt.Write(key)
		copy(dst[:], s.gRldz9mNt.Sum(s.fO1kMPCd0gG[:0]))
	} else {
		copy(dst[:], key)
	}
	for i := range dst {
		dst[i] ^= pad
	}
}
func (s *MXXLkaPB) wJRMV1dV(salt []byte) { // 0x10a6ea0
	if s.zW1JAGMGD != nil && bytes.Equal(s.b8aGJ4aIvJE, salt) {
		return
	}
	s.b8aGJ4aIvJE = append(s.b8aGJ4aIvJE[:0], salt...)
	s.eO4Tdy(&s.fN3a3_, salt, 0x36)
	s.gRldz9mNt.Reset()
	s.gRldz9mNt.Write(s.fN3a3_[:])
	var err error
	s.zW1JAGMGD, err = s.gRldz9mNt.(encoding.BinaryAppender).AppendBinary(s.zW1JAGMGD[:0])
	if err != nil {
		panic(err)
	}
	s.eO4Tdy(&s.aLAk_YZ9, salt, 0x5c)
	s.xMuBKew5Sx.Reset()
	s.xMuBKew5Sx.Write(s.aLAk_YZ9[:])
	s.vjV_TXt9, err = s.xMuBKew5Sx.(encoding.BinaryAppender).AppendBinary(s.vjV_TXt9[:0])
	if err != nil {
		panic(err)
	}
}
func (s *MXXLkaPB) b0pUkTAq(key []byte) { // 0x10a72a0
	if err := s.gRldz9mNt.(encoding.BinaryUnmarshaler).UnmarshalBinary(s.zW1JAGMGD); err != nil {
		panic(err)
	}
	s.gRldz9mNt.Write(key)
	inner := s.gRldz9mNt.Sum(s.fO1kMPCd0gG[:0])
	if err := s.xMuBKew5Sx.(encoding.BinaryUnmarshaler).UnmarshalBinary(s.vjV_TXt9); err != nil {
		panic(err)
	}
	s.xMuBKew5Sx.Write(inner)
	copy(s.sa6pd3D[:], s.xMuBKew5Sx.Sum(s.fO1kMPCd0gG[:0]))
}
func (s *MXXLkaPB) dHQaCID_x(dst, info []byte) { // 0x10a7580
	s.eO4Tdy(&s.fN3a3_, s.sa6pd3D[:], 0x36)
	s.eO4Tdy(&s.aLAk_YZ9, s.sa6pd3D[:], 0x5c)
	var counter [1]byte
	var previous []byte
	for offset := 0; offset < len(dst); {
		counter[0]++ // Original counter wraps; there is no HKDF length rejection.
		s.gRldz9mNt.Reset()
		s.gRldz9mNt.Write(s.fN3a3_[:])
		s.gRldz9mNt.Write(previous)
		s.gRldz9mNt.Write(info)
		s.gRldz9mNt.Write(counter[:])
		inner := s.gRldz9mNt.Sum(s.fO1kMPCd0gG[:0])
		s.xMuBKew5Sx.Reset()
		s.xMuBKew5Sx.Write(s.aLAk_YZ9[:])
		s.xMuBKew5Sx.Write(inner)
		previous = s.xMuBKew5Sx.Sum(s.rYY2XWXjYjq7[:0])
		offset += copy(dst[offset:], previous)
	}
}
func (s *MXXLkaPB) oE369eL1UY(dst, key, salt []byte) { // 0x10a7920
	s.lGjGYnir7l_L = append(s.lGjGYnir7l_L[:0], key...)
	s.lGjGYnir7l_L = append(s.lGjGYnir7l_L, salt...)
	blake3.DeriveKey(dst, "shadowsocks 2022 session subkey", s.lGjGYnir7l_L)
}
