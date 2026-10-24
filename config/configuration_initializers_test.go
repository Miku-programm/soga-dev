package config

import (
	"reflect"
	"testing"
)

func TestRecoveredConfigurationInitOrder(t *testing.T) {
	isolateRegistry(t)
	oldBlock, oldWhite, oldVersion := blockListAt5095A90.Load(), whiteListAt5095AD8.Load(), versionAt5029EE0
	t.Cleanup(func() {
		blockListAt5095A90.Store(oldBlock)
		whiteListAt5095AD8.Store(oldWhite)
		versionAt5029EE0 = oldVersion
	})
	versionAt5029EE0 = " \n2.17.0\u3000"
	initializeRegisteredConfigurations()
	want := []any{&configAt509B660, &configAt50A2260, &configAt50992E0, &nodeConfigurationAt50A2AE0, &configAt5096DC0, &configAt509F9C0, &configAt509CB40, &configAt509E2A0, &configAt50C8440, &configAt509B380, &configAt509B6A0, &configAt50C6C49, &configAt509BEA0, &configAt50A2900, &configAt509B3A0, &configAt509B3C0, &configAt509C940, &configAt509CB80, &configAt5096DD0, &configAt509B3E0, &configAt5099340, &configAt509B6E0, &configAt509EBE0, &configAt5099360, &configAt509E300, &configAt50A02C0, &configAt50A2300, &configAt5096E00}
	if len(registeredAt5099320) != len(want) {
		t.Fatalf("got %d registrations, want %d", len(registeredAt5099320), len(want))
	}
	for i, pointer := range want {
		if registeredAt5099320[i] != pointer {
			t.Errorf("registration %d differs: %T", i, registeredAt5099320[i])
		}
	}
	callbacks := []func(){init0func1, uQm8ubPKg, jkMcB1_, init6func1, vlfjWv1b4Bbg, mvo6TVLM, init10func1, qZO6G3A, tzfa9ae, init14func1, init15func1, ekZtUL9pJS, mQ5mMKnkLjl, rHAcgy, eifTpu4B}
	if len(afterLoadAt5099300) != len(callbacks) {
		t.Fatalf("got %d callbacks, want %d", len(afterLoadAt5099300), len(callbacks))
	}
	for i, callback := range callbacks {
		if reflect.ValueOf(afterLoadAt5099300[i]).Pointer() != reflect.ValueOf(callback).Pointer() {
			t.Errorf("callback %d differs", i)
		}
	}
	if versionAt5029EE0 != "2.17.0" {
		t.Fatalf("version not trimmed: %q", versionAt5029EE0)
	}
	block, white := blockListAt5095A90.Load(), whiteListAt5095AD8.Load()
	if block == nil || white == nil || *block == nil || *white == nil || len(*block) != 0 || cap(*block) != 0 || len(*white) != 0 || cap(*white) != 0 {
		t.Fatalf("expected distinct initialized pointers to nonnil empty named slices: %v, %v", block, white)
	}
	if block == white {
		t.Fatal("original initializers allocate separate slice headers")
	}
}
