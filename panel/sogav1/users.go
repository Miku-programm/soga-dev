// 根据原样本 0x16d71e0、0x16d76e0、0x16d5a40 恢复。
package sogav1

import (
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	D "recovered.local/soga/audit/matcher"
)

// 比较顺序保留原指令的 NaN 分支，不用 math.Min/Max 替代。
func mergeSpeed(a, b float64) float64 {
	if a < 0 {
		a = 0
	} else if a > 10000 {
		a = 10000
	}
	if b < 0 {
		b = 0
	} else if b > 10000 {
		b = 10000
	}
	if a > 0 && b > 0 {
		if !(b > a) {
			a = b
		}
	} else if !(a > b) {
		a = b
	}
	return a
}
func mergeCount(a, b int) int {
	if a < 0 {
		a = 0
	}
	if b < 0 {
		b = 0
	}
	if a > 0 && b > 0 {
		if b < a {
			a = b
		}
	} else if a < b {
		a = b
	}
	return a
}
func (u *D6uXyl88ab7) ToPanelUser(kind C.Mpn2JltM) *L.Aga91y {
	var credential L.ADVqOHSx
	switch kind {
	case "ss":
		credential = &L.YlfOV0DBBnx{A_wNUTkMh6pH: u.DHknGz}
	case "ssr":
		credential = &L.QcqpSL3i6ri{YyAgnMatmj: u.DHknGz}
	case "v2ray":
		credential = &L.Yn8BJoL{SfiEGpJo: u.S7USjTWE8o}
	case "vless":
		credential = &L.IImIwacp{Dh5iiakU: u.S7USjTWE8o}
	case "trojan":
		credential = &L.JxTGUar{ARARP1z0yA: u.DHknGz}
	case "hysteria":
		credential = &L.WQM2tKA{CBWeRfrO1ahs: u.DHknGz}
	case "anytls":
		credential = &L.ZtMFlBTZXyD{CBWeRfrO1ahs: u.DHknGz}
	case "mieru":
		credential = &L.DO9qqM9{Ejujyaj: u.WUNMcwtiDw, GltkJGvV: u.DHknGz}
	}
	return &L.Aga91y{WZ89pTAOrHK4: u.IMizGzpEPYfj, LmOKwjha: mergeSpeed(u.GZJoKAlg8K9, C.PanelUserSpeedLimit()), Xmrna8: mergeCount(C.PanelUserConnectionLimit(), u.LNDBhsiH), B4DYSxuA31za: mergeCount(C.PanelUserTCPLimit(), u.DGYgUabBB6Jx), JAfXWwzYrW: credential}
}

// 原 ELF 0x3787335 初值为 1；保留原授权状态下的用户数量限制。
var stateAt3787335 byte = 1

func (users DdDyLf6D) ToPanelUsers(kind C.Mpn2JltM) L.Sxcf3wu3L {
	if stateAt3787335 == 0 && C.Q0SDIN() && len(users) > 88 {
		users = users[:88]
	}
	out := make(L.Sxcf3wu3L, 0, len(users))
	for _, u := range users {
		out = append(out, u.ToPanelUser(kind))
	}
	stateAt3787335 = 0
	return out
}
func (rules AQ3wazHBAA) ToDetectors() D.Ij9r21E0UOt3 {
	out := make(D.Ij9r21E0UOt3, 0, len(rules))
	for _, r := range rules {
		out = append(out, D.NewPanelDetector(r.RFcywH9K, r.FI29zO1tSx))
	}
	return out
}
func (r *OFFj2pY8Z) Success() bool { return r.NIg3b55LCOZ == 0 }
