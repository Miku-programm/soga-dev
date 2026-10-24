package EP3M1uC

import (
	"net"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/G7A7csgrs62p"
	"recovered.local/soga/xx5fRyrFw"
	"testing"
	"unsafe"
)

func preparedRule() *DgAjL9UN {
	return &DgAjL9UN{aru24wr: xx5fRyrFw.SekQr0GHk(nil), ov3c4tIGkqkb: G7A7csgrs62p.PQQKui[int](), cPaGes: xx5fRyrFw.HcZBFktYT(nil)}
}
func TestRuleRTTILayout(t *testing.T) {
	if unsafe.Sizeof(DgAjL9UN{}) != 168 || unsafe.Sizeof(AbtxqK3iZa2{}) != 64 || unsafe.Sizeof(UIfnUsim{}) != 32 || unsafe.Sizeof(Qi1aOz5oQ{}) != 32 {
		t.Fatal("rule RTTI sizes differ")
	}
}
func TestRuleProtocolGateAndORMatching(t *testing.T) {
	target, err := G6bHD5JA.U80FDt("example.com:443")
	if err != nil {
		t.Fatal(err)
	}
	req := &AbtxqK3iZa2{BqOvrz: target, WBmwgcAHw: "tcp", YnL8s4_1IH: 12}
	rule := preparedRule()
	rule.yr2tuiHRVF = xx5fRyrFw.XSe2ktMsZnSQ("udp")
	rule.__ev4XIbr = true
	if rule.IsMatch(req) {
		t.Fatal("all rule bypassed protocol gate")
	}
	req.WBmwgcAHw = "udp"
	if !rule.IsMatch(req) {
		t.Fatal("all rule missed")
	}
	rule.__ev4XIbr = false
	rule.aru24wr = xx5fRyrFw.SekQr0GHk([]string{"443"})
	if !rule.IsMatch(req) {
		t.Fatal("port match missed")
	}
	rule.aru24wr = xx5fRyrFw.SekQr0GHk(nil)
	rule.ov3c4tIGkqkb.Add(12)
	if !rule.IsMatch(req) {
		t.Fatal("node match missed")
	}
	rule.ov3c4tIGkqkb.Del(12)
	rule.hSORasT5h = xx5fRyrFw.Ij9r21E0UOt3{xx5fRyrFw.NewDNSMatcher("full:example.com")}
	if !rule.IsMatch(req) {
		t.Fatal("domain match missed")
	}
	ipTarget, err := G6bHD5JA.RNFJeaRY(&net.TCPAddr{IP: net.ParseIP("192.0.2.1"), Port: 1})
	if err != nil {
		t.Fatal(err)
	}
	req.BqOvrz = ipTarget
	rule.hSORasT5h = nil
	rule.cPaGes = xx5fRyrFw.HcZBFktYT([]string{"192.0.2.0/24"})
	if !rule.IsMatch(req) {
		t.Fatal("IP matcher missed")
	}
}
func TestRoutePriorityNoMatchAndHealthySnapshot(t *testing.T) {
	target, err := G6bHD5JA.U80FDt("example.com:80")
	if err != nil {
		t.Fatal(err)
	}
	req := &AbtxqK3iZa2{BqOvrz: target, WBmwgcAHw: "tcp"}
	first, second := preparedRule(), preparedRule()
	a, b := new(GyyO51), new(GyyO51)
	first.__ev4XIbr = true
	second.__ev4XIbr = true
	first.UwkgrlFf8 = J_enG0wE{a}
	second.UwkgrlFf8 = J_enG0wE{b}
	routes := &UIfnUsim{CnK0a4: true, MMKbp60np: ViF0a1nppLG{first, second}}
	if got, e := routes.SelectOut(req); e != nil || got != a {
		t.Fatal("first rule must win")
	}
	snapshot := J_enG0wE{b}
	first.jf_lZ8wS2GNO.Store(&snapshot)
	first.yvcsQOnq = new(h1kE2OhgN)
	if got, e := routes.SelectOut(req); e != nil || got != b {
		t.Fatal("healthy snapshot ignored")
	}
	empty := J_enG0wE{}
	first.jf_lZ8wS2GNO.Store(&empty)
	if got := first.oImRPaEmL(req); got != a {
		t.Fatal("empty healthy list must fall back to configured list")
	}
	first.__ev4XIbr = false
	second.__ev4XIbr = false
	if got, e := routes.SelectOut(req); got != nil || e == nil || e.Error() != "not route match host: example.com, protocol: tcp" {
		t.Fatalf("no-match must error: %v %v", got, e)
	}
	routes.CnK0a4 = false
	if got, e := routes.SelectOut(nil); e != nil || got != defaultOutboundAt5029210 || !got.IsDirect() {
		t.Fatal("disabled routes must use static direct")
	}
}
