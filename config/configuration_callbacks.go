package config

import "strings"

// 以下函数根据所提供样本的 IDA 反汇编恢复；全局结构和字段由 RTTI 核对。
// init6func1 对应 config.init.6.func1，地址 0x1346d80。
func init6func1() {
	if configAt509CB40.PBvoBS2NX5 <= 0 {
		configAt509CB40.PBvoBS2NX5 = 10
	}
}

// qZO6G3A 对应地址 0x1329620。TCP 的零值与 UDP 的零值处理不同。
func qZO6G3A() {
	if configAt509BEA0.Hni9vNgOp4g < 0 {
		configAt509BEA0.Hni9vNgOp4g = 120
	}
	if configAt509BEA0.JtFIOYSoCNi <= 0 {
		configAt509BEA0.JtFIOYSoCNi = 2
	}
}

// tzfa9ae 对应地址 0x132a220。
func tzfa9ae() {
	c := &configAt50A2900
	if c.Qxaw2_S <= 0 {
		c.Qxaw2_S = 60
	}
	if c.Fak_7LgVV <= 0 {
		c.Fak_7LgVV = 60
	}
	if c.KS63SD < 0 {
		c.KS63SD = 10
	}
	if c.GozCogFPm <= 0 {
		c.GozCogFPm = 2048
	}
	// 原机器码的双 MINSD、按位 OR 和符号翻转保留 NaN，且将负零变为正零。
	c.QzlEP1r3t = max(c.QzlEP1r3t, 0)
	c.YyLGOC = max(c.YyLGOC, 0)
	c.ItKGl9keQrlY = max(c.ItKGl9keQrlY, 0)
	c.P6KMmld = max(c.P6KMmld, 0)
	if uint(c.JQC3mlZzpS69) > 32 {
		c.JQC3mlZzpS69 = 32
	}
	if uint(c.BkelkzTWy) > 128 {
		c.BkelkzTWy = 128
	}
	if c.BIdU1vYs <= 0 {
		c.BIdU1vYs = 24
	}
	c.KI4eIM8J = max(c.KI4eIM8J, 0)
	c.DhXdMvN = max(c.DhXdMvN, 0)
}

// init14func1 对应 config.init.14.func1，地址 0x1346ca0。
func init14func1() {
	if uint(configAt509B3A0.DXN_loYa-1) > 65534 {
		configAt509B3A0.DXN_loYa = 21010
	}
}

// init15func1 对应 config.init.15.func1，地址 0x1346c80。
func init15func1() {
	if uint(configAt509B3C0.DXN_loYa-1) > 65534 {
		configAt509B3C0.DXN_loYa = 19090
	}
}

// ekZtUL9pJS 对应地址 0x132a7a0。
func ekZtUL9pJS() {
	if configAt509C940.Eb9vzdoE <= 0 {
		configAt509C940.Eb9vzdoE = 60
	}
}

// init10func1 对应 config.init.10.func1，地址 0x1346cc0。
func init10func1() {
	if configAt509B6A0.HSBWD7 == "false" {
		configAt509B6A0.HSBWD7 = ""
	}
	configAt509B6A0.HSBWD7 = strings.TrimRight(configAt509B6A0.HSBWD7, "/")
}
