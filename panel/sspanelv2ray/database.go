// Original DB constructor @0x1741960 and GetNode/GetAPIData methods.
package sspanelv2ray

import (
	"fmt"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	P "recovered.local/soga/panel/sspanel"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

type DFA2BtCSa7 struct{ P.N3owFvm }

func lApu0FSLw0r4(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &DFA2BtCSa7{P.N3owFvm{TMfNUn: c}} }
func init()                                    { Q.PanelFactoriesAt5095E50["sspanel-uim-v2ray-db"] = lApu0FSLw0r4 }
func (p *DFA2BtCSa7) GetNode() (*L.Wr5SVP, error) {
	r, err := p.GetNodeResult()
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	return r.ToVMessNode(p.TMfNUn.Yf1KqSqJpNI7)
}
func (p *DFA2BtCSa7) GetAPIData() (*Q.D5gTIZ5aVg, error) {
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
