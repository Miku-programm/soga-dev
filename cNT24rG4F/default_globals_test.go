package cNT24rG4F

import (
	"reflect"
	"testing"
)

func TestDefaultLogFilenamePatterns(t *testing.T) {
	for _, name := range []string{
		"soga-access-2026-10-23 09:00:00.log",
		"soga-access-0000-99-99 99:00:00.log",
	} {
		match := accessLogPatternAt5095A98.FindStringSubmatch(name)
		if len(match) != 2 || len(match[1]) != 13 {
			t.Fatalf("access filename %q: %q", name, match)
		}
	}
	for _, name := range []string{
		"/etc/soga/soga-access-2026-10-23 09:00:00.log",
		"soga-access-2026-10-23 09:00:00.log.old",
		"soga-access-2026-10-23 09:01:00.log",
		"soga-access-2026-10-23 09:00:00Xlog",
	} {
		if accessLogPatternAt5095A98.MatchString(name) {
			t.Fatalf("unexpected access filename match: %q", name)
		}
	}
	// 普通日志的原表达式没有起止锚点；\w 仅匹配 ASCII 单词字符。
	match := logPatternAt5095AB8.FindStringSubmatch("prefix/soga-anyTLS_2-2026-10-23.log.old")
	if !reflect.DeepEqual(match, []string{"soga-anyTLS_2-2026-10-23.log", "anyTLS_2", "2026-10-23"}) {
		t.Fatalf("ordinary log captures: %q", match)
	}
	for _, name := range []string{"soga-节点-2026-10-23.log", "soga-any-tls-2026-10-23.log"} {
		if logPatternAt5095AB8.MatchString(name) {
			t.Fatalf("unexpected ordinary filename match: %q", name)
		}
	}
}

func TestDefaultDependencyContainer(t *testing.T) {
	// 使用仅在此测试定义的类型，验证默认容器保留按需构造和实例缓存。
	type dependency struct{ value int }
	type consumer struct{ dep *dependency }
	scope := containerAt5095AC8.Scope(t.Name())
	created := 0
	if err := scope.Provide(func() *dependency {
		created++
		return &dependency{value: 27}
	}); err != nil {
		t.Fatal(err)
	}
	if err := scope.Provide(func(dep *dependency) *consumer { return &consumer{dep: dep} }); err != nil {
		t.Fatal(err)
	}
	if created != 0 {
		t.Fatal("constructor executed before Invoke")
	}
	for i := 0; i < 2; i++ {
		if err := scope.Invoke(func(c *consumer) {
			if c.dep.value != 27 {
				t.Fatalf("dependency value: %d", c.dep.value)
			}
		}); err != nil {
			t.Fatal(err)
		}
	}
	if created != 1 {
		t.Fatalf("dependency constructed %d times", created)
	}
}
