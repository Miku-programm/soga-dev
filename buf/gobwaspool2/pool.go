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
package gobwaspool2

import (
	pmath "recovered.local/soga/buf/gobwaspool4"
	"sync"
)

type LVPKRSqpONvB struct {
	lxoG0T       map[int]*sync.Pool
	jI0Cpy13WusE func(int) int
}
type qp7HhDhYVKA LVPKRSqpONvB
type MpotdXINa2 interface {
	AddSize(int)
	SetSizeMapping(func(int) int)
}

// Option 是原函数参数的源码表达；该函数类型没有保留原命名 RTTI。
type Option func(MpotdXINa2)

var poolAt5096780 = L06Zy06(128, 65536)

func L06Zy06(min, max int) *LVPKRSqpONvB {
	return QPgdbVKbwji(R_bnSZT(pmath.WDHhd0lpnI3), FP4M0qQ6(min, max))
}
func QPgdbVKbwji(opts ...Option) *LVPKRSqpONvB {
	p := &LVPKRSqpONvB{lxoG0T: make(map[int]*sync.Pool), jI0Cpy13WusE: pmath.NH6pIGDz}
	c := (*qp7HhDhYVKA)(p)
	for _, opt := range opts {
		opt(c)
	}
	return p
}
func (p *LVPKRSqpONvB) Get(size int) (any, int) {
	n := p.jI0Cpy13WusE(size)
	if q := p.lxoG0T[n]; q != nil {
		return q.Get(), n
	}
	return nil, size
}

// Put 的逻辑在原调用者内联，由 0x11ca8a0 的 map 查询及 sync.Pool.Put 恢复。
func (p *LVPKRSqpONvB) Put(v any, size int) {
	if q := p.lxoG0T[size]; q != nil {
		q.Put(v)
	}
}
func (p *qp7HhDhYVKA) AddSize(n int)                  { p.lxoG0T[n] = new(sync.Pool) }
func (p *qp7HhDhYVKA) SetSizeMapping(f func(int) int) { p.jI0Cpy13WusE = f }

// 构造本身已内联；以下闭包分别对应 0x11a02e0 / 0x11a03a0。
func FP4M0qQ6(min, max int) Option {
	return func(c MpotdXINa2) { pmath.YxT95tOm(min, max, func(n int) { c.AddSize(n) }) }
}
func R_bnSZT(f func(int) int) Option { return func(c MpotdXINa2) { c.SetSizeMapping(f) } }
