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
// 原 ELF 0x11b8760 的 raw Upgrade 及解析链逐分支对照 gobwas/ws v1.4.0 恢复。
// 只恢复机器码中已证实的函数；未存活的局部类型和字段使用明确恢复名。
package FgcbiePIJLm

import (
	"bytes"
	httphead "recovered.local/soga/EsS8DL8PE9"
)

type httpRequestLine struct {
	method, uri  []byte
	major, minor int
}

func wBe3m0MV0Cy(line []byte) (req httpRequestLine, err error) {
	var proto []byte
	req.method, req.uri, proto = cU_pjSUO(line, ' ')

	var ok bool
	req.major, req.minor, ok = lJiBRXY(proto)
	if !ok {
		err = ErrMalformedRequest
	}
	return req, err
}

func lJiBRXY(bts []byte) (major, minor int, ok bool) {
	switch {
	case bytes.Equal(bts, httpVersion1_0):
		return 1, 0, true
	case bytes.Equal(bts, httpVersion1_1):
		return 1, 1, true
	case len(bts) < 8:
		return 0, 0, false
	case !bytes.Equal(bts[:5], httpVersionPrefix):
		return 0, 0, false
	}

	bts = bts[5:]

	dot := bytes.IndexByte(bts, '.')
	if dot == -1 {
		return 0, 0, false
	}
	var err error
	major, err = fm9r93qPp(bts[:dot])
	if err != nil {
		return major, 0, false
	}
	minor, err = fm9r93qPp(bts[dot+1:])
	if err != nil {
		return major, minor, false
	}

	return major, minor, true
}

func mdOcFVt(line []byte) (k, v []byte, ok bool) {
	colon := bytes.IndexByte(line, ':')
	if colon == -1 {
		return nil, nil, false
	}

	k = btrim(line[:colon])

	canonicalizeHeaderKey(k)

	v = btrim(line[colon+1:])

	return k, v, true
}

func cyqy5Vn4u(h []byte, check func([]byte) bool) (ret string, ok bool) {
	var selected []byte
	ok = httphead.PgwXhS7T(h, func(v []byte) bool {
		if check(v) {
			selected = v
			return false
		}
		return true
	})
	if ok && selected != nil {
		return string(selected), true
	}
	return ret, ok
}
