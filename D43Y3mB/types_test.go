package D43Y3mB

import (
	"recovered.local/soga/cNT24rG4F"
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
		{reflect.TypeFor[GMLWI3](), 8, []field{{"TMfNUn", 0, "", reflect.TypeFor[*cNT24rG4F.FCEAfQJwDA7e](), false}}}, {reflect.TypeFor[EIx7G7QCW](), 24, []field{{"NIg3b55LCOZ", 0, "json:\"code\"", reflect.TypeFor[int](), false}, {"TKIFA83m_r", 8, "json:\"message\"", reflect.TypeFor[string](), false}}}, {reflect.TypeFor[IJeD7r5Uw_7_](), 40, []field{{"CYPDxL3K", 0, "json:\"mode\"", reflect.TypeFor[string](), false}, {"NPpVYG5M9", 16, "json:\"rules\"", reflect.TypeFor[[]*Lv22mM8eYq](), false}}}, {reflect.TypeFor[ucFGrGVJeia](), 24, []field{{"M9gh81OOM4ca", 0, "json:\"uid\"", reflect.TypeFor[int64](), false}, {"A6Zb5j5", 8, "json:\"ip\"", reflect.TypeFor[string](), false}}}, {reflect.TypeFor[Lv22mM8eYq](), 40, []field{{"DPx7K7u", 0, "json:\"id\"", reflect.TypeFor[int](), false}, {"USQvAO4WtdOX", 8, "json:\"type\"", reflect.TypeFor[string](), false}, {"LgZENghw", 24, "json:\"pattern\"", reflect.TypeFor[string](), false}}}, {reflect.TypeFor[Y4Pw3Ojd6HO](), 64, []field{{"IdNatD8MGczW", 0, "json:\"code\"", reflect.TypeFor[int](), false}, {"UzQ0WhgKfU", 8, "json:\"message\"", reflect.TypeFor[string](), false}, {"GPETCf2", 24, "json:\"data\"", reflect.TypeFor[IJeD7r5Uw_7_](), false}}}, {reflect.TypeFor[ksuCP0Hi](), 32, []field{{"CYhv312VDo", 0, "json:\"uid\"", reflect.TypeFor[int64](), false}, {"F5beiROWvEE", 8, "json:\"rule_id\"", reflect.TypeFor[int](), false}, {"L6tA4h_", 16, "json:\"reason\"", reflect.TypeFor[string](), false}}}, {reflect.TypeFor[ymId_GfJFS](), 24, []field{{"W0KlV7rEtxgz", 0, "json:\"uid\"", reflect.TypeFor[int64](), false}, {"QavUxYiaO", 8, "json:\"upload\"", reflect.TypeFor[int64](), false}, {"Vdf1KaaiPt", 16, "json:\"download\"", reflect.TypeFor[int64](), false}}},
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
