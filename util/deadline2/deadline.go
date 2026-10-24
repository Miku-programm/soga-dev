package deadline2

import (
	"sync"
	"sync/atomic"
	"time"
)

// TErFW7l9B2S 为原 32 字节结构，字段偏移 0/8/16/24。
type TErFW7l9B2S struct {
	kO0YDf    sync.Mutex
	uwIvic    *time.Timer
	qb0z25Hn  chan struct{}
	rEp0gZcar atomic.Bool
}

// NewDeadline 表达 HpSkqj @ 0x18ae4e0 中被内联的构造。
// 此名称是恢复工程辅助名称，不冒充样本中存在独立的构造函数。
func NewDeadline() *TErFW7l9B2S { return &TErFW7l9B2S{qb0z25Hn: make(chan struct{})} }

// Set @ 0x124ee80，timer 回调 @ 0x124f780。
func (d *TErFW7l9B2S) Set(deadline time.Time) {
	d.kO0YDf.Lock()
	defer d.kO0YDf.Unlock()
	if d.uwIvic != nil && !d.uwIvic.Stop() {
		<-d.qb0z25Hn
	}
	d.uwIvic = nil
	closed := false
	select {
	case <-d.qb0z25Hn:
		closed = true
	default:
	}
	if deadline.IsZero() {
		if closed {
			d.qb0z25Hn = make(chan struct{})
		}
		d.rEp0gZcar.Store(false)
		return
	}
	d.rEp0gZcar.Store(true)
	duration := time.Until(deadline)
	if duration <= 0 {
		if !closed {
			close(d.qb0z25Hn)
		}
		return
	}
	if closed {
		d.qb0z25Hn = make(chan struct{})
	}
	channel := d.qb0z25Hn
	d.uwIvic = time.AfterFunc(duration, func() { close(channel) })
}
func (d *TErFW7l9B2S) Active() bool { return d.rEp0gZcar.Load() } // 0x124f180
func (d *TErFW7l9B2S) Wait() <-chan struct{} {
	d.kO0YDf.Lock()
	defer d.kO0YDf.Unlock()
	return d.qb0z25Hn
} // 0x124f1a0
