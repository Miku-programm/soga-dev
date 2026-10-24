package xx5fRyrFw

import (
	"inet.af/netaddr"
	"net"
	"testing"
	"unsafe"
)

func TestNetworkMatcherLayouts(t *testing.T) {
	if unsafe.Sizeof(ZLM7bds8T{}) != 16 || unsafe.Sizeof(Z5Jdu9{}) != 8 || unsafe.Sizeof(DaCngxGKN{}) != 3 || unsafe.Sizeof(netaddr.IPSetBuilder{}) != 72 {
		t.Fatal("RTTI layout mismatch")
	}
}
func TestProtocolMatcherSemantics(t *testing.T) {
	for _, pattern := range []string{"", "invalid", " tcp"} {
		m := XSe2ktMsZnSQ(pattern)
		if !m.Match("quic") {
			t.Fatalf("no recognized token must allow all: %q", pattern)
		}
	}
	m := XSe2ktMsZnSQ("TCP, udp")
	if !m.Match("tcp") || m.Match("udp") || m.Match("TCP") {
		t.Fatal("lowercase config only; no whitespace trim or query normalization")
	}
	if !(*DaCngxGKN)(nil).Match("unknown") {
		t.Fatal("nil matcher should accept")
	}
}
func TestPortMatcherRangesAndInvalid(t *testing.T) {
	m := SekQr0GHk([]string{"port:80,443,103-100", "0,65535,65536", " 81", "bad", "PORT:82", "90-90", "2--1"})
	for _, p := range []int{1, 2, 80, 443, 100, 101, 102, 103, 65535, 90} {
		if !m.Match(p) {
			t.Fatalf("missing port %d", p)
		}
	}
	for _, p := range []int{-1, 0, 81, 82, 99, 104, 65536} {
		if m.Match(p) {
			t.Fatalf("unexpected port %d", p)
		}
	}
	if (*Z5Jdu9)(nil).Match(0) {
		t.Fatal("invalid port should return before dereference")
	}
}
func TestIPMatcherFamiliesNormalizationAndInvalid(t *testing.T) {
	m := HcZBFktYT([]string{"ip:192.0.2.7", "198.51.100.7/24", "2001:db8::123/32", "bad", "ip:geoip:private"})
	for _, s := range []string{"192.0.2.7", "::ffff:192.0.2.7", "198.51.100.255", "2001:db8:abcd::1"} {
		if !m.TestIP(net.ParseIP(s)) {
			t.Fatalf("missing %s", s)
		}
	}
	for _, s := range []string{"192.0.2.8", "198.51.101.1", "2001:db9::1", "10.0.0.1"} {
		if m.TestIP(net.ParseIP(s)) {
			t.Fatalf("unexpected %s", s)
		}
	}
	if m.TestIP(nil) || m.TestIP(net.IP{1, 2, 3}) {
		t.Fatal("malformed IP accepted")
	}
}
func TestOriginalGeoIPPrivateRules(t *testing.T) {
	m := HcZBFktYT([]string{"geoip:private"})
	for _, s := range []string{"10.0.0.1", "192.168.5.6", "127.0.0.1", "::1"} {
		if !m.TestIP(net.ParseIP(s)) {
			t.Fatalf("private database omitted %s", s)
		}
	}
	if m.TestIP(net.ParseIP("8.8.8.8")) {
		t.Fatal("public address matched private")
	}
}
