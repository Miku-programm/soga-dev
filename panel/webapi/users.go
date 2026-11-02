package webapi

import (
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
)

// 原 ELF 数据字节 0x3787334 初值 1；保留首轮和后续轮次的原限制分支。
var stateAt3787334 byte = 1

func (u *MACjYTsGk7G) ToPanelUser(protocol C.Mpn2JltM) *L.Aga91y {
	var credential L.ADVqOHSx
	switch protocol {
	case "ss":
		credential = &L.YlfOV0DBBnx{A_wNUTkMh6pH: u.WtNHeT4}
	case "v2ray":
		credential = &L.Yn8BJoL{SfiEGpJo: u.WtNHeT4}
	case "vless":
		credential = &L.IImIwacp{Dh5iiakU: u.WtNHeT4}
	case "trojan":
		credential = &L.JxTGUar{ARARP1z0yA: u.WtNHeT4}
	case "hysteria":
		credential = &L.WQM2tKA{CBWeRfrO1ahs: u.WtNHeT4}
	}
	speed := u.YQbaJusq0g6N
	if speed < 0 {
		speed = 0
	} else if speed > 10000 {
		speed = 10000
	}
	global := C.PanelUserSpeedLimit()
	if global < 0 {
		global = 0
	} else if global > 10000 {
		global = 10000
	}
	if speed > 0 && global > 0 {
		if !(global > speed) {
			speed = global
		}
	} else if !(speed > global) {
		speed = global
	}
	count, globalCount := u.D7H3u6wL2q5Y, C.PanelUserConnectionLimit()
	if count < 0 {
		count = 0
	}
	if globalCount < 0 {
		globalCount = 0
	}
	if count > 0 && globalCount > 0 {
		if globalCount > count {
			globalCount = count
		}
	} else if globalCount < count {
		globalCount = count
	}
	return &L.Aga91y{WZ89pTAOrHK4: u.AnbiaLml0Bs, LmOKwjha: speed, Xmrna8: globalCount, B4DYSxuA31za: C.PanelUserTCPLimit(), JAfXWwzYrW: credential}
}
func (users T8lmI8Ic) ToPanelUsers(protocol C.Mpn2JltM) L.Sxcf3wu3L {
	out := make(L.Sxcf3wu3L, 0, len(users))
	for _, u := range users {
		out = append(out, u.ToPanelUser(protocol))
	}
	stateAt3787334 = 0
	return out
}
