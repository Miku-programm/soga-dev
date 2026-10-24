package G6bHD5JA

import (
	"net"
	"testing"
)

type tcpAddrFixture string

func (a tcpAddrFixture) Network() string { return "fixture" }
func (a tcpAddrFixture) String() string  { return string(a) }
func TestTCPAddrOriginalConversion(t *testing.T) {
	a := &net.TCPAddr{IP: net.IPv6loopback, Port: 80, Zone: "zone"}
	got, e := JY6gu_Rqfz(a)
	if e != nil || got != a {
		t.Fatal(got, e)
	}
	u := &net.UDPAddr{IP: net.IPv6loopback, Port: 81, Zone: "zone"}
	got, e = JY6gu_Rqfz(u)
	if e != nil || got.Zone != "" || &got.IP[0] != &u.IP[0] {
		t.Fatal(got, e)
	}
	for _, s := range []string{"host.test:80", "[fe80::1%eth0]:80", "127.0.0.1:bad", "no-port"} {
		if _, e = JY6gu_Rqfz(tcpAddrFixture(s)); e == nil {
			t.Fatal(s)
		}
	}
	got, e = JY6gu_Rqfz(tcpAddrFixture("127.0.0.1:-7"))
	if e != nil || got.Port != -7 || len(got.IP) != 16 {
		t.Fatal(got, e)
	}
}
