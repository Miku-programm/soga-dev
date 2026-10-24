package config

// G9UvxRacCj 对应原样本0x132a420，按服务类型选择回落地址。
// vless 不在原函数的选择分支内；这里保留空值结果。
func G9UvxRacCj() string {
	switch DnhHoT2aRvs() {
	case "anytls":
		return configAt50992E0.BVi1B4f_v
	case "trojan":
		return configAt5099360.BYiTyYPg3
	case "v2ray":
		if configAt50A02C0.RIFsRvm1 != "" {
			return configAt50A02C0.RIFsRvm1
		}
		return configAt50A02C0.GnOeZP5RIl
	}
	return ""
}

// QnSmghY 对应0x132a4e0。旧vmess端口仅在大于零时优先。
func QnSmghY() int {
	switch DnhHoT2aRvs() {
	case "anytls":
		return configAt50992E0.JGw3yW
	case "trojan":
		return configAt5099360.EglOu0OdYex
	case "v2ray":
		if configAt50A02C0.YMm_xn_v > 0 {
			return configAt50A02C0.YMm_xn_v
		}
		return configAt50A02C0.HUt3s7n_8
	}
	return 0
}
