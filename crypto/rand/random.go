// Package uLofn6j restores the executable's cryptographic random integers.
package rand

import (
	"crypto/rand"
	"encoding/binary"
)

// FXllqewWB restores 0xe8e120. Each result consumes eight random bytes.
func FXllqewWB() int32 { return int32(Ctna67jV() >> 32) }

// Ctna67jV restores 0xe8e180.
func Ctna67jV() int64 {
	var b [8]byte
	rand.Read(b[:])
	return int64(binary.BigEndian.Uint64(b[:]) & 0x7fffffffffffffff)
}

// NMT2SOqf restores 0xe8e1e0, including rejection sampling.
func NMT2SOqf(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 {
		return FXllqewWB() & (n - 1)
	}
	max := int32(0x7fffffff - (uint32(1)<<31)%uint32(n))
	v := FXllqewWB()
	for v > max {
		v = FXllqewWB()
	}
	return v % n
}

// TfHpoOcv restores 0xe8e2c0.
func TfHpoOcv(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 {
		return Ctna67jV() & (n - 1)
	}
	max := int64(0x7fffffffffffffff - (uint64(1)<<63)%uint64(n))
	v := Ctna67jV()
	for v > max {
		v = Ctna67jV()
	}
	return v % n
}

// Co6D_zfls restores 0xe8e420; panic strings are decoder-derived.
func Co6D_zfls(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 0x7fffffff {
		return int(NMT2SOqf(int32(n)))
	}
	return int(TfHpoOcv(int64(n)))
}
