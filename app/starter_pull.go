package app

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/modern-go/reflect2"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/util/errutil"
	"recovered.local/soga/protocol/inbound"
	"recovered.local/soga/config"
	"recovered.local/soga/util/logx"
	"recovered.local/soga/panel/registry"
)

// starterPullLoop 提取原 avZC_rQmNC.func5@0x1ff4700。
// start 参数对应原接收者 vDpj7g 方法；get/set 对应原互斥锁保护的两个闭包。
// 首轮直接拉取，后续等待；更新返回 true 时重建。两个间隔仅接收正值。
func starterPullLoop(ctx context.Context, panel registry.VyFk5uT,
	get func() inbound.Kz_8ajU7, set func(inbound.Kz_8ajU7),
	start func(context.Context, *registry.D5gTIZ5aVg) (inbound.Kz_8ajU7, error),
	check, submit *atomic.Int64, stopAfterFunc *func() bool,
) {
	for {
		func() {
			data, err := panel.GetAPIData()
			if err != nil {
				logx.Jxy_B4(err)
				return
			}
			if data.Xdo1D9 > 0 {
				check.Store(int64(data.Xdo1D9))
			}
			if data.DDu2Nm_P > 0 {
				submit.Store(int64(data.DDu2Nm_P))
			}
			mergeStarterRules(data)
			current := get()
			if reflect2.IsNil(current) {
				created, err := start(ctx, data)
				if err != nil {
					logx.Jxy_B4("new server err:", err)
					return
				}
				*stopAfterFunc = context.AfterFunc(ctx, func() { created.Stop() })
				set(created)
				return
			}
			if !current.UpdateAPIData(data) {
				return
			}
			(*stopAfterFunc)()
			current.Stop()
			set(nil)
			created, err := start(ctx, data)
			if err != nil {
				logx.Jxy_B4("new server err:", err)
				return
			}
			*stopAfterFunc = context.AfterFunc(ctx, func() { created.Stop() })
			set(created)
		}()
		delay := config.EVzHyyNDe(time.Duration(check.Load()))
		timer := time.NewTimer(delay)
		logx.JucPcHhJ("wait %.2fs for next check", delay.Seconds())
		if !rtx.UXNvag9(ctx, timer) {
			return
		}
	}
}

// 原 0x200c9c0 与 0x200cb40：两个字段分别并行合并，错误只记录日志；
// 失败不会覆盖面板规则，也不改写面板自身的规则错误字段。
func mergeStarterRules(data *registry.D5gTIZ5aVg) {
	_ = errutil.D5aDHlQDLp(func() error {
		rules, err := config.Klz7sa_01()
		if err != nil {
			logx.GigMlIhMM("get block list err: %v", err)
		} else {
			data.ATGYGi0 = data.ATGYGi0.Merge(rules)
		}
		return nil
	}, func() error {
		rules, err := config.Cf33JCugNk1()
		if err != nil {
			logx.GigMlIhMM("get white list err: %v", err)
		} else {
			data.NjjWbjH = data.NjjWbjH.Merge(rules)
		}
		return nil
	})
}
