package vKl1IHr3

import "strings"

// DP0s5dasePo 对应 0xb7d4e0，0x607440 核定为 TrimSpace。
// 每项先去空白，空项跳过，最后只去除末尾一个分隔符。
func DP0s5dasePo(values []string, separator string) string {
	var b strings.Builder
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			b.WriteString(value)
			b.WriteString(separator)
		}
	}
	out := b.String()
	if len(out) != 0 && len(separator) != 0 {
		return out[:len(out)-len(separator)]
	}
	return out
}
