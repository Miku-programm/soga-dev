package cNT24rG4F

import (
	"net"
	"strconv"
	"strings"
	"time"
)

// ZNxaPWNlDqa 是原 RTTI 的 96 字节类型。显示前缀 ftLBcUpmkQ 与
// uncommon package path cNT24rG4F 不同；这里采用后者，不声称是原文件路径。
type ZNxaPWNlDqa struct {
	ASvYULtJ7   string
	ZH4VO_H4JZ  string
	IJFBj1lIX   string
	H0x74Y      string
	SNWlSe      string
	BaLySDum3K  int
	EVFQh90cvJK bool
}

// FullTable: 0x131cd00。保留原始反引号拼接，无额外转义。
func (c *ZNxaPWNlDqa) FullTable() string { return "`" + c.ZH4VO_H4JZ + "`.`" + c.IJFBj1lIX + "`" }

// E7kAu2lOVtXY: 0x131cd80。读取实际配置全局，每次返回独立配置值。
func E7kAu2lOVtXY() *ZNxaPWNlDqa {
	c := &configAt50A2260
	// 原编译器内联分支把精确的 false 文本视为空值。
	clean := func(s string) string {
		if s == "false" {
			return ""
		}
		return s
	}
	r := &ZNxaPWNlDqa{ASvYULtJ7: bgj94i7et(clean(c.QpeDOsW)), ZH4VO_H4JZ: clean(c.CmPnvU), IJFBj1lIX: clean(c.G8aycKMf5aa), H0x74Y: clean(c.HnvQ_1vV), SNWlSe: clean(c.ACWNbHb1C), BaLySDum3K: c.DOZx3XRAm01, EVFQh90cvJK: c.D4WJQARA}
	if r.ZH4VO_H4JZ == "" {
		r.ZH4VO_H4JZ = "default"
	}
	if r.IJFBj1lIX == "" {
		r.IJFBj1lIX = "access_log"
	}
	if r.BaLySDum3K <= 0 {
		r.BaLySDum3K = 1000
	}
	return r
}

// bgj94i7et: 0x131d020。保留原始大小写敏感的 scheme、SplitHostPort 失败后
// 直接拼接默认端口的行为，包括原代码对带路径和已加括号 IPv6 的处理。
func bgj94i7et(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	s = strings.TrimRight(s, "/")
	scheme := "http://"
	if strings.HasPrefix(s, "http://") {
		s = s[7:]
	} else if strings.HasPrefix(s, "https://") {
		scheme = "https://"
		s = s[8:]
	}
	if s == "" {
		return ""
	}
	host, port, err := net.SplitHostPort(s)
	if err != nil {
		port = "8123"
		if scheme == "https://" {
			port = "8443"
		}
		return scheme + net.JoinHostPort(s, port)
	}
	if _, err = strconv.Atoi(port); err != nil {
		return ""
	}
	return scheme + net.JoinHostPort(host, port)
}

// O__f2DDogBIK: 0x131d4c0。按记录时间所在时区逐小时选文件。
func O__f2DDogBIK(t time.Time) string {
	return "soga-access-" + t.Format("2006-01-02 15") + ":00:00.log"
}
