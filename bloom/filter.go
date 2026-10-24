// Package bloom 逐函数恢复样本所链接github.com/riobard/go-bloom；
// 本目录为恢复工程路径，未用未经匹配的公开库替换。
package bloom

import "math"

type Filter interface {
	Add([]byte)
	Reset()
	Size() int
	Test([]byte) bool
}
type classicFilter struct {
	b []byte
	k int
	h func([]byte) (uint64, uint64)
}

// New 对应0x1173e40。保留浮点计算顺序和向零截断，无新增参数校验。
func New(n int, p float64, h func([]byte) (uint64, uint64)) Filter {
	k := -math.Log(p) * 1.4426950408889634
	size := int(0.125 * (1.4426950408889634 * (float64(n) * k)))
	return &classicFilter{b: make([]byte, size), k: int(k), h: h}
}

// Add 对应0x1173f80，uint64哈希溢出后再按位数取余。
func (f *classicFilter) Add(data []byte) {
	a, b := f.h(data)
	for i := 0; i < f.k; i++ {
		offset := (a + b*uint64(i)) % uint64(8*len(f.b))
		f.b[offset>>3] |= 1 << (offset & 7)
	}
}

// Test 对应0x1174060。
func (f *classicFilter) Test(data []byte) bool {
	a, b := f.h(data)
	for i := 0; i < f.k; i++ {
		offset := (a + b*uint64(i)) % uint64(8*len(f.b))
		if f.b[offset>>3]&(1<<(offset&7)) == 0 {
			return false
		}
	}
	return true
}

// Size 对应0x1174160，返回字节数。
func (f *classicFilter) Size() int { return len(f.b) }

// Reset 对应0x1174180。
func (f *classicFilter) Reset() { clear(f.b) }
