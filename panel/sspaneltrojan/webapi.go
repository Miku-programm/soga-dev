package sspaneltrojan

import (
	"fmt"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	P "recovered.local/soga/panel/sspanel"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

type FU_j8iQ struct{ P.JYkPXk4ZdW }

func aoRZHdant(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &FU_j8iQ{P.JYkPXk4ZdW{EA_5pkz: c}} }
func init()                                 { Q.PanelFactoriesAt5095E50["sspanel-uim-trojan-webapi"] = aoRZHdant }
func (p *FU_j8iQ) GetNode() (*L.Wr5SVP, error) {
	r, err := p.GetNodeResult()
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	return r.ToTrojanNode(p.EA_5pkz.Yf1KqSqJpNI7)
}
func (p *FU_j8iQ) GetAPIData() (*Q.D5gTIZ5aVg, error) {
	out := new(Q.D5gTIZ5aVg)
	err := H.D5aDHlQDLp(func() error {
		v, e := p.GetNode()
		if e == nil {
			out.H47AzZsYHB1N = v
		}
		return e
	}, func() error {
		v, e := p.GetUsers()
		if e == nil {
			out.VTme1SuHL6i_ = v
		}
		return e
	}, func() error {
		v, e := p.GetDetectors()
		if e == nil {
			out.HYnfYWa = v
		}
		return e
	})
	return out, err
}
