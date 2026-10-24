package m3Cw8P3_Le

import "crypto/md5"

// ArNrGS_EV: 0x194e9c0. Original EVP_BytesToKey-compatible MD5 expansion.
// Each digest includes the preceding digest followed by the password; the
// returned slice retains append's allocation capacity and zero length is nil.
func ArNrGS_EV(password string, length int) []byte {
	var key, previous []byte
	hash := md5.New()
	for len(key) < length {
		hash.Write(previous)
		hash.Write([]byte(password))
		key = hash.Sum(key)
		previous = key[len(key)-md5.Size:]
		hash.Reset()
	}
	return key[:length]
}
