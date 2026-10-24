package bloom

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestBloomOriginalFormulaAndBitOffsets(t *testing.T) {
	f := New(100000, 1e-12, func([]byte) (uint64, uint64) { return 7, 8 }).(*classicFilter)
	if unsafe.Sizeof(*f) != 40 || f.k != 39 || f.Size() != 718879 {
		t.Fatal(f.k, f.Size())
	}
	f = New(1, .1, func([]byte) (uint64, uint64) { return ^uint64(0), 2 }).(*classicFilter)
	if f.k != 3 || f.Size() != 0 {
		t.Fatal(f.k, f.Size())
	}
	// 一字节位图明确验证uint64加法溢出：偏移7、1、3。
	f = &classicFilter{b: make([]byte, 1), k: 3, h: func([]byte) (uint64, uint64) { return ^uint64(0), 2 }}
	if f.Test(nil) {
		t.Fatal("empty")
	}
	f.Add(nil)
	if !reflect.DeepEqual(f.b, []byte{0x8a}) || !f.Test(nil) {
		t.Fatal(f.b)
	}
	f.Reset()
	if f.b[0] != 0 || f.Test(nil) {
		t.Fatal("reset")
	}
}
