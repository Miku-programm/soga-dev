package CbO_7wp3B

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"recovered.local/soga/MFFa9yT"
	"recovered.local/soga/q5bqCgTdxaTM"
)

type pullLoopPanel struct {
	submitLoopPanel
	fetch func() (*q5bqCgTdxaTM.D5gTIZ5aVg, error)
}

func (p *pullLoopPanel) GetAPIData() (*q5bqCgTdxaTM.D5gTIZ5aVg, error) { return p.fetch() }

type pullLoopInbound struct {
	submitLoopInbound
	update func(*q5bqCgTdxaTM.D5gTIZ5aVg) bool
	stop   func()
}

func (i *pullLoopInbound) UpdateAPIData(d *q5bqCgTdxaTM.D5gTIZ5aVg) bool { return i.update(d) }
func (i *pullLoopInbound) Stop()                                         { i.stop() }

func TestStarterPullLoopRestartFailureAndRetry(t *testing.T) {
	zeroSubmitJitter(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var mu sync.Mutex
	var events []string
	record := func(s string) { mu.Lock(); events = append(events, s); mu.Unlock() }
	stopped := make(chan struct{})
	first := &pullLoopInbound{update: func(*q5bqCgTdxaTM.D5gTIZ5aVg) bool { record("update-first"); return true }, stop: func() { record("stop-first") }}
	last := &pullLoopInbound{stop: func() { record("stop-last"); close(stopped) }}
	var current MFFa9yT.Kz_8ajU7
	var after func() bool
	var check, submit atomic.Int64
	data := &q5bqCgTdxaTM.D5gTIZ5aVg{Xdo1D9: 1, DDu2Nm_P: 456}
	panel := &pullLoopPanel{fetch: func() (*q5bqCgTdxaTM.D5gTIZ5aVg, error) { record("fetch"); return data, nil }}
	builds := 0
	starterPullLoop(ctx, panel, func() MFFa9yT.Kz_8ajU7 { return current }, func(i MFFa9yT.Kz_8ajU7) {
		current = i
		if i == nil {
			record("clear")
		} else {
			record("publish")
		}
	},
		func(context.Context, *q5bqCgTdxaTM.D5gTIZ5aVg) (MFFa9yT.Kz_8ajU7, error) {
			builds++
			switch builds {
			case 1:
				record("build-first")
				return first, nil
			case 2:
				record("build-failed")
				return nil, errors.New("bind failed")
			default:
				record("build-last")
				if check.Load() != 1 || submit.Load() != 456 {
					t.Errorf("intervals %d %d", check.Load(), submit.Load())
				}
				check.Store(int64(time.Hour))
				cancel()
				return last, nil
			}
		}, &check, &submit, &after)
	select {
	case <-stopped:
	case <-time.After(time.Second):
		t.Fatal("cancellation did not stop last inbound")
	}
	mu.Lock()
	got := append([]string(nil), events...)
	mu.Unlock()
	// The final AfterFunc is allowed to run before publication after cancellation.
	withoutStop := make([]string, 0, len(got))
	for _, e := range got {
		if e != "stop-last" {
			withoutStop = append(withoutStop, e)
		}
	}
	want := []string{"fetch", "build-first", "publish", "fetch", "update-first", "stop-first", "clear", "build-failed", "fetch", "build-last", "publish"}
	if !reflect.DeepEqual(withoutStop, want) {
		t.Fatalf("events %v", got)
	}
	if check.Load() != int64(time.Hour) || submit.Load() != 456 {
		t.Fatalf("intervals %d %d", check.Load(), submit.Load())
	}
}

func TestStarterPullLoopNoRestartKeepsCurrent(t *testing.T) {
	zeroSubmitJitter(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	current := &pullLoopInbound{update: func(*q5bqCgTdxaTM.D5gTIZ5aVg) bool { cancel(); return false }, stop: func() { t.Error("unexpected stop") }}
	panel := &pullLoopPanel{fetch: func() (*q5bqCgTdxaTM.D5gTIZ5aVg, error) {
		return &q5bqCgTdxaTM.D5gTIZ5aVg{Xdo1D9: -1, DDu2Nm_P: -2}, nil
	}}
	var check, submit atomic.Int64
	check.Store(int64(time.Hour))
	submit.Store(456)
	var after func() bool
	starterPullLoop(ctx, panel, func() MFFa9yT.Kz_8ajU7 { return current }, func(MFFa9yT.Kz_8ajU7) { t.Fatal("unexpected publication") },
		func(context.Context, *q5bqCgTdxaTM.D5gTIZ5aVg) (MFFa9yT.Kz_8ajU7, error) {
			t.Fatal("unexpected rebuild")
			return nil, nil
		}, &check, &submit, &after)
	if check.Load() != int64(time.Hour) || submit.Load() != 456 {
		t.Fatal("nonpositive panel interval replaced current interval")
	}
}
