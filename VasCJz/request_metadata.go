package VasCJz

import (
	"net/http"
	"strings"
)

// 0x1252a00。调用者负责此前缀存在；本函数只以其字节长度切片。
func (c *R0J19auWof) ExtractMetaFromRequest(r *http.Request, prefix string) (session, sequence string) {
	sp, qp := c.GetNormalizedSessionPlacement(), c.GetNormalizedSeqPlacement()
	sk, qk := c.GetNormalizedSessionKey(), c.GetNormalizedSeqKey()
	var parts []string
	if sp == "path" || qp == "path" {
		parts = strings.Split(r.URL.Path[len(prefix):], "/")
	}
	next := 0
	get := func(place, key string) string {
		switch place {
		case "path":
			if next < len(parts) {
				value := parts[next]
				next++
				return value
			}
		case "header":
			return r.Header.Get(key)
		case "query":
			return r.URL.Query().Get(key)
		case "cookie":
			if x, e := r.Cookie(key); e == nil {
				return x.Value
			}
		}
		return ""
	}
	session = get(sp, sk)
	sequence = get(qp, qk)
	return
}

// 0x1251c40。这里只设置原 CORS 响应字段，不写状态码。
func (c *R0J19auWof) WriteResponseHeader(w http.ResponseWriter, method string, request http.Header) {
	origin := request.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)
	if c.GetNormalizedSessionPlacement() == "cookie" || c.GetNormalizedSeqPlacement() == "cookie" || c.V0qykwP == "cookie" || c.GetNormalizedUplinkDataPlacement() == "cookie" {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	if method == "OPTIONS" {
		value := request.Get("Access-Control-Request-Method")
		if value == "" {
			value = "*"
		}
		w.Header().Set("Access-Control-Allow-Methods", value)
		value = request.Get("Access-Control-Request-Headers")
		if value == "" {
			value = "*"
		}
		w.Header().Set("Access-Control-Allow-Headers", value)
	}
}
