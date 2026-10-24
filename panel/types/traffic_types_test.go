package types

import (
	"reflect"
	"testing"
)

func TestTrafficRTTIStructures(t *testing.T) {
	for _, tc := range []struct {
		typ         reflect.Type
		size        uintptr
		names, tags []string
		offsets     []uintptr
		types       []reflect.Type
	}{
		{reflect.TypeOf(DLxslMwYJXQw{}), 24, []string{"Hy7Q97RROK", "XJkGTQf_qkAt", "I6LXObODav"}, []string{`json:"user_id"`, `json:"u"`, `json:"d"`}, []uintptr{0, 8, 16}, []reflect.Type{reflect.TypeFor[int64](), reflect.TypeFor[int64](), reflect.TypeFor[int64]()}},
		{reflect.TypeOf(LmNKoY94{}), 32, []string{"Vy9ybu", "PNtn9KAV"}, []string{`json:"user_id"`, `json:"ips"`}, []uintptr{0, 8}, []reflect.Type{reflect.TypeFor[int64](), reflect.TypeFor[[]string]()}},
		{reflect.TypeOf(B9FS114R4F9h{}), 16, []string{"GbzQUT_4f", "CAgUUdrv9"}, []string{`json:"user_id"`, `json:"detect_id"`}, []uintptr{0, 8}, []reflect.Type{reflect.TypeFor[int64](), reflect.TypeFor[int]()}},
	} {
		if tc.typ.Size() != tc.size || tc.typ.NumField() != len(tc.names) {
			t.Fatalf("RTTI大小不符 %v", tc.typ)
		}
		for i, name := range tc.names {
			f := tc.typ.Field(i)
			if f.Name != name || f.Offset != tc.offsets[i] || string(f.Tag) != tc.tags[i] || f.Type != tc.types[i] {
				t.Fatalf("RTTI字段不符 %v %v", tc.typ, f)
			}
		}
	}
	for _, tc := range []struct{ slice, elem reflect.Type }{{reflect.TypeFor[CWfcUv](), reflect.TypeFor[*DLxslMwYJXQw]()}, {reflect.TypeFor[CQhZrJ4dK](), reflect.TypeFor[*LmNKoY94]()}, {reflect.TypeFor[QmiwDVjV00lZ](), reflect.TypeFor[*B9FS114R4F9h]()}} {
		if tc.slice.Kind() != reflect.Slice || tc.slice.Elem() != tc.elem {
			t.Fatalf("RTTI切片元素不符 %v", tc.slice)
		}
	}
}
