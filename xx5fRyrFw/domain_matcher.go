// 根据样本 RTTI 及 0xebb540—0xebc240、0xebe940—0xebefa8 指令恢复。
// 这是二进制的 Go 源码表示，不是作者源码文件原文。
package xx5fRyrFw

import (
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/oaoVvJ"
	"regexp"
	"strings"
	"sync"
)

type VkyfPk_JtMAm interface{ Match(string) bool }
type TaBfSSyPGYr interface{ MatchPacket([]byte) bool }
type lF6FBapn func(string) VkyfPk_JtMAm
type Q2jO4SS []string
type Ij9r21E0UOt3 []*IXZNQqqXDsV

// IXZNQqqXDsV 对应 RTTI 0x36c8c48，amd64 大小 72 字节。
type IXZNQqqXDsV struct {
	_fXIaVa   int
	qwkN4ar_B string
	wg8voMH   VkyfPk_JtMAm
	d9LcIO    sync.Once
	iTu75U    TaBfSSyPGYr
}
type cgjsz8_osTn3 struct{}
type oKUqB8kA struct{}
type lw3P3pi struct{ byZuFBprBR string }
type ajX96fwy struct{ t0DAgBajt string }
type lW2GR2MgrH struct{ dunzUGo string }
type nP9kahl struct{ dsOTghDJS *oaoVvJ.HajppSMFH }

// hGv42jgZ.OlsdK3aZO2 已对应到标准库 regexp.Regexp。
type geXMKdJw struct {
	uUU6ClDeF0 string
	uwcSJEzXa  *regexp.Regexp
	ahz2UWm    sync.Once
}

// 原全局地址 0x5095db0；变量名称由分析者指定。
// geosite:、keyword: 由原 init 的两处字符串 decoder 解出。
var factoriesAt5095DB0 = map[string]lF6FBapn{
	"regexp:":  func(rule string) VkyfPk_JtMAm { return &geXMKdJw{uUU6ClDeF0: rule[7:]} },
	"domain:":  func(rule string) VkyfPk_JtMAm { return &ajX96fwy{t0DAgBajt: rule[7:]} },
	"full:":    func(rule string) VkyfPk_JtMAm { return &lW2GR2MgrH{dunzUGo: rule[5:]} },
	"geosite:": func(rule string) VkyfPk_JtMAm { return &nP9kahl{dsOTghDJS: oaoVvJ.N1tP6M0OX(rule[8:])} },
	"keyword:": func(rule string) VkyfPk_JtMAm { return &lw3P3pi{byZuFBprBR: rule[8:]} },
}

// NewDNSMatcher 是为源码组织新增的导出辅助，不是原 metadata 函数。
// 精确表示 DNS 构造 o4RFaBT0MwY3@0x1586860 内联的私有字段初始化。
// 原 DNS 空规则预置全匹配对象；一般 ToDetectors 的空规则则不匹配。
func NewDNSMatcher(pattern string) *IXZNQqqXDsV {
	if pattern == "" {
		return &IXZNQqqXDsV{wg8voMH: &oKUqB8kA{}}
	}
	return &IXZNQqqXDsV{qwkN4ar_B: pattern}
}

// NewRouteMatcher 为 0x1b87308 附近的跨包私有字段初始化提供源码桥接。
// 与 DNS 的空规则处理不同，这里仅设置原始规则，编号保持零值。
func NewRouteMatcher(pattern string) *IXZNQqqXDsV {
	return &IXZNQqqXDsV{qwkN4ar_B: pattern}
}

// ToDetectors 对应 0xebb540，空输入返回非 nil 空切片，规则 ID 是原下标。
func (rules Q2jO4SS) ToDetectors() Ij9r21E0UOt3 {
	out := make(Ij9r21E0UOt3, 0, len(rules))
	for i, rule := range rules {
		out = append(out, &IXZNQqqXDsV{_fXIaVa: i, qwkN4ar_B: rule})
	}
	return out
}

// WUWGex 对应 0xebb7a0，识别已知规则前缀或 ip:/geoip:/port:。
func WUWGex(rule string) bool {
	if i := strings.IndexByte(rule, ':'); i >= 0 {
		if _, ok := factoriesAt5095DB0[rule[:i+1]]; ok {
			return true
		}
	}
	return strings.HasPrefix(rule, "ip:") || strings.HasPrefix(rule, "geoip:") || strings.HasPrefix(rule, "port:")
}
func (r *IXZNQqqXDsV) GetId() int      { return r._fXIaVa }
func (r *IXZNQqqXDsV) GetRule() string { return r.qwkN4ar_B }

// mLjyfL5UDf 对应 0xebb940；闭包的完整控制流来自 0xebed20.asm。
func (r *IXZNQqqXDsV) mLjyfL5UDf() {
	r.d9LcIO.Do(func() {
		if i := strings.IndexByte(r.qwkN4ar_B, ':'); i >= 0 {
			if factory, ok := factoriesAt5095DB0[r.qwkN4ar_B[:i+1]]; ok {
				r.wg8voMH = factory(r.qwkN4ar_B)
			}
		}
		if r.wg8voMH == nil {
			switch {
			case strings.HasPrefix(r.qwkN4ar_B, "ip:"), strings.HasPrefix(r.qwkN4ar_B, "geoip:"), r.qwkN4ar_B == "":
				r.wg8voMH = &cgjsz8_osTn3{}
			default:
				r.wg8voMH = &lw3P3pi{byZuFBprBR: r.qwkN4ar_B}
			}
		}
		if packet, ok := r.wg8voMH.(TaBfSSyPGYr); ok {
			r.iTu75U = packet
		}
	})
}
func (r *IXZNQqqXDsV) Test(value string) bool { r.mLjyfL5UDf(); return r.wg8voMH.Match(value) }
func (r *IXZNQqqXDsV) MatchPacket(packet []byte) bool {
	r.mLjyfL5UDf()
	return r.iTu75U != nil && r.iTu75U.MatchPacket(packet)
}
func (rules Ij9r21E0UOt3) Test(value string) (*IXZNQqqXDsV, bool) {
	for _, r := range rules {
		if r.Test(value) {
			return r, true
		}
	}
	return nil, false
}
func (rules Ij9r21E0UOt3) MatchPacket(packet []byte) (*IXZNQqqXDsV, bool) {
	for _, r := range rules {
		if r.MatchPacket(packet) {
			return r, true
		}
	}
	return nil, false
}
func (*cgjsz8_osTn3) Match(string) bool    { return false }
func (*oKUqB8kA) Match(string) bool        { return true }
func (r *lw3P3pi) Match(value string) bool { return strings.Contains(value, r.byZuFBprBR) }
func (r *ajX96fwy) Match(value string) bool {
	return strings.HasSuffix(value, r.t0DAgBajt) && (len(value) == len(r.t0DAgBajt) || value[len(value)-len(r.t0DAgBajt)-1] == '.')
}
func (r *lW2GR2MgrH) Match(value string) bool { return value == r.dunzUGo }
func (r *nP9kahl) Match(value string) bool    { return r.dsOTghDJS.Match(value) }

// x5z41sAeR 对应 0xebbd20。正则编译失败只记录日志，不回退普通字符串。
func (r *geXMKdJw) x5z41sAeR() {
	compiled, err := regexp.Compile(r.uUU6ClDeF0)
	if err != nil {
		muc279Mpa6S4.V8GN5Stk(8, "regex %s compile error: %s", r.uUU6ClDeF0, err)
		return
	}
	r.uwcSJEzXa = compiled
}
func (r *geXMKdJw) Match(value string) bool {
	r.ahz2UWm.Do(r.x5z41sAeR)
	return r.uwcSJEzXa != nil && r.uwcSJEzXa.MatchString(value)
}

// 原字节分支调用 Find 后判断非 nil，保留该行为而非用文本分支替换。
func (r *geXMKdJw) MatchPacket(packet []byte) bool {
	r.ahz2UWm.Do(r.x5z41sAeR)
	return r.uwcSJEzXa != nil && r.uwcSJEzXa.Find(packet) != nil
}
