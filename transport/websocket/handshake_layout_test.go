package websocket

import (
	"reflect"
	"testing"
)

func TestNativeHandshakeRTTI(t *testing.T) {
	{
		typ := reflect.TypeOf(jC7sxat{})
		if typ.Size() != 24 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "zqxPTJEN" || f.Offset != 0 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "h7OgHq0GODs" || f.Offset != 8 || f.Type.Size() != 16 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(Zbu9jHc{})
		if typ.Size() != 40 || typ.NumField() != 3 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "lPbsfO" || f.Offset != 0 || f.Type.Size() != 16 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "uIlV5a" || f.Offset != 16 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(2)
			if f.Name != "oCKugs" || f.Offset != 24 || f.Type.Size() != 16 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
}
