package aRZNjEvm

import (
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
)

// 原 ELF 0x378733C 初始字节 1；初始值直接从 ELF 数据段读取。
// 地址命名由恢复工程指定。保留首次调用与后续调用的原检查差异。
var stateAt378733C byte = 1

// dwkvm4ffDDTc 对应 0x17a5480。浮点比较顺序来自汇编，保留 NaN 行为。
func (u *AOa7p4h) dwkvm4ffDDTc(serverType cNT24rG4F.Mpn2JltM, index int) *LFWTTXF7JR.Aga91y {
	var credential LFWTTXF7JR.ADVqOHSx
	switch serverType {
	case "ss":
		credential = &LFWTTXF7JR.YlfOV0DBBnx{A_wNUTkMh6pH: u.UH0HXa5O}
	case "v2ray":
		credential = &LFWTTXF7JR.Yn8BJoL{SfiEGpJo: u.UH0HXa5O}
	case "vless":
		credential = &LFWTTXF7JR.IImIwacp{Dh5iiakU: u.UH0HXa5O}
	case "anytls":
		credential = &LFWTTXF7JR.ZtMFlBTZXyD{CBWeRfrO1ahs: u.UH0HXa5O}
	case "hysteria":
		credential = &LFWTTXF7JR.WQM2tKA{CBWeRfrO1ahs: u.UH0HXa5O}
	case "trojan":
		credential = &LFWTTXF7JR.JxTGUar{ARARP1z0yA: u.UH0HXa5O}
	}
	globalSpeed := cNT24rG4F.PanelUserSpeedLimit()
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
	globalConnections := cNT24rG4F.PanelUserConnectionLimit()
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
	out := &LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: u.IDLJXX6F, LmOKwjha: globalSpeed, Xmrna8: connections, B4DYSxuA31za: cNT24rG4F.PanelUserTCPLimit(), JAfXWwzYrW: credential}
	if stateAt378733C == 0 && cNT24rG4F.Rs_wP3R() && index >= 88 {
		out.JAfXWwzYrW = nil
	}
	return out
}

// ToPanelUsers 对应 0x17a5900；保留原切片长度修改和末尾共享字节复位。
func (r *IMyQ5CxaikH) ToPanelUsers(serverType cNT24rG4F.Mpn2JltM) LFWTTXF7JR.Sxcf3wu3L {
	if stateAt378733C == 0 && cNT24rG4F.Q0SDIN() && len(r.Ta593Rq2lSa) > 88 {
		r.Ta593Rq2lSa = r.Ta593Rq2lSa[:88]
	}
	out := make(LFWTTXF7JR.Sxcf3wu3L, 0)
	for i, user := range r.Ta593Rq2lSa {
		out = append(out, user.dwkvm4ffDDTc(serverType, i))
	}
	stateAt378733C = 0
	return out
}
