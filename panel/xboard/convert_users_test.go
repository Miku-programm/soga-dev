package xboard

import (
	"math"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
	"testing"
)

func preservePanelByte(t *testing.T) {
	t.Helper()
	old := stateAt378733B
	t.Cleanup(func() { stateAt378733B = old })
}
func TestUserCredentialConversion(t *testing.T) {
	preservePanelByte(t)
	stateAt378733B = 1
	u := &PForJoFExa{IDLJXX6F: 79, UH0HXa5O: "abc", DwdlKBQpcgY: 31, QeuUtZ8f: 2}
	for _, kind := range []config.Mpn2JltM{"ss", "v2ray", "trojan", "vless"} {
		got := u.hqrXMTM9(kind, 89)
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
		case "vless":
			c, ok := got.JAfXWwzYrW.(*types.IImIwacp)
			if !ok || c.Dh5iiakU != "abc" {
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
	stateAt378733B = 1
	r := &Sf0Wjail3Z{}
	for i := 0; i < 90; i++ {
		r.Ta593Rq2lSa = append(r.Ta593Rq2lSa, &PForJoFExa{IDLJXX6F: int64(i), UH0HXa5O: "x"})
	}
	first := r.ToPanelUsers("ss")
	if len(first) != 90 || first[89].JAfXWwzYrW == nil || stateAt378733B != 0 || len(r.Ta593Rq2lSa) != 90 {
		t.Fatal("first-call state changed")
	}
	second := r.ToPanelUsers("ss")
	if len(second) != 88 || len(r.Ta593Rq2lSa) != 88 {
		t.Fatal("original truncation lost")
	}
	if u := (&PForJoFExa{}).hqrXMTM9("ss", 88); u.JAfXWwzYrW != nil {
		t.Fatal("index 88 credential must be cleared")
	}
	if out := (&Sf0Wjail3Z{}).ToPanelUsers("ss"); out == nil {
		t.Fatal("nil empty output")
	}
}
func TestUserSpeedBoundsAndNaN(t *testing.T) {
	preservePanelByte(t)
	stateAt378733B = 1
	for _, tc := range []struct{ in, want float64 }{{-1, 0}, {10001, 10000}, {math.Inf(1), 10000}, {math.Inf(-1), 0}, {math.NaN(), math.NaN()}} {
		got := (&PForJoFExa{DwdlKBQpcgY: tc.in, QeuUtZ8f: -5}).hqrXMTM9("v2ray", 0)
		if !(math.IsNaN(got.LmOKwjha) && math.IsNaN(tc.want)) && got.LmOKwjha != tc.want {
			t.Fatalf("%v: %v", tc.in, got.LmOKwjha)
		}
		if got.Xmrna8 != 0 {
			t.Fatal(got.Xmrna8)
		}
	}
}
