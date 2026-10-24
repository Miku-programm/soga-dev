package genmap

import "testing"

func TestSetOperationsAndRangeStop(t *testing.T) {
	s := PQQKui[int64]()
	s.Add(7)
	s.Add(7)
	s.Add(-1)
	if s.Len() != 2 || !s.Has(7) || s.Has(0) {
		t.Fatal("set membership differs")
	}
	n := 0
	s.Range(func(value int64) bool { n++; return false })
	if n != 1 {
		t.Fatal("Range must stop on false")
	}
	s.Del(7)
	s.Del(99)
	if s.Len() != 1 || s.Has(7) {
		t.Fatal("delete differs")
	}
	var zero Fw_lgKgam_jg[int64]
	if zero.Len() != 0 || zero.Has(1) {
		t.Fatal("nil map reads differ")
	}
	zero.Del(1)
	zero.Range(nil)
	defer func() {
		if recover() == nil {
			t.Fatal("Add to a nil map must panic")
		}
	}()
	zero.Add(1)
}
