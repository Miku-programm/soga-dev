package BVM7xtVwiJ

import (
	"io"
	"os"
	"sync"
	"time"
)

// TyIy_s：原 24 字节可重设截止时间，0x18a0e00/0x18a10e0。
type TyIy_s struct {
	pa5spB7    sync.Mutex
	k8uNfDo7r  *time.Timer
	f4UBm3WNmj chan struct{}
}

// NewDeadline 提取原调用处内联的通道初始化，未改变零值行为。
func NewDeadline() TyIy_s { return TyIy_s{f4UBm3WNmj: make(chan struct{})} }

func (d *TyIy_s) Set(t time.Time) {
	d.pa5spB7.Lock()
	defer d.pa5spB7.Unlock()
	if d.k8uNfDo7r != nil {
		if !d.k8uNfDo7r.Stop() {
			<-d.f4UBm3WNmj
		}
	}
	d.k8uNfDo7r = nil
	closed := false
	select {
	case <-d.f4UBm3WNmj:
		closed = true
	default:
	}
	if t.IsZero() {
		if closed {
			d.f4UBm3WNmj = make(chan struct{})
		}
		return
	}
	duration := time.Until(t)
	if duration <= 0 {
		if !closed {
			close(d.f4UBm3WNmj)
		}
		return
	}
	if closed {
		d.f4UBm3WNmj = make(chan struct{})
	}
	d.k8uNfDo7r = time.AfterFunc(duration, func() { close(d.f4UBm3WNmj) })
}
func (d *TyIy_s) Wait() <-chan struct{} {
	d.pa5spB7.Lock()
	defer d.pa5spB7.Unlock()
	return d.f4UBm3WNmj
}

type id7WVlQ struct {
	sync.Mutex
	xVMnwB9ImnWn error
}

func (e *id7WVlQ) Store(err error) { // 0x18a11a0，保留第一次非 nil 错误。
	e.Lock()
	defer e.Unlock()
	if e.xVMnwB9ImnWn == nil {
		e.xVMnwB9ImnWn = err
	}
}
func (e *id7WVlQ) Load() error { // 0x18a12c0
	e.Lock()
	defer e.Unlock()
	return e.xVMnwB9ImnWn
}

type sKIvAo01VX6 struct {
	jUK5p6B      sync.Mutex
	pOkhBeBkJX   chan []byte
	vcvttFgE     chan int
	iHEgQeEGnzXj sync.Once
	k4parX0      chan struct{}
	cAVB2yd1FN7  id7WVlQ
	mB_XlhrDC    id7WVlQ
	aS3h9m       TyIy_s
	mc4LogHcQM   TyIy_s
}
type P848NOz3su9 struct{ sKIvAo01VX6 }
type EjFsCHfka9Zc struct{ i9gO0Lgd P848NOz3su9 }

// NewPipe 提取 vtQf23A9G.OpenStream @0x18a29a0 中原内联构造。
// 读端指向写端对象内部的同一块 144 字节状态。
func NewPipe() (*P848NOz3su9, *EjFsCHfka9Zc) {
	w := &EjFsCHfka9Zc{i9gO0Lgd: P848NOz3su9{sKIvAo01VX6: sKIvAo01VX6{
		pOkhBeBkJX: make(chan []byte), vcvttFgE: make(chan int), k4parX0: make(chan struct{}),
		aS3h9m: NewDeadline(), mc4LogHcQM: NewDeadline(),
	}}}
	return &w.i9gO0Lgd, w
}
func (p *sKIvAo01VX6) pIqVbeo(b []byte) (int, error) { // 0x18a13a0
	select {
	case <-p.k4parX0:
		return 0, p.oxvEkw()
	case <-p.aS3h9m.Wait():
		return 0, os.ErrDeadlineExceeded
	default:
	}
	select {
	case src := <-p.pOkhBeBkJX:
		n := copy(b, src)
		p.vcvttFgE <- n
		return n, nil
	case <-p.k4parX0:
		return 0, p.oxvEkw()
	case <-p.aS3h9m.Wait():
		return 0, os.ErrDeadlineExceeded
	}
}
func (p *sKIvAo01VX6) iOkJGP(b []byte) (n int, err error) { // 0x18a16a0
	select {
	case <-p.k4parX0:
		return 0, p.nQCo0nPvCm6X()
	case <-p.mc4LogHcQM.Wait():
		return 0, os.ErrDeadlineExceeded
	default:
	}
	p.jUK5p6B.Lock()
	defer p.jUK5p6B.Unlock()
	for once := true; once || len(b) > 0; once = false {
		select {
		case p.pOkhBeBkJX <- b:
			used := <-p.vcvttFgE
			b = b[used:]
			n += used
		case <-p.k4parX0:
			return n, p.nQCo0nPvCm6X()
		case <-p.mc4LogHcQM.Wait():
			return n, os.ErrDeadlineExceeded
		}
	}
	return n, nil
}
func (p *sKIvAo01VX6) iDyRGKLKD(err error) error { // 0x18a1600
	if err == nil {
		err = io.ErrClosedPipe
	}
	p.cAVB2yd1FN7.Store(err)
	p.iHEgQeEGnzXj.Do(func() { close(p.k4parX0) })
	return nil
}
func (p *sKIvAo01VX6) rFXK39Yrhg(err error) error { // 0x18a1ac0
	if err == nil {
		err = io.EOF
	}
	p.mB_XlhrDC.Store(err)
	p.iHEgQeEGnzXj.Do(func() { close(p.k4parX0) })
	return nil
}
func (p *sKIvAo01VX6) oxvEkw() error { // 0x18a1b60
	r, w := p.cAVB2yd1FN7.Load(), p.mB_XlhrDC.Load()
	if r == nil && w != nil {
		return w
	}
	return io.ErrClosedPipe
}
func (p *sKIvAo01VX6) nQCo0nPvCm6X() error { // 0x18a1be0
	w, r := p.mB_XlhrDC.Load(), p.cAVB2yd1FN7.Load()
	if w == nil && r != nil {
		return r
	}
	return io.ErrClosedPipe
}
func (r *P848NOz3su9) Read(b []byte) (int, error)      { return r.pIqVbeo(b) }
func (r *P848NOz3su9) Close() error                    { return r.iDyRGKLKD(nil) }
func (r *P848NOz3su9) CloseWithError(err error) error  { return r.iDyRGKLKD(err) }
func (w *EjFsCHfka9Zc) Write(b []byte) (int, error)    { return w.i9gO0Lgd.iOkJGP(b) }
func (w *EjFsCHfka9Zc) Close() error                   { return w.i9gO0Lgd.rFXK39Yrhg(nil) }
func (w *EjFsCHfka9Zc) CloseWithError(err error) error { return w.i9gO0Lgd.rFXK39Yrhg(err) }
func (r *P848NOz3su9) SetReadDeadline(t time.Time) error { // 0x18a1e60
	select {
	case <-r.k4parX0:
		return io.ErrClosedPipe
	default:
	}
	r.aS3h9m.Set(t)
	return nil
}
func (w *EjFsCHfka9Zc) SetWriteDeadline(t time.Time) error { // 0x18a1f20
	select {
	case <-w.i9gO0Lgd.k4parX0:
		return io.ErrClosedPipe
	default:
	}
	w.i9gO0Lgd.mc4LogHcQM.Set(t)
	return nil
}
