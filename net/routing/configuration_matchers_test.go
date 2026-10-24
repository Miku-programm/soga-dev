package routing

import (
	"net"
	"recovered.local/soga/net/relay"
	"strconv"
	"testing"
)

func TestRouteConfigurationDefaultsAndSnapshots(t *testing.T) {
	a := &GyyO51{AHTlIEynXL: "peer.example"}
	b := &GyyO51{AHTlIEynXL: "other.example", SRegcRbFQo: true}
	r := &DgAjL9UN{UwkgrlFf8: J_enG0wE{a, b}, TnykhVDupaKh: "round_robin"}
	r.tIXCrVLXdv2w()
	if a.LUjxWBsS2Jl != "tcp" || a.JaH5hv != "ipv4_first" || a.JqKHTUS0FUa != "peer.example" || b.JqKHTUS0FUa != "" {
		t.Fatal(a, b)
	}
	if r.yvcsQOnq.Select(r.UwkgrlFf8, nil, 0) != b {
		t.Fatal("original round-robin starts at index 1")
	}
	a.ebDTwNX(false)
	a.ebDTwNX(false)
	a.ebDTwNX(false)
	r.tIXCrVLXdv2w()
	snapshot := r.jf_lZ8wS2GNO.Load()
	if a.IsEnabled() || snapshot == nil || len(*snapshot) != 1 || (*snapshot)[0] != b {
		t.Fatal("initialization re-enabled unhealthy outbound")
	}
}

func TestCompiledRouteMatchersAndOriginalReinitialization(t *testing.T) {
	rule := &DgAjL9UN{UwkgrlFf8: J_enG0wE{new(GyyO51)}, Ehbw9wa: "TCP", OZmvCwn: []string{"ip:192.0.2.0/24", "port:443", "node_id: 12,broken,-3,12", "full:example.org"}}
	if err := rule.initializeMatchers(); err != nil {
		t.Fatal(err)
	}
	if !rule.ov3c4tIGkqkb.Has(12) || !rule.ov3c4tIGkqkb.Has(-3) || rule.ov3c4tIGkqkb.Len() != 2 {
		t.Fatal("node identifiers differ")
	}
	for _, tc := range []struct {
		host       string
		port, node int
		match      bool
	}{{"192.0.2.1", 80, 0, true}, {"198.51.100.2", 443, 0, true}, {"unmatched.example", 80, 12, true}, {"example.org", 80, 0, true}, {"unmatched.example", 80, 0, false}} {
		target, err := relay.U80FDt(net.JoinHostPort(tc.host, strconv.Itoa(tc.port)))
		if err != nil {
			t.Fatal(err)
		}
		req := &AbtxqK3iZa2{BqOvrz: target, YnL8s4_1IH: tc.node, WBmwgcAHw: "tcp"}
		if got := rule.IsMatch(req); got != tc.match {
			t.Fatalf("%+v: %v", tc, got)
		}
		req.WBmwgcAHw = "udp"
		if rule.IsMatch(req) {
			t.Fatal("protocol gate bypassed")
		}
	}
	rule.OZmvCwn = []string{"*"}
	if err := rule.initializeMatchers(); err != nil {
		t.Fatal(err)
	}
	rule.OZmvCwn = []string{"full:another.org"}
	if err := rule.initializeMatchers(); err != nil {
		t.Fatal(err)
	}
	if !rule.__ev4XIbr || len(rule.hSORasT5h) != 2 {
		t.Fatal("original reinitialization retains wildcard flag and old domains")
	}
}

func TestRouteProtocolWhitespaceAsymmetryAndValidation(t *testing.T) {
	r := &DgAjL9UN{UwkgrlFf8: J_enG0wE{new(GyyO51)}, Ehbw9wa: " tcp "}
	if err := r.initializeMatchers(); err != nil {
		t.Fatal(err)
	}
	if !r.yr2tuiHRVF.Match("udp") {
		t.Fatal("original validation trims, matcher does not; must preserve this behavior")
	}
	r.Ehbw9wa = "tcp,quic"
	if err := r.initializeMatchers(); err == nil || err.Error() != "invalid route protocol: tcp,quic, only support tcp/udp" {
		t.Fatal(err)
	}
	r.Ehbw9wa = ""
	if err := r.initializeMatchers(); err == nil || err.Error() != "please configure at least one rules or protocol" {
		t.Fatal(err)
	}
	r.UwkgrlFf8 = nil
	if err := r.initializeMatchers(); err == nil || err.Error() != "please configure at least one outs" {
		t.Fatal(err)
	}
}
