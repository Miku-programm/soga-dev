package xiaov2board

import (
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
)

// 原 ELF 0x378733C 初始字节 1；初始值直接从 ELF 数据段读取。
// 地址命名由恢复工程指定。保留首次调用与后续调用的原检查差异。
var stateAt378733C byte = 1

// dwkvm4ffDDTc 对应 0x17a5480。浮点比较顺序来自汇编，保留 NaN 行为。
func (u *AOa7p4h) dwkvm4ffDDTc(serverType config.Mpn2JltM, index int) *types.Aga91y {
	var credential types.ADVqOHSx
	switch serverType {
	case "ss":
		credential = &types.YlfOV0DBBnx{A_wNUTkMh6pH: u.UH0HXa5O}
	case "v2ray":
		credential = &types.Yn8BJoL{SfiEGpJo: u.UH0HXa5O}
	case "vless":
		credential = &types.IImIwacp{Dh5iiakU: u.UH0HXa5O}
	case "anytls":
		credential = &types.ZtMFlBTZXyD{CBWeRfrO1ahs: u.UH0HXa5O}
	case "hysteria":
		credential = &types.WQM2tKA{CBWeRfrO1ahs: u.UH0HXa5O}
	case "trojan":
		credential = &types.JxTGUar{ARARP1z0yA: u.UH0HXa5O}
	}
	globalSpeed := config.PanelUserSpeedLimit()
	if globalSpeed < 0 {
		globalSpeed = 0
	} else if globalSpeed > 10000 {
		globalSpeed = 10000
	}
	speed := u.DwdlKBQpcgY
	if speed < 0 {
		speed = 0
	} else if speed > 10000 {
		speed = 10000
	}
	if globalSpeed > 0 && speed > 0 {
		if !(speed > globalSpeed) {
			globalSpeed = speed
		}
	} else if !(globalSpeed > speed) {
		globalSpeed = speed
	}
	globalConnections := config.PanelUserConnectionLimit()
	connections := u.QeuUtZ8f
	if globalConnections < 0 {
		globalConnections = 0
	}
	if connections < 0 {
		connections = 0
	}
	if globalConnections > 0 && connections > 0 {
		if connections > globalConnections {
			connections = globalConnections
		}
	} else if connections < globalConnections {
		connections = globalConnections
	}
	out := &types.Aga91y{WZ89pTAOrHK4: u.IDLJXX6F, LmOKwjha: globalSpeed, Xmrna8: connections, B4DYSxuA31za: config.PanelUserTCPLimit(), JAfXWwzYrW: credential}
	if stateAt378733C == 0 && config.Rs_wP3R() && index >= 88 {
		out.JAfXWwzYrW = nil
	}
	return out
}

// ToPanelUsers 对应 0x17a5900；保留原切片长度修改和末尾共享字节复位。
func (r *IMyQ5CxaikH) ToPanelUsers(serverType config.Mpn2JltM) types.Sxcf3wu3L {
	if stateAt378733C == 0 && config.Q0SDIN() && len(r.Ta593Rq2lSa) > 88 {
		r.Ta593Rq2lSa = r.Ta593Rq2lSa[:88]
	}
	out := make(types.Sxcf3wu3L, 0)
	for i, user := range r.Ta593Rq2lSa {
		out = append(out, user.dwkvm4ffDDTc(serverType, i))
	}
	stateAt378733C = 0
	return out
}
