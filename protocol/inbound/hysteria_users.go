package inbound

import (
	"net"
	Native "recovered.local/soga/protocol/hysteria"
	G "recovered.local/soga/util/genmap"
	L "recovered.local/soga/panel/types"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/user"
	Q "recovered.local/soga/panel/registry"
	"runtime"
	"slices"
)

// Aq5i0Bwez8 的 88 字节结构与原 RTTI 相符；密码索引为原值嵌入 map。
type Aq5i0Bwez8 struct {
	*JWN154qJ4
	uqbfp_xRjia []*Native.P3_qJL5v
	fhPQBadj    G.GTFAWF[string, U.PgaNd3014U]
}

func (v *Aq5i0Bwez8) Authenticate(_ net.Addr, password string, _ uint64) (U.PgaNd3014U, bool) {
	return v.fhPQBadj.Get(password)
} //0x1d56b20

func (v *Aq5i0Bwez8) UpdateAPIData(data *Q.D5gTIZ5aVg) bool { //0x1d57760
	if v.xBzM2TWz4(data.H47AzZsYHB1N) {
		return true
	}
	v.spcNlUL(data.VTme1SuHL6i_)
	v.afT7skk(data.HYnfYWa)
	if data.BeqZvr6 == nil && !slices.Equal(v.xxf4dimGm, data.ATGYGi0) {
		v.v5VChcYe(data.ATGYGi0)
		v.xxf4dimGm = data.ATGYGi0
	}
	if data.IGP6TVQazVxp == nil && !slices.Equal(v.x1djY10H, data.NjjWbjH) {
		v.ctArZJ(data.NjjWbjH)
		v.x1djY10H = data.NjjWbjH
	}
	return false
}

func (v *Aq5i0Bwez8) spcNlUL(panels L.Sxcf3wu3L) { //0x1d57a60
	if panels == nil {
		panels = L.Sxcf3wu3L{}
	}
	seen := make(map[int64]struct{})
	for _, panel := range panels {
		if speed, ok := v.oSsZZT4EmeG.GetUserSpeedLimit(panel.WZ89pTAOrHK4); ok {
			panel.LmOKwjha = speed
		}
		if base, ok := v.jc_QPN6.Get(panel.WZ89pTAOrHK4); ok {
			u := base.(*U.V0wEAoUPFk)
			old := u.GetPassword()
			changed, err := u.UpdatePanelUser(panel)
			if err != nil {
				Log.V8GN5Stk(8, "update user %v err: %v", u.GetId(), err)
				continue
			}
			if changed {
				v.jc_QPN6.Add(u)
				v.fhPQBadj.Delete(old)
				v.fhPQBadj.Set(u.GetPassword(), u)
				Log.TRTbrquD(-4, "changed user:", u)
			}
		} else {
			u, err := U.PB4jVTrQHa(panel, v.dwfzZM)
			if err != nil {
				Log.TRTbrquD(8, err)
				continue
			}
			v.jc_QPN6.Add(u)
			v.fhPQBadj.Set(u.GetPassword(), u)
			Log.TRTbrquD(-4, "add user:", u)
		}
		seen[panel.WZ89pTAOrHK4] = struct{}{}
	}
	v.jc_QPN6.Range(func(id int64, base U.PgaNd3014U) bool { //0x1e6d1e0
		if _, ok := seen[id]; !ok {
			u := base.(*U.V0wEAoUPFk)
			u.ResetCtx()
			v.jc_QPN6.Delete(id)
			v.fhPQBadj.Delete(u.GetPassword())
			Log.TRTbrquD(-4, "remove user:", u)
		}
		return true
	})
	userCountAt50C6DA0 = len(panels)
	Log.V8GN5Stk(0, "current user num: %v", len(panels))
	v.fUxO7UYAH = len(panels)
}

func (v *Aq5i0Bwez8) Stop() { //0x1d57600
	Log.TRTbrquD(0, "close hysteria server")
	v.yPUbKO()
	for _, l := range v.uqbfp_xRjia {
		l.Close()
	}
	v.jc_QPN6.Range(func(_ int64, u U.PgaNd3014U) bool { u.ResetCtx(); return true })
	runtime.GC()
}
