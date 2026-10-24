package matcher

import (
	"inet.af/netaddr"
	"log/slog"
	"net"
	"recovered.local/soga/geo"
	"recovered.local/soga/util/logx"
	"strings"
)

// ZLM7bds8T follows the original two *netaddr.IPSet fields (16 bytes).
// The original passes these percent signs to Sprintln, not Sprintf.
var invalidGeoIPLog = "unknown cidr in geoip:%v:%v"

type ZLM7bds8T struct{ uNRj4pznONc, kaTS6UFYvuwx *netaddr.IPSet }

// HcZBFktYT restores 0xebcba0. Invalid patterns log warnings and do not abort.
func HcZBFktYT(patterns []string) *ZLM7bds8T {
	m := new(ZLM7bds8T)
	var ipv4, ipv6 netaddr.IPSetBuilder
	for _, pattern := range patterns {
		m.iF3QxW1j2A(&ipv4, &ipv6, pattern)
	}
	var err error
	m.uNRj4pznONc, err = ipv4.IPSet()
	if err != nil {
		logx.TRTbrquD(slog.LevelError, err)
	}
	m.kaTS6UFYvuwx, err = ipv6.IPSet()
	if err != nil {
		logx.TRTbrquD(slog.LevelError, err)
	}
	return m
}

// iF3QxW1j2A restores 0xebcd80. geoip: is tested before optional ip: removal.
func (m *ZLM7bds8T) iF3QxW1j2A(ipv4, ipv6 *netaddr.IPSetBuilder, pattern string) {
	if m.jrhCUUp(ipv4, ipv6, pattern) {
		return
	}
	pattern = strings.TrimPrefix(pattern, "ip:")
	if m.biOzcBi(ipv4, ipv6, pattern) || m.i7iRt1(ipv4, ipv6, pattern) {
		return
	}
	logx.TRTbrquD(slog.LevelWarn, "unknown ip rule:", pattern)
}

// jrhCUUp restores 0xebcf40, using the original GeoIP records in order.
func (m *ZLM7bds8T) jrhCUUp(ipv4, ipv6 *netaddr.IPSetBuilder, pattern string) bool {
	if !strings.HasPrefix(pattern, "geoip:") {
		return false
	}
	code := pattern[6:]
	data, err := geo.YugPy6709RY(code)
	if err != nil {
		logx.TRTbrquD(slog.LevelWarn, err)
		return true
	}
	for _, cidr := range data.GetCidr() {
		ip, ok := netaddr.FromStdIP(cidr.GetIp())
		if !ok {
			logx.TRTbrquD(slog.LevelWarn, invalidGeoIPLog, code, cidr)
			continue
		}
		prefix := netaddr.IPPrefixFrom(ip, uint8(cidr.GetPrefix()))
		if ip.Is4() {
			ipv4.AddPrefix(prefix)
		} else if ip.Is6() {
			ipv6.AddPrefix(prefix)
		}
	}
	return true
}

// biOzcBi restores 0xebd220.
func (m *ZLM7bds8T) biOzcBi(ipv4, ipv6 *netaddr.IPSetBuilder, pattern string) bool {
	parsed := net.ParseIP(pattern)
	if parsed == nil {
		return false
	}
	ip, ok := netaddr.FromStdIP(parsed)
	if !ok {
		logx.TRTbrquD(slog.LevelWarn, "invalid ip:", parsed)
		return true
	}
	if ip.Is4() {
		ipv4.Add(ip)
	} else if ip.Is6() {
		ipv6.Add(ip)
	}
	return true
}

// i7iRt1 restores 0xebd420, retaining net.ParseCIDR normalization.
func (m *ZLM7bds8T) i7iRt1(ipv4, ipv6 *netaddr.IPSetBuilder, pattern string) bool {
	_, network, err := net.ParseCIDR(pattern)
	if err != nil {
		return false
	}
	prefix, ok := netaddr.FromStdIPNet(network)
	if !ok {
		logx.TRTbrquD(slog.LevelWarn, "invalid ip cidr:", pattern)
		return true
	}
	if prefix.IP().Is4() {
		ipv4.AddPrefix(prefix)
	} else if prefix.IP().Is6() {
		ipv6.AddPrefix(prefix)
	}
	return true
}

// TestIP restores 0xebd5a0; FromStdIP unmaps IPv4-mapped IPv6 inputs.
func (m *ZLM7bds8T) TestIP(value net.IP) bool {
	ip, ok := netaddr.FromStdIP(value)
	if !ok {
		return false
	}
	if ip.Is4() {
		return m.uNRj4pznONc.Contains(ip)
	}
	if !ip.Is6() {
		return false
	}
	return m.kaTS6UFYvuwx.Contains(ip)
}
