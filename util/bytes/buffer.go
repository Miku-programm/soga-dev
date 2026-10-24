package bytes

import "bytes"

// BaDekw7qe is the original named []byte type, not bytes.Buffer.
type BaDekw7qe []byte

// Grow: 0x192b0a0. It extends the length and returns the old length. Negative
// counts can shrink the slice; the original leaves bounds enforcement to slicing.
func (b *BaDekw7qe) Grow(n int) int {
	m := len(*b)
	if n > cap(*b)-m {
		return b.hahJfwxTr35o(n)
	}
	*b = (*b)[:m+n]
	return m
}

// hahJfwxTr35o: 0x192af20.
func (b *BaDekw7qe) hahJfwxTr35o(n int) int {
	m, c := len(*b), cap(*b)
	if n <= c-m {
		*b = (*b)[:m+n]
		return m
	}
	if *b == nil && n <= 64 {
		*b = make([]byte, n, 64)
		return 0
	}
	if c > int(^uint(0)>>1)-c-n {
		panic(bytes.ErrTooLarge)
	}
	if n > c/2-m {
		*b = k9c6jLHi(*b, n)
	}
	*b = (*b)[:m+n]
	return m
}

// k9c6jLHi: 0x192ad60; recovery closure: 0x192bb20. The original allocation
// uses append to a nil slice, so Go's allocation-size capacity rounding is retained.
func k9c6jLHi(b []byte, n int) []byte {
	defer func() {
		if recover() != nil {
			panic(bytes.ErrTooLarge)
		}
	}()
	size := len(b) + n
	if size < 2*cap(b) {
		size = 2 * cap(b)
	}
	result := append([]byte(nil), make([]byte, size)...)
	copy(result, b)
	return result[:len(b)]
}
