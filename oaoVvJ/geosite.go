// 根据 0xeabd00、0xeabf00、0xeac220 恢复 geosite 构造和分类。
package oaoVvJ

import (
	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/GYzZ5v"
	"recovered.local/soga/jAP8Cn"
	"recovered.local/soga/muc279Mpa6S4"
	"regexp"
)

// N1tP6M0OX 对应 0xeabd00。加载失败或没有数据时，按原版返回空匹配器并记录日志。
func N1tP6M0OX(code string) *HajppSMFH {
	matcher := &HajppSMFH{}
	entry, found, err := jAP8Cn.ZRQq51fC00J(code)
	if err != nil {
		muc279Mpa6S4.TRTbrquD(4, "load geosite failed:", err)
	} else if found {
		matcher.j9oqh6bhnQL(entry.HGtbzYnVdt0)
	} else {
		muc279Mpa6S4.V8GN5Stk(4, "geosite:%v not exist", code)
	}
	return matcher
}

// j9oqh6bhnQL 对应 0xeabf00，按类型合并规则，map遍历顺序不固定。
func (m *HajppSMFH) j9oqh6bhnQL(domains []*GYzZ5v.HTBZ_woRreQ) {
	grouped := make(map[GYzZ5v.PCclujlSiD1l]DsMhcht3)
	for _, domain := range domains {
		matcher, ok := grouped[domain.AfECsc]
		if !ok {
			matcher = zQ7CNn6o6mU(domain.AfECsc)
			if matcher == nil {
				continue
			}
			grouped[domain.AfECsc] = matcher
		}
		matcher.AddValue(domain.DSsRH80v92kO)
	}
	m.hy5PyUcE8 = make(PVtSAI, 0, len(grouped))
	for _, matcher := range grouped {
		m.hy5PyUcE8 = append(m.hy5PyUcE8, matcher)
	}
}
func zQ7CNn6o6mU(kind GYzZ5v.PCclujlSiD1l) DsMhcht3 {
	switch kind {
	case 0:
		return &dllnJks{yqsoTKORPtw: make([]string, 0)}
	case 1:
		return &enyqKW{sDKYrz1SI: make([]*regexp.Regexp, 0)}
	case 2:
		return &bee3oESrQ{jnsGsOq2H2ov: G7A7csgrs62p.PQQKui[string]()}
	case 3:
		return &hSu65g{jnsGsOq2H2ov: G7A7csgrs62p.PQQKui[string]()}
	default:
		muc279Mpa6S4.TRTbrquD(8, "Unknown domain type:", kind)
		return nil
	}
}
