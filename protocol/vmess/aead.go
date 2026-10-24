package vmess

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/binary"
	"golang.org/x/crypto/chacha20poly1305"
)

type BwCT6jfasz6 interface {
	Encrypt([]byte, []byte) (int, error)
	GetEncryptLen([]byte) int
}
type G1Wopa interface {
	Decrypt([]byte, []byte) (int, error)
	GetDecryptLen([]byte) int
}
type M3rebiEbLfQx struct {
	yRSunnMW   cipher.AEAD
	pufJF2     uint16
	fKs57NhW6o []byte
}

func m1kq85rPFq(key []byte) (cipher.AEAD, error) { // 0x1d13080
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}
func p3FOeQkjR(key []byte) (cipher.AEAD, error) { // 0x1d13100
	first := md5.Sum(key)
	second := md5.Sum(first[:])
	var expanded [32]byte
	copy(expanded[:16], first[:])
	copy(expanded[16:], second[:])
	return chacha20poly1305.New(expanded[:])
}
func eG_qzDpw(key, iv []byte) (*M3rebiEbLfQx, error) { // 0x1d13660
	a, err := m1kq85rPFq(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, a.NonceSize())
	copy(nonce, iv)
	return &M3rebiEbLfQx{yRSunnMW: a, fKs57NhW6o: nonce}, nil
}
func vLtpACzP6C(key, iv []byte) (*M3rebiEbLfQx, error) { // 0x1d137a0
	a, err := p3FOeQkjR(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, a.NonceSize())
	copy(nonce, iv)
	return &M3rebiEbLfQx{yRSunnMW: a, fKs57NhW6o: nonce}, nil
}
func hYUXBKqsWpGa(key, iv []byte) (BwCT6jfasz6, error) { v, e := eG_qzDpw(key, iv); return v, e }   // 0x1d13460
func bYH58jCAX(key, iv []byte) (G1Wopa, error)         { v, e := eG_qzDpw(key, iv); return v, e }   // 0x1d134e0
func q4ORlubuA2c(key, iv []byte) (BwCT6jfasz6, error)  { v, e := vLtpACzP6C(key, iv); return v, e } // 0x1d13560
func a1jSJ0g(key, iv []byte) (G1Wopa, error)           { v, e := vLtpACzP6C(key, iv); return v, e } // 0x1d135e0
func (a *M3rebiEbLfQx) GetEncryptLen(src []byte) int   { return len(src) + a.yRSunnMW.Overhead() }  // 0x1d138e0
func (a *M3rebiEbLfQx) GetDecryptLen(src []byte) int   { return len(src) - a.yRSunnMW.Overhead() }  // 0x1d13960
func (a *M3rebiEbLfQx) Encrypt(dst, src []byte) (int, error) { // 0x1d139e0
	binary.BigEndian.PutUint16(a.fKs57NhW6o, a.pufJF2)
	a.pufJF2++
	a.yRSunnMW.Seal(dst[:0], a.fKs57NhW6o, src, nil)
	return len(src) + a.yRSunnMW.Overhead(), nil
}
func (a *M3rebiEbLfQx) Decrypt(dst, src []byte) (int, error) { // 0x1d13b00
	binary.BigEndian.PutUint16(a.fKs57NhW6o, a.pufJF2)
	a.pufJF2++
	if _, err := a.yRSunnMW.Open(dst[:0], a.fKs57NhW6o, src, nil); err != nil {
		return 0, err
	}
	return len(src) - a.yRSunnMW.Overhead(), nil
}

type P8gDoqFWbS struct{}

func (a *P8gDoqFWbS) GetEncryptLen(src []byte) int { return len(src) } // 0x1d1c140
func (a *P8gDoqFWbS) GetDecryptLen(src []byte) int { return len(src) } // 0x1d1c160
func (a *P8gDoqFWbS) Encrypt(dst, src []byte) (int, error) { // 0x1d1c180
	if len(src) == 0 {
		return 0, nil
	}
	_ = dst[0]
	return copy(dst, src), nil
}
func (a *P8gDoqFWbS) Decrypt(dst, src []byte) (int, error) { // 0x1d1c260
	if len(src) == 0 {
		return 0, nil
	}
	_ = dst[0]
	return copy(dst, src), nil
}
