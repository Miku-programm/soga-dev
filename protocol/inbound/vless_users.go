package inbound

import (
	"github.com/gofrs/uuid"
	"net"
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/protocol/vless"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/user"
	"recovered.local/soga/panel/registry"
	"slices"
)

// KHonGaqbIy_6：原RTTI的48字节VLESS入站状态。
// 仅由原二进制字段及方法恢复，非原作者源码文件。
type KHonGaqbIy_6 struct {
	*JWN154qJ4
	kX5i8U      []net.Listener
	vU51MPl_at_ *genmap.GTFAWF[uuid.UUID, *user.A6CvRbtca]
	i420mh4Yr   *vless.LgF3OB1QX
}

// 原全局0x50c6da0，名称由恢复工程按地址指定。
var userCountAt50C6DA0 int

// UpdateAPIData 对应0x1d97920。节点需重启时立即返回，不先更新用户/规则。
func (v *KHonGaqbIy_6) UpdateAPIData(data *registry.D5gTIZ5aVg) bool {
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

// spcNlUL 对应0x1d97c20及func1@0x1e72a80、func2@0x1e729e0。
// 保留原两次Q0SDIN检查及88上限，也保留失败条目仍进入seen/计数的行为。
func (v *KHonGaqbIy_6) spcNlUL(panels types.Sxcf3wu3L) {
	if panels == nil {
		panels = types.Sxcf3wu3L{}
	}
	if config.Q0SDIN() && len(panels) > 88 {
		panels = panels[:88]
	}
	seen := make(map[int64]*types.Aga91y)
	for _, panel := range panels {
		speed, limited := v.oSsZZT4EmeG.GetUserSpeedLimit(panel.WZ89pTAOrHK4)
		if existing, ok := v.jc_QPN6.Get(panel.WZ89pTAOrHK4); ok {
			user := existing.(*user.A6CvRbtca)
			old, changed, err := user.UpdatePanelUser(panel)
			if err != nil {
				logx.V8GN5Stk(4, "update user %v err: %v", user.GetId(), err)
			} else if changed {
				v.vU51MPl_at_.Delete(old)
				if id, ok := user.GetUuid(); ok {
					v.vU51MPl_at_.Set(id, user)
				}
				logx.TRTbrquD(-4, "changed user:", user)
			}
			if limited {
				user.SetSpeedLimit(speed)
			}
		} else {
			user, err := user.Sa_NiNyuJ_ax(panel, v.dwfzZM)
			if err != nil {
				logx.V8GN5Stk(4, "new user %v err: %v", panel.WZ89pTAOrHK4, err)
			} else if id, ok := user.GetUuid(); ok {
				v.jc_QPN6.Add(user)
				v.vU51MPl_at_.Set(id, user)
				logx.TRTbrquD(-4, "add user:", user)
				if limited {
					user.SetSpeedLimit(speed)
				}
			}
		}
		seen[panel.WZ89pTAOrHK4] = panel
	}
	v.jc_QPN6.Range(func(id int64, base user.PgaNd3014U) bool {
		if _, ok := seen[id]; !ok {
			user := base.(*user.A6CvRbtca)
			user.ResetCtx()
			v.jc_QPN6.Delete(id)
			if uuid, ok := user.GetUuid(); ok {
				v.vU51MPl_at_.Delete(uuid)
			}
			logx.TRTbrquD(-4, "remove user:", user)
		}
		return true
	})
	if config.Q0SDIN() && len(panels) > 88 {
		v.jc_QPN6.Range(func(id int64, user user.PgaNd3014U) bool { user.ResetCtx(); v.jc_QPN6.Delete(id); return true })
	}
	userCountAt50C6DA0 = len(panels)
	logx.V8GN5Stk(0, "current user num: %v", len(panels))
	v.fUxO7UYAH = len(panels)
}

// Stop 对应0x1d97720、func1@0x1eb85c0。关闭监听器后重置所有用户上下文，
// 原函数不会清空用户映射或监听器切片。
func (v *KHonGaqbIy_6) Stop() {
	logx.TRTbrquD(0, "close VLESS server")
	v.yPUbKO()
	for _, listener := range v.kX5i8U {
		if err := listener.Close(); err != nil {
			logx.TRTbrquD(0, "close VLESS server error:", err)
		}
	}
	v.jc_QPN6.Range(func(_ int64, user user.PgaNd3014U) bool { user.ResetCtx(); return true })
}
