package cNT24rG4F

import (
	"recovered.local/soga/CAodbhZf"
	"reflect"
	"testing"
)

func TestListenOriginalSplitting(t *testing.T) {
	old := configAt509BEA0
	defer func() { configAt509BEA0 = old }()
	for _, tc := range []struct {
		input string
		want  []string
	}{{"", []string{""}}, {" all ", []string{"all"}}, {"ALL", []string{"ALL"}}, {" 127.0.0.1 , , ::1,", []string{"127.0.0.1", "", "::1", ""}}} {
		configAt509BEA0.VerfnYj = tc.input
		if got := EdlWFAt(); !reflect.DeepEqual(got, tc.want) {
			t.Fatal(tc, got)
		}
	}
	ips, err := CAodbhZf.N4EF0n2()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{}
	for _, ip := range ips {
		if ip.IsLoopback() || ip.IsLinkLocalUnicast() {
			t.Fatal(ip)
		}
		want = append(want, ip.String())
	}
	if len(want) == 0 {
		want = []string{""}
	}
	configAt509BEA0.VerfnYj = "all"
	if got := EdlWFAt(); !reflect.DeepEqual(got, want) {
		t.Fatal(got, want)
	}
}
