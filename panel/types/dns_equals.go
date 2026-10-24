package types

import "slices"

// Equals 对应 0x12741a0。切片按顺序比较内容，nil 与空切片相等；
// 原函数未对 nil 结构指针做保护，因此不增加特殊分支。
func (r *Axa8uKIB4JL) Equals(other *Axa8uKIB4JL) bool {
	return slices.Equal(r.K8BNfziIgb, other.K8BNfziIgb) && r.CM0eMsqyP == other.CM0eMsqyP && slices.Equal(r.CZ80ik, other.CZ80ik)
}

// Equals 对应 0x1274360。
func (r U863Mm01Nre) Equals(other U863Mm01Nre) bool {
	if len(r) != len(other) {
		return false
	}
	for i := range r {
		if !r[i].Equals(other[i]) {
			return false
		}
	}
	return true
}
