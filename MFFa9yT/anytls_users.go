package MFFa9yT

import (
	"net"
	CA "recovered.local/soga/CAodbhZf"
	G "recovered.local/soga/G7A7csgrs62p"
	P "recovered.local/soga/HHGbXWqYaYt_"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Log "recovered.local/soga/muc279Mpa6S4"
	U "recovered.local/soga/nzATvajt"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"slices"
	"strings"
)

// R_n5ucb 是原生 AnyTLS 入站的 72 字节结构（RTTI 0x3727a78）。
type R_n5ucb struct {
	*JWN154qJ4
	qjukoB    []net.Listener
	lLSuDoo   *G.GTFAWF[[32]byte, *U.HaB1NEnqcph]
	aeJMgHLjP string
	yO3uyY    CA.EahpvU_iseQ[*P.FZBbzkJPWbrk]
}

func (v *R_n5ucb) UpdateAPIData(data *Q.D5gTIZ5aVg) bool { // 0x1d34040
	if v.xBzM2TWz4(data.H47AzZsYHB1N) {
		return true
	}
	scheme := strings.TrimSpace(data.H47AzZsYHB1N.QaZjDzw95.UPnToXrM)
	if scheme != v.aeJMgHLjP {
		if padding := P.R_GY_FGx([]byte(scheme)); padding != nil {
			v.aeJMgHLjP = scheme
			v.yO3uyY.Store(padding)
			Log.V8GN5Stk(0, "update padding scheme\n%s\n", scheme)
		} else {
			Log.V8GN5Stk(0, "invalid padding scheme\n%s\n", scheme)
		}
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

func (v *R_n5ucb) spcNlUL(panels L.Sxcf3wu3L) { // 0x1d345a0
	if panels == nil {
		panels = L.Sxcf3wu3L{}
	}
	if C.Q0SDIN() && len(panels) > 88 {
		panels = panels[:88]
	}
	seen := make(map[int64]*L.Aga91y)
	for _, panel := range panels {
		speed, limited := v.oSsZZT4EmeG.GetUserSpeedLimit(panel.WZ89pTAOrHK4)
		if base, ok := v.jc_QPN6.Get(panel.WZ89pTAOrHK4); ok {
			user := base.(*U.HaB1NEnqcph)
			old, changed, err := user.UpdatePanelUser(panel)
			if err != nil {
				Log.V8GN5Stk(4, "update user %v err: %v", user.GetId(), err)
			} else if changed {
				v.lLSuDoo.Delete(old)
				v.lLSuDoo.Set(user.GetAuth(), user)
				Log.TRTbrquD(-4, "changed user:", user)
			}
			if limited {
				user.SetSpeedLimit(speed)
			}
		} else {
			user, err := U.NZl_wC1qYi(panel, v.dwfzZM)
			if err != nil {
				Log.V8GN5Stk(4, "new user %v err: %v", panel.WZ89pTAOrHK4, err)
			} else {
				v.jc_QPN6.Add(user)
				v.lLSuDoo.Set(user.GetAuth(), user)
				Log.TRTbrquD(-4, "add user:", user)
				if limited {
					user.SetSpeedLimit(speed)
				}
			}
		}
		seen[panel.WZ89pTAOrHK4] = panel
	}
	v.jc_QPN6.Range(func(id int64, base U.PgaNd3014U) bool {
		if _, ok := seen[id]; !ok {
			user := base.(*U.HaB1NEnqcph)
			user.ResetCtx()
			v.jc_QPN6.Delete(id)
			v.lLSuDoo.Delete(user.GetAuth())
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

func (v *R_n5ucb) Stop() { // 0x1d33e40
	Log.TRTbrquD(0, "close Anytls server")
	v.yPUbKO()
	for _, listener := range v.qjukoB {
		if err := listener.Close(); err != nil {
			Log.TRTbrquD(0, "close Anytls server error:", err)
		}
	}
	v.jc_QPN6.Range(func(_ int64, user U.PgaNd3014U) bool { user.ResetCtx(); return true })
}
