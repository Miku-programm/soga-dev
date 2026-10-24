package CbO_7wp3B

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/modern-go/reflect2"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/MFFa9yT"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/q5bqCgTdxaTM"
)

// starterSubmitLoop 是原 avZC_rQmNC.func4@0x1ffa640 的可命名提取。
// 原始算术控制流及成功、失败、空入站、取消四条路径分别由原机码确认。
// 每轮先等待再上报；读取当前入站及间隔，支持后续替换。
func starterSubmitLoop(ctx context.Context, panel q5bqCgTdxaTM.VyFk5uT, get func() MFFa9yT.Kz_8ajU7, interval *atomic.Int64) {
	for {
		delay := cNT24rG4F.EVzHyyNDe(time.Duration(interval.Load()))
		muc279Mpa6S4.JucPcHhJ("wait %.2fs for next submit", delay.Seconds())
		if !CAodbhZf.UXNvag9(ctx, time.NewTimer(delay)) {
			return
		}
		inbound := get()
		if reflect2.IsNil(inbound) {
			continue
		}
		data := inbound.GetSubmitData()
		if err := panel.SubmitData(data); err != nil {
			muc279Mpa6S4.GigMlIhMM("submit data failed, alive user: %v, error: %v", len(data.JNcHHwTZ7), err)
		} else {
			muc279Mpa6S4.VDyYELbQKX("submit data success, alive user: %v", len(data.JNcHHwTZ7))
		}
	}
}
