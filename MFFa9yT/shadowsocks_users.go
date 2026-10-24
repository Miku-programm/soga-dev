package MFFa9yT

import (
	"net"
	"recovered.local/soga/GTiWvwS"
	"recovered.local/soga/J1P4KNvfI_J"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/NtxdAE"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/nzATvajt"
	"recovered.local/soga/q5bqCgTdxaTM"
	"runtime"
	"slices"
)

// Lp07FzBLx9V is the original 80-byte native Shadowsocks inbound state
// (RTTI 0x3728b28), reconstructed from the target binary.
type Lp07FzBLx9V struct {
	*JWN154qJ4
	vLRjbS7NDeZZ []net.Listener
	sSqGyaLW     []J1P4KNvfI_J.KsKg3yls
	wZhn2Y       *J1P4KNvfI_J.EafxyUZD
	gtKatlWm0    *NtxdAE.M1Qn9BA
	mCYkzsUd     *GTiWvwS.N07NaA
}

// UpdateAPIData corresponds to 0x1d6e9a0.
func (v *Lp07FzBLx9V) UpdateAPIData(data *q5bqCgTdxaTM.D5gTIZ5aVg) bool {
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

// spcNlUL corresponds to 0x1d6eca0, with closures 0x1e6f6c0/0x1e6f600.
// Failed updates are deliberately not marked seen, and can therefore be included
// twice in the cache-invalidation slice. The two original license checks differ.
func (v *Lp07FzBLx9V) spcNlUL(panels LFWTTXF7JR.Sxcf3wu3L) {
	if panels == nil {
		panels = LFWTTXF7JR.Sxcf3wu3L{}
	}
	if cNT24rG4F.Q0SDIN() && len(panels) > 88 {
		panels = panels[:88]
	}
	seen := make(map[int64]struct{})
	removed := []nzATvajt.PgaNd3014U{}
	for _, panel := range panels {
		speed, limited := v.oSsZZT4EmeG.GetUserSpeedLimit(panel.WZ89pTAOrHK4)
		if existing, ok := v.jc_QPN6.Get(panel.WZ89pTAOrHK4); ok {
			user := existing.(*nzATvajt.WtK9HQSy66)
			changed, err := user.UpdatePanelUser(panel)
			if err != nil {
				muc279Mpa6S4.V8GN5Stk(8, "update user %v err: %v", user.GetId(), err)
				removed = append(removed, user)
			} else {
				if changed {
					removed = append(removed, user)
					v.jc_QPN6.Add(user)
					if v.mCYkzsUd != nil {
						v.mCYkzsUd.UpdateUser(user)
					}
					muc279Mpa6S4.TRTbrquD(-4, "changed user:", user)
				}
				seen[panel.WZ89pTAOrHK4] = struct{}{}
				if limited {
					user.SetSpeedLimit(speed)
				}
			}
		} else {
			user, err := nzATvajt.MbpJoYH(panel, v.dwfzZM)
			if err != nil {
				muc279Mpa6S4.TRTbrquD(8, err)
			} else {
				v.jc_QPN6.Add(user)
				if v.mCYkzsUd != nil {
					v.mCYkzsUd.UpdateUser(user)
				}
				muc279Mpa6S4.TRTbrquD(-4, "add user:", user)
				seen[panel.WZ89pTAOrHK4] = struct{}{}
				if limited {
					user.SetSpeedLimit(speed)
				}
			}
		}
	}
	v.jc_QPN6.Range(func(id int64, base nzATvajt.PgaNd3014U) bool {
		if _, ok := seen[id]; !ok {
			user := base.(*nzATvajt.WtK9HQSy66)
			user.ResetCtx()
			removed = append(removed, user)
			v.jc_QPN6.Delete(id)
			if v.mCYkzsUd != nil {
				v.mCYkzsUd.RemoveUser(id)
			}
			muc279Mpa6S4.TRTbrquD(-4, "remove user:", user)
		}
		return true
	})
	if v.wZhn2Y != nil {
		v.wZhn2Y.RemoveUsers(removed)
	}
	if cNT24rG4F.HuQEoQ() && len(panels) > 88 {
		v.jc_QPN6.Range(func(id int64, user nzATvajt.PgaNd3014U) bool {
			user.ResetCtx()
			v.jc_QPN6.Delete(id)
			if v.mCYkzsUd != nil {
				v.mCYkzsUd.RemoveUser(id)
			}
			return true
		})
	}
	userCountAt50C6DA0 = len(panels)
	muc279Mpa6S4.V8GN5Stk(0, "current user num: %v", len(panels))
	v.fUxO7UYAH = len(panels)
}

// Stop corresponds to 0x1d6e820 and closure 0x1eb86e0. Close errors are
// discarded, user contexts reset, and an explicit GC follows, as in the target.
func (v *Lp07FzBLx9V) Stop() {
	muc279Mpa6S4.TRTbrquD(0, "close shadowsocks server")
	v.yPUbKO()
	for _, l := range v.vLRjbS7NDeZZ {
		_ = l.Close()
	}
	for _, p := range v.sSqGyaLW {
		_ = p.Close()
	}
	v.jc_QPN6.Range(func(_ int64, u nzATvajt.PgaNd3014U) bool { u.ResetCtx(); return true })
	runtime.GC()
}
