package config

import (
	"reflect"
	"testing"
)

func TestPanelAPIConfigRTTI(t *testing.T) {
	rt := reflect.TypeOf(FCEAfQJwDA7e{})
	if rt.Size() != 160 || rt.NumField() != 11 {
		t.Fatalf("layout: %v/%d", rt, rt.Size())
	}
	if f := rt.Field(0); f.Name != "OS91X026uML" || f.Offset != 0 || f.Type.String() != "config.FSOZSH7g" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 0: %#v", f)
	}
	if f := rt.Field(1); f.Name != "ZyAhkq5th8bf" || f.Offset != 16 || f.Type.String() != "config.Mpn2JltM" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 1: %#v", f)
	}
	if f := rt.Field(2); f.Name != "F3PCiaIra" || f.Offset != 32 || f.Type.String() != "config.HWLVDK2" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 2: %#v", f)
	}
	if f := rt.Field(3); f.Name != "W_ogGtF" || f.Offset != 48 || f.Type.String() != "string" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 3: %#v", f)
	}
	if f := rt.Field(4); f.Name != "Bd08Li" || f.Offset != 64 || f.Type.String() != "string" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 4: %#v", f)
	}
	if f := rt.Field(5); f.Name != "Yf1KqSqJpNI7" || f.Offset != 80 || f.Type.String() != "int" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 5: %#v", f)
	}
	if f := rt.Field(6); f.Name != "EotADOuOVu" || f.Offset != 88 || f.Type.String() != "string" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 6: %#v", f)
	}
	if f := rt.Field(7); f.Name != "Z9qKfN7e_" || f.Offset != 104 || f.Type.String() != "int" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 7: %#v", f)
	}
	if f := rt.Field(8); f.Name != "XU27D_AD" || f.Offset != 112 || f.Type.String() != "string" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 8: %#v", f)
	}
	if f := rt.Field(9); f.Name != "GcBQVp1f" || f.Offset != 128 || f.Type.String() != "string" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 9: %#v", f)
	}
	if f := rt.Field(10); f.Name != "JWhomGW" || f.Offset != 144 || f.Type.String() != "string" || f.Tag != "" || f.Anonymous {
		t.Fatalf("field 10: %#v", f)
	}
}
