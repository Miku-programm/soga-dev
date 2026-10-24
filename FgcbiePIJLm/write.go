/*
The MIT License (MIT)

Copyright (c) 2017-2021 Sergey Kamardin <gobwas@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
// 依据原 ELF 的 RTTI、IDA 与 gobwas/ws v1.4.0 对照恢复。
// 主函数机器码已匹配；原被内联的小辅助表达式保留明确恢复名称，不计独立原函数。
package FgcbiePIJLm

import (
	"encoding/binary"
	"io"
)

const (
	MaxHeaderSize = 14
	MinHeaderSize = 2
)

const (
	bit0 = 0x80
	bit1 = 0x40
	bit2 = 0x20
	bit3 = 0x10
	bit4 = 0x08
	bit5 = 0x04
	bit6 = 0x02
	bit7 = 0x01

	len7  = int64(125)
	len16 = int64(^(uint16(0)))
	len64 = int64(^(uint64(0)) >> 1)
)

func HeaderSize(h Wmyq5K0sQ0w) (n int) {
	switch {
	case h.L4ixlgT8VpTo < 126:
		n = 2
	case h.L4ixlgT8VpTo <= len16:
		n = 4
	case h.L4ixlgT8VpTo <= len64:
		n = 10
	default:
		return -1
	}
	if h.X4IS_ER04TVv {
		n += len(h.BiCmXtyECQl)
	}
	return n
}

func KzLwHzhWB90h(w io.Writer, h Wmyq5K0sQ0w) error {

	bts := make([]byte, MaxHeaderSize)

	if h.HcIhwPLvD {
		bts[0] |= bit0
	}
	bts[0] |= h.PwwYpeWQ << 4
	bts[0] |= byte(h.IsHe8C)

	var n int
	switch {
	case h.L4ixlgT8VpTo <= len7:
		bts[1] = byte(h.L4ixlgT8VpTo)
		n = 2

	case h.L4ixlgT8VpTo <= len16:
		bts[1] = 126
		binary.BigEndian.PutUint16(bts[2:4], uint16(h.L4ixlgT8VpTo))
		n = 4

	case h.L4ixlgT8VpTo <= len64:
		bts[1] = 127
		binary.BigEndian.PutUint64(bts[2:10], uint64(h.L4ixlgT8VpTo))
		n = 10

	default:
		return ErrHeaderLengthUnexpected
	}

	if h.X4IS_ER04TVv {
		bts[1] |= bit0
		n += copy(bts[n:], h.BiCmXtyECQl[:])
	}

	_, err := w.Write(bts[:n])

	return err
}

func VzPOf1_5QN(w io.Writer, f PFIuXmv) error {
	err := KzLwHzhWB90h(w, f.IBIsKFqaR)
	if err != nil {
		return err
	}
	_, err = w.Write(f.CUdi3A1jBq)
	return err
}
