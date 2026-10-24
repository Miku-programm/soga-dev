package ssrobfs

import (
	"encoding/hex"
	"testing"
)

func TestPercentOriginalTraces(t *testing.T) {
	for _, r := range []struct {
		in, out   string
		nilResult bool
		capacity  int
	}{
		{"", "", false, 0},
		{"474554202f20485454502f312e31", "", false, 0},
		{"253431", "41", false, 1},
		{"253431253432", "4142", false, 2},
		{"25343125343220485454502f312e31", "4142", false, 2},
		{"474554202f25303025666625614220485454502f312e31", "00ffab", false, 3},
		{"2561", "0a", false, 1},
		{"2561253432", "0a", false, 1},
		{"25", "", true, 0},
		{"25253431", "", true, 0},
		{"256730", "", true, 0},
		{"253467", "", true, 0},
		{"2534313432253433", "41", false, 1},
		{"25343178253432", "41", false, 1},
		{"2530256666", "00", false, 1},
		{"25343125", "", true, 0},
	} {
		data, _ := hex.DecodeString(r.in)
		got := (&nArN8vNMTaR{}).bDsaTMgl(data)
		if hex.EncodeToString(got) != r.out || (got == nil) != r.nilResult || cap(got) != r.capacity {
			t.Fatalf("input=%s got=%x nil=%v cap=%d", r.in, got, got == nil, cap(got))
		}
	}
}
