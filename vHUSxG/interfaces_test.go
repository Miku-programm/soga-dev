package vHUSxG

import (
	"net"
	"testing"
	"unsafe"
)

func subnet(s string) *net.IPNet {
	ip, n, err := net.ParseCIDR(s)
	if err != nil {
		panic(err)
	}
	n.IP = ip
	return n
}
func TestAddressSelection(t *testing.T) {
	link := subnet("fe80::1/64")
	first := subnet("2001:db8::1/64")
	second := subnet("2001:db8:2::1/64")
	v4 := subnet("192.0.2.1/24")
	a := &Ah4poD{JNv6km: []*net.IPNet{link, first, second, v4}}
	for _, test := range []struct {
		dst  net.IP
		want *net.IPNet
	}{{nil, first}, {net.ParseIP("2001:db8:2::9"), second}, {net.ParseIP("fe80::9"), link}, {net.ParseIP("2001:db8:9::9"), first}} {
		got, err := a.PickIPv6Addr(test.dst)
		if err != nil || got != test.want {
			t.Fatal(got, err)
		}
	}
	got, err := a.PickIPv4Addr(nil)
	if got != v4 || err != nil {
		t.Fatal(got, err)
	}
	a.JNv6km = []*net.IPNet{link}
	got, err = a.PickIPv6Addr(nil)
	if got != nil || err != AddrNotFoundAt5098660 {
		t.Fatal(got, err)
	}
	if unsafe.Sizeof(*a) != 72 {
		t.Fatal("RTTI layout")
	}
}
func TestInterfaceEnumerationAndCache(t *testing.T) {
	value, err := dinaz4func1()
	if err != nil {
		t.Fatal(err)
	}
	m := value.(map[string]*Ah4poD)
	if len(m) == 0 {
		t.Skip("no interfaces")
	}
	for name, iface := range m {
		got, err := DINaz4(name)
		if err != nil || got.Arz5udf != iface.Arz5udf || got.BdMyLmuSMQA != iface.BdMyLmuSMQA {
			t.Fatal(got, err)
		}
		same, err := DINaz4(name)
		if same != got || err != nil {
			t.Fatal("cache identity")
		}
		break
	}
	if _, err := DINaz4("nonexistent-soga-interface"); err != InterfaceNotFoundAt5098650 {
		t.Fatal(err)
	}
}
