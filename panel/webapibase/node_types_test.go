package webapibase

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNodeRTTILayout(t *testing.T) {
	{
		rt := reflect.TypeOf(DX8SNQb7vp{})
		if rt.Size() != 64 || rt.NumField() != 4 {
			t.Fatalf("DX8SNQb7vp layout: %v/%d", rt, rt.Size())
		}
		if f := rt.Field(0); f.Name != "CEfy7Qj" || f.Offset != 0 || f.Type.Size() != 8 || string(f.Tag) != "json:\"id\"" || f.Anonymous || f.Type.String() != "int" {
			t.Fatalf("DX8SNQb7vp field 0: %#v", f)
		}
		if f := rt.Field(1); f.Name != "Fu2GAU2hpa1b" || f.Offset != 8 || f.Type.Size() != 24 || string(f.Tag) != "json:\"match\"" || f.Anonymous || f.Type.String() != "jsontext.Value" {
			t.Fatalf("DX8SNQb7vp field 1: %#v", f)
		}
		if f := rt.Field(2); f.Name != "ReTngXeu" || f.Offset != 32 || f.Type.Size() != 16 || string(f.Tag) != "json:\"action\"" || f.Anonymous || f.Type.String() != "string" {
			t.Fatalf("DX8SNQb7vp field 2: %#v", f)
		}
		if f := rt.Field(3); f.Name != "CoShU6j" || f.Offset != 48 || f.Type.Size() != 16 || string(f.Tag) != "json:\"action_value,omitempty\"" || f.Anonymous || f.Type.String() != "string" {
			t.Fatalf("DX8SNQb7vp field 3: %#v", f)
		}
	}
	{
		rt := reflect.TypeOf(U3nQ9iaQj{})
		if rt.Size() != 192 || rt.NumField() != 12 {
			t.Fatalf("U3nQ9iaQj layout: %v/%d", rt, rt.Size())
		}
		if f := rt.Field(0); f.Name != "VKhFmhBGQ" || f.Offset != 0 || f.Type.Size() != 16 || string(f.Tag) != "json:\"message\"" || f.Anonymous || f.Type.String() != "string" {
			t.Fatalf("U3nQ9iaQj field 0: %#v", f)
		}
		if f := rt.Field(1); f.Name != "BDLnnhaR5" || f.Offset != 16 || f.Type.Size() != 8 || string(f.Tag) != "json:\"server_port\"" || f.Anonymous || f.Type.String() != "int" {
			t.Fatalf("U3nQ9iaQj field 1: %#v", f)
		}
		if f := rt.Field(2); f.Name != "RqE21w" || f.Offset != 24 || f.Type.Size() != 24 || string(f.Tag) != "json:\"routes\"" || f.Anonymous || f.Type.String() != "webapibase.IUXU5ZXqeiRC" {
			t.Fatalf("U3nQ9iaQj field 2: %#v", f)
		}
		if f := rt.Field(3); f.Name != "WNvzoiMPp" || f.Offset != 48 || f.Type.Size() != 16 || string(f.Tag) != "json:\"base_config\"" || f.Anonymous {
			t.Fatalf("U3nQ9iaQj field 3: %#v", f)
		}
		if f := rt.Field(4); f.Name != "X6Jwaaf492C" || f.Offset != 64 || f.Type.Size() != 16 || string(f.Tag) != "json:\"network\"" || f.Anonymous || f.Type.String() != "types.BoulTn31" {
			t.Fatalf("U3nQ9iaQj field 4: %#v", f)
		}
		if f := rt.Field(5); f.Name != "MDoXQFch8aWk" || f.Offset != 80 || f.Type.Size() != 24 || string(f.Tag) != "json:\"networkSettings\"" || f.Anonymous || f.Type.String() != "jsontext.Value" {
			t.Fatalf("U3nQ9iaQj field 5: %#v", f)
		}
		if f := rt.Field(6); f.Name != "PqqmTY8L" || f.Offset != 104 || f.Type.Size() != 8 || string(f.Tag) != "json:\"tls\"" || f.Anonymous || f.Type.String() != "int" {
			t.Fatalf("U3nQ9iaQj field 6: %#v", f)
		}
		if f := rt.Field(7); f.Name != "ELUl7_FnS" || f.Offset != 112 || f.Type.Size() != 16 || string(f.Tag) != "json:\"host\"" || f.Anonymous || f.Type.String() != "string" {
			t.Fatalf("U3nQ9iaQj field 7: %#v", f)
		}
		if f := rt.Field(8); f.Name != "BNKMfAf5" || f.Offset != 128 || f.Type.Size() != 16 || string(f.Tag) != "json:\"server_name\"" || f.Anonymous || f.Type.String() != "string" {
			t.Fatalf("U3nQ9iaQj field 8: %#v", f)
		}
		if f := rt.Field(9); f.Name != "FYeHum" || f.Offset != 144 || f.Type.Size() != 16 || string(f.Tag) != "json:\"cipher\"" || f.Anonymous || f.Type.String() != "ssmethod.C6q3Upu" {
			t.Fatalf("U3nQ9iaQj field 9: %#v", f)
		}
		if f := rt.Field(10); f.Name != "D07VWhUsih" || f.Offset != 160 || f.Type.Size() != 16 || string(f.Tag) != "json:\"server_key\"" || f.Anonymous || f.Type.String() != "string" {
			t.Fatalf("U3nQ9iaQj field 10: %#v", f)
		}
		if f := rt.Field(11); f.Name != "L6f4wB51G0cC" || f.Offset != 176 || f.Type.Size() != 16 || string(f.Tag) != "json:\"obfs\"" || f.Anonymous || f.Type.String() != "ssrnames.PMorzD7u" {
			t.Fatalf("U3nQ9iaQj field 11: %#v", f)
		}
	}
}
func TestBaseConfigRTTI(t *testing.T) {
	rt := reflect.TypeOf(U3nQ9iaQj{}).Field(3).Type
	if rt.Size() != 16 || rt.NumField() != 2 {
		t.Fatalf("base layout: %v", rt)
	}
	for i, want := range []struct{ name, tag string }{{"F2WCOEc5Eu6", `json:"pull_interval"`}, {"I8kigte", `json:"push_interval"`}} {
		f := rt.Field(i)
		if f.Name != want.name || string(f.Tag) != want.tag || f.Offset != uintptr(i*8) || f.Type.String() != "runtime.JlKvxH[int]" || f.Anonymous {
			t.Fatalf("base field: %#v", f)
		}
	}
}

func TestNodeJSONAndMatchParsing(t *testing.T) {
	var n U3nQ9iaQj
	input := []byte(`{"server_port":443,"routes":[{"id":3,"match":[" a ","",null],"action":"block"}],"base_config":{"pull_interval":"60.9","push_interval":30},"network":"ws","networkSettings":{"path":"/ws"},"tls":1,"host":"node.test","server_name":"tls.test","cipher":"aes-128-gcm","server_key":"key","obfs":"http"}`)
	if err := json.Unmarshal(input, &n); err != nil {
		t.Fatal(err)
	}
	if n.BDLnnhaR5 != 443 || n.WNvzoiMPp.F2WCOEc5Eu6.Value() != 60 || n.WNvzoiMPp.I8kigte.Value() != 30 || n.X6Jwaaf492C != "ws" || string(n.MDoXQFch8aWk) != `{"path":"/ws"}` || n.FYeHum != "aes-128-gcm" || n.L6f4wB51G0cC != "http" {
		t.Fatalf("node: %#v", n)
	}
	for _, tc := range []struct {
		raw  string
		want []string
	}{
		{`[" a ","",null]`, []string{" a ", "", ""}},
		{`null`, nil}, {`[]`, []string{}},
		{` alpha , , beta `, []string{"alpha", "beta"}},
		{`"alpha,beta"`, []string{`"alpha`, `beta"`}},
		{`["a",42]`, []string{`["a"`, `42]`}},
	} {
		r := DX8SNQb7vp{Fu2GAU2hpa1b: []byte(tc.raw)}
		if got := r.d09KQUdaP05h(); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("match %s: %#v", tc.raw, got)
		}
	}
	b, err := json.Marshal(&n)
	if err != nil {
		t.Fatal(err)
	}
	var wire map[string]any
	if err = json.Unmarshal(b, &wire); err != nil {
		t.Fatal(err)
	}
	if _, ok := wire["networkSettings"].(map[string]any); !ok {
		t.Fatalf("raw JSON was not preserved: %s", b)
	}
}
