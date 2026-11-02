package xiaov2boardnode

import (
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
)

// Original ELF static byte0x3787320 starts at1; retain both original checks.
var stateAt3787320 byte = 1

// Named slice inferred from the original GetUsers method's Go ABI, no surviving RTTI.
type QgFZpfpg8DH []*NzbHXtDtzF

func (in QgFZpfpg8DH) GetUsers() L.Sxcf3wu3L {
	n := len(in)
	out := L.Sxcf3wu3L{}
	for i := 0; i < n; i++ {
		out = append(out, in[i].GetUser(i))
	}
	stateAt3787320 = 0
	return out
}
func (u *NzbHXtDtzF) GetUser(index int) *L.Aga91y {
	speed := C.PanelUserSpeedLimit()
	out := &L.Aga91y{WZ89pTAOrHK4: u.LMCS6ltENGT, LmOKwjha: speed, Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit(), XuBBcTThBW: u.OpTNDA2LU, UWsYsmytTi: u.ObCFHYm9Op, IUlI_15Qf_: u.GfoBcVVws1j_, 	JAfXWwzYrW: &L.Yn8BJoL{SfiEGpJo: u.DhIHEmq}}
	return out
}
