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
// 由原 ELF/IDA 与 gobwas/ws v1.4.0 对照恢复，非作者源码原文。
// 保留原 RTTI 名称；内联辅助表达式使用恢复名称，不计独立原函数。
package wsutil

import (
	"fmt"
	"io"

	ws "recovered.local/soga/transport/websocket"
	pbytes "recovered.local/soga/buf/gobwaspool"
)

var (
	ErrNotEmpty = fmt.Errorf("writer not empty")

	ErrControlOverflow = fmt.Errorf("control frame payload overflow")
)

const (
	len7  = int64(125)
	len16 = int64(^uint16(0))
	len64 = int64((^uint64(0)) >> 1)
)

type DH32d4Xf struct {
	mpzgoFUbRt *M6JkkEsBtZH
	jJkHg4kv   int
	fZvXjIM    int
}

func ZB0fR0Aw(dest io.Writer, state ws.DwfKMV, op ws.Ao3txMm8, buf []byte) *DH32d4Xf {
	max := ws.MaxControlFramePayloadSize + headerSize(state, ws.MaxControlFramePayloadSize)
	if len(buf) > max {
		buf = buf[:max]
	}

	mpzgoFUbRt := NewWriterBuffer(dest, state, op, buf)

	return &DH32d4Xf{
		mpzgoFUbRt: mpzgoFUbRt,
		jJkHg4kv:   len(mpzgoFUbRt.ariAbujn0),
	}
}

func (c *DH32d4Xf) Write(p []byte) (fZvXjIM int, err error) {
	if c.fZvXjIM+len(p) > c.jJkHg4kv {
		return 0, ErrControlOverflow
	}
	return c.mpzgoFUbRt.Write(p)
}

func (c *DH32d4Xf) Flush() error {
	return c.mpzgoFUbRt.Flush()
}

type M6JkkEsBtZH struct {
	bEEZIz5 io.Writer

	pIitEG ws.Ao3txMm8

	sSjoMgATz0p ws.DwfKMV

	v78A5VtJX []KzT9kPu3yJD

	bKFgx3c bool

	gGZQkcSDF1jp []byte

	ariAbujn0 []byte

	jfnerh_x6 int

	e_fTv76YzJG  bool
	vB4Yl8PWQC   int
	vQT2iWLEILFo error
}

func NewWriterBuffer(bEEZIz5 io.Writer, sSjoMgATz0p ws.DwfKMV, pIitEG ws.Ao3txMm8, ariAbujn0 []byte) *M6JkkEsBtZH {
	w := &M6JkkEsBtZH{
		bEEZIz5:      bEEZIz5,
		sSjoMgATz0p:  sSjoMgATz0p,
		pIitEG:       pIitEG,
		gGZQkcSDF1jp: ariAbujn0,
	}
	w.cJBSsW7X9F()
	return w
}

func (w *M6JkkEsBtZH) cJBSsW7X9F() {
	offset := reserve(w.sSjoMgATz0p, len(w.gGZQkcSDF1jp))
	if len(w.gGZQkcSDF1jp) <= offset {
		panic("wsutil: writer buffer is too small")
	}
	w.ariAbujn0 = w.gGZQkcSDF1jp[offset:]
}

func (w *M6JkkEsBtZH) Reset(bEEZIz5 io.Writer, sSjoMgATz0p ws.DwfKMV, pIitEG ws.Ao3txMm8) {
	w.bEEZIz5 = bEEZIz5
	w.sSjoMgATz0p = sSjoMgATz0p
	w.pIitEG = pIitEG

	w.cJBSsW7X9F()

	w.jfnerh_x6 = 0
	w.e_fTv76YzJG = false
	w.vB4Yl8PWQC = 0
	w.v78A5VtJX = w.v78A5VtJX[:0]
	w.bKFgx3c = false
}

func (w *M6JkkEsBtZH) ResetOp(pIitEG ws.Ao3txMm8) {
	w.pIitEG = pIitEG
	w.jfnerh_x6 = 0
	w.e_fTv76YzJG = false
	w.vB4Yl8PWQC = 0
}

func (w *M6JkkEsBtZH) SetExtensions(xs ...KzT9kPu3yJD) {
	w.v78A5VtJX = xs
}

func (w *M6JkkEsBtZH) DisableFlush() {
	w.bKFgx3c = true
}

func (w *M6JkkEsBtZH) Size() int {
	return len(w.ariAbujn0)
}

func (w *M6JkkEsBtZH) Available() int {
	return len(w.ariAbujn0) - w.jfnerh_x6
}

func (w *M6JkkEsBtZH) Buffered() int {
	return w.jfnerh_x6
}

func (w *M6JkkEsBtZH) Write(p []byte) (jfnerh_x6 int, vQT2iWLEILFo error) {

	w.e_fTv76YzJG = true

	var nn int
	for len(p) > w.Available() && w.vQT2iWLEILFo == nil {
		if w.bKFgx3c {
			w.Grow(len(p))
			continue
		}
		if w.Buffered() == 0 {

			nn, _ = w.WriteThrough(p)
		} else {
			nn = copy(w.ariAbujn0[w.jfnerh_x6:], p)
			w.jfnerh_x6 += nn
			w.FlushFragment()
		}
		jfnerh_x6 += nn
		p = p[nn:]
	}
	if w.vQT2iWLEILFo != nil {
		return jfnerh_x6, w.vQT2iWLEILFo
	}
	nn = copy(w.ariAbujn0[w.jfnerh_x6:], p)
	w.jfnerh_x6 += nn
	jfnerh_x6 += nn

	return jfnerh_x6, w.vQT2iWLEILFo
}

func ceilPowerOfTwo(jfnerh_x6 int) int {
	jfnerh_x6 |= jfnerh_x6 >> 1
	jfnerh_x6 |= jfnerh_x6 >> 2
	jfnerh_x6 |= jfnerh_x6 >> 4
	jfnerh_x6 |= jfnerh_x6 >> 8
	jfnerh_x6 |= jfnerh_x6 >> 16
	jfnerh_x6 |= jfnerh_x6 >> 32
	jfnerh_x6++
	return jfnerh_x6
}

func (w *M6JkkEsBtZH) Grow(jfnerh_x6 int) {

	var (
		size       = len(w.gGZQkcSDF1jp)
		prevOffset = len(w.gGZQkcSDF1jp) - len(w.ariAbujn0)
		nextOffset = len(w.gGZQkcSDF1jp) - len(w.ariAbujn0)
		buffered   = w.Buffered()
	)
	for cap := size - nextOffset - buffered; cap < jfnerh_x6; {

		size = ceilPowerOfTwo(nextOffset + buffered + jfnerh_x6)
		nextOffset = reserve(w.sSjoMgATz0p, size)
		cap = size - nextOffset - buffered
	}
	if size < len(w.gGZQkcSDF1jp) {
		panic("wsutil: buffer grow leads to its reduce")
	}
	if size == len(w.gGZQkcSDF1jp) {
		return
	}
	p := make([]byte, size)
	copy(p[nextOffset-prevOffset:], w.gGZQkcSDF1jp[:prevOffset+buffered])
	w.gGZQkcSDF1jp = p
	w.ariAbujn0 = w.gGZQkcSDF1jp[nextOffset:]
}

func (w *M6JkkEsBtZH) WriteThrough(p []byte) (jfnerh_x6 int, vQT2iWLEILFo error) {
	if w.vQT2iWLEILFo != nil {
		return 0, w.vQT2iWLEILFo
	}
	if w.Buffered() != 0 {
		return 0, ErrNotEmpty
	}

	var frame ws.PFIuXmv
	frame.IBIsKFqaR = ws.Wmyq5K0sQ0w{
		IsHe8C:       w.opCode(),
		HcIhwPLvD:    false,
		L4ixlgT8VpTo: int64(len(p)),
	}
	for _, x := range w.v78A5VtJX {
		frame.IBIsKFqaR, vQT2iWLEILFo = x.SetBits(frame.IBIsKFqaR)
		if vQT2iWLEILFo != nil {
			return 0, vQT2iWLEILFo
		}
	}
	if w.sSjoMgATz0p.ClientSide() {

		payload := pbytes.GetLen(len(p))
		defer pbytes.KqwWvOwoRkN(payload)
		copy(payload, p)

		frame.CUdi3A1jBq = payload
		frame = ws.GU1uDCm70(frame)
	} else {
		frame.CUdi3A1jBq = p
	}

	w.vQT2iWLEILFo = ws.VzPOf1_5QN(w.bEEZIz5, frame)
	if w.vQT2iWLEILFo == nil {
		jfnerh_x6 = len(p)
	}

	w.e_fTv76YzJG = true
	w.vB4Yl8PWQC++

	return jfnerh_x6, w.vQT2iWLEILFo
}

func (w *M6JkkEsBtZH) ReadFrom(src io.Reader) (jfnerh_x6 int64, vQT2iWLEILFo error) {
	var nn int
	for vQT2iWLEILFo == nil {
		if w.Available() == 0 {
			if w.bKFgx3c {
				w.Grow(w.Buffered())
			} else {
				vQT2iWLEILFo = w.FlushFragment()
			}
			continue
		}

		const maxEmptyReads = 100
		var nr int
		for nr < maxEmptyReads {
			nn, vQT2iWLEILFo = src.Read(w.ariAbujn0[w.jfnerh_x6:])
			if nn != 0 || vQT2iWLEILFo != nil {
				break
			}
			nr++
		}
		if nr == maxEmptyReads {
			return jfnerh_x6, io.ErrNoProgress
		}

		w.jfnerh_x6 += nn
		jfnerh_x6 += int64(nn)
	}
	if vQT2iWLEILFo == io.EOF {

		vQT2iWLEILFo = nil
		w.e_fTv76YzJG = true
	}
	return jfnerh_x6, vQT2iWLEILFo
}

func (w *M6JkkEsBtZH) Flush() error {
	if (!w.e_fTv76YzJG && w.Buffered() == 0) || w.vQT2iWLEILFo != nil {
		return w.vQT2iWLEILFo
	}

	w.vQT2iWLEILFo = w.xGAnyTaY7pd(true)
	w.jfnerh_x6 = 0
	w.e_fTv76YzJG = false
	w.vB4Yl8PWQC = 0

	return w.vQT2iWLEILFo
}

func (w *M6JkkEsBtZH) FlushFragment() error {
	if w.Buffered() == 0 || w.vQT2iWLEILFo != nil {
		return w.vQT2iWLEILFo
	}

	w.vQT2iWLEILFo = w.xGAnyTaY7pd(false)
	w.jfnerh_x6 = 0
	w.vB4Yl8PWQC++

	return w.vQT2iWLEILFo
}

func (w *M6JkkEsBtZH) xGAnyTaY7pd(fin bool) (vQT2iWLEILFo error) {
	var (
		payload = w.ariAbujn0[:w.jfnerh_x6]
		header  = ws.Wmyq5K0sQ0w{
			IsHe8C:       w.opCode(),
			HcIhwPLvD:    fin,
			L4ixlgT8VpTo: int64(len(payload)),
		}
	)
	for _, ext := range w.v78A5VtJX {
		header, vQT2iWLEILFo = ext.SetBits(header)
		if vQT2iWLEILFo != nil {
			return vQT2iWLEILFo
		}
	}
	if w.sSjoMgATz0p.ClientSide() {
		header.X4IS_ER04TVv = true
		header.BiCmXtyECQl = ws.NewMask()
		ws.W2gws0VQo(payload, header.BiCmXtyECQl, 0)
	}

	var (
		offset = len(w.gGZQkcSDF1jp) - len(w.ariAbujn0)
		skip   = offset - ws.HeaderSize(header)
	)
	ariAbujn0 := zvrEv0{
		b5FO9d: w.gGZQkcSDF1jp[skip:offset],
	}
	if vQT2iWLEILFo := ws.KzLwHzhWB90h(&ariAbujn0, header); vQT2iWLEILFo != nil {

		panic("dump header error: " + vQT2iWLEILFo.Error())
	}
	_, vQT2iWLEILFo = w.bEEZIz5.Write(w.gGZQkcSDF1jp[skip : offset+w.jfnerh_x6])
	return vQT2iWLEILFo
}

func (w *M6JkkEsBtZH) opCode() ws.Ao3txMm8 {
	if w.vB4Yl8PWQC > 0 {
		return ws.OpContinuation
	}
	return w.pIitEG
}

var errNoSpace = fmt.Errorf("not enough buffer space")

type zvrEv0 struct {
	b5FO9d      []byte
	lFd2tGa8_bb int
}

func (w *zvrEv0) Write(p []byte) (int, error) {
	n := copy(w.b5FO9d[w.lFd2tGa8_bb:], p)
	w.lFd2tGa8_bb += n
	if n != len(p) {
		return n, errNoSpace
	}
	return n, nil
}

func z6Gqa5XP(w io.Writer, s ws.DwfKMV, pIitEG ws.Ao3txMm8, fin bool, p []byte) error {
	var frame ws.PFIuXmv
	if s.ClientSide() {

		payload := pbytes.GetLen(len(p))
		defer pbytes.KqwWvOwoRkN(payload)

		copy(payload, p)

		frame = ws.NewFrame(pIitEG, fin, payload)
		frame = ws.GU1uDCm70(frame)
	} else {
		frame = ws.NewFrame(pIitEG, fin, p)
	}

	return ws.VzPOf1_5QN(w, frame)
}

func reserve(sSjoMgATz0p ws.DwfKMV, jfnerh_x6 int) (offset int) {
	var mask int
	if sSjoMgATz0p.ClientSide() {
		mask = 4
	}
	switch {
	case jfnerh_x6 <= int(len7)+mask+2:
		return mask + 2
	case jfnerh_x6 <= int(len16)+mask+4:
		return mask + 4
	default:
		return mask + 10
	}
}

func headerSize(s ws.DwfKMV, jfnerh_x6 int) int {
	return ws.HeaderSize(ws.Wmyq5K0sQ0w{
		L4ixlgT8VpTo: int64(jfnerh_x6),
		X4IS_ER04TVv: s.ClientSide(),
	})
}
