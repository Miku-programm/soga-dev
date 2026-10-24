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
	"io"
	ws "recovered.local/soga/transport/websocket"
)

// 构造内联；返回闭包对应 0x11d4180。
func L9DiQX(w io.Writer, state ws.DwfKMV) KTXwG8eroaXF {
	return func(h ws.Wmyq5K0sQ0w, r io.Reader) error {
		return (GdvjKg5RudlP{DisableSrcCiphering: true, Src: r, Dst: w, State: state}).Handle(h)
	}
}
