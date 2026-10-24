package config

import "recovered.local/soga/panel/types"

// ResolverConfiguration 为原全局配置提供源码桥接。规则和上游切片
// 保留原底层存储，TTL 保持原分钟单位，由调用者按原流程换算。
func ResolverConfiguration() (rules types.U863Mm01Nre, servers []string, ttlMinutes int, strategy string) {
	if current := dnsRulesAt5095AD0.Load(); current != nil {
		rules = *current
	}
	return rules, configAt509CB40.RHZA1kqpMSz, configAt509CB40.PBvoBS2NX5, configAt509CB40.EAOjNaQ6
}
