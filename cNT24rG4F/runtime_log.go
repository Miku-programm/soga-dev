package cNT24rG4F

import (
	"fmt"
	"log/slog"
	"os"
	"sync/atomic"
	"time"

	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/muc279Mpa6S4"
)

var logPathAt5095AC0 atomic.Pointer[string]

// _RTOfcJE0BpE 对应 0x1327360，使用本地日期和规范化的 server_type。
func _RTOfcJE0BpE() string {
	if configAt509B6A0.HSBWD7 == "" {
		return ""
	}
	return fmt.Sprintf("%v/soga-%v-%v.log", configAt509B6A0.HSBWD7, DnhHoT2aRvs(), time.Now().Format("2006-01-02"))
}

// W2uyDtOBN2 对应 0x1327040。原比较键只有日志路径；路径未变时，
// 即使 log_level 已更改也不会重建日志器。
func W2uyDtOBN2() {
	path := _RTOfcJE0BpE()
	if old := logPathAt5095AC0.Load(); old != nil && *old == path {
		return
	}
	if path != "" && !CAodbhZf.VSCJLy2No(path) {
		dir := CAodbhZf.EVZvjBCf(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			muc279Mpa6S4.V8GN5Stk(8, "make log dir <%v> failed: %v", dir, err)
			path = ""
		}
	}
	logPathAt5095AC0.Store(&path)
	level := slog.LevelInfo
	switch configAt509B6A0.NaavbD4r35 {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	case "info":
		level = slog.LevelInfo
	}
	muc279Mpa6S4.LQ3CVLXgCaCF(level, path)
	if path != "" {
		muc279Mpa6S4.V8GN5Stk(0, "init log file: %v", path)
	}
}
