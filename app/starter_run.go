package app

import (
	"context"
	"sync"
	"sync/atomic"

	inb "recovered.local/soga/protocol/inbound"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/registry"
)

// runOriginalStarter 提取 WrdrARwuhuha.avZC_rQmNC@0x1fa5a20 的实际装配。
// start 对应接收者的 vDpj7g 方法；实际入口绑定真实原生八协议工厂。
// 三个 goroutine 按原顺序发起；初始空上报的错误被忽略。
func runOriginalStarter(ctx context.Context, panel registry.VyFk5uT,
	start func(context.Context, *registry.D5gTIZ5aVg) (inb.Kz_8ajU7, error),
) {
	var mu sync.Mutex
	var inbound inb.Kz_8ajU7
	get := func() inb.Kz_8ajU7 {
		mu.Lock()
		value := inbound
		mu.Unlock()
		return value
	}
	set := func(value inb.Kz_8ajU7) {
		mu.Lock()
		inbound = value
		mu.Unlock()
	}
	var stopAfterFunc func() bool
	var check, submit atomic.Int64
	checkDelay, submitDelay := config.PanelPollingIntervals()
	check.Store(int64(checkDelay))
	submit.Store(int64(submitDelay))
	go func() { _ = panel.SubmitData(new(registry.FGmNba4MhioY)) }()
	go starterSubmitLoop(ctx, panel, get, &submit)
	go starterPullLoop(ctx, panel, get, set, start, &check, &submit, &stopAfterFunc)
	<-ctx.Done()
}
