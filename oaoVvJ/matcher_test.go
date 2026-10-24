package oaoVvJ

import (
	"recovered.local/soga/GYzZ5v"
	"reflect"
	"testing"
)

func TestGeositeGroupedMatchers(t *testing.T) {
	m := &HajppSMFH{}
	m.j9oqh6bhnQL([]*GYzZ5v.HTBZ_woRreQ{
		{AfECsc: 0, DSsRH80v92kO: "needle"}, {AfECsc: 0, DSsRH80v92kO: "second"},
		{AfECsc: 1, DSsRH80v92kO: "^rx[0-9]+$"}, {AfECsc: 1, DSsRH80v92kO: "["},
		{AfECsc: 2, DSsRH80v92kO: "example.com"}, {AfECsc: 3, DSsRH80v92kO: "exact.test"},
		{AfECsc: 99, DSsRH80v92kO: "unsupported"},
	})
	if len(m.hy5PyUcE8) != 4 {
		t.Fatalf("分组数量 %d", len(m.hy5PyUcE8))
	}
	for _, s := range []string{"xneedley", "rx123", "example.com", "子域.example.com", "exact.test"} {
		if !m.Match(s) {
			t.Errorf("未匹配 %q", s)
		}
	}
	for _, s := range []string{"NEEDLE", "arx123", "notexample.com", "x.exact.test", "unsupported", "["} {
		if m.Match(s) {
			t.Errorf("错误匹配 %q", s)
		}
	}
	// 再次构建替换旧规则，不累积。
	m.j9oqh6bhnQL(nil)
	if m.hy5PyUcE8 == nil || len(m.hy5PyUcE8) != 0 || m.Match("needle") {
		t.Fatal("空输入必须替换旧规则")
	}
}
func TestDomainSetBoundaryAndUnknownType(t *testing.T) {
	matcher := zQ7CNn6o6mU(2)
	matcher.AddValue("example.com")
	matcher.AddValue("example.com")
	for _, tc := range []struct {
		s    string
		want bool
	}{{"example.com", true}, {"a.example.com", true}, {"notexample.com", false}, {"example.com.", false}, {"EXAMPLE.COM", false}} {
		if matcher.Match(tc.s) != tc.want {
			t.Fatalf("%q", tc.s)
		}
	}
	if zQ7CNn6o6mU(-1) != nil {
		t.Fatal("未知类型应返回nil")
	}
}
func TestGeositeTypeLayout(t *testing.T) {
	for _, tc := range []struct {
		v     any
		size  uintptr
		field string
	}{{HajppSMFH{}, 24, "hy5PyUcE8"}, {dllnJks{}, 24, "yqsoTKORPtw"}, {enyqKW{}, 24, "sDKYrz1SI"}, {bee3oESrQ{}, 8, "jnsGsOq2H2ov"}, {hSu65g{}, 8, "jnsGsOq2H2ov"}} {
		typ := reflect.TypeOf(tc.v)
		if typ.Size() != tc.size || typ.NumField() != 1 || typ.Field(0).Name != tc.field || typ.Field(0).Offset != 0 {
			t.Fatalf("RTTI不匹配: %v", typ)
		}
	}
}
