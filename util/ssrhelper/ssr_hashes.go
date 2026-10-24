package ssrhelper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
)

// GhU1IKb4 restores 0x181b920; funcval0x376a380 is crypto/md5.New.
func GhU1IKb4(key []byte, parts ...[]byte) []byte {
	h := hmac.New(md5.New, key)
	for _, p := range parts {
		h.Write(p)
	}
	return h.Sum(nil)[:10]
}
func KWa8h0(p []byte) []byte      { sum := md5.Sum(p); return append([]byte(nil), sum[:]...) } //0x181bb20
func EyTkbUUkjjn(p []byte) []byte { h := sha1.New(); h.Write(p); return h.Sum(nil) }           //0x181bc60
