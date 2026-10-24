package vmess

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"hash/fnv"
)

// NzmxEE 对应原 RTTI 的 16 字节 cipher.Stream 包装。
type NzmxEE struct{ pUW5E6ld6 cipher.Stream }

func txowtYpKbXXG(key, iv []byte) (*NzmxEE, error) { // 0x1d14d80
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 原构造始终使用 Decrypter，包括 Encrypt 接口一侧。
	return &NzmxEE{pUW5E6ld6: cipher.NewCFBDecrypter(block, iv)}, nil
}
func (c *NzmxEE) GetEncryptLen(src []byte) int { return len(src) + 4 } // 0x1d14e80
func (c *NzmxEE) GetDecryptLen(src []byte) int { return len(src) - 4 } // 0x1d14ea0
func (c *NzmxEE) Encrypt(dst, src []byte) (int, error) { // 0x1d14ec0
	copy(dst[4:], src)
	payload := dst[4 : len(src)+4]
	h := fnv.New32a()
	h.Write(payload)
	// 原机器码将校验追加在 len(dst) 处并丢弃返回切片，而非写入前四字节。
	// 已用原指令的有界执行核实，不在源码恢复阶段修改此行为。
	h.Sum(dst)
	c.pUW5E6ld6.XORKeyStream(dst, dst[:len(src)+4])
	return len(src) + 4, nil
}
func (c *NzmxEE) Decrypt(dst, src []byte) (int, error) { // 0x1d15080
	c.pUW5E6ld6.XORKeyStream(dst, src)
	payload := dst[:len(src)][4:]
	h := fnv.New32a()
	h.Write(payload)
	if binary.BigEndian.Uint32(dst[:4]) != h.Sum32() {
		return 0, errors.New("data fnv32 wrong")
	}
	copy(dst, payload)
	return len(src) - 4, nil
}
