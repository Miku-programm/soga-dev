package cNT24rG4F

import "testing"

func TestFallbackSelectionOriginalBranches(t *testing.T) {
	a, b, c, n := configAt50992E0, configAt5099360, configAt50A02C0, nodeConfigurationAt50A2AE0
	defer func() { configAt50992E0 = a; configAt5099360 = b; configAt50A02C0 = c; nodeConfigurationAt50A2AE0 = n }()
	configAt50992E0.BVi1B4f_v = "any"
	configAt50992E0.JGw3yW = -1
	configAt5099360.BYiTyYPg3 = "trojan"
	configAt5099360.EglOu0OdYex = 443
	configAt50A02C0.GnOeZP5RIl = "new"
	configAt50A02C0.HUt3s7n_8 = 80
	configAt50A02C0.RIFsRvm1 = "old"
	configAt50A02C0.YMm_xn_v = -8
	for _, tc := range []struct {
		kind, host string
		port       int
	}{{"ANYTLS", "any", -1}, {"trojan", "trojan", 443}, {"vmess", "old", 80}, {"v2ray", "old", 80}, {"vless", "", 0}, {"ss", "", 0}} {
		nodeConfigurationAt50A2AE0.MnXMlyI = tc.kind
		if G9UvxRacCj() != tc.host || QnSmghY() != tc.port {
			t.Fatal(tc, G9UvxRacCj(), QnSmghY())
		}
	}
	nodeConfigurationAt50A2AE0.MnXMlyI = "v2ray"
	configAt50A02C0.RIFsRvm1 = ""
	configAt50A02C0.YMm_xn_v = 25
	if G9UvxRacCj() != "new" || QnSmghY() != 25 {
		t.Fatal("independent old host/port fallback")
	}
}
