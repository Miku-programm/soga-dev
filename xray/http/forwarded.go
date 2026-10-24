package http

import (
	"net/http"
	N "recovered.local/soga/xray/net"
	"strings"
)

// ParseXForwardedFor 0x1250760。保留所有项（包括空项和域名），不跳过非法 IP。
func ParseXForwardedFor(h http.Header) []N.Address {
	v := h.Get("X-Forwarded-For")
	if v == "" {
		return nil
	}
	parts := strings.Split(v, ",")
	out := make([]N.Address, 0, len(parts))
	for _, part := range parts {
		out = append(out, N.ParseAddress(part))
	}
	return out
}
