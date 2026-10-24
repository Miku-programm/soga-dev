package routing

import (
	"context"
	"net"
	"testing"
	"time"
)

func TestHealthProbeAndStateWithRealTCP(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	accepted := make(chan struct{})
	go func() {
		defer close(accepted)
		for i := 0; i < 2; i++ {
			c, e := listener.Accept()
			if e != nil {
				return
			}
			c.Close()
		}
	}()
	addr := listener.Addr().(*net.TCPAddr)
	out := &GyyO51{AHTlIEynXL: addr.IP.String(), JO1YIpXk: addr.Port}
	manager := &Qi1aOz5oQ{bXx0sIP: context.Background()}
	if manager.nvI5aMbq(out) || out.IsEnabled() {
		t.Fatal("one success must not enable")
	}
	if !manager.nvI5aMbq(out) || !out.IsEnabled() {
		t.Fatal("second success must enable")
	}
	<-accepted
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	manager.bXx0sIP = canceled
	for i := 0; i < 2; i++ {
		if manager.nvI5aMbq(out) {
			t.Fatal("early disable")
		}
	}
	if !manager.nvI5aMbq(out) || out.IsEnabled() {
		t.Fatal("third failure must disable")
	}
}
func TestHealthCancellationAndEligibility(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	a := &GyyO51{AHTlIEynXL: "192.0.2.1", JO1YIpXk: 1}
	a.uymza6wOIV6()
	direct := &GyyO51{NyLbiXb4OrlL: "direct", AHTlIEynXL: "192.0.2.2", JO1YIpXk: 1}
	rule := &DgAjL9UN{UwkgrlFf8: J_enG0wE{a, a, direct}}
	manager := &Qi1aOz5oQ{bXx0sIP: ctx, b0Ftnk: &UIfnUsim{MMKbp60np: ViF0a1nppLG{rule}}}
	manager.qPMo12()
	if a.pJpTed != 0 || rule.jf_lZ8wS2GNO.Load() != nil {
		t.Fatal("canceled jitter must not probe or publish")
	}
	done := make(chan struct{})
	go func() { manager.Run(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("manager did not cancel")
	}
}
func TestHealthManagerReplacement(t *testing.T) {
	tCGadgsq(nil)
	t.Cleanup(func() { tCGadgsq(nil) })
	routes := &UIfnUsim{CnK0a4: true, H839j9DYDD5k: true}
	tCGadgsq(routes)
	first := healthManagerAt5095C60
	if first == nil {
		t.Fatal("health manager not installed")
	}
	tCGadgsq(&UIfnUsim{})
	if healthManagerAt5095C60 != nil {
		t.Fatal("disabled config kept manager")
	}
	select {
	case <-first.bXx0sIP.Done():
	default:
		t.Fatal("old manager not canceled")
	}
}
func TestHealthySnapshotAndDisplayName(t *testing.T) {
	a, b := new(GyyO51), new(GyyO51)
	a.uymza6wOIV6()
	rule := &DgAjL9UN{UwkgrlFf8: J_enG0wE{b, a, a}}
	rule.eEYGK2LJf()
	first := rule.jf_lZ8wS2GNO.Load()
	if first == nil || len(*first) != 2 || (*first)[0] != a {
		t.Fatal("snapshot lost order or duplicates")
	}
	b.uymza6wOIV6()
	rule.eEYGK2LJf()
	if len(*first) != 2 || len(*rule.jf_lZ8wS2GNO.Load()) != 3 {
		t.Fatal("snapshot mutated in place")
	}
	a.NyLbiXb4OrlL = "direct"
	if a.l9fABxN() != "direct" {
		t.Fatal("type fallback")
	}
	a.AHTlIEynXL = "2001:db8::1"
	a.JO1YIpXk = 443
	if a.l9fABxN() != "[2001:db8::1]:443" {
		t.Fatal("IPv6 display")
	}
	a.HG8A97UIPf1A = "named"
	if a.l9fABxN() != "named" {
		t.Fatal("name precedence")
	}
}
