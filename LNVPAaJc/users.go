package LNVPAaJc

import (
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
)

// Original ELF static byte0x378731D starts at1; retain both original checks.
var stateAt378731D byte = 1

// Named slice inferred from the original GetUsers method's Go ABI, no surviving RTTI.
type RFBSFd3 []*A4MyCjSpEhR

func (in RFBSFd3) GetUsers() L.Sxcf3wu3L {
	n := len(in)
	if stateAt378731D == 0 && C.JatfUF38O() && n > 88 {
		n = 88
	}
	out := L.Sxcf3wu3L{}
	for i := 0; i < n; i++ {
		out = append(out, in[i].GetUser(i))
	}
	stateAt378731D = 0
	return out
}
func (u *A4MyCjSpEhR) GetUser(index int) *L.Aga91y {
	speed := C.PanelUserSpeedLimit()
	local := float64(u.DZfKpi2g)
	if local < 0 {
		local = 0
	} else if local > 10000 {
		local = 10000
	}
	if speed < 0 {
		speed = 0
	} else if speed > 10000 {
		speed = 10000
	}
	if local > 0 && speed > 0 {
		if !(speed > local) {
			local = speed
		}
	} else if !(local > speed) {
		local = speed
	}
	speed = local
	out := &L.Aga91y{WZ89pTAOrHK4: u.UZaO5D8ajd1, LmOKwjha: speed, Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit(), XuBBcTThBW: u.BdFEnjMRK, UWsYsmytTi: u.FtW49iWwvF, IUlI_15Qf_: u.HXazhQBL, JAfXWwzYrW: &L.YlfOV0DBBnx{A_wNUTkMh6pH: u.SvKI66}}
	if stateAt378731D == 0 && C.Gh2DnstC() && index >= 88 {
		out.JAfXWwzYrW = nil
	}
	return out
}
