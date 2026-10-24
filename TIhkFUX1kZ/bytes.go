/*
The MIT License (MIT)

Copyright (c) 2017-2019 Sergey Kamardin <gobwas@gmail.com>

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
// 由原 ELF/IDA 与 gobwas/pool v0.2.1 对照恢复，非作者源码原文。
// 保留原 RTTI 名称；内联辅助表达式使用恢复名称，不计独立原函数。
package TIhkFUX1kZ

import pool "recovered.local/soga/g8mMV5wjV"

// 类型名见 0x11ca940 元数据，私有字段名未保留；8B 指针布局来自其机器码。
type E0PtY88jkHw struct{ poolAtOffset0 *pool.LVPKRSqpONvB }

var defaultPoolAt5029D70 = &E0PtY88jkHw{pool.L06Zy06(128, 65536)}

func (p *E0PtY88jkHw) Get(n, c int) []byte {
	if n > c {
		panic("requested length is greater than capacity")
	}
	v, x := p.poolAtOffset0.Get(c)
	if v != nil {
		return v.([]byte)[:n]
	}
	return make([]byte, n, x)
}
func KqwWvOwoRkN(b []byte) { defaultPoolAt5029D70.poolAtOffset0.Put(b, cap(b)) }

// GetLen 是调用者原内联表达式的源码接口，不是额外独立原函数。
func GetLen(n int) []byte { return defaultPoolAt5029D70.Get(n, n) }
