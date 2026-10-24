package RoaZP1oJy

import (
	"crypto/hmac"
	"crypto/md5"
	S "recovered.local/soga/ad5KEbw_EzWk"
)

// hmcSfmiGu restores 0x1ccc1a0, returning the complete 16-byte HMAC-MD5
// (unlike the separate Jafv truncated helper used by auth_aes).
func (*h2QD2cI1) hmcSfmiGu(key []byte, parts ...[]byte) []byte {
	h := hmac.New(md5.New, key)
	for _, p := range parts {
		h.Write(p)
	}
	return h.Sum(nil)
}

// z0O7jemcj restores 0x1ccc2a0, including returning before modifying the
// supplied PRNG for payloads over1440 and treating negative lengths as <=400.
func (*h2QD2cI1) z0O7jemcj(length int, seed []byte, r *S.HelGdo2ETK) int {
	if length > 1440 {
		return 0
	}
	r.InitFromBytesLen(seed, length)
	n := r.Next()
	switch {
	case length > 1300:
		return int(n % 31)
	case length > 900:
		return int(n % 127)
	case length > 400:
		return int(n % 521)
	default:
		return int(n % 1021)
	}
}

// vuOMfc restores 0x1ccc540. InitFromBytes is inlined in the original and
// pads the caller's backing array when capacity permits.
func (*h2QD2cI1) vuOMfc(seed []byte, r *S.HelGdo2ETK) int {
	r.InitFromBytes(seed)
	return int(r.Next() % 127)
}
