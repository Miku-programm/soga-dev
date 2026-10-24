// Package MFFa9yT 是样本的节点入站装配模块。这里仅包含已由二进制
// 证据恢复的函数，不将未恢复的监听/转发方法写成空实现。
package inbound

import "recovered.local/soga/audit/matcher"

// rHfs_oNXEJ 原 RTTI 40 字节，字段偏移 0、24、32；显示前缀
// oUi8Xtk3yv 与实际方法包 MFFa9yT 不同，采用方法元数据包路径。
type rHfs_oNXEJ struct {
	a73gMUVz  matcher.Ij9r21E0UOt3
	fj_WUx4re *matcher.ZLM7bds8T
	xDjyw2UC  *matcher.Z5Jdu9
}

// 原汇编0x1d3e7b3—0x1d3e7bd明确返回三组计数，Hex-Rays遗漏了这些返回值。
// elCwduDdyGl 对应 0x1d3e540；域名规则编号按分组后的索引设置。
func elCwduDdyGl(rules matcher.VHWq49iRat) (*rHfs_oNXEJ, int, int, int) {
	domains, ips, ports := rules.Split()
	matchers := make(matcher.Ij9r21E0UOt3, 0, len(domains))
	for i, rule := range domains {
		matchers = append(matchers, matcher.NewPanelDetector(i, rule))
	}
	return &rHfs_oNXEJ{matchers, matcher.HcZBFktYT(ips), matcher.SekQr0GHk(ports)}, len(domains), len(ips), len(ports)
}
