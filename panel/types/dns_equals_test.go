package types

import "testing"

func TestOriginalDNSRuleEquality(t *testing.T) {
	if !(U863Mm01Nre(nil)).Equals(U863Mm01Nre{}) {
		t.Fatal("nil/empty list")
	}
	a := &Axa8uKIB4JL{K8BNfziIgb: []string{"a", "b"}, CM0eMsqyP: "v4", CZ80ik: []string{"x", "y"}}
	b := &Axa8uKIB4JL{K8BNfziIgb: []string{"a", "b"}, CM0eMsqyP: "v4", CZ80ik: []string{"x", "y"}}
	if !a.Equals(b) {
		t.Fatal("content equality")
	}
	b.K8BNfziIgb[0], b.K8BNfziIgb[1] = "b", "a"
	if a.Equals(b) {
		t.Fatal("order lost")
	}
	if (U863Mm01Nre{a}).Equals(nil) {
		t.Fatal("length lost")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("original nil rule panic must remain")
		}
	}()
	(U863Mm01Nre{nil}).Equals(U863Mm01Nre{nil})
}
