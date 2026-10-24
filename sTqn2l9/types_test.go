package sTqn2l9

import (
	"recovered.local/soga/D43Y3mB"
	"reflect"
	"testing"
)

func TestOriginalRTTILayout(t *testing.T) {
	type field struct {
		name     string
		offset   uintptr
		tag      string
		typ      reflect.Type
		embedded bool
	}
	cases := []struct {
		typ    reflect.Type
		size   uintptr
		fields []field
	}{
		{reflect.TypeFor[J5tD2c9](), 8, []field{{"GMLWI3", 0, "", reflect.TypeFor[D43Y3mB.GMLWI3](), true}}}, {reflect.TypeFor[PqISXuLk](), 32, []field{{"Jrq3niigv", 0, "json:\"uid\"", reflect.TypeFor[int64](), false}, {"Jm6azqIw", 8, "json:\"password\"", reflect.TypeFor[string](), false}, {"Bexa45zkL", 24, "json:\"speed_limit\"", reflect.TypeFor[float64](), false}}}, {reflect.TypeFor[BoM67db](), 48, []field{{"IdNatD8MGczW", 0, "json:\"code\"", reflect.TypeFor[int](), false}, {"UzQ0WhgKfU", 8, "json:\"message\"", reflect.TypeFor[string](), false}, {"GPETCf2", 24, "json:\"data\"", reflect.TypeFor[[]*PqISXuLk](), false}}}, {reflect.TypeFor[XhDo0vyZH](), 24, []field{{"JLAJhSUZPpQv", 0, "json:\"speed_limit\"", reflect.TypeFor[float64](), false}, {"JQsJVj", 8, "json:\"client_limit\"", reflect.TypeFor[int](), false}, {"Jbn82w0HZp", 16, "json:\"trojan_port\"", reflect.TypeFor[int](), false}}}, {reflect.TypeFor[WmjoeO](), 48, []field{{"IdNatD8MGczW", 0, "json:\"code\"", reflect.TypeFor[int](), false}, {"UzQ0WhgKfU", 8, "json:\"message\"", reflect.TypeFor[string](), false}, {"GPETCf2", 24, "json:\"data\"", reflect.TypeFor[XhDo0vyZH](), false}}},
	}
	for _, c := range cases {
		if c.typ.Size() != c.size || c.typ.NumField() != len(c.fields) {
			t.Errorf("%s size/field count mismatch", c.typ)
		}
		for i, f := range c.fields {
			x := c.typ.Field(i)
			if x.Name != f.name || x.Offset != f.offset || string(x.Tag) != f.tag || x.Type != f.typ || x.Anonymous != f.embedded {
				t.Errorf("%s.%s differs: %+v", c.typ, f.name, x)
			}
		}
	}
}
