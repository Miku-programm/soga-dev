/*
MIT License

Copyright (c) 2016-2017 Daniel Fu

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

// 依据原 ELF（SHA256 6660dcfb…）IDA/RTTI 恢复；并非作者原文件。
// 以 MIT 许可 smux v1.5.24 为可比对基线，原 fork 的批量发送、锁范围、
// 版本协商、异步窗口更新及缓冲池差异均按机器码恢复，不能用上游包直接替代。
// RTTI 显示前缀 ji2LadBZQZW6，uncommon package path 为 I4AWsAkBLRP。
package I4AWsAkBLRP

import (
	"errors"
	"io"
	"net"
	"recovered.local/soga/D_zCag1qTV"
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultAcceptBacklog = 1024
	openCloseTimeout     = 30 * time.Second
)

type JZpBeJdC int

const (
	CLSCTRL JZpBeJdC = iota
	CLSDATA
)

var (
	invalidProtocolAt5097F40 = errors.New("invalid protocol")
	consumedAt5097F50        = errors.New("peer consumed more than sent")
	goAwayAt5097F60          = errors.New("stream id overflows, should start a new connection")
	timeoutAt502B390         = errors.New("timeout")
	wouldBlockAt5097F70      = errors.New("operation would block on IO")
)

type vwv8dqM struct {
	e9dK1WBuMa   JZpBeJdC
	fN47GQEn     TKUtOq
	vlOxyuBq8ThB uint32
	wbTHvwrgbD   chan yoowfr
}

type yoowfr struct {
	b1C__L1 int
	uRSOFlo error
}

type DZOCfwWJsy struct {
	mlu21WY6 io.ReadWriteCloser

	wpUkUgrJJHwN *BWpFHH3x
	mTqgfhNdlT82 uint32
	p2tD2AhCZnC  sync.Mutex

	zbny3Aedj int32
	tnmOyFJvf chan struct{}

	wjubqp map[uint32]*YboZfdKYbbwF
	r6klva sync.RWMutex

	nHOcCGRza7V chan struct{}
	h9RMpMidR   sync.Once

	l4IUcN       atomic.Value
	xnMhVM5Fh    atomic.Value
	dq5F2u2      chan struct{}
	nvN2Bu       chan struct{}
	_5AKaibLIdIF sync.Once
	gFei_ksGsjwr sync.Once

	u1sgxOsIJxL atomic.Value
	eLUgKTG     chan struct{}
	eR6WSnZ4i   sync.Once

	cf0Ovj chan *YboZfdKYbbwF

	nZwZauYdh int32

	jE8dYZ8F_ int32

	a2A6Mhwyca atomic.Value

	bVqjxgmOu  uint32
	hHxKb51kQa chan vwv8dqM
}

// 原函数 0x1192c60。
func yfgupp(wpUkUgrJJHwN *BWpFHH3x, mlu21WY6 io.ReadWriteCloser, client bool) *DZOCfwWJsy {
	s := new(DZOCfwWJsy)
	s.nHOcCGRza7V = make(chan struct{})
	s.mlu21WY6 = mlu21WY6
	s.wpUkUgrJJHwN = wpUkUgrJJHwN
	s.wjubqp = make(map[uint32]*YboZfdKYbbwF)
	s.cf0Ovj = make(chan *YboZfdKYbbwF, defaultAcceptBacklog)
	s.zbny3Aedj = int32(wpUkUgrJJHwN.K3nkPQuOI)
	s.tnmOyFJvf = make(chan struct{}, 1)
	s.hHxKb51kQa = make(chan vwv8dqM, 16)
	s.dq5F2u2 = make(chan struct{})
	s.nvN2Bu = make(chan struct{})
	s.eLUgKTG = make(chan struct{})

	if client {
		s.mTqgfhNdlT82 = 1
	} else {
		s.mTqgfhNdlT82 = 0
	}

	go s.fg7YBsqh()
	go s.hm9Q3Wo()
	if !wpUkUgrJJHwN.DU0iTVAg {
		go s.fgZVNR()
	}
	return s
}

// 原函数 0x1193000。
func (s *DZOCfwWJsy) OpenStream() (*YboZfdKYbbwF, error) {
	if s.IsClosed() {
		return nil, io.ErrClosedPipe
	}

	s.p2tD2AhCZnC.Lock()
	if s.jE8dYZ8F_ > 0 {
		s.p2tD2AhCZnC.Unlock()
		return nil, goAwayAt5097F60
	}

	s.mTqgfhNdlT82 += 2
	a_9mrWgHy := s.mTqgfhNdlT82
	if a_9mrWgHy == a_9mrWgHy%2 {
		s.jE8dYZ8F_ = 1
		s.p2tD2AhCZnC.Unlock()
		return nil, goAwayAt5097F60
	}
	s.p2tD2AhCZnC.Unlock()

	stream := newStream(a_9mrWgHy, s.wpUkUgrJJHwN.XsqVEjw4, s)

	if _, err := s.ei19Cr(newFrame(byte(s.wpUkUgrJJHwN.RTNOROJC5), cmdSYN, a_9mrWgHy)); err != nil {
		return nil, err
	}

	select {
	case <-s.dq5F2u2:
		return nil, s.l4IUcN.Load().(error)
	case <-s.nvN2Bu:
		return nil, s.xnMhVM5Fh.Load().(error)
	case <-s.nHOcCGRza7V:
		return nil, io.ErrClosedPipe
	default:
		s.r6klva.Lock()
		s.wjubqp[a_9mrWgHy] = stream
		s.r6klva.Unlock()
		return stream, nil
	}
}

// 原函数 0x1193560。
func (s *DZOCfwWJsy) Open() (io.ReadWriteCloser, error) {
	return s.OpenStream()
}

// 原函数 0x11935a0。
func (s *DZOCfwWJsy) AcceptStream() (*YboZfdKYbbwF, error) {
	var a2A6Mhwyca <-chan time.Time
	if d, ok := s.a2A6Mhwyca.Load().(time.Time); ok && !d.IsZero() {
		timer := time.NewTimer(time.Until(d))
		defer timer.Stop()
		a2A6Mhwyca = timer.C
	}

	select {
	case stream := <-s.cf0Ovj:
		return stream, nil
	case <-a2A6Mhwyca:
		return nil, timeoutAt502B390
	case <-s.dq5F2u2:
		return nil, s.l4IUcN.Load().(error)
	case <-s.eLUgKTG:
		return nil, s.u1sgxOsIJxL.Load().(error)
	case <-s.nHOcCGRza7V:
		return nil, io.ErrClosedPipe
	}
}

// 原函数 0x1193aa0。
func (s *DZOCfwWJsy) Accept() (io.ReadWriteCloser, error) {
	return s.AcceptStream()
}

// 原函数 0x1193ae0。
func (s *DZOCfwWJsy) Close() error {
	var once bool
	s.h9RMpMidR.Do(func() {
		close(s.nHOcCGRza7V)
		once = true
	})

	if once {
		s.r6klva.RLock()
		for k := range s.wjubqp {
			s.wjubqp[k].kG4hgG()
		}
		s.r6klva.RUnlock()
		return s.mlu21WY6.Close()
	} else {
		return io.ErrClosedPipe
	}
}

// 原函数 0x1193c60。
func (s *DZOCfwWJsy) CloseChan() <-chan struct{} {
	return s.nHOcCGRza7V
}

// 此辅助名未保留在元数据中；函数体来自调用者内联指令，非业务占位。
func (s *DZOCfwWJsy) notifyBucket() {
	select {
	case s.tnmOyFJvf <- struct{}{}:
	default:
	}
}

// 原函数 0x1193c80。
func (s *DZOCfwWJsy) pIBdZC0s(err error) {
	s._5AKaibLIdIF.Do(func() {
		s.l4IUcN.Store(err)
		close(s.dq5F2u2)
	})
}

// 原函数 0x1193d00。
func (s *DZOCfwWJsy) dp0GL_8V7wP(err error) {
	s.gFei_ksGsjwr.Do(func() {
		s.xnMhVM5Fh.Store(err)
		close(s.nvN2Bu)
	})
}

// 原函数 0x1193d80。
func (s *DZOCfwWJsy) aAQsNYZ12B(err error) {
	s.eR6WSnZ4i.Do(func() {
		s.u1sgxOsIJxL.Store(err)
		close(s.eLUgKTG)
	})
}

// 原函数 0x1193e00。
func (s *DZOCfwWJsy) IsClosed() bool {
	select {
	case <-s.nHOcCGRza7V:
		return true
	default:
		return false
	}
}

// 原函数 0x1193e60。
func (s *DZOCfwWJsy) NumStreams() int {
	if s.IsClosed() {
		return 0
	}
	s.r6klva.RLock()
	defer s.r6klva.RUnlock()
	return len(s.wjubqp)
}

// 原函数 0x1193f60。
func (s *DZOCfwWJsy) SetDeadline(t time.Time) error {
	s.a2A6Mhwyca.Store(t)
	return nil
}

// 原函数 0x1194000。
func (s *DZOCfwWJsy) LocalAddr() net.Addr {
	if ts, ok := s.mlu21WY6.(interface {
		LocalAddr() net.Addr
	}); ok {
		return ts.LocalAddr()
	}
	return nil
}

// 原函数 0x11940c0。
func (s *DZOCfwWJsy) RemoteAddr() net.Addr {
	if ts, ok := s.mlu21WY6.(interface {
		RemoteAddr() net.Addr
	}); ok {
		return ts.RemoteAddr()
	}
	return nil
}

// 原函数 0x1194180。
func (s *DZOCfwWJsy) gHio09O(a_9mrWgHy uint32) {
	s.r6klva.Lock()
	if n := s.wjubqp[a_9mrWgHy].hgJudzwz(); n > 0 {
		if atomic.AddInt32(&s.zbny3Aedj, int32(n)) > 0 {
			s.notifyBucket()
		}
	}
	delete(s.wjubqp, a_9mrWgHy)
	s.r6klva.Unlock()
}

// 此辅助名未保留在元数据中；函数体来自调用者内联指令，非业务占位。
func (s *DZOCfwWJsy) returnTokens(n int) {
	if atomic.AddInt32(&s.zbny3Aedj, int32(n)) > 0 {
		s.notifyBucket()
	}
}

// 原函数 0x1194260。
func (s *DZOCfwWJsy) fg7YBsqh() {
	var hdr rawHeader
	var updHdr xnabvi1sOhR
	for {
		for atomic.LoadInt32(&s.zbny3Aedj) <= 0 && !s.IsClosed() {
			select {
			case <-s.tnmOyFJvf:
			case <-s.nHOcCGRza7V:
				return
			}
		}
		if _, err := io.ReadFull(s.mlu21WY6, hdr[:]); err != nil {
			s.pIBdZC0s(err)
			return
		}
		atomic.StoreInt32(&s.nZwZauYdh, 1)
		if hdr.RTNOROJC5() == 0 || hdr.RTNOROJC5() > 2 {
			s.aAQsNYZ12B(invalidProtocolAt5097F40)
			return
		}
		s.wpUkUgrJJHwN.RTNOROJC5 = int(hdr.RTNOROJC5())
		a_9mrWgHy := hdr.StreamID()
		switch hdr.Cmd() {
		case cmdNOP:
		case cmdSYN:
			s.r6klva.Lock()
			if _, ok := s.wjubqp[a_9mrWgHy]; !ok {
				stream := newStream(a_9mrWgHy, s.wpUkUgrJJHwN.XsqVEjw4, s)
				s.wjubqp[a_9mrWgHy] = stream
				s.r6klva.Unlock()
				select {
				case s.cf0Ovj <- stream:
				case <-s.nHOcCGRza7V:
				}
			} else {
				s.r6klva.Unlock()
			}
		case cmdFIN:
			s.r6klva.RLock()
			stream, ok := s.wjubqp[a_9mrWgHy]
			s.r6klva.RUnlock()
			if ok {
				stream.qaoU8ZUm_V()
				stream.notifyReadEvent()
			}
		case cmdPSH:
			if hdr.Length() > 0 {
				newbuf := D_zCag1qTV.PoolAt50961D8.Get(int(hdr.Length()))
				written, err := io.ReadFull(s.mlu21WY6, newbuf)
				if err != nil {
					D_zCag1qTV.PoolAt50961D8.Put(newbuf)
					s.pIBdZC0s(err)
					return
				}
				s.r6klva.RLock()
				stream, ok := s.wjubqp[a_9mrWgHy]
				s.r6klva.RUnlock()
				if ok {
					stream.nF0n3Xsat(newbuf)
					atomic.AddInt32(&s.zbny3Aedj, -int32(written))
					stream.notifyReadEvent()
				} else {
					D_zCag1qTV.PoolAt50961D8.Put(newbuf)
				}
			}
		case cmdUPD:
			if _, err := io.ReadFull(s.mlu21WY6, updHdr[:]); err != nil {
				s.pIBdZC0s(err)
				return
			}
			s.r6klva.RLock()
			stream, ok := s.wjubqp[a_9mrWgHy]
			s.r6klva.RUnlock()
			if ok {
				stream.update(updHdr.Consumed(), updHdr.Window())
			}
		default:
			s.aAQsNYZ12B(invalidProtocolAt5097F40)
			return
		}
	}
}

// 原函数 0x1194ca0。
func (s *DZOCfwWJsy) fgZVNR() {
	tickerPing := time.NewTicker(s.wpUkUgrJJHwN.PGgIC2)
	tickerTimeout := time.NewTicker(s.wpUkUgrJJHwN.PK9EpD)
	defer tickerPing.Stop()
	defer tickerTimeout.Stop()
	for {
		select {
		case <-tickerPing.C:
			s.l1G8yaCPEm(newFrame(byte(s.wpUkUgrJJHwN.RTNOROJC5), cmdNOP, 0), tickerPing.C, CLSCTRL)
			s.notifyBucket()
		case <-tickerTimeout.C:
			if !atomic.CompareAndSwapInt32(&s.nZwZauYdh, 1, 0) {
				if atomic.LoadInt32(&s.zbny3Aedj) > 0 {
					s.Close()
					return
				}
			}
		case <-s.nHOcCGRza7V:
			return
		}
	}
}

// 原函数 0x1195880。
func (s *DZOCfwWJsy) ei19Cr(f TKUtOq) (n int, err error) {
	timer := time.NewTimer(openCloseTimeout)
	defer timer.Stop()
	return s.l1G8yaCPEm(f, timer.C, CLSCTRL)
}

// 原函数 0x1195d40。
func (s *DZOCfwWJsy) l1G8yaCPEm(f TKUtOq, a2A6Mhwyca <-chan time.Time, e9dK1WBuMa JZpBeJdC) (int, error) {
	req := vwv8dqM{
		e9dK1WBuMa:   e9dK1WBuMa,
		fN47GQEn:     f,
		vlOxyuBq8ThB: atomic.AddUint32(&s.bVqjxgmOu, 1),
		wbTHvwrgbD:   make(chan yoowfr, 1),
	}
	select {
	case s.hHxKb51kQa <- req:
	case <-s.nHOcCGRza7V:
		return 0, io.ErrClosedPipe
	case <-s.nvN2Bu:
		return 0, s.xnMhVM5Fh.Load().(error)
	case <-a2A6Mhwyca:
		return 0, timeoutAt502B390
	}

	select {
	case wbTHvwrgbD := <-req.wbTHvwrgbD:
		return wbTHvwrgbD.b1C__L1, wbTHvwrgbD.uRSOFlo
	case <-s.nHOcCGRza7V:
		return 0, io.ErrClosedPipe
	case <-s.nvN2Bu:
		return 0, s.xnMhVM5Fh.Load().(error)
	case <-a2A6Mhwyca:
		return 0, timeoutAt502B390
	}
}

// 原函数 0x1194ee0。
func (s *DZOCfwWJsy) hm9Q3Wo() {
	buf := D_zCag1qTV.S4OQXWIZ(20480)
	defer buf.Free()
	requests := make([]vwv8dqM, 0, 16)
	for {
		requests = requests[:0]
	collect:
		for len(requests) < 16 {
			select {
			case request := <-s.hHxKb51kQa:
				requests = append(requests, request)
			case <-s.nHOcCGRza7V:
				if len(requests) == 0 {
					return
				}
				break collect
			}
			if len(s.hHxKb51kQa) == 0 {
				break
			}
		}
		buf.Reset()
		start := 0
		var err error
		for i, request := range requests {
			if !buf.IsEmpty() && buf.Len()+len(request.fN47GQEn.cAKGyUaj)+8 > 20480 {
				_, err = s.mlu21WY6.Write(buf.Bytes())
				for j := start; j < i; j++ {
					n := len(requests[j].fN47GQEn.cAKGyUaj)
					if err != nil {
						n = 0
					}
					requests[j].wbTHvwrgbD <- yoowfr{n, err}
					close(requests[j].wbTHvwrgbD)
					requests[j] = vwv8dqM{}
				}
				if err != nil {
					break
				}
				buf.Reset()
				start = i
			}
			buf.WriteByte(request.fN47GQEn.xlaU77zvaW7W)
			buf.WriteByte(request.fN47GQEn.gUa4uFKR6)
			buf.WriteUint16Le(uint16(len(request.fN47GQEn.cAKGyUaj)))
			buf.WriteUint32Le(request.fN47GQEn.a_9mrWgHy)
			buf.Write(request.fN47GQEn.cAKGyUaj)
		}
		if err == nil && start < len(requests) {
			_, err = s.mlu21WY6.Write(buf.Bytes())
			for j := start; j < len(requests); j++ {
				n := len(requests[j].fN47GQEn.cAKGyUaj)
				if err != nil {
					n = 0
				}
				requests[j].wbTHvwrgbD <- yoowfr{n, err}
				close(requests[j].wbTHvwrgbD)
				requests[j] = vwv8dqM{}
			}
		}
		if err != nil {
			s.dp0GL_8V7wP(err)
			return
		}
	}
}

// 原函数 0x1195a00。
func (s *DZOCfwWJsy) dI7L3D(f TKUtOq, a2A6Mhwyca <-chan time.Time, e9dK1WBuMa JZpBeJdC) error {
	req := vwv8dqM{e9dK1WBuMa: e9dK1WBuMa, fN47GQEn: f, vlOxyuBq8ThB: atomic.AddUint32(&s.bVqjxgmOu, 1), wbTHvwrgbD: make(chan yoowfr, 1)}
	select {
	case s.hHxKb51kQa <- req:
		return nil
	case <-s.nHOcCGRza7V:
		return io.ErrClosedPipe
	case <-s.nvN2Bu:
		return s.xnMhVM5Fh.Load().(error)
	case <-a2A6Mhwyca:
		return timeoutAt502B390
	}
}
