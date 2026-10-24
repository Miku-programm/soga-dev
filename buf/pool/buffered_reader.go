// 原 NADyavA 的 Go bufio 派生实现：字段与业务入口恢复原符号。
// 对照原指令保留共享池分配、Reset 自引用行为、WriteTo 空写及 FreeBuf。
// 以下保留所用 Go 标准库源码的完整许可。
/*
Copyright 2009 The Go Authors.
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:
   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google LLC nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package pool

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode/utf8"
)

const (
	defaultBufSize = 4096
)

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

type NADyavA struct {
	bkCmHugyPy                []byte
	c6HUuPfxb526              io.Reader
	rfe2UwaF1FkL, inKjzRUMN3d int
	lsDnEVuZ                  error
	lS7kcz                    int
	zyE274D49tb               int
}

const minReadBufferSize = 16
const maxConsecutiveEmptyReads = 100

func C8FauP3Auy(c6HUuPfxb526 io.Reader, size int) *NADyavA {
	b, ok := c6HUuPfxb526.(*NADyavA)
	if ok && len(b.bkCmHugyPy) >= size {
		return b
	}
	rfe2UwaF1FkL := new(NADyavA)
	rfe2UwaF1FkL.reset(PoolAt50961D8.Get(max(size, minReadBufferSize)), c6HUuPfxb526)
	return rfe2UwaF1FkL
}
func (b *NADyavA) Size() int { return len(b.bkCmHugyPy) }
func (b *NADyavA) Reset(rfe2UwaF1FkL io.Reader) {
	if b.bkCmHugyPy == nil {
		b.bkCmHugyPy = PoolAt50961D8.Get(defaultBufSize)
	}
	b.reset(b.bkCmHugyPy, rfe2UwaF1FkL)
}
func (b *NADyavA) reset(bkCmHugyPy []byte, rfe2UwaF1FkL io.Reader) {
	*b = NADyavA{
		bkCmHugyPy:   bkCmHugyPy,
		c6HUuPfxb526: rfe2UwaF1FkL,
		lS7kcz:       -1,
		zyE274D49tb:  -1,
	}
}

var errNegativeRead = errors.New("bufio: reader returned negative count from Read")

func (b *NADyavA) rwazSq1Jm() {
	if b.rfe2UwaF1FkL > 0 {
		copy(b.bkCmHugyPy, b.bkCmHugyPy[b.rfe2UwaF1FkL:b.inKjzRUMN3d])
		b.inKjzRUMN3d -= b.rfe2UwaF1FkL
		b.rfe2UwaF1FkL = 0
	}
	if b.inKjzRUMN3d >= len(b.bkCmHugyPy) {
		panic("bufio: tried to rwazSq1Jm full buffer")
	}
	for i := maxConsecutiveEmptyReads; i > 0; i-- {
		n, lsDnEVuZ := b.c6HUuPfxb526.Read(b.bkCmHugyPy[b.inKjzRUMN3d:])
		if n < 0 {
			panic(errNegativeRead)
		}
		b.inKjzRUMN3d += n
		if lsDnEVuZ != nil {
			b.lsDnEVuZ = lsDnEVuZ
			return
		}
		if n > 0 {
			return
		}
	}
	b.lsDnEVuZ = io.ErrNoProgress
}
func (b *NADyavA) readErr() error {
	lsDnEVuZ := b.lsDnEVuZ
	b.lsDnEVuZ = nil
	return lsDnEVuZ
}
func (b *NADyavA) Peek(n int) ([]byte, error) {
	if n < 0 {
		return nil, ErrNegativeCount
	}
	b.lS7kcz = -1
	b.zyE274D49tb = -1
	for b.inKjzRUMN3d-b.rfe2UwaF1FkL < n && b.inKjzRUMN3d-b.rfe2UwaF1FkL < len(b.bkCmHugyPy) && b.lsDnEVuZ == nil {
		b.rwazSq1Jm()
	}
	if n > len(b.bkCmHugyPy) {
		return b.bkCmHugyPy[b.rfe2UwaF1FkL:b.inKjzRUMN3d], ErrBufferFull
	}
	var lsDnEVuZ error
	if avail := b.inKjzRUMN3d - b.rfe2UwaF1FkL; avail < n {
		n = avail
		lsDnEVuZ = b.readErr()
		if lsDnEVuZ == nil {
			lsDnEVuZ = ErrBufferFull
		}
	}
	return b.bkCmHugyPy[b.rfe2UwaF1FkL : b.rfe2UwaF1FkL+n], lsDnEVuZ
}
func (b *NADyavA) Discard(n int) (discarded int, lsDnEVuZ error) {
	if n < 0 {
		return 0, ErrNegativeCount
	}
	if n == 0 {
		return
	}
	b.lS7kcz = -1
	b.zyE274D49tb = -1
	remain := n
	for {
		skip := b.Buffered()
		if skip == 0 {
			b.rwazSq1Jm()
			skip = b.Buffered()
		}
		if skip > remain {
			skip = remain
		}
		b.rfe2UwaF1FkL += skip
		remain -= skip
		if remain == 0 {
			return n, nil
		}
		if b.lsDnEVuZ != nil {
			return n - remain, b.readErr()
		}
	}
}
func (b *NADyavA) Read(p []byte) (n int, lsDnEVuZ error) {
	n = len(p)
	if n == 0 {
		if b.Buffered() > 0 {
			return 0, nil
		}
		return 0, b.readErr()
	}
	if b.rfe2UwaF1FkL == b.inKjzRUMN3d {
		if b.lsDnEVuZ != nil {
			return 0, b.readErr()
		}
		if len(p) >= len(b.bkCmHugyPy) {
			n, b.lsDnEVuZ = b.c6HUuPfxb526.Read(p)
			if n < 0 {
				panic(errNegativeRead)
			}
			if n > 0 {
				b.lS7kcz = int(p[n-1])
				b.zyE274D49tb = -1
			}
			return n, b.readErr()
		}
		b.rfe2UwaF1FkL = 0
		b.inKjzRUMN3d = 0
		n, b.lsDnEVuZ = b.c6HUuPfxb526.Read(b.bkCmHugyPy)
		if n < 0 {
			panic(errNegativeRead)
		}
		if n == 0 {
			return 0, b.readErr()
		}
		b.inKjzRUMN3d += n
	}
	n = copy(p, b.bkCmHugyPy[b.rfe2UwaF1FkL:b.inKjzRUMN3d])
	b.rfe2UwaF1FkL += n
	b.lS7kcz = int(b.bkCmHugyPy[b.rfe2UwaF1FkL-1])
	b.zyE274D49tb = -1
	return n, nil
}
func (b *NADyavA) ReadByte() (byte, error) {
	b.zyE274D49tb = -1
	for b.rfe2UwaF1FkL == b.inKjzRUMN3d {
		if b.lsDnEVuZ != nil {
			return 0, b.readErr()
		}
		b.rwazSq1Jm()
	}
	c := b.bkCmHugyPy[b.rfe2UwaF1FkL]
	b.rfe2UwaF1FkL++
	b.lS7kcz = int(c)
	return c, nil
}
func (b *NADyavA) UnreadByte() error {
	if b.lS7kcz < 0 || b.rfe2UwaF1FkL == 0 && b.inKjzRUMN3d > 0 {
		return ErrInvalidUnreadByte
	}
	if b.rfe2UwaF1FkL > 0 {
		b.rfe2UwaF1FkL--
	} else {
		b.inKjzRUMN3d = 1
	}
	b.bkCmHugyPy[b.rfe2UwaF1FkL] = byte(b.lS7kcz)
	b.lS7kcz = -1
	b.zyE274D49tb = -1
	return nil
}
func (b *NADyavA) ReadRune() (rfe2UwaF1FkL rune, size int, lsDnEVuZ error) {
	for b.rfe2UwaF1FkL+utf8.UTFMax > b.inKjzRUMN3d && !utf8.FullRune(b.bkCmHugyPy[b.rfe2UwaF1FkL:b.inKjzRUMN3d]) && b.lsDnEVuZ == nil && b.inKjzRUMN3d-b.rfe2UwaF1FkL < len(b.bkCmHugyPy) {
		b.rwazSq1Jm()
	}
	b.zyE274D49tb = -1
	if b.rfe2UwaF1FkL == b.inKjzRUMN3d {
		return 0, 0, b.readErr()
	}
	rfe2UwaF1FkL, size = utf8.DecodeRune(b.bkCmHugyPy[b.rfe2UwaF1FkL:b.inKjzRUMN3d])
	b.rfe2UwaF1FkL += size
	b.lS7kcz = int(b.bkCmHugyPy[b.rfe2UwaF1FkL-1])
	b.zyE274D49tb = size
	return rfe2UwaF1FkL, size, nil
}
func (b *NADyavA) UnreadRune() error {
	if b.zyE274D49tb < 0 || b.rfe2UwaF1FkL < b.zyE274D49tb {
		return ErrInvalidUnreadRune
	}
	b.rfe2UwaF1FkL -= b.zyE274D49tb
	b.lS7kcz = -1
	b.zyE274D49tb = -1
	return nil
}
func (b *NADyavA) Buffered() int { return b.inKjzRUMN3d - b.rfe2UwaF1FkL }
func (b *NADyavA) ReadSlice(delim byte) (line []byte, lsDnEVuZ error) {
	s := 0
	for {
		if i := bytes.IndexByte(b.bkCmHugyPy[b.rfe2UwaF1FkL+s:b.inKjzRUMN3d], delim); i >= 0 {
			i += s
			line = b.bkCmHugyPy[b.rfe2UwaF1FkL : b.rfe2UwaF1FkL+i+1]
			b.rfe2UwaF1FkL += i + 1
			break
		}
		if b.lsDnEVuZ != nil {
			line = b.bkCmHugyPy[b.rfe2UwaF1FkL:b.inKjzRUMN3d]
			b.rfe2UwaF1FkL = b.inKjzRUMN3d
			lsDnEVuZ = b.readErr()
			break
		}
		if b.Buffered() >= len(b.bkCmHugyPy) {
			b.rfe2UwaF1FkL = b.inKjzRUMN3d
			line = b.bkCmHugyPy
			lsDnEVuZ = ErrBufferFull
			break
		}
		s = b.inKjzRUMN3d - b.rfe2UwaF1FkL
		b.rwazSq1Jm()
	}
	if i := len(line) - 1; i >= 0 {
		b.lS7kcz = int(line[i])
		b.zyE274D49tb = -1
	}
	return
}
func (b *NADyavA) ReadLine() (line []byte, isPrefix bool, lsDnEVuZ error) {
	line, lsDnEVuZ = b.ReadSlice('\n')
	if lsDnEVuZ == ErrBufferFull {
		if len(line) > 0 && line[len(line)-1] == '\r' {
			if b.rfe2UwaF1FkL == 0 {
				panic("bufio: tried to rewind past start of buffer")
			}
			b.rfe2UwaF1FkL--
			line = line[:len(line)-1]
		}
		return line, true, nil
	}
	if len(line) == 0 {
		if lsDnEVuZ != nil {
			line = nil
		}
		return
	}
	lsDnEVuZ = nil
	if line[len(line)-1] == '\n' {
		drop := 1
		if len(line) > 1 && line[len(line)-2] == '\r' {
			drop = 2
		}
		line = line[:len(line)-drop]
	}
	return
}
func (b *NADyavA) cIV8s3xFuN(delim byte) (fullBuffers [][]byte, finalFragment []byte, totalLen int, lsDnEVuZ error) {
	var frag []byte
	for {
		var e error
		frag, e = b.ReadSlice(delim)
		if e == nil {
			break
		}
		if e != ErrBufferFull {
			lsDnEVuZ = e
			break
		}
		bkCmHugyPy := bytes.Clone(frag)
		fullBuffers = append(fullBuffers, bkCmHugyPy)
		totalLen += len(bkCmHugyPy)
	}
	totalLen += len(frag)
	return fullBuffers, frag, totalLen, lsDnEVuZ
}
func (b *NADyavA) ReadBytes(delim byte) ([]byte, error) {
	full, frag, n, lsDnEVuZ := b.cIV8s3xFuN(delim)
	bkCmHugyPy := make([]byte, n)
	n = 0
	for i := range full {
		n += copy(bkCmHugyPy[n:], full[i])
	}
	copy(bkCmHugyPy[n:], frag)
	return bkCmHugyPy, lsDnEVuZ
}
func (b *NADyavA) ReadString(delim byte) (string, error) {
	full, frag, n, lsDnEVuZ := b.cIV8s3xFuN(delim)
	var bkCmHugyPy strings.Builder
	bkCmHugyPy.Grow(n)
	for _, fb := range full {
		bkCmHugyPy.Write(fb)
	}
	bkCmHugyPy.Write(frag)
	return bkCmHugyPy.String(), lsDnEVuZ
}
func (b *NADyavA) WriteTo(inKjzRUMN3d io.Writer) (n int64, lsDnEVuZ error) {
	b.lS7kcz = -1
	b.zyE274D49tb = -1
	n, lsDnEVuZ = b.im509M7Md(inKjzRUMN3d)
	if lsDnEVuZ != nil {
		return
	}
	if rfe2UwaF1FkL, ok := b.c6HUuPfxb526.(io.WriterTo); ok {
		m, lsDnEVuZ := rfe2UwaF1FkL.WriteTo(inKjzRUMN3d)
		n += m
		return n, lsDnEVuZ
	}
	if inKjzRUMN3d, ok := inKjzRUMN3d.(io.ReaderFrom); ok {
		m, lsDnEVuZ := inKjzRUMN3d.ReadFrom(b.c6HUuPfxb526)
		n += m
		return n, lsDnEVuZ
	}
	if b.inKjzRUMN3d-b.rfe2UwaF1FkL < len(b.bkCmHugyPy) {
		b.rwazSq1Jm()
	}
	for b.rfe2UwaF1FkL < b.inKjzRUMN3d {
		m, lsDnEVuZ := b.im509M7Md(inKjzRUMN3d)
		n += m
		if lsDnEVuZ != nil {
			return n, lsDnEVuZ
		}
		b.rwazSq1Jm()
	}
	if b.lsDnEVuZ == io.EOF {
		b.lsDnEVuZ = nil
	}
	return n, b.readErr()
}

var errNegativeWrite = errors.New("bufio: writer returned negative count from Write")

func (b *NADyavA) im509M7Md(inKjzRUMN3d io.Writer) (int64, error) {
	n, lsDnEVuZ := inKjzRUMN3d.Write(b.bkCmHugyPy[b.rfe2UwaF1FkL:b.inKjzRUMN3d])
	if n < 0 {
		panic(errNegativeWrite)
	}
	b.rfe2UwaF1FkL += n
	return int64(n), lsDnEVuZ
}

// FreeBuf @ 0xe92560，归还共享池并保留底层 Reader。
func (b *NADyavA) FreeBuf() {
	PoolAt50961D8.Put(b.bkCmHugyPy)
	*b = NADyavA{c6HUuPfxb526: b.c6HUuPfxb526, lS7kcz: -1, zyE274D49tb: -1}
}
