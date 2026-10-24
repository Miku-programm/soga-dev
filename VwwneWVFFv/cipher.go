package VwwneWVFFv

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rc4"
	"errors"
	"golang.org/x/crypto/chacha20"
	"strconv"
)

type ZDHlm7vHo interface {
	IVSize() int
	Encrypter([]byte) cipher.Stream
	Decrypter([]byte) cipher.Stream
}
type XJzap3 int

// Error: 0x19496e0; prefix recovered by bounded decoder emulation.
func (e XJzap3) Error() string { return "key size error: need " + strconv.Itoa(int(e)) + " bytes" }

var errorAt5098550 = errors.New("short packet") // init 0x1949620

type sBODxR3OBCW struct{ cipher.Block }
type lBN76XO struct{ cipher.Block }
type hSaexzkFHX []byte
type z9Wa9EpYx []byte
type hNflyaqjzpAO []byte

// AES CTR: 0x1949780/19497c0/1949840/19498c0.
func (c *sBODxR3OBCW) IVSize() int                       { return c.BlockSize() }
func (c *sBODxR3OBCW) Encrypter(iv []byte) cipher.Stream { return cipher.NewCTR(c, iv) }
func (c *sBODxR3OBCW) Decrypter(iv []byte) cipher.Stream { return cipher.NewCTR(c, iv) }
func E6vQRRYI(key []byte) (ZDHlm7vHo, error) {
	b, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	return &sBODxR3OBCW{b}, nil
}

// AES CFB: 0x1949980/19499c0/1949a40/1949ac0.
func (c *lBN76XO) IVSize() int                       { return c.BlockSize() }
func (c *lBN76XO) Encrypter(iv []byte) cipher.Stream { return cipher.NewCFBEncrypter(c, iv) }
func (c *lBN76XO) Decrypter(iv []byte) cipher.Stream { return cipher.NewCFBDecrypter(c, iv) }
func HQZtzHEUA3(key []byte) (ZDHlm7vHo, error) {
	b, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	return &lBN76XO{b}, nil
}

// ChaCha20-IETF: 0x1949b80/1949ba0/1949c20/1949d00.
func (c hSaexzkFHX) IVSize() int { return 12 }
func (c hSaexzkFHX) Encrypter(iv []byte) cipher.Stream {
	s, e := chacha20.NewUnauthenticatedCipher(c, iv)
	if e != nil {
		panic(e)
	}
	return s
}
func (c hSaexzkFHX) Decrypter(iv []byte) cipher.Stream { return c.Encrypter(iv) }
func E9kDmJ55P(key []byte) (ZDHlm7vHo, error) {
	if len(key) != 32 {
		return nil, XJzap3(32)
	}
	return hSaexzkFHX(key), nil
}

// XChaCha20: 0x1949d80/1949da0/1949e20/1949f00.
func (c z9Wa9EpYx) IVSize() int { return 24 }
func (c z9Wa9EpYx) Encrypter(iv []byte) cipher.Stream {
	s, e := chacha20.NewUnauthenticatedCipher(c, iv)
	if e != nil {
		panic(e)
	}
	return s
}
func (c z9Wa9EpYx) Decrypter(iv []byte) cipher.Stream { return c.Encrypter(iv) }
func AfbgTLVHrju(key []byte) (ZDHlm7vHo, error) {
	if len(key) != 32 {
		return nil, XJzap3(32)
	}
	return z9Wa9EpYx(key), nil
}

// RC4-MD5: 0x1949f80/1949fa0/194a140/194a1c0. Constructor accepts all
// key lengths; the stream key is always the 16-byte MD5(key || IV).
func (c hNflyaqjzpAO) IVSize() int { return 16 }
func (c hNflyaqjzpAO) Encrypter(iv []byte) cipher.Stream {
	h := md5.New()
	h.Write(c)
	h.Write(iv)
	s, _ := rc4.NewCipher(h.Sum(nil))
	return s
}
func (c hNflyaqjzpAO) Decrypter(iv []byte) cipher.Stream { return c.Encrypter(iv) }
func F3Yl5rl(key []byte) (ZDHlm7vHo, error)              { return hNflyaqjzpAO(key), nil }
