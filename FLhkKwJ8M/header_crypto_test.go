package FLhkKwJ8M

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"github.com/gofrs/uuid/v5"
	"hash"
	"hash/crc32"
	"strings"
	"testing"
	"time"
	"unsafe"
)

// Separate functional hash factory for the peer/fixtures; no production KDF call.
func peerKDF(key []byte, labels ...string) []byte {
	var build func(int) hash.Hash
	build = func(n int) hash.Hash {
		if n == 0 {
			return hmac.New(sha256.New, []byte("VMess AEAD KDF"))
		}
		return hmac.New(func() hash.Hash { return build(n - 1) }, []byte(labels[n-1]))
	}
	h := build(len(labels))
	h.Write(key)
	return h.Sum(nil)
}
func peerGCM(key []byte) cipher.AEAD {
	b, e := aes.NewCipher(key)
	if e != nil {
		panic(e)
	}
	a, e := cipher.NewGCM(b)
	if e != nil {
		panic(e)
	}
	return a
}
func TestKDFIndependentHashlibVectors(t *testing.T) {
	for _, v := range []struct {
		key  []byte
		path []string
		want string
	}{{nil, nil, "515f37c9eaa2812ab48bfb4d3d6799a480e3d0a5872db4775328cf3a1b839b10"}, {[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, []string{"AES Auth ID Encryption"}, "9fa4289c41650861a45b34aeab3879fe4785dce57ab3f68cfb0cc60fca69460a"}, {[]byte("input"), []string{"one", "two"}, "5f9a44819493bff0acf34e4520ed9c5e3794f507aa1bad49a73509dcc0ef3f90"}, {[]byte("input"), []string{strings.Repeat("A", 80), strings.Repeat("B", 100)}, "8aa0c53ccbebd6c98f9a6ed82640a1f5143e6cdfea630381cc77d1ae77b84b29"}} {
		got := xXVvCU5cjuB(v.key, v.path...)
		if hex.EncodeToString(got) != v.want {
			t.Fatalf("%x", got)
		}
	}
}
func TestAEADRequestEnvelope(t *testing.T) {
	key := [16]byte{1, 2, 3, 4, 5, 6, 7, 8}
	payload := []byte("independent encrypted request")
	stamp := time.Unix(1720000000, 987)
	wire := s9ADKgL(key, payload, stamp)
	if len(wire) != len(payload)+58 {
		t.Fatal(len(wire))
	}
	var auth [16]byte
	b, _ := aes.NewCipher(peerKDF(key[:], "AES Auth ID Encryption")[:16])
	b.Decrypt(auth[:], wire[:16])
	if binary.BigEndian.Uint64(auth[:8]) != uint64(stamp.Unix()) || binary.BigEndian.Uint32(auth[12:]) != crc32.ChecksumIEEE(auth[:12]) {
		t.Fatal("authID")
	}
	labels := []string{string(wire[:16]), string(wire[34:42])}
	lk := peerKDF(key[:], append([]string{"VMess Header AEAD Key_Length"}, labels...)...)
	ln := peerKDF(key[:], append([]string{"VMess Header AEAD Nonce_Length"}, labels...)...)
	length, err := peerGCM(lk[:16]).Open(nil, ln[:12], wire[16:34], wire[:16])
	if err != nil || binary.BigEndian.Uint16(length) != uint16(len(payload)) {
		t.Fatal(err)
	}
	pk := peerKDF(key[:], append([]string{"VMess Header AEAD Key"}, labels...)...)
	pn := peerKDF(key[:], append([]string{"VMess Header AEAD Nonce"}, labels...)...)
	plain, err := peerGCM(pk[:16]).Open(nil, pn[:12], wire[42:], wire[:16])
	if err != nil || !bytes.Equal(plain, payload) {
		t.Fatal(err)
	}
}
func TestClientLayoutAndIdentitySharing(t *testing.T) {
	if unsafe.Sizeof(ClientConfiguration{}) != 80 || unsafe.Offsetof(ClientConfiguration{}.AEAD) != 72 || unsafe.Sizeof(GC3qACg18wgq{}) != 40 || unsafe.Sizeof(HRzHVG1yP5O{}) != 168 || unsafe.Sizeof(IxhMZ66CqoH1{}) != 40 {
		t.Fatal("RTTI/call layout")
	}
	for _, kind := range []string{"auto", "aes-128-gcm", "chacha20-poly1305", "none", "zero"} {
		c, err := Iw0IEMM3ha(ClientConfiguration{UUID: "00010203-0405-0607-0809-0a0b0c0d0e0f", AlterID: 2, Security: kind})
		if err != nil {
			t.Fatal(err)
		}
		if len(c.roEClw) != 3 || c.roEClw[2].AJxI2iBn != c.gyJC3dk_8 {
			t.Fatal("primary order")
		}
		if &c.roEClw[0].P3zlkY2W1n[0] != &c.roEClw[2].P3zlkY2W1n[0] {
			t.Fatal("key not shared")
		}
		id := uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		if *c.gyJC3dk_8 != id {
			t.Fatal("uuid")
		}
	}
	if _, err := Iw0IEMM3ha(ClientConfiguration{UUID: "not uuid", Security: "bad"}); err == nil || strings.Contains(err.Error(), "unknown security") {
		t.Fatal("parse ordering", err)
	}
	if _, err := Iw0IEMM3ha(ClientConfiguration{UUID: "00000000-0000-0000-0000-000000000000", Security: "bad"}); err == nil || err.Error() != "unknown security type: bad" {
		t.Fatal(err)
	}
}
