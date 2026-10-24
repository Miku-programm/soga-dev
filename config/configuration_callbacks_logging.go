package config

import (
	"strings"

	"recovered.local/soga/util/logx"
)

// uQm8ubPKg 对应地址 0x131c9a0；混淆字符串由受限指令仿真恢复。
func uQm8ubPKg() {
	c := &configAt50A2260
	mode := c.WO1HZMaFg6Mx
	if mode == "false" {
		mode = ""
	}
	mode = strings.ToLower(strings.TrimSpace(mode))
	if mode == "" {
		mode = "local"
	} else if mode != "local" && mode != "clickhouse" {
		if c.PmIFLM {
			logx.V8GN5Stk(4, "unknown access_log_mode <%v>, use %v", mode, "local")
		}
		mode = "local"
	}
	c.WO1HZMaFg6Mx = mode
	directory := c.INBdnGu56l5
	if directory == "false" {
		directory = ""
	}
	c.INBdnGu56l5 = strings.TrimRight(directory, "/")
	if c.PmIFLM && c.WO1HZMaFg6Mx == "clickhouse" && c.EuOaOO8qBn0Y > 0 {
		logx.V8GN5Stk(4, "access_log_retention_hours does not work in %v mode, please clean the access log table by yourself (e.g. ALTER TABLE ... MODIFY TTL)", "clickhouse")
	}
}

// mQ5mMKnkLjl 对应地址 0x132a880。警告文本保留原字节及顺序。
func mQ5mMKnkLjl() {
	old := &configAt509C940
	c := &configAt509CB80
	if old.H2FCRVc {
		if c.KCYNbwBsZ {
			logx.TRTbrquD(4, "redis_enable 等旧配置已废弃且被忽略，当前生效的是 redis_ip_limit_* 配置，请删除旧配置")
		} else {
			logx.TRTbrquD(4, "redis_enable/redis_addr/redis_password/redis_db/redis_tls/conn_limit_expiry 配置已废弃，即将在后续版本移除，已自动转换为 redis_ip_limit_* 配置，请尽快改用新配置")
			c.KCYNbwBsZ = true
			c.FXMKW6VcHH = old.Ij4bWLsuo_x4
			c.AhScWSQ = old.A9GLBE4PY
			c.ADewfG9Z64W = old.ZDYXh3Cuvxpw
			c.NDRQGsmc132Y = old.IDyeXiEv
			c.BFkqcX7Dx = old.Eb9vzdoE
		}
	}
	if c.BFkqcX7Dx <= 0 {
		c.BFkqcX7Dx = 60
	}
	if c.KCYNbwBsZ && c.FXMKW6VcHH == "" {
		logx.TRTbrquD(4, "redis_ip_limit_addr 未配置，全局限制用户 IP 数不会启用")
		c.KCYNbwBsZ = false
	}
}

// eifTpu4B 对应地址 0x1333100。
func eifTpu4B() {
	c := &configAt50A02C0
	if c.SVuGy3CoRze {
		c.dZWc0A2uIe = "aead"
		logx.V8GN5Stk(4, "force_vmess_aead is deprecated, please use vmess_auth_mode=aead instead")
	} else {
		c.dZWc0A2uIe = zaWmXx3rGK35(c.ZFK_OBT1)
	}
}
