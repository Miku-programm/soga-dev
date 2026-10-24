package genmap

import (
	"sync"
	"testing"
)

func TestConcurrentMapValueIdentityAndCount(t *testing.T) {
	var m GTFAWF[string, *int]
	if value, ok := m.Get("absent"); value != nil || ok {
		t.Fatal(value, ok)
	}
	a, b := 1, 2
	m.Set("key", &a)
	m.Set("key", &b)
	if m.yaLIYjUJ.Load() != 1 {
		t.Fatal("replacement increments count")
	}
	if value, loaded := m.GetOrSet("key", &a); value != &b || !loaded {
		t.Fatal(value, loaded)
	}
	if value, loaded := m.GetOrSet("new", nil); value != nil || loaded {
		t.Fatal(value, loaded)
	}
	if value, loaded := m.Get("new"); value != nil || !loaded {
		t.Fatal("typed nil loses presence")
	}
	if value, loaded := m.GetAndDelete("key"); value != &b || !loaded {
		t.Fatal(value, loaded)
	}
	if value, loaded := m.GetAndDelete("key"); value != nil || loaded {
		t.Fatal(value, loaded)
	}
	m.Delete("missing")
	m.Delete("new")
	if m.yaLIYjUJ.Load() != 0 {
		t.Fatal(m.yaLIYjUJ.Load())
	}
}

func TestConcurrentMapAtomicInsertion(t *testing.T) {
	var m GTFAWF[string, *int]
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); value := 1; m.GetOrSet("one", &value) }()
	}
	wg.Wait()
	if m.yaLIYjUJ.Load() != 1 {
		t.Fatal(m.yaLIYjUJ.Load())
	}
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); m.Delete("one") }()
	}
	wg.Wait()
	if m.yaLIYjUJ.Load() != 0 {
		t.Fatal(m.yaLIYjUJ.Load())
	}
}

func TestConcurrentMapPreservesTypeAssertionPanic(t *testing.T) {
	var m GTFAWF[string, *int]
	m.r5_Hdcu3GxXC.Store("wrong", 7)
	defer func() {
		if recover() == nil {
			t.Fatal("wrong stored type must panic")
		}
	}()
	m.Get("wrong")
}
