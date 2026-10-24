package muc279Mpa6S4

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unsafe"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestHandlerLayoutAndLevels(t *testing.T) {
	h := i_AiDbrLsgcV{}
	s := bVYSWZ_b_pm1{}
	if unsafe.Sizeof(h) != 48 || unsafe.Offsetof(h.xOqyEML2mC) != 8 || unsafe.Offsetof(h.t0TglgfHe) != 32 || unsafe.Sizeof(s) != 32 || unsafe.Offsetof(s.le0DUsW) != 16 || unsafe.Offsetof(s.xCgQb_dljfc) != 24 {
		t.Fatal("RTTI layout mismatch")
	}
	for _, tc := range []struct {
		in  slog.Level
		out zapcore.Level
	}{{-100, -1}, {-4, -1}, {-3, 0}, {0, 0}, {3, 0}, {4, 1}, {7, 1}, {8, 2}, {100, 2}} {
		if got := zapLevel(tc.in); got != tc.out {
			t.Errorf("level %d: %d", tc.in, got)
		}
	}
	ctx := context.Background()
	var nilHandler *i_AiDbrLsgcV
	if nilHandler.Enabled(ctx, 0) || h.Enabled(ctx, 0) {
		t.Fatal("nil handler/logger enabled")
	}
	if nilHandler.Handle(ctx, slog.Record{}) != nil || h.Handle(ctx, slog.Record{}) != nil {
		t.Fatal("nil Handle error")
	}
	core, _ := observer.New(zapcore.WarnLevel)
	h.bmcty2 = zap.New(core)
	if h.Enabled(ctx, 3) || !h.Enabled(ctx, 4) {
		t.Fatal("level mapping differs")
	}
}

func TestHandlerGroupsAndRecordFields(t *testing.T) {
	core, logs := observer.New(zapcore.DebugLevel)
	h := &i_AiDbrLsgcV{bmcty2: zap.New(core)}
	if h.WithAttrs(nil) != h || h.WithGroup("") != h {
		t.Fatal("empty modifiers must retain receiver")
	}
	parent := h.WithAttrs([]slog.Attr{slog.Int("outer", 1)}).(*i_AiDbrLsgcV)
	child := parent.WithGroup("group").WithAttrs([]slog.Attr{slog.String("child", "yes")}).(*i_AiDbrLsgcV)
	if len(parent.xOqyEML2mC) != 1 || len(child.xOqyEML2mC) != 2 {
		t.Fatal("WithAttrs mutated original")
	}
	record := slog.NewRecord(time.Unix(1, 0), slog.LevelWarn, "message", 0)
	record.AddAttrs(slog.Group("nested", slog.Int("value", 7)))
	// 覆盖 Record 前五个内联属性及后续切片两个路径。
	for i := 0; i < 7; i++ {
		record.AddAttrs(slog.Int(fmt.Sprintf("n%d", i), i))
	}
	if err := child.Handle(context.Background(), record); err != nil {
		t.Fatal(err)
	}
	entries := logs.All()
	if len(entries) != 1 || entries[0].Level != zapcore.WarnLevel || entries[0].Message != "message" {
		t.Fatalf("entry: %#v", entries)
	}
	fields := entries[0].ContextMap()
	if fields["outer"] != int64(1) || fields["group.child"] != "yes" || fields["group.nested.value"] != int64(7) || fields["group.n6"] != int64(6) {
		t.Fatalf("fields: %#v", fields)
	}
	if entries[0].Time.Equal(record.Time) {
		t.Fatal("record timestamp must not replace zap-generated timestamp")
	}
	if err := parent.Handle(context.Background(), slog.NewRecord(time.Time{}, slog.LevelInfo, "no attrs", 0)); err != nil {
		t.Fatal(err)
	}
	if len(logs.All()[1].Context) != 1 {
		t.Fatal("stored fields lost for attribute-free record")
	}
}

type countedStringer struct{ calls *int }

func (s countedStringer) String() string { *s.calls++; return "converted" }

func TestLogWrappersFormatOnlyEnabledMessages(t *testing.T) {
	previous := loggerAt5095C80.Load()
	t.Cleanup(func() { loggerAt5095C80.Store(previous) })
	loggerAt5095C80.Store(nil)
	TRTbrquD(0, "ignored")
	V8GN5Stk(0, "ignored")
	core, logs := observer.New(zapcore.InfoLevel)
	logger := zap.New(core)
	loggerAt5095C80.Store(&bVYSWZ_b_pm1{wNpsYA6: slog.New(&i_AiDbrLsgcV{bmcty2: logger}), hRJPE3zodxa: logger})
	calls := 0
	JucPcHhJ("%s", countedStringer{&calls})
	if calls != 0 {
		t.Fatal("disabled log must not format")
	}
	TRTbrquD(0, "one", "two\n")
	VDyYELbQKX("number=%d", 4)
	Jxy_B4("warn")
	GigMlIhMM("%s", "warn formatted")
	FcL3BVMFuVAn("error")
	got := logs.All()
	if len(got) != 5 || got[0].Message != "one two\n" || got[1].Message != "number=4" || got[2].Level != zapcore.WarnLevel || got[4].Level != zapcore.ErrorLevel {
		t.Fatalf("logs: %#v", got)
	}
}

func TestInitializationAppendsClosesOldFileAndUsesConsoleFormat(t *testing.T) {
	previous := loggerAt5095C80.Load()
	t.Cleanup(func() {
		state := loggerAt5095C80.Swap(previous)
		if state != nil && state.xCgQb_dljfc != nil {
			state.xCgQb_dljfc.Close()
		}
	})
	path := filepath.Join(t.TempDir(), "log")
	if err := os.WriteFile(path, []byte("prefix\n"), 0600); err != nil {
		t.Fatal(err)
	}
	LQ3CVLXgCaCF(slog.LevelWarn, path)
	old := loggerAt5095C80.Load()
	VDyYELbQKX("hidden")
	Jxy_B4("visible")
	LQ3CVLXgCaCF(slog.LevelInfo, "")
	if _, err := old.xCgQb_dljfc.Write([]byte("closed")); err == nil {
		t.Fatal("old file remained open")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) != 3 || lines[0] != "prefix" {
		t.Fatalf("file: %q", data)
	}
	parts := strings.Split(lines[1], "\t")
	if len(parts) != 3 || parts[1] != "WARN" || parts[2] != "visible" {
		t.Fatalf("console output: %q", data)
	}
	if _, err := time.Parse("2006/01/02 15:04:05", parts[0]); err != nil {
		t.Fatal(err)
	}
}

func TestFatalWrappersKeepOriginalZeroExitStatus(t *testing.T) {
	if mode := os.Getenv("SOGA_RECOVERED_LOG_EXIT_CASE"); mode != "" {
		switch mode {
		case "plain":
			ROUJ82_bIQ("exit")
		case "format":
			Wa0FarWk("%s", "exit")
		case "open":
			LQ3CVLXgCaCF(0, os.Getenv("SOGA_RECOVERED_BAD_LOG_PATH"))
		}
		os.Exit(92)
	}
	for _, mode := range []string{"plain", "format", "open"} {
		cmd := exec.Command(os.Args[0], "-test.run=^TestFatalWrappersKeepOriginalZeroExitStatus$")
		cmd.Env = append(os.Environ(), "SOGA_RECOVERED_LOG_EXIT_CASE="+mode, "SOGA_RECOVERED_BAD_LOG_PATH="+filepath.Join(t.TempDir(), "missing", "log"))
		if data, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%s: %s %v", mode, data, err)
		}
	}
}
