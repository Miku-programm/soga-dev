// 由原样本本包构造、请求、闭包及init.0恢复，文字由原decoder受限仿真验证。
package sTqn2l9

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

// sTqn2l9.init.0：原样本入口 0x16c92a0。
func init() { Q.PanelFactoriesAt5095E50["proxypanel-trojan-webapi"] = uZ2az_6UM }

// sTqn2l9.uZ2az_6UM：原样本入口 0x16c9360。
func uZ2az_6UM(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &J5tD2c9{GMLWI3: B.GMLWI3{TMfNUn: c}} }

// sTqn2l9.(*J5tD2c9).i5lLCOPR7yXe：原样本入口 0x16c93e0。
func (p *J5tD2c9) i5lLCOPR7yXe(path string) string { return p.TMfNUn.W_ogGtF + "api/trojan/v1/" + path }

// sTqn2l9.(*J5tD2c9).GetNode：原样本入口 0x16c94a0。
func (p *J5tD2c9) GetNode() (*L.Wr5SVP, error) {
	body, err := p.GetCommon(p.i5lLCOPR7yXe(fmt.Sprintf("node/%v", p.TMfNUn.Yf1KqSqJpNI7)))
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	var r WmjoeO
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("parse node failed: %v", err)
	}
	if r.IdNatD8MGczW != 200 {
		return nil, fmt.Errorf("get node failed, code: %v", r.IdNatD8MGczW)
	}
	return r.GPETCf2.ToPanelNode(p.TMfNUn.Yf1KqSqJpNI7), nil
}

// sTqn2l9.(*J5tD2c9).eJPdTSQz4uN：原样本入口 0x16c9920。
func (p *J5tD2c9) eJPdTSQz4uN() (L.Sxcf3wu3L, error) {
	body, err := p.GetCommon(p.i5lLCOPR7yXe(fmt.Sprintf("userList/%v", p.TMfNUn.Yf1KqSqJpNI7)))
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	var r BoM67db
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("parse users failed: %v", err)
	}
	if r.IdNatD8MGczW != 200 {
		return nil, fmt.Errorf("get users failed, code: %v", r.IdNatD8MGczW)
	}
	return r.GetUsers(), nil
}

// sTqn2l9.(*J5tD2c9).b0d4GGxgk：原样本入口 0x16c9dc0。
func (p *J5tD2c9) b0d4GGxgk() (X.Ij9r21E0UOt3, error) {
	return p.GetDetectors(p.i5lLCOPR7yXe(fmt.Sprintf("nodeRule/%v", p.TMfNUn.Yf1KqSqJpNI7)))
}

// sTqn2l9.(*J5tD2c9).GetAPIData：原样本入口 0x16c9ea0。
func (p *J5tD2c9) GetAPIData() (*Q.D5gTIZ5aVg, error) {
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
			v, e := p.eJPdTSQz4uN()
			if e != nil {
				return e
			}
			out.VTme1SuHL6i_ = v
			return nil
		},
		func() error {
			v, e := p.b0d4GGxgk()
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
		}
	}
	return out, err
}

// sTqn2l9.(*J5tD2c9).SubmitData：原样本入口 0x16ca060。
func (p *J5tD2c9) SubmitData(data *Q.FGmNba4MhioY) error {
	id := p.TMfNUn.Yf1KqSqJpNI7
	return H.D5aDHlQDLp(
		func() error { return p.SubmitNodeInfo(p.i5lLCOPR7yXe(fmt.Sprintf("nodeStatus/%v", id))) },
		func() error {
			return p.SubmitTraffic(p.i5lLCOPR7yXe(fmt.Sprintf("userTraffic/%v", id)), data.HTgPyE3VytZd)
		},
		func() error { return p.SubmitAliveIP(p.i5lLCOPR7yXe(fmt.Sprintf("nodeOnline/%v", id)), data.JNcHHwTZ7) },
		func() error {
			return p.SubmitDetectLog(p.i5lLCOPR7yXe(fmt.Sprintf("trigger/%v", id)), data.Y0ys4ma4Tno4)
		})
}
