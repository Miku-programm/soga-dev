package xboardnode

import (
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
)

// Original ELF static byte0x378731F starts at1; retain both original checks.
var stateAt378731F byte = 1

// Named slice inferred from the original GetUsers method's Go ABI, no surviving RTTI.
type YQ5jWLXE []*M3DtrvpLfJH6

func (in YQ5jWLXE) GetUsers() L.Sxcf3wu3L {
	n := len(in)
	out := L.Sxcf3wu3L{}
	for i := 0; i < n; i++ {
		out = append(out, in[i].GetUser(i))
	}
	stateAt378731F = 0
	return out
}
func (u *M3DtrvpLfJH6) GetUser(index int) *L.Aga91y {
	speed := C.PanelUserSpeedLimit()
	out := &L.Aga91y{WZ89pTAOrHK4: u.LMCS6ltENGT, LmOKwjha: speed, Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit(), XuBBcTThBW: u.OpTNDA2LU, UWsYsmytTi: u.ObCFHYm9Op, IUlI_15Qf_: u.GfoBcVVws1j_, 	JAfXWwzYrW: &L.JxTGUar{ARARP1z0yA: u.DhIHEmq}}
	return out
}
