// Package Ualx5UM_D6I 恢复原样本缓冲接口的真实包路径；RTTI 显示前缀为 wPyWyp。
package Ualx5UM_D6I

import (
	"crypto/cipher"
	"io"
)

// Mf3phSOVqgf 来自 RTTI 0x3713920 的完整38方法集，未缩小为使用点的子集。
type Mf3phSOVqgf interface {
	Bytes() []byte
	BytesFrom(int) []byte
	BytesRange(int, int) []byte
	BytesTo(int) []byte
	Cap() int
	Free()
	GrowBy(int)
	IsEmpty() bool
	Len() int
	Open(cipher.AEAD, []byte, []byte) (int, error)
	OpenWrite(cipher.AEAD, []byte, []byte, []byte) (int, error)
	Read([]byte) (int, error)
	ReadAllFrom(io.Reader, int) (int, error)
	ReadBytes(int) ([]byte, error)
	ReadFrom(io.Reader) (int, error)
	ReadFullFrom(io.Reader, int) (int, error)
	ReadMaxFrom(io.Reader, int) (int, error)
	ReadUint16() (uint16, error)
	ReadUint16Le() (uint16, error)
	ReadUint32() (uint32, error)
	ReadUint32Le() (uint32, error)
	ReadUint64() (uint64, error)
	ReadUint8() (uint8, error)
	Reset()
	ReverseEnd(int) int
	Seal(cipher.AEAD, []byte, []byte)
	SealWrite(cipher.AEAD, []byte, []byte, []byte) int
	SealWriteWithStart(cipher.AEAD, int, []byte, []byte, []byte) int
	SkipEnd(int)
	Write([]byte) (int, error)
	WriteByte(uint8)
	WriteRandom(int) []byte
	WriteString(string)
	WriteUint16(uint16)
	WriteUint32(uint32)
	WriteUint64(uint64)
	XORKeyStreamRead(cipher.Stream, []byte) (int, error)
	XORKeyStreamWrite(cipher.Stream, []byte) int
}
