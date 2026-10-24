package ssrhelper

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

var kdfBaseAt509A710 = []byte("VMess AEAD KDF")

// yzvWj4wW has the original single embedded hash.Hash interface. Its promoted
// methods preserve the wrappers at 0x181cbe0 through 0x181ce80.
type yzvWj4wW struct{ hash.Hash }

// CVqKigbE26 @ 0x181c1a0, including its original shared-hash constructor
// closure @ 0x181cb60. The wrapper gives the first interface a distinct dynamic
// type while retaining the same underlying hash state.
func CVqKigbE26(key []byte, path ...[]byte) []byte {
	h := hmac.New(sha256.New, kdfBaseAt509A710)
	for _, p := range path {
		first := true
		h = hmac.New(func() hash.Hash {
			if first {
				first = false
				return &yzvWj4wW{h}
			}
			return h
		}, p)
	}
	h.Write(key)
	return h.Sum(nil)
}
