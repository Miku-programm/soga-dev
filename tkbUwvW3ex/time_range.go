// 本文件依据 IDA 指令和 Go 类型元数据恢复，非作者文件原文。
// 样本 SHA256：6660dcfbc7971188dcb751b440f4269aa718491b569d4c8629b751e6e339f92a。
// 类型名字符串使用 pTVEKuLf 前缀，方法及 uncommon 包路径使用 tkbUwvW3ex。
// 此处使用后者组织可编译包，不代表已确定作者的原始目录名。
package tkbUwvW3ex

import (
	"strconv"
	"time"

	"recovered.local/soga/vKl1IHr3"
)

// D1Teq5S1f 对应类型描述符 0x36b2240；字段名、顺序和标签来自 RTTI。
type D1Teq5S1f struct {
	LXcSWg26p int `json:"min_hour"`
	RBNUY7G   int `json:"min_minute"`
	BZwEA15n  int `json:"max_hour"`
	G1LUbU    int `json:"max_minute"`
}

// ITMfakhJ 对应类型描述符 0x35eabd8；元素描述符为 0x3426a48。
type ITMfakhJ []*D1Teq5S1f

// DJr_rBzh 对应 0xec1380。反向时段交换两端，不拆分成跨午夜时段。
// 原整数解析包装器的错误文本不参与此函数结果，故使用标准库解析。
func DJr_rBzh(text string) (*D1Teq5S1f, bool) {
	ends := vKl1IHr3.DAaBs7_C(text, "-", 2)
	if len(ends) != 2 {
		return nil, false
	}
	start := vKl1IHr3.DAaBs7_C(ends[0], ":", 2)
	end := vKl1IHr3.DAaBs7_C(ends[1], ":", 2)
	if len(start) != 2 || len(end) != 2 {
		return nil, false
	}
	h1, err := strconv.Atoi(start[0])
	if err != nil {
		return nil, false
	}
	m1, err := strconv.Atoi(start[1])
	if err != nil || h1 < 0 || h1 > 24 || m1 < 0 || m1 >= 60 || h1 == 24 && m1 != 0 {
		return nil, false
	}
	h2, err := strconv.Atoi(end[0])
	if err != nil {
		return nil, false
	}
	m2, err := strconv.Atoi(end[1])
	if err != nil || h2 < 0 || h2 > 24 || m2 < 0 || m2 >= 60 || h2 == 24 && m2 != 0 {
		return nil, false
	}
	if h2 < h1 || h2 == h1 && m2 < m1 {
		h1, h2 = h2, h1
		m1, m2 = m2, m1
	}
	return &D1Teq5S1f{h1, m1, h2, m2}, true
}

// IsInTime 对应 0xec1600，按传入时间所在时区取时、分，左闭右开。
// 原参数为混淆的 time.Time 类型；这里使用标准库的公开名称。
func (r *D1Teq5S1f) IsInTime(t time.Time) bool {
	hour, minute := t.Hour(), t.Minute()
	if r.LXcSWg26p > hour || r.LXcSWg26p == hour && r.RBNUY7G > minute {
		return false
	}
	return r.BZwEA15n > hour || r.BZwEA15n == hour && r.G1LUbU > minute
}

// HasCross 对应 0xec1760。保留按起点包含判断的行为，包括零长度区间。
func (r *D1Teq5S1f) HasCross(other *D1Teq5S1f) bool {
	hour, minute := other.LXcSWg26p, other.RBNUY7G
	if (r.LXcSWg26p < hour || r.LXcSWg26p == hour && r.RBNUY7G <= minute) &&
		(r.BZwEA15n > hour || r.BZwEA15n == hour && r.G1LUbU > minute) {
		return true
	}
	hour, minute = r.LXcSWg26p, r.RBNUY7G
	return (other.LXcSWg26p < hour || other.LXcSWg26p == hour && other.RBNUY7G <= minute) &&
		(other.BZwEA15n > hour || other.BZwEA15n == hour && other.G1LUbU > minute)
}

// IsInTime 对应 0xec1e00；空列表表示不限制时段。
func (ranges ITMfakhJ) IsInTime(t time.Time) bool {
	if len(ranges) == 0 {
		return true
	}
	for _, r := range ranges {
		if r.IsInTime(t) {
			return true
		}
	}
	return false
}

// HasCross 对应 0xec1f00；空接收者返回真，非空接收者与空参数返回假。
func (ranges ITMfakhJ) HasCross(other ITMfakhJ) bool {
	if len(ranges) == 0 {
		return true
	}
	for _, r := range ranges {
		for _, candidate := range other {
			if r.HasCross(candidate) {
				return true
			}
		}
	}
	return false
}
