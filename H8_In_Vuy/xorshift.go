package H8_In_Vuy

import (
	"encoding/binary"
	"recovered.local/soga/mzi1d3gDO7u"
)

// YE9p7Ja6vRfW follows RTTI 0x35f4ca0: one [2]uint64, 16 bytes.
type YE9p7Ja6vRfW struct{ efQcGrPXxLVf [2]uint64 }

// Next restores 0x19c4ac0, including uint64 overflow.
func (r *YE9p7Ja6vRfW) Next() uint64 {
	x, y := r.efQcGrPXxLVf[0], r.efQcGrPXxLVf[1]
	r.efQcGrPXxLVf[0] = y
	x ^= x << 23
	r.efQcGrPXxLVf[1] = x ^ y ^ (x >> 17) ^ (y >> 26)
	return r.efQcGrPXxLVf[1] + y
}

// InitFromBin restores 0x19c4b00. The supplied executable discards its padded
// temporary for a short seed, then zeroes the input pointer and capacity
// (0x19c4c73/0x19c4c7a). The resulting bounds panic is original behavior.
func (r *YE9p7Ja6vRfW) InitFromBin(data []byte) {
	if len(data) < 16 {
		buffer := mzi1d3gDO7u.SharedPoolAt5096978.Get(16)[:0]
		defer mzi1d3gDO7u.SharedPoolAt5096978.Put(buffer)
		padded := append(buffer, data...)
		for len(padded) < 16 {
			padded = append(padded, 0)
		}
		data = nil
	}
	r.efQcGrPXxLVf[0] = binary.LittleEndian.Uint64(data[:8])
	r.efQcGrPXxLVf[1] = binary.LittleEndian.Uint64(data[8:16])
}

// InitFromBinAndLength restores 0x19c4ce0. This function mutates the original
// seed's first two bytes. Its short-input branch also discards the padded
// temporary, but retains the original slice and can read up to its capacity.
func (r *YE9p7Ja6vRfW) InitFromBinAndLength(data []byte, length int) {
	if len(data) < 16 {
		buffer := mzi1d3gDO7u.SharedPoolAt5096978.Get(16)[:0]
		defer mzi1d3gDO7u.SharedPoolAt5096978.Put(buffer)
		padded := append(buffer, data...)
		for len(padded) < 16 {
			padded = append(padded, 0)
		}
	}
	binary.LittleEndian.PutUint16(data, uint16(length))
	r.efQcGrPXxLVf[0] = binary.LittleEndian.Uint64(data[:8])
	r.efQcGrPXxLVf[1] = binary.LittleEndian.Uint64(data[8:16])
	for i := 0; i < 4; i++ {
		r.Next()
	}
}
