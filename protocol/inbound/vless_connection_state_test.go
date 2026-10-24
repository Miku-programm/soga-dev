package inbound

import (
	"errors"
	"fmt"
	CA "recovered.local/soga/util/rtx"
	"strings"
	"testing"
)

// 期望事件直接来自原1ea6080的受限机器迹证；另外三个原闭包独立验证同值。
// 删除/退出只由fixture记录，不访问授权服务器，也不执行系统副作用。
func TestVLESSOriginalStateTraces(t *testing.T) {
	for _, tc := range []struct {
		code          string
		empty         bool
		before, after int
		events        string
	}{
		{"0", false, 0, 0, "key,check"},
		{"0", false, 10, 0, "key,check"},
		{"0", true, 0, 0, "key,state:true,check"},
		{"100001", false, 0, 0, "key,check,state:true"},
		{"100001", false, 10, 0, "key,check,state:true"},
		{"100001", true, 0, 0, "key,state:true,check,state:true"},
		{"100002", false, 0, 1, "key,check,Exit:0"},
		{"100002", false, 10, 11, "key,check,Exit:0"},
		{"100002", true, 0, 1, "key,state:true,check,Exit:0"},
		{"100003", false, 0, 1, "key,check"},
		{"100003", false, 10, 11, "key,check,Exit:0"},
		{"100003", true, 0, 1, "key,state:true,check"},
		{"100004", false, 0, 1, "key,check"},
		{"100004", false, 10, 11, "key,check,Exit:0"},
		{"100004", true, 0, 1, "key,state:true,check"},
		{"100005", false, 0, 1, "key,check,Exit:0"},
		{"100005", false, 10, 11, "key,check,Exit:0"},
		{"100005", true, 0, 1, "key,state:true,check,Exit:0"},
		{"100006", false, 0, 100, "key,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"100006", false, 10, 110, "key,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"100006", true, 0, 100, "key,state:true,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"100007", false, 0, 1, "key,check"},
		{"100007", false, 10, 11, "key,check,Exit:0"},
		{"100007", true, 0, 1, "key,state:true,check"},
		{"100008", false, 0, 100, "key,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"100008", false, 10, 110, "key,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"100008", true, 0, 100, "key,state:true,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"200001", false, 0, 100, "key,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"200001", false, 10, 110, "key,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"200001", true, 0, 100, "key,state:true,check,RemoveAll:/usr/local/soga/,Remove:/etc/systemd/system/soga.service,Remove:/usr/local/bin/soga,Remove:/usr/bin/soga,Exit:0"},
		{"unknown", false, 0, 100, "key,check,Exit:0"},
		{"unknown", false, 10, 110, "key,check,Exit:0"},
		{"unknown", true, 0, 100, "key,state:true,check,Exit:0"},
		{"error", false, 0, 1, "key,check,key"},
		{"error", false, 10, 11, "key,check,key"},
		{"error", true, 0, 0, "key,state:true,check,key"},
	} {
		n := tc.before
		var events []string
		key := func() string {
			events = append(events, "key")
			if tc.empty {
				return ""
			}
			return "fixture"
		}
		check := func() (string, error) {
			events = append(events, "check")
			if tc.code == "error" {
				return "", errors.New("fixture")
			}
			return tc.code, nil
		}
		connectionStateResponse(&n, key, check, func(x bool) { events = append(events, fmt.Sprint("state:", x)) }, func(p string) error { events = append(events, "RemoveAll:"+p); return nil }, func(p string) error { events = append(events, "Remove:"+p); return nil }, func(x int) { events = append(events, fmt.Sprint("Exit:", x)) })
		if n != tc.after || strings.Join(events, ",") != tc.events {
			t.Fatalf("%+v: %d %v", tc, n, events)
		}
	}
}
func TestVLESSOriginalCurrentHourSharedCount(t *testing.T) {
	v := &KHonGaqbIy_6{JWN154qJ4: &JWN154qJ4{fUxO7UYAH: 73}}
	v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
	old := userCountAt50C6DA0
	defer func() { userCountAt50C6DA0 = old }()
	v.bj10P7X()
	if userCountAt50C6DA0 != 73 {
		t.Fatal("close-path count")
	}
	v.fUxO7UYAH = 81
	v.sbbdX3aT6E()
	if userCountAt50C6DA0 != 81 {
		t.Fatal("open-path count")
	}
}
