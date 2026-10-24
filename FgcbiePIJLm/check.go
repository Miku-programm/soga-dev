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

import "unicode/utf8"

type DwfKMV uint8

const (
	StateServerSide DwfKMV = 0x1 << iota

	StateClientSide

	StateExtended

	StateFragmented
)

func (s DwfKMV) Is(v DwfKMV) bool {
	return uint8(s)&uint8(v) != 0
}

func (s DwfKMV) Set(v DwfKMV) DwfKMV {
	return s | v
}

func (s DwfKMV) Clear(v DwfKMV) DwfKMV {
	return s & (^v)
}

func (s DwfKMV) ServerSide() bool { return s.Is(StateServerSide) }

func (s DwfKMV) ClientSide() bool { return s.Is(StateClientSide) }

func (s DwfKMV) Extended() bool { return s.Is(StateExtended) }

func (s DwfKMV) Fragmented() bool { return s.Is(StateFragmented) }

type FsEzVo5mX string

func (p FsEzVo5mX) Error() string { return string(p) }

var (
	ErrProtocolOpCodeReserved             = FsEzVo5mX("use of reserved op code")
	ErrProtocolControlPayloadOverflow     = FsEzVo5mX("control frame payload limit exceeded")
	ErrProtocolControlNotFinal            = FsEzVo5mX("control frame is not final")
	ErrProtocolNonZeroRsv                 = FsEzVo5mX("non-zero rsv bits with no extension negotiated")
	ErrProtocolMaskRequired               = FsEzVo5mX("frames from client to server must be masked")
	ErrProtocolMaskUnexpected             = FsEzVo5mX("frames from server to client must be not masked")
	ErrProtocolContinuationExpected       = FsEzVo5mX("unexpected non-continuation data frame")
	ErrProtocolContinuationUnexpected     = FsEzVo5mX("unexpected continuation data frame")
	ErrProtocolStatusCodeNotInUse         = FsEzVo5mX("status code is not in use")
	ErrProtocolStatusCodeApplicationLevel = FsEzVo5mX("status code is only application level")
	ErrProtocolStatusCodeNoMeaning        = FsEzVo5mX("status code has no meaning yet")
	ErrProtocolStatusCodeUnknown          = FsEzVo5mX("status code is not defined in spec")
	ErrProtocolInvalidUTF8                = FsEzVo5mX("invalid utf8 sequence in close reason")
)

func IYI3RQDfj4(h Wmyq5K0sQ0w, s DwfKMV) error {
	if h.IsHe8C.IsReserved() {
		return ErrProtocolOpCodeReserved
	}
	if h.IsHe8C.IsControl() {
		if h.L4ixlgT8VpTo > MaxControlFramePayloadSize {
			return ErrProtocolControlPayloadOverflow
		}
		if !h.HcIhwPLvD {
			return ErrProtocolControlNotFinal
		}
	}

	switch {

	case h.PwwYpeWQ != 0 && !s.Extended():
		return ErrProtocolNonZeroRsv

	case s.ServerSide() && !h.X4IS_ER04TVv:
		return ErrProtocolMaskRequired
	case s.ClientSide() && h.X4IS_ER04TVv:
		return ErrProtocolMaskUnexpected

	case s.Fragmented() && !h.IsHe8C.IsControl() && h.IsHe8C != OpContinuation:
		return ErrProtocolContinuationExpected
	case !s.Fragmented() && h.IsHe8C == OpContinuation:
		return ErrProtocolContinuationUnexpected

	default:
		return nil
	}
}

func QXKFKa(code K9jDRT9Z, reason string) error {
	switch {
	case code.IsNotUsed():
		return ErrProtocolStatusCodeNotInUse

	case code.IsProtocolReserved():
		return ErrProtocolStatusCodeApplicationLevel

	case code == StatusNoMeaningYet:
		return ErrProtocolStatusCodeNoMeaning

	case code.IsProtocolSpec() && !code.IsProtocolDefined():
		return ErrProtocolStatusCodeUnknown

	case !utf8.ValidString(reason):
		return ErrProtocolInvalidUTF8

	default:
		return nil
	}
}
