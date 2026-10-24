package app

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
	"recovered.local/soga/config"
	"recovered.local/soga/util/cronx"
)

// WdkJetA 是原 dig.In 的混淆嵌入名；别名保留字段名及 dig 的实际标记类型。
type WdkJetA = dig.In

// C7gpSMD8r：原 RTTI 0x364ff78，32 字节。
type C7gpSMD8r struct {
	WdkJetA
	XwSaXtyHwQR *cron.Cron
	E7fCaMMT    []cronx.YNfyrQAMx `group:"JOBS"`
}

// M0mmPxPJsa：原 RTTI 0x36c73d0，80 字节。
type M0mmPxPJsa struct {
	spsw8VktR9z  context.Context
	un5PcsX7aN53 context.CancelFunc
	bvSWEWNR     *cron.Cron
	jicpAs       []cronx.YNfyrQAMx
	kxWyUwUkwBr  []cron.EntryID
}

// 原 init.0@0x1f98cc0；0x5097860 已由 cron.init 的写入确认是 DiscardLogger。
func init() {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("load location failed:", err)
	}
	scheduler := cron.New(
		cron.WithLogger(cron.DiscardLogger),
		cron.WithLocation(location),
		cron.WithParser(cron.NewParser(cron.ParseOption(381))),
		cron.WithChain(cron.SkipIfStillRunning(cron.DiscardLogger)),
	)
	scheduler.Start()
	config.Bnm1qf(func() *cron.Cron { return scheduler })
	// 原 init.2@0x1f99840。
	config.Bnm1qf(lKahKJ0kn)
}

// lKahKJ0kn 对应 0x1f9a6e0，仅装配字段；context 留待 InitLinux 初始化。
func lKahKJ0kn(input C7gpSMD8r) *M0mmPxPJsa {
	return &M0mmPxPJsa{bvSWEWNR: input.XwSaXtyHwQR, jicpAs: input.E7fCaMMT}
}

// vVeTh4DwS 对应 0x1f9a7c0。失败时保留此前已加入的任务，不自动回滚；
// 再次调用只清空记录 ID 的切片，不移除旧任务。
func (app *M0mmPxPJsa) vVeTh4DwS() error {
	app.kxWyUwUkwBr = []cron.EntryID{}
	for _, job := range app.jicpAs {
		id, err := app.bvSWEWNR.AddJob(job.Cron(), job)
		if err != nil {
			return err
		}
		app.kxWyUwUkwBr = append(app.kxWyUwUkwBr, id)
	}
	return nil
}
