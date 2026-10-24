package MFFa9yT

import (
	"net"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"slices"
	"sync"
)

// EfaIFpp 对应原生Trojan入站80字节RTTI，认证索引为原sync.Map。
type EfaIFpp struct {
	*JWN154qJ4
	cDyEMW2    []net.Listener
	d9Cxrbh1g7 sync.Map
}

func (v *EfaIFpp) UpdateAPIData(data *Q.D5gTIZ5aVg) bool { // 0x1d8f220
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

// spcNlUL @0x1d8f520及闭包0x1e72160/0x1e72200。
func (v *EfaIFpp) spcNlUL(panels L.Sxcf3wu3L) {
	if panels == nil {
		panels = L.Sxcf3wu3L{}
	}
	if C.Q0SDIN() && len(panels) > 88 {
		panels = panels[:88]
	}
	seen := make(map[int64]*L.Aga91y)
	for _, panel := range panels {
		speed, limited := v.oSsZZT4EmeG.GetUserSpeedLimit(panel.WZ89pTAOrHK4)
		if existing, ok := v.jc_QPN6.Get(panel.WZ89pTAOrHK4); ok {
			user := existing.(*U.DWaKIV1T)
			if previous := user.UpdatePanelUser(panel); previous != nil {
				v.d9Cxrbh1g7.Delete(*previous)
				v.d9Cxrbh1g7.Store(user.IdeY3L, user)
				Log.TRTbrquD(-4, "changed user:", user)
			}
			if limited {
				user.SetSpeedLimit(speed)
			}
		} else {
			user, err := U.KUa7qViw(panel, v.dwfzZM)
			if err != nil {
				Log.TRTbrquD(8, err)
				continue
			}
			v.jc_QPN6.Add(user)
			v.d9Cxrbh1g7.Store(user.IdeY3L, user)
			Log.TRTbrquD(-4, "add user:", user)
			if limited {
				user.SetSpeedLimit(speed)
			}
		}
		seen[panel.WZ89pTAOrHK4] = panel
	}
	v.jc_QPN6.Range(func(id int64, base U.PgaNd3014U) bool {
		if _, ok := seen[id]; !ok {
			user := base.(*U.DWaKIV1T)
			user.ResetCtx()
			v.jc_QPN6.Delete(id)
			v.d9Cxrbh1g7.Delete(user.IdeY3L)
			Log.TRTbrquD(-4, "remove user:", user)
		}
		return true
	})
	if C.Q0SDIN() && len(panels) > 88 {
		v.jc_QPN6.Range(func(id int64, user U.PgaNd3014U) bool { user.ResetCtx(); v.jc_QPN6.Delete(id); return true })
	}
	userCountAt50C6DA0 = len(panels)
	Log.V8GN5Stk(0, "current user num: %v", len(panels))
	v.fUxO7UYAH = len(panels)
}

// Stop @0x1d8f020；与原VMess不同，原Trojan不清空用户表或认证索引。
func (v *EfaIFpp) Stop() {
	Log.TRTbrquD(0, "close trojan server listener")
	v.yPUbKO()
	for _, listener := range v.cDyEMW2 {
		if err := listener.Close(); err != nil {
			Log.TRTbrquD(0, "close trojan server error:", err)
		}
	}
	v.jc_QPN6.Range(func(_ int64, user U.PgaNd3014U) bool { user.ResetCtx(); return true })
}
