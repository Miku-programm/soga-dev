package routing

import (
	"net"
	"testing"
)

func TestSourceFamilyPreferenceAndFallback(t *testing.T) {
	v4, v6 := net.ParseIP("192.0.2.10"), net.ParseIP("2001:db8::10")
	r4, r6 := &net.TCPAddr{IP: net.ParseIP("192.0.2.20"), Port: 80}, &net.UDPAddr{IP: net.ParseIP("2001:db8::20"), Port: 80}
	out := &GyyO51{iCySnjtd: v4, xk9pAT: v6, JaH5hv: "ipv6_first"}
	source, remote, err := out.w4foeAH7FyI(nil, []net.Addr{r4, r6})
	if err != nil || !source.Equal(v6) || remote != r6 {
		t.Fatalf("v6 preference %v %v %v", source, remote, err)
	}
	out.JaH5hv = "IPV6_FIRST"
	source, remote, err = out.w4foeAH7FyI(nil, []net.Addr{r6, r4})
	if err != nil || !source.Equal(v4) || remote != r4 {
		t.Fatal("default v4 preference lost")
	}
	source, remote, err = out.w4foeAH7FyI(nil, nil)
	if err != nil || !source.Equal(v4) || remote != nil {
		t.Fatal("no destination must retain first source")
	}
	out.xk9pAT = nil
	source, remote, err = out.w4foeAH7FyI(nil, []net.Addr{r6})
	if err != nil || !source.Equal(v4) || remote != nil {
		t.Fatal("family mismatch must preserve source with nil destination")
	}
	out.iCySnjtd = net.IPv4(127, 0, 0, 1)
	source, remote, err = out.w4foeAH7FyI(nil, []net.Addr{r6})
	if err != nil || source != nil || remote != r6 {
		t.Fatal("loopback source must be filtered")
	}
	out.iCySnjtd = net.IPv4zero
	source, remote, err = out.w4foeAH7FyI(nil, []net.Addr{r6})
	if err != nil || !source.Equal(net.IPv4zero) || remote != r6 {
		t.Fatal("unspecified source must accept either family")
	}
}
func TestServerLiteralOnly(t *testing.T) {
	out := &GyyO51{AHTlIEynXL: "localhost"}
	if out.vhKZu6vCmJ_() != nil {
		t.Fatal("server name must not trigger DNS here")
	}
	out.AHTlIEynXL = "2001:db8::1"
	got := out.vhKZu6vCmJ_()
	if len(got) != 1 || got[0].(*net.TCPAddr).Port != 0 || !got[0].(*net.TCPAddr).IP.Equal(net.ParseIP(out.AHTlIEynXL)) {
		t.Fatal("literal server candidate changed")
	}
}
