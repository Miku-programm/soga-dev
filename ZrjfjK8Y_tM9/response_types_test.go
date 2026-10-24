package ZrjfjK8Y_tM9

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestResponseRTTILayout(t *testing.T) {
	{
		rt := reflect.TypeOf(GIAIBDeJa25{})
		if rt.Size() != 16 || rt.NumField() != 1 {
			t.Fatalf("GIAIBDeJa25 布局: %v/%d", rt, rt.Size())
		}
		if f := rt.Field(0); f.Name != "RVKq_aLvfWyF" || f.Offset != 0 || f.Type.String() != "string" || string(f.Tag) != "json:\"serviceName\"" || f.Anonymous {
			t.Fatalf("GIAIBDeJa25 字段 0: %#v", f)
		}
	}
	{
		rt := reflect.TypeOf(Re7aaf{})
		if rt.Size() != 16 || rt.NumField() != 1 {
			t.Fatalf("Re7aaf 布局: %v/%d", rt, rt.Size())
		}
		if f := rt.Field(0); f.Name != "JCRhkrluQQY" || f.Offset != 0 || f.Type.String() != "string" || string(f.Tag) != "json:\"path\"" || f.Anonymous {
			t.Fatalf("Re7aaf 字段 0: %#v", f)
		}
	}
	{
		rt := reflect.TypeOf(D9MPLRTuA{})
		if rt.Size() != 40 || rt.NumField() != 2 {
			t.Fatalf("D9MPLRTuA 布局: %v/%d", rt, rt.Size())
		}
		if f := rt.Field(0); f.Name != "EWFMfZX" || f.Offset != 0 || f.Type.String() != "string" || string(f.Tag) != "json:\"message\"" || f.Anonymous {
			t.Fatalf("D9MPLRTuA 字段 0: %#v", f)
		}
		if f := rt.Field(1); f.Name != "Ta593Rq2lSa" || f.Offset != 16 || f.Type.String() != "ZrjfjK8Y_tM9.Uhu9JY7Vj" || string(f.Tag) != "json:\"users\"" || f.Anonymous {
			t.Fatalf("D9MPLRTuA 字段 1: %#v", f)
		}
	}
	{
		rt := reflect.TypeOf(WVfKv8zPvt{})
		if rt.Size() != 40 || rt.NumField() != 4 {
			t.Fatalf("WVfKv8zPvt 布局: %v/%d", rt, rt.Size())
		}
		if f := rt.Field(0); f.Name != "IDLJXX6F" || f.Offset != 0 || f.Type.String() != "int64" || string(f.Tag) != "json:\"id\"" || f.Anonymous {
			t.Fatalf("WVfKv8zPvt 字段 0: %#v", f)
		}
		if f := rt.Field(1); f.Name != "UH0HXa5O" || f.Offset != 8 || f.Type.String() != "string" || string(f.Tag) != "json:\"uuid\"" || f.Anonymous {
			t.Fatalf("WVfKv8zPvt 字段 1: %#v", f)
		}
		if f := rt.Field(2); f.Name != "DwdlKBQpcgY" || f.Offset != 24 || f.Type.String() != "float64" || string(f.Tag) != "json:\"speed_limit\"" || f.Anonymous {
			t.Fatalf("WVfKv8zPvt 字段 2: %#v", f)
		}
		if f := rt.Field(3); f.Name != "QeuUtZ8f" || f.Offset != 32 || f.Type.String() != "int" || string(f.Tag) != "json:\"device_limit\"" || f.Anonymous {
			t.Fatalf("WVfKv8zPvt 字段 3: %#v", f)
		}
	}
	if rt := reflect.TypeOf(Uhu9JY7Vj{}); rt.Kind() != reflect.Slice || rt.Elem() != reflect.TypeOf((*WVfKv8zPvt)(nil)) {
		t.Fatalf("用户列表元素: %v", rt)
	}
}
func TestResponseJSONFields(t *testing.T) {
	var got D9MPLRTuA
	input := []byte(`{"message":"","users":[{"id":9223372036854775807,"uuid":"id-value","speed_limit":12.5,"device_limit":3},null]}`)
	if err := json.Unmarshal(input, &got); err != nil {
		t.Fatal(err)
	}
	if got.EWFMfZX != "" || len(got.Ta593Rq2lSa) != 2 || got.Ta593Rq2lSa[1] != nil {
		t.Fatalf("response: %#v", got)
	}
	u := got.Ta593Rq2lSa[0]
	if u.IDLJXX6F != 9223372036854775807 || u.UH0HXa5O != "id-value" || u.DwdlKBQpcgY != 12.5 || u.QeuUtZ8f != 3 {
		t.Fatalf("user: %#v", u)
	}
	encoded, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != string(input) {
		t.Fatalf("JSON: %s", encoded)
	}
	for _, v := range []struct {
		value any
		want  string
	}{
		{GIAIBDeJa25{RVKq_aLvfWyF: "svc"}, `{"serviceName":"svc"}`},
		{Re7aaf{JCRhkrluQQY: "/ws"}, `{"path":"/ws"}`},
	} {
		b, err := json.Marshal(v.value)
		if err != nil || string(b) != v.want {
			t.Fatalf("JSON: %s %v", b, err)
		}
	}
}
