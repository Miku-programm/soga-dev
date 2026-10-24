package MFFa9yT

import (
	"recovered.local/soga/BaEWsPy"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"slices"
	"time"
)

// vkyAjc4W 对应0x1d434e0；只根据合并规则变化更新解析器，其他配置变化单独不触发替换。
func (b *JWN154qJ4) vkyAjc4W(node *LFWTTXF7JR.Wr5SVP) {
	global, servers, ttl, strategy := cNT24rG4F.ResolverConfiguration()
	merged := slices.Concat(node.AjXcBfRcAXR, global)
	if b.ffju0Enjo.Load() == nil || !b.icS5YvmZ.Equals(merged) {
		r := BaEWsPy.DAlA5L(b.zb8f3cHWW8v, &BaEWsPy.PQL1aLrtju{MLo0o543OAaK: merged, Zfr5Q8ybETV: servers, X6T7XOmp: cNT24rG4F.WbFtukHA9(strategy), C_TYygjEa0: time.Duration(ttl) * time.Minute, DMa5Vg24rVR: true})
		b.ffju0Enjo.Store(r)
		b.icS5YvmZ = merged
	}
}

// tA0P9D6faxH 对应0x1d43180。原 UCOMISD/JB 对 NaN 也进入转换分支。
func (b *JWN154qJ4) tA0P9D6faxH() {
	speed := b.wiy6P3.Load().TtWKyh
	if speed <= 0 {
		b.dwfzZM.UpdateLimit(0)
	} else {
		b.dwfzZM.UpdateLimit(int64(1024 * (0.125 * (speed * 1024))))
	}
}
