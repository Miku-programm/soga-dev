package MFFa9yT

import (
	"context"
	"recovered.local/soga/G6bHD5JA"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/qwn0vjA"
	"recovered.local/soga/xx5fRyrFw"
	"testing"
)

func TestOriginalBaseInitializationAndRuleReplacement(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bt, _ := qwn0vjA.NHzYZHDSd("bt")
	b := &JWN154qJ4{zb8f3cHWW8v: ctx, cEUo2gYKI: qwn0vjA.KZgBMVcpL{bt}}
	if e := b.init(); e != nil {
		t.Fatal(e)
	}
	if b.e3JHP2MzU.Load() == nil || *b.e3JHP2MzU.Load() == nil || b.oIgEuicmPL.Load() == nil || b._IoxtH.Load() == nil {
		t.Fatal("missing initial rule state")
	}
	expected := 1
	for _, name := range cNT24rG4F.DomainSniffProtocols() {
		if _, e := qwn0vjA.NHzYZHDSd(qwn0vjA.MCe3PxlaMYM(name)); e == nil {
			expected++
		}
	}
	if len(b.cEUo2gYKI) != expected || b.cEUo2gYKI[0] != bt {
		t.Fatal("existing sniffers replaced")
	}
	before := b.oIgEuicmPL.Load()
	b.v5VChcYe(xx5fRyrFw.VHWq49iRat{"full:block.test"})
	first := b.oIgEuicmPL.Load()
	_, blockMatch := first.a73gMUVz.Test("block.test")
	if before == first || !blockMatch {
		t.Fatal("block rules not replaced")
	}
	b.v5VChcYe(xx5fRyrFw.VHWq49iRat{"full:block.test"})
	if first == b.oIgEuicmPL.Load() {
		t.Fatal("identical update incorrectly cached")
	}
	b.ctArZJ(xx5fRyrFw.VHWq49iRat{"full:white.test"})
	if _, whiteMatch := b._IoxtH.Load().a73gMUVz.Test("white.test"); !whiteMatch {
		t.Fatal("white update")
	}
	if b.x1djY10H != nil || b.xxf4dimGm != nil {
		t.Fatal("update helper must not change saved lists")
	}
}

func TestOriginalNodeUpdateFirstBTAndPathRules(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	b := &JWN154qJ4{zb8f3cHWW8v: ctx, dwfzZM: &G6bHD5JA.L337vk2G{}}
	first := &LFWTTXF7JR.Wr5SVP{Z1TIgHOQ: "first", HSA2b5V: true, TtWKyh: 8}
	first.SwxJNgPpg.ChJnngzb = "a//b"
	first.SwxJNgPpg.UVeahqDBT = true
	if b.xBzM2TWz4(first) {
		t.Fatal("first assignment reported change")
	}
	if first.SwxJNgPpg.ChJnngzb != "/a//b" || b.wiy6P3.Load() != first || b.dwfzZM.GetSpeedLimit() != 8 {
		t.Fatal("node state")
	}
	if first.SwxJNgPpg.UVeahqDBT == cNT24rG4F.ForceCloseSSL() {
		t.Fatal("force close mapping")
	}
	if len(b.cEUo2gYKI) != 1 || b.cEUo2gYKI[0].Type() != "bt" {
		t.Fatal("first BT sniffer missing")
	}
	resolver := b.ffju0Enjo.Load()
	next := &LFWTTXF7JR.Wr5SVP{Z1TIgHOQ: "first", HSA2b5V: true, TtWKyh: 0}
	if b.xBzM2TWz4(next) || len(b.cEUo2gYKI) != 1 || b.ffju0Enjo.Load() != resolver || next.SwxJNgPpg.ChJnngzb != "/" || b.dwfzZM.GetSpeedLimit() != 0 {
		t.Fatal("unchanged node behavior")
	}
	changed := &LFWTTXF7JR.Wr5SVP{Z1TIgHOQ: "second"}
	if !b.xBzM2TWz4(changed) {
		t.Fatal("changed identity ignored")
	}
	later := &JWN154qJ4{zb8f3cHWW8v: ctx, dwfzZM: &G6bHD5JA.L337vk2G{}}
	later.xBzM2TWz4(&LFWTTXF7JR.Wr5SVP{})
	later.xBzM2TWz4(&LFWTTXF7JR.Wr5SVP{HSA2b5V: true})
	if len(later.cEUo2gYKI) != 0 {
		t.Fatal("later BT setting must not append")
	}
}
