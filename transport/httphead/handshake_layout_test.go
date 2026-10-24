package httphead

import (
	"reflect"
	"testing"
)

func TestNativeHandshakeRTTI(t *testing.T) {
	{
		typ := reflect.TypeOf(BIMZTmbYHb{})
		if typ.Size() != 448 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "KXa0Kt2I" || f.Offset != 0 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "Jd2j6aD7" || f.Offset != 24 || f.Type.Size() != 424 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(MVgyIAGYmsKQ{})
		if typ.Size() != 424 || typ.NumField() != 4 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "clXYVDW3J" || f.Offset != 0 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "l6UF09" || f.Offset != 8 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(2)
			if f.Name != "rraG37Ve" || f.Offset != 16 || f.Type.Size() != 384 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(3)
			if f.Name != "hMg6gO_J" || f.Offset != 400 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(blm5Wb{})
		if typ.Size() != 48 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "z9Lbkony" || f.Offset != 0 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "yZSzEF" || f.Offset != 24 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
}
