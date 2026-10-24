/*
The MIT License (MIT)

Copyright (c) 2017 Sergey Kamardin

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
// 根据原 ELF/IDA 和 gobwas/httphead v0.1.0 逐函数对照恢复；非作者原始文件。
// 存活类型保留 RTTI 名称；未存活的值类型及内联函数保留明确对照名称。
package httphead

type BIMZTmbYHb struct {
	KXa0Kt2I []byte
	Jd2j6aD7 MVgyIAGYmsKQ
}

func (opt BIMZTmbYHb) Size() int {
	return len(opt.KXa0Kt2I) + opt.Jd2j6aD7.l6UF09
}

func (opt BIMZTmbYHb) Copy(p []byte) BIMZTmbYHb {
	n := copy(p, opt.KXa0Kt2I)
	opt.KXa0Kt2I = p[:n]
	opt.Jd2j6aD7, p = opt.Jd2j6aD7.Copy(p[n:])
	return opt
}

type MVgyIAGYmsKQ struct {
	clXYVDW3J int
	l6UF09    int
	rraG37Ve  [8]blm5Wb
	hMg6gO_J  []blm5Wb
}

func (p *MVgyIAGYmsKQ) Size() int {
	return p.l6UF09
}

func (p *MVgyIAGYmsKQ) Copy(dst []byte) (MVgyIAGYmsKQ, []byte) {
	ret := MVgyIAGYmsKQ{
		clXYVDW3J: p.clXYVDW3J,
		l6UF09:    p.l6UF09,
	}
	if p.hMg6gO_J != nil {
		ret.hMg6gO_J = make([]blm5Wb, len(p.hMg6gO_J))
		for i, v := range p.hMg6gO_J {
			ret.hMg6gO_J[i], dst = v.copy(dst)
		}
	} else {
		for i, p := range p.rraG37Ve {
			ret.rraG37Ve[i], dst = p.copy(dst)
		}
	}
	return ret, dst
}

func (p *MVgyIAGYmsKQ) Set(z9Lbkony, yZSzEF []byte) {
	p.l6UF09 += len(z9Lbkony) + len(yZSzEF)

	if p.clXYVDW3J < len(p.rraG37Ve) {
		p.rraG37Ve[p.clXYVDW3J] = blm5Wb{z9Lbkony, yZSzEF}
		p.clXYVDW3J++
		return
	}

	if p.hMg6gO_J == nil {
		p.hMg6gO_J = make([]blm5Wb, len(p.rraG37Ve), len(p.rraG37Ve)+1)
		copy(p.hMg6gO_J, p.rraG37Ve[:])
	}
	p.hMg6gO_J = append(p.hMg6gO_J, blm5Wb{z9Lbkony, yZSzEF})
}

func (p *MVgyIAGYmsKQ) data() []blm5Wb {
	if p.hMg6gO_J != nil {
		return p.hMg6gO_J
	}
	return p.rraG37Ve[:p.clXYVDW3J]
}

type blm5Wb struct {
	z9Lbkony, yZSzEF []byte
}

func (p blm5Wb) copy(dst []byte) (blm5Wb, []byte) {
	n := copy(dst, p.z9Lbkony)
	p.z9Lbkony = dst[:n]
	m := n + copy(dst[n:], p.yZSzEF)
	p.yZSzEF = dst[n:m]

	dst = dst[m:]

	return p, dst
}
