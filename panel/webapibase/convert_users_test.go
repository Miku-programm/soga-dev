package webapibase

import (
	"math"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
	"testing"
)

func preservePanelByte(t *testing.T) {
	t.Helper()
	old := stateAt3787339
	t.Cleanup(func() { stateAt3787339 = old })
}
func TestUserCredentialConversion(t *testing.T) {
	preservePanelByte(t)
	stateAt3787339 = 1
	u := &WVfKv8zPvt{IDLJXX6F: 79, UH0HXa5O: "abc", DwdlKBQpcgY: 31, QeuUtZ8f: 2}
	for _, kind := range []config.Mpn2JltM{"ss", "v2ray", "trojan", "vless"} {
		got := u.wfpy4HFAx21(kind, 89)
		if got.WZ89pTAOrHK4 != 79 || got.LmOKwjha != 31 || got.Xmrna8 != 2 {
			t.Fatalf("%+v", got)
		}
		switch kind {
		case "ss":
			c, ok := got.JAfXWwzYrW.(*types.YlfOV0DBBnx)
			if !ok || c.A_wNUTkMh6pH != "abc" || c.T1oDDIgYJHh != "" {
				t.Fatal(got)
			}
		case "v2ray":
			c, ok := got.JAfXWwzYrW.(*types.Yn8BJoL)
			if !ok || c.SfiEGpJo != "abc" {
				t.Fatal(got)
			}
		case "trojan":
			c, ok := got.JAfXWwzYrW.(*types.JxTGUar)
			if !ok || c.ARARP1z0yA != "abc" {
				t.Fatal(got)
			}
		default:
			if got.JAfXWwzYrW != nil {
				t.Fatal(got)
			}
		}
	}
}
func TestUsersOriginalFirstCallAndLimit(t *testing.T) {
	preservePanelByte(t)
	stateAt3787339 = 1
	r := &D9MPLRTuA{}
	for i := 0; i < 90; i++ {
		r.Ta593Rq2lSa = append(r.Ta593Rq2lSa, &WVfKv8zPvt{IDLJXX6F: int64(i), UH0HXa5O: "x"})
	}
	first := r.ToPanelUsers("ss")
	if len(first) != 90 || first[89].JAfXWwzYrW == nil || stateAt3787339 != 0 || len(r.Ta593Rq2lSa) != 90 {
		t.Fatal("first-call state changed")
	}
	second := r.ToPanelUsers("ss")
	if len(second) != 90 || len(r.Ta593Rq2lSa) != 90 {
		t.Fatal("user list must not be truncated")
	}
	if u := (&WVfKv8zPvt{}).wfpy4HFAx21("ss", 88); u.JAfXWwzYrW == nil {
		t.Fatal("index 88 credential must be present")
	}
	if out := (&D9MPLRTuA{}).ToPanelUsers("ss"); out == nil {
		t.Fatal("nil empty output")
	}
}
func TestUserSpeedBoundsAndNaN(t *testing.T) {
	preservePanelByte(t)
	stateAt3787339 = 1
	for _, tc := range []struct{ in, want float64 }{{-1, 0}, {10001, 10000}, {math.Inf(1), 10000}, {math.Inf(-1), 0}, {math.NaN(), math.NaN()}} {
		got := (&WVfKv8zPvt{DwdlKBQpcgY: tc.in, QeuUtZ8f: -5}).wfpy4HFAx21("v2ray", 0)
		if !(math.IsNaN(got.LmOKwjha) && math.IsNaN(tc.want)) && got.LmOKwjha != tc.want {
			t.Fatalf("%v: %v", tc.in, got.LmOKwjha)
		}
		if got.Xmrna8 != 0 {
			t.Fatal(got.Xmrna8)
		}
	}
}
