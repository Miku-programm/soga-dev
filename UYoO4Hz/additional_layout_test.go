package UYoO4Hz

import (
	"reflect"
	"testing"
)

func TestNativeAdditionalRTTI(t *testing.T) {
	{
		typ := reflect.TypeOf(M6JkkEsBtZH{})
		if typ.Size() != 144 || typ.NumField() != 11 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "bEEZIz5" || f.Offset != 0 || f.Type.Size() != 16 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "pIitEG" || f.Offset != 16 || f.Type.Size() != 1 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(2)
			if f.Name != "sSjoMgATz0p" || f.Offset != 17 || f.Type.Size() != 1 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(3)
			if f.Name != "v78A5VtJX" || f.Offset != 24 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(4)
			if f.Name != "bKFgx3c" || f.Offset != 48 || f.Type.Size() != 1 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(5)
			if f.Name != "gGZQkcSDF1jp" || f.Offset != 56 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(6)
			if f.Name != "ariAbujn0" || f.Offset != 80 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(7)
			if f.Name != "jfnerh_x6" || f.Offset != 104 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(8)
			if f.Name != "e_fTv76YzJG" || f.Offset != 112 || f.Type.Size() != 1 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(9)
			if f.Name != "vB4Yl8PWQC" || f.Offset != 120 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(10)
			if f.Name != "vQT2iWLEILFo" || f.Offset != 128 || f.Type.Size() != 16 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(DH32d4Xf{})
		if typ.Size() != 24 || typ.NumField() != 3 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "mpzgoFUbRt" || f.Offset != 0 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "jJkHg4kv" || f.Offset != 8 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(2)
			if f.Name != "fZvXjIM" || f.Offset != 16 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(zvrEv0{})
		if typ.Size() != 32 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "b5FO9d" || f.Offset != 0 || f.Type.Size() != 24 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "lFd2tGa8_bb" || f.Offset != 24 || f.Type.Size() != 8 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
	{
		typ := reflect.TypeOf(VGrYaNsp{})
		if typ.Size() != 24 || typ.NumField() != 2 {
			t.Fatal(typ, typ.Size(), typ.NumField())
		}
		{
			f := typ.Field(0)
			if f.Name != "EAGe4eFzGle" || f.Offset != 0 || f.Type.Size() != 2 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
		{
			f := typ.Field(1)
			if f.Name != "Nd7HXs1GwPQ" || f.Offset != 8 || f.Type.Size() != 16 || string(f.Tag) != "" {
				t.Fatal(typ, f)
			}
		}
	}
}
