package MFFa9yT

import (
	"net"
	J "recovered.local/soga/J1P4KNvfI_J"
	L "recovered.local/soga/LFWTTXF7JR"
	N "recovered.local/soga/NtxdAE"
	C "recovered.local/soga/cNT24rG4F"
	V "recovered.local/soga/gk1z9v"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"slices"
)

// K4avFiYOPi 对应原 VMess 入站 RTTI 的64字节状态。
type K4avFiYOPi struct {
	*JWN154qJ4
	l8Uq4M      []net.Listener
	dj_cTnaUp6f *V.E3b13Z3FO
	p54XvyV     *U.ZpNFqUVZp
	l_ZoZQDxG7  *J.EafxyUZD
	qgz9PfTS    *N.M1Qn9BA
}

func (v *K4avFiYOPi) UpdateAPIData(data *Q.D5gTIZ5aVg) bool { // 0x1da1e80
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
func (v *K4avFiYOPi) jYtyxkIgmHp(user *U.QKaIWV2) { // 0x1da2180
	if v.p54XvyV == nil {
		return
	}
	if _, ok := v.jc_QPN6.Get(user.GetId()); ok && user.GetBlock() != nil {
		v.p54XvyV.Add(user)
	} else {
		v.p54XvyV.Delete(user.GetId())
	}
}
func (v *K4avFiYOPi) spcNlUL(panels L.Sxcf3wu3L) { // 0x1da2260
	if panels == nil {
		panels = L.Sxcf3wu3L{}
	}
	if C.Q0SDIN() && len(panels) > 88 {
		panels = panels[:88]
	}
	seen := make(map[int64]*L.Aga91y, len(panels))
	added := make([]*U.QKaIWV2, 0)
	for _, panel := range panels {
		speed, limited := v.oSsZZT4EmeG.GetUserSpeedLimit(panel.WZ89pTAOrHK4)
		if existing, ok := v.jc_QPN6.Get(panel.WZ89pTAOrHK4); ok {
			user := existing.(*U.QKaIWV2)
			if _, changed := user.UpdatePanelUser(panel); changed {
				v.jYtyxkIgmHp(user)
				Log.TRTbrquD(-4, "changed user:", user)
			}
			if limited {
				user.SetSpeedLimit(speed)
			}
		} else {
			user, err := U.EOkO09L(panel, v.dwfzZM)
			if err != nil {
				Log.TRTbrquD(8, err)
				continue
			}
			v.jc_QPN6.Add(user)
			v.jYtyxkIgmHp(user)
			Log.TRTbrquD(-4, "add user:", user)
			added = append(added, user)
			if limited {
				user.SetSpeedLimit(speed)
			}
		}
		seen[panel.WZ89pTAOrHK4] = panel
	}
	removed := make([]U.PgaNd3014U, 0)
	v.jc_QPN6.Range(func(id int64, base U.PgaNd3014U) bool { // 0x1e73e00
		if _, ok := seen[id]; !ok {
			user := base.(*U.QKaIWV2)
			user.ResetCtx()
			v.jc_QPN6.Delete(id)
			v.jYtyxkIgmHp(user)
			Log.TRTbrquD(-4, "remove user:", user)
			removed = append(removed, user)
		}
		return true
	})
	if v.l_ZoZQDxG7 != nil {
		v.l_ZoZQDxG7.RemoveUsers(removed)
	}
	if C.Q0SDIN() && len(panels) > 88 {
		v.jc_QPN6.Range(func(id int64, user U.PgaNd3014U) bool { user.ResetCtx(); v.jc_QPN6.Delete(id); return true }) // 0x1e73d60
	} else if len(added) > 0 && v.dj_cTnaUp6f != nil {
		v.dj_cTnaUp6f.GenAuthForNewUsers(added)
	}
	userCountAt50C6DA0 = len(panels)
	Log.V8GN5Stk(0, "current user num: %v", len(panels))
	v.fUxO7UYAH = len(panels)
}
func (v *K4avFiYOPi) Stop() { // 0x1da1c60 and 0x1e73ce0
	Log.TRTbrquD(0, "close VMess server")
	v.yPUbKO()
	for _, listener := range v.l8Uq4M {
		if err := listener.Close(); err != nil {
			Log.TRTbrquD(0, "close VMess server error:", err)
		}
	}
	// 与 VLESS/SS 不同，原 VMess 停止闭包先删除主用户表条目再重置上下文。
	v.jc_QPN6.Range(func(id int64, user U.PgaNd3014U) bool { v.jc_QPN6.Delete(id); user.ResetCtx(); return true })
}
