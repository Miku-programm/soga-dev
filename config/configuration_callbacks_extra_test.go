package config

import (
	"math"
	"reflect"
	"testing"
)

func TestRecoveredBlockListWhiteIDs(t *testing.T) {
	oldConfig, oldIDs := configAt509B660, blockListWhiteIDsAt5095AE0
	defer func() { configAt509B660, blockListWhiteIDsAt5095AE0 = oldConfig, oldIDs }()
	blockListWhiteIDsAt5095AE0 = map[int64]struct{}{99: {}}
	configAt509B660.U2hwBFC = []string{"1", "+2", "-3", "0004", "1", " 5", "6 ", "0x7", "8_0", "", "9223372036854775807", "9223372036854775808", "-9223372036854775808", "-9223372036854775809"}
	init0func1()
	want := map[int64]struct{}{99: {}, 1: {}, 2: {}, -3: {}, 4: {}, math.MaxInt64: {}, math.MinInt64: {}}
	if !reflect.DeepEqual(blockListWhiteIDsAt5095AE0, want) {
		t.Fatalf("got %v, want %v", blockListWhiteIDsAt5095AE0, want)
	}
	configAt509B660.U2hwBFC = nil
	init0func1()
	if !reflect.DeepEqual(blockListWhiteIDsAt5095AE0, want) {
		t.Fatal("empty refresh cleared accumulated IDs")
	}
}

func TestRecoveredObfsModeCallback(t *testing.T) {
	old := configAt5099340
	defer func() { configAt5099340 = old }()
	for _, tc := range []struct {
		input string
		want  uint8
	}{
		{"strict", 1}, {"STRICT", 1}, {"Strict_Path", 2}, {"strict_host", 3},
		{"relaxed", 0}, {"strict ", 0}, {" strict_host", 0}, {"", 0}, {"unknown", 0},
	} {
		configAt5099340.DlaO1l = tc.input
		configAt5099340.zdRAxv6CNPMA = 99
		rHAcgy()
		if got := uint8(configAt5099340.zdRAxv6CNPMA); got != tc.want {
			t.Errorf("%q: got %d, want %d", tc.input, got, tc.want)
		}
		if configAt5099340.DlaO1l != tc.input {
			t.Fatal("source mode string was modified")
		}
	}
}

func TestRecoveredIPCacheDirectoryCallback(t *testing.T) {
	old := configAt509B380
	defer func() { configAt509B380 = old }()
	for _, tc := range []struct{ input, want string }{
		{"", "/etc/soga/"}, {"/", "/"}, {"/etc/soga", "/etc/soga/"},
		{"relative///", "relative///"}, {"false", "false/"}, {"  ", "  /"},
	} {
		configAt509B380.FpIPwenv = tc.input
		configAt509B380.TsRVDHP0SA = 0
		mvo6TVLM()
		if got := configAt509B380.FpIPwenv; got != tc.want {
			t.Errorf("%q: got %q, want %q", tc.input, got, tc.want)
		}
		if configAt509B380.TsRVDHP0SA != 24 {
			t.Fatal("nonpositive cache time not normalized")
		}
	}
	configAt509B380.TsRVDHP0SA = 1
	mvo6TVLM()
	if configAt509B380.TsRVDHP0SA != 1 {
		t.Fatal("positive cache time modified")
	}
}
