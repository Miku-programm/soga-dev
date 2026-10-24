package MFFa9yT

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"

	"recovered.local/soga/CAodbhZf"
)

// Expected effects below come from independent execution of all four original
// closures in a bounded Unicorn harness. File deletion and exit are recorded
// only; this test never calls either OS operation or a licensing endpoint.
func TestVMessStateMachineFixtures(t *testing.T) {
	for _, tc := range []struct {
		code          string
		empty         bool
		before, after int
		effects       string
	}{
		{"0", false, 9, 0, "key,check"},
		{"0", false, 10, 0, "key,check"},
		{"100000", false, 0, 100, "key,check,exit:0"},
		{"100001", false, 0, 0, "key,check,state:true"},
		{"100001", false, 4, 0, "key,check,state:true"},
		{"100001", false, 99, 0, "key,check,state:true"},
		{"100001", true, 4, 0, "key,state:true,check,state:true"},
		{"100002", false, 0, 1, "key,check,exit:0"},
		{"100003", false, 0, 1, "key,check"},
		{"100003", false, 9, 10, "key,check"},
		{"100003", false, 10, 11, "key,check,exit:0"},
		{"100003", false, 99, 100, "key,check,exit:0"},
		{"100004", false, 0, 1, "key,check"},
		{"100005", false, 0, 1, "key,check,exit:0"},
		{"100006", false, 0, 100, "key,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"100007", false, 0, 1, "key,check"},
		{"100008", false, 0, 100, "key,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"100009", false, 0, 100, "key,check,exit:0"},
		{"100010", false, 0, 100, "key,check,exit:0"},
		{"100011", false, 0, 100, "key,check,exit:0"},
		{"100012", false, 0, 100, "key,check,exit:0"},
		{"200000", false, 99, 199, "key,check,exit:0"},
		{"200001", false, 99, 199, "key,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"200002", false, 99, 199, "key,check,exit:0"},
		{"error", false, 0, 1, "key,check,key"},
		{"error", false, 4, 5, "key,check,key"},
		{"error", false, 9, 10, "key,check,key"},
		{"error", false, 10, 11, "key,check,key"},
		{"error", true, 4, 4, "key,state:true,check,key"},
		{"unknown", false, 0, 100, "key,check,exit:0"},
	} {
		t.Run(fmt.Sprintf("%s/%t/%d", tc.code, tc.empty, tc.before), func(t *testing.T) {
			counter := tc.before
			var events []string
			key := func() string {
				events = append(events, "key")
				if tc.empty {
					return ""
				}
				return "fixture-key"
			}
			check := func() (string, error) {
				events = append(events, "check")
				if tc.code == "error" {
					return "", errors.New("offline fixture")
				}
				return tc.code, nil
			}
			set := func(v bool) { events = append(events, fmt.Sprintf("state:%t", v)) }
			rmAll := func(p string) error {
				events = append(events, "removeAll:"+p)
				return errors.New("ignored original removal error")
			}
			rm := func(p string) error {
				events = append(events, "remove:"+p)
				return errors.New("ignored original removal error")
			}
			exit := func(code int) { events = append(events, fmt.Sprintf("exit:%d", code)) }
			connectionStateResponse(&counter, key, check, set, rmAll, rm, exit)
			if counter != tc.after || strings.Join(events, ",") != tc.effects {
				t.Fatalf("counter %d, events %v; want counter %d, events %s", counter, events, tc.after, tc.effects)
			}
		})
	}
}

func TestVMessStateCheckUint32Boundaries(t *testing.T) {
	for _, tc := range []struct {
		previous uint32
		claim    bool
	}{
		{6399, true}, {6400, false}, {6401, false}, {10001, true}, {0xffffffff, true},
	} {
		var last atomic.Uint32
		last.Store(tc.previous)
		if got := claimConnectionStateCheck(&last, 10000); got != tc.claim {
			t.Fatalf("previous %d: %v", tc.previous, got)
		}
		want := tc.previous
		if tc.claim {
			want = 10000
		}
		if last.Load() != want {
			t.Fatal("CAS result")
		}
	}
	var last atomic.Uint32
	var winners atomic.Int32
	done := make(chan struct{}, 32)
	for range 32 {
		go func() {
			if claimConnectionStateCheck(&last, 10000) {
				winners.Add(1)
			}
			done <- struct{}{}
		}()
	}
	for range 32 {
		<-done
	}
	if winners.Load() != 1 {
		t.Fatalf("claim winners %d", winners.Load())
	}
}

func TestVMessStateMethodsWithinInterval(t *testing.T) {
	old := userCountAt50C6DA0
	defer func() { userCountAt50C6DA0 = old }()
	v := &K4avFiYOPi{JWN154qJ4: &JWN154qJ4{fUxO7UYAH: 83, mKB02yhVfTfP: 4}}
	// This is the original proved no-goroutine branch, not an authorization bypass.
	v.jfLH8mvq5e03.Store(uint32(CAodbhZf.C5Jff5vs0()))
	for _, method := range []func(){v.sbbdX3aT6E, v.bj10P7X} {
		method()
		if userCountAt50C6DA0 != 83 || v.mKB02yhVfTfP != 4 {
			t.Fatal("state changed within interval")
		}
	}
}
