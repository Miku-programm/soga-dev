// 根据原样本 oaoVvJ 匹配器指令及 RTTI 恢复，不是作者源码原文。
// RTTI 类型前缀 giWev7LTQDa，uncommon 包路径为 oaoVvJ。
package geosite

import (
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/util/logx"
	"regexp"
	"strings"
)

type DsMhcht3 interface {
	AddValue(string)
	Match(string) bool
}
type PVtSAI []DsMhcht3
type HajppSMFH struct{ hy5PyUcE8 PVtSAI }
type dllnJks struct{ yqsoTKORPtw []string }
type enyqKW struct{ sDKYrz1SI []*regexp.Regexp }
type bee3oESrQ struct {
	jnsGsOq2H2ov *genmap.Fw_lgKgam_jg[string]
}
type hSu65g struct {
	jnsGsOq2H2ov *genmap.Fw_lgKgam_jg[string]
}

// AddValue/Match 对应 0xeab5c0、0xeab6c0，普通规则按子串匹配。
func (m *dllnJks) AddValue(value string) { m.yqsoTKORPtw = append(m.yqsoTKORPtw, value) }
func (m *dllnJks) Match(value string) bool {
	for _, rule := range m.yqsoTKORPtw {
		if strings.Contains(value, rule) {
			return true
		}
	}
	return false
}

// 0xeab780、0xeab920；无效正则不加入切片。
func (m *enyqKW) AddValue(value string) {
	compiled, err := regexp.Compile(value)
	if err != nil {
		logx.TRTbrquD(8, "regex error:", value)
		return
	}
	m.sDKYrz1SI = append(m.sDKYrz1SI, compiled)
}
func (m *enyqKW) Match(value string) bool {
	for _, rule := range m.sDKYrz1SI {
		if rule.MatchString(value) {
			return true
		}
	}
	return false
}

// 0xeaba00、0xeaba60；按 rune 遍历，只在点之后继续查询完整后缀。
func (m *bee3oESrQ) AddValue(value string) { m.jnsGsOq2H2ov.Add(value) }
func (m *bee3oESrQ) Match(value string) bool {
	if m.jnsGsOq2H2ov.Has(value) {
		return true
	}
	for i, r := range value {
		if r == '.' && m.jnsGsOq2H2ov.Has(value[i+1:]) {
			return true
		}
	}
	return false
}

// 0xeabba0、0xeabc00；完整字符串匹配。
func (m *hSu65g) AddValue(value string)   { m.jnsGsOq2H2ov.Add(value) }
func (m *hSu65g) Match(value string) bool { return m.jnsGsOq2H2ov.Has(value) }

// Match 对应 0xeabc60；与上层 nP9kahl.Match 的内联逻辑一致。
func (m *HajppSMFH) Match(value string) bool {
	for _, matcher := range m.hy5PyUcE8 {
		if matcher.Match(value) {
			return true
		}
	}
	return false
}
