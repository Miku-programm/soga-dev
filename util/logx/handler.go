// 根据 0xb1b740—0xb1bee0 的 IDA 指令恢复。
package logx

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// RTTI 显示名称为 dg9pOBXC.i_AiDbrLsgcV；方法所在包为 muc279Mpa6S4。
type i_AiDbrLsgcV struct {
	bmcty2     *zap.Logger
	xOqyEML2mC []zapcore.Field
	t0TglgfHe  string
}

// 原函数内联的级别映射。此辅助函数名由分析者指定。
func zapLevel(level slog.Level) zapcore.Level {
	if level <= slog.LevelDebug {
		return zapcore.DebugLevel
	}
	if level < slog.LevelWarn {
		return zapcore.InfoLevel
	}
	if level < slog.LevelError {
		return zapcore.WarnLevel
	}
	return zapcore.ErrorLevel
}

func (h *i_AiDbrLsgcV) Enabled(_ context.Context, level slog.Level) bool {
	return h != nil && h.bmcty2 != nil && h.bmcty2.Core().Enabled(zapLevel(level))
}

func (h *i_AiDbrLsgcV) Handle(_ context.Context, record slog.Record) error {
	if h == nil || h.bmcty2 == nil {
		return nil
	}
	entry := h.bmcty2.Check(zapLevel(record.Level), record.Message)
	if entry == nil {
		return nil
	}
	if record.NumAttrs() == 0 {
		entry.Write(h.xOqyEML2mC...)
		return nil
	}
	fields := make([]zapcore.Field, 0, len(h.xOqyEML2mC)+record.NumAttrs())
	fields = append(fields, h.xOqyEML2mC...)
	record.Attrs(func(attr slog.Attr) bool {
		fields = wFqiIG4V(fields, attr, h.t0TglgfHe)
		return true
	})
	entry.Write(fields...)
	return nil
}

func (h *i_AiDbrLsgcV) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	next := &i_AiDbrLsgcV{
		bmcty2:     h.bmcty2,
		xOqyEML2mC: make([]zapcore.Field, len(h.xOqyEML2mC), len(h.xOqyEML2mC)+len(attrs)),
		t0TglgfHe:  h.t0TglgfHe,
	}
	copy(next.xOqyEML2mC, h.xOqyEML2mC)
	for _, attr := range attrs {
		next.xOqyEML2mC = wFqiIG4V(next.xOqyEML2mC, attr, h.t0TglgfHe)
	}
	return next
}

func (h *i_AiDbrLsgcV) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	return &i_AiDbrLsgcV{bmcty2: h.bmcty2, xOqyEML2mC: h.xOqyEML2mC, t0TglgfHe: h.t0TglgfHe + name + "."}
}
