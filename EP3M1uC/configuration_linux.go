//go:build linux

package EP3M1uC

import (
	"fmt"
	A "recovered.local/soga/LCe4Wkk2"
	T "recovered.local/soga/umGtgyvgLEWz"
)

// KmKjIZi1 restores 0x1b79820, including defaults before enabled validation.
func KmKjIZi1(data string) (*UIfnUsim, error) {
	routes := &UIfnUsim{CnK0a4: true, H839j9DYDD5k: true}
	if _, err := T.Ikm5Xa(data, routes); err != nil {
		return nil, err
	}
	for _, rule := range routes.MMKbp60np {
		rule.tIXCrVLXdv2w()
	}
	if err := routes.o7v2ON9D(); err != nil {
		return nil, err
	}
	return routes, nil
}

// o7v2ON9D restores 0x1b79a60. Disabled configurations skip construction.
func (routes *UIfnUsim) o7v2ON9D() error {
	if !routes.CnK0a4 {
		return nil
	}
	for _, rule := range routes.MMKbp60np {
		if err := rule.o7v2ON9D(); err != nil {
			return err
		}
	}
	return nil
}

// o7v2ON9D restores 0x1b86a80; adapter failures precede matcher validation.
func (rule *DgAjL9UN) o7v2ON9D() error {
	for _, out := range rule.UwkgrlFf8 {
		if err := out.o7v2ON9D(); err != nil {
			return err
		}
	}
	return rule.initializeMatchers()
}

// o7v2ON9D restores 0x1b7c4c0 and its original constant option templates.
// HTTP/SOCKS do not inherit the route TLS fields. Trojan inherits only SNI
// and skip-cert-verify; its network/path/gRPC fields stay at template values.
func (out *GyyO51) o7v2ON9D() error {
	if err := out.initializeAddresses(); err != nil {
		return err
	}
	var err error
	switch out.NyLbiXb4OrlL {
	case "ss":
		out.ucp4iC, err = A.PQ9jOjtX(A.ShadowsocksOptions{Server: out.AHTlIEynXL, Port: out.JO1YIpXk, Password: out.CiOEvesmFF, Cipher: out.ExA8yD, UDP: true})
	case "ssr":
		out.ucp4iC, err = A.CR6bC5(A.ShadowsocksROptions{Server: out.AHTlIEynXL, Port: out.JO1YIpXk, Password: out.CiOEvesmFF, Cipher: out.ExA8yD, Obfs: out.TmSpNP, ObfsParam: out.Ca3YcWv688v, Protocol: out.JuTg4y, ProtocolParam: out.P57BOZ, UDP: true})
	case "http":
		out.ucp4iC = A.CJW16StU3I(A.HTTPOptions{Server: out.AHTlIEynXL, Port: out.JO1YIpXk, Username: out.XFrkscEI, Password: out.CiOEvesmFF})
	case "socks":
		out.ucp4iC = A.Auqn90u2NBQ(A.SocksOptions{Server: out.AHTlIEynXL, Port: out.JO1YIpXk, Username: out.XFrkscEI, Password: out.CiOEvesmFF, UDP: true})
	case "vmess":
		out.ucp4iC, err = A.Fr9G0QQhtaac(A.YaiP1Z10VuOK{PaWTH6: out.AHTlIEynXL, KQ7hyM: out.JO1YIpXk, MMeZdiHjQ: out.PJkA7ZsW4FB, TJqENyK7: out.FegVmoOB, YKnBs48P7XY: "auto", DS2NjE9s_: true, EtDA3oIPdZ7y: string(out.LUjxWBsS2Jl), MLU1J_: out.UkmhqYCQ9u, VPGpSFLRW: out.SRegcRbFQo, JPPDxQ: out.JqKHTUS0FUa, FN179GVH_: A.THKLdNRQ{IDDJaTZ: out.ZNfRIgJOM4uR}, U7nq5HsQY: A.HrHLa45bjuV{TTsHbH5qb: out.AJMYgKaXJz}})
	case "direct":
		out.ucp4iC = A.NewDirectAdapter()
	case "trojan":
		out.ucp4iC, err = A.RFqiXSepf_9(A.ZxrFcPHYpol{HF6Vma: out.AHTlIEynXL, GRc9GM8NFW: out.JO1YIpXk, Anl_Yj: out.CiOEvesmFF, IGRMFmQSwT: out.JqKHTUS0FUa, R1aqZuYO: out.SRegcRbFQo, TNtpmYwuy: true})
	case "redirect":
		// No adapter assignment occurs in the original redirect branch.
	default:
		return fmt.Errorf("unknown out type: %v", out.NyLbiXb4OrlL)
	}
	return err
}
