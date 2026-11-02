package xiaov2board

import (
	"net/http"
	"net/url"
	"recovered.local/soga/config"
	"strconv"
)

func (p *LRlqOaYmO) GetAPIConfig() *config.FCEAfQJwDA7e { return p.EA_5pkz }

func (p *LRlqOaYmO) cIbNGPeZYyl() string {
	if config.XiaoV2NodeMode() {
		return "v2node"
	}
	switch p.EA_5pkz.ZyAhkq5th8bf {
	case "ss":
		return "shadowsocks"
	case "v2ray":
		return "v2ray"
	case "vless", "anytls", "hysteria", "mieru":
		return string(p.EA_5pkz.ZyAhkq5th8bf)
	case "trojan":
		return "trojan"
	}
	return ""
}

func (p *LRlqOaYmO) fUwRFGkp() url.Values {
	values := url.Values{"token": {p.EA_5pkz.Bd08Li}, "node_id": {strconv.Itoa(p.EA_5pkz.Yf1KqSqJpNI7)}}
	if nodeType := p.cIbNGPeZYyl(); nodeType != "" {
		values["node_type"] = []string{nodeType}
	}
	return values
}

func (p *LRlqOaYmO) GetUniUrl(action string) string {

	return p.EA_5pkz.W_ogGtF + "api/v1/server/UniProxy/" + action
}

func (p *LRlqOaYmO) iU3FqWNT(key string) http.Header {
	etag := p.eyq32s.GetETag(key)
	if etag == "" {
		return nil
	}
	return http.Header{"IF-NONE-MATCH": {etag}}
}

// kwoj3qch_ 对应 0x17a62e0。
func (p *LRlqOaYmO) kwoj3qch_() string {
	if config.XiaoV2NodeMode() {
		return p.EA_5pkz.W_ogGtF + "api/v2/server/config"
	}
	return p.GetUniUrl("config")
}
