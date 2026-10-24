package JlGfHYMcZ

import (
	"errors"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

func (p *Ckybjbgxz3Y) GetNode() (*L.Wr5SVP, error) { //0x177ae60
	c := C.WHMCSSettings()
	in := &TvejtjZJ{c.VD6q_75hbej2}
	if in.CcjL4eyDez == 0 {
		return nil, errors.New("please set whmcs_port")
	}
	return in.ToPanelNode(), nil
}
func (in *TvejtjZJ) ToPanelNode() *L.Wr5SVP { //0x1781060
	n := &L.Wr5SVP{Z1TIgHOQ: Q.T0j_Rco(in), HSA2b5V: C.PanelForbiddenBitTorrent(), TtWKyh: C.PanelNodeSpeedLimit(), W3Dw_B4wgbM: C.PanelUserConnectionLimit()}
	n.JhIYwEukjd.GbD_2PU = in.CcjL4eyDez
	n.JhIYwEukjd.T4cuhG7MTbG = "tcp"
	return n
}
