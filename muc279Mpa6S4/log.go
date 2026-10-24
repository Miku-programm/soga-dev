// 根据 0xb1b1e0—0xb1b740 恢复的日志入口。
package muc279Mpa6S4

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"
)

type bVYSWZ_b_pm1 struct {
	wNpsYA6     *slog.Logger
	hRJPE3zodxa *zap.Logger
	le0DUsW     zap.AtomicLevel
	xCgQb_dljfc *os.File
}

// 全局变量原地址分别为 0x5095c80、0x50c7158；变量名由分析者指定。
var loggerAt5095C80 atomic.Pointer[bVYSWZ_b_pm1]
var loggerMutexAt50C7158 sync.Mutex

// TRTbrquD 对应原 tRTbrquD；首字母改大写以保留 Go 跨包调用的可编译性。
func TRTbrquD(level slog.Level, args ...any) {
	state := loggerAt5095C80.Load()
	if state == nil {
		return
	}
	ctx := context.Background()
	if !state.wNpsYA6.Enabled(ctx, level) {
		return
	}
	message := fmt.Sprintln(args...)
	if len(message) != 0 {
		message = message[:len(message)-1]
	}
	state.wNpsYA6.Log(ctx, level, message)
}

// V8GN5Stk 对应原 v8GN5Stk，使用相同的导出名称映射。
func V8GN5Stk(level slog.Level, format string, args ...any) {
	state := loggerAt5095C80.Load()
	if state == nil {
		return
	}
	ctx := context.Background()
	if !state.wNpsYA6.Enabled(ctx, level) {
		return
	}
	state.wNpsYA6.Log(ctx, level, fmt.Sprintf(format, args...))
}

func JucPcHhJ(format string, args ...any)   { V8GN5Stk(slog.LevelDebug, format, args...) }
func VDyYELbQKX(format string, args ...any) { V8GN5Stk(slog.LevelInfo, format, args...) }
func Jxy_B4(args ...any)                    { TRTbrquD(slog.LevelWarn, args...) }
func GigMlIhMM(format string, args ...any)  { V8GN5Stk(slog.LevelWarn, format, args...) }
func FcL3BVMFuVAn(args ...any)              { TRTbrquD(slog.LevelError, args...) }
func ROUJ82_bIQ(args ...any)                { TRTbrquD(slog.LevelError, args...); os.Exit(0) }
func Wa0FarWk(format string, args ...any)   { V8GN5Stk(slog.LevelError, format, args...); os.Exit(0) }
