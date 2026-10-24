package cNT24rG4F

import "testing"

func TestRecoveredVMessAuthenticationMode(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  F6DuWgn9N
	}{
		{"md5", "md5"}, {" MD5\n", "md5"}, {"\u3000AeAd\u00a0", "aead"},
		{"auto", "auto"}, {"", "auto"}, {"md5finptr", "auto"}, {"aeadx", "auto"}, {"unknown", "auto"},
	} {
		if got := zaWmXx3rGK35(tc.input); got != tc.want {
			t.Errorf("%q: got %q, want %q", tc.input, got, tc.want)
		}
	}
}
