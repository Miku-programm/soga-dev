package sspaneltypes

import (
	L "recovered.local/soga/panel/types"
	M "recovered.local/soga/util/logx"
	"strconv"
)

// UpdateSSRNode @0x16e99c0 retains the original validation order and warning exits.
func (c *CEJsaaVtomQg) UpdateSSRNode(n *L.Dv6QJyFW6bU) {
	if !n.QawbKS2x {
		return
	}
	v := c.HS63ppHsdS
	port, _ := strconv.Atoi(v.SMaxEaB)
	if port <= 0 {
		port = c.uX0Wxx()
	}
	if port <= 0 {
		return
	}
	if v.RNq55eY == "" {
		M.TRTbrquD(4, "ssr password can not be empty")
		return
	}
	method, protocol, obfs := v.GetMethod(), v.GetProtocol(), v.GetObfs()
	switch method {
	case "rc4", "none", "bf-cfb", "des-cfb", "salsa20", "rc2-cfb", "rc4-md5", "chacha20", "idea-cfb", "cast5-cfb", "rc4-md5-6", "aes-128-ofb", "aes-192-ofb", "aes-256-ofb", "aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr", "aes-128-cfb", "aes-192-cfb", "aes-256-cfb", "chacha20-ietf", "camellia-128-cfb", "camellia-192-cfb", "camellia-256-cfb", "chacha20-ietf-poly1305":
	default:
		M.V8GN5Stk(4, "unknown ssr method: <%v>", method)
		return
	}
	switch protocol {
	case "origin", "auth_aes128_md5", "auth_aes128_sha1", "auth_chain_a", "auth_chain_b", "auth_chain_c", "auth_chain_d", "auth_chain_e", "auth_chain_f":
	default:
		M.V8GN5Stk(4, "unknown ssr protocol: <%v>", protocol)
		return
	}
	switch obfs {
	case "none", "plain", "http_post", "http_simple", "simple_obfs_tls", "simple_obfs_http", "tls1.2_ticket_auth":
	default:
		M.V8GN5Stk(4, "unknown ssr obfs: <%v>", obfs)
		return
	}
	n.AddSSRUser(&L.Aga91y{JAfXWwzYrW: &L.QcqpSL3i6ri{YrjVl_2g_S: port, YyAgnMatmj: v.RNq55eY, Z_taAb: method, H_y2xZht7S8: protocol, ZZIDsJ: obfs, BskzckkE7qI: 2}})
}
