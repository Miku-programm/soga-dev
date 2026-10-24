package QHUZzZ2d

import (
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestDecoderRTTIAndActualPluginShape(t *testing.T) {
	if unsafe.Sizeof(MBi603C{}) != 24 || unsafe.Offsetof(MBi603C{}.RlQxpRzya) != 16 || unsafe.Sizeof(LsMT6BB{}) != 8 {
		t.Fatal("RTTI mismatch")
	}
	d := K6YuYvsMD5Y(MBi603C{RlQxpRzya: true})
	if d.d01d_V594LQ.YONNn23 != "structure" {
		t.Fatal(d)
	}
	var output struct {
		Mode string `structure:"mode"`
		Host string `structure:"host,omitempty"`
		Port int    `structure:"port"`
		TLS  bool   `structure:"tls"`
	}
	if err := d.Decode(map[string]any{"mode": "websocket", "port": "0x1bb", "tls": 1}, &output); err != nil {
		t.Fatal(err)
	}
	if output.Mode != "websocket" || output.Port != 443 || !output.TLS || output.Host != "" {
		t.Fatal(output)
	}
	if err := d.Decode(map[string]any{"mode": "websocket", "port": 443, "tls": "true"}, &output); err == nil {
		t.Fatal("sample does not parse bool strings")
	}
}

func TestOriginalDecoderTypeBoundaries(t *testing.T) {
	weak := K6YuYvsMD5Y(MBi603C{RlQxpRzya: true})
	strict := K6YuYvsMD5Y(MBi603C{})
	for _, tc := range []struct {
		data     any
		expected int
		valid    bool
	}{{int(12), 12, true}, {float64(12.9), 12, true}, {"077", 63, true}, {int64(12), 0, false}, {uint(12), 0, false}, {true, 0, false}, {"08", 0, false}} {
		value := 0
		err := weak.kFuaz5EqcMhd("port", tc.data, reflect.ValueOf(&value).Elem())
		if (err == nil) != tc.valid || tc.valid && value != tc.expected {
			t.Fatalf("%T(%v) => %d,%v", tc.data, tc.data, value, err)
		}
	}
	var i int
	if err := strict.kFuaz5EqcMhd("n", float64(1), reflect.ValueOf(&i).Elem()); err == nil {
		t.Fatal("strict float accepted")
	}
	var unsupported int64
	if err := weak.kFuaz5EqcMhd("n", int64(1), reflect.ValueOf(&unsupported).Elem()); err == nil || err.Error() != "type int64 not support" {
		t.Fatal(err)
	}
	var out struct{ X string }
	if err := weak.Decode(map[string]any{"X": "value"}, &out); err == nil || err.Error() != "key '' missing" {
		t.Fatal(err)
	}
	if err := weak.Decode(map[string]any{"": "value"}, &out); err != nil || out.X != "value" {
		t.Fatal(out, err)
	}
}

func TestDecoderSlicesRetainTailAndPartialWrites(t *testing.T) {
	d := K6YuYvsMD5Y(MBi603C{RlQxpRzya: true})
	xs := []int{7, 8, 9}
	if err := d.kFuaz5EqcMhd("xs", []any{1, nil}, reflect.ValueOf(&xs).Elem()); err != nil || !reflect.DeepEqual(xs, []int{1, 8, 9}) {
		t.Fatal(xs, err)
	}
	strict := K6YuYvsMD5Y(MBi603C{})
	if err := strict.kFuaz5EqcMhd("xs", []any{2, nil}, reflect.ValueOf(&xs).Elem()); err == nil || err.Error() != "'xs[1]' can not be null" || xs[0] != 2 {
		t.Fatal(xs, err)
	}
	var maps []map[string]int
	if err := strict.kFuaz5EqcMhd("maps", []any{nil}, reflect.ValueOf(&maps).Elem()); err != nil || len(maps) != 1 || maps[0] != nil {
		t.Fatal(maps, err)
	}
}

func TestDecoderMapPartialSuccessAndOriginalErrorFormatting(t *testing.T) {
	d := K6YuYvsMD5Y(MBi603C{})
	output := map[string]int{"old": 3}
	err := d.kFuaz5EqcMhd("entry", map[string]any{"good": 7, "bad": true}, reflect.ValueOf(&output).Elem())
	if err == nil || output["good"] != 7 || output["old"] != 3 {
		t.Fatal(output, err)
	}
	err = d.kFuaz5EqcMhd("%v", map[string]any{"nil": nil}, reflect.ValueOf(&output).Elem())
	if err == nil || !strings.Contains(err.Error(), "%!v(MISSING)") {
		t.Fatal("original aggregate-error format processing changed", err)
	}
	var nilMap map[string]int
	if err := d.kFuaz5EqcMhd("entry", nilMap, reflect.ValueOf(&output).Elem()); err != nil || output != nil {
		t.Fatal(output, err)
	}
}

type embeddedFixture struct {
	Value string `structure:"value"`
}

func TestNestedStructSquashAndCaseFold(t *testing.T) {
	d := K6YuYvsMD5Y(MBi603C{RlQxpRzya: true})
	var out struct {
		Nested struct {
			Embed embeddedFixture `structure:",squash"`
			Count int             `structure:"count"`
		} `structure:"nested"`
	}
	err := d.Decode(map[string]any{"nested": map[string]any{"VALUE": "ok", "COUNT": "12", "unknown": 3}}, &out)
	if err != nil || out.Nested.Embed.Value != "ok" || out.Nested.Count != 12 {
		t.Fatal(out, err)
	}
	var x struct {
		Nested struct {
			Number int `structure:",squash"`
		} `structure:"nested"`
	}
	if err := d.Decode(map[string]any{"nested": map[string]any{}}, &x); err == nil || !strings.Contains(err.Error(), "unsupported type for squash: int") {
		t.Fatal(err)
	}
}

func TestDecoderNilDestinationRetainsOriginalPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("original Decode(nil) panics")
		}
	}()
	K6YuYvsMD5Y(MBi603C{}).Decode(nil, nil)
}
