package uazGa0Ap

import (
	"encoding/json"
	"fmt"
	"net/http"
	G "recovered.local/soga/G6bHD5JA"
	H "recovered.local/soga/HxKViEp2a"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	W "recovered.local/soga/wd42zQ0FT"
	X "recovered.local/soga/xx5fRyrFw"
)

func JLydOCi(config *C.FCEAfQJwDA7e) Q.VyFk5uT { //0x174fb80
	return &W3tZSymoah{iAFjt_tT9ZmG: config, oDI12LC: new(W.NG0hCZyRHK)}
}
func (p *W3tZSymoah) adwTxBv(action string) (string, error) { //0x174fc40
	kind := ""
	switch p.iAFjt_tT9ZmG.ZyAhkq5th8bf {
	case "ss":
		kind = "shadowsocks"
	case "v2ray", "vless":
		kind = "v2ray"
	case "trojan":
		kind = "trojan"
	case "hysteria":
		kind = "hysteria2"
	default:
		return "", fmt.Errorf("unsupported server type: %v", p.iAFjt_tT9ZmG.ZyAhkq5th8bf)
	}
	// Original interpolation does not add slash or escape token/action.
	return fmt.Sprintf("%vmodules/addons/v2raysocks_nodes/web_api.php/?token=%v&node_id=%v&node_type=%v&act=%v", p.iAFjt_tT9ZmG.W_ogGtF, p.iAFjt_tT9ZmG.Bd08Li, p.iAFjt_tT9ZmG.Yf1KqSqJpNI7, kind, action), nil
}
func (p *W3tZSymoah) eGEwNwIMv(address string) http.Header { //0x1750040
	tag := p.oDI12LC.GetETag(address)
	if tag == "" {
		return nil
	}
	return http.Header{"IF-NONE-MATCH": {tag}}
}
func (p *W3tZSymoah) mSSnouymyTV1(action string) ([]byte, error) { //0x17501c0
	address, err := p.adwTxBv(action)
	if err != nil {
		return nil, err
	}
	status, data, headers, err := G.IlWBbp1xdPw(address, nil, p.eGEwNwIMv(address))
	if err != nil {
		return nil, fmt.Errorf("result: %v, err: %v, status code: %v", string(data), err, status)
	}
	if status == 304 {
		obj, ok := p.oDI12LC.GetObj(address)
		if !ok {
			return nil, fmt.Errorf("can not find eTag: %v", address)
		}
		return obj.([]byte), nil
	}
	if status != 200 {
		return nil, fmt.Errorf("request failed, result: %v, status code: %v", string(data), status)
	}
	if tag := headers.Get("etag"); tag != "" {
		if old, ok := p.oDI12LC.CheckETag(address, tag); ok {
			return old.([]byte), nil
		}
		p.oDI12LC.AddETag(address, tag, data)
	}
	return data, nil
}
func (p *W3tZSymoah) cpDvcLHkAY(action string, value any) error { //0x17507e0
	address, err := p.adwTxBv(action)
	if err != nil {
		return err
	}
	data, status, err := G.ZBovtj1xky(address, nil, value, nil)
	if err != nil {
		return fmt.Errorf("result: %v, err: %v, status code: %v", string(data), err, status)
	}
	return nil // Original ignores POST status code and JSON body on transport success.
}
func (p *W3tZSymoah) qN4sSoabr() (*L.Wr5SVP, X.Ij9r21E0UOt3, error) { //0x17509e0
	data, err := p.mSSnouymyTV1("config")
	if err != nil {
		return nil, nil, fmt.Errorf("get node failed: %v", err)
	}
	r := new(C7MFTE4gD)
	if err = json.Unmarshal(data, r); err != nil {
		return nil, nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(data))
	}
	node, err := r.ToPanelNode(p.iAFjt_tT9ZmG.Yf1KqSqJpNI7)
	if err != nil {
		return nil, nil, fmt.Errorf("convert node failed: %v, result: %v", err, string(data))
	}
	return node, r.GetDetectors(), nil
}
func (p *W3tZSymoah) kctmoq() (L.Sxcf3wu3L, error) { //0x1750e80
	data, err := p.mSSnouymyTV1("user")
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	r := new(QbI1asEYYq)
	if err = json.Unmarshal(data, r); err != nil {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(data))
	}
	return r.ToPanelUsers(p.iAFjt_tT9ZmG.ZyAhkq5th8bf), nil // Original does not call IsOk here.
}
func (p *W3tZSymoah) GetAPIConfig() *C.FCEAfQJwDA7e { return p.iAFjt_tT9ZmG } //0x17511c0
func (p *W3tZSymoah) GetAPIData() (*Q.D5gTIZ5aVg, error) { //0x17511e0, closures175afe0/175af60
	out := new(Q.D5gTIZ5aVg)
	err := H.D5aDHlQDLp(func() error {
		n, d, e := p.qN4sSoabr()
		if e != nil {
			return e
		}
		out.H47AzZsYHB1N = n
		out.HYnfYWa = d
		return nil
	}, func() error {
		u, e := p.kctmoq()
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
