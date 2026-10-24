package config

import (
	"regexp"

	"go.uber.org/dig"
)

// 以下默认值来自 init（0x1319080）的实际解码结果；变量名中的地址
// 用于对应二进制全局，并非推测原作者的变量名。
var (
	blockListPathAt5096DA0    = "/etc/soga/blockList"
	accessLogPatternAt5095A98 = regexp.MustCompile(`^soga-access-([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}):00:00\.log$`)
	logPatternAt5095AB8       = regexp.MustCompile(`soga-(\w+)-([0-9]{4}-[0-9]{2}-[0-9]{2})\.log`)
	dnsPathAt5096DE0          = "/etc/soga/dns.yml"
	whiteListPathAt5096DF0    = "/etc/soga/whiteList"
)

// 0x1319292 调用 HHqPxE8H.T8uNE2(nil)，对应 dig.New()。
// 该包经 Container/Scope 布局及构造流程识别为 Uber dig；v1.15.0 是
// 已核对 Scope 的 15 个字段和 152 字节布局的依赖版本，不代表已确定
// 原二进制所用的精确上游提交。直接使用该依赖保留真实依赖注入功能。
var containerAt5095AC8 = dig.New()
