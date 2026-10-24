package registry

import (
	"recovered.local/soga/panel/types"
	"recovered.local/soga/audit/matcher"
	"reflect"
	"testing"
	"time"
)

func TestPanelNodeRTTI(t *testing.T) {
	type field struct {
		name   string
		offset uintptr
		typ    reflect.Type
		tag    string
	}
	for _, tc := range []struct {
		typ    reflect.Type
		size   uintptr
		fields []field
	}{
		{reflect.TypeOf(D5gTIZ5aVg{}), 152, []field{
			{"H47AzZsYHB1N", 0, reflect.TypeFor[*types.Wr5SVP](), ""},
			{"VTme1SuHL6i_", 8, reflect.TypeFor[types.Sxcf3wu3L](), ""},
			{"HYnfYWa", 32, reflect.TypeFor[matcher.Ij9r21E0UOt3](), ""},
			{"Xdo1D9", 56, reflect.TypeFor[time.Duration](), ""},
			{"DDu2Nm_P", 64, reflect.TypeFor[time.Duration](), ""},
			{"ATGYGi0", 72, reflect.TypeFor[matcher.VHWq49iRat](), ""},
			{"BeqZvr6", 96, reflect.TypeFor[error](), ""},
			{"NjjWbjH", 112, reflect.TypeFor[matcher.VHWq49iRat](), ""},
			{"IGP6TVQazVxp", 136, reflect.TypeFor[error](), ""},
		}},
	} {
		if tc.typ.Size() != tc.size || tc.typ.NumField() != len(tc.fields) {
			t.Fatalf("RTTI大小/字段数量不符: %v", tc.typ)
		}
		for i, w := range tc.fields {
			f := tc.typ.Field(i)
			if f.Name != w.name || f.Offset != w.offset || f.Type != w.typ || string(f.Tag) != w.tag {
				t.Fatalf("RTTI字段不符: %v field %d %+v", tc.typ, i, f)
			}
		}
	}
}
