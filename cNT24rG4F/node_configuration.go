// 根据原样本 IDA 指令、RTTI 和静态解码字符串恢复；非作者文件原文。
package cNT24rG4F

import (
	"strconv"
	"strings"
)

// nodeConfigurationAt50A2AE0 对应 init.3 在 0x1321032—0x1321040
// 以类型描述符 0x33f2088 注册的全局配置结构；变量名由分析者指定。
// 声明时保持零值；configuration_initializers.go 恢复的 init.3 注册该结构，
// _NA1kft 收集 default 标签，TvBECnD9fM6 经 bFnBl5Ak9 装载默认值及配置，
// 随后调用本文件的 jkMcB1_ 规范化节点字段。
var nodeConfigurationAt50A2AE0 i7BY8ClkrDX

// jkMcB1_ 对应 0x13210e0。只规范化 type、api、sub_type，跳过无效节点 ID。
func jkMcB1_() {
	c := &nodeConfigurationAt50A2AE0
	c.EtRWBSEhLQc = strings.ToLower(c.EtRWBSEhLQc)
	c.NJSpG6 = strings.ToLower(c.NJSpG6)
	c.AG_mSFklcvj = strings.ToLower(c.AG_mSFklcvj)
	c.t09QzN = make([]int, 0, len(c.N1Z81aDzg))
	for _, text := range c.N1Z81aDzg {
		if n, err := strconv.Atoi(text); err == nil {
			c.t09QzN = append(c.t09QzN, n)
		}
	}
}

// DnhHoT2aRvs 对应 0x1321f40。读取 server_type，返回名称，不回写配置。
// 三个混淆常量分别在 0x1322011、0x1322087、0x1322116 的调用点
// 静态解码为 shadowsocks、shadowsocksr、hysteria。
func DnhHoT2aRvs() string {
	value := strings.ToLower(nodeConfigurationAt50A2AE0.MnXMlyI)
	switch value {
	case "v2ray", "vmess":
		return "v2ray"
	case "ss", "shadowsocks":
		return "ss"
	case "ssr", "shadowsocksr":
		return "ssr"
	case "trojan":
		return "trojan"
	case "vless":
		return "vless"
	case "hysteria":
		return "hysteria"
	case "mieru":
		return "mieru"
	case "anytls":
		return "anytls"
	default:
		return value
	}
}
