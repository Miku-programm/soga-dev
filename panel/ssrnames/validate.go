package ssrnames

import "fmt"

// PsDhAx0aUZ @0x1249b40: independently recovered three ordered allowlists.
func PsDhAx0aUZ(method ZvAxOt1l7Ar, protocol Mu0vMIIs, obfs PMorzD7u) error {
	switch method {
	case "rc4", "none", "bf-cfb", "des-cfb", "salsa20", "rc2-cfb", "rc4-md5", "chacha20", "idea-cfb", "cast5-cfb", "rc4-md5-6", "aes-128-ofb", "aes-192-ofb", "aes-256-ofb", "aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr", "aes-128-cfb", "aes-192-cfb", "aes-256-cfb", "chacha20-ietf", "camellia-128-cfb", "camellia-192-cfb", "camellia-256-cfb", "chacha20-ietf-poly1305":
	default:
		return fmt.Errorf("not support ssr method: %v", method)
	}
	switch protocol {
	case "origin", "auth_aes128_md5", "auth_aes128_sha1", "auth_chain_a", "auth_chain_b", "auth_chain_c", "auth_chain_d", "auth_chain_e", "auth_chain_f":
	default:
		return fmt.Errorf("not support ssr protocol: %v", protocol)
	}
	switch obfs {
	case "none", "plain", "http_post", "http_simple", "simple_obfs_tls", "simple_obfs_http", "tls1.2_ticket_auth":
	default:
		return fmt.Errorf("not support ssr obfs: %v", obfs)
	}
	return nil
}
