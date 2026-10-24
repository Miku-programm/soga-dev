package rtx

import (
	"context"
	"time"
)

// SwUrlM3M: 0xb43a80；闭包 0xb4eb40 和 0xb4ec60。
// 首次调用等到 ticker 到期，回调串行运行，取消 context 后停止 ticker。
func SwUrlM3M(ctx context.Context, fn func(), interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fn()
			case <-ctx.Done():
				return
			}
		}
	}()
}
