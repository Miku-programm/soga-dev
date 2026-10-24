package VWxNG7

import (
	"encoding/json"
	"fmt"
	"recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	D "recovered.local/soga/xx5fRyrFw"
)

// GetNode：0x16d8d40；间隔方法虽存在，但原聚合闭包不写入它们。
func (p *CsKk5VzGb) GetNode() (*L.Wr5SVP, error) {
	data, err := p.sXWeuE("node")
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	var n GX5_vLs
	if err = json.Unmarshal(data, &n); err != nil {
		return nil, fmt.Errorf("parse node failed: %v, data: %v", err, string(data))
	}
	node, err := n.ToPanelNode(p.gAaW6Z7Uh3.Yf1KqSqJpNI7)
	if err != nil {
		return nil, fmt.Errorf("parse node failed: %v, data: %v", err, string(data))
	}
	return node, nil
}
func (p *CsKk5VzGb) jB_aWXym() (L.Sxcf3wu3L, error) {
	data, err := p.sXWeuE("users")
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	var result DdDyLf6D
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("parse users failed: %v, data: %v", err, string(data))
	}
	return result.ToPanelUsers(p.gAaW6Z7Uh3.ZyAhkq5th8bf), nil
}
func (p *CsKk5VzGb) fyDKlia() (D.Ij9r21E0UOt3, error) {
	data, err := p.sXWeuE("audit_rules")
	if err != nil {
		return nil, fmt.Errorf("get rules failed: %v", err)
	}
	var result AQ3wazHBAA
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("parse rules failed: %v, data: %v", err, string(data))
	}
	return result.ToDetectors(), nil
}
func (p *CsKk5VzGb) jkZU_d4h() (D.VHWq49iRat, error) {
	data, err := p.sXWeuE("white_list")
	if err != nil {
		return nil, fmt.Errorf("get white list failed: %v", err)
	}
	var result D.VHWq49iRat
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("parse white list failed: %v, data: %v", err, string(data))
	}
	return result, nil
}

// GetAPIData：0x16d9d80，四个闭包：0x16e6e40/6dc0/6d40/6cc0。
func (p *CsKk5VzGb) GetAPIData() (*Q.D5gTIZ5aVg, error) {
	out := new(Q.D5gTIZ5aVg)
	err := HxKViEp2a.D5aDHlQDLp(
		func() error {
			n, e := p.GetNode()
			if e == nil {
				out.H47AzZsYHB1N = n
			}
			return e
		},
		func() error {
			u, e := p.jB_aWXym()
			if e == nil {
				out.VTme1SuHL6i_ = u
			}
			return e
		},
		func() error {
			r, e := p.fyDKlia()
			if e == nil {
				out.HYnfYWa = r
			}
			return e
		},
		func() error {
			w, e := p.jkZU_d4h()
			if e == nil {
				out.NjjWbjH = w
			}
			return e
		},
	)
	if err != nil {
		return nil, err
	}
	for _, u := range out.VTme1SuHL6i_ {
		switch c := u.JAfXWwzYrW.(type) {
		case *L.YlfOV0DBBnx:
			c.T1oDDIgYJHh = out.H47AzZsYHB1N.FnMUny.HMkc4BgR
		case *L.QcqpSL3i6ri:
			n := &out.H47AzZsYHB1N.ASJxBbFS
			c.YrjVl_2g_S = n.PU1Di18
			c.Z_taAb = n.Qggyovo
			c.ZZIDsJ = n.ND3AWUWg
			c.H_y2xZht7S8 = n.D2rWMm0
		}
	}
	return out, nil
}
