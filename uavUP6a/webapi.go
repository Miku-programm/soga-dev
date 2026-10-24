package uavUP6a

import (
	"fmt"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	P "recovered.local/soga/PqJzgkx7s6D8"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

type Ch9OFE struct{ P.JYkPXk4ZdW }

func m0djx6vX4eRo(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &Ch9OFE{P.JYkPXk4ZdW{EA_5pkz: c}} }
func init()                                    { Q.PanelFactoriesAt5095E50["sspanel-uim-v2ray-webapi"] = m0djx6vX4eRo }
func (p *Ch9OFE) GetNode() (*L.Wr5SVP, error) {
	r, err := p.GetNodeResult()
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	return r.ToVMessNode(p.EA_5pkz.Yf1KqSqJpNI7)
}
func (p *Ch9OFE) GetAPIData() (*Q.D5gTIZ5aVg, error) {
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
	if out.H47AzZsYHB1N != nil {
		for _, u := range out.VTme1SuHL6i_ {
			if v, ok := u.JAfXWwzYrW.(*L.Yn8BJoL); ok {
				v.FvQ2ltKOM7 = out.H47AzZsYHB1N.SwxJNgPpg.VUwf1n9g
			}
		}
	}
	return out, err
}
