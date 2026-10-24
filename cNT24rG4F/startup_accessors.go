package cNT24rG4F

// 以下桥接仅表达原 main/中央启动函数直接读取的包全局；不是新增配置。

// ConfigurationFlagNames 对应 main.init.0@0x20187c0 对 0x5095AA0 的键枚举。
// 保留 map 的无序枚举，返回独立切片。
func ConfigurationFlagNames() []string {
	names := make([]string, 0, len(defaultsAt5095AA0))
	for name := range defaultsAt5095AA0 {
		names = append(names, name)
	}
	return names
}

// StartupVersion 对应 main.main/InitLinux 对 0x5029EE0 的读取。
func StartupVersion() string { return versionAt5029EE0 }

// StartupPanelService 对应 InitLinux 的 0x509B3E0/0x509B3F0 字符串。
func StartupPanelService() (url, key string) {
	return configAt509B3E0.ZkkjP9, configAt509B3E0.TzpbaXGDeaz
}

// StartupAutoUpdate 对应 so4m0_oRopC 的 0x50A2938。
func StartupAutoUpdate() bool { return configAt50A2900.F3BbbG_1qsA }

// StartupGeoUpdate 对应 0x50A29A2、0x50A29A8、0x50A29B0、0x50A29C0。
func StartupGeoUpdate() (enabled bool, hours int, siteURL, ipURL string) {
	c := &configAt50A2900
	return c.JRQZn_, c.BIdU1vYs, c.YGrHEaVoNKL, c.YVMoJnAY
}

// StartupPrometheus 对应中央启动对 0x509B3C0 的内联字段读取。
func StartupPrometheus() (enabled bool, addr string, port int) {
	c := &configAt509B3C0
	return c.PVSao2JIq, c.OAHNXtGh, c.DXN_loYa
}

// StartupPprof 对应中央启动对 0x509B3A0 的内联字段读取。
func StartupPprof() (enabled bool, addr string, port int) {
	c := &configAt509B3A0
	return c.PVSao2JIq, c.OAHNXtGh, c.DXN_loYa
}
