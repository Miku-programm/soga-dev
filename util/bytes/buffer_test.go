package bytes

import (
	"bytes"
	"testing"
)

func TestGrowPreservesDataAndNilDistinction(t *testing.T) {
	var b BaDekw7qe
	if b.Grow(1) != 0 || len(b) != 1 || cap(b) != 64 {
		t.Fatal(len(b), cap(b))
	}
	b[0] = 9
	if b.Grow(64) != 1 || len(b) != 65 || b[0] != 9 {
		t.Fatal(b)
	}
	if b.Grow(-1) != 65 || len(b) != 64 {
		t.Fatal(len(b))
	}
	nonnil := BaDekw7qe{}
	nonnil.Grow(1)
	if cap(nonnil) == 64 {
		t.Fatal("nonnil empty must bypass nil-only 64-byte allocation")
	}
}
func TestOversizePanic(t *testing.T) {
	b := make(BaDekw7qe, 1, 1)
	defer func() {
		if got := recover(); got != bytes.ErrTooLarge {
			t.Fatalf("panic=%v", got)
		}
	}()
	b.Grow(int(^uint(0) >> 1))
}
