package kdf

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestOriginalAEADMethods(t *testing.T) {
	for _, tc := range []struct {
		method RQvD2v3oZ
		size   int
	}{{"aes-128-gcm", 16}, {"aes-192-gcm", 24}, {"aes-256-gcm", 32}, {"chacha20-ietf-poly1305", 32}} {
		p, err := MakT9YbIpyNJ(tc.method)
		if err != nil {
			t.Fatal(err)
		}
		again, _ := MakT9YbIpyNJ(tc.method)
		if p != again || p.KeyLen() != tc.size || p.NonceLen() != 12 {
			t.Fatal(tc)
		}
		a, err := p.NewAEAD(make([]byte, tc.size))
		if err != nil {
			t.Fatal(err)
		}
		nonce := make([]byte, 12)
		plain := []byte("native tunnel record")
		encrypted := a.Seal(nil, nonce, plain, []byte("aad"))
		decrypted, err := a.Open(nil, nonce, encrypted, []byte("aad"))
		if err != nil || !bytes.Equal(decrypted, plain) {
			t.Fatal(err)
		}
		encrypted[0] ^= 1
		if _, err := a.Open(nil, nonce, encrypted, []byte("aad")); err == nil {
			t.Fatal("tamper accepted")
		}
		if _, err := p.NewAEAD(nil); err == nil || !strings.HasPrefix(err.Error(), "new aead key len wrong: 0/") {
			t.Fatal(err)
		}
	}
	if _, err := MakT9YbIpyNJ("AES-128-GCM"); err == nil || err.Error() != "not support method: AES-128-GCM" {
		t.Fatal(err)
	}
}
func TestDerivationKnownVectorsAndBounds(t *testing.T) {
	// RFC5869测试向量1，SHA256、42字节输出。
	h := func(s string) []byte {
		b, e := hex.DecodeString(s)
		if e != nil {
			t.Fatal(e)
		}
		return b
	}
	got := ENKzXYE1a(sha256.New, bytes.Repeat([]byte{0x0b}, 22), h("000102030405060708090a0b0c"), h("f0f1f2f3f4f5f6f7f8f9"), 42)
	want := h("3cb25f25faacd57a90434f64d0362f2a2d2d0a90cf1a5a4c5db02d56ecc4c5bf34007208d5b887185865")
	if !bytes.Equal(got, want) {
		t.Fatalf("%x", got)
	}
	first := sha512.Sum512([]byte("password"))
	second := sha512.Sum512(append(first[:], []byte("password")...))
	want = append(first[:], second[:]...)
	if got := BMQB8vpnQXXB(sha512.New, []byte("password"), 80); !bytes.Equal(got, want[:80]) {
		t.Fatalf("%x", got)
	}
	if BMQB8vpnQXXB(sha512.New, nil, 0) != nil {
		t.Fatal("zero-size must retain nil")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("negative size must panic")
		}
	}()
	BMQB8vpnQXXB(sha512.New, nil, -1)
}
func TestHKDFFailureRetainsFatalExit(t *testing.T) {
	if os.Getenv("SOGA_TEST_HKDF_FATAL") == "1" {
		ENKzXYE1a(sha256.New, nil, nil, nil, 8161)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHKDFFailureRetainsFatalExit$")
	cmd.Env = append(os.Environ(), "SOGA_TEST_HKDF_FATAL=1")
	out, err := cmd.CombinedOutput()
	exit, ok := err.(*exec.ExitError)
	if !ok || exit.ExitCode() != 1 || !bytes.Contains(out, []byte("hkdf len: 8161, err: hkdf: entropy limit reached")) {
		t.Fatalf("%v %s", err, out)
	}
}
