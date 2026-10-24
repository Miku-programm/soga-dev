package JafvCs5

import (
	"bytes"
	"encoding/hex"
	"testing"
)

// Long-path fixture retains the original shared-state HMAC behavior, independently
// calculated with Python hashlib in work/ida-20261023/credential_hash_vectors.py.
func TestKDFIndependentNestedHMACVectors(t *testing.T) {
	key := make([]byte, 16)
	for i := range key {
		key[i] = byte(i)
	}
	cases := []struct {
		key  []byte
		path [][]byte
		want string
	}{
		{nil, nil, "515f37c9eaa2812ab48bfb4d3d6799a480e3d0a5872db4775328cf3a1b839b10"},
		{key, [][]byte{[]byte("AES Auth ID Encryption")}, "9fa4289c41650861a45b34aeab3879fe4785dce57ab3f68cfb0cc60fca69460a"},
		{[]byte("input"), [][]byte{[]byte("one"), []byte("two")}, "5f9a44819493bff0acf34e4520ed9c5e3794f507aa1bad49a73509dcc0ef3f90"},
		{[]byte("input"), [][]byte{bytes.Repeat([]byte("A"), 80), bytes.Repeat([]byte("B"), 100)}, "8e424bc683ff619fa7740315ca7f097433b55baa3b1948c4efc1443a2d78bcca"},
	}
	for _, c := range cases {
		if got := hex.EncodeToString(CVqKigbE26(c.key, c.path...)); got != c.want {
			t.Errorf("KDF=%s want %s", got, c.want)
		}
	}
}
