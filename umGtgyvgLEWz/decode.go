// Package umGtgyvgLEWz connects the TOML dependency identified in the binary.
// The parser/lexer/error layouts and Decode/parse control flow match
// github.com/BurntSushi/toml v1.4.0. The original dependency's exact commit
// is not present in the executable's stripped module information.
package umGtgyvgLEWz

import "github.com/BurntSushi/toml"

// XykYEQs is the 96-byte metadata value returned by 0x18c72a0 on amd64.
type XykYEQs = toml.MetaData

// Q2Erylds4g corresponds to the reader-backed decoder at 0x18c7480.
type Q2Erylds4g = toml.Decoder

// Ikm5Xa restores 0x18c72a0: string reader -> Decoder.Decode, with both
// the metadata and error returned unchanged.
func Ikm5Xa(data string, dst any) (XykYEQs, error) {
	return toml.Decode(data, dst)
}
