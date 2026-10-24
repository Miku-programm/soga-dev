package kdf

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
)

// RQvD2v3oZ 来自原RTTI的命名string类型，显示前缀cqvWuv3A4YoV。
type RQvD2v3oZ string

// 原init@0x10a9500创建的四项映射，返回共享构造器指针。
var methodsAt5095DC0 = map[RQvD2v3oZ]*PAEsiXjhBp_C{
	"aes-128-gcm":            {16, 12, zMTZ28J},
	"aes-192-gcm":            {24, 12, zMTZ28J},
	"aes-256-gcm":            {32, 12, zMTZ28J},
	"chacha20-ietf-poly1305": {32, 12, ytXhWTPtl},
}

// NewAEAD 对应0x10a9740。
func (p *PAEsiXjhBp_C) NewAEAD(key []byte) (cipher.AEAD, error) {
	if len(key) != p.e8j2caBkTaf {
		return nil, fmt.Errorf("new aead key len wrong: %v/%v", len(key), p.e8j2caBkTaf)
	}
	return p.voQyhU6T(key)
}

// KeyLen 对应0x10a98c0。
func (p *PAEsiXjhBp_C) KeyLen() int { return p.e8j2caBkTaf }

// NonceLen 对应0x10a98e0。
func (p *PAEsiXjhBp_C) NonceLen() int { return p.kzkZ9jF }

// zMTZ28J 对应0x10a9900。
func zMTZ28J(key []byte) (cipher.AEAD, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(b)
}

// ytXhWTPtl 对应0x10a9960；原内联chacha20poly1305.New。
func ytXhWTPtl(key []byte) (cipher.AEAD, error) { return chacha20poly1305.New(key) }

// MakT9YbIpyNJ 对应0x10a9a40，不规范化名称。
func MakT9YbIpyNJ(method RQvD2v3oZ) (*PAEsiXjhBp_C, error) {
	if p, ok := methodsAt5095DC0[method]; ok {
		return p, nil
	}
	return nil, fmt.Errorf("not support method: %v", method)
}
