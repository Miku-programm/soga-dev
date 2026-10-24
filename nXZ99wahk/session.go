package nXZ99wahk

import (
	"context"
	"errors"
	"fmt"
	"net"
	"recovered.local/soga/BaEWsPy"
	"recovered.local/soga/CpHCdepV"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/qcqMSUWwIx40"
)

// Ps5tpGhQ 保留原80字节布局与字段；源码由机器码/RTTI重建。
type Ps5tpGhQ struct {
	aojobe      context.Context
	bYFyxREdBX  context.CancelFunc
	mECvR4TiwY  net.Conn
	gIPtn7Z3pn_ *BaEWsPy.D23y2888a_
	jtfRLia     nzATvajt.PgaNd3014U
	zbRz7R      *XtJioO
	i1xUuKWJ    chan J1P4KNvfI_J.Ki8urSHTx44
}

// KJwDQE2y3 对应0x1ced4c0；串行化底层写入后启动真实帧接收循环。
func KJwDQE2y3(ctx context.Context, c J1P4KNvfI_J.Ki8urSHTx44, r *BaEWsPy.D23y2888a_) *Ps5tpGhQ {
	ctx, cancel := context.WithCancel(ctx)
	p := &Ps5tpGhQ{aojobe: ctx, bYFyxREdBX: cancel, mECvR4TiwY: &CpHCdepV.MhXtuhyPkR5{DrpXao: c}, gIPtn7Z3pn_: r, jtfRLia: c.GetUser(), zbRz7R: &XtJioO{}, i1xUuKWJ: make(chan J1P4KNvfI_J.Ki8urSHTx44, 8)}
	go p.bjXNbL()
	return p
}

// bjXNbL 对应0x1ced780及清理闭包0x1cf2180；原清理先关子流，再取消会话并关连接。
func (p *Ps5tpGhQ) bjXNbL() {
	defer func() { p.zbRz7R.Clear(); p.Close() }()
	for {
		select {
		case <-p.aojobe.Done():
			return
		default:
		}
		f, err := HlVrAnC111(p.mECvR4TiwY)
		if err != nil {
			muc279Mpa6S4.JucPcHhJ("mux session recv loop read request: %v", err)
			return
		}
		switch f.xvTw9VQDQdY {
		case 1:
			err = p.caYqS8sJ(f)
		case 2:
			err = p.l4vHJHVTLmEy(f)
		case 3:
			err = p.ie6rc5Jas(f)
		case 4:
			f.FreeData()
		}
		if err != nil {
			muc279Mpa6S4.JucPcHhJ("mux session recv loop handle cmd <%v>: %v", f.xvTw9VQDQdY, err)
			return
		}
	}
}

// caYqS8sJ 对应0x1cedac0。原逻辑先送入 Accept 通道，再登记映射；保留此顺序。
func (p *Ps5tpGhQ) caYqS8sJ(f *CUwik6VP2WU0) error {
	if _, ok := p.zbRz7R.f56bH6.Get(f.ljtkJw); ok {
		f.FreeData()
		return fmt.Errorf("repeat mux id: %d", f.ljtkJw)
	}
	muc279Mpa6S4.JucPcHhJ("mux new: %v", f.ljtkJw)
	if f.ePmcTrBC != 1 && f.ePmcTrBC != 2 {
		f.FreeData()
		return fmt.Errorf("unknown mux net type: %v", f.ePmcTrBC)
	}
	c := &Ldrr25M{gwE4Or1X8JW: f.ljtkJw, fj8yHOR: make(chan *CUwik6VP2WU0, 2), hiMx_c: p.mECvR4TiwY, iBwiz6Mebe: p.jtfRLia, yFmjxzi: f.hUHpeGfV, altuteX: f.c8AtRj7hb, s1w1RA: qcqMSUWwIx40.NewDeadline()}
	if f.ePmcTrBC == 1 {
		p.i1xUuKWJ <- c
		p.zbRz7R.f56bH6.Set(f.ljtkJw, c)
	} else {
		udp := &WSnkIcJKU{Ldrr25M: c, oNZFwhbp85ce: p.gIPtn7Z3pn_}
		c.altuteX = nil
		if f.HasData() {
			c.fj8yHOR <- f
		}
		p.i1xUuKWJ <- udp
		p.zbRz7R.f56bH6.Set(f.ljtkJw, udp)
	}
	return nil
}
func (p *Ps5tpGhQ) l4vHJHVTLmEy(f *CUwik6VP2WU0) error { //0x1cee2e0
	if c, ok := p.zbRz7R.f56bH6.Get(f.ljtkJw); ok {
		if f.HasData() {
			c.Feed(f)
		}
		return nil
	}
	f.FreeData()
	muc279Mpa6S4.JucPcHhJ("mux keep: id %v not found", f.ljtkJw)
	return nil
}
func (p *Ps5tpGhQ) ie6rc5Jas(f *CUwik6VP2WU0) error { //0x1cee440
	muc279Mpa6S4.JucPcHhJ("mux end: %v", f.ljtkJw)
	if c, ok := p.zbRz7R.f56bH6.GetAndDelete(f.ljtkJw); ok {
		c.Close()
	}
	f.FreeData()
	if f.HasError() {
		return errors.New("session has error")
	}
	return nil
}
func (p *Ps5tpGhQ) Accept() (J1P4KNvfI_J.Ki8urSHTx44, bool) { //0x1cee620
	select {
	case <-p.aojobe.Done():
		return nil, false
	case c, ok := <-p.i1xUuKWJ:
		return c, ok
	}
}
func (p *Ps5tpGhQ) Close() { p.bYFyxREdBX(); p.mECvR4TiwY.Close() } //0x1cee6e0
