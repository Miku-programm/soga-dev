package app

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/modern-go/reflect2"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/protocol/inbound"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/panel/registry"
)

// starterSubmitLoop 是原 avZC_rQmNC.func4@0x1ffa640 的可命名提取。
// 原始算术控制流及成功、失败、空入站、取消四条路径分别由原机码确认。
// 每轮先等待再上报；读取当前入站及间隔，支持后续替换。
func starterSubmitLoop(ctx context.Context, panel registry.VyFk5uT, get func() inbound.Kz_8ajU7, interval *atomic.Int64) {
	for {
		delay := config.EVzHyyNDe(time.Duration(interval.Load()))
		logx.JucPcHhJ("wait %.2fs for next submit", delay.Seconds())
		if !rtx.UXNvag9(ctx, time.NewTimer(delay)) {
			return
		}
		inbound := get()
		if reflect2.IsNil(inbound) {
			continue
		}
		data := inbound.GetSubmitData()
		if err := panel.SubmitData(data); err != nil {
			logx.GigMlIhMM("submit data failed, alive user: %v, error: %v", len(data.JNcHHwTZ7), err)
		} else {
			logx.VDyYELbQKX("submit data success, alive user: %v", len(data.JNcHHwTZ7))
		}
	}
}
