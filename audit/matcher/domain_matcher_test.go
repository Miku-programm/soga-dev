package matcher

import (
	"reflect"
	"sync"
	"testing"
)

func TestDomainMatcherRTTI(t *testing.T) {
	typ := reflect.TypeOf(IXZNQqqXDsV{})
	fields := []struct {
		name   string
		offset uintptr
		typ    reflect.Type
	}{
		{"_fXIaVa", 0, reflect.TypeOf(int(0))}, {"qwkN4ar_B", 8, reflect.TypeOf("")},
		{"wg8voMH", 24, reflect.TypeFor[VkyfPk_JtMAm]()}, {"d9LcIO", 40, reflect.TypeOf(sync.Once{})},
		{"iTu75U", 56, reflect.TypeFor[TaBfSSyPGYr]()},
	}
	if typ.Size() != 72 || typ.NumField() != len(fields) {
		t.Fatalf("RTTI大小/字段数量: %v", typ)
	}
	for i, want := range fields {
		f := typ.Field(i)
		if f.Name != want.name || f.Offset != want.offset || f.Type != want.typ || f.Tag != "" {
			t.Fatalf("字段 %d: %+v", i, f)
		}
	}
}
func TestDomainRulesExactBehavior(t *testing.T) {
	for _, tc := range []struct {
		rule, value string
		want        bool
	}{
		{"", "anything", false}, {"example", "xexample.org", true}, {"example", "EXAMPLE", false},
		{"domain:example.com", "example.com", true}, {"domain:example.com", "a.example.com", true}, {"domain:example.com", "notexample.com", false}, {"domain:example.com", "example.com.", false},
		{"domain:", "", true}, {"domain:", "a.", true}, {"domain:", "a", false},
		{"full:example.com", "example.com", true}, {"full:example.com", "a.example.com", false},
		{"keyword:", "anything", true}, {"keyword:ab", "cabc", true},
		{"regexp:^a+$", "aaa", true}, {"regexp:^a+$", "baaa", false}, {"regexp:[", "[", false},
		{"ip:127.0.0.1", "ip:127.0.0.1", false}, {"geoip:cn", "geoip:cn", false},
		{"port:443", "port:443", true}, {"unknown:x", "unknown:x", true}, {"DOMAIN:example.com", "example.com", false},
	} {
		t.Run(tc.rule+"/"+tc.value, func(t *testing.T) {
			r := Q2jO4SS{tc.rule}.ToDetectors()[0]
			if got := r.Test(tc.value); got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
func TestDNSAndGeneralEmptyRuleDiffer(t *testing.T) {
	dns := NewDNSMatcher("")
	if !dns.Test("anything") || dns.MatchPacket([]byte("anything")) {
		t.Fatal("DNS空规则只匹配字符串")
	}
	if (Q2jO4SS{""}).ToDetectors()[0].Test("anything") {
		t.Fatal("一般空规则应不匹配")
	}
	rules := Q2jO4SS{"full:one", "keyword:o", "full:two"}.ToDetectors()
	match, ok := rules.Test("one")
	if !ok || match != rules[0] || match.GetId() != 0 || match.GetRule() != "full:one" {
		t.Fatal("首命中/ID/规则内容错误")
	}
	if r, ok := rules.Test("absent"); ok || r != nil {
		t.Fatal("未命中返回错误")
	}
	if empty := Q2jO4SS(nil).ToDetectors(); empty == nil || len(empty) != 0 {
		t.Fatal("空输入应产生非nil空结果")
	}
}
func TestRulePrefixRecognition(t *testing.T) {
	for _, rule := range []string{"regexp:", "domain:", "full:", "geosite:", "keyword:", "ip:", "geoip:", "port:"} {
		if !WUWGex(rule) {
			t.Fatalf("未识别 %q", rule)
		}
	}
	for _, rule := range []string{"", "example.com", "DOMAIN:x", "other:x", "foo:regexp:"} {
		if WUWGex(rule) {
			t.Fatalf("错误识别 %q", rule)
		}
	}
}
func TestPacketAndConcurrentInitialization(t *testing.T) {
	rules := Q2jO4SS{"keyword:abc", "regexp:abc"}.ToDetectors()
	found, ok := rules.MatchPacket([]byte("xabc"))
	if !ok || found != rules[1] {
		t.Fatal("字节规则只应命中正则")
	}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !rules[1].Test("abc") {
				t.Error("并发匹配失败")
			}
		}()
	}
	wg.Wait()
	if r, ok := rules.MatchPacket([]byte("zzz")); ok || r != nil {
		t.Fatal("未命中返回错误")
	}
}

func TestPacketEmptyMatchUsesFindResult(t *testing.T) {
	matcher := Q2jO4SS{"regexp:^$"}.ToDetectors()[0]
	if matcher.MatchPacket(nil) {
		t.Fatal("原 Find(nil) 返回 nil，不能改成 Match(nil)")
	}
	if !matcher.MatchPacket([]byte{}) {
		t.Fatal("非 nil 空切片应保留空匹配")
	}
}
