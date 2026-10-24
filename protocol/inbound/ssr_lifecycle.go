package inbound

import (
	"context"
	"net"
	G "recovered.local/soga/util/genmap"
	H "recovered.local/soga/util/errutil"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/user"
	Q "recovered.local/soga/panel/registry"
	"runtime"
	"slices"
	"sync"
)

// 原二进制 RTTI：48字节，类型显示前缀 oUi8Xtk3yv，包路径 MFFa9yT。
// 这是原类型和实际行为的源码表示，不是作者原文件。
type Ze0fOaAko struct {
	*JWN154qJ4
	mc0MXUU     *xdWl88
	jTyGfNkAw   *G.GTFAWF[int64, *U.LJJkaD]
	tpQEL_qKogl *G.GTFAWF[int64, string]
	ktET5DFha   *G.GTFAWF[string, *U.LJJkaD]
	g4fBrm0Bxl  bool
}

// 原 RTTI 40 字节。
type xdWl88 struct {
	sync.RWMutex
	spzpWtrEXC1G *Ze0fOaAko
	sSKJ79jFt    map[int]*pxPRmFCpW6SU
}

// 原 RTTI 88 字节。
type pxPRmFCpW6SU struct {
	ot1kuknWf35  context.Context
	iHn9sGij     context.CancelFunc
	yF2TzminVQ1M int
	gHBoVOaQY    *U.LJJkaD
	woqAbWjQH_A  []net.Listener
	aaBpAQ44Zfs  []net.PacketConn
}

// 0x1d7dac0；监听端口在 UpdateAPIData 内同步建立，Start 本来只写日志。
func x4oEZV(arg *DV_XPJ4z0) (Kz_8ajU7, error) {
	base, err := n2s6Gf9ey(arg)
	if err != nil {
		return nil, err
	}
	v := &Ze0fOaAko{JWN154qJ4: base, jTyGfNkAw: &G.GTFAWF[int64, *U.LJJkaD]{}, tpQEL_qKogl: &G.GTFAWF[int64, string]{}, ktET5DFha: &G.GTFAWF[string, *U.LJJkaD]{}, g4fBrm0Bxl: arg.HG5k0365.H47AzZsYHB1N.ASJxBbFS.QawbKS2x}
	v.mc0MXUU = &xdWl88{spzpWtrEXC1G: v, sSKJ79jFt: make(map[int]*pxPRmFCpW6SU)}
	v.UpdateAPIData(arg.HG5k0365)
	Log.TRTbrquD(0, "create shadowsocksr server")
	return v, nil
}
func (v *Ze0fOaAko) Start() error { Log.TRTbrquD(0, "start shadowsocksr server"); return nil }

// 0x1d7ddc0 / Stop.func1 0x1eb8680；保留取消、关全部端口、ResetCtx、GC 顺序。
func (v *Ze0fOaAko) Stop() {
	Log.TRTbrquD(0, "close shadowsocksr server")
	v.yPUbKO()
	v.mc0MXUU.StopAll()
	v.jc_QPN6.Range(func(_ int64, u U.PgaNd3014U) bool { u.ResetCtx(); return true })
	runtime.GC()
}

// 0x1d7dec0。
func (v *Ze0fOaAko) UpdateAPIData(d *Q.D5gTIZ5aVg) bool {
	if v.xBzM2TWz4(d.H47AzZsYHB1N) {
		return true
	}
	v.spcNlUL(d.VTme1SuHL6i_)
	v.afT7skk(d.HYnfYWa)
	if d.BeqZvr6 == nil && !slices.Equal(v.xxf4dimGm, d.ATGYGi0) {
		v.v5VChcYe(d.ATGYGi0)
		v.xxf4dimGm = d.ATGYGi0
	}
	if d.IGP6TVQazVxp == nil && !slices.Equal(v.x1djY10H, d.NjjWbjH) {
		v.ctArZJ(d.NjjWbjH)
		v.x1djY10H = d.NjjWbjH
	}
	return false
}

// 0x1d8ae40；串行收集 Close 错误后返回原命名切片错误，不改用 errors.Join。
func (p *pxPRmFCpW6SU) ojfxJfwLalz3() error {
	if p.iHn9sGij != nil {
		p.iHn9sGij()
	}
	errs := []error{}
	for _, l := range p.woqAbWjQH_A {
		errs = append(errs, l.Close())
	}
	for _, c := range p.aaBpAQ44Zfs {
		errs = append(errs, c.Close())
	}
	return H.T2i1ItMaa(errs...)
}

// 0x1d8c6c0；原函数持读锁关闭端口，但不清空表。
func (m *xdWl88) StopAll() {
	m.RLock()
	defer m.RUnlock()
	for _, p := range m.sSKJ79jFt {
		_ = p.ojfxJfwLalz3()
	}
}
