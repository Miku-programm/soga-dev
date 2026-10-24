package gobwaspool

import (
	"reflect"
	"testing"
)

func TestNativeCapacityClasses(t *testing.T) {
	for _, tc := range []struct{ n, c int }{{0, 0}, {1, 1}, {127, 128}, {128, 128}, {129, 256}, {65535, 65536}, {65536, 65536}, {65537, 65537}} {
		b := GetLen(tc.n)
		if len(b) != tc.n || cap(b) != tc.c {
			t.Fatal(tc, len(b), cap(b))
		}
		KqwWvOwoRkN(b)
	}
	if reflect.TypeOf(E0PtY88jkHw{}).Size() != 8 {
		t.Fatal("native pointer layout")
	}
}
func TestNativeGetLengthPanic(t *testing.T) {
	defer func() {
		if e := recover(); e != "requested length is greater than capacity" {
			t.Fatal(e)
		}
	}()
	defaultPoolAt5029D70.Get(2, 1)
}
