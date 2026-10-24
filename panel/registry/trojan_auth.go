package registry

import (
	"crypto/sha256"
	"encoding/hex"
)

// J26YsI2DwxPQ @ 0x15cc460: lowercase hexadecimal SHA-224 of the password.
func J26YsI2DwxPQ(password string) string {
	digest := sha256.Sum224([]byte(password))
	return hex.EncodeToString(digest[:])
}
