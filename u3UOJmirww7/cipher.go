package u3UOJmirww7

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"errors"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/hkdf"
	"io"
	"strconv"
)

// HugVbq is the original AEAD cipher interface (RTTI ildQ8zN9vi7.HugVbq).
type HugVbq interface {
	KeySize() int
	SaltSize() int
	Encrypter([]byte) (cipher.AEAD, error)
	Decrypter([]byte) (cipher.AEAD, error)
}

type ZPe9opSsSr9k int

// Error: 0x1943b40.
func (e ZPe9opSsSr9k) Error() string {
	return "key size error: need " + strconv.Itoa(int(e)) + " bytes"
}

type o4NlPb3xf8 struct {
	gyXKbc  []byte
	jQJmxUl func([]byte) (cipher.AEAD, error)
}

// Original init 0x19439c0. Decoded literals retain separate error identities.
var errorAt5098530 = errors.New("short packet")
var errorAt5098540 = errors.New("zero chunk")

// fIP4S_hGK0: 0x1943be0; hash funcval 0x376a388 -> crypto/sha1.New.
func fIP4S_hGK0(key, salt, info, out []byte) {
	if _, err := io.ReadFull(hkdf.New(sha1.New, key, salt, info), out); err != nil {
		panic(err)
	}
}

// KeySize: 0x1943ce0; SaltSize: 0x1943d00.
func (c *o4NlPb3xf8) KeySize() int  { return len(c.gyXKbc) }
func (c *o4NlPb3xf8) SaltSize() int { return max(16, len(c.gyXKbc)) }

// Encrypter: 0x1943d20; Decrypter: 0x1943e80. Neither validates salt length.
func (c *o4NlPb3xf8) Encrypter(salt []byte) (cipher.AEAD, error) {
	key := make([]byte, len(c.gyXKbc))
	fIP4S_hGK0(c.gyXKbc, salt, []byte("ss-subkey"), key)
	return c.jQJmxUl(key)
}
func (c *o4NlPb3xf8) Decrypter(salt []byte) (cipher.AEAD, error) {
	key := make([]byte, len(c.gyXKbc))
	fIP4S_hGK0(c.gyXKbc, salt, []byte("ss-subkey"), key)
	return c.jQJmxUl(key)
}

// a4z7CGof: 0x1943fc0.
func a4z7CGof(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

// Rk8ppUT2aLh: 0x1944020. The input key slice is retained without cloning.
func Rk8ppUT2aLh(key []byte) (HugVbq, error) {
	switch len(key) {
	case 16, 24, 32:
		return &o4NlPb3xf8{key, a4z7CGof}, nil
	}
	return nil, aes.KeySizeError(len(key))
}

// IEFKpuUGsJ: 0x1944120; BeqNJFSF: 0x1944200. Their error stores the
// required size (32), rather than the supplied key length.
func IEFKpuUGsJ(key []byte) (HugVbq, error) {
	if len(key) != 32 {
		return nil, ZPe9opSsSr9k(32)
	}
	return &o4NlPb3xf8{key, chacha20poly1305.New}, nil
}
func BeqNJFSF(key []byte) (HugVbq, error) {
	if len(key) != 32 {
		return nil, ZPe9opSsSr9k(32)
	}
	return &o4NlPb3xf8{key, chacha20poly1305.NewX}, nil
}
