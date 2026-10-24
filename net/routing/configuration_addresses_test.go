package routing

import (
	"net"
	"testing"
)

func TestOutboundAddressValidationOrder(t *testing.T) {
	for _, tc := range []struct {
		kind   Fdaw35mqIz
		server string
		port   int
		want   string
	}{
		{"ss", "", 0, "please set out server"},
		{"ss", "localhost", 0, "please set out port"},
		{"ss", "localhost", -1, "out port invalid: -1"},
		{"ss", "localhost", 65536, "out port invalid: 65536"},
		{"ss", "localhost", 65535, ""},
		{"redirect", "localhost", 0, ""},
		{"redirect", "localhost", -1, "out port invalid: -1"},
		{"direct", "", -1, ""},
	} {
		out := &GyyO51{NyLbiXb4OrlL: tc.kind, AHTlIEynXL: tc.server, JO1YIpXk: tc.port, JaH5hv: "ipv4_first"}
		err := out.initializeAddresses()
		if tc.want == "" {
			if err != nil {
				t.Fatal(err)
			}
		} else if err == nil || err.Error() != tc.want {
			t.Fatalf("%+v: %v", tc, err)
		}
	}
	for _, strategy := range []U0f6rIw7p{"", "IPV4_FIRST", "ipv4_only"} {
		out := &GyyO51{NyLbiXb4OrlL: "direct", JaH5hv: strategy}
		if err := out.initializeAddresses(); err == nil || err.Error() != "unknown listen_strategy: "+string(strategy) {
			t.Fatalf("%q: %v", strategy, err)
		}
	}
}

func TestListenFamiliesAndPartialMutation(t *testing.T) {
	out := &GyyO51{NyLbiXb4OrlL: "direct", JaH5hv: "ipv6_first", DM8KdAE6ZiZu: "192.0.2.1", LXpap7F: "192.0.2.2"}
	err := out.initializeAddresses()
	if err == nil || err.Error() != "listen_ipv6 is not ipv6: 192.0.2.2" {
		t.Fatal(err)
	}
	// The original commits the IPv4 field before rejecting the IPv6 field.
	if out.ZK61Na2nsqVQ != "192.0.2.1" || !out.iCySnjtd.Equal(net.ParseIP("192.0.2.1")) || len(out.iCySnjtd) != 16 {
		t.Fatal("lost partial address update")
	}
	out.LXpap7F = "2001:db8::1"
	if err := out.initializeAddresses(); err != nil {
		t.Fatal(err)
	}
	out.DM8KdAE6ZiZu = ""
	out.ZK61Na2nsqVQ = ""
	out.LXpap7F = ""
	if err := out.initializeAddresses(); err != nil {
		t.Fatal(err)
	}
	if out.iCySnjtd == nil || out.xk9pAT == nil {
		t.Fatal("empty fields must retain existing parsed addresses")
	}
	out.DM8KdAE6ZiZu = "localhost"
	if err := out.initializeAddresses(); err == nil || err.Error() != "listen is not ip: localhost" {
		t.Fatal(err)
	}
	out.DM8KdAE6ZiZu = "2001:db8::9"
	out.ZK61Na2nsqVQ = "2001:db8::8"
	if err := out.initializeAddresses(); err == nil || err.Error() != "listen_ipv4 is not ipv4: 2001:db8::8" {
		t.Fatal(err)
	}
	if out.LXpap7F != "2001:db8::9" {
		t.Fatal("IPv6 listen inference not retained on later failure")
	}
}
