package hysteriaauth

import (
	"bytes"
	"recovered.local/soga/crypto/misc"
	"reflect"
	"testing"
)

func TestRTTILayout(t *testing.T) {
	ty := reflect.TypeFor[YWXI4jW3]()
	if ty.Size() != 48 {
		t.Fatal(ty.Size())
	}
	names := []string{"sQXA1D9aRa", "mk2YbtR", "uGzaQtaWh", "zLEI1e06OnKH"}
	off := []uintptr{0, 8, 32, 40}
	types := []reflect.Type{reflect.TypeFor[uint16](), reflect.TypeFor[[]*misc.ItLMumzBf](), reflect.TypeFor[uint8](), reflect.TypeFor[int]()}
	for i, n := range names {
		f := ty.Field(i)
		if f.Name != n || f.Offset != off[i] || f.Type != types[i] || f.Tag != "" {
			t.Fatal(f)
		}
	}
}
func TestFragmentReorderAndOriginalMutation(t *testing.T) {
	m := &misc.ItLMumzBf{D35ycLc5aB6r: 42, CBHlJenOD: 7, XrDcBXcD: "x", HmoFPk: []byte("abcdefgh")}
	f := TOF1KM6hi(m, m.HeaderSize()+3)
	if len(f) != 3 || string(f[2].HmoFPk) != "gh" {
		t.Fatal(f)
	}
	f[0].HmoFPk[0] = 'A'
	if m.HmoFPk[0] != 'A' {
		t.Fatal("fragment copy")
	}
	var d YWXI4jW3
	if d.Feed(&f[2]) != nil || d.Feed(&f[0]) != nil || d.Feed(&f[0]) != nil {
		t.Fatal("premature/duplicate")
	}
	f[1].XrDcBXcD = "last"
	f[1].D35ycLc5aB6r = 99
	got := d.Feed(&f[1])
	if got != &f[1] || got.QmDvxK_8Q != 0 || got.Kb9FLg7wuNE != 1 || got.XrDcBXcD != "last" || got.D35ycLc5aB6r != 99 || !bytes.Equal(got.HmoFPk, []byte("Abcdefgh")) {
		t.Fatal(got)
	}
	if d.Feed(&f[0]) != nil {
		t.Fatal("state unexpectedly reset")
	}
}
func TestOriginalInvalidAndResetPaths(t *testing.T) {
	var d YWXI4jW3
	m := &misc.ItLMumzBf{Kb9FLg7wuNE: 1}
	if d.Feed(m) != m {
		t.Fatal("unfragmented")
	}
	m.Kb9FLg7wuNE = 2
	m.QmDvxK_8Q = 2
	if d.Feed(m) != nil {
		t.Fatal("invalid id")
	}
	m.QmDvxK_8Q = 0
	if d.Feed(m) != nil {
		t.Fatal("first")
	}
	n := *m
	n.CBHlJenOD++
	n.QmDvxK_8Q = 1
	d.Feed(&n)
	if d.uGzaQtaWh != 1 || d.sQXA1D9aRa != n.CBHlJenOD {
		t.Fatal(d)
	}
}
func TestOriginalMissingFragmentLimitGuard(t *testing.T) {
	for _, extra := range []int{0, -1} {
		func() {
			defer func() {
				if recover() == nil {
					t.Fatal("original invalid maxSize must panic", extra)
				}
			}()
			m := &misc.ItLMumzBf{XrDcBXcD: "x", HmoFPk: []byte("abc")}
			TOF1KM6hi(m, m.HeaderSize()+extra)
		}()
	}
	m := &misc.ItLMumzBf{XrDcBXcD: "x", HmoFPk: []byte("abc")}
	f := TOF1KM6hi(m, 100)
	if len(f) != 1 || &f[0] == m || &f[0].HmoFPk[0] != &m.HmoFPk[0] {
		t.Fatal("single shallow value copy")
	}
}
