package ppanelnode

import (
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
)

// Original ELF static byte0x378731E starts at1; retain both original checks.
var stateAt378731E byte = 1

// Named slice inferred from the original GetUsers method's Go ABI, no surviving RTTI.
type CgPle0Z []*Ya7brDb

func (in CgPle0Z) GetUsers() L.Sxcf3wu3L {
	n := len(in)
	out := L.Sxcf3wu3L{}
	for i := 0; i < n; i++ {
		out = append(out, in[i].GetUser(i))
	}
	stateAt378731E = 0
	return out
}
func (u *Ya7brDb) GetUser(index int) *L.Aga91y {
	speed := C.PanelUserSpeedLimit()
	out := &L.Aga91y{WZ89pTAOrHK4: u.SlNaxp, LmOKwjha: speed, Xmrna8: C.PanelUserConnectionLimit(), B4DYSxuA31za: C.PanelUserTCPLimit(), XuBBcTThBW: u.DPB_Ss, UWsYsmytTi: u.ZI0nwO, IUlI_15Qf_: u.TC6zmSFvzw, 	JAfXWwzYrW: &L.QcqpSL3i6ri{YrjVl_2g_S: u.So23GIUGYxZ8, YyAgnMatmj: u.GUDL30zd}}
	return out
}
