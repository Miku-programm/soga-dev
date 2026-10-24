package inbound

import (
	"recovered.local/soga/audit/matcher"
	"testing"
)

func TestDetectorUpdateRetainsOriginalIDsAndMatcherIdentity(t *testing.T) {
	b := &JWN154qJ4{}
	old := matcher.Ij9r21E0UOt3{matcher.NewPanelDetector(7, "full:a"), matcher.NewPanelDetector(8, "full:b")}
	b.e3JHP2MzU.Store(&old)
	old[0].Test("a")
	next := matcher.Ij9r21E0UOt3{matcher.NewPanelDetector(70, "full:a"), matcher.NewPanelDetector(80, "full:b")}
	b.afT7skk(next)
	if b.e3JHP2MzU.Load() != &old || (*b.e3JHP2MzU.Load())[0].GetId() != 7 {
		t.Fatal("original ignores ID-only change")
	}
	next[1] = matcher.NewPanelDetector(90, "full:c")
	b.afT7skk(next)
	got := *b.e3JHP2MzU.Load()
	if got[0] != old[0] || got[1] != next[1] {
		t.Fatal("must reuse unchanged matcher")
	}
	b.afT7skk(next[:1])
	got = *b.e3JHP2MzU.Load()
	if len(got) != 1 || got[0] != next[0] {
		t.Fatal("length change replaces all")
	}
}
