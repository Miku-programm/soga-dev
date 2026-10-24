package cNT24rG4F

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestDynamicConfigurationDefaultsAndWhitelist(t *testing.T) {
	old, oldResult := configAt509E2A0, dynamicConfigurationAt5095AB0
	t.Cleanup(func() { configAt509E2A0, dynamicConfigurationAt5095AB0 = old, oldResult })
	configAt509E2A0 = rXC9QSbKAoV0{EFqkTZF6Wdql: true, IrsBkXPlx: []string{"1", "-2", "1", "bad", " 3", "9223372036854775808"}}
	value, err := oatbDNUZrild()
	if err != nil {
		t.Fatal(err)
	}
	if !value.DiN7MxyB7 || value.BgrTZk.String() != "00:00-24:00" || value.Vs46JphL != time.Minute || value.Yt1TMC2 != 100 || value.M6lj3M4NlG != 30 || value.ZxdryEoK != 10*time.Minute {
		t.Fatalf("default conversion differs: %#v", value)
	}
	if value.GK1kRuhBCq.Len() != 2 || !value.GK1kRuhBCq.Has(-2) || value.GK1kRuhBCq.Has(3) {
		t.Fatal("whitelist parsing differs")
	}
	vlfjWv1b4Bbg()
	if dynamicConfigurationAt5095AB0 == nil || dynamicConfigurationAt5095AB0 == value {
		t.Fatal("callback must install a newly built configuration")
	}
	configAt509E2A0.CGGYGi0 = "00:00-01:00,bad"
	if result, err := oatbDNUZrild(); result != nil || err == nil {
		t.Fatal("invalid duration must return nil and error")
	}
}

func TestDynamicConfigurationNumericEdges(t *testing.T) {
	old := configAt509E2A0
	t.Cleanup(func() { configAt509E2A0 = old })
	for _, tc := range []struct {
		seconds int
		want    time.Duration
	}{{-1, time.Minute}, {0, time.Minute}, {1, 5 * time.Second}, {5, 5 * time.Second}, {6, 6 * time.Second}, {math.MaxInt64, time.Minute}} {
		configAt509E2A0 = rXC9QSbKAoV0{PLAo2k09hxd: tc.seconds, ILhRYq68NSUC: math.NaN(), LxiE3fm4BLS: math.Inf(1), Zao5joI9JY: math.MaxInt64}
		got, err := oatbDNUZrild()
		if err != nil || got.Vs46JphL != tc.want || got.ZxdryEoK != 10*time.Minute || !math.IsNaN(got.Yt1TMC2) || !math.IsInf(got.M6lj3M4NlG, 1) {
			t.Fatalf("seconds %d: %#v %v", tc.seconds, got, err)
		}
	}
	typ := reflect.TypeOf(HwIE6xOrX1OH{})
	if typ.Size() != 72 {
		t.Fatal("runtime type size differs")
	}
	for i, want := range []uintptr{0, 8, 32, 40, 48, 56, 64} {
		if typ.Field(i).Offset != want {
			t.Errorf("field %d offset differs", i)
		}
	}
}
