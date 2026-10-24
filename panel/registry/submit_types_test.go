package registry

import (
	"recovered.local/soga/panel/types"
	"reflect"
	"testing"
)

func TestSubmitRTTI(t *testing.T) {
	typ := reflect.TypeOf(FGmNba4MhioY{})
	names := []string{"HTgPyE3VytZd", "JNcHHwTZ7", "Y0ys4ma4Tno4"}
	types := []reflect.Type{reflect.TypeFor[types.CWfcUv](), reflect.TypeFor[types.CQhZrJ4dK](), reflect.TypeFor[types.QmiwDVjV00lZ]()}
	if typ.Size() != 72 || typ.NumField() != 3 {
		t.Fatal("RTTI结构大小不符")
	}
	for i, name := range names {
		f := typ.Field(i)
		if f.Name != name || f.Type != types[i] || f.Offset != uintptr(24*i) || f.Tag != "" {
			t.Fatalf("RTTI字段不符 %+v", f)
		}
	}
}
