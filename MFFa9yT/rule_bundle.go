// Package MFFa9yT 是样本的节点入站装配模块。这里仅包含已由二进制
// 证据恢复的函数，不将未恢复的监听/转发方法写成空实现。
package MFFa9yT

import "recovered.local/soga/xx5fRyrFw"

// rHfs_oNXEJ 原 RTTI 40 字节，字段偏移 0、24、32；显示前缀
// oUi8Xtk3yv 与实际方法包 MFFa9yT 不同，采用方法元数据包路径。
type rHfs_oNXEJ struct {
	a73gMUVz  xx5fRyrFw.Ij9r21E0UOt3
	fj_WUx4re *xx5fRyrFw.ZLM7bds8T
	xDjyw2UC  *xx5fRyrFw.Z5Jdu9
}

// 原汇编0x1d3e7b3—0x1d3e7bd明确返回三组计数，Hex-Rays遗漏了这些返回值。
// elCwduDdyGl 对应 0x1d3e540；域名规则编号按分组后的索引设置。
func elCwduDdyGl(rules xx5fRyrFw.VHWq49iRat) (*rHfs_oNXEJ, int, int, int) {
	domains, ips, ports := rules.Split()
	matchers := make(xx5fRyrFw.Ij9r21E0UOt3, 0, len(domains))
	for i, rule := range domains {
		matchers = append(matchers, xx5fRyrFw.NewPanelDetector(i, rule))
	}
	return &rHfs_oNXEJ{matchers, xx5fRyrFw.HcZBFktYT(ips), xx5fRyrFw.SekQr0GHk(ports)}, len(domains), len(ips), len(ports)
}
