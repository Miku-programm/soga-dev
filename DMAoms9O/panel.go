// 由原样本本包构造、请求、闭包及init.0恢复，文字由原decoder受限仿真验证。
package DMAoms9O

import (
	"encoding/json"
	"fmt"
	B "recovered.local/soga/D43Y3mB"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	X "recovered.local/soga/xx5fRyrFw"
)

// DMAoms9O.init.0：原样本入口 0x16cf400。
func init() { Q.PanelFactoriesAt5095E50["proxypanel-v2ray-webapi"] = gHV0TDBY }

// DMAoms9O.gHV0TDBY：原样本入口 0x16cf4c0。
func gHV0TDBY(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &UqRMl2W{GMLWI3: B.GMLWI3{TMfNUn: c}} }

// DMAoms9O.(*UqRMl2W).vnFrlhlVa1EL：原样本入口 0x16cf540。
func (p *UqRMl2W) vnFrlhlVa1EL(path string) string { return p.TMfNUn.W_ogGtF + "api/v2ray/v1/" + path }

// DMAoms9O.(*UqRMl2W).GetNode：原样本入口 0x16cf600。
func (p *UqRMl2W) GetNode() (*L.Wr5SVP, error) {
	body, err := p.GetCommon(p.vnFrlhlVa1EL(fmt.Sprintf("node/%v", p.TMfNUn.Yf1KqSqJpNI7)))
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	var r YEaGcD
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("parse node failed: %v", err)
	}
	if r.IdNatD8MGczW != 200 {
		return nil, fmt.Errorf("get node failed, code: %v", r.IdNatD8MGczW)
	}
	return r.GPETCf2.ToPanelNode(p.TMfNUn.Yf1KqSqJpNI7), nil
}

// DMAoms9O.(*UqRMl2W).hkSE81：原样本入口 0x16cfa60。
func (p *UqRMl2W) hkSE81() (L.Sxcf3wu3L, error) {
	body, err := p.GetCommon(p.vnFrlhlVa1EL(fmt.Sprintf("userList/%v", p.TMfNUn.Yf1KqSqJpNI7)))
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	var r FaNycS0m
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("parse users failed: %v", err)
	}
	if r.IdNatD8MGczW != 200 {
		return nil, fmt.Errorf("get users failed, code: %v", r.IdNatD8MGczW)
	}
	return r.GetUsers(), nil
}

// DMAoms9O.(*UqRMl2W).evR9LWMz：原样本入口 0x16cff20。
func (p *UqRMl2W) evR9LWMz() (X.Ij9r21E0UOt3, error) {
	return p.GetDetectors(p.vnFrlhlVa1EL(fmt.Sprintf("nodeRule/%v", p.TMfNUn.Yf1KqSqJpNI7)))
}

// DMAoms9O.(*UqRMl2W).GetAPIData：原样本入口 0x16d0000。
func (p *UqRMl2W) GetAPIData() (*Q.D5gTIZ5aVg, error) {
	out := new(Q.D5gTIZ5aVg)
	err := H.D5aDHlQDLp(
		func() error {
			v, e := p.GetNode()
			if e != nil {
				return e
			}
			out.H47AzZsYHB1N = v
			return nil
		},
		func() error {
			v, e := p.hkSE81()
			if e != nil {
				return e
			}
			out.VTme1SuHL6i_ = v
			return nil
		},
		func() error {
			v, e := p.evR9LWMz()
			if e != nil {
				return e
			}
			out.HYnfYWa = v
			return nil
		})
	// 原行为即使并行任务部分失败，仍按已取得节点同步用户限制。
	if n := out.H47AzZsYHB1N; n != nil {
		for _, u := range out.VTme1SuHL6i_ {
			u.Xmrna8 = n.W3Dw_B4wgbM
			if credential, ok := u.JAfXWwzYrW.(*L.Yn8BJoL); ok {
				credential.FvQ2ltKOM7 = n.SwxJNgPpg.VUwf1n9g
			}
		}
	}
	return out, err
}

// DMAoms9O.(*UqRMl2W).SubmitData：原样本入口 0x16d01e0。
func (p *UqRMl2W) SubmitData(data *Q.FGmNba4MhioY) error {
	id := p.TMfNUn.Yf1KqSqJpNI7
	return H.D5aDHlQDLp(
		func() error { return p.SubmitNodeInfo(p.vnFrlhlVa1EL(fmt.Sprintf("nodeStatus/%v", id))) },
		func() error {
			return p.SubmitTraffic(p.vnFrlhlVa1EL(fmt.Sprintf("userTraffic/%v", id)), data.HTgPyE3VytZd)
		},
		func() error { return p.SubmitAliveIP(p.vnFrlhlVa1EL(fmt.Sprintf("nodeOnline/%v", id)), data.JNcHHwTZ7) },
		func() error {
			return p.SubmitDetectLog(p.vnFrlhlVa1EL(fmt.Sprintf("trigger/%v", id)), data.Y0ys4ma4Tno4)
		})
}
