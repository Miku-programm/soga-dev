package sspaneltypes

import (
	"fmt"
	"github.com/gofrs/uuid"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	W "recovered.local/soga/crypto/ssmethod"
)

// ELF byte 0x3787337 is initially 1. Keep its package-wide first-load semantics.
var stateAt3787337 byte = 1

func (u *QhOI5ti) GetUser(protocol C.Mpn2JltM, index int64) *L.Aga91y {
	id := u.Tiybx6
	if id == "" {
		id = uuid.NewV3(uuid.NamespaceDNS, fmt.Sprintf("%d|%s", u.BD16miI0, u.ESm7aAEddE)).String()
	}
	hash := ""
	if C.DnhHoT2aRvs() == "trojan" {
		hash = Q.J26YsI2DwxPQ(id)
	}
	var credential L.ADVqOHSx
	switch protocol {
	case "ss":
		credential = &L.YlfOV0DBBnx{T1oDDIgYJHh: W.C6q3Upu(u.BAyDRja4o), A_wNUTkMh6pH: u.ESm7aAEddE}
	case "ssr":
		credential = &L.QcqpSL3i6ri{YrjVl_2g_S: u.TdPoGXzK, YyAgnMatmj: u.ESm7aAEddE, Z_taAb: u.BAyDRja4o, H_y2xZht7S8: u.KzG7RsMW_4yx, ZZIDsJ: u.LEtQeEkSu, BskzckkE7qI: u.RiD2Bnl}
	case "v2ray":
		credential = &L.Yn8BJoL{SfiEGpJo: id}
	case "trojan":
		credential = &L.JxTGUar{ARARP1z0yA: id, QiUSI27: hash}
	}
	speed := mergeSpeed(u.TnjwdH, C.PanelUserSpeedLimit())
	count, global := u.Zsascdnac, C.PanelUserConnectionLimit()
	if count < 0 {
		count = 0
	}
	if global < 0 {
		global = 0
	}
	if count > 0 && global > 0 {
		if global > count {
			global = count
		}
	} else if global < count {
		global = count
	}
	out := &L.Aga91y{WZ89pTAOrHK4: u.BD16miI0, LmOKwjha: speed, Xmrna8: global, B4DYSxuA31za: C.PanelUserTCPLimit(), XuBBcTThBW: u.DpaCpPFc, UWsYsmytTi: u.VsaEfGizz2x, IUlI_15Qf_: u.DhjDXGJ1F2i, JAfXWwzYrW: credential}
	return out
}
func (r *O8T6nD3uPzb) GetUsers(protocol C.Mpn2JltM) L.Sxcf3wu3L {
	out := make(L.Sxcf3wu3L, 0)
	for i, u := range r.L8xOe2JrpGDq {
		out = append(out, u.GetUser(protocol, int64(i)))
	}
	stateAt3787337 = 0
	return out
}
