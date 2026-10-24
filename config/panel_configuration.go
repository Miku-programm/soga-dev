package config

import "recovered.local/soga/util/logx"

// 面板转换函数对原配置全局字段的直接读取桥接。
func PanelForbiddenBitTorrent() bool { return configAt50A2900.JGgVNMp }
func PanelNodeSpeedLimit() float64   { return configAt50A2900.QzlEP1r3t }
func PanelUserConnectionLimit() int  { return configAt50A2900.ItKGl9keQrlY }
func PanelUserSpeedLimit() float64   { return configAt50A2900.YyLGOC }
func PanelUserTCPLimit() int         { return configAt50A2900.P6KMmld }

// H7eU0uGXU 对应 0x13216c0；旧字段仅在新字段为空时使用。
func H7eU0uGXU() string {
	c := &nodeConfigurationAt50A2AE0
	if c.UnkaBGBuOUXV != "" {
		return c.UnkaBGBuOUXV
	}
	if c.MKrU9XEJ != "" {
		logx.TRTbrquD(4, "'webapi_mukey' is deprecated, please use 'webapi_key'")
	}
	return c.MKrU9XEJ
}

// SGx0HPPP 对应 0x13212e0。每个节点构造独立配置，URL 仅补末尾斜杠。
func SGx0HPPP() []*FCEAfQJwDA7e {
	var result []*FCEAfQJwDA7e
	c := &nodeConfigurationAt50A2AE0
	for _, nodeID := range c.t09QzN {
		serverType := DnhHoT2aRvs()
		url := c.VfiX6TsWZxF
		if len(url) == 0 || url[len(url)-1] != '/' {
			url += "/"
		}
		key := H7eU0uGXU()
		result = append(result, &FCEAfQJwDA7e{
			OS91X026uML: FSOZSH7g(c.EtRWBSEhLQc), ZyAhkq5th8bf: Mpn2JltM(serverType),
			F3PCiaIra: HWLVDK2(c.NJSpG6), W_ogGtF: url, Bd08Li: key, Yf1KqSqJpNI7: nodeID,
			EotADOuOVu: c.ApfXLZLmFXh, Z9qKfN7e_: c.Mk5cDiQzaj,
			XU27D_AD: c.CsBHZVUxTO, GcBQVp1f: c.ZGaSmv, JWhomGW: c.ASSy13OjR6aJ,
		})
	}
	return result
}

// SIZIzCa、KvDEtVwdr1rT 分别对应 0x13217a0、0x13217e0。
func SIZIzCa() bool      { return DnhHoT2aRvs() == "v2ray" }
func KvDEtVwdr1rT() bool { return DnhHoT2aRvs() == "vless" }

// PanelType 连接原启动函数直接读取的全局字段。
func PanelType() string { return nodeConfigurationAt50A2AE0.EtRWBSEhLQc }
