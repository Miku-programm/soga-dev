package _vl4eANj83Ck

import (
	"reflect"
	"testing"
)

func TestNativePowerOfTwoBoundaries(t *testing.T) {
	for _, tc := range [][2]int{{-3, -3}, {0, 0}, {1, 1}, {2, 2}, {3, 4}, {128, 128}, {129, 256}, {1 << 62, 1 << 62}} {
		if n := WDHhd0lpnI3(tc[0]); n != tc[1] {
			t.Fatal(tc, n)
		}
	}
	var got []int
	YxT95tOm(0, 9, func(n int) { got = append(got, n) })
	if !reflect.DeepEqual(got, []int{1, 2, 4, 8}) {
		t.Fatal(got)
	}
	defer func() {
		if e := recover(); e != "argument is too large" {
			t.Fatal(e)
		}
	}()
	WDHhd0lpnI3((1 << 62) + 1)
}
