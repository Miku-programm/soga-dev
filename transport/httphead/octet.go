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

type OctetType byte

func (t OctetType) IsChar() bool { return t&octetChar != 0 }

func (t OctetType) IsControl() bool { return t&octetControl != 0 }

func (t OctetType) IsSeparator() bool { return t&octetSeparator != 0 }

func (t OctetType) IsSpace() bool { return t&octetSpace != 0 }

func (t OctetType) IsToken() bool { return t&octetToken != 0 }

const (
	octetChar OctetType = 1 << iota
	octetControl
	octetSpace
	octetSeparator
	octetToken
)

var OctetTypes [256]OctetType

func init() {
	for c := 32; c < 256; c++ {
		var t OctetType
		if c <= 127 {
			t |= octetChar
		}
		if 0 <= c && c <= 31 || c == 127 {
			t |= octetControl
		}
		switch c {
		case '(', ')', '<', '>', '@', ',', ';', ':', '"', '/', '[', ']', '?', '=', '{', '}', '\\':
			t |= octetSeparator
		case ' ', '\t':
			t |= octetSpace | octetSeparator
		}

		if t.IsChar() && !t.IsControl() && !t.IsSeparator() && !t.IsSpace() {
			t |= octetToken
		}

		OctetTypes[c] = t
	}
}
