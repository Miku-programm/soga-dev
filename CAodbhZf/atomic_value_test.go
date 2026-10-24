package CAodbhZf

import "testing"

func TestNativeAtomicNilWrapper(t *testing.T) {
	var v EahpvU_iseQ[*int]
	if v.Load() != nil {
		t.Fatal("zero")
	}
	x := 7
	if v.CompareAndSwap(nil, &x) {
		t.Fatal("uninitialized wrapper CAS should fail")
	}
	v.Store(nil)
	if !v.CompareAndSwap(nil, &x) || v.Load() != &x {
		t.Fatal("typed nil CAS")
	}
	if v.Swap(nil) != &x || v.Load() != nil {
		t.Fatal("swap")
	}
}
