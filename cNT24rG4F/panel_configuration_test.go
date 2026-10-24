package cNT24rG4F

import "testing"

func TestPanelConfigurationCopiesNodesAndPreservesURL(t *testing.T) {
	old := nodeConfigurationAt50A2AE0
	t.Cleanup(func() { nodeConfigurationAt50A2AE0 = old })
	nodeConfigurationAt50A2AE0 = i7BY8ClkrDX{
		EtRWBSEhLQc: "xiaov2board", MnXMlyI: "VMess", NJSpG6: "webapi",
		VfiX6TsWZxF: "https://example.invalid/base?x=1", UnkaBGBuOUXV: "new", MKrU9XEJ: "old",
		ApfXLZLmFXh: "db", Mk5cDiQzaj: 3307, CsBHZVUxTO: "name", ZGaSmv: "user", ASSy13OjR6aJ: "password",
		t09QzN: []int{7, -2, 7},
	}
	got := SGx0HPPP()
	if len(got) != 3 || got[0] == got[2] {
		t.Fatalf("independent nodes: %#v", got)
	}
	for i, c := range got {
		if c.Yf1KqSqJpNI7 != []int{7, -2, 7}[i] || c.OS91X026uML != "xiaov2board" || c.ZyAhkq5th8bf != "v2ray" || c.F3PCiaIra != "webapi" || c.W_ogGtF != "https://example.invalid/base?x=1/" || c.Bd08Li != "new" || c.EotADOuOVu != "db" || c.Z9qKfN7e_ != 3307 || c.XU27D_AD != "name" || c.GcBQVp1f != "user" || c.JWhomGW != "password" {
			t.Fatalf("node %d: %#v", i, c)
		}
	}
	got[0].Bd08Li = "changed"
	if got[2].Bd08Li != "new" || nodeConfigurationAt50A2AE0.UnkaBGBuOUXV != "new" {
		t.Fatal("configuration aliases node")
	}
	if !SIZIzCa() || KvDEtVwdr1rT() {
		t.Fatal("VMess canonical type")
	}
	nodeConfigurationAt50A2AE0.MnXMlyI = "VLESS"
	if SIZIzCa() || !KvDEtVwdr1rT() {
		t.Fatal("VLESS canonical type")
	}
	nodeConfigurationAt50A2AE0.UnkaBGBuOUXV = ""
	if H7eU0uGXU() != "old" {
		t.Fatal("legacy key fallback")
	}
	nodeConfigurationAt50A2AE0.VfiX6TsWZxF = ""
	if SGx0HPPP()[0].W_ogGtF != "/" {
		t.Fatal("empty URL suffix")
	}
	nodeConfigurationAt50A2AE0.t09QzN = nil
	if SGx0HPPP() != nil {
		t.Fatal("zero nodes must retain nil slice")
	}
}
