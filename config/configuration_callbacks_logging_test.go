package config

import "testing"

func TestRecoveredAccessLogNormalization(t *testing.T) {
	old := configAt50A2260
	defer func() { configAt50A2260 = old }()
	for _, tc := range []struct{ mode, directory, wantMode, wantDirectory string }{
		{"", "false", "local", ""}, {"false", "/", "local", ""},
		{"  ClickHouse\n", "/etc/soga///", "clickhouse", "/etc/soga"},
		{"unknown", "relative", "local", "relative"}, {"FALSE", "FALSE", "local", "FALSE"},
	} {
		configAt50A2260 = tUl2hoOYI0Ao{PmIFLM: true, WO1HZMaFg6Mx: tc.mode, INBdnGu56l5: tc.directory, EuOaOO8qBn0Y: 72}
		uQm8ubPKg()
		c := configAt50A2260
		if c.WO1HZMaFg6Mx != tc.wantMode || c.INBdnGu56l5 != tc.wantDirectory {
			t.Errorf("%+v -> mode %q directory %q", tc, c.WO1HZMaFg6Mx, c.INBdnGu56l5)
		}
		if c.EuOaOO8qBn0Y != 72 {
			t.Fatal("retention changed instead of warning")
		}
	}
}

func TestRecoveredRedisMigrationAndValidation(t *testing.T) {
	oldLegacy, oldCurrent := configAt509C940, configAt509CB80
	defer func() { configAt509C940, configAt509CB80 = oldLegacy, oldCurrent }()
	configAt509C940 = x7sf3Tkb{H2FCRVc: true, IDyeXiEv: true, Ij4bWLsuo_x4: "legacy:6379", A9GLBE4PY: "test", ZDYXh3Cuvxpw: 3, Eb9vzdoE: 25}
	configAt509CB80 = vHzUrjrUZcY8{FXMKW6VcHH: "unused", BFkqcX7Dx: 12}
	mQ5mMKnkLjl()
	c := configAt509CB80
	if !c.KCYNbwBsZ || c.FXMKW6VcHH != "legacy:6379" || c.AhScWSQ != "test" || c.ADewfG9Z64W != 3 || !c.NDRQGsmc132Y || c.BFkqcX7Dx != 25 {
		t.Fatalf("migration differs: %+v", c)
	}
	configAt509CB80 = vHzUrjrUZcY8{KCYNbwBsZ: true, FXMKW6VcHH: "current:6379", BFkqcX7Dx: 12}
	mQ5mMKnkLjl()
	if configAt509CB80.FXMKW6VcHH != "current:6379" || configAt509CB80.BFkqcX7Dx != 12 || configAt509CB80.NDRQGsmc132Y {
		t.Fatal("enabled new settings were overwritten")
	}
	configAt509C940.H2FCRVc = false
	configAt509CB80 = vHzUrjrUZcY8{KCYNbwBsZ: true, BFkqcX7Dx: -1}
	mQ5mMKnkLjl()
	if configAt509CB80.KCYNbwBsZ || configAt509CB80.BFkqcX7Dx != 60 {
		t.Fatal("missing address or expiry validation differs")
	}
	configAt509CB80 = vHzUrjrUZcY8{KCYNbwBsZ: true, FXMKW6VcHH: " ", BFkqcX7Dx: 1}
	mQ5mMKnkLjl()
	if !configAt509CB80.KCYNbwBsZ {
		t.Fatal("address was trimmed although original only checks length")
	}
}

func TestRecoveredVMessLegacySwitch(t *testing.T) {
	old := configAt50A02C0
	defer func() { configAt50A02C0 = old }()
	configAt50A02C0 = xIoYVOH{SVuGy3CoRze: true, ZFK_OBT1: "md5"}
	eifTpu4B()
	if configAt50A02C0.dZWc0A2uIe != "aead" || configAt50A02C0.ZFK_OBT1 != "md5" {
		t.Fatal("legacy forced mode mapping differs")
	}
	configAt50A02C0.SVuGy3CoRze = false
	eifTpu4B()
	if configAt50A02C0.dZWc0A2uIe != "md5" {
		t.Fatal("normal mode not restored")
	}
}
