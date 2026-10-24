package EP3M1uC

import (
	"recovered.local/soga/E3F12urwswL"
	"testing"
	"unsafe"
)

func TestOutboundRTTILayout(t *testing.T) {
	var out GyyO51
	if unsafe.Sizeof(out) != 448 || unsafe.Offsetof(out.ucp4iC) != 328 || unsafe.Offsetof(out.mJTqKai) != 392 || unsafe.Offsetof(out.vWopFdjg) != 440 {
		t.Fatalf("outbound layout differs from RTTI: size=%d", unsafe.Sizeof(out))
	}
	if unsafe.Sizeof(E3F12urwswL.D0mLgQF{}) != 160 || unsafe.Sizeof(E3F12urwswL.Kna5ouY0MG9d{}) != 32 {
		t.Fatal("adapter metadata layout differs")
	}
}
func TestOutboundHealthTransitions(t *testing.T) {
	out := new(GyyO51)
	out.uymza6wOIV6()
	sequence := []struct{ ok, disabled, enabled, state bool }{
		{false, false, false, true}, {false, false, false, true}, {true, false, false, true},
		{false, false, false, true}, {false, false, false, true}, {false, true, false, false},
		{false, false, false, false}, {true, false, false, false}, {false, false, false, false},
		{true, false, false, false}, {true, false, true, true}, {true, false, false, true},
	}
	for i, s := range sequence {
		d, e := out.ebDTwNX(s.ok)
		if d != s.disabled || e != s.enabled || out.IsEnabled() != s.state {
			t.Fatalf("step %d: transitions=%v,%v state=%v", i, d, e, out.IsEnabled())
		}
	}
	for i := 0; i < 3; i++ {
		out.ebDTwNX(false)
	}
	out.uymza6wOIV6()
	if out.IsEnabled() {
		t.Fatal("repeated initialization revived disabled outbound")
	}
}
func TestEnabledPreservesOrderAndEmpty(t *testing.T) {
	if got := (J_enG0wE(nil)).Enabled(); got == nil || len(got) != 0 {
		t.Fatal("nil input must produce nonnil empty result")
	}
	a, b, c := new(GyyO51), new(GyyO51), new(GyyO51)
	a.uymza6wOIV6()
	c.uymza6wOIV6()
	got := (J_enG0wE{a, b, c, a}).Enabled()
	if len(got) != 3 || got[0] != a || got[1] != c || got[2] != a {
		t.Fatal("enabled order or duplicates lost")
	}
}
