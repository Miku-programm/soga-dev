package bQxeORT

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
	"lukechampine.com/blake3"
)

// RaS2myUz retains the original 32-byte wrapper and 96-bit nonce counter.
type RaS2myUz struct {
	cipher.AEAD
	NQGg9x [12]byte
}

// WrnUSnC6 recovers 0x1cf7ee0. Context bytes become the BLAKE3 context string.
func WrnUSnC6(context, material []byte, useAES bool) *RaS2myUz {
	var key [32]byte
	blake3.DeriveKey(key[:], string(context), material)
	var aead cipher.AEAD
	if useAES {
		block, _ := aes.NewCipher(key[:])
		aead, _ = cipher.NewGCM(block)
	} else {
		aead, _ = chacha20poly1305.New(key[:])
	}
	return &RaS2myUz{AEAD: aead}
}

// Seal and Open preserve the pre-call increment, including failed Open calls.
func (a *RaS2myUz) Seal(dst, nonce, plaintext, additional []byte) []byte {
	if nonce == nil {
		for i := 11; i >= 0; i-- {
			a.NQGg9x[i]++
			if a.NQGg9x[i] != 0 {
				break
			}
		}
		nonce = a.NQGg9x[:]
	}
	return a.AEAD.Seal(dst, nonce, plaintext, additional)
}

func (a *RaS2myUz) Open(dst, nonce, ciphertext, additional []byte) ([]byte, error) {
	if nonce == nil {
		for i := 11; i >= 0; i-- {
			a.NQGg9x[i]++
			if a.NQGg9x[i] != 0 {
				break
			}
		}
		nonce = a.NQGg9x[:]
	}
	return a.AEAD.Open(dst, nonce, ciphertext, additional)
}

// GdvNyLM recovers 0x1d011c0.
func GdvNyLM(material, iv []byte) cipher.Stream {
	var key [32]byte
	blake3.DeriveKey(key[:], "VLESS", material)
	block, _ := aes.NewCipher(key[:])
	return cipher.NewCTR(block, iv)
}

// QkDUgv recovers 0x1cf82e0. Header length bounds are checked by the original
// byte accesses; callers passing fewer than five bytes still panic.
func QkDUgv(header []byte) (int, error) {
	_ = header[3]
	_ = header[4]
	n := int(binary.BigEndian.Uint16(header[3:5]))
	if header[0] != 23 || header[1] != 3 || header[2] != 3 {
		n = 0
	}
	if n < 17 || n > 16640 {
		return n, fmt.Errorf("invalid header: %v", header[:5])
	}
	return n, nil
}
