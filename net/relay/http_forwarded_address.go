package relay

import (
	"net"
	"net/http"
	"strings"
)

// SOmphOG1Ty @0x15a5be0。原类型为HTTP头映射；先读X-Real-Ip，
// 仅当它为空才取X-Forwarded-For首项并TrimSpace。非空但非法的Real-IP不回退。
func SOmphOG1Ty(header http.Header) net.Addr {
	text := header.Get("X-Real-Ip")
	if text == "" {
		if forwarded := header.Get("X-Forwarded-For"); forwarded != "" {
			text = strings.TrimSpace(strings.Split(forwarded, ",")[0])
		}
	}
	if text != "" {
		if ip := net.ParseIP(text); ip != nil {
			return &net.TCPAddr{IP: ip}
		}
	}
	return nil
}
