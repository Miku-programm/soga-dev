package cNT24rG4F

import (
	"reflect"
	"testing"
	"time"
)

func TestAccessLogConfigurationFromOriginalLayout(t *testing.T) {
	typ := reflect.TypeFor[ZNxaPWNlDqa]()
	if typ.Size() != 96 {
		t.Fatal(typ.Size())
	}
	names := []string{"ASvYULtJ7", "ZH4VO_H4JZ", "IJFBj1lIX", "H0x74Y", "SNWlSe", "BaLySDum3K", "EVFQh90cvJK"}
	offsets := []uintptr{0, 16, 32, 48, 64, 80, 88}
	for i, n := range names {
		f := typ.Field(i)
		if f.Name != n || f.Offset != offsets[i] || f.Tag != "" {
			t.Fatalf("field %d: %+v", i, f)
		}
	}
	old := configAt50A2260
	defer func() { configAt50A2260 = old }()
	configAt50A2260 = tUl2hoOYI0Ao{QpeDOsW: " https://127.0.0.1/// ", CmPnvU: "false", G8aycKMf5aa: "false", HnvQ_1vV: "false", ACWNbHb1C: "FALSE", DOZx3XRAm01: -1, D4WJQARA: true}
	c := E7kAu2lOVtXY()
	want := ZNxaPWNlDqa{"https://127.0.0.1:8443", "default", "access_log", "", "FALSE", 1000, true}
	if *c != want {
		t.Fatalf("%+v", c)
	}
	c.IJFBj1lIX = "a`b"
	if c.FullTable() != "`default`.`a`b`" || E7kAu2lOVtXY().IJFBj1lIX != "access_log" {
		t.Fatal("quoting or snapshot mismatch")
	}
}

func TestClickhouseAddressOriginalEdgeCases(t *testing.T) {
	for _, tc := range [][2]string{
		{"", ""}, {" / ", ""}, {"http://", ""}, {"https://", ""}, {"host", "http://host:8123"},
		{"https://host/", "https://host:8443"}, {"host:9000", "http://host:9000"},
		{"host:+99999", "http://host:+99999"}, {"host:-1", "http://host:-1"}, {"host:x", ""},
		{"host:", ""}, {"2001:db8::1", "http://[2001:db8::1]:8123"},
		{"[::1]", "http://[[::1]]:8123"}, {"[::1]:80", "http://[::1]:80"},
		{"https://host/path", "https://host/path:8443"}, {"HTTP://host", ""},
	} {
		if got := bgj94i7et(tc[0]); got != tc[1] {
			t.Errorf("%q -> %q want %q", tc[0], got, tc[1])
		}
	}
	now := time.Date(2026, 10, 23, 23, 59, 59, 0, time.FixedZone("test", 8*3600))
	if O__f2DDogBIK(now) != "soga-access-2026-10-23 23:00:00.log" || O__f2DDogBIK(now.UTC()) != "soga-access-2026-10-23 15:00:00.log" {
		t.Fatal("time location or hourly bucket mismatch")
	}
}
