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
	"bytes"
	"encoding/binary"
	"math/rand"
)

const (
	MaxControlFramePayloadSize = 125
)

type Ao3txMm8 byte

const (
	OpContinuation Ao3txMm8 = 0x0
	OpText         Ao3txMm8 = 0x1
	OpBinary       Ao3txMm8 = 0x2
	OpClose        Ao3txMm8 = 0x8
	OpPing         Ao3txMm8 = 0x9
	OpPong         Ao3txMm8 = 0xa
)

func (c Ao3txMm8) IsControl() bool {

	return c&0x8 != 0
}

func (c Ao3txMm8) IsData() bool {

	return c&0x8 == 0
}

func (c Ao3txMm8) IsReserved() bool {

	return (0x3 <= c && c <= 0x7) || (0xb <= c && c <= 0xf)
}

type K9jDRT9Z uint16

type T9tLfPri0v struct {
	HHa11Na0i, TCmULKM K9jDRT9Z
}

var (
	StatusRangeNotInUse    = T9tLfPri0v{0, 999}
	StatusRangeProtocol    = T9tLfPri0v{1000, 2999}
	StatusRangeApplication = T9tLfPri0v{3000, 3999}
	StatusRangePrivate     = T9tLfPri0v{4000, 4999}
)

const (
	StatusNormalClosure           K9jDRT9Z = 1000
	StatusGoingAway               K9jDRT9Z = 1001
	StatusProtocolError           K9jDRT9Z = 1002
	StatusUnsupportedData         K9jDRT9Z = 1003
	StatusNoMeaningYet            K9jDRT9Z = 1004
	StatusInvalidFramePayloadData K9jDRT9Z = 1007
	StatusPolicyViolation         K9jDRT9Z = 1008
	StatusMessageTooBig           K9jDRT9Z = 1009
	StatusMandatoryExt            K9jDRT9Z = 1010
	StatusInternalServerError     K9jDRT9Z = 1011
	StatusTLSHandshake            K9jDRT9Z = 1015

	StatusAbnormalClosure K9jDRT9Z = 1006

	StatusNoStatusRcvd K9jDRT9Z = 1005
)

func (s K9jDRT9Z) In(r T9tLfPri0v) bool {
	return r.HHa11Na0i <= s && s <= r.TCmULKM
}

func (s K9jDRT9Z) Empty() bool {
	return s == 0
}

func (s K9jDRT9Z) IsNotUsed() bool {
	return s.In(StatusRangeNotInUse)
}

func (s K9jDRT9Z) IsApplicationSpec() bool {
	return s.In(StatusRangeApplication)
}

func (s K9jDRT9Z) IsPrivateSpec() bool {
	return s.In(StatusRangePrivate)
}

func (s K9jDRT9Z) IsProtocolSpec() bool {
	return s.In(StatusRangeProtocol)
}

func (s K9jDRT9Z) IsProtocolDefined() bool {
	switch s {
	case StatusNormalClosure,
		StatusGoingAway,
		StatusProtocolError,
		StatusUnsupportedData,
		StatusInvalidFramePayloadData,
		StatusPolicyViolation,
		StatusMessageTooBig,
		StatusMandatoryExt,
		StatusInternalServerError,
		StatusNoStatusRcvd,
		StatusAbnormalClosure,
		StatusTLSHandshake:
		return true
	}
	return false
}

func (s K9jDRT9Z) IsProtocolReserved() bool {
	switch s {

	case StatusNoStatusRcvd, StatusAbnormalClosure, StatusTLSHandshake:
		return true
	default:
		return false
	}
}

type Wmyq5K0sQ0w struct {
	HcIhwPLvD    bool
	PwwYpeWQ     byte
	IsHe8C       Ao3txMm8
	X4IS_ER04TVv bool
	BiCmXtyECQl  [4]byte
	L4ixlgT8VpTo int64
}

func (h Wmyq5K0sQ0w) Rsv1() bool { return h.PwwYpeWQ&bit5 != 0 }

func (h Wmyq5K0sQ0w) Rsv2() bool { return h.PwwYpeWQ&bit6 != 0 }

func (h Wmyq5K0sQ0w) Rsv3() bool { return h.PwwYpeWQ&bit7 != 0 }

type PFIuXmv struct {
	IBIsKFqaR  Wmyq5K0sQ0w
	CUdi3A1jBq []byte
}

func NewFrame(op Ao3txMm8, fin bool, p []byte) PFIuXmv {
	return PFIuXmv{
		IBIsKFqaR: Wmyq5K0sQ0w{
			HcIhwPLvD:    fin,
			IsHe8C:       op,
			L4ixlgT8VpTo: int64(len(p)),
		},
		CUdi3A1jBq: p,
	}
}

func NewCloseFrame(p []byte) PFIuXmv {
	return NewFrame(OpClose, true, p)
}

func NewCloseFrameBody(code K9jDRT9Z, reason string) []byte {
	n := min(2+len(reason), MaxControlFramePayloadSize)
	p := make([]byte, n)

	crop := min(MaxControlFramePayloadSize-2, len(reason))
	PutCloseFrameBody(p, code, reason[:crop])

	return p
}

func PutCloseFrameBody(p []byte, code K9jDRT9Z, reason string) {
	_ = p[1+len(reason)]
	binary.BigEndian.PutUint16(p, uint16(code))
	copy(p[2:], reason)
}

func GU1uDCm70(f PFIuXmv) PFIuXmv {
	return MaskFrameInPlaceWith(f, NewMask())
}

func MaskFrameInPlaceWith(f PFIuXmv, m [4]byte) PFIuXmv {
	f.IBIsKFqaR.X4IS_ER04TVv = true
	f.IBIsKFqaR.BiCmXtyECQl = m
	W2gws0VQo(f.CUdi3A1jBq, m, 0)
	return f
}

func NewMask() (ret [4]byte) {
	binary.BigEndian.PutUint32(ret[:], rand.Uint32())
	return ret
}

func DN2uvj(f PFIuXmv) (bts []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, 16))
	err = VzPOf1_5QN(buf, f)
	bts = buf.Bytes()
	return bts, err
}
