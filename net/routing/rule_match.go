package routing

import (
	"fmt"
	"recovered.local/soga/crypto/rand"
)

// Static pointer 0x5029210 points at 0x50781a8; its sole initialized field is
// NyLbiXb4OrlL="direct" at offset 16. Its health flag is initially false.
var defaultOutboundAt5029210 = &GyyO51{NyLbiXb4OrlL: "direct"}

// IsMatch restores 0x1b877e0. Protocol is mandatory; other matchers are ORed.
func (rule *DgAjL9UN) IsMatch(request *AbtxqK3iZa2) bool {
	if !rule.yr2tuiHRVF.Match(request.WBmwgcAHw) {
		return false
	}
	if rule.__ev4XIbr {
		return true
	}
	target := request.BqOvrz
	if rule.aru24wr.Match(target.GetPort()) {
		return true
	}
	if rule.ov3c4tIGkqkb.Has(request.YnL8s4_1IH) {
		return true
	}
	host := target.GetHost()
	if target.IsIP() {
		host = ""
	}
	firstIP := target.GetFirstIp()
	if host != "" {
		for _, matcher := range rule.hSORasT5h {
			if matcher.Test(host) {
				return true
			}
		}
	}
	return firstIP != nil && rule.cPaGes.TestIP(firstIP)
}

// oImRPaEmL restores 0x1b87a20. Empty/missing healthy snapshots deliberately
// use a random configured outbound, even if it is currently unhealthy.
func (rule *DgAjL9UN) oImRPaEmL(request *AbtxqK3iZa2) *GyyO51 {
	if healthy := rule.jf_lZ8wS2GNO.Load(); healthy != nil && len(*healthy) > 0 {
		return rule.yvcsQOnq.Select(*healthy, request.Kbj7uES, request.DmaA6p)
	}
	return rule.UwkgrlFf8[rand.Co6D_zfls(len(rule.UwkgrlFf8))]
}

// HasRoute restores 0x1b79ae0.
func (routes *UIfnUsim) HasRoute() bool { return routes.CnK0a4 && len(routes.MMKbp60np) != 0 }

// b7i5V4Y4B restores 0x1b79b00, preserving first-rule priority.
func (routes *UIfnUsim) b7i5V4Y4B(request *AbtxqK3iZa2) (*DgAjL9UN, error) {
	for _, rule := range routes.MMKbp60np {
		if rule.IsMatch(request) {
			return rule, nil
		}
	}
	return nil, fmt.Errorf("not route match host: %v, protocol: %v", request.BqOvrz.GetHost(), request.WBmwgcAHw)
}

// SelectOut restores 0x1b79cc0. Enabled routes with no matching rule return an
// error; the static direct outbound is only for disabled/empty routing.
func (routes *UIfnUsim) SelectOut(request *AbtxqK3iZa2) (*GyyO51, error) {
	if !routes.HasRoute() {
		return defaultOutboundAt5029210, nil
	}
	rule, err := routes.b7i5V4Y4B(request)
	if err != nil {
		return nil, err
	}
	return rule.oImRPaEmL(request), nil
}
