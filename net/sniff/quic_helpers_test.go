package sniff

import (
	"bytes"
	"crypto"
	_ "crypto/sha256"
	"encoding/hex"
	"io"
	"testing"
)

func TestQUICVarintWidthsAndTruncation(t *testing.T) {
	for _, tc := range []struct {
		hex  string
		want uint64
	}{{"00", 0}, {"3f", 63}, {"4001", 1}, {"7fff", 16383}, {"80000001", 1}, {"bfffffff", 1073741823}, {"c000000000000001", 1}, {"ffffffffffffffff", 4611686018427387903}} {
		data, _ := hex.DecodeString(tc.hex)
		r := bytes.NewReader(append(data, 99))
		got, e := vFUaVI1hy(r)
		if e != nil || got != tc.want || r.Len() != 1 {
			t.Fatal(tc, got, e, r.Len())
		}
		for n := 0; n < len(data); n++ {
			got, e := vFUaVI1hy(bytes.NewReader(data[:n]))
			if got != 0 || e != io.EOF {
				t.Fatal(n, got, e)
			}
		}
	}
}
func TestQUICHKDFLabelAndOriginalPanic(t *testing.T) {
	secret := make([]byte, 32)
	for i := range secret {
		secret[i] = byte(i)
	}
	// 独立HMAC-SHA256两块展开计算得到，不通过待测函数产生。
	want, _ := hex.DecodeString("e5e7ce4cf665b7d530e88b738f8cba5c432373ad9b3e8ee54f7f9e460c7c86978217a90182bcab0ac14c")
	if got := qhXufgHMZAdy(crypto.SHA256, secret, []byte("ctx"), "quic hp", 42); !bytes.Equal(got, want) {
		t.Fatalf("%x", got)
	}
	if got := qhXufgHMZAdy(crypto.SHA256, secret, nil, "", 0); got == nil || len(got) != 0 {
		t.Fatal(got)
	}
	defer func() {
		if v := recover(); v != "quic: HKDF-Expand-Label invocation failed unexpectedly" {
			t.Fatal(v)
		}
	}()
	qhXufgHMZAdy(crypto.SHA256, secret, nil, "x", 8161)
}
