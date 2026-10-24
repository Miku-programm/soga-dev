package limit

import (
	"testing"
	"unsafe"
)

func TestOriginalTrafficWindow(t *testing.T) {
	if unsafe.Sizeof(Mblp8L{}) != 48 {
		t.Fatal("RTTI")
	}
	w := &Mblp8L{d5H9TK: make([]int64, 3)}
	for i, s := range []struct {
		input, sum int64
		index      int
	}{{0, 0, 0}, {3, 3, 1}, {5, 8, 2}, {7, 15, 0}, {0, 12, 1}, {0, 7, 2}, {0, 0, 0}, {0, 0, 0}, {-4, -4, 1}, {2, -2, 2}, {2, 0, 0}, {0, 0, 0}} {
		w.AddTraffic(s.input)
		if w.Sum() != s.sum || w.mgvGqsp != s.index {
			t.Fatalf("step%d sum%d index%d", i, w.Sum(), w.mgvGqsp)
		}
	}
	var empty Mblp8L
	empty.AddTraffic(0)
	defer func() {
		if recover() == nil {
			t.Fatal("nonzero empty-ring update must panic")
		}
	}()
	empty.AddTraffic(1)
}
