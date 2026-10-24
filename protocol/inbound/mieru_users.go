package inbound

import (
	"errors"
	"net"
	L "recovered.local/soga/panel/types"
	N "recovered.local/soga/net/limit"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/user"
	Q "recovered.local/soga/panel/registry"
	"slices"
)

// ONGT2EHM40r retains original 40-byte native Mieru inbound state.
type ONGT2EHM40r struct {
	*JWN154qJ4
	d7Fav9U0ykgg []net.Listener
	xVMggjOZQW   *N.O0UldNL4Eq
}

func (v *ONGT2EHM40r) UpdateAPIData(data *Q.D5gTIZ5aVg) bool { //0x1d602a0
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
	v.xVMggjOZQW.UpdateResolver(data.H47AzZsYHB1N.AjXcBfRcAXR)
	return false
}
func (v *ONGT2EHM40r) spcNlUL(panels L.Sxcf3wu3L) { //0x1d605c0; closures1e6dae0/1e6db80
	if panels == nil {
		panels = L.Sxcf3wu3L{}
	}
	if C.Q0SDIN() && len(panels) > 88 {
		panels = panels[:88]
	}
	seen := make(map[int64]struct{})
	for _, p := range panels {
		if speed, limited := v.oSsZZT4EmeG.GetUserSpeedLimit(p.WZ89pTAOrHK4); limited {
			p.LmOKwjha = speed
		}
		if existing, ok := v.jc_QPN6.Get(p.WZ89pTAOrHK4); ok {
			user := existing.(*U.WIZQ28)
			changed, err := user.UpdatePanelUser(p)
			if err != nil {
				Log.V8GN5Stk(8, "update user %v err: %v", existing.GetId(), err)
				continue
			}
			if changed {
				v.jc_QPN6.Add(existing)
				Log.TRTbrquD(-4, "changed user:", existing)
			}
		} else {
			user, err := U.U1KSvhrlJFSB(p, v.dwfzZM)
			if err != nil {
				Log.TRTbrquD(8, err)
				continue
			}
			v.jc_QPN6.Add(user)
			Log.TRTbrquD(-4, "add user:", user)
		}
		seen[p.WZ89pTAOrHK4] = struct{}{}
	}
	v.jc_QPN6.Range(func(id int64, base U.PgaNd3014U) bool {
		if _, ok := seen[id]; !ok {
			user := base.(*U.WIZQ28)
			user.ResetCtx()
			v.jc_QPN6.Delete(id)
			Log.TRTbrquD(-4, "remove user:", user)
		}
		return true
	})
	if C.HuQEoQ() && len(panels) > 88 {
		v.jc_QPN6.Range(func(id int64, user U.PgaNd3014U) bool { user.ResetCtx(); v.jc_QPN6.Delete(id); return true })
	}
	userCountAt50C6DA0 = len(panels)
	Log.V8GN5Stk(0, "current user num: %v", len(panels))
	v.fUxO7UYAH = len(panels)
}
func (v *ONGT2EHM40r) Stop() { //0x1d600a0; closure1eb8740
	Log.TRTbrquD(0, "close mieru server")
	v.yPUbKO()
	for _, l := range v.d7Fav9U0ykgg {
		if err := l.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
			Log.TRTbrquD(0, "close mieru tcp listener error:", err)
		}
	}
	v.jc_QPN6.Range(func(_ int64, user U.PgaNd3014U) bool { user.ResetCtx(); return true })
}
