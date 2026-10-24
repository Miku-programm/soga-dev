package cNT24rG4F

import (
	"recovered.local/soga/CAodbhZf"
	L "recovered.local/soga/muc279Mpa6S4"
	"strings"
)

// EdlWFAt 对应0x1329660。只有未经修剪的精确值all触发接口枚举；
// 普通列表逐项修剪但不删除空项，空配置或枚举失败返回单个空监听地址。
func EdlWFAt() []string {
	value := configAt509BEA0.VerfnYj
	if value == "" {
		return []string{""}
	}
	if value == "all" {
		ips, err := CAodbhZf.N4EF0n2()
		if err != nil {
			L.V8GN5Stk(8, "listen [all] failed: %v", err)
			return []string{""}
		}
		if len(ips) == 0 {
			L.TRTbrquD(4, "listen [all] failed: no listen ips found")
			return []string{""}
		}
		result := make([]string, 0, len(ips))
		for _, ip := range ips {
			result = append(result, ip.String())
		}
		return result
	}
	result := []string{}
	for _, part := range strings.Split(value, ",") {
		result = append(result, strings.TrimSpace(part))
	}
	return result
}
