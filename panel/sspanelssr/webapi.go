package sspanelssr

import (
	"fmt"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	P "recovered.local/soga/panel/sspanel"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

type MpE9pa4fA struct{ P.JYkPXk4ZdW }

func tIgW4wJnP(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &MpE9pa4fA{P.JYkPXk4ZdW{EA_5pkz: c}} }
func init()                                 { Q.PanelFactoriesAt5095E50["sspanel-uim-ssr-webapi"] = tIgW4wJnP }
func (p *MpE9pa4fA) GetNode() (*L.Wr5SVP, error) {
	r, err := p.GetNodeResult()
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	if server, err := p.viBFPHIjmS5(); err == nil {
		r.L8xOe2JrpGDq.HIgDb6THe1A9 = server
	}
	return r.ToSSRNode(p.EA_5pkz.Yf1KqSqJpNI7)
}
func (p *MpE9pa4fA) GetAPIData() (*Q.D5gTIZ5aVg, error) {
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
	if err == nil {
		addSSRUsers(out)
	}
	return out, err
}
