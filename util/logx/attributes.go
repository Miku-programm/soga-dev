package logx

import (
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// jh9bDVK_MFF 对应地址 0xb1d120。原函数不捕获 LogValue 的 panic，
// 也没有 slog.Value.Resolve 的迭代上限，因此直接保留原循环。
func jh9bDVK_MFF(value slog.Value) slog.Value {
	for value.Kind() == slog.KindLogValuer {
		value = value.Any().(slog.LogValuer).LogValue()
	}
	return value
}

// wFqiIG4V 对应地址 0xb1bee0，将嵌套 slog 属性展开成带点号前缀的 zap 字段。
func wFqiIG4V(fields []zapcore.Field, attr slog.Attr, prefix string) []zapcore.Field {
	attr.Value = jh9bDVK_MFF(attr.Value)
	if attr.Equal(slog.Attr{}) {
		return fields
	}
	if attr.Value.Kind() == slog.KindGroup {
		children := attr.Value.Group()
		if len(children) == 0 {
			return fields
		}
		if attr.Key != "" {
			prefix += attr.Key + "."
		}
		for _, child := range children {
			fields = wFqiIG4V(fields, child, prefix)
		}
		return fields
	}
	key := attr.Key
	if prefix != "" {
		key = prefix + key
	}
	var field zapcore.Field
	switch attr.Value.Kind() {
	case slog.KindBool:
		field = zap.Bool(key, attr.Value.Bool())
	case slog.KindDuration:
		field = zap.Duration(key, attr.Value.Duration())
	case slog.KindFloat64:
		field = zap.Float64(key, attr.Value.Float64())
	case slog.KindInt64:
		field = zap.Int64(key, attr.Value.Int64())
	case slog.KindString:
		field = zap.String(key, attr.Value.String())
	case slog.KindTime:
		field = zap.Time(key, attr.Value.Time())
	case slog.KindUint64:
		field = zap.Uint64(key, attr.Value.Uint64())
	default:
		field = zap.Any(key, attr.Value.Any())
	}
	return append(fields, field)
}
