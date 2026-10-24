package errors

import (
	"context"
	Log "recovered.local/soga/util/logx"
	"strings"
)

// LogInfoInner 0x124faa0。样本不读取传入 context。
func LogInfoInner(_ context.Context, err error, values ...string) {
	Log.V8GN5Stk(0, "%v: %v", strings.Join(values, " "), err)
}
