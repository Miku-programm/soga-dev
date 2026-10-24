package gobwaspool2

import (
	"reflect"
	"testing"
)

func TestNativeMappedAndMissingClasses(t *testing.T) {
	p := L06Zy06(128, 65536)
	if len(p.lxoG0T) != 10 {
		t.Fatal(len(p.lxoG0T))
	}
	for _, tc := range []struct{ n, want int }{{0, 0}, {3, 3}, {127, 128}, {128, 128}, {129, 256}, {65535, 65536}, {65537, 65537}} {
		v, n := p.Get(tc.n)
		if v != nil || n != tc.want {
			t.Fatal(tc, v, n)
		}
	}
	p.Put("wrong class", 129)
	if _, exists := p.lxoG0T[129]; exists {
		t.Fatal("Put must not register classes")
	}
	if reflect.TypeOf(LVPKRSqpONvB{}).Size() != 16 {
		t.Fatal("layout")
	}
}
func TestNativeCustomOptionsOrder(t *testing.T) {
	var calls []int
	p := QPgdbVKbwji(func(c MpotdXINa2) { calls = append(calls, 1); c.AddSize(3) }, func(c MpotdXINa2) { calls = append(calls, 2); c.SetSizeMapping(func(int) int { return 3 }) })
	if !reflect.DeepEqual(calls, []int{1, 2}) {
		t.Fatal(calls)
	}
	if _, n := p.Get(8); n != 3 {
		t.Fatal(n)
	}
}
