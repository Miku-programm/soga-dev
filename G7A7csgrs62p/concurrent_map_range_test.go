package G7A7csgrs62p

import "testing"

func TestRangeAndCount(t *testing.T) {
	var m GTFAWF[int, string]
	if m.Len() != 0 {
		t.Fatal(m.Len())
	}
	m.Set(1, "a")
	m.Set(1, "b")
	m.Set(2, "c")
	if m.Len() != 2 {
		t.Fatal(m.Len())
	}
	seen := map[int]string{}
	m.Range(func(k int, v string) bool { seen[k] = v; return true })
	if len(seen) != 2 || seen[1] != "b" || seen[2] != "c" {
		t.Fatal(seen)
	}
	calls := 0
	m.Range(func(int, string) bool { calls++; return false })
	if calls != 1 {
		t.Fatal(calls)
	}
	m.Delete(1)
	m.GetAndDelete(2)
	if m.Len() != 0 {
		t.Fatal(m.Len())
	}
}
func TestRangeNilInterfacePanics(t *testing.T) {
	var m GTFAWF[int, any]
	m.Set(1, nil)
	defer func() {
		if recover() == nil {
			t.Fatal("原类型断言应拒绝 nil 接口")
		}
	}()
	m.Range(func(int, any) bool { return true })
}
