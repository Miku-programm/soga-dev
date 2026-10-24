package w2YaTd9_

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
)

var subkeyInfoAt50998F0 = []byte("ss-subkey") // init @ 0x10a11e0.
func (c C6q3Upu) Is2022() bool {
	return c == "2022-blake3-aes-128-gcm" || c == "2022-blake3-aes-256-gcm"
}                                       // 0x10a1280
func (c C6q3Upu) SaltLen() (int, error) { return c.KeyLen() } // 0x10a1520
func (c C6q3Upu) NonceLen() int         { return 12 }         // 0x10a1560
func (c C6q3Upu) Overhead() int         { return 16 }         // 0x10a1580
func (c C6q3Upu) PSK(password string) ([]byte, error) { // 0x10a15a0
	n, err := c.KeyLen()
	if err != nil {
		return nil, err
	}
	if !c.Is2022() {
		return giRMs6N(password, n), nil
	}
	if password == "" {
		return nil, errors.New("empty password")
	}
	key, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return nil, fmt.Errorf("decode ss2022 password <%v> err: %v", password, err)
	}
	if len(key) < n {
		return nil, fmt.Errorf("ss2022 password key length too short: %v < %v", len(key), n)
	}
	if len(key) > n {
		digest := sha256.Sum256(key)
		key = digest[:n]
	}
	return key, nil
}
func giRMs6N(password string, n int) []byte { // 0x10a7d40
	h := md5.New()
	var result, previous []byte
	for len(result) < n {
		h.Write(previous)
		h.Write([]byte(password))
		result = h.Sum(result)
		previous = result[len(result)-16:]
		h.Reset()
	}
	return result[:n]
}
func (c C6q3Upu) SubKey(key, salt []byte) []byte { // 0x10a1ac0
	result := make([]byte, len(key))
	c.SubKeyTo(SaYhx5xog9(), result, key, salt)
	return result
}
func (c C6q3Upu) SubKeyTo(s *MXXLkaPB, dst, key, salt []byte) { // 0x10a1c00
	if c.Is2022() {
		s.oE369eL1UY(dst, key, salt)
		return
	}
	s.wJRMV1dV(salt)
	s.b0pUkTAq(key)
	s.dHQaCID_x(dst, subkeyInfoAt50998F0)
}
func (c C6q3Upu) NewBlock(key []byte) (cipher.Block, error) { return aes.NewCipher(key) } // 0x10a1de0
func (c C6q3Upu) NewAEAD(key []byte) (cipher.AEAD, error) { // 0x10a1e60
	switch c {
	case "aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm":
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		return cipher.NewGCM(block)
	case "chacha20-ietf-poly1305":
		return chacha20poly1305.New(key)
	default:
		return nil, fmt.Errorf("unknown ss cipher: %v", c)
	}
}
