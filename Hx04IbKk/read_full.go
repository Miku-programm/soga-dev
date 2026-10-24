package Hx04IbKk

import "io"

// ReadFull: 0x192b8e0. Failed reads restore the previous logical length;
// partial bytes and any capacity growth remain in the underlying allocation.
func (b *BaDekw7qe) ReadFull(r io.Reader, n int) error {
	old := len(*b)
	pos := b.Grow(n)
	_, err := io.ReadFull(r, (*b)[pos:pos+n])
	if err != nil {
		*b = (*b)[:old]
	}
	return err
}
