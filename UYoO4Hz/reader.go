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
package UYoO4Hz

import (
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"

	ws "recovered.local/soga/FgcbiePIJLm"
)

var ErrNoFrameAdvance = errors.New("no frame advance")

var ErrFrameTooLarge = errors.New("frame too large")

type KTXwG8eroaXF func(ws.Wmyq5K0sQ0w, io.Reader) error

type HLKO3Cdqw struct {
	GX04GFKh1Gz io.Reader
	WXUh5psqY7  ws.DwfKMV

	QGRPzIklOviT bool

	CaUjxvGB2Ne bool

	JNz1DAm1v []Lwh4ZhXkF4

	ZbPIcS int64

	SIMJzjvB    KTXwG8eroaXF
	EbVEHBvY4lh KTXwG8eroaXF

	ya34HnplNfq ws.Ao3txMm8
	iGvszJg     io.Reader
	w7jqWX      io.LimitedReader
	a9xrva7p    Wvd84IIKB
	bEc6Yalku   [ws.MaxHeaderSize - 2]byte
	sEnaudEEn   *T3e3P4wOxooC
}

func (r *HLKO3Cdqw) Read(p []byte) (n int, err error) {
	if r.iGvszJg == nil {
		if !r.fragmented() {

			return 0, ErrNoFrameAdvance
		}

		_, err := r.NextFrame()
		if err != nil {
			return 0, err
		}
		if r.iGvszJg == nil {

			return 0, nil
		}
	}

	n, err = r.iGvszJg.Read(p)
	if err != nil && err != io.EOF {
		return n, err
	}
	if err == nil && r.w7jqWX.N != 0 {
		return n, nil
	}

	switch {
	case r.w7jqWX.N != 0:
		err = io.ErrUnexpectedEOF

	case r.fragmented():
		err = nil
		r.resetFragment()

	case r.CaUjxvGB2Ne && !r.a9xrva7p.Valid():

		n = r.a9xrva7p.Accepted()
		err = ErrInvalidUTF8

	default:
		r.reset()
		err = io.EOF
	}

	return n, err
}

func (r *HLKO3Cdqw) Discard() (err error) {
	for {
		_, err = io.Copy(ioutil.Discard, &r.w7jqWX)
		if err != nil {
			break
		}
		if !r.fragmented() {
			break
		}
		if _, err = r.NextFrame(); err != nil {
			break
		}
	}
	r.reset()
	return err
}

func (r *HLKO3Cdqw) NextFrame() (hdr ws.Wmyq5K0sQ0w, err error) {
	hdr, err = r.anP9kfyom(r.GX04GFKh1Gz)
	if err == io.EOF && r.fragmented() {

		err = io.ErrUnexpectedEOF
	}
	if err == nil && !r.QGRPzIklOviT {
		err = ws.IYI3RQDfj4(hdr, r.WXUh5psqY7)
	}
	if err != nil {
		return hdr, err
	}

	if n := r.ZbPIcS; n > 0 && hdr.L4ixlgT8VpTo > n {
		return hdr, ErrFrameTooLarge
	}

	r.w7jqWX = io.LimitedReader{
		R: r.GX04GFKh1Gz,
		N: hdr.L4ixlgT8VpTo,
	}

	iGvszJg := io.Reader(&r.w7jqWX)
	if hdr.X4IS_ER04TVv {
		if r.sEnaudEEn == nil {
			r.sEnaudEEn = NewCipherReader(iGvszJg, hdr.BiCmXtyECQl)
		} else {
			r.sEnaudEEn.Reset(iGvszJg, hdr.BiCmXtyECQl)
		}
		iGvszJg = r.sEnaudEEn
	}

	for _, x := range r.JNz1DAm1v {
		hdr, err = x.UnsetBits(hdr)
		if err != nil {
			return hdr, err
		}
	}

	if r.fragmented() {
		if hdr.IsHe8C.IsControl() {
			if cb := r.EbVEHBvY4lh; cb != nil {
				err = cb(hdr, iGvszJg)
			}
			if err == nil {

				_, err = io.Copy(ioutil.Discard, &r.w7jqWX)
			}
			return hdr, err
		}
	} else {
		r.ya34HnplNfq = hdr.IsHe8C
	}
	if r.CaUjxvGB2Ne && (hdr.IsHe8C == ws.OpText || (r.fragmented() && r.ya34HnplNfq == ws.OpText)) {
		r.a9xrva7p.RPKyxC = iGvszJg
		iGvszJg = &r.a9xrva7p
	}

	r.iGvszJg = iGvszJg

	if hdr.IsHe8C == ws.OpContinuation {
		if cb := r.SIMJzjvB; cb != nil {
			err = cb(hdr, iGvszJg)
		}
	}

	if hdr.HcIhwPLvD {
		r.WXUh5psqY7 = r.WXUh5psqY7.Clear(ws.StateFragmented)
	} else {
		r.WXUh5psqY7 = r.WXUh5psqY7.Set(ws.StateFragmented)
	}

	return hdr, err
}

func (r *HLKO3Cdqw) fragmented() bool {
	return r.WXUh5psqY7.Fragmented()
}

func (r *HLKO3Cdqw) resetFragment() {
	r.w7jqWX = io.LimitedReader{}
	r.iGvszJg = nil

	r.a9xrva7p.RPKyxC = nil
}

func (r *HLKO3Cdqw) reset() {
	r.w7jqWX = io.LimitedReader{}
	r.iGvszJg = nil
	r.a9xrva7p = Wvd84IIKB{}
	r.ya34HnplNfq = 0
}

func (r *HLKO3Cdqw) anP9kfyom(in io.Reader) (h ws.Wmyq5K0sQ0w, err error) {

	bts := r.bEc6Yalku[:2]

	_, err = io.ReadFull(in, bts)
	if err != nil {
		return h, err
	}
	const bit0 = 0x80

	h.HcIhwPLvD = bts[0]&bit0 != 0
	h.PwwYpeWQ = (bts[0] & 0x70) >> 4
	h.IsHe8C = ws.Ao3txMm8(bts[0] & 0x0f)

	var extra int

	if bts[1]&bit0 != 0 {
		h.X4IS_ER04TVv = true
		extra += 4
	}

	length := bts[1] & 0x7f
	switch {
	case length < 126:
		h.L4ixlgT8VpTo = int64(length)

	case length == 126:
		extra += 2

	case length == 127:
		extra += 8

	default:
		err = ws.ErrHeaderLengthUnexpected
		return h, err
	}

	if extra == 0 {
		return h, err
	}

	bts = bts[:extra]
	_, err = io.ReadFull(in, bts)
	if err != nil {
		return h, err
	}

	switch {
	case length == 126:
		h.L4ixlgT8VpTo = int64(binary.BigEndian.Uint16(bts[:2]))
		bts = bts[2:]

	case length == 127:
		if bts[0]&0x80 != 0 {
			err = ws.ErrHeaderLengthMSB
			return h, err
		}
		h.L4ixlgT8VpTo = int64(binary.BigEndian.Uint64(bts[:8]))
		bts = bts[8:]
	}

	if h.X4IS_ER04TVv {
		copy(h.BiCmXtyECQl[:], bts)
	}

	return h, nil
}
