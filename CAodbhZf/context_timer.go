package CAodbhZf

import (
	"context"
	"time"
)

// UXNvag9 对应 0xb439c0。定时器触发时返回 true；取消时停止定时器。
func UXNvag9(ctx context.Context, timer *time.Timer) bool {
	select {
	case <-timer.C:
		return true
	case <-ctx.Done():
		timer.Stop()
		return false
	}
}
