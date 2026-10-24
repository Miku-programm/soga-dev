package inbound

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// Expected effects come from 148 independent SSR closure traces at
// 0x1e96c80, 0x1e98820, 0x1e9a380 and 0x1e9bf80; all four agree.
// See ssr-state-all-traces.json. Deletes and exits are recorded only.
func TestSSRStateResponseOriginalFixtures(t *testing.T) {
	for _, tc := range []struct {
		code          string
		empty         bool
		before, after int
		effects       string
	}{
		{"0", true, 9, 0, "key,state:true,check"},
		{"100000", true, 9, 109, "key,state:true,check,exit:0"},
		{"100001", true, 9, 0, "key,state:true,check,state:true"},
		{"100002", true, 9, 10, "key,state:true,check,exit:0"},
		{"100003", true, 9, 10, "key,state:true,check"},
		{"100004", true, 9, 10, "key,state:true,check"},
		{"100005", true, 9, 10, "key,state:true,check,exit:0"},
		{"100006", true, 9, 109, "key,state:true,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"100007", true, 9, 10, "key,state:true,check"},
		{"100008", true, 9, 109, "key,state:true,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"100009", true, 9, 109, "key,state:true,check,exit:0"},
		{"100010", true, 9, 109, "key,state:true,check,exit:0"},
		{"100011", true, 9, 109, "key,state:true,check,exit:0"},
		{"100012", true, 9, 109, "key,state:true,check,exit:0"},
		{"200001", true, 9, 109, "key,state:true,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"unknown", true, 9, 109, "key,state:true,check,exit:0"},
		{"error", true, 9, 9, "key,state:true,check,key"},
		{"0", false, 9, 0, "key,check"},
		{"100000", false, 9, 109, "key,check,exit:0"},
		{"100001", false, 9, 0, "key,check,state:true"},
		{"100002", false, 9, 10, "key,check,exit:0"},
		{"100003", false, 9, 10, "key,check"},
		{"100004", false, 9, 10, "key,check"},
		{"100005", false, 9, 10, "key,check,exit:0"},
		{"100006", false, 9, 109, "key,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"100007", false, 9, 10, "key,check"},
		{"100008", false, 9, 109, "key,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"100009", false, 9, 109, "key,check,exit:0"},
		{"100010", false, 9, 109, "key,check,exit:0"},
		{"100011", false, 9, 109, "key,check,exit:0"},
		{"100012", false, 9, 109, "key,check,exit:0"},
		{"200001", false, 9, 109, "key,check,removeAll:/usr/local/soga/,remove:/etc/systemd/system/soga.service,remove:/usr/local/bin/soga,remove:/usr/bin/soga,exit:0"},
		{"unknown", false, 9, 109, "key,check,exit:0"},
		{"error", false, 9, 10, "key,check,key"},
		{"100003", false, 10, 11, "key,check,exit:0"},
		{"100004", false, 10, 11, "key,check,exit:0"},
		{"100007", false, 10, 11, "key,check,exit:0"},
	} {
		t.Run(fmt.Sprintf("%s/%t/%d", tc.code, tc.empty, tc.before), func(t *testing.T) {
			n := tc.before
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
					return "", errors.New("fixture")
				}
				return tc.code, nil
			}
			connectionStateResponse(&n, key, check, func(b bool) { events = append(events, fmt.Sprintf("state:%t", b)) }, func(p string) error { events = append(events, "removeAll:"+p); return errors.New("ignored") }, func(p string) error { events = append(events, "remove:"+p); return errors.New("ignored") }, func(n int) { events = append(events, fmt.Sprintf("exit:%d", n)) })
			if n != tc.after || strings.Join(events, ",") != tc.effects {
				t.Fatalf("counter=%d effects=%v, expected %d %s", n, events, tc.after, tc.effects)
			}
		})
	}
}
