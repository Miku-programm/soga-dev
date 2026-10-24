package inbound

import (
	"recovered.local/soga/panel/types"
	"recovered.local/soga/net/limit"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/net/sniff"
	"recovered.local/soga/audit/matcher"
	"strings"
)

// init 对应0x1d3ea80；未知嗅探类型仅警告并跳过，不令初始化失败。
// 原函数追加已有嗅探列表，不在重复调用时清空它。
func (b *JWN154qJ4) init() error {
	for _, kind := range config.DomainSniffProtocols() {
		s, e := sniff.NHzYZHDSd(sniff.MCe3PxlaMYM(kind))
		if e != nil {
			logx.V8GN5Stk(4, "sniff type <%v> not exist", kind)
		} else {
			b.cEUo2gYKI = append(b.cEUo2gYKI, s)
		}
	}
	enabled, _, _, _, _, _ := config.RedisIPLimitConfiguration()
	if enabled {
		b.af8f_J = limit.Cder0rQ(b.zb8f3cHWW8v)
	}
	_, _, banPrivate := config.InspectionConfiguration()
	logx.TRTbrquD(0, "ban private ip:", banPrivate)
	detectors := matcher.Ij9r21E0UOt3{}
	b.e3JHP2MzU.Store(&detectors)
	black, _, _, _ := elCwduDdyGl(nil)
	b.oIgEuicmPL.Store(black)
	white, _, _, _ := elCwduDdyGl(nil)
	b._IoxtH.Store(white)
	return nil
}

// xBzM2TWz4 对应0x1d43260；返回值仅表示原有节点的 Z1TIgHOQ 文本变化。
// BT 嗅探只在第一次节点赋值且 HSA2b5V 为真时追加；保留这一原始行为。
func (b *JWN154qJ4) xBzM2TWz4(node *types.Wr5SVP) bool {
	if config.ForceCloseSSL() {
		node.SwxJNgPpg.UVeahqDBT = false
	}
	if !strings.HasPrefix(node.SwxJNgPpg.ChJnngzb, "/") {
		node.SwxJNgPpg.ChJnngzb = "/" + node.SwxJNgPpg.ChJnngzb
	}
	b.vkyAjc4W(node)
	previous := b.wiy6P3.Load()
	if previous != nil {
		changed := previous.Z1TIgHOQ != node.Z1TIgHOQ
		b.wiy6P3.Store(node)
		b.tA0P9D6faxH()
		return changed
	}
	b.wiy6P3.Store(node)
	b.tA0P9D6faxH()
	if node.HSA2b5V {
		s, e := sniff.NHzYZHDSd("bt")
		if e != nil {
			logx.TRTbrquD(8, e)
		} else {
			b.cEUo2gYKI = append(b.cEUo2gYKI, s)
		}
	}
	return false
}
