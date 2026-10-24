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

import (
	"bytes"
)

type ItemType int

const (
	ItemUndef ItemType = iota

	ItemToken

	ItemSeparator

	ItemString

	ItemComment

	ItemOctet
)

type K1cBAl struct {
	data []byte
	pos  int

	itemType  ItemType
	itemBytes []byte

	err bool
}

func (l *K1cBAl) Next() bool {
	c, ok := l.t9YEol3tOA()
	if !ok {
		return false
	}
	switch c {
	case '"':
		return l.irYm9NYUw1D()

	case '(':
		return l.fjYFxdPF()

	case '\\', ')':
		l.err = true
		return false

	default:
		return l.rQt3ptU0()
	}
}

func (l *K1cBAl) SkipEscaped(c byte) {
	if l.err {
		return
	}

	l.resetItem()

	if i := EQ1c264(l.data[l.pos:], c); i == -1 {

		l.pos = len(l.data)
	} else {
		l.pos += i + 1
	}
}

func (l *K1cBAl) Type() ItemType {
	return l.itemType
}

func (l *K1cBAl) Bytes() []byte {
	return l.itemBytes
}

func (l *K1cBAl) t9YEol3tOA() (byte, bool) {

	l.resetItem()

	if l.err {
		return 0, false
	}
	l.pos += SkipSpace(l.data[l.pos:])
	if l.pos == len(l.data) {
		return 0, false
	}
	return l.data[l.pos], true
}

func (l *K1cBAl) resetItem() {
	l.itemType = ItemUndef
	l.itemBytes = nil
}

func (l *K1cBAl) rQt3ptU0() bool {
	n, t := ScanToken(l.data[l.pos:])
	if n == -1 {
		l.err = true
		return false
	}

	l.itemType = t
	l.itemBytes = l.data[l.pos : l.pos+n]
	l.pos += n

	return true
}

func (l *K1cBAl) irYm9NYUw1D() (ok bool) {
	l.pos++

	n := EQ1c264(l.data[l.pos:], '"')
	if n == -1 {
		l.err = true
		return false
	}

	l.itemType = ItemString
	l.itemBytes = EFBQu4(l.data[l.pos:l.pos+n], '\\')
	l.pos += n + 1

	return true
}

func (l *K1cBAl) fjYFxdPF() (ok bool) {
	l.pos++

	n := AndxaUbuz(l.data[l.pos:], '(', ')')
	if n == -1 {
		l.err = true
		return false
	}

	l.itemType = ItemComment
	l.itemBytes = EFBQu4(l.data[l.pos:l.pos+n], '\\')
	l.pos += n + 1

	return true
}

func EQ1c264(data []byte, c byte) (n int) {
	for {
		i := bytes.IndexByte(data[n:], c)
		if i == -1 {
			return -1
		}
		n += i
		if n == 0 || data[n-1] != '\\' {
			break
		}
		n++
	}
	return
}

func AndxaUbuz(data []byte, open, close byte) (n int) {
	var m int
	opened := 1
	for {
		i := bytes.IndexByte(data[n:], close)
		if i == -1 {
			return -1
		}
		n += i

		if n == 0 || data[n-1] != '\\' {
			opened--
		}

		for m < i {
			j := bytes.IndexByte(data[m:i], open)
			if j == -1 {
				break
			}
			m += j + 1
			opened++
		}

		if opened == 0 {
			break
		}

		n++
		m = n
	}
	return
}

func EFBQu4(data []byte, c byte) []byte {
	j := bytes.IndexByte(data, c)
	if j == -1 {
		return data
	}

	n := len(data) - 1

	result := make([]byte, n)
	k := copy(result, data[:j])

	for i := j + 1; i < n; {
		j = bytes.IndexByte(data[i:], c)
		if j != -1 {
			k += copy(result[k:], data[i:i+j])
			i = i + j + 1
		} else {
			k += copy(result[k:], data[i:])
			break
		}
	}

	return result[:k]
}

func SkipSpace(p []byte) (n int) {
	for len(p) > 0 {
		switch {
		case len(p) >= 3 &&
			p[0] == '\r' &&
			p[1] == '\n' &&
			OctetTypes[p[2]].IsSpace():
			p = p[3:]
			n += 3
		case OctetTypes[p[0]].IsSpace():
			p = p[1:]
			n++
		default:
			return
		}
	}
	return
}

func ScanToken(p []byte) (n int, t ItemType) {
	if len(p) == 0 {
		return 0, ItemUndef
	}

	c := p[0]
	switch {
	case OctetTypes[c].IsSeparator():
		return 1, ItemSeparator

	case OctetTypes[c].IsToken():
		for n = 1; n < len(p); n++ {
			c := p[n]
			if !OctetTypes[c].IsToken() {
				break
			}
		}
		return n, ItemToken

	default:
		return -1, ItemUndef
	}
}
