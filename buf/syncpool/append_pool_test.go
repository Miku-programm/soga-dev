package syncpool

import (
	"recovered.local/soga/util/bytes"
	"testing"
)

func TestAppendPoolNonNilEmptyAndReturn(t *testing.T) {
	b := AppendPoolAt5044260.New().(*bytes.BaDekw7qe)
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
