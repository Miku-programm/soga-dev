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
package EsS8DL8PE9

import "io"

var (
	comma     = []byte{','}
	equality  = []byte{'='}
	semicolon = []byte{';'}
	quote     = []byte{'"'}
	escape    = []byte{'\\'}
)

func KBgYhFImpuOF(dest io.Writer, options []BIMZTmbYHb) (n int, err error) {
	w := i2Jgi93li{w: dest}
	for i, opt := range options {
		if i > 0 {
			w.xLiYco3gck1(comma)
		}

		eU0XAX(&w, opt.KXa0Kt2I)

		for _, p := range opt.Jd2j6aD7.data() {
			w.xLiYco3gck1(semicolon)
			eU0XAX(&w, p.z9Lbkony)
			if len(p.yZSzEF) != 0 {
				w.xLiYco3gck1(equality)
				eU0XAX(&w, p.yZSzEF)
			}
		}
	}
	return w.result()
}

func eU0XAX(bw *i2Jgi93li, bts []byte) {
	var qt bool
	var pos int
	for i := 0; i < len(bts); i++ {
		c := bts[i]
		if !OctetTypes[c].IsToken() && !qt {
			qt = true
			bw.xLiYco3gck1(quote)
		}
		if OctetTypes[c].IsControl() || c == '"' {
			if !qt {
				qt = true
				bw.xLiYco3gck1(quote)
			}
			bw.xLiYco3gck1(bts[pos:i])
			bw.xLiYco3gck1(escape)
			bw.xLiYco3gck1(bts[i : i+1])
			pos = i + 1
		}
	}
	if !qt {
		bw.xLiYco3gck1(bts)
	} else {
		bw.xLiYco3gck1(bts[pos:])
		bw.xLiYco3gck1(quote)
	}
}

type i2Jgi93li struct {
	w   io.Writer
	n   int
	err error
}

func (w *i2Jgi93li) xLiYco3gck1(p []byte) {
	if w.err != nil {
		return
	}
	var n int
	n, w.err = w.w.Write(p)
	w.n += n
	return
}

func (w *i2Jgi93li) result() (int, error) {
	return w.n, w.err
}
