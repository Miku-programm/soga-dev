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
package realityconn

import (
	"encoding/binary"
	"io"
	"net"
	"recovered.local/soga/buf/pool"
	"sync"
	"sync/atomic"
	"time"
)

type YboZfdKYbbwF struct {
	aJW2r4quXUZ0 uint32
	ax9eSX       *DZOCfwWJsy

	dkrfsm       [][]byte
	cCO3_XPTNnnL [][]byte

	jv_tE_9jqFgQ sync.Mutex
	_oSX6fTfG    int

	fQuoCeZnfHdC chan struct{}

	kPA6eSs3XGMS chan struct{}
	odTLapRPCdm  sync.Once

	rcczhz   chan struct{}
	cs8cXKxa sync.Once

	uWxaXR0T atomic.Value
	dDtIx_   atomic.Value

	fTT0Ew3m  uint32
	cGnO8z    uint32
	xwvruuYKD uint32

	f8xi72G1Ov uint32
	cXWhLPc    uint32
	waGBYT     chan struct{}
}

// 此辅助名未保留在元数据中；函数体来自调用者内联指令，非业务占位。
func newStream(aJW2r4quXUZ0 uint32, _oSX6fTfG int, ax9eSX *DZOCfwWJsy) *YboZfdKYbbwF {
	s := new(YboZfdKYbbwF)
	s.aJW2r4quXUZ0 = aJW2r4quXUZ0
	s.fQuoCeZnfHdC = make(chan struct{}, 1)
	s.waGBYT = make(chan struct{}, 1)
	s._oSX6fTfG = _oSX6fTfG
	s.ax9eSX = ax9eSX
	s.kPA6eSs3XGMS = make(chan struct{})
	s.rcczhz = make(chan struct{})
	s.cXWhLPc = initialPeerWindow
	return s
}

// 原函数 0x1196ea0。
func (s *YboZfdKYbbwF) ID() uint32 {
	return s.aJW2r4quXUZ0
}

// 原函数 0x1196ec0。
func (s *YboZfdKYbbwF) Read(b []byte) (n int, err error) {
	for {
		n, err = s.f8IU0XaRuv(b)
		if err == wouldBlockAt5097F70 {
			if ew := s.yLH6hSi(); ew != nil {
				return 0, ew
			}
		} else {
			return n, err
		}
	}
}

// 原函数 0x1196fa0。
func (s *YboZfdKYbbwF) f8IU0XaRuv(b []byte) (n int, err error) {
	if s.ax9eSX.wpUkUgrJJHwN.RTNOROJC5 == 2 {
		return s.aAn_GSQU(b)
	}

	if len(b) == 0 {
		return 0, nil
	}

	var view, recycle []byte
	s.jv_tE_9jqFgQ.Lock()
	if len(s.dkrfsm) > 0 {
		n = min(len(b), len(s.dkrfsm[0]))
		view = s.dkrfsm[0]
		s.dkrfsm[0] = s.dkrfsm[0][n:]
		if len(s.dkrfsm[0]) == 0 {
			s.dkrfsm[0] = nil
			s.dkrfsm = s.dkrfsm[1:]
			recycle = s.cCO3_XPTNnnL[0]
			s.cCO3_XPTNnnL = s.cCO3_XPTNnnL[1:]
		}
	}
	s.jv_tE_9jqFgQ.Unlock()

	copy(b, view)
	if len(recycle) > 0 {
		pool.PoolAt50961D8.Put(recycle)
	}

	if n > 0 {
		s.ax9eSX.returnTokens(n)
		return n, nil
	}

	select {
	case <-s.kPA6eSs3XGMS:
		return 0, io.EOF
	default:
		return 0, wouldBlockAt5097F70
	}
}

// 原函数 0x1197320。
func (s *YboZfdKYbbwF) aAn_GSQU(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}

	var notifyConsumed uint32
	var view, recycle []byte
	s.jv_tE_9jqFgQ.Lock()
	if len(s.dkrfsm) > 0 {
		n = min(len(b), len(s.dkrfsm[0]))
		view = s.dkrfsm[0]
		s.dkrfsm[0] = s.dkrfsm[0][n:]
		if len(s.dkrfsm[0]) == 0 {
			s.dkrfsm[0] = nil
			s.dkrfsm = s.dkrfsm[1:]
			recycle = s.cCO3_XPTNnnL[0]
			s.cCO3_XPTNnnL = s.cCO3_XPTNnnL[1:]
		}
	}
	s.jv_tE_9jqFgQ.Unlock()

	s.fTT0Ew3m += uint32(n)
	s.xwvruuYKD += uint32(n)
	if s.xwvruuYKD >= uint32(s.ax9eSX.wpUkUgrJJHwN.CUuHLYfWnC8q/2) || s.fTT0Ew3m == uint32(n) {
		notifyConsumed = s.fTT0Ew3m
		s.xwvruuYKD = 0
	}

	copy(b, view)
	if len(recycle) > 0 {
		pool.PoolAt50961D8.Put(recycle)
	}

	if n > 0 {
		s.ax9eSX.returnTokens(n)
		if notifyConsumed > 0 {
			err := s.zH509K(notifyConsumed)
			return n, err
		} else {
			return n, nil
		}
	}

	select {
	case <-s.kPA6eSs3XGMS:
		return 0, io.EOF
	default:
		return 0, wouldBlockAt5097F70
	}
}

// 原函数 0x1197720。
func (s *YboZfdKYbbwF) WriteTo(w io.Writer) (n int64, err error) {
	if s.ax9eSX.wpUkUgrJJHwN.RTNOROJC5 == 2 {
		return s.xYDbJhDiF(w)
	}

	for {
		var buf []byte
		s.jv_tE_9jqFgQ.Lock()
		if len(s.dkrfsm) > 0 {
			buf = s.dkrfsm[0]
			s.dkrfsm = s.dkrfsm[1:]
			s.cCO3_XPTNnnL = s.cCO3_XPTNnnL[1:]
		}
		s.jv_tE_9jqFgQ.Unlock()

		if buf != nil {
			nw, ew := w.Write(buf)
			s.ax9eSX.returnTokens(len(buf))
			pool.PoolAt50961D8.Put(buf)
			if nw > 0 {
				n += int64(nw)
			}

			if ew != nil {
				return n, ew
			}
		} else if ew := s.yLH6hSi(); ew != nil {
			return n, ew
		}
	}
}

// 原函数 0x1197a20。
func (s *YboZfdKYbbwF) xYDbJhDiF(w io.Writer) (n int64, err error) {
	for {
		var notifyConsumed uint32
		var buf []byte
		s.jv_tE_9jqFgQ.Lock()
		if len(s.dkrfsm) > 0 {
			buf = s.dkrfsm[0]
			s.dkrfsm = s.dkrfsm[1:]
			s.cCO3_XPTNnnL = s.cCO3_XPTNnnL[1:]
		}
		s.fTT0Ew3m += uint32(len(buf))
		s.xwvruuYKD += uint32(len(buf))
		if s.xwvruuYKD >= uint32(s.ax9eSX.wpUkUgrJJHwN.CUuHLYfWnC8q/2) || s.fTT0Ew3m == uint32(len(buf)) {
			notifyConsumed = s.fTT0Ew3m
			s.xwvruuYKD = 0
		}
		s.jv_tE_9jqFgQ.Unlock()

		if buf != nil {
			nw, ew := w.Write(buf)
			s.ax9eSX.returnTokens(len(buf))
			pool.PoolAt50961D8.Put(buf)
			if nw > 0 {
				n += int64(nw)
			}

			if ew != nil {
				return n, ew
			}

			if notifyConsumed > 0 {
				if err := s.zH509K(notifyConsumed); err != nil {
					return n, err
				}
			}
		} else if ew := s.yLH6hSi(); ew != nil {
			return n, ew
		}
	}
}

// 原函数 0x1197da0。
func (s *YboZfdKYbbwF) zH509K(consumed uint32) error {
	var timer *time.Timer
	var deadline <-chan time.Time
	if d, ok := s.uWxaXR0T.Load().(time.Time); ok && !d.IsZero() {
		timer = time.NewTimer(time.Until(d))
		defer timer.Stop()
		deadline = timer.C
	}

	fN47GQEn := newFrame(byte(s.ax9eSX.wpUkUgrJJHwN.RTNOROJC5), cmdUPD, s.aJW2r4quXUZ0)
	var hdr xnabvi1sOhR
	binary.LittleEndian.PutUint32(hdr[:], consumed)
	binary.LittleEndian.PutUint32(hdr[4:], uint32(s.ax9eSX.wpUkUgrJJHwN.CUuHLYfWnC8q))
	fN47GQEn.cAKGyUaj = hdr[:]
	err := s.ax9eSX.dI7L3D(fN47GQEn, deadline, CLSDATA)
	return err
}

// 原函数 0x1197fa0。
func (s *YboZfdKYbbwF) yLH6hSi() error {
	var timer *time.Timer
	var deadline <-chan time.Time
	if d, ok := s.uWxaXR0T.Load().(time.Time); ok && !d.IsZero() {
		timer = time.NewTimer(time.Until(d))
		defer timer.Stop()
		deadline = timer.C
	}

	select {
	case <-s.fQuoCeZnfHdC:
		return nil
	case <-s.rcczhz:
		s.jv_tE_9jqFgQ.Lock()
		remaining := len(s.dkrfsm)
		s.jv_tE_9jqFgQ.Unlock()
		if remaining > 0 {
			return nil
		}
		return io.EOF
	case <-s.ax9eSX.dq5F2u2:
		return s.ax9eSX.l4IUcN.Load().(error)
	case <-s.ax9eSX.eLUgKTG:
		return s.ax9eSX.u1sgxOsIJxL.Load().(error)
	case <-deadline:
		return timeoutAt502B390
	case <-s.kPA6eSs3XGMS:
		return io.ErrClosedPipe
	}

}

// 原函数 0x1198560。
func (s *YboZfdKYbbwF) Write(b []byte) (n int, err error) {
	if s.ax9eSX.wpUkUgrJJHwN.RTNOROJC5 == 2 {
		return s.kcdbgkEayd(b)
	}

	var deadline <-chan time.Time
	if d, ok := s.dDtIx_.Load().(time.Time); ok && !d.IsZero() {
		timer := time.NewTimer(time.Until(d))
		defer timer.Stop()
		deadline = timer.C
	}

	select {
	case <-s.kPA6eSs3XGMS:
		return 0, io.ErrClosedPipe
	default:
	}

	sent := 0
	fN47GQEn := newFrame(byte(s.ax9eSX.wpUkUgrJJHwN.RTNOROJC5), cmdPSH, s.aJW2r4quXUZ0)
	bts := b
	for len(bts) > 0 {
		sz := len(bts)
		if sz > s._oSX6fTfG {
			sz = s._oSX6fTfG
		}
		fN47GQEn.cAKGyUaj = bts[:sz]
		n, err := s.ax9eSX.l1G8yaCPEm(fN47GQEn, deadline, CLSDATA)
		s.cGnO8z++
		sent += n
		if err != nil {
			return sent, err
		}
		bts = bts[sz:]
	}

	return sent, nil
}

// 原函数 0x1198960。
func (s *YboZfdKYbbwF) kcdbgkEayd(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}

	select {
	case <-s.kPA6eSs3XGMS:
		return 0, io.ErrClosedPipe
	default:
	}

	var deadline <-chan time.Time
	if d, ok := s.dDtIx_.Load().(time.Time); ok && !d.IsZero() {
		timer := time.NewTimer(time.Until(d))
		defer timer.Stop()
		deadline = timer.C
	}

	sent := 0
	fN47GQEn := newFrame(byte(s.ax9eSX.wpUkUgrJJHwN.RTNOROJC5), cmdPSH, s.aJW2r4quXUZ0)

	for {
		var bts []byte
		inflight := int32(atomic.LoadUint32(&s.cGnO8z) - atomic.LoadUint32(&s.f8xi72G1Ov))
		if inflight < 0 {
			return 0, consumedAt5097F50
		}

		win := int32(atomic.LoadUint32(&s.cXWhLPc)) - inflight
		if win > 0 {
			if win > int32(len(b)) {
				bts = b
				b = nil
			} else {
				bts = b[:win]
				b = b[win:]
			}

			for len(bts) > 0 {
				sz := len(bts)
				if sz > s._oSX6fTfG {
					sz = s._oSX6fTfG
				}
				fN47GQEn.cAKGyUaj = bts[:sz]
				bts = bts[sz:]
				n, err := s.ax9eSX.l1G8yaCPEm(fN47GQEn, deadline, CLSDATA)
				atomic.AddUint32(&s.cGnO8z, uint32(sz))
				sent += n
				if err != nil {
					return sent, err
				}
			}
		}

		if len(b) > 0 {
			select {
			case <-s.rcczhz:
				return 0, io.EOF
			case <-s.kPA6eSs3XGMS:
				return sent, io.ErrClosedPipe
			case <-deadline:
				return sent, timeoutAt502B390
			case <-s.ax9eSX.nvN2Bu:
				return sent, s.ax9eSX.xnMhVM5Fh.Load().(error)
			case <-s.waGBYT:
				continue
			}
		} else {
			return sent, nil
		}
	}
}

// 原函数 0x11991e0。
func (s *YboZfdKYbbwF) Close() error {
	var once bool
	var err error
	s.odTLapRPCdm.Do(func() {
		close(s.kPA6eSs3XGMS)
		once = true
	})

	if once {
		_, err = s.ax9eSX.ei19Cr(newFrame(byte(s.ax9eSX.wpUkUgrJJHwN.RTNOROJC5), cmdFIN, s.aJW2r4quXUZ0))
		s.ax9eSX.gHio09O(s.aJW2r4quXUZ0)
		return err
	} else {
		return io.ErrClosedPipe
	}
}

// 原函数 0x11992c0。
func (s *YboZfdKYbbwF) GetDieCh() <-chan struct{} {
	return s.kPA6eSs3XGMS
}

// 原函数 0x11992e0。
func (s *YboZfdKYbbwF) SetReadDeadline(t time.Time) error {
	s.uWxaXR0T.Store(t)
	s.notifyReadEvent()
	return nil
}

// 原函数 0x11993a0。
func (s *YboZfdKYbbwF) SetWriteDeadline(t time.Time) error {
	s.dDtIx_.Store(t)
	return nil
}

// 原函数 0x1199440。
func (s *YboZfdKYbbwF) SetDeadline(t time.Time) error {
	if err := s.SetReadDeadline(t); err != nil {
		return err
	}
	if err := s.SetWriteDeadline(t); err != nil {
		return err
	}
	return nil
}

// 原函数 0x1199560。
func (s *YboZfdKYbbwF) kG4hgG() { s.odTLapRPCdm.Do(func() { close(s.kPA6eSs3XGMS) }) }

// 原函数 0x11995c0。
func (s *YboZfdKYbbwF) LocalAddr() net.Addr {
	if ts, ok := s.ax9eSX.mlu21WY6.(interface {
		LocalAddr() net.Addr
	}); ok {
		return ts.LocalAddr()
	}
	return nil
}

// 原函数 0x1199680。
func (s *YboZfdKYbbwF) RemoteAddr() net.Addr {
	if ts, ok := s.ax9eSX.mlu21WY6.(interface {
		RemoteAddr() net.Addr
	}); ok {
		return ts.RemoteAddr()
	}
	return nil
}

// 原函数 0x1199740。
func (s *YboZfdKYbbwF) nF0n3Xsat(buf []byte) (written int, err error) {
	s.jv_tE_9jqFgQ.Lock()
	s.dkrfsm = append(s.dkrfsm, buf)
	s.cCO3_XPTNnnL = append(s.cCO3_XPTNnnL, buf)
	s.jv_tE_9jqFgQ.Unlock()
	return
}

// 原函数 0x1199940。
func (s *YboZfdKYbbwF) hgJudzwz() (n int) {
	s.jv_tE_9jqFgQ.Lock()
	dkrfsm, cCO3_XPTNnnL := s.dkrfsm, s.cCO3_XPTNnnL
	s.dkrfsm = nil
	s.cCO3_XPTNnnL = nil
	s.jv_tE_9jqFgQ.Unlock()
	for k := range dkrfsm {
		n += len(dkrfsm[k])
		pool.PoolAt50961D8.Put(cCO3_XPTNnnL[k])
	}
	return
}

// 此辅助名未保留在元数据中；函数体来自调用者内联指令，非业务占位。
func (s *YboZfdKYbbwF) notifyReadEvent() {
	select {
	case s.fQuoCeZnfHdC <- struct{}{}:
	default:
	}
}

// 此辅助名未保留在元数据中；函数体来自调用者内联指令，非业务占位。
func (s *YboZfdKYbbwF) update(consumed uint32, window uint32) {
	atomic.StoreUint32(&s.f8xi72G1Ov, consumed)
	atomic.StoreUint32(&s.cXWhLPc, window)
	select {
	case s.waGBYT <- struct{}{}:
	default:
	}
}

// 原函数 0x1199aa0。
func (s *YboZfdKYbbwF) qaoU8ZUm_V() {
	s.cs8cXKxa.Do(func() {
		close(s.rcczhz)
	})
}
