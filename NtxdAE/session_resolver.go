package NtxdAE

import (
	"context"
	"recovered.local/soga/BaEWsPy"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/nzATvajt"
	"slices"
	"sync/atomic"
	"time"
)

// O0UldNL4Eq 来自原 RTTI 0x36cc830，64 字节。显示前缀 fHkHpz8
// 与方法包路径 NtxdAE 不同；这是原二进制类型的源码表示。
type O0UldNL4Eq struct {
	RCaEGDlm   context.Context
	Cebit_sS   int64
	RacA3he0TX *nzATvajt.ZpNFqUVZp
	kyWKarzX   atomic.Pointer[BaEWsPy.D23y2888a_]
	lNzIhzjHYz LFWTTXF7JR.U863Mm01Nre
}

// GetResolver 对应 0x183d2a0。
func (s *O0UldNL4Eq) GetResolver() *BaEWsPy.D23y2888a_ { return s.kyWKarzX.Load() }

// UpdateResolver 对应 0x183cec0，先合并节点规则与全局规则；原逻辑只比较
// 规则列表，独立改变上游/TTL/策略不会触发替换，保留这一判断条件。
func (s *O0UldNL4Eq) UpdateResolver(rules LFWTTXF7JR.U863Mm01Nre) {
	global, servers, ttlMinutes, strategy := cNT24rG4F.ResolverConfiguration()
	merged := slices.Concat(rules, global)
	if s.kyWKarzX.Load() == nil || !s.lNzIhzjHYz.Equals(merged) {
		resolver := BaEWsPy.DAlA5L(s.RCaEGDlm, &BaEWsPy.PQL1aLrtju{
			MLo0o543OAaK: merged, Zfr5Q8ybETV: servers, X6T7XOmp: cNT24rG4F.WbFtukHA9(strategy),
			C_TYygjEa0: time.Duration(ttlMinutes) * time.Minute, DMa5Vg24rVR: true,
		})
		s.kyWKarzX.Store(resolver)
		s.lNzIhzjHYz = merged
	}
}
