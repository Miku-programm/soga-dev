package tkbUwvW3ex

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
	"testing"
	"unsafe"
)

func TestAliveIPLayoutAndConstruction(t *testing.T) {
	a := NewJW2SKH2ajwT6()
	if unsafe.Sizeof(*a) != 32 || unsafe.Offsetof(a.x8InevskM) != 8 || unsafe.Offsetof(a.rl2GEL) != 16 || unsafe.Offsetof(a.gRUHgk0Uufx) != 24 {
		t.Fatal("original type layout changed")
	}
	if a.aM9TY50Pn == nil || a.x8InevskM == nil || a.rl2GEL == nil {
		t.Fatal("constructor must allocate all three maps")
	}
	got := a.GetAndClearAliveIps()
	if got == nil || len(got) != 0 || cap(got) != 0 {
		t.Fatalf("empty result: %#v", got)
	}
}

func TestAliveIPLimitAdmission(t *testing.T) {
	a := NewJW2SKH2ajwT6()
	steps := []struct {
		first, second string
		limit         int
		want          bool
	}{
		{"device1", "192.0.2.1", 1, true},
		{"device1", "192.0.2.2", 1, true},
		{"device2", "192.0.2.1", 1, false},
		{"device2", "192.0.2.3", 2, true},
		{"device1", "192.0.2.4", 1, false},
		{"device3", "192.0.2.5", 0, true},
		{"device4", "192.0.2.6", -1, true},
	}
	for _, step := range steps {
		if got := a.AddCount(step.first, step.second, step.limit); got != step.want {
			t.Fatalf("AddCount(%q,%q,%d)=%v", step.first, step.second, step.limit, got)
		}
	}
	if !reflect.DeepEqual(a.aM9TY50Pn, map[string]int{"device1": 2, "device2": 1, "device3": 1, "device4": 1}) || len(a.x8InevskM) != 5 {
		t.Fatalf("counts first=%v second=%v", a.aM9TY50Pn, a.x8InevskM)
	}
	if got := a.GetAndClearAliveIps(); len(got) != 0 {
		t.Fatalf("AddCount must not call Record: %v", got)
	}
}

func TestAliveIPReportRetainsOnlyActiveRecords(t *testing.T) {
	a := NewJW2SKH2ajwT6()
	a.AddCount("device1", "192.0.2.1", 0)
	a.AddCount("device2", "192.0.2.1", 0)
	a.Record("192.0.2.1")
	a.Record("192.0.2.1")
	a.Record("2001:db8::1")
	got := a.GetAndClearAliveIps()
	sort.Strings(got)
	if !reflect.DeepEqual(got, []string{"192.0.2.1", "2001:db8::1"}) || cap(got) != 2 {
		t.Fatalf("first report %v cap=%d", got, cap(got))
	}
	a.SubCount("device1", "192.0.2.1")
	if got := a.GetAndClearAliveIps(); !reflect.DeepEqual(got, []string{"192.0.2.1"}) {
		t.Fatalf("shared IP still active: %v", got)
	}
	a.SubCount("device2", "192.0.2.1")
	if got := a.GetAndClearAliveIps(); !reflect.DeepEqual(got, []string{"192.0.2.1"}) {
		t.Fatalf("last disconnected IP not reported: %v", got)
	}
	if got := a.GetAndClearAliveIps(); len(got) != 0 {
		t.Fatalf("stale record retained: %v", got)
	}
}

func TestAliveIPSubCountIndependentAndNonpositive(t *testing.T) {
	a := NewJW2SKH2ajwT6()
	a.aM9TY50Pn["first"] = 2
	a.x8InevskM["second"] = 2
	a.SubCount("absent", "second")
	a.SubCount("first", "absent")
	if a.aM9TY50Pn["first"] != 1 || a.x8InevskM["second"] != 1 {
		t.Fatal("keys were not independently decremented")
	}
	a.SubCount("first", "second")
	a.aM9TY50Pn["negative"] = -1
	a.x8InevskM["zero"] = 0
	a.Record("zero")
	// Membership, not a positive counter, keeps a record alive.
	a.GetAndClearAliveIps()
	if _, ok := a.rl2GEL["zero"]; !ok {
		t.Fatal("zero count key lost its record")
	}
	a.SubCount("negative", "zero")
	if len(a.aM9TY50Pn) != 0 || len(a.x8InevskM) != 0 {
		t.Fatal("nonpositive counters must be removed")
	}
}

func TestAliveIPZeroValuePanicReleasesMutex(t *testing.T) {
	for _, operation := range []string{"add", "record"} {
		t.Run(operation, func(t *testing.T) {
			a := new(JW2SKH2ajwT6)
			var recovered any
			func() {
				defer func() { recovered = recover() }()
				if operation == "add" {
					a.AddCount("a", "b", 0)
				} else {
					a.Record("b")
				}
			}()
			if recovered == nil {
				t.Fatal("nil maps must retain original panic")
			}
			if !a.gRUHgk0Uufx.TryLock() {
				t.Fatal("deferred unlock not performed")
			}
			a.gRUHgk0Uufx.Unlock()
			a.SubCount("absent", "absent")
			if a.GetAndClearAliveIps() == nil {
				t.Fatal("empty report must be allocated")
			}
		})
	}
}

func TestAliveIPConcurrentAdmissionAndRecords(t *testing.T) {
	a := NewJW2SKH2ajwT6()
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprint(i)
			for j := 0; j < 50; j++ {
				if a.AddCount(key, "192.0.2.1", 4) {
					a.Record("192.0.2.1")
					a.GetAndClearAliveIps()
					a.SubCount(key, "192.0.2.1")
				}
			}
		}(i)
	}
	wg.Wait()
	if len(a.aM9TY50Pn) != 0 || len(a.x8InevskM) != 0 {
		t.Fatal("unbalanced counts")
	}
	a.GetAndClearAliveIps()
	if len(a.GetAndClearAliveIps()) != 0 {
		t.Fatal("inactive records were not cleared")
	}
}
