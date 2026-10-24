// Original DB constructor @0x173b6c0 and GetNode/GetAPIData methods.
package sspanelssr

import (
	"fmt"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	P "recovered.local/soga/panel/sspanel"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

type Bnz_6dIDxvXS struct{ P.N3owFvm }

func iHVZeKHxxPT(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &Bnz_6dIDxvXS{P.N3owFvm{TMfNUn: c}} }
func init()                                   { Q.PanelFactoriesAt5095E50["sspanel-uim-ssr-db"] = iHVZeKHxxPT }
func (p *Bnz_6dIDxvXS) GetNode() (*L.Wr5SVP, error) {
	r, err := p.GetNodeResult()
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	return r.ToSSRNode(p.TMfNUn.Yf1KqSqJpNI7)
}
func (p *Bnz_6dIDxvXS) GetAPIData() (*Q.D5gTIZ5aVg, error) {
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
