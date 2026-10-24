package CAodbhZf

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNumberRTTIAndConversion(t *testing.T) {
	rt := reflect.TypeOf(JlKvxH[int]{})
	if rt.Size() != 8 || rt.NumField() != 1 || rt.Field(0).Name != "_2L2dhl" || rt.Field(0).Offset != 0 || rt.Field(0).Type.Kind() != reflect.Int || rt.Field(0).IsExported() {
		t.Fatalf("RTTI: %v", rt)
	}
	for _, tc := range []struct {
		text string
		want int
		bad  bool
	}{
		{"60", 60, false}, {"12.9", 12, false}, {"-12.9", -12, false}, {`"61.5"`, 61, false}, {"6e1", 60, false},
		{"null", 7, true}, {`"invalid"`, 7, true}, {"true", 7, true}, {"{}", 7, true}, {"1e400", 7, true},
	} {
		n := JlKvxH[int]{_2L2dhl: 7}
		err := json.Unmarshal([]byte(tc.text), &n)
		if (err != nil) != tc.bad || n.Value() != tc.want {
			t.Fatalf("%s: %d %v", tc.text, n.Value(), err)
		}
		b, err := json.Marshal(&n)
		if err != nil {
			t.Fatal(err)
		}
		var raw int
		if err = json.Unmarshal(b, &raw); err != nil || raw != tc.want {
			t.Fatalf("MarshalJSON: %s %v", b, err)
		}
	}
}
