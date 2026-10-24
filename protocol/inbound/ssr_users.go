package inbound

import (
	"fmt"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/user"
)

// spcNlUL@1d7e1c0，三个闭包@1e705a0/1e70280/1e70160。
// 原函数先启动单端口承载用户，再同步普通凭据用户；失败用户不进入 seen。
func (v *Ze0fOaAko) spcNlUL(panels L.Sxcf3wu3L) {
	if panels == nil {
		panels = L.Sxcf3wu3L{}
	}
	if v.g4fBrm0Bxl {
		carriers := v.wiy6P3.Load().ASJxBbFS.SinglePortUsers()
		if len(carriers) == 0 {
			Log.TRTbrquD(8, "未找到 ssr 承载用户")
			return
		}
		for _, p := range carriers {
			if err := v.mc0MXUU.Update(p); err != nil {
				Log.V8GN5Stk(8, "update ssr user id %v error: %v", p.WZ89pTAOrHK4, err)
			}
		}
	}
	if C.Q0SDIN() && len(panels) > 88 {
		panels = panels[:88]
	}
	seen := make(map[int64]*L.Aga91y)
	for _, p := range panels {
		credential, ok := p.JAfXWwzYrW.(*L.QcqpSL3i6ri)
		if !ok || credential.BskzckkE7qI != 0 {
			continue
		}
		speed, limited := v.oSsZZT4EmeG.GetUserSpeedLimit(p.WZ89pTAOrHK4)
		if existing, ok := v.jc_QPN6.Get(p.WZ89pTAOrHK4); ok {
			u := existing.(*U.LJJkaD)
			if v.g4fBrm0Bxl {
				changed, oldPassword := u.UpdatePanelUser(p)
				if changed {
					if oldPassword != u.GetPassword() {
						v.jTyGfNkAw.Set(u.GetUID(), u)
						v.tpQEL_qKogl.Set(u.GetUID(), u.GetPassword())
						Log.TRTbrquD(-4, "changed user:", u)
					}
					v.ktET5DFha.Range(func(host string, other *U.LJJkaD) bool {
						if other.GetId() == u.GetId() {
							v.ktET5DFha.Delete(host)
							return false
						}
						return true
					})
					host, err := u.GetSSRHost()
					if err != nil {
						Log.TRTbrquD(8, err)
						continue
					}
					v.ktET5DFha.Set(host, u)
				}
			} else if err := v.mc0MXUU.Update(p); err != nil {
				Log.TRTbrquD(8, err)
				continue
			}
			if limited {
				u.SetSpeedLimit(speed)
			}
		} else {
			u, err := U.PXHoe5(p, v.dwfzZM)
			if err != nil {
				Log.TRTbrquD(8, err)
				continue
			}
			v.jc_QPN6.Add(u)
			v.jTyGfNkAw.Set(u.GetUID(), u)
			v.tpQEL_qKogl.Set(u.GetUID(), u.GetPassword())
			host, err := u.GetSSRHost()
			if err != nil {
				Log.TRTbrquD(8, err)
				continue
			}
			v.ktET5DFha.Set(host, u)
			if !v.g4fBrm0Bxl {
				if err := v.mc0MXUU.v_Q8IiV(u); err != nil {
					Log.TRTbrquD(8, err)
					continue
				}
			}
			Log.TRTbrquD(-4, "add user:", u)
			if limited {
				u.SetSpeedLimit(speed)
			}
		}
		seen[p.WZ89pTAOrHK4] = p
	}
	v.jc_QPN6.Range(func(id int64, base U.PgaNd3014U) bool {
		if _, ok := seen[id]; ok {
			return true
		}
		u := base.(*U.LJJkaD)
		u.ResetCtx()
		v.jc_QPN6.Delete(id)
		v.jTyGfNkAw.Delete(u.GetUID())
		v.tpQEL_qKogl.Delete(u.GetUID())
		if host, err := u.GetSSRHost(); err == nil {
			v.ktET5DFha.Delete(host)
		}
		if !v.g4fBrm0Bxl {
			if err := v.mc0MXUU.hmycHk6ebPwS(u.SSRPort()); err != nil {
				Log.TRTbrquD(8, "stop port", u.SSRPort(), "error:", err)
			}
		}
		Log.TRTbrquD(-4, "remove user:", u)
		return true
	})
	if C.Q0SDIN() && len(panels) > 88 {
		v.jc_QPN6.Range(func(id int64, base U.PgaNd3014U) bool {
			u := base.(*U.LJJkaD)
			u.ResetCtx()
			v.jc_QPN6.Delete(id)
			v.jTyGfNkAw.Delete(u.GetUID())
			v.tpQEL_qKogl.Delete(u.GetUID())
			if host, err := u.GetSSRHost(); err == nil {
				v.ktET5DFha.Delete(host)
			}
			return true
		})
	}
	userCountAt50C6DA0 = len(panels)
	Log.V8GN5Stk(0, "current user num: %v", len(panels))
	v.fUxO7UYAH = len(panels)
}

// Update@1d8b1e0。已有端口的 credential 更新后，先关闭旧监听再新建；
// 新用户对象在这里不加入公共用户集合，这与主同步循环是两条独立路径。
func (m *xdWl88) Update(p *L.Aga91y) error {
	c, ok := p.JAfXWwzYrW.(*L.QcqpSL3i6ri)
	if !ok {
		return fmt.Errorf("user %v has no ssr credential", p.WZ89pTAOrHK4)
	}
	m.RLock()
	old := m.sSKJ79jFt[c.YrjVl_2g_S]
	m.RUnlock()
	if old != nil {
		if changed, _ := old.gHBoVOaQY.UpdatePanelUser(p); changed {
			if err := m.hmycHk6ebPwS(c.YrjVl_2g_S); err != nil {
				Log.TRTbrquD(4, "stop port err:", err)
			}
			return m.v_Q8IiV(old.gHBoVOaQY)
		}
		return nil
	}
	u, err := U.PXHoe5(p, m.spzpWtrEXC1G.dwfzZM)
	if err != nil {
		return err
	}
	return m.v_Q8IiV(u)
}

// hmycHk6ebPwS@1d8c4c0。先从表删除再关闭，不在锁内执行 Close。
func (m *xdWl88) hmycHk6ebPwS(port int) error {
	m.Lock()
	p := m.sSKJ79jFt[port]
	delete(m.sSKJ79jFt, port)
	m.Unlock()
	if p == nil {
		return fmt.Errorf("can not find ssr port: %v", port)
	}
	Log.TRTbrquD(0, "stop ssr port")
	return p.ojfxJfwLalz3()
}
