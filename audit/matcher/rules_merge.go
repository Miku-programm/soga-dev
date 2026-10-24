package matcher

import "slices"

// Merge 对应 0xebb360；空接收者原样返回 other，其余只过滤追加项中的已存在值。
func (rules VHWq49iRat) Merge(other VHWq49iRat) VHWq49iRat {
	if len(rules) == 0 {
		return other
	}
	for _, rule := range other {
		if !slices.Contains(rules, rule) {
			rules = append(rules, rule)
		}
	}
	return rules
}
