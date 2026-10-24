package JafvCs5

import (
	"crypto/md5"
	"crypto/sha1"
	"golang.org/x/crypto/hkdf"
	"io"
)

// Is8IqhjVE6 restores 0x181bd40: EVP_BytesToKey's repeated MD5 construction.
// A zero length returns a nil slice. The final slice retains append capacity.
func Is8IqhjVE6(password string, keyLen int) []byte {
	h := md5.New()
	var key, last []byte
	for len(key) < keyLen {
		h.Write(last)
		h.Write([]byte(password))
		key = append(key, h.Sum(nil)...)
		last = key[len(key)-16:]
		h.Reset()
	}
	return key[:keyLen]
}

// THj8k5 restores 0x181bfc0. The hash funcval at 0x376a388 is sha1.New;
// io.ReadFull's error is propagated as a panic in the original function.
func THj8k5(key, salt, info, out []byte) {
	if _, err := io.ReadFull(hkdf.New(sha1.New, key, salt, info), out); err != nil {
		panic(err)
	}
}
