package _ZwYXGIdVc

import (
	"encoding/json"
	"fmt"
	G "recovered.local/soga/G6bHD5JA"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	X "recovered.local/soga/xx5fRyrFw"
)

func D_5dOciQK_A(config *C.FCEAfQJwDA7e) Q.VyFk5uT { //0x175bb00
	return &DrEoaRCt2{pBXis_w: config}
}
func (p *DrEoaRCt2) d5nKoaGxYfB(action string) (string, error) { //0x175bb80
	kind := ""
	switch p.pBXis_w.ZyAhkq5th8bf {
	case "ss":
		kind = "shadowsocks"
	case "v2ray":
		kind = "v2ray"
	case "trojan":
		kind = "trojan"
	default:
		return "", fmt.Errorf("unsupported server type: %v", p.pBXis_w.ZyAhkq5th8bf)
	}
	// Original interpolation does not add slash or escape token/action.
	return fmt.Sprintf("%vweb_api/?token=%v&node_id=%v&node_type=%v&act=%v", p.pBXis_w.W_ogGtF, p.pBXis_w.Bd08Li, p.pBXis_w.Yf1KqSqJpNI7, kind, action), nil
}

// hc5XlrePB5JY restores0x175bee0: HTTP status is not inspected.
func (p *DrEoaRCt2) hc5XlrePB5JY(action string) ([]byte, error) {
	address, err := p.d5nKoaGxYfB(action)
	if err != nil {
		return nil, err
	}
	status, data, _, err := G.IlWBbp1xdPw(address, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("result: %v, err: %v, status code: %v", string(data), err, status)
	}
	return data, nil
}
func (p *DrEoaRCt2) rQrimKF088a(action string, value any) error { //0x175c0e0
	address, err := p.d5nKoaGxYfB(action)
	if err != nil {
		return err
	}
	data, status, err := G.ZBovtj1xky(address, nil, value, nil)
	if err != nil {
		return fmt.Errorf("result: %v, err: %v, status code: %v", string(data), err, status)
	}
	return nil // Original ignores POST status code and JSON body on transport success.
}
func (p *DrEoaRCt2) wLjVU6() (*L.Wr5SVP, X.Ij9r21E0UOt3, error) { //0x175c300
	data, err := p.hc5XlrePB5JY("config")
	if err != nil {
		return nil, nil, fmt.Errorf("get node failed: %v", err)
	}
	r := new(AnlFZbNq)
	if err = json.Unmarshal(data, r); err != nil {
		return nil, nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(data))
	}
	node, err := r.ToPanelNode()
	if err != nil {
		return nil, nil, fmt.Errorf("convert node failed: %v, result: %v", err, string(data))
	}
	return node, r.GetDetectors(), nil
}
func (p *DrEoaRCt2) f6jVOJvD4y7() (L.Sxcf3wu3L, error) { //0x175c960
	data, err := p.hc5XlrePB5JY("user")
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	r := new(Oh1oNuTp)
	if err = json.Unmarshal(data, r); err != nil {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(data))
	}
	users, err := r.ToPanelUsers(p.pBXis_w.ZyAhkq5th8bf)
	if err != nil {
		return nil, fmt.Errorf("convert users failed: %v, result: %v", err, string(data))
	}
	return users, nil
}
func (p *DrEoaRCt2) GetAPIConfig() *C.FCEAfQJwDA7e { return p.pBXis_w } //0x175cdc0
func (p *DrEoaRCt2) GetAPIData() (*Q.D5gTIZ5aVg, error) { //0x175cde0, closures1769360/17692e0
	out := new(Q.D5gTIZ5aVg)
	err := H.D5aDHlQDLp(func() error {
		n, d, e := p.wLjVU6()
		if e != nil {
			return e
		}
		out.H47AzZsYHB1N = n
		out.HYnfYWa = d
		return nil
	}, func() error {
		u, e := p.f6jVOJvD4y7()
		if e != nil {
			return e
		}
		out.VTme1SuHL6i_ = u
		return nil
	})
	if err != nil {
		return nil, err
	}
	for _, u := range out.VTme1SuHL6i_ {
		if c, ok := u.JAfXWwzYrW.(*L.YlfOV0DBBnx); ok {
			c.T1oDDIgYJHh = out.H47AzZsYHB1N.FnMUny.HMkc4BgR
		}
	}
	return out, nil
}
