package logx

import (
	"io"
	"log/slog"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 原 init.0 @ 0xb1a880。
func init() { LQ3CVLXgCaCF(slog.LevelInfo, "") }

// LQ3CVLXgCaCF @ 0xb1a8c0。原实现使用 zap consoleEncoder 和自定义 slog.Handler。
// zap 依赖版本选用本地可用的 1.27.1；原构建的精确依赖版本尚无法由元数据确认。
func LQ3CVLXgCaCF(level slog.Level, path string) {
	var file *os.File
	var writer io.Writer = os.Stdout
	if path != "" {
		var err error
		file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			ROUJ82_bIQ(err)
		}
		writer = io.MultiWriter(os.Stdout, file)
	}
	// 静态结构 0x2169fc8 与函数值 0x3768528—0x3768550。
	encoder := zap.NewProductionEncoderConfig()
	encoder.MessageKey = "message"
	encoder.TimeKey = "time"
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		// func1 @ 0xb1de00；格式字串来自原解码函数的受限仿真。
		enc.AppendString(t.Format("2006/01/02 15:04:05"))
	}
	atomicLevel := zap.NewAtomicLevelAt(zapLevel(level))
	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.AddSync(writer), atomicLevel))
	handler := &i_AiDbrLsgcV{bmcty2: logger}
	state := &bVYSWZ_b_pm1{wNpsYA6: slog.New(handler), hRJPE3zodxa: logger, le0DUsW: atomicLevel, xCgQb_dljfc: file}
	loggerMutexAt50C7158.Lock()
	previous := loggerAt5095C80.Swap(state)
	loggerMutexAt50C7158.Unlock()
	if previous != nil && previous.xCgQb_dljfc != nil {
		_ = previous.xCgQb_dljfc.Close()
	}
}
