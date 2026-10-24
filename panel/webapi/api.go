package webapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	G "recovered.local/soga/net/relay"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	W "recovered.local/soga/panel/etagcache"
	"strconv"
)

func WF3vigUr(c *C.FCEAfQJwDA7e) Q.VyFk5uT {
	return &Zkt5w_2o{iAFjt_tT9ZmG: c, oDI12LC: new(W.NG0hCZyRHK)}
}
func (p *Zkt5w_2o) GetAPIConfig() *C.FCEAfQJwDA7e { return p.iAFjt_tT9ZmG }
func (p *Zkt5w_2o) rDzO_k() string {
	switch p.iAFjt_tT9ZmG.ZyAhkq5th8bf {
	case "ss":
		return "shadowsocks"
	case "v2ray":
		return "vmess"
	case "vless":
		return "vless"
	case "trojan":
		return "trojan"
	case "hysteria":
		return "hysteria"
	}
	return ""
}
func (p *Zkt5w_2o) m15fwKapDO() url.Values {
	return url.Values{"protocol": {p.rDzO_k()}, "server_id": {strconv.FormatInt(int64(p.iAFjt_tT9ZmG.Yf1KqSqJpNI7), 10)}, "secret_key": {p.iAFjt_tT9ZmG.Bd08Li}}
}
func (p *Zkt5w_2o) qOheTtVSGU(address string) http.Header {
	etag := p.oDI12LC.GetETag(address)
	if etag == "" {
		return nil
	}
	return http.Header{"IF-NONE-MATCH": {etag}}
}
func (p *Zkt5w_2o) mTbgoUSgl(action string) string {
	return fmt.Sprintf("%vv1/server/%v", p.iAFjt_tT9ZmG.W_ogGtF, action)
}

// aCLea6Z 对应0x16ad120；同一 ETag 的 HTTP 200 仍返回旧 body，304 无缓存时报错。
func (p *Zkt5w_2o) aCLea6Z(action string) ([]byte, error) {
	address := p.mTbgoUSgl(action)
	headers := p.qOheTtVSGU(address)
	query := p.m15fwKapDO()
	status, body, response, err := G.IlWBbp1xdPw(address, query, headers)
	if err != nil {
		return nil, fmt.Errorf("%v, result: %v, status code: %v", err, string(body), status)
	}
	if status == 304 {
		if value, ok := p.oDI12LC.GetObj(address); ok {
			return value.([]byte), nil
		}
		return nil, fmt.Errorf("can not find eTag: %v", address)
	}
	if status != 200 {
		return nil, fmt.Errorf("request failed, result: %v, status code: %v", string(body), status)
	}
	if etag := response.Get("etag"); etag != "" {
		if value, ok := p.oDI12LC.CheckETag(address, etag); ok {
			return value.([]byte), nil
		}
		p.oDI12LC.AddETag(address, etag, body)
	}
	return body, nil
}
func (p *Zkt5w_2o) wVyldOdHZ3gi(action string, value any) error {
	address := p.mTbgoUSgl(action)
	query := p.m15fwKapDO()
	body, status, err := G.ZBovtj1xky(address, query, value, nil)
	if err != nil {
		return fmt.Errorf("error: %v, status code: %v, result: %v", err, status, string(body))
	}
	if status != 200 {
		return fmt.Errorf("statuc code: %v, result: %v", status, string(body))
	}
	return nil
}
func (p *Zkt5w_2o) zg3DBxs7Nv() (*Yq4D8wXbIQyY, error) {
	body, err := p.aCLea6Z("config")
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	config := new(Yq4D8wXbIQyY)
	if err = json.Unmarshal(body, config); err != nil {
		return nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(body))
	}
	return config, nil
}
func (p *Zkt5w_2o) n2CFbQ1Ws() (T8lmI8Ic, error) {
	body, err := p.aCLea6Z("user")
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	var response struct {
		Gd8mcBP80d T8lmI8Ic `json:"users"`
	}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(body))
	}
	return response.Gd8mcBP80d, nil
}
func (p *Zkt5w_2o) GetNode() (*L.Wr5SVP, error) {
	config, err := p.zg3DBxs7Nv()
	if err != nil {
		return nil, err
	}
	node, err := config.ToPanelNode(p.iAFjt_tT9ZmG.Yf1KqSqJpNI7)
	if err != nil {
		return nil, fmt.Errorf("parse node failed: %v", err)
	}
	return node, nil
}
func (p *Zkt5w_2o) GetAPIData() (*Q.D5gTIZ5aVg, error) {
	out := new(Q.D5gTIZ5aVg)
	err := H.D5aDHlQDLp(func() error {
		config, e := p.zg3DBxs7Nv()
		if e != nil {
			return e
		}
		out.Xdo1D9 = config.GetPullInterval()
		out.DDu2Nm_P = config.GetPushInterval()
		out.H47AzZsYHB1N, e = config.ToPanelNode(p.iAFjt_tT9ZmG.Yf1KqSqJpNI7)
		return e
	}, func() error {
		users, e := p.n2CFbQ1Ws()
		if e != nil {
			return e
		}
		out.VTme1SuHL6i_ = users.ToPanelUsers(p.iAFjt_tT9ZmG.ZyAhkq5th8bf)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if p.iAFjt_tT9ZmG.ZyAhkq5th8bf == "ss" {
		cipher := out.H47AzZsYHB1N.FnMUny.HMkc4BgR
		encoded := cipher == "2022-blake3-aes-128-gcm" || cipher == "2022-blake3-aes-256-gcm"
		keyLen, e := cipher.KeyLen()
		if e != nil {
			return nil, fmt.Errorf("ss node err: %v", e)
		}
		for _, user := range out.VTme1SuHL6i_ {
			if credential, ok := user.JAfXWwzYrW.(*L.YlfOV0DBBnx); ok {
				credential.T1oDDIgYJHh = cipher
				if encoded && keyLen <= len(credential.A_wNUTkMh6pH) {
					credential.A_wNUTkMh6pH = base64.StdEncoding.EncodeToString([]byte(credential.A_wNUTkMh6pH[:keyLen]))
				}
			}
		}
	}
	return out, nil
}
