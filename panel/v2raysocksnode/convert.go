package v2raysocksnode

import (
	"errors"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	W "recovered.local/soga/crypto/ssmethod"
	X "recovered.local/soga/audit/matcher"
)

// ToPanelNode restores 0x174e480; select the first exact proxy tag, hash only
// that inbound, and preserve the original shared Reality slice backing arrays.
func (r *C7MFTE4gD) ToPanelNode(id int) (*L.Wr5SVP, error) {
	for _, in := range r.Dps5ZWL {
		if in.IouPnj9l != "proxy" {
			continue
		}
		n := &L.Wr5SVP{EMlXEc: id, Z1TIgHOQ: Q.T0j_Rco(in), HSA2b5V: C.PanelForbiddenBitTorrent(), TtWKyh: C.PanelNodeSpeedLimit(), W3Dw_B4wgbM: C.PanelUserConnectionLimit()}
		s := in.IHDmZMGzd
		n.SwxJNgPpg.CKsOpayd = in.CMxAlUxE
		n.SwxJNgPpg.F2GXyZZ = L.BoulTn31(s.BmgWKVKBTb5Y)
		n.SwxJNgPpg.UVeahqDBT = s.FlU0GQuhpmy == "tls"
		n.SwxJNgPpg.ChJnngzb = s.BPlxuJE.JCRhkrluQQY
		n.SwxJNgPpg.KWffzRDF = s.JhnfLInJR.RVKq_aLvfWyF
		n.JhIYwEukjd.GbD_2PU = in.CMxAlUxE
		n.JhIYwEukjd.T4cuhG7MTbG = L.BoulTn31(s.BmgWKVKBTb5Y)
		n.JhIYwEukjd.PBDb3Wc = s.BPlxuJE.JCRhkrluQQY
		n.JhIYwEukjd.QVlaSCZ4RCK = s.JhnfLInJR.RVKq_aLvfWyF
		n.FnMUny.Am3QoeF = in.CMxAlUxE
		n.FnMUny.HMkc4BgR = W.C6q3Upu(in.IyMnP39SHgC.P9ydbTayFu)
		n.FnMUny.VuhWqou = in.IyMnP39SHgC.XUfl9V
		n.Ck6DxjB.GJtOdKOWBiFP = in.CMxAlUxE
		n.Ck6DxjB.XmIaPj3CzOF = L.BoulTn31(s.BmgWKVKBTb5Y)
		n.Ck6DxjB.WAdJ4ftdZFk = L.D7FHIoaa8s4a(s.FlU0GQuhpmy)
		n.Ck6DxjB.Nd3mBux1 = s.BPlxuJE.JCRhkrluQQY
		n.Ck6DxjB.J1efyloFyhk = s.JhnfLInJR.RVKq_aLvfWyF
		if s.FlU0GQuhpmy == "reality" {
			n.Ck6DxjB.NMzMUQ2 = "xtls-rprx-vision"
		}
		n.Ck6DxjB.BANVIZ72wh = s.WmebQF3Gl7.DbrtlXUx
		n.Ck6DxjB.CrS6C1 = s.WmebQF3Gl7.HCxLjZtIa9f
		n.Ck6DxjB.DxKJ7lSe3m8F = s.WmebQF3Gl7.WuQfmyoetUwc
		n.Ck6DxjB.IqSfNzea = s.WmebQF3Gl7.FGJM_s
		return n, nil
	}
	return nil, errors.New("invalid node config: can not find \"proxy\" tag")
}

// GetDetectors restores 0x174ea60: non-nil empty, regardless of routing.
func (*C7MFTE4gD) GetDetectors() X.Ij9r21E0UOt3 { return X.Ij9r21E0UOt3{} }
func (r *QbI1asEYYq) IsOk() bool                { return r.DCbciSyJ_c == "ok" } //0x174f600

// Original ELF byte 0x378733a starts at1; no authorization branch is removed.
var stateAt378733A byte = 1

func (r *QbI1asEYYq) ToPanelUsers(kind C.Mpn2JltM) L.Sxcf3wu3L { //0x174f620
	out := make(L.Sxcf3wu3L, 0, len(r.XH32wVY))
	for i, u := range r.XH32wVY {
		out = append(out, u.ToPanelUser(kind, i))
	}
	stateAt378733A = 0
	return out
}
func (u *WsPqZpEZ) ToPanelUser(kind C.Mpn2JltM, index int) *L.Aga91y { //0x174f7e0
	var credential L.ADVqOHSx
	switch kind {
	case "ss":
		credential = &L.YlfOV0DBBnx{A_wNUTkMh6pH: u.PEtkkff}
	case "v2ray":
		credential = &L.Yn8BJoL{SfiEGpJo: u.CQ5RzKX}
	case "vless":
		credential = &L.IImIwacp{Dh5iiakU: u.CQ5RzKX}
	case "trojan":
		credential = &L.JxTGUar{ARARP1z0yA: u.JBz4bq}
	}
	globalSpeed, speed := C.PanelUserSpeedLimit(), u.TjNsBPv
	if globalSpeed < 0 {
		globalSpeed = 0
	} else if globalSpeed > 10000 {
		globalSpeed = 10000
	}
	if speed < 0 {
		speed = 0
	} else if speed > 10000 {
		speed = 10000
	}
	// UCOMISD/JA order preserves unordered NaN behavior from original ASM.
	if globalSpeed > 0 && speed > 0 {
		if !(speed > globalSpeed) {
			globalSpeed = speed
		}
	} else if !(globalSpeed > speed) {
		globalSpeed = speed
	}
	globalConnections, connections := C.PanelUserConnectionLimit(), u.U1cKBG6G62
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
	out := &L.Aga91y{WZ89pTAOrHK4: u.Gae2lLMaC4h, LmOKwjha: globalSpeed, Xmrna8: connections, B4DYSxuA31za: C.PanelUserTCPLimit(), JAfXWwzYrW: credential}
	return out
}
