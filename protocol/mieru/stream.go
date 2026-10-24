package mieru

import (
	"context"
	"io"
	"net"
	D "recovered.local/soga/buf/pool"
	G "recovered.local/soga/net/relay"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	U "recovered.local/soga/user"
	"sync"
	"sync/atomic"
	"time"
)

// xjTVP8 retains original 184-byte RTTI 0x371a870 and separate residual/owned
// slices. Close signals the stream; Free returns its owned buffers separately.
type xjTVP8 struct {
	xIwb7TaM     uint32
	tduCFDwT     U.PgaNd3014U
	gUoJvESp     context.Context
	xfdnp9xA     chan []byte
	nxYA_EWIa6   []byte
	dGT9ultPgr   []byte
	lnAkdLIFF    net.Conn
	kmcW3r       *qhFcsZC
	wsyay7SHLui  *G.AHU2MRmjzQu
	wP6Ql2PLz    J.ZES6F5
	auBuNh       func(*xjTVP8)
	qYZ2Uu       atomic.Uint32
	cf9qwXK6I5K2 atomic.Int64
	a8D7t3SKMs5P sync.Once
	zhIY4QEke    chan struct{}
}

func vnqzugKQF1jX(id uint32, user U.PgaNd3014U, conn net.Conn, writer *qhFcsZC) *xjTVP8 { //0x1c37f20
	v := &xjTVP8{xIwb7TaM: id, tduCFDwT: user, xfdnp9xA: make(chan []byte, 64), lnAkdLIFF: conn, kmcW3r: writer, zhIY4QEke: make(chan struct{})}
	if user != nil {
		v.gUoJvESp = user.GetCtx()
	}
	return v
}
func (v *xjTVP8) SetDestAddr(a *G.AHU2MRmjzQu) { v.wsyay7SHLui = a }               //0x1c380a0
func (v *xjTVP8) GetDestAddr() *G.AHU2MRmjzQu  { return v.wsyay7SHLui }            //0x1c38e80
func (v *xjTVP8) GetUser() U.PgaNd3014U        { return v.tduCFDwT }               //0x1c38e60
func (v *xjTVP8) NetworkType() J.ZES6F5        { return v.wP6Ql2PLz }              //0x1c38ea0
func (v *xjTVP8) LocalAddr() net.Addr          { return v.lnAkdLIFF.LocalAddr() }  //0x1c38c00
func (v *xjTVP8) RemoteAddr() net.Addr         { return v.lnAkdLIFF.RemoteAddr() } //0x1c38c40
func (v *xjTVP8) SetDeadline(t time.Time) error { //0x1c38c80
	if t.IsZero() {
		v.cf9qwXK6I5K2.Store(0)
	} else {
		v.cf9qwXK6I5K2.Store(t.UnixMicro())
	}
	return nil
}
func (v *xjTVP8) SetReadDeadline(t time.Time) error { return v.SetDeadline(t) } //0x1c38d00
func (v *xjTVP8) SetWriteDeadline(time.Time) error  { return nil }              //0x1c38d80
func (v *xjTVP8) Close() error { //0x1c38aa0 and closure0x1c5c3c0
	v.a8D7t3SKMs5P.Do(func() {
		close(v.zhIY4QEke)
		if v.auBuNh != nil {
			v.auBuNh(v)
		}
	})
	return nil
}
func (v *xjTVP8) Free() { //0x1c38b00
	for {
		select {
		case b := <-v.xfdnp9xA:
			D.PoolAt50961D8.Put(b)
		default:
			goto drained
		}
	}
drained:
	if v.dGT9ultPgr != nil {
		D.PoolAt50961D8.Put(v.dGT9ultPgr)
		v.dGT9ultPgr = nil
	}
	v.nxYA_EWIa6 = nil
}

func (v *xjTVP8) w7MbqIUnAX() error { //0x1c38da0
	if v.tduCFDwT != nil {
		if v.tduCFDwT.IsTrafficOut() {
			return trafficOutAt50983D0
		}
		select {
		case <-v.gUoJvESp.Done():
			return H.ErrorAt5097800
		default:
		}
	}
	return nil
}
func (v *xjTVP8) Read(b []byte) (int, error) { //0x1c38100
	n, err := v.u8WQklo(b)
	if v.tduCFDwT != nil {
		v.tduCFDwT.AddUp(n)
	}
	if check := v.w7MbqIUnAX(); check != nil {
		return 0, check
	}
	return n, err
}
func (v *xjTVP8) Write(b []byte) (int, error) { //0x1c38860
	if v.tduCFDwT != nil {
		v.tduCFDwT.AddDown(len(b))
	}
	if err := v.w7MbqIUnAX(); err != nil {
		return 0, err
	}
	return v.e91JivRNgi(b)
}
func (v *xjTVP8) e91JivRNgi(b []byte) (int, error) { //0x1c38920
	total := len(b)
	for len(b) > 0 {
		seq := v.qYZ2Uu.Add(1) - 1
		m := &IqB1foHWg{Rf6N7kY6GVCs: 7, HNaaDh: v.xIwb7TaM, Tt1OfazA: seq, UlMT6MeaV2: 4096}
		n := min(len(b), 32768)
		if err := v.kmcW3r.WriteDataSegment(m, b[:n]); err != nil {
			return total - len(b), err
		}
		b = b[n:]
	}
	return total, nil
}

// u8WQklo @0x1c381c0. Buffered bytes precede deadline/close checks. A changed
// deadline does not wake an already pending select. Two atomic loads reproduce
// the original seconds/microseconds construction of that pending deadline.
func (v *xjTVP8) u8WQklo(b []byte) (int, error) {
	if len(v.nxYA_EWIa6) > 0 {
		n := copy(b, v.nxYA_EWIa6)
		v.nxYA_EWIa6 = v.nxYA_EWIa6[n:]
		if len(v.nxYA_EWIa6) == 0 {
			v.nxYA_EWIa6 = nil
			D.PoolAt50961D8.Put(v.dGT9ultPgr)
			v.dGT9ultPgr = nil
		}
		return n, nil
	}
	var timeout <-chan time.Time
	if value := v.cf9qwXK6I5K2.Load(); value != 0 {
		until := time.Until(time.Unix(value/1000000, 1000*(v.cf9qwXK6I5K2.Load()%1000000)))
		if until <= 0 {
			return 0, cKThZd6LHk4{}
		}
		timer := time.NewTimer(until)
		defer timer.Stop()
		timeout = timer.C
	}
	select {
	case <-timeout:
		return 0, cKThZd6LHk4{}
	case <-v.zhIY4QEke:
		return 0, io.EOF
	case data, ok := <-v.xfdnp9xA:
		if !ok {
			return 0, io.EOF
		}
		n := copy(b, data)
		if n >= len(data) {
			D.PoolAt50961D8.Put(data)
		} else {
			v.nxYA_EWIa6 = data[n:]
			v.dGT9ultPgr = data
		}
		return n, nil
	}
}
