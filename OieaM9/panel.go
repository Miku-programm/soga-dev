// 由原样本本包构造、请求、闭包及init.0恢复，文字由原decoder受限仿真验证。
package OieaM9

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

// OieaM9.init.0：原样本入口 0x16c2520。
func init() { Q.PanelFactoriesAt5095E50["proxypanel-ssr-webapi"] = vIsEGhbrfBJK }

// OieaM9.vIsEGhbrfBJK：原样本入口 0x16c25e0。
func vIsEGhbrfBJK(c *C.FCEAfQJwDA7e) Q.VyFk5uT { return &HbUGOX{GMLWI3: B.GMLWI3{TMfNUn: c}} }

// OieaM9.(*HbUGOX).aBxp9PQk：原样本入口 0x16c2660。
func (p *HbUGOX) aBxp9PQk(path string) string { return p.TMfNUn.W_ogGtF + "api/vnet/v2/" + path }

// OieaM9.(*HbUGOX).GetNode：原样本入口 0x16c2720。
func (p *HbUGOX) GetNode() (*L.Wr5SVP, error) {
	body, err := p.GetCommon(p.aBxp9PQk(fmt.Sprintf("node/%v", p.TMfNUn.Yf1KqSqJpNI7)))
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	var r TEsbPj5T
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("parse node failed: %v", err)
	}
	if r.IdNatD8MGczW != 200 {
		return nil, fmt.Errorf("get node failed, code: %v", r.IdNatD8MGczW)
	}
	return r.GPETCf2.ToPanelNode(p.TMfNUn.Yf1KqSqJpNI7)
}

// OieaM9.(*HbUGOX).w6XnzT：原样本入口 0x16c2b80。
func (p *HbUGOX) w6XnzT() (L.Sxcf3wu3L, error) {
	body, err := p.GetCommon(p.aBxp9PQk(fmt.Sprintf("userList/%v", p.TMfNUn.Yf1KqSqJpNI7)))
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	var r JpYAk64U2
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("parse users failed: %v", err)
	}
	if r.IdNatD8MGczW != 200 {
		return nil, fmt.Errorf("get users failed, code: %v", r.IdNatD8MGczW)
	}
	return r.GetUsers(), nil
}

// OieaM9.(*HbUGOX).z_ImEB0DAdLt：原样本入口 0x16c3020。
func (p *HbUGOX) z_ImEB0DAdLt() (X.Ij9r21E0UOt3, error) {
	return p.GetDetectors(p.aBxp9PQk(fmt.Sprintf("nodeRule/%v", p.TMfNUn.Yf1KqSqJpNI7)))
}

// OieaM9.(*HbUGOX).GetAPIData：原样本入口 0x16c3100。
func (p *HbUGOX) GetAPIData() (*Q.D5gTIZ5aVg, error) {
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
			v, e := p.w6XnzT()
			if e != nil {
				return e
			}
			out.VTme1SuHL6i_ = v
			return nil
		},
		func() error {
			v, e := p.z_ImEB0DAdLt()
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
		credential := &L.QcqpSL3i6ri{YrjVl_2g_S: n.ASJxBbFS.PU1Di18, YyAgnMatmj: n.ASJxBbFS.JvlcaS, Z_taAb: n.ASJxBbFS.Qggyovo, H_y2xZht7S8: n.ASJxBbFS.D2rWMm0, ZZIDsJ: n.ASJxBbFS.ND3AWUWg, BskzckkE7qI: 2}
		n.ASJxBbFS.AddSSRUser(&L.Aga91y{JAfXWwzYrW: credential})
	}
	return out, err
}

// OieaM9.(*HbUGOX).SubmitData：原样本入口 0x16c3400。
func (p *HbUGOX) SubmitData(data *Q.FGmNba4MhioY) error {
	id := p.TMfNUn.Yf1KqSqJpNI7
	return H.D5aDHlQDLp(
		func() error { return p.SubmitNodeInfo(p.aBxp9PQk(fmt.Sprintf("nodeStatus/%v", id))) },
		func() error { return p.SubmitTraffic(p.aBxp9PQk(fmt.Sprintf("userTraffic/%v", id)), data.HTgPyE3VytZd) },
		func() error { return p.SubmitAliveIP(p.aBxp9PQk(fmt.Sprintf("nodeOnline/%v", id)), data.JNcHHwTZ7) },
		func() error { return p.SubmitDetectLog(p.aBxp9PQk(fmt.Sprintf("trigger/%v", id)), data.Y0ys4ma4Tno4) })
}
