package G6bHD5JA

import (
	"net"
	"testing"
)

func TestOriginalIPMaskPrefixAndAliasing(t *testing.T) {
	ip := net.ParseIP("192.0.2.137")
	full := HYeSYzPzN(ip, 32, 128)
	if len(full) != 16 || &full[0] != &ip[0] {
		t.Fatal("full prefix must retain original mapped slice")
	}
	masked := HYeSYzPzN(ip, 24, 64)
	if masked.String() != "192.0.2.0" || len(masked) != 4 || ip.String() != "192.0.2.137" {
		t.Fatal(masked, ip)
	}
	if HYeSYzPzN(ip, -1, 128) != nil || HYeSYzPzN(ip, 33, 128) != nil {
		t.Fatal("invalid v4 prefix")
	}
	v6 := net.ParseIP("2001:db8:abcd:1234::12")
	if got := HYeSYzPzN(v6, 32, 48).String(); got != "2001:db8:abcd::" {
		t.Fatal(got)
	}
	if HYeSYzPzN(v6, 32, 129) != nil {
		t.Fatal("invalid v6 prefix")
	}
	invalid := net.IP{1, 2, 3}
	if got := HYeSYzPzN(invalid, 32, 128); len(got) != 3 || &got[0] != &invalid[0] {
		t.Fatal(got)
	}
}
