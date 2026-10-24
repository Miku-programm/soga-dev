package dns

import (
	"context"
	"net"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/audit/matcher"
)

// o4RFaBT0MwY3: 0x1586860。空规则仍创建恒真匹配器。
func o4RFaBT0MwY3(cfg *types.Axa8uKIB4JL, ip net.IP) *I1JijVatatst {
	servers := make(GG2P4UAyQqI, 0, len(cfg.K8BNfziIgb))
	for _, server := range cfg.K8BNfziIgb {
		servers = append(servers, cDpHJRFxZ(server, ip))
	}
	matchers := make(matcher.Ij9r21E0UOt3, 0, len(cfg.CZ80ik))
	for _, pattern := range cfg.CZ80ik {
		matchers = append(matchers, matcher.NewDNSMatcher(pattern))
	}
	return &I1JijVatatst{servers, matchers, dnsTypes(cfg.CM0eMsqyP)}
}

// Test: 0x1586d60。
func (r *I1JijVatatst) Test(host string) bool {
	for _, matcher := range r.luBuRE {
		if matcher.Test(host) {
			return true
		}
	}
	return false
}

// LookupDNSType: 0x1586e20。
func (r *I1JijVatatst) LookupDNSType(ctx context.Context, host string, typ uint16) ([]net.IP, error) {
	return r.aSg142Hhy1.LookupDNSType(ctx, host, typ)
}

// Test: 0x1586ec0。按配置顺序选择第一条规则。
func (rules ZrMh8EfBQar) Test(host string) (*I1JijVatatst, bool) {
	for _, rule := range rules {
		if rule.Test(host) {
			return rule, true
		}
	}
	return nil, false
}
