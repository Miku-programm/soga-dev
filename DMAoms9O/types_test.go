package DMAoms9O

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
		{reflect.TypeFor[UqRMl2W](), 8, []field{{"GMLWI3", 0, "", reflect.TypeFor[D43Y3mB.GMLWI3](), true}}}, {reflect.TypeFor[Pn5vgsrEm5](), 32, []field{{"EgeihzZFfj", 0, "json:\"uid\"", reflect.TypeFor[int64](), false}, {"Z6TxVE4XB", 8, "json:\"vmess_uid\"", reflect.TypeFor[string](), false}, {"Z_Y4R2Rbg", 24, "json:\"speed_limit\"", reflect.TypeFor[float64](), false}}}, {reflect.TypeFor[FaNycS0m](), 48, []field{{"IdNatD8MGczW", 0, "json:\"code\"", reflect.TypeFor[int](), false}, {"UzQ0WhgKfU", 8, "json:\"message\"", reflect.TypeFor[string](), false}, {"GPETCf2", 24, "json:\"data\"", reflect.TypeFor[[]*Pn5vgsrEm5](), false}}}, {reflect.TypeFor[YEaGcD](), 128, []field{{"IdNatD8MGczW", 0, "json:\"code\"", reflect.TypeFor[int](), false}, {"UzQ0WhgKfU", 8, "json:\"message\"", reflect.TypeFor[string](), false}, {"GPETCf2", 24, "json:\"data\"", reflect.TypeFor[GDzADaHU](), false}}}, {reflect.TypeFor[GDzADaHU](), 104, []field{{"WO4euQQs", 0, "json:\"speed_limit\"", reflect.TypeFor[float64](), false}, {"AzPzYuvK9", 8, "json:\"client_limit\"", reflect.TypeFor[int](), false}, {"I6bjCC2H4sJm", 16, "json:\"v2_port\"", reflect.TypeFor[int](), false}, {"QgIbz8hHD0", 24, "json:\"v2_alter_id\"", reflect.TypeFor[int](), false}, {"A05Abab", 32, "json:\"v2_net\"", reflect.TypeFor[string](), false}, {"LVgd5PJKI", 48, "json:\"v2_path\"", reflect.TypeFor[string](), false}, {"HaC47go2d8Dz", 64, "json:\"v2_tls\"", reflect.TypeFor[bool](), false}, {"NHfgG1rTsQ", 72, "json:\"pem\"", reflect.TypeFor[string](), false}, {"SOjYQldEtR5", 88, "json:\"key\"", reflect.TypeFor[string](), false}}},
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
