package R95yk_mwAr

import (
	"errors"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

func (p *MrOlpocy) GetNode() (*L.Wr5SVP, error) { //0x1781c60
	c := C.WHMCSSettings()
	in := &GBljZRVUbG0m{c.VD6q_75hbej2, c.X4Dd6nOZ_7, c.Dd5Xza, c.W2WFtBWP, c.NQ3CpRlonU, c.Ew45bIap}
	if in.Kp1UgWbcxUg == 0 {
		return nil, errors.New("please set whmcs_port")
	}
	return in.ToPanelNode(), nil
}
func (in *GBljZRVUbG0m) ToPanelNode() *L.Wr5SVP { //0x1787fa0
	n := &L.Wr5SVP{Z1TIgHOQ: Q.T0j_Rco(in), HSA2b5V: C.PanelForbiddenBitTorrent(), TtWKyh: C.PanelNodeSpeedLimit(), W3Dw_B4wgbM: C.PanelUserConnectionLimit()}
	n.SwxJNgPpg.CKsOpayd = in.Kp1UgWbcxUg
	n.SwxJNgPpg.VUwf1n9g = in.CnPfr4P
	n.SwxJNgPpg.F2GXyZZ = L.BoulTn31(in.F2WQsJbM0cXY)
	n.SwxJNgPpg.UVeahqDBT = in.GY6Wi2
	n.SwxJNgPpg.ChJnngzb = in.WE23i9V_
	n.SwxJNgPpg.KWffzRDF = in.G5G15a
	return n
}
