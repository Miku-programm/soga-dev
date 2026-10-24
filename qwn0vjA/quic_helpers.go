package qwn0vjA

import (
	"crypto"
	"encoding/binary"
	"golang.org/x/crypto/hkdf"
	"io"
)

// vFUaVI1hy 对应0x18786a0。读取首字节给出的1/2/4/8字节无符号数；
// 非最短表示可被接受；任意后续 ReadByte 错误使数值返回零。
func vFUaVI1hy(r io.ByteReader) (uint64, error) {
	first, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	n := 1 << ((first & 0xc0) >> 6)
	v := uint64(first & 0x3f)
	for i := 1; i < n; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		v = (v << 8) + uint64(b)
	}
	return v, nil
}

// qhXufgHMZAdy 对应0x1878360。保留原标签长度窄化、单次读取和短读 panic。
func qhXufgHMZAdy(hash crypto.Hash, secret, context []byte, label string, length int) []byte {
	info := make([]byte, 3, len(label)+len(context)+10)
	binary.BigEndian.PutUint16(info, uint16(length))
	info[2] = byte(len(label) + 6)
	info = append(info, "tls13 "...)
	info = append(info, label...)
	info = append(info, byte(len(context)))
	info = append(info, context...)
	out := make([]byte, length)
	n, _ := hkdf.Expand(hash.New, secret, info).Read(out)
	if n != length {
		panic("quic: HKDF-Expand-Label invocation failed unexpectedly")
	}
	return out
}
