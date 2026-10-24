// 依据 0x16d78a0—0x16d8d40 的控制流及受限解码字符串恢复。
package sogav1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recovered.local/soga/net/relay"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/registry"
	"recovered.local/soga/panel/etagcache"
	"strconv"
)

func DDky8V(c *config.FCEAfQJwDA7e) registry.VyFk5uT {
	return &CsKk5VzGb{cI03v8q7: new(etagcache.NG0hCZyRHK), gAaW6Z7Uh3: c}
}
func (p *CsKk5VzGb) GetAPIConfig() *config.FCEAfQJwDA7e { return p.gAaW6Z7Uh3 }
func (p *CsKk5VzGb) GetUrl(path string) string             { return p.gAaW6Z7Uh3.W_ogGtF + "api/v1/" + path }
func (p *CsKk5VzGb) rrcjzNnhs() string {
	switch p.gAaW6Z7Uh3.ZyAhkq5th8bf {
	case "ss":
		return "shadowsocks"
	case "ssr":
		return "shadowsocksr"
	case "v2ray":
		return "vmess"
	case "trojan", "vless", "anytls", "hysteria", "mieru":
		return string(p.gAaW6Z7Uh3.ZyAhkq5th8bf)
	default:
		return "unknown"
	}
}
func (p *CsKk5VzGb) ojEs5pzln(address string) http.Header {
	h := http.Header{"API-KEY": {p.gAaW6Z7Uh3.Bd08Li}, "NODE-ID": {strconv.Itoa(p.gAaW6Z7Uh3.Yf1KqSqJpNI7)}, "NODE-TYPE": {p.rrcjzNnhs()}}
	if etag := p.cI03v8q7.GetETag(address); etag != "" {
		h.Set("IF-NONE-MATCH", etag)
	}
	return h
}
func (p *CsKk5VzGb) olDPHF3j(path string, value any) error {
	address := p.GetUrl(path)
	data, status, err := relay.ZBovtj1xky(address, nil, value, p.ojEs5pzln(address))
	if err != nil {
		return fmt.Errorf("post %v err: %v, status code: %v, result: %v", address, err, status, string(data))
	}
	if status != 200 {
		return fmt.Errorf("post status code: %v, result: %v", status, string(data))
	}
	var result OFFj2pY8Z
	if err = json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf("parse post result failed: %v, data: %v", err, string(data))
	}
	if !result.Success() {
		return fmt.Errorf("post failed: %v, data: %v", result.TKIFA83m_r, string(data))
	}
	return nil
}
func (p *CsKk5VzGb) sXWeuE(path string) ([]byte, error) {
	address := p.GetUrl(path)
	status, data, headers, err := relay.IlWBbp1xdPw(address, nil, p.ojEs5pzln(address))
	if err != nil {
		return nil, fmt.Errorf("%v, result: %v, status code: %v", err, string(data), status)
	}
	switch status {
	case 200:
		if etag := headers.Get("etag"); etag != "" {
			if v, ok := p.cI03v8q7.CheckETag(address, etag); ok {
				return v.([]byte), nil
			}
			p.cI03v8q7.AddETag(address, etag, data)
		}
		return data, nil
	case 304:
		if v, ok := p.cI03v8q7.GetObj(address); ok {
			return v.([]byte), nil
		}
		return nil, fmt.Errorf("can not find eTag: %v", address)
	default:
		return nil, fmt.Errorf("request failed, result: %v, status code: %v", string(data), status)
	}
}
