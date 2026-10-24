package g0MlUOhvO

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/sha3"
	"io"
	D "recovered.local/soga/DDYEn9Z_b"
	"testing"
	"unsafe"
)

func recordFixture(t *testing.T) (*Cb7vj3KQA8, *wireFixture) {
	t.Helper()
	block, e := aes.NewCipher(make([]byte, 16))
	if e != nil {
		t.Fatal(e)
	}
	a, e := cipher.NewGCM(block)
	if e != nil {
		t.Fatal(e)
	}
	recv, send := sha3.NewShake128(), sha3.NewShake128()
	recv.Write([]byte("record-test"))
	send.Write([]byte("record-test"))
	raw := &wireFixture{Reader: bytes.NewReader(nil)}
	return &Cb7vj3KQA8{DrpXao: raw, e8aahi_R3S: a, bUWvbqWzlH: a, rOfYjOzCdA1L: make([]byte, 12), cdwwdRLQG: make([]byte, 12), w1vg8Dd: &ErrfcUUB{pzOWW_: recv}, r6A494: &ErrfcUUB{pzOWW_: send}}, raw
}
func TestTunnelCryptoRTTIAndDerivation(t *testing.T) {
	var c Cb7vj3KQA8
	if unsafe.Sizeof(c) != 136 || unsafe.Offsetof(c.e8aahi_R3S) != 24 || unsafe.Offsetof(c.rOfYjOzCdA1L) != 40 || unsafe.Offsetof(c.bUWvbqWzlH) != 64 || unsafe.Offsetof(c.cdwwdRLQG) != 80 || unsafe.Offsetof(c.w1vg8Dd) != 104 || unsafe.Offsetof(c.r6A494) != 112 || unsafe.Offsetof(c.npFKkBd4j) != 120 || unsafe.Sizeof(ErrfcUUB{}) != 24 || unsafe.Sizeof(T4k2R7cJ{}) != 24 {
		t.Fatal("RTTI mismatch")
	}
	method, e := D.MakT9YbIpyNJ("aes-128-gcm")
	if e != nil {
		t.Fatal(e)
	}
	secret, salt := make([]byte, 32), make([]byte, 12)
	for i := range secret {
		secret[i] = byte(i)
	}
	for i := range salt {
		salt[i] = byte(i)
	}
	cfg := &VTl94pzz2IkB{zkn35IN: method, oWvGY4I9cA: secret}
	key, a, e := cfg.NewAEAD(salt)
	if e != nil || hex.EncodeToString(key) != "7441837dbd1a34e69c0804146d4993d3" || a.NonceSize() != 12 {
		t.Fatalf("%x %v", key, e)
	}
	if !cfg.Verify(make([]byte, 16)) || cfg.Verify(make([]byte, 15)) {
		t.Fatal("verification size")
	}
}
func TestEncryptedRecordsBufferingPaddingAndNonce(t *testing.T) {
	for _, size := range []int{0, 1, 123, 60000} {
		c, raw := recordFixture(t)
		data := bytes.Repeat([]byte{'a'}, size)
		if n, e := c.oH5wsFE(nil, data); n != size || e != nil {
			t.Fatal(n, e)
		}
		if c.cdwwdRLQG[0] != 2 {
			t.Fatal("two AEAD records per frame")
		}
		raw.Reader = bytes.NewReader(raw.out.Bytes())
		var out bytes.Buffer
		p := make([]byte, 79)
		for out.Len() < size || size == 0 && out.Len() == 0 {
			n, e := c.e21qHx(p)
			if e != nil {
				t.Fatal(e)
			}
			out.Write(p[:n])
			if size == 0 {
				break
			}
		}
		if !bytes.Equal(out.Bytes(), data) || c.rOfYjOzCdA1L[0] != 2 || c.npFKkBd4j != nil {
			t.Fatal(size, out.Len())
		}
	}
}
func TestRecordAuthenticationErrorAsymmetry(t *testing.T) {
	for _, header := range []bool{true, false} {
		c, raw := recordFixture(t)
		c.oH5wsFE(nil, []byte("payload"))
		wire := append([]byte(nil), raw.out.Bytes()...)
		if header {
			wire[0] ^= 1
		} else {
			wire[len(wire)-1] ^= 1
		}
		raw.Reader = bytes.NewReader(wire)
		n, e := c.e21qHx(make([]byte, 100))
		if n != 0 {
			t.Fatal(n)
		}
		if header && e == nil {
			t.Fatal("header authentication must return error")
		}
		if !header && e != nil {
			t.Fatal("original payload auth returns zero,nil", e)
		}
		want := byte(1)
		if !header {
			want = 2
		}
		if c.rOfYjOzCdA1L[0] != want {
			t.Fatal("nonce failure progression")
		}
	}
}
func TestRecordShortWriteErrorAndInputTail(t *testing.T) {
	c, raw := recordFixture(t)
	raw.shortWrite = true
	if n, e := c.oH5wsFE(nil, []byte("abc")); n != 3 || e != nil {
		t.Fatal(n, e)
	}
	raw.writeErr = io.ErrClosedPipe
	if n, e := c.oH5wsFE(nil, []byte("abc")); n != 0 || !errors.Is(e, io.ErrClosedPipe) {
		t.Fatal(n, e)
	}
	c, raw = recordFixture(t)
	backing := make([]byte, 300)
	copy(backing, []byte("abc"))
	if _, e := c.oH5wsFE(nil, backing[:3]); e != nil {
		t.Fatal(e)
	}
	if bytes.Equal(backing[3:], make([]byte, 297)) {
		t.Fatal("original writes random padding into spare input capacity")
	}
	c.Free()
	if !raw.freed {
		t.Fatal("nested Free")
	}
}
func TestRecordUint16LengthWrap(t *testing.T) {
	c, raw := recordFixture(t)
	data := make([]byte, 65535)
	if n, e := c.oH5wsFE(nil, data); n != 65535 || e != nil {
		t.Fatal(n, e)
	}
	raw.Reader = bytes.NewReader(raw.out.Bytes())
	length, padding, e := c.bpzXJ_7IMXKd()
	if e != nil || padding == 0 || length != padding-1 {
		t.Fatal(length, padding, e)
	}
}
