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
package wsutil

import (
	"io"

	ws "recovered.local/soga/transport/websocket"
)

type T3e3P4wOxooC struct {
	zGmeMGooB io.Reader
	xVfRjn07o [4]byte
	lu_5t1TLE int
}

func NewCipherReader(zGmeMGooB io.Reader, xVfRjn07o [4]byte) *T3e3P4wOxooC {
	return &T3e3P4wOxooC{zGmeMGooB, xVfRjn07o, 0}
}

func (c *T3e3P4wOxooC) Reset(zGmeMGooB io.Reader, xVfRjn07o [4]byte) {
	c.zGmeMGooB = zGmeMGooB
	c.xVfRjn07o = xVfRjn07o
	c.lu_5t1TLE = 0
}

func (c *T3e3P4wOxooC) Read(p []byte) (n int, err error) {
	n, err = c.zGmeMGooB.Read(p)
	ws.W2gws0VQo(p[:n], c.xVfRjn07o, c.lu_5t1TLE)
	c.lu_5t1TLE += n
	return n, err
}
