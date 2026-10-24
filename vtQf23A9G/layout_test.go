package vtQf23A9G

import (
	"reflect"
	"testing"
)

// 原样本 all_types.json 两个 RTTI 结构的全部大小/字段名/偏移/tag，不由恢复源码生成。
func TestOriginalRTTIWholeLayouts(t *testing.T) {
	for _, tc := range []struct {
		v      any
		size   uintptr
		fields []struct {
			name   string
			offset uintptr
			size   uintptr
			tag    string
		}
	}{
		{BIkgBMz3M{}, 192, []struct {
			name   string
			offset uintptr
			size   uintptr
			tag    string
		}{
			{"c740FZf", 0, 16, ""},
			{"a5L1EeV4yvaw", 16, 8, ""},
			{"mW8PvJYxYup9", 24, 8, ""},
			{"iX5FOl", 32, 4, ""},
			{"ukg_2xDnH9s", 36, 24, ""},
			{"pGT2EZ20", 60, 12, ""},
			{"sF3ka3A2dYX", 72, 8, ""},
			{"twJTEeJ0S", 80, 8, ""},
			{"c5NUXMoBI", 88, 8, ""},
			{"wAiQ9IC2UZ", 96, 8, ""},
			{"xZ8dRw2_KR", 104, 8, ""},
			{"re4cr8xRiVdH", 112, 24, ""},
			{"bmjW7a", 136, 8, ""},
			{"gAJaebL", 144, 1, ""},
			{"pNuGhHxIQdo", 145, 1, ""},
			{"rVguB66", 146, 1, ""},
			{"d0Ze8qja", 147, 1, ""},
			{"vV_K36", 152, 24, ""},
			{"wLYsfs", 176, 4, ""},
			{"ulr87QdD", 184, 8, ""},
		}},
		{CCXLCurYiT7{}, 112, []struct {
			name   string
			offset uintptr
			size   uintptr
			tag    string
		}{
			{"y7z3pbAl", 0, 4, ""},
			{"hnbi8qavj6Y", 8, 8, ""},
			{"k_6JEbXwYJY", 16, 8, ""},
			{"z8qCZQcE", 24, 8, ""},
			{"ko_ou22dHrhf", 32, 24, ""},
			{"yTV1rdHC", 56, 12, ""},
			{"kCaaEFymiBDr", 72, 8, ""},
			{"yKgxtGJV", 80, 16, ""},
			{"bo7peua2l", 96, 12, ""},
		}},
	} {
		typ := reflect.TypeOf(tc.v)
		if typ.Size() != tc.size || typ.NumField() != len(tc.fields) {
			t.Fatalf("%s size=%d fields=%d", typ, typ.Size(), typ.NumField())
		}
		for i, want := range tc.fields {
			f := typ.Field(i)
			if f.Name != want.name || f.Offset != want.offset || f.Type.Size() != want.size || string(f.Tag) != want.tag {
				t.Fatalf("%s field%d=%+v want=%+v", typ, i, f, want)
			}
		}
	}
	if reflect.TypeFor[cNZriK]().Size() != 7 {
		t.Fatal("header layout")
	}
}
