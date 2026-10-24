// Original DB constructor @0x173fee0 and GetNode/GetAPIData methods.
package Se1QrdYD7Aa

import (
	"fmt"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	P "recovered.local/soga/PqJzgkx7s6D8"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
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
