package aead

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"testing"
	"unsafe"
)

func TestCipherConstructorsAndKeyAliasing(t *testing.T) {
	if unsafe.Sizeof(o4NlPb3xf8{}) != 32 {
		t.Fatal("RTTI")
	}
	for _, n := range []int{16, 24, 32} {
		key := make([]byte, n)
		c, err := Rk8ppUT2aLh(key)
		if err != nil || c.KeySize() != n || c.SaltSize() != n {
			t.Fatal(c, err)
		}
		before, _ := c.Encrypter(nil)
		nonce := make([]byte, before.NonceSize())
		sealed := before.Seal(nil, nonce, []byte("data"), nil)
		key[0] = 1
		after, _ := c.Decrypter(nil)
		if _, err = after.Open(nil, nonce, sealed, nil); err == nil {
			t.Fatal("input key was cloned")
		}
	}
	if _, err := Rk8ppUT2aLh(make([]byte, 7)); err != aes.KeySizeError(7) {
		t.Fatal(err)
	}
	for _, ctor := range []func([]byte) (HugVbq, error){IEFKpuUGsJ, BeqNJFSF} {
		if _, err := ctor(nil); err != ZPe9opSsSr9k(32) {
			t.Fatal(err)
		}
		c, _ := ctor(make([]byte, 32))
		encrypter, _ := c.Encrypter([]byte("short salt"))
		decrypter, _ := c.Decrypter([]byte("short salt"))
		nonce := make([]byte, encrypter.NonceSize())
		message := []byte("authenticated data")
		wire := encrypter.Seal(nil, nonce, message, nil)
		plain, err := decrypter.Open(nil, nonce, wire, nil)
		if err != nil || !bytes.Equal(plain, message) {
			t.Fatal(err)
		}
	}
	if ZPe9opSsSr9k(32).Error() != "key size error: need 32 bytes" {
		t.Fatal("error text")
	}
}
func TestHKDFAndAESIndependentVector(t *testing.T) {
	// SHA1 HKDF vector independently computed from extract/expand HMACs.
	key := []byte("0123456789abcdef")
	salt := []byte("fixed salt")
	derived := make([]byte, 16)
	fIP4S_hGK0(key, salt, []byte("ss-subkey"), derived)
	expected, _ := hex.DecodeString("625554d07361723a89d496fade4cf693")
	if !bytes.Equal(derived, expected) {
		t.Fatalf("%x", derived)
	}
	block, _ := aes.NewCipher(expected)
	reference, _ := cipher.NewGCM(block)
	c, _ := Rk8ppUT2aLh(key)
	actual, _ := c.Encrypter(salt)
	nonce := make([]byte, 12)
	if !bytes.Equal(reference.Seal(nil, nonce, key, nil), actual.Seal(nil, nonce, key, nil)) {
		t.Fatal("AEAD mismatch")
	}
}
