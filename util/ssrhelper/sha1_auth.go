package ssrhelper

import (
	"crypto/hmac"
	"crypto/sha1"
)

// L01ck7znS restores 0x181ba20: HMAC-SHA1 over each variadic slice,
// returning the first ten bytes with the complete digest's capacity.
// ELF funcval 0x376a388 points to crypto/sha1.New@0x740020.
func L01ck7znS(key []byte, parts ...[]byte) []byte {
	h := hmac.New(sha1.New, key)
	for _, part := range parts {
		h.Write(part)
	}
	return h.Sum(nil)[:10]
}
