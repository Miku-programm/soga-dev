// SSPanel-UIM WebAPI：按原 0x172b420—0x172cf40 与 RTTI 恢复。
package sspanel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	G "recovered.local/soga/net/relay"
	L "recovered.local/soga/panel/types"
	U "recovered.local/soga/panel/sspaneltypes"
	C "recovered.local/soga/config"
	W "recovered.local/soga/panel/etagcache"
	X "recovered.local/soga/audit/matcher"
	"strconv"
)

// 原 RTTI 56 字节，缓存是值字段，不是指针。
type JYkPXk4ZdW struct {
	EA_5pkz *C.FCEAfQJwDA7e
	eyq32s  W.NG0hCZyRHK
}

func (p *JYkPXk4ZdW) GetAPIConfig() *C.FCEAfQJwDA7e { return p.EA_5pkz }
func (p *JYkPXk4ZdW) _fIxSv() url.Values {
	v := make(url.Values)
	v.Add("key", p.EA_5pkz.Bd08Li)
	v.Add("node_id", strconv.FormatInt(int64(p.EA_5pkz.Yf1KqSqJpNI7), 10))
	return v
}
func (p *JYkPXk4ZdW) dtazxgZ4i(address string) http.Header {
	etag := p.eyq32s.GetETag(address)
	if etag == "" {
		return nil
	}
	return http.Header{"IF-NONE-MATCH": {etag}}
}
func (p *JYkPXk4ZdW) GetCommon(action string) ([]byte, error) {
	address := p.EA_5pkz.W_ogGtF + "mod_mu/" + action
	query := p._fIxSv()
	headers := p.dtazxgZ4i(address)
	status, body, response, err := G.IlWBbp1xdPw(address, query, headers)
	if err != nil {
		return nil, fmt.Errorf("%v, result: %v, status code: %v", err, string(body), status)
	}
	if status == 304 {
		if obj, ok := p.eyq32s.GetObj(address); ok {
			return obj.([]byte), nil
		}
		return nil, fmt.Errorf("can not find eTag: %v", address)
	}
	// 原函数对非 304 的 HTTP 状态不作拒绝。
	if etag := response.Get("etag"); etag != "" {
		if obj, ok := p.eyq32s.CheckETag(address, etag); ok {
			return obj.([]byte), nil
		}
		p.eyq32s.AddETag(address, etag, body)
	}
	return body, nil
}
func (p *JYkPXk4ZdW) PostCommon(action string, value any) error {
	address := p.EA_5pkz.W_ogGtF + "mod_mu/" + action
	body, status, err := G.ZBovtj1xky(address, p._fIxSv(), value, nil)
	if err != nil {
		return fmt.Errorf("%v, result: %v, status code: %v", err, string(body), status)
	}
	if status != 200 {
		return fmt.Errorf("request failed, result: %v, status code: %v", string(body), status)
	}
	var result U.BqC2_84zPQi
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.LTQEC_KXaAS2 != 1 {
		return fmt.Errorf("result: %v", string(body))
	}
	return nil
}
func (p *JYkPXk4ZdW) GetNodeResult() (*U.H0izB2agvG, error) {
	body, err := p.GetCommon(fmt.Sprintf("nodes/%v/info", p.EA_5pkz.Yf1KqSqJpNI7))
	if err != nil {
		return nil, fmt.Errorf("get node failed: %v", err)
	}
	out := new(U.H0izB2agvG)
	if err = json.Unmarshal(body, out); err != nil {
		return nil, fmt.Errorf("parse node failed: %v, result: %v", err, string(body))
	}
	return out, nil
}
func (p *JYkPXk4ZdW) GetUsersResult() (*U.O8T6nD3uPzb, error) {
	body, err := p.GetCommon("users")
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	out := new(U.O8T6nD3uPzb)
	if err = json.Unmarshal(body, out); err != nil {
		return nil, fmt.Errorf("parse users failed: %v, result: %v", err, string(body))
	}
	return out, nil
}
func (p *JYkPXk4ZdW) GetDetectRulesResult() (*U.Nq0XWDbiQ, error) {
	body, err := p.GetCommon("func/detect_rules")
	if err != nil {
		return nil, fmt.Errorf("get detect rules failed: %v", err)
	}
	out := new(U.Nq0XWDbiQ)
	if err = json.Unmarshal(body, out); err != nil {
		return nil, fmt.Errorf("parse detect rules failed: %v, result: %v", err, string(body))
	}
	return out, nil
}
func (p *JYkPXk4ZdW) GetUsers() (L.Sxcf3wu3L, error) {
	result, err := p.GetUsersResult()
	if err != nil {
		return nil, fmt.Errorf("get users failed: %v", err)
	}
	return result.GetUsers(p.EA_5pkz.ZyAhkq5th8bf), nil
}
func (p *JYkPXk4ZdW) GetDetectors() (X.Ij9r21E0UOt3, error) {
	result, err := p.GetDetectRulesResult()
	if err != nil {
		return nil, fmt.Errorf("get detect rules failed: %v", err)
	}
	return result.ToDetectors(), nil
}
