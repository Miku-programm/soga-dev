package OieaM9

import (
	"recovered.local/soga/D43Y3mB"
	"recovered.local/soga/hfOuV8fOe4L"
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
		{reflect.TypeFor[HbUGOX](), 8, []field{{"GMLWI3", 0, "", reflect.TypeFor[D43Y3mB.GMLWI3](), true}}}, {reflect.TypeFor[JpYAk64U2](), 48, []field{{"IdNatD8MGczW", 0, "json:\"code\"", reflect.TypeFor[int](), false}, {"UzQ0WhgKfU", 8, "json:\"message\"", reflect.TypeFor[string](), false}, {"GPETCf2", 24, "json:\"data\"", reflect.TypeFor[[]*CXFPuaTkSm](), false}}}, {reflect.TypeFor[TEsbPj5T](), 128, []field{{"IdNatD8MGczW", 0, "json:\"code\"", reflect.TypeFor[int](), false}, {"UzQ0WhgKfU", 8, "json:\"message\"", reflect.TypeFor[string](), false}, {"GPETCf2", 24, "json:\"data\"", reflect.TypeFor[QVdbsbqVT3](), false}}}, {reflect.TypeFor[CXFPuaTkSm](), 40, []field{{"QZuKzrg", 0, "json:\"uid\"", reflect.TypeFor[int64](), false}, {"SXYRZj", 8, "json:\"port\"", reflect.TypeFor[int](), false}, {"JaWN6wQMY_8k", 16, "json:\"passwd\"", reflect.TypeFor[string](), false}, {"T1hUUd2", 32, "json:\"speed_limit\"", reflect.TypeFor[float64](), false}}}, {reflect.TypeFor[QVdbsbqVT3](), 104, []field{{"AZlE6hE7M7", 0, "json:\"port\"", reflect.TypeFor[string](), false}, {"UGE6Z0a", 16, "json:\"passwd\"", reflect.TypeFor[string](), false}, {"G4HJNQ1", 32, "json:\"method\"", reflect.TypeFor[hfOuV8fOe4L.ZvAxOt1l7Ar](), false}, {"DXDUf20", 48, "json:\"protocol\"", reflect.TypeFor[hfOuV8fOe4L.Mu0vMIIs](), false}, {"SCImeBpi43", 64, "json:\"obfs\"", reflect.TypeFor[hfOuV8fOe4L.PMorzD7u](), false}, {"DR80QsI", 80, "json:\"speed_limit\"", reflect.TypeFor[float64](), false}, {"K9N3GWPMdM3", 88, "json:\"client_limit\"", reflect.TypeFor[int](), false}, {"C8gP2P", 96, "json:\"single\"", reflect.TypeFor[int](), false}}},
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
