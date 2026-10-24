package H8_In_Vuy

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
)

// YPrsw38MN restores 0x19c4720. Function pointer 0x376a380 is md5.New.
func YPrsw38MN(key, data []byte) []byte {
	h := hmac.New(md5.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// B4_4_BnW restores 0x19c47e0. Function pointer 0x376a388 is sha1.New.
func B4_4_BnW(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// RP_qK7a restores 0x19c48a0; YhCPGpz restores 0x19c49e0.
func RP_qK7a(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}

func YhCPGpz(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
