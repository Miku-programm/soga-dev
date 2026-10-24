package Hx04IbKk

import "testing"

func TestPutSliceAliasAndReset(t *testing.T) {
	b := BaDekw7qe(make([]byte, 3, 32))
	copy(b, "abc")
	b.PutSlice(b[:2])
	if string(b) != "abcab" {
		t.Fatal(string(b))
	}
	view := b.Bytes()
	view[0] = 'X'
	if b[0] != 'X' {
		t.Fatal("Bytes copied")
	}
	b.Reset()
	if b == nil || len(b) != 0 || cap(b) != 32 {
		t.Fatal(b)
	}
	var nilBuffer BaDekw7qe
	nilBuffer.Reset()
	if nilBuffer != nil {
		t.Fatal("nil changed")
	}
}
