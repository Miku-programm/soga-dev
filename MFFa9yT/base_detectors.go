package MFFa9yT

import (
	"recovered.local/soga/jAP8Cn"
	"recovered.local/soga/xx5fRyrFw"
)

// afT7skk 对应0x1d438e0。Geo数据更新或条目数量改变时直接替换；
// 否则仅比较同位置规则文本，复用旧matcher及其已初始化状态（包括旧ID）。
func (b *JWN154qJ4) afT7skk(next xx5fRyrFw.Ij9r21E0UOt3) {
	timestamp, updated := jAP8Cn.MCOm37H(b.f6mOarj)
	b.f6mOarj = timestamp
	previous := *b.e3JHP2MzU.Load()
	if len(previous) != len(next) || updated {
		b.e3JHP2MzU.Store(&next)
		return
	}
	merged := make(xx5fRyrFw.Ij9r21E0UOt3, len(next))
	changed := false
	for i, old := range previous {
		if old.GetRule() == next[i].GetRule() {
			merged[i] = old
		} else {
			merged[i] = next[i]
			changed = true
		}
	}
	if changed {
		b.e3JHP2MzU.Store(&merged)
	}
}
