package bytes

// PutSlice: 0x192b680. Grow precedes the overlapping-safe copy.
func (b *BaDekw7qe) PutSlice(p []byte) {
	i := b.Grow(len(p))
	copy((*b)[i:i+len(p)], p)
}

// Bytes: 0x192bae0. The returned slice aliases the buffer.
func (b *BaDekw7qe) Bytes() []byte { return *b }

// Reset: 0x192bb00. Keep the allocation and nil/non-nil distinction.
func (b *BaDekw7qe) Reset() { *b = (*b)[:0] }
