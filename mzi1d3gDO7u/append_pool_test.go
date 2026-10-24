package mzi1d3gDO7u

import (
	"recovered.local/soga/Hx04IbKk"
	"testing"
)

func TestAppendPoolNonNilEmptyAndReturn(t *testing.T) {
	b := AppendPoolAt5044260.New().(*Hx04IbKk.BaDekw7qe)
	if *b == nil || len(*b) != 0 || cap(*b) != 0 {
		t.Fatal("original initializer mismatch")
	}
	b.PutSlice([]byte("payload"))
	capacity := cap(*b)
	Wr41CoUmtL0(b)
	if len(*b) != 0 || cap(*b) != capacity {
		t.Fatal("reset discarded allocation")
	}
}
