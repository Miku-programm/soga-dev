package httphead

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNativeTokensAndGrammar(t *testing.T) {
	for _, tc := range []struct {
		in   string
		ok   bool
		want []string
	}{{"one,two", true, []string{"one", "two"}}, {"one, two", true, []string{"one", "two"}}, {",", false, nil}, {"one,(comment)", false, []string{"one"}}, {"one\t,two", false, []string{"one"}}} {
		var got []string
		ok := PgwXhS7T([]byte(tc.in), func(b []byte) bool { got = append(got, string(b)); return true })
		if ok != tc.ok || !reflect.DeepEqual(got, tc.want) {
			t.Fatal(tc, ok, got)
		}
	}
}
func TestNativeOptionsInlineDynamicCopyAndEncoding(t *testing.T) {
	input := []byte(`sample;a=one;b="two words";c=three;d=four;e=five;f=six;g=seven;h=eight;i=nine;flag,other`)
	options, ok := (ExK8bKI2Xox{Flags: SelectCopy}).Select(input, nil)
	if !ok || len(options) != 2 || string(options[0].KXa0Kt2I) != "sample" || len(options[0].Jd2j6aD7.data()) != 10 {
		t.Fatal(ok, len(options))
	}
	for i := range input {
		input[i] = 'x'
	}
	var wire bytes.Buffer
	n, e := KBgYhFImpuOF(&wire, options)
	want := `sample;a=one;b="two words";c=three;d=four;e=five;f=six;g=seven;h=eight;i=nine;flag,other`
	if e != nil || n != len(want) || wire.String() != want {
		t.Fatal(n, e, wire.String())
	}
}
func TestNativeScannerControls(t *testing.T) {
	var got []string
	ok := LU7plsn8k([]byte("skip;a=1;b=2,keep;c=3"), func(_ int, name, attr, val []byte) Control {
		got = append(got, string(name))
		if string(name) == "skip" {
			return ControlSkip
		}
		return ControlBreak
	})
	if !ok || !reflect.DeepEqual(got, []string{"skip", "keep"}) {
		t.Fatal(ok, got)
	}
	func() {
		defer func() {
			if e := recover(); e != "unexpected control value" {
				t.Fatal(e)
			}
		}()
		LU7plsn8k([]byte("a;b=c"), func(int, []byte, []byte, []byte) Control { return 9 })
	}()
	if EQ1c264([]byte(`a\",b"`), '"') != 5 {
		t.Fatal("escaped delimiter")
	}
}

type failingHeadWriter struct{ n int }

func (w *failingHeadWriter) Write(p []byte) (int, error) { w.n++; return 0, errors.New("failure") }
func TestNativeEncodingStopsAfterError(t *testing.T) {
	w := new(failingHeadWriter)
	n, e := KBgYhFImpuOF(w, []BIMZTmbYHb{{KXa0Kt2I: []byte("a")}, {KXa0Kt2I: []byte("b")}})
	if n != 0 || e == nil || w.n != 1 {
		t.Fatal(fmt.Sprint(n, e, w.n))
	}
}
