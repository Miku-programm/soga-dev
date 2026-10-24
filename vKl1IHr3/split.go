// 本文件由样本的 IDA 指令恢复，非作者源码原文。
package vKl1IHr3

import "strings"

// DAaBs7_C 对应 0xb7d720：分割前后去除空白，再丢弃空项。
func DAaBs7_C(text, separator string, n int) []string {
	parts := strings.SplitN(strings.TrimSpace(text), separator, n)
	result := make([]string, 0)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}
