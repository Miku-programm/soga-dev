package cNT24rG4F

import "time"

// 恢复桥接：原 VMess 构造内联读取 0x50a02c0/0x50a02e0。
func VmessReduceMemory() bool { return configAt50A02C0.Ff3xSe || configAt50A02C0.SElnoFh_fIwr }

// 原构造 0x1da0d40 读取 0x50c8460 起四项；非正数使用 30/60/600。
func VmessInvalidAccessSettings() *Nl62Lnd_ {
	c := configAt50C8440
	count, duration, forbidden := c.Zdj5DFrb4WqO, c.BPUCSK_mS81, c.ITmtDnIa8J
	if count <= 0 {
		count = 30
	}
	if duration <= 0 {
		duration = 60
	}
	if forbidden <= 0 {
		forbidden = 600
	}
	return &Nl62Lnd_{SEU_Z687ugp: c.H1kZa6, GxuEkr: count, QS255wlN7m2W: time.Duration(duration) * time.Second, NjiRkFvidOaA: time.Duration(forbidden) * time.Second}
}
