package kdf

import (
	"golang.org/x/crypto/hkdf"
	"hash"
	"io"
	"log"
)

// BMQB8vpnQXXB 对应0x10aae60，连续散列previousDigest||password后截取。
func BMQB8vpnQXXB(newHash func() hash.Hash, password []byte, keyLen int) []byte {
	h := newHash()
	var result, last []byte
	for len(result) < keyLen {
		h.Write(last)
		h.Write(password)
		result = h.Sum(result)
		last = result[len(result)-h.Size():]
		h.Reset()
	}
	return result[:keyLen]
}

// ENKzXYE1a 对应0x10ab020，HKDF读取失败保留原log.Fatalf退出行为。
func ENKzXYE1a(newHash func() hash.Hash, secret, salt, info []byte, size int) []byte {
	reader := hkdf.New(newHash, secret, salt, info)
	output := make([]byte, size)
	if _, err := io.ReadFull(reader, output); err != nil {
		log.Fatalf("hkdf len: %v, err: %v", size, err)
	}
	return output
}
