package config

import (
	"reflect"
	"testing"
)

func TestProtocolNames(t *testing.T) {
	saved := nodeConfigurationAt50A2AE0
	t.Cleanup(func() { nodeConfigurationAt50A2AE0 = saved })
	for _, tc := range []struct{ input, want string }{
		{"VMESS", "v2ray"}, {"V2RAY", "v2ray"}, {"Shadowsocks", "ss"},
		{"SS", "ss"}, {"SHADOWSOCKSR", "ssr"}, {"SSR", "ssr"},
		{"Trojan", "trojan"}, {"VLESS", "vless"}, {"Hysteria", "hysteria"},
		{"MIERU", "mieru"}, {"AnyTLS", "anytls"}, {"HYSTERIA2", "hysteria2"},
		{" VMESS ", " vmess "}, {"CustomÄ", "customä"}, {"", ""},
	} {
		nodeConfigurationAt50A2AE0.MnXMlyI = tc.input
		if got := DnhHoT2aRvs(); got != tc.want {
			t.Errorf("normalize %q = %q, want %q", tc.input, got, tc.want)
		}
		if nodeConfigurationAt50A2AE0.MnXMlyI != tc.input {
			t.Fatal("normalization must not write back to the global")
		}
	}
}

func TestNodeIDsAndSelectedCaseConversion(t *testing.T) {
	saved := nodeConfigurationAt50A2AE0
	t.Cleanup(func() { nodeConfigurationAt50A2AE0 = saved })
	nodeConfigurationAt50A2AE0 = i7BY8ClkrDX{
		EtRWBSEhLQc: "XIAOV2BOARD", AG_mSFklcvj: "TypE", NJSpG6: "WEBAPI",
		MnXMlyI: "VMESS", VfiX6TsWZxF: "https://EXAMPLE.test/Case",
		N1Z81aDzg: []string{"1", "-2", "+3", " 4", "05", "1", "x", "9223372036854775808"},
		t09QzN:    []int{99},
	}
	jkMcB1_()
	c := &nodeConfigurationAt50A2AE0
	if c.EtRWBSEhLQc != "xiaov2board" || c.AG_mSFklcvj != "type" || c.NJSpG6 != "webapi" {
		t.Fatal("selected fields must be lowercased")
	}
	if c.MnXMlyI != "VMESS" || c.VfiX6TsWZxF != "https://EXAMPLE.test/Case" {
		t.Fatal("other fields must be retained")
	}
	if !reflect.DeepEqual(c.t09QzN, []int{1, -2, 3, 5, 1}) || cap(c.t09QzN) != len(c.N1Z81aDzg) {
		t.Fatalf("node IDs differ: %v (cap %d)", c.t09QzN, cap(c.t09QzN))
	}
	c.N1Z81aDzg = nil
	jkMcB1_()
	if c.t09QzN == nil || len(c.t09QzN) != 0 {
		t.Fatal("empty input must replace the old result with a non-nil empty slice")
	}
}
