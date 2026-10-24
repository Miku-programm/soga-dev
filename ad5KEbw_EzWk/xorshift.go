package ad5KEbw_EzWk

import "encoding/binary"

// HelGdo2ETK follows RTTI 0x3631df0, two uint64 words (16 bytes).
type HelGdo2ETK struct{ mFYAwci, bWvBk2x4oPae uint64 }

// Next restores 0x1c844e0, including unsigned overflow.
func (r *HelGdo2ETK) Next() uint64 {
	x, y := r.mFYAwci, r.bWvBk2x4oPae
	r.mFYAwci = y
	x ^= x << 23
	r.bWvBk2x4oPae = x ^ y ^ (x >> 17) ^ (y >> 26)
	return r.bWvBk2x4oPae + y
}

// InitFromBytes restores 0x1c84520. Short seeds append zero bytes to the
// supplied slice, including its available backing capacity. This native
// implementation differs from the separate outbound H8_In_Vuy implementation.
func (r *HelGdo2ETK) InitFromBytes(data []byte) {
	for len(data) < 16 {
		data = append(data, 0)
	}
	r.mFYAwci = binary.LittleEndian.Uint64(data[:8])
	r.bWvBk2x4oPae = binary.LittleEndian.Uint64(data[8:])
}

// InitFromBytesLen restores 0x1c84680: a private 16-byte zeroed seed, a
// truncated little-endian length and four discarded xorshift outputs.
func (r *HelGdo2ETK) InitFromBytesLen(data []byte, length int) {
	var seed [16]byte
	copy(seed[:], data)
	binary.LittleEndian.PutUint16(seed[:], uint16(length))
	r.mFYAwci = binary.LittleEndian.Uint64(seed[:8])
	r.bWvBk2x4oPae = binary.LittleEndian.Uint64(seed[8:])
	for i := 0; i < 4; i++ {
		r.Next()
	}
}
