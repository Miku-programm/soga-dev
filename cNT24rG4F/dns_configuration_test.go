package cNT24rG4F

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDNSConfigurationLayoutsAndConversion(t *testing.T) {
	rt := reflect.TypeFor[AorTZHf4G9Yq]()
	if rt.Size() != 40 || rt.Field(0).Offset != 0 || rt.Field(1).Offset != 16 || rt.Field(0).Tag.Get("yaml") != "strategy" || rt.Field(1).Tag.Get("yaml") != "rules" {
		t.Fatal(rt)
	}
	rule := &AorTZHf4G9Yq{ZkuUXsveafC: "ipv6_first", IxW5Yxj: []string{"udp://127.0.0.1:53"}}
	got := D2YhfBI(map[string]*AorTZHf4G9Yq{" , a.test, ,b.test,a.test , ": rule, " , ": nil})
	if len(got) != 1 || !reflect.DeepEqual(got[0].K8BNfziIgb, []string{"a.test", "b.test", "a.test"}) || got[0].CM0eMsqyP != "ipv6_first" {
		t.Fatal(got)
	}
	rule.IxW5Yxj[0] = "changed"
	if got[0].CZ80ik[0] != "changed" {
		t.Fatal("server slice unexpectedly copied")
	}
	if rule.String() != "{ strategy: ipv6_first, rules: [changed] }" {
		t.Fatal(rule.String())
	}
	if got := D2YhfBI(nil); got == nil || len(got) != 0 {
		t.Fatal("empty conversion must be allocated empty slice")
	}
}

func TestDNSYAMLFormatsAndAtomicFailure(t *testing.T) {
	old := dnsRulesAt5095AD0.Load()
	t.Cleanup(func() { dnsRulesAt5095AD0.Store(old) })
	if err := as3gRNOKx9KD("test", []byte("'a.test, b.test':\n  strategy: ipv6_only\n  rules: [1.1.1.1, 'tls://example.test']\n")); err != nil {
		t.Fatal(err)
	}
	got := *dnsRulesAt5095AD0.Load()
	if len(got) != 1 || got[0].CM0eMsqyP != "ipv6_only" || !reflect.DeepEqual(got[0].CZ80ik, []string{"1.1.1.1", "tls://example.test"}) {
		t.Fatal(got)
	}
	previous := dnsRulesAt5095AD0.Load()
	for _, data := range []string{"[:", "a: {strategy: []}", "a: []\na: []"} {
		err := as3gRNOKx9KD("invalid", []byte(data))
		if err == nil || errors.Unwrap(err) != nil || dnsRulesAt5095AD0.Load() != previous {
			t.Fatalf("failed parse changed snapshot: %q, %v", data, err)
		}
	}
	if err := as3gRNOKx9KD("legacy", []byte("'a, b': [1.1.1.1, 8.8.8.8]")); err != nil {
		t.Fatal(err)
	}
	got = *dnsRulesAt5095AD0.Load()
	if len(got) != 1 || got[0].CM0eMsqyP != "ipv4_first" || len(got[0].CZ80ik) != 2 {
		t.Fatal(got)
	}
	if err := as3gRNOKx9KD("empty", nil); err != nil {
		t.Fatal(err)
	}
	if got := dnsRulesAt5095AD0.Load(); got == nil || *got == nil || len(*got) != 0 {
		t.Fatal("empty YAML did not publish empty snapshot")
	}
}

func TestDNSFileReadAndParseErrorsPreserveSnapshot(t *testing.T) {
	old, oldPath := dnsRulesAt5095AD0.Load(), dnsPathAt5096DE0
	t.Cleanup(func() { dnsRulesAt5095AD0.Store(old); dnsPathAt5096DE0 = oldPath })
	dnsPathAt5096DE0 = filepath.Join(t.TempDir(), "dns.yml")
	if err := os.WriteFile(dnsPathAt5096DE0, []byte("example.test: {rules: [127.0.0.1]}"), 0600); err != nil {
		t.Fatal(err)
	}
	ka8yZsm4z()
	previous := dnsRulesAt5095AD0.Load()
	if previous == nil || len(*previous) != 1 {
		t.Fatal("local file not loaded")
	}
	if err := os.WriteFile(dnsPathAt5096DE0, []byte("["), 0600); err != nil {
		t.Fatal(err)
	}
	ka8yZsm4z()
	if dnsRulesAt5095AD0.Load() != previous {
		t.Fatal("invalid YAML replaced snapshot")
	}
	if err := os.Remove(dnsPathAt5096DE0); err != nil {
		t.Fatal(err)
	}
	ka8yZsm4z()
	if dnsRulesAt5095AD0.Load() != previous {
		t.Fatal("missing file replaced snapshot")
	}
}
