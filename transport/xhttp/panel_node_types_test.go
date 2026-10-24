package xhttp

import (
	"reflect"
	"testing"
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
		{reflect.TypeOf(R0J19auWof{}), 296, []field{
			{"TzkkcVTusJD", 0, reflect.TypeFor[string](), "json:\"host,omitempty\""},
			{"HoBz4V", 16, reflect.TypeFor[string](), "json:\"path,omitempty\""},
			{"MfWFUwd1", 32, reflect.TypeFor[string](), "json:\"mode,omitempty\""},
			{"F0XSwqNYoaP", 48, reflect.TypeFor[map[string]string](), "json:\"headers,omitempty\""},
			{"N2C4XMihD", 56, reflect.TypeFor[*MUVoWl0](), "json:\"xPaddingBytes,omitempty\""},
			{"A6NxppvA", 64, reflect.TypeFor[bool](), "json:\"noSSEHeader,omitempty\""},
			{"ULnN_iDWLIXN", 72, reflect.TypeFor[*MUVoWl0](), "json:\"scMaxEachPostBytes,omitempty\""},
			{"LEaRo0", 80, reflect.TypeFor[int64](), "json:\"scMaxBufferedPosts,omitempty\""},
			{"Iv950ouQ5AW", 88, reflect.TypeFor[*MUVoWl0](), "json:\"scStreamUpServerSecs,omitempty\""},
			{"PnRF9ifkXKD", 96, reflect.TypeFor[bool](), "json:\"xPaddingObfsMode,omitempty\""},
			{"H6PfvFysST", 104, reflect.TypeFor[string](), "json:\"xPaddingKey,omitempty\""},
			{"LjJbxC6vDL", 120, reflect.TypeFor[string](), "json:\"xPaddingHeader,omitempty\""},
			{"V0qykwP", 136, reflect.TypeFor[string](), "json:\"xPaddingPlacement,omitempty\""},
			{"MLlGWkds", 152, reflect.TypeFor[string](), "json:\"xPaddingMethod,omitempty\""},
			{"KJMkBJ", 168, reflect.TypeFor[string](), "json:\"uplinkHTTPMethod,omitempty\""},
			{"XEKp9SbqKQ", 184, reflect.TypeFor[string](), "json:\"sessionPlacement,omitempty\""},
			{"G9oU9diBqr", 200, reflect.TypeFor[string](), "json:\"sessionKey,omitempty\""},
			{"HqtZJaRZNw", 216, reflect.TypeFor[string](), "json:\"seqPlacement,omitempty\""},
			{"NPLvCM", 232, reflect.TypeFor[string](), "json:\"seqKey,omitempty\""},
			{"Aonc19Hp33E", 248, reflect.TypeFor[string](), "json:\"uplinkDataPlacement,omitempty\""},
			{"HFEkUe4a9Sas", 264, reflect.TypeFor[string](), "json:\"uplinkDataKey,omitempty\""},
			{"JvaHmm", 280, reflect.TypeFor[*MUVoWl0](), "json:\"uplinkChunkSize,omitempty\""},
			{"Diwr8ey", 288, reflect.TypeFor[int32](), "json:\"serverMaxHeaderBytes,omitempty\""},
		}},
		{reflect.TypeOf(MUVoWl0{}), 8, []field{
			{"AS9zG362", 0, reflect.TypeFor[int32](), "json:\"from,omitempty\""},
			{"HHTPwX0N", 4, reflect.TypeFor[int32](), "json:\"to,omitempty\""},
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
