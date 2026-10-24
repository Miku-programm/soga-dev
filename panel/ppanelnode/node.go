package ppanelnode

import (
	"errors"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	F "recovered.local/soga/panel/ssrnames"
	Q "recovered.local/soga/panel/registry"
)

func (p *UXWxgJe4slp) GetNode() (*L.Wr5SVP, error) { //0x1771d80
	c := C.WHMCSSettings()
	in := &KDNfxh{c.VD6q_75hbej2, c.Z9bFboOQlPiw, F.ZvAxOt1l7Ar(c.Et04j58qH), F.Mu0vMIIs(c.EKv8Zjg7dn85), F.PMorzD7u(c.HmKQa3xt)}
	if in.ImzyOwUUCbr == 0 {
		return nil, errors.New("please set whmcs_port")
	}
	if in.VeJdy0T6 == "" {
		return nil, errors.New("please set whmcs_password")
	}
	if in.SpgrmCJh1rgI == "" {
		return nil, errors.New("please set whmcs_method")
	}
	if in.WhAH8ORz == "" {
		return nil, errors.New("please set whmcs_protocol")
	}
	if in.AoHvzHNkflAi == "" {
		return nil, errors.New("please set whmcs_obfs")
	}
	return in.ToPanelNode(), nil
}
func (in *KDNfxh) ToPanelNode() *L.Wr5SVP { //0x1779ce0
	n := &L.Wr5SVP{Z1TIgHOQ: Q.T0j_Rco(in), HSA2b5V: C.PanelForbiddenBitTorrent(), TtWKyh: C.PanelNodeSpeedLimit(), W3Dw_B4wgbM: C.PanelUserConnectionLimit()}
	n.ASJxBbFS.PU1Di18 = in.ImzyOwUUCbr
	n.ASJxBbFS.Qggyovo = in.SpgrmCJh1rgI
	n.ASJxBbFS.D2rWMm0 = in.WhAH8ORz
	n.ASJxBbFS.ND3AWUWg = in.AoHvzHNkflAi
	n.ASJxBbFS.JvlcaS = in.VeJdy0T6
	n.ASJxBbFS.QawbKS2x = true
	n.ASJxBbFS.AddSSRUser(&L.Aga91y{JAfXWwzYrW: &L.QcqpSL3i6ri{YrjVl_2g_S: in.ImzyOwUUCbr, YyAgnMatmj: in.VeJdy0T6, Z_taAb: in.SpgrmCJh1rgI, H_y2xZht7S8: in.WhAH8ORz, ZZIDsJ: in.AoHvzHNkflAi, BskzckkE7qI: 2}})
	return n
}
