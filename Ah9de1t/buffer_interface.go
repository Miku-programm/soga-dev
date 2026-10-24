// 原RTTI 0x3702e80的完整29方法接口，显示前缀x2NpasZPR9b，实际包Ah9de1t。
package Ah9de1t

import (
	"crypto/cipher"
	"io"
)

type DwBIUXLAa interface {
	Bytes() []uint8
	BytesFrom(int) []uint8
	Free()
	IsEmpty() bool
	Len() int
	Open(cipher.AEAD, []uint8, []uint8) (int, error)
	OpenWrite(cipher.AEAD, []uint8, []uint8, []uint8) (int, error)
	Read([]uint8) (int, error)
	ReadAllFrom(io.Reader, int) (int, error)
	ReadBytes(int) ([]uint8, error)
	ReadFrom(io.Reader) (int, error)
	ReadFullFrom(io.Reader, int) (int, error)
	ReadMaxFrom(io.Reader, int) (int, error)
	ReadUint16() (uint16, error)
	ReadUint16Le() (uint16, error)
	ReadUint32() (uint32, error)
	ReadUint32Le() (uint32, error)
	ReadUint64() (uint64, error)
	ReadUint8() (uint8, error)
	ReverseEnd(int) int
	Seal(cipher.AEAD, []uint8, []uint8)
	SealWrite(cipher.AEAD, []uint8, []uint8, []uint8) int
	Write([]uint8) (int, error)
	WriteByte(uint8)
	WriteRandom(int) []uint8
	WriteString(string)
	WriteUint16(uint16)
	WriteUint32(uint32)
	WriteUint64(uint64)
}
