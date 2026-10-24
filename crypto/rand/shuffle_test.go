package rand

import "testing"

func TestShuffleOriginalRandomSourceAndOrder(t *testing.T) {
	withRandomWords(t, 0x0000000100000000, 0x0000000200000000, 0)
	values := []int{0, 1, 2, 3}
	var calls [][2]int
	Tij6hx(len(values), func(i, j int) { calls = append(calls, [2]int{i, j}); values[i], values[j] = values[j], values[i] })
	want := [][2]int{{3, 1}, {2, 2}, {1, 0}}
	for i := range want {
		if calls[i] != want[i] {
			t.Fatalf("calls %v", calls)
		}
	}
	if values[0] != 3 || values[1] != 0 || values[2] != 2 || values[3] != 1 {
		t.Fatalf("values %v", values)
	}
	Tij6hx(0, func(int, int) { t.Fatal("empty shuffle called swap") })
	Tij6hx(1, func(int, int) { t.Fatal("singleton shuffle called swap") })
}
func TestShuffleNegativePanics(t *testing.T) {
	defer func() {
		if v := recover(); v != "invalid argument to Shuffle" {
			t.Fatalf("panic %v", v)
		}
	}()
	Tij6hx(-1, nil)
}
func TestShuffleLargeRangeUsesInt63n(t *testing.T) {
	withRandomWords(t, 5)
	defer func() {
		if v := recover(); v != "stop after first swap" {
			t.Fatalf("panic %v", v)
		}
	}()
	Tij6hx(1<<31, func(i, j int) {
		if i != (1<<31)-1 || j != 5 {
			t.Fatalf("swap(%d,%d)", i, j)
		}
		panic("stop after first swap")
	})
}
