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
// 原 0x11a0400..0x11a0b20 与 gobwas/pool v0.2.1 对照；非作者原始文件。
// Package pbufio contains tools for pooling bufio.Reader and bufio.Writers.
package gPJ6UJF

import (
	"bufio"
	"io"

	pool "recovered.local/soga/g8mMV5wjV"
)

var (
	DefaultQBjm6CMcar = NewQBjm6CMcar(256, 65536)
	DefaultRqHcmi     = NewRqHcmi(256, 65536)
)

// GetWriter returns bufio.Writer whose buffer has at least size bytes.
// Note that size could be ceiled to the next power of two.
// GetWriter is a wrapper around DefaultQBjm6CMcar.Get().
func GetWriter(w io.Writer, size int) *bufio.Writer { return DefaultQBjm6CMcar.Get(w, size) }

// PutWriter takes bufio.Writer for future reuse.
// It does not reuse bufio.Writer which underlying buffer size is not power of
// PutWriter is a wrapper around DefaultQBjm6CMcar.Put().
func PutWriter(bw *bufio.Writer) { DefaultQBjm6CMcar.Put(bw) }

// GetReader returns bufio.Reader whose buffer has at least size bytes. It returns
// its capacity for further pass to Put().
// Note that size could be ceiled to the next power of two.
// GetReader is a wrapper around DefaultRqHcmi.Get().
func GetReader(w io.Reader, size int) *bufio.Reader { return DefaultRqHcmi.Get(w, size) }

// PutReader takes bufio.Reader and its size for future reuse.
// It does not reuse bufio.Reader if size is not power of two or is out of pool
// min/max range.
// PutReader is a wrapper around DefaultRqHcmi.Put().
func PutReader(bw *bufio.Reader) { DefaultRqHcmi.Put(bw) }

// QBjm6CMcar contains logic of *bufio.Writer reuse with various size.
type QBjm6CMcar struct {
	pool *pool.LVPKRSqpONvB
}

// NewQBjm6CMcar creates new QBjm6CMcar that reuses writers which size is in
// logarithmic range [min, max].
func NewQBjm6CMcar(min, max int) *QBjm6CMcar {
	return &QBjm6CMcar{pool.L06Zy06(min, max)}
}

// CustomQBjm6CMcar creates new QBjm6CMcar with given options.

// Get returns bufio.Writer whose buffer has at least size bytes.
func (wp *QBjm6CMcar) Get(w io.Writer, size int) *bufio.Writer {
	v, n := wp.pool.Get(size)
	if v != nil {
		bw := v.(*bufio.Writer)
		bw.Reset(w)
		return bw
	}
	return bufio.NewWriterSize(w, n)
}

// Put takes ownership of bufio.Writer for further reuse.
func (wp *QBjm6CMcar) Put(bw *bufio.Writer) {
	// Should reset even if we do Reset() inside Get().
	// This is done to prevent locking underlying io.Writer from GC.
	bw.Reset(nil)
	wp.pool.Put(bw, bw.Size())
}

// RqHcmi contains logic of *bufio.Reader reuse with various size.
type RqHcmi struct {
	pool *pool.LVPKRSqpONvB
}

// NewRqHcmi creates new RqHcmi that reuses writers which size is in
// logarithmic range [min, max].
func NewRqHcmi(min, max int) *RqHcmi {
	return &RqHcmi{pool.L06Zy06(min, max)}
}

// CustomRqHcmi creates new RqHcmi with given options.

// Get returns bufio.Reader whose buffer has at least size bytes.
func (rp *RqHcmi) Get(r io.Reader, size int) *bufio.Reader {
	v, n := rp.pool.Get(size)
	if v != nil {
		br := v.(*bufio.Reader)
		br.Reset(r)
		return br
	}
	return bufio.NewReaderSize(r, n)
}

// Put takes ownership of bufio.Reader for further reuse.
func (rp *RqHcmi) Put(br *bufio.Reader) {
	// Should reset even if we do Reset() inside Get().
	// This is done to prevent locking underlying io.Reader from GC.
	br.Reset(nil)
	rp.pool.Put(br, br.Size())
}

// 原类型的私有字段名未存活；两个类都仅含偏移 0 的池指针。
