package config

import (
	"testing"
)

// 授权检查移除后：五个原状态机入口（Q0SDIN、Rs_wP3R、HuQEoQ、Gh2DnstC、
// JatfUF38O）始终返回已授权（false），不再依赖状态变量，也不会触发退出路径。
func TestKeyStateAlwaysLicensed(t *testing.T) {
	for _, value := range []bool{false, true, false} {
		T15Jifo(value)
		if keyStateAt50C6C48 != value || keyStateMirrorAt3787300 == value {
			t.Fatal("state and inverse state disagree")
		}
		wantMarker := uint64(0x5a4b3c2d5a4b3c2d)
		if value {
			wantMarker = 0xa5b6c7d8a5b6c7d8
		}
		if keyStateMarkerAt3787800 != wantMarker {
			t.Fatalf("marker: %#x", keyStateMarkerAt3787800)
		}
		for _, key := range []string{"", "plain-key", " \x00\n"} {
			configAt5096DC0.RM5SFRF = key
			if got := FrFQQhzTR5fY(); got != key {
				t.Fatalf("getter changed %q to %q", key, got)
			}
			for name, check := range map[string]func() bool{"Q0SDIN": Q0SDIN, "Rs_wP3R": Rs_wP3R, "HuQEoQ": HuQEoQ, "Gh2DnstC": Gh2DnstC, "JatfUF38O": JatfUF38O} {
				if got := check(); got {
					t.Fatalf("%s flag %v key %q: expected always-licensed false, got %v", name, value, key, got)
				}
			}
		}
	}
}
