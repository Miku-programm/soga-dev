package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"recovered.local/soga/audit/matcher"
)

func TestRuleFileSnapshots(t *testing.T) {
	oldBlockPath, oldWhitePath := blockListPathAt5096DA0, whiteListPathAt5096DF0
	oldBlockURL, oldWhiteURL := configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi
	oldPorts := configAt50A2900.DUpePpa39q
	oldBlock, oldWhite := blockListAt5095A90.Load(), whiteListAt5095AD8.Load()
	t.Cleanup(func() {
		blockListPathAt5096DA0, whiteListPathAt5096DF0 = oldBlockPath, oldWhitePath
		configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi = oldBlockURL, oldWhiteURL
		configAt50A2900.DUpePpa39q = oldPorts
		blockListAt5095A90.Store(oldBlock)
		whiteListAt5095AD8.Store(oldWhite)
	})
	path := filepath.Join(t.TempDir(), "rules")
	content := "\n # comment\r\n  domain:example.test # trailing\r\nport:80\nport:80\n\u2003ip:192.0.2.1\u2003\n" + strings.Repeat("a", 4100) + "\nlast-no-newline"
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	blockListPathAt5096DA0, whiteListPathAt5096DF0 = path, path
	configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi = "", ""
	configAt50A2900.DUpePpa39q = []string{"443", "100-200", ""}
	LnLDMWbjk3()
	I5PWiaSBwdar()
	wantWhite := matcher.VHWq49iRat{"domain:example.test", "port:80", "port:80", "ip:192.0.2.1", strings.Repeat("a", 4096), "aaaa", "last-no-newline"}
	wantBlock := append(append(matcher.VHWq49iRat{}, wantWhite...), "port:443", "port:100-200", "port:")
	if !reflect.DeepEqual(*whiteListAt5095AD8.Load(), wantWhite) || !reflect.DeepEqual(*blockListAt5095A90.Load(), wantBlock) {
		t.Fatal("comment trimming, duplicate preservation, line fragments or forbidden ports mismatch")
	}
	previousBlock, previousWhite := blockListAt5095A90.Load(), whiteListAt5095AD8.Load()
	configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi = "https://example.invalid/block", "https://example.invalid/white"
	if err := os.WriteFile(path, []byte("replacement"), 0600); err != nil {
		t.Fatal(err)
	}
	LnLDMWbjk3()
	I5PWiaSBwdar()
	if blockListAt5095A90.Load() != previousBlock || whiteListAt5095AD8.Load() != previousWhite {
		t.Fatal("remote configuration did not preserve snapshots")
	}
	configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi = "", ""
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	LnLDMWbjk3()
	I5PWiaSBwdar()
	if blockListAt5095A90.Load() != previousBlock || whiteListAt5095AD8.Load() != previousWhite {
		t.Fatal("open failure replaced old rules")
	}
	// 目录可以打开，但读取失败；原循环仍发布空规则，并追加禁止端口。
	blockListPathAt5096DA0, whiteListPathAt5096DF0 = filepath.Dir(path), filepath.Dir(path)
	LnLDMWbjk3()
	I5PWiaSBwdar()
	if len(*whiteListAt5095AD8.Load()) != 0 || !reflect.DeepEqual(*blockListAt5095A90.Load(), matcher.VHWq49iRat{"port:443", "port:100-200", "port:"}) {
		t.Fatal("read error publication mismatch")
	}
}

func TestDNSPathLiteralPrefix(t *testing.T) {
	old := dnsPathAt5096DE0
	t.Cleanup(func() { dnsPathAt5096DE0 = old })
	for _, tc := range []struct{ input, want string }{
		{"", "/etc/soga/"}, {"dns.yml", "/etc/soga/dns.yml"}, {"../dns.yml", "/etc/soga/../dns.yml"},
		{"/tmp//dns.yml", "/tmp//dns.yml"}, {`C:\dns.yml`, `/etc/soga/C:\dns.yml`},
	} {
		QJZTKBo(tc.input)
		if dnsPathAt5096DE0 != tc.want {
			t.Fatalf("path %q = %q", tc.input, dnsPathAt5096DE0)
		}
	}
}
