package mzi1d3gDO7u

import (
	"testing"
	"unsafe"
)

func TestOriginalPoolClasses(t *testing.T) {
	if unsafe.Sizeof(L4upasOm{}) != 24 {
		t.Fatal("layout")
	}
	p := newSharedPool()
	for _, v := range []struct{ n, c int }{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {1025, 2048}, {20480, 32768}, {65536, 65536}, {65537, 65537}} {
		b := p.Get(v.n)
		if len(b) != v.n || cap(b) != v.c {
			t.Fatalf("%d: %d/%d", v.n, len(b), cap(b))
		}
		if err := p.Put(b); err != nil {
			t.Fatal(err)
		}
	}
	if err := p.Put(make([]byte, 1, 3)); err == nil || err.Error() != "allocator Put() incorrect buffer size" {
		t.Fatal(err)
	}
	defer func() {
		if r := recover(); r != "alloc.Get: len out of range" {
			t.Fatalf("panic %v", r)
		}
	}()
	p.Get(-1)
}
