// Original DB constructor @0x173fee0 and GetNode/GetAPIData methods.
package sspaneltrojan

import (
	"fmt"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	P "recovered.local/soga/panel/sspanel"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

type Lvx7xjWU struct{ P.N3owFvm }

func rOyGkf(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &Lvx7xjWU{P.N3owFvm{TMfNUn: c}} }
func init()                              { Q.PanelFactoriesAt5095E50["sspanel-uim-trojan-db"] = rOyGkf }
func (p *Lvx7xjWU) GetNode() (*L.Wr5SVP, error) {
	r, err := p.GetNodeResult()
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	return r.ToTrojanNode(p.TMfNUn.Yf1KqSqJpNI7)
}
func (p *Lvx7xjWU) GetAPIData() (*Q.D5gTIZ5aVg, error) {
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
