package config

import (
	"testing"
)

// 授权检查移除后：Gh2DnstC（0x1342300）与 JatfUF38O（0x1342900）
// 始终返回已授权（false），不再读取状态变量，也没有退出路径。
func TestPanelStateAlwaysLicensed(t *testing.T) {
	for _, value := range []bool{false, true} {
		for _, marker := range []uint64{0x0, 0xa5b6c7d8a5b6c7d8, 0x5a4b3c2d5a4b3c2d} {
			for _, key := range []string{"", "x"} {
				keyStateAt50C6C48, keyStateMirrorAt3787300, keyStateMarkerAt3787800 = value, !value, marker
				configAt5096DC0.RM5SFRF = key
				if Gh2DnstC() {
					t.Fatalf("Gh2DnstC flag %v marker %#x key %q: expected always-licensed false", value, marker, key)
				}
				if JatfUF38O() {
					t.Fatalf("JatfUF38O flag %v marker %#x key %q: expected always-licensed false", value, marker, key)
				}
			}
		}
	}
}
