/*
MIT License

Copyright (c) 2016-2017 Daniel Fu

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

// 依据原 ELF（SHA256 6660dcfb…）IDA/RTTI 恢复；并非作者原文件。
// 以 MIT 许可 smux v1.5.24 为可比对基线，原 fork 的批量发送、锁范围、
// 版本协商、异步窗口更新及缓冲池差异均按机器码恢复，不能用上游包直接替代。
// RTTI 显示前缀 ji2LadBZQZW6，uncommon package path 为 I4AWsAkBLRP。
package realityconn

import (
	"encoding/binary"
)

const (
	cmdSYN byte = iota
	cmdFIN
	cmdPSH
	cmdNOP

	cmdUPD
)

const (
	szCmdUPD = 8
)

const (
	initialPeerWindow = 262144
)

const (
	sizeOfVer    = 1
	sizeOfCmd    = 1
	sizeOfLength = 2
	sizeOfSid    = 4
	headerSize   = sizeOfVer + sizeOfCmd + sizeOfSid + sizeOfLength
)

type TKUtOq struct {
	xlaU77zvaW7W byte
	gUa4uFKR6    byte
	a_9mrWgHy    uint32
	cAKGyUaj     []byte
}

// 此辅助名未保留在元数据中；函数体来自调用者内联指令，非业务占位。
func newFrame(version byte, gUa4uFKR6 byte, a_9mrWgHy uint32) TKUtOq {
	return TKUtOq{xlaU77zvaW7W: version, gUa4uFKR6: gUa4uFKR6, a_9mrWgHy: a_9mrWgHy}
}

type rawHeader [headerSize]byte

func (h rawHeader) RTNOROJC5() byte {
	return h[0]
}

func (h rawHeader) Cmd() byte {
	return h[1]
}

func (h rawHeader) Length() uint16 {
	return binary.LittleEndian.Uint16(h[2:])
}

func (h rawHeader) StreamID() uint32 {
	return binary.LittleEndian.Uint32(h[4:])
}

type xnabvi1sOhR [szCmdUPD]byte

func (h xnabvi1sOhR) Consumed() uint32 {
	return binary.LittleEndian.Uint32(h[:])
}
func (h xnabvi1sOhR) Window() uint32 {
	return binary.LittleEndian.Uint32(h[4:])
}
