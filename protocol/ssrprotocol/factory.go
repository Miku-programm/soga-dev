package ssrprotocol

import (
	"fmt"
	T "recovered.local/soga/protocol/ssrcipher"
	E "recovered.local/soga/panel/ssrnames"
)

// Original map 0x5096970 and init.0 through init.8, in registration order.
// All nine registered constructors and their IO paths are restored here.
var protocolsAt5096970 = map[E.Mu0vMIIs]ojt4CWtzUu{
	"auth_aes128_md5":  jKGRC86Lx9g,
	"auth_aes128_sha1": dffBLS,
	"auth_chain_a":     oPPzNk,
	"auth_chain_b":     gYQ1Wtm,
	"auth_chain_c":     vHtqJ7xSQiSj,
	"auth_chain_d":     daajdwCbv7,
	"auth_chain_e":     xrwCp6,
	"auth_chain_f":     hH3aFx,
	"origin":           jyvuea4xsHn,
}

// HrQWlmA restores 0x1cdc720; it does not normalize names or alias none.
func HrQWlmA(name E.Mu0vMIIs, c *T.KrOWYtzUYft) (Vzuv_s9ZS5, error) {
	constructor, ok := protocolsAt5096970[name]
	if !ok {
		return nil, fmt.Errorf("protocol %v not support", name)
	}
	return constructor(c), nil
}
