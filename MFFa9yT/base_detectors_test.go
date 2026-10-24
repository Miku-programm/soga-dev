package MFFa9yT

import (
	"recovered.local/soga/xx5fRyrFw"
	"testing"
)

func TestDetectorUpdateRetainsOriginalIDsAndMatcherIdentity(t *testing.T) {
	b := &JWN154qJ4{}
	old := xx5fRyrFw.Ij9r21E0UOt3{xx5fRyrFw.NewPanelDetector(7, "full:a"), xx5fRyrFw.NewPanelDetector(8, "full:b")}
	b.e3JHP2MzU.Store(&old)
	old[0].Test("a")
	next := xx5fRyrFw.Ij9r21E0UOt3{xx5fRyrFw.NewPanelDetector(70, "full:a"), xx5fRyrFw.NewPanelDetector(80, "full:b")}
	b.afT7skk(next)
	if b.e3JHP2MzU.Load() != &old || (*b.e3JHP2MzU.Load())[0].GetId() != 7 {
		t.Fatal("original ignores ID-only change")
	}
	next[1] = xx5fRyrFw.NewPanelDetector(90, "full:c")
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
