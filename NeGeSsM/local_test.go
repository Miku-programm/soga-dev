package NeGeSsM

import (
	"fmt"
	"os"
	"path/filepath"
	C "recovered.local/soga/cNT24rG4F"
	"strings"
	"sync"
	"testing"
	"time"
)

// 使用实际配置文件加载入口，不向生产包加入可替换的配置钩子。
func loadAccessFixture(t *testing.T, dir string, enabled bool, extra string) {
	t.Helper()
	conf := filepath.Join(t.TempDir(), "test.conf")
	text := fmt.Sprintf("access_log_enable=%t\naccess_log_dir=%s\nlog_file=false\nip_cache_file=false\n%s", enabled, dir, extra)
	if err := os.WriteFile(conf, []byte(text), 0600); err != nil {
		t.Fatal(err)
	}
	C.MkwNWciPHS(conf)
	if err := C.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
}

func TestLocalWriterHourlyRotationAndDisabledDirectory(t *testing.T) {
	oldPath := C.ConfigurationPath()
	defer C.MkwNWciPHS(oldPath)
	dir := t.TempDir()
	loadAccessFixture(t, dir, false, "")
	s := &cToP4Y8{}
	defer s.t_DddcBjF7()
	r := cUwNbDJ6opWy{time.Date(2026, 10, 23, 22, 59, 59, 0, time.UTC), "tcp", 7, 42, "a|b\nc", "x\r\ny", "::1", 443}
	var b strings.Builder
	s.ajkzPuuQ(&b, &r)
	first := s.uzBEsO7m
	if first == nil || s.wBou6BBU0y2S.Size() != 65536 {
		t.Fatal("writer initialization")
	}
	path := filepath.Join(dir, "soga-access-2026-10-23 22:00:00.log")
	data, err := os.ReadFile(path)
	if err != nil || len(data) != 0 {
		t.Fatalf("unexpected pre-flush bytes: %q %v", data, err)
	}
	s.eyAgzSwv7d()
	data, err = os.ReadFile(path)
	if err != nil || string(data) != "2026-10-23 22:59:59|tcp|7|42|a_bc|xy|::1|443\n" {
		t.Fatalf("line: %q %v", data, err)
	}
	s.ajkzPuuQ(&b, &r)
	if s.uzBEsO7m != first {
		t.Fatal("same hour reopened")
	}
	r.pv0C8wVKLJh = r.pv0C8wVKLJh.Add(time.Second)
	s.ajkzPuuQ(&b, &r)
	if s.uzBEsO7m == first {
		t.Fatal("hour did not rotate")
	}
	if _, err := first.Write(nil); err == nil {
		t.Fatal("old handle not closed")
	}
	data, _ = os.ReadFile(path)
	if strings.Count(string(data), "\n") != 2 {
		t.Fatal("rotation omitted buffered data")
	}
	loadAccessFixture(t, "false", false, "")
	s.ajkzPuuQ(&b, &r)
	if *s != (cToP4Y8{}) {
		t.Fatal("disabled directory not fully reset")
	}
}

func TestOriginalQueueDropAndConfigurationGate(t *testing.T) {
	oldPath := C.ConfigurationPath()
	defer C.MkwNWciPHS(oldPath)
	loadAccessFixture(t, t.TempDir(), true, "")
	// 固定已启动的 once 和独立缓冲队列，隔离后台无限循环，仅验证原提交入口。
	oldQueue := queueAt5096948
	queueAt5096948 = make(chan cUwNbDJ6opWy, 8192)
	defer func() { queueAt5096948 = oldQueue; onceAt50C7D00 = sync.Once{}; droppedAt50C79D8.Store(0) }()
	onceAt50C7D00.Do(func() {})
	before := time.Now()
	IPblyWsS("udp", 12, 34, "source", "host", 53, "destination")
	r := <-queueAt5096948
	if r.uUbcC0nMK1Z9 != "udp" || r.iJYI5hR != 12 || r.zBzHxRx4qP0 != 34 || r.c7a4q7l7VIYD != "source" || r.mv3eQCXbV != "host" || r.vz3eAFzt != "destination" || r.g8o8V93EQtw != 53 || r.pv0C8wVKLJh.Before(before) {
		t.Fatalf("ABI fields: %+v", r)
	}
	for i := 0; i < 8193; i++ {
		IPblyWsS("tcp", 1, 2, "a", "b", 3, "c")
	}
	if len(queueAt5096948) != 8192 || droppedAt50C79D8.Load() != 1 {
		t.Fatal("nonblocking overflow")
	}
	loadAccessFixture(t, "false", false, "")
	IPblyWsS("tcp", 1, 2, "a", "b", 3, "c")
	if droppedAt50C79D8.Load() != 1 {
		t.Fatal("disabled entry reached queue")
	}
}
