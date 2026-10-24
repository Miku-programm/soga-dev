package DsyR31Nu

import (
	"net"
	"testing"
)

func TestNativeUDPAddrConversion(t *testing.T) {
	a := &net.UDPAddr{IP: net.ParseIP("fe80::1"), Zone: "eth0", Port: 1}
	if got, e := Q_ZqnPkUlb9(a); got != a || e != nil {
		t.Fatal(got, e)
	}
	for _, tc := range []struct {
		a    net.Addr
		host string
		port int
		fail bool
	}{
		{&net.TCPAddr{IP: net.IPv4(1, 2, 3, 4), Port: -5}, "1.2.3.4", -5, false},
		{&net.TCPAddr{IP: net.ParseIP("fe80::1"), Zone: "eth0", Port: 5}, "", 0, true},
		{&net.UnixAddr{Name: "example.org:8"}, "", 0, true},
	} {
		got, err := Q_ZqnPkUlb9(tc.a)
		if tc.fail {
			if err == nil {
				t.Fatal(got)
			}
			continue
		}
		if err != nil || got.IP.String() != tc.host || got.Port != tc.port {
			t.Fatal(got, err)
		}
	}
}
