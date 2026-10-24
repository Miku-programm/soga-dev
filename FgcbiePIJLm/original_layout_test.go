package FgcbiePIJLm

import (
	"reflect"
	"testing"
)

func TestRecoveredOriginalRTTI(t *testing.T) {
	{
		typ := reflect.TypeOf(Wmyq5K0sQ0w{})
		if typ.Size() != 16 || typ.NumField() != 6 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "HcIhwPLvD" || f.Offset != 0 || f.Type.Size() != 1 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "PwwYpeWQ" || f.Offset != 1 || f.Type.Size() != 1 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(2)
			if f.Name != "IsHe8C" || f.Offset != 2 || f.Type.Size() != 1 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(3)
			if f.Name != "X4IS_ER04TVv" || f.Offset != 3 || f.Type.Size() != 1 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(4)
			if f.Name != "BiCmXtyECQl" || f.Offset != 4 || f.Type.Size() != 4 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(5)
			if f.Name != "L4ixlgT8VpTo" || f.Offset != 8 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(PFIuXmv{})
		if typ.Size() != 40 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "IBIsKFqaR" || f.Offset != 0 || f.Type.Size() != 16 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "CUdi3A1jBq" || f.Offset != 16 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(T9tLfPri0v{})
		if typ.Size() != 4 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "HHa11Na0i" || f.Offset != 0 || f.Type.Size() != 2 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "TCmULKM" || f.Offset != 2 || f.Type.Size() != 2 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
}
