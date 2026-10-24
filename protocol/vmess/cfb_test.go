package vmess

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"testing"
)

type observedStream struct {
	src    []byte
	dstLen int
}

func (s *observedStream) XORKeyStream(dst, src []byte) {
	s.src = append([]byte(nil), src...)
	s.dstLen = len(dst)
}
func TestCFBOriginalMachineTraceFraming(t *testing.T) {
	for _, tc := range []struct {
		n, cap int
		want   string
	}{{9, 13, "aaaaaaaa68656c6c6f4f9f2cab"}, {9, 9, "aaaaaaaa68656c6c6f"}, {16, 20, "aaaaaaaa68656c6c6faaaaaaaaaaaaaa4f9f2cab"}} {
		buf := bytes.Repeat([]byte{0xaa}, tc.cap)
		stream := new(observedStream)
		c := &NzmxEE{stream}
		n, err := c.Encrypt(buf[:tc.n], []byte("hello"))
		if err != nil || n != 9 || hex.EncodeToString(buf) != tc.want || hex.EncodeToString(stream.src) != "aaaaaaaa68656c6c6f" || stream.dstLen != tc.n {
			t.Fatal(n, err, hex.EncodeToString(buf), stream)
		}
	}
}
func TestCFBDecryptActualClientCipher(t *testing.T) {
	key, iv := []byte("0123456789abcdef"), []byte("abcdefghijklmnop")
	block, _ := aes.NewCipher(key)
	raw, _ := hex.DecodeString("4f9f2cab68656c6c6f")
	encrypted := make([]byte, len(raw))
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(encrypted, raw)
	c, err := txowtYpKbXXG(key, iv)
	if err != nil {
		t.Fatal(err)
	}
	dst := make([]byte, len(encrypted))
	n, err := c.Decrypt(dst, encrypted)
	if n != 5 || err != nil || string(dst[:n]) != "hello" {
		t.Fatal(n, err, dst)
	}
	c, _ = txowtYpKbXXG(key, iv)
	encrypted[0] ^= 1
	n, err = c.Decrypt(dst, encrypted)
	if n != 0 || err == nil || err.Error() != "data fnv32 wrong" {
		t.Fatal(n, err)
	}
	if _, err = txowtYpKbXXG([]byte{1}, iv); err == nil {
		t.Fatal("invalid key accepted")
	}
}
func TestCFBEncryptUsesDecrypterAndLeavesPrefix(t *testing.T) {
	key, iv := []byte("0123456789abcdef"), []byte("abcdefghijklmnop")
	c, _ := txowtYpKbXXG(key, iv)
	block, _ := aes.NewCipher(key)
	reference := cipher.NewCFBDecrypter(block, iv)
	for i := 0; i < 2; i++ {
		payload := bytes.Repeat([]byte{byte(i + 1)}, 64)
		dst := bytes.Repeat([]byte{0xaa}, 68)
		expected := append(bytes.Repeat([]byte{0xaa}, 4), payload...)
		reference.XORKeyStream(expected, expected)
		n, err := c.Encrypt(dst, payload)
		if n != 68 || err != nil || !bytes.Equal(dst, expected) {
			t.Fatal(n, err)
		}
	}
}
