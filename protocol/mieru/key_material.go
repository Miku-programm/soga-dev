package mieru

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

// gUBL99 @0x1c393e0: the NUL separator is present even with an empty salt.
func gUBL99(password string, salt []byte) *[32]byte {
	b := append([]byte(password), 0)
	b = append(b, salt...)
	sum := sha256.Sum256(b)
	return &sum
}

// pVQO7MqTE4ao @0x1c39580. The order is previous, current, next two-minute
// slot after time.Round; each timestamp is an eight-byte big endian integer.
func pVQO7MqTE4ao(now time.Time) (out [3][]byte) {
	rounded := now.Round(2 * time.Minute)
	for i, offset := range [...]time.Duration{-2 * time.Minute, 0, 2 * time.Minute} {
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], uint64(rounded.Add(offset).Unix()))
		sum := sha256.Sum256(b[:])
		out[i] = sum[:]
	}
	return
}

// pyUlp642nl @0x1c397a0 compares the first four digest bytes with the last
// four nonce bytes. The original comparison is bytes.Equal, not constant-time.
func pyUlp642nl(password string, nonce []byte) bool {
	if len(nonce) < 20 {
		return false
	}
	sum := sha256.Sum256(append([]byte(password), nonce[:16]...))
	return bytes.Equal(sum[:4], nonce[len(nonce)-4:])
}

// o0fNBhk @0x1c54340. Entropy errors propagate here; padding elsewhere ignores
// them. The generated 24-byte nonce carries a password tag only when nonempty.
func o0fNBhk(password string) ([]byte, error) {
	nonce := make([]byte, 24)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	if password != "" {
		uJghIeaJPse(nonce, password)
	}
	return nonce, nil
}

// uJghIeaJPse @0x1c54400. It always writes at offset 20, at most four bytes;
// the verifier above instead reads the final four. Preserve that distinction
// for unusual slice lengths and the original slicing panics.
func uJghIeaJPse(nonce []byte, password string) {
	sum := sha256.Sum256(append([]byte(password), nonce[:16]...))
	copy(nonce[20:], sum[:4])
}
