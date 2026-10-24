package genmap

import (
	"reflect"
	"testing"
)

func TestConcurrentSetOriginalMethods(t *testing.T) {
	var s Fxs7SY[int64]
	if reflect.TypeOf(s).Size() != 56 || s.ToList() == nil {
		t.Fatal("layout/empty slice")
	}
	s.Add(4)
	s.Add(4)
	s.Add(9)
	if s.Len() != 2 || !s.Has(4) {
		t.Fatal(s.ToList())
	}
	n := 0
	s.Range(func(int64) bool { n++; return false })
	if n != 1 {
		t.Fatal(n)
	}
	list := s.ToList()
	if len(list) != 2 || cap(list) != 2 {
		t.Fatal(list)
	}
	s.Delete(4)
	s.Delete(4)
	if s.Len() != 1 || s.Has(4) {
		t.Fatal(s.ToList())
	}
	s.Clear()
	if s.Len() != 0 || s.Has(9) {
		t.Fatal(s.ToList())
	}
	s.Add(1)
	if !s.Has(1) {
		t.Fatal("reuse")
	}
}
