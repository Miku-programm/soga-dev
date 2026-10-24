package inbound

import (
	"recovered.local/soga/util/logx"
	"recovered.local/soga/audit/matcher"
)

// v5VChcYe 对应0x1d3ee00；每次调用都建立并原子替换黑名单，再记录三类数量。
func (b *JWN154qJ4) v5VChcYe(rules matcher.VHWq49iRat) {
	bundle, domains, ips, ports := elCwduDdyGl(rules)
	b.oIgEuicmPL.Store(bundle)
	logx.V8GN5Stk(0, "refresh block list: domain rules: %v, ip rules: %v, port rules: %v", domains, ips, ports)
}

// ctArZJ 对应0x1d3ef80；原函数不修改保存的 x1djY10H 切片。
func (b *JWN154qJ4) ctArZJ(rules matcher.VHWq49iRat) {
	bundle, domains, ips, ports := elCwduDdyGl(rules)
	b._IoxtH.Store(bundle)
	logx.V8GN5Stk(0, "refresh white list: domain rules: %v, ip rules: %v, port rules: %v", domains, ips, ports)
}
