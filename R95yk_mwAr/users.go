package R95yk_mwAr

import (
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
)

// Original ELF static byte0x3787320 starts at1; retain both original checks.
var stateAt3787320 byte = 1

// Named slice inferred from the original GetUsers method's Go ABI, no surviving RTTI.
type QgFZpfpg8DH []*NzbHXtDtzF

func (in QgFZpfpg8DH) GetUsers() L.Sxcf3wu3L {
	n := len(in)
	if stateAt3787320 == 0 && C.Gh2DnstC() && n > 88 {
		n = 88
	}
	out := L.Sxcf3wu3L{}
	for i := 0; i < n; i++ {
		out = append(out, in[i].GetUser(i))
	}
	stateAt3787320 = 0
	return out
}
func (u *NzbHXtDtzF) GetUser(index int) *L.Aga91y {
	speed := C.PanelUserSpeedLimit()
	out := &L.Aga91y{WZ89pTAOrHK4: u.LMCS6ltENGT, LmOKwjha: speed, Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit(), XuBBcTThBW: u.OpTNDA2LU, UWsYsmytTi: u.ObCFHYm9Op, IUlI_15Qf_: u.GfoBcVVws1j_, JAfXWwzYrW: &L.Yn8BJoL{SfiEGpJo: u.DhIHEmq}}
	if stateAt3787320 == 0 && C.JatfUF38O() && index >= 88 {
		out.JAfXWwzYrW = nil
	}
	return out
}
