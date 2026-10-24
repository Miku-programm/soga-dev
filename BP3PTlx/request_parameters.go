package BP3PTlx

import (
	"net/http"
	"net/url"
	"recovered.local/soga/cNT24rG4F"
	"strconv"
)

func (p *KlO65BuwOt) GetAPIConfig() *cNT24rG4F.FCEAfQJwDA7e { return p.EA_5pkz }

func (p *KlO65BuwOt) f6nmeuaym() string {
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

func (p *KlO65BuwOt) tgg_NuFN() url.Values {
	values := url.Values{"token": {p.EA_5pkz.Bd08Li}, "node_id": {strconv.Itoa(p.EA_5pkz.Yf1KqSqJpNI7)}}
	if nodeType := p.f6nmeuaym(); nodeType != "" {
		values["node_type"] = []string{nodeType}
	}
	return values
}

func (p *KlO65BuwOt) rTN7wPdD(action string) string {

	return p.EA_5pkz.W_ogGtF + "api/v1/server/UniProxy/" + action
}

func (p *KlO65BuwOt) x7aWIuw06(key string) http.Header {
	etag := p.eyq32s.GetETag(key)
	if etag == "" {
		return nil
	}
	return http.Header{"IF-NONE-MATCH": {etag}}
}
