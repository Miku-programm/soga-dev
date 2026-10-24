package mieru

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"testing"
	"time"
)

func TestKeyMaterialIndependentVectors(t *testing.T) {
	cases := []struct{ key, salt, want string }{
		{"", "", "6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d"},
		{"p\u0000ass", "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", "6317f82d0a3c5fb38280c82cee6f345da45849db9b39a923173dad82366336db"},
		{"中文密码", "ff0010", "090054495da47b2d9810f06f5bc3f36e5a7a3a2df6c449b870fce4fa4c917cb7"},
	}
	for _, c := range cases {
		salt, _ := hex.DecodeString(c.salt)
		sum := gUBL99(c.key, salt)
		if hex.EncodeToString(sum[:]) != c.want {
			t.Fatal("password key mismatch")
		}
	}
	for _, size := range []int{20, 21, 22, 23, 24, 25, 29} {
		nonce := make([]byte, size)
		for i := range nonce {
			nonce[i] = byte(i)
		}
		before := append([]byte(nil), nonce...)
		sum := sha256.Sum256(append([]byte("password"), before[:16]...))
		uJghIeaJPse(nonce, "password")
		for i := range nonce {
			want := before[i]
			if i >= 20 && i < 24 {
				want = sum[i-20]
			}
			if nonce[i] != want {
				t.Fatalf("size %d byte %d", size, i)
			}
		}
		if pyUlp642nl("password", nonce) != (size == 24) {
			t.Fatalf("unexpected nonce acceptance size %d", size)
		}
	}
	nonce, err := o0fNBhk("password")
	if err != nil || len(nonce) != 24 || !pyUlp642nl("password", nonce) || pyUlp642nl("wrong", nonce) {
		t.Fatalf("generated nonce %x %v", nonce, err)
	}
	for _, seconds := range []int64{0, 59, 60, 119, 120, 1700000000} {
		keys := pVQO7MqTE4ao(time.Unix(seconds, 0))
		rounded := ((seconds + 60) / 120) * 120
		for i, k := range keys {
			var raw [8]byte
			binary.BigEndian.PutUint64(raw[:], uint64(rounded+int64(i-1)*120))
			want := sha256.Sum256(raw[:])
			if hex.EncodeToString(k) != hex.EncodeToString(want[:]) {
				t.Fatalf("time key %d/%d", seconds, i)
			}
		}
	}
}
func TestPaddingDeterministicFingerprint(t *testing.T) {
	for _, key := range []string{"", "abc", "中文", "3.32.0-soga"} {
		sum := sha256.Sum256([]byte(key + " 3.32.0-soga"))
		want := int(binary.BigEndian.Uint32(sum[:4]) & 0x7fffffff)
		for _, n := range []int{-1, 0, 1, 2, 17, 999} {
			got := jyDI5g(n, key)
			w := 0
			if n > 0 {
				w = want % n
			}
			if got != w {
				t.Fatalf("fingerprint %q/%d: %d != %d", key, n, got, w)
			}
		}
	}
	b := []byte{0, 31, 32, 126, 127, 128, 160, 254, 255}
	goAOlbfD50OV(b, 0, len(b))
	for _, v := range b {
		if v < 32 || v > 126 {
			t.Fatalf("nonprintable %x", b)
		}
	}
	if b[2] != 32 || b[3] != 126 || b[6] != 32 || b[7] != 126 {
		t.Fatal("original printable/high-bit bytes changed")
	}
}
