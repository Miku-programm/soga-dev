package realitykdf

import (
	"reflect"
	"testing"
	"time"
)

func TestParseObservedBoundaries(t *testing.T) {
	for _, tc := range []struct {
		text string
		want D1Teq5S1f
	}{
		{"20:00-06:00", D1Teq5S1f{6, 0, 20, 0}},
		{"0:00-24:00", D1Teq5S1f{0, 0, 24, 0}},
		{"12:59-12:01", D1Teq5S1f{12, 1, 12, 59}},
		{"24:00-24:00", D1Teq5S1f{24, 0, 24, 0}},
		{"\u2003+06 : 05 - 8 : 9\t", D1Teq5S1f{6, 5, 8, 9}},
		{"0008:0009-08:09", D1Teq5S1f{8, 9, 8, 9}},
	} {
		got, ok := DJr_rBzh(tc.text)
		if !ok || got == nil || *got != tc.want {
			t.Errorf("parse(%q) = %v, %v; want %v, true", tc.text, got, ok, tc.want)
		}
	}
	for _, text := range []string{
		"", "0:00", "0:00-", ":00-1:00", "0:-1:00", "24:01-0:00", "25:00-0:00",
		"0:60-1:00", "0:00-24:01", "0:00-25:00", "0:00-1:60", "-1:00-2:00",
		"0:00-1:00-2:00", "0:00:00-1:00", "０:00-1:00", "0x1:00-2:00",
		"0:00-9223372036854775808:00", "0:00-1:00 # x",
	} {
		if got, ok := DJr_rBzh(text); ok || got != nil {
			t.Errorf("invalid %q accepted: %v, %v", text, got, ok)
		}
	}
}

func TestTimeUsesLocalMinuteAndOpenEnd(t *testing.T) {
	r := &D1Teq5S1f{6, 15, 20, 30}
	zone := time.FixedZone("UTC+8", 8*3600)
	for _, tc := range []struct {
		h, m, s int
		want    bool
	}{
		{6, 14, 59, false}, {6, 15, 0, true}, {6, 15, 59, true},
		{20, 29, 59, true}, {20, 30, 0, false}, {23, 0, 0, false},
	} {
		when := time.Date(2026, 10, 23, tc.h, tc.m, tc.s, 0, zone)
		if got := r.IsInTime(when); got != tc.want {
			t.Errorf("time %s: got %v, want %v", when, got, tc.want)
		}
	}
	when := time.Date(2026, 10, 23, 6, 15, 0, 0, zone)
	if r.IsInTime(when.UTC()) {
		t.Fatal("must use the argument's time zone")
	}
}

func TestCrossObservedSpecialCases(t *testing.T) {
	for _, tc := range []struct {
		a, b D1Teq5S1f
		want bool
	}{
		{D1Teq5S1f{6, 0, 8, 0}, D1Teq5S1f{8, 0, 10, 0}, false},
		{D1Teq5S1f{6, 0, 8, 0}, D1Teq5S1f{7, 59, 10, 0}, true},
		{D1Teq5S1f{6, 0, 8, 0}, D1Teq5S1f{7, 0, 7, 0}, true},
		{D1Teq5S1f{6, 0, 8, 0}, D1Teq5S1f{8, 0, 8, 0}, false},
		{D1Teq5S1f{7, 0, 7, 0}, D1Teq5S1f{7, 0, 7, 0}, false},
		{D1Teq5S1f{6, 15, 6, 30}, D1Teq5S1f{6, 20, 6, 25}, true},
	} {
		if got := tc.a.HasCross(&tc.b); got != tc.want {
			t.Errorf("%v vs %v = %v, want %v", tc.a, tc.b, got, tc.want)
		}
		if got := tc.b.HasCross(&tc.a); got != tc.want {
			t.Errorf("reverse %v vs %v = %v, want %v", tc.b, tc.a, got, tc.want)
		}
	}
}

func TestListEmptyAndAnyMatch(t *testing.T) {
	var empty ITMfakhJ
	ranges := ITMfakhJ{&D1Teq5S1f{0, 0, 6, 0}, &D1Teq5S1f{20, 0, 24, 0}}
	if !empty.IsInTime(time.Time{}) || !empty.HasCross(ranges) || !empty.HasCross(nil) {
		t.Fatal("empty receiver must return true")
	}
	if ranges.HasCross(empty) {
		t.Fatal("nonempty receiver with empty argument must return false")
	}
	if !ranges.HasCross(ITMfakhJ{&D1Teq5S1f{23, 0, 24, 0}}) {
		t.Fatal("must examine later elements")
	}
	if !ranges.IsInTime(time.Date(2026, 1, 1, 23, 0, 0, 0, time.UTC)) ||
		ranges.IsInTime(time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)) {
		t.Fatal("list time membership incorrect")
	}
}

func TestOriginalTypeLayout(t *testing.T) {
	typ := reflect.TypeOf(D1Teq5S1f{})
	want := []struct{ name, tag string }{
		{"LXcSWg26p", `json:"min_hour"`}, {"RBNUY7G", `json:"min_minute"`},
		{"BZwEA15n", `json:"max_hour"`}, {"G1LUbU", `json:"max_minute"`},
	}
	if typ.Size() != 32 || typ.NumField() != 4 {
		t.Fatal("amd64 type layout differs from RTTI")
	}
	for i, expected := range want {
		f := typ.Field(i)
		if f.Name != expected.name || string(f.Tag) != expected.tag || f.Offset != uintptr(i*8) || f.Type.Kind() != reflect.Int {
			t.Errorf("field %d differs from RTTI: %v", i, f)
		}
	}
}
