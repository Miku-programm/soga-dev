package BaEWsPy

import (
	"context"
	"net"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/xx5fRyrFw"
	"reflect"
	"testing"
)

func TestDNSRulesOrderAndSeparateTypes(t *testing.T) {
	first := &I1JijVatatst{luBuRE: xx5fRyrFw.Ij9r21E0UOt3{xx5fRyrFw.NewDNSMatcher("domain:example.test")}, qBHWeHSCLI: []uint16{28}}
	second := &I1JijVatatst{luBuRE: xx5fRyrFw.Ij9r21E0UOt3{xx5fRyrFw.NewDNSMatcher("")}, qBHWeHSCLI: []uint16{1}}
	rules := ZrMh8EfBQar{first, second}
	if got, ok := rules.Test("www.example.test"); !ok || got != first {
		t.Fatal("first matching rule")
	}
	if got, ok := rules.Test("other.test"); !ok || got != second {
		t.Fatal("empty DNS pattern must match all")
	}
	first.aSg142Hhy1 = GG2P4UAyQqI{lookupFunc(func(_ context.Context, _ string, typ uint16) ([]net.IP, error) {
		if typ != 28 {
			t.Error("rule type not selected")
		}
		return []net.IP{net.ParseIP("2001:db8::4")}, nil
	})}
	r := testResolver(t, 0)
	r.a06M4WZc = rules
	got, err := r.Resolve("www.example.test")
	if err != nil || len(got) != 1 || got[0].String() != "2001:db8::4" {
		t.Fatalf("routed result %v %v", got, err)
	}
}
func TestDNSStrategyAndConstructor(t *testing.T) {
	cases := map[string][]uint16{"ipv4_only": {1}, "ipv6_only": {28}, "ipv6_first": {28, 1}, "ipv4_first": {1, 28}, " IPv6_only ": {1, 28}, "": {1, 28}}
	for strategy, want := range cases {
		if got := dnsTypes(strategy); !reflect.DeepEqual(got, want) {
			t.Fatalf("%q %v != %v", strategy, got, want)
		}
	}
	cfg := &LFWTTXF7JR.Axa8uKIB4JL{K8BNfziIgb: []string{"local"}, CM0eMsqyP: "ipv6_first", CZ80ik: []string{"full:example.test", ""}}
	rule := o4RFaBT0MwY3(cfg, nil)
	if len(rule.aSg142Hhy1) != 1 || len(rule.luBuRE) != 2 || !reflect.DeepEqual(rule.qBHWeHSCLI, []uint16{28, 1}) || !rule.Test("all.test") {
		t.Fatal("rule constructor")
	}
	r := testResolver(t, 0)
	if r.For(nil) != r {
		t.Fatal("nil address must preserve receiver")
	}
}
