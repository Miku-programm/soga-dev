package g8mMV5wjV

import (
	"reflect"
	"testing"
)

func TestNativeAdditionalRTTI(t *testing.T) {
	{
		typ := reflect.TypeOf(LVPKRSqpONvB{})
		if typ.Size() != 16 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "lxoG0T" || f.Offset != 0 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "jI0Cpy13WusE" || f.Offset != 8 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(qp7HhDhYVKA{})
		if typ.Size() != 16 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "lxoG0T" || f.Offset != 0 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "jI0Cpy13WusE" || f.Offset != 8 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
}
