package ssmethod

import "testing"

func TestKeyLenOriginalAlgorithms(t *testing.T) {
	for cipher, want := range map[C6q3Upu]int{"aes-128-gcm": 16, "aes-192-gcm": 24, "aes-256-gcm": 32, "chacha20-ietf-poly1305": 32, "2022-blake3-aes-128-gcm": 16, "2022-blake3-aes-256-gcm": 32} {
		if got, err := cipher.KeyLen(); err != nil || got != want {
			t.Fatalf("%s: %d %v", cipher, got, err)
		}
	}
	for _, cipher := range []C6q3Upu{"", "AES-128-GCM", "none", "2022-blake3-chacha20-poly1305"} {
		n, err := cipher.KeyLen()
		if n != 0 || err == nil || err.Error() != "unknown ss cipher <"+string(cipher)+"> key len" {
			t.Fatalf("%s: %d %v", cipher, n, err)
		}
	}
}
