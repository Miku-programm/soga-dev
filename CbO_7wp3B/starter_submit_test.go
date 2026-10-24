package CbO_7wp3B

import (
	"context"
	"crypto/rand"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"recovered.local/soga/MFFa9yT"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/q5bqCgTdxaTM"
)

type submitLoopPanel struct {
	report func(*q5bqCgTdxaTM.FGmNba4MhioY) error
}

func (*submitLoopPanel) GetAPIConfig() *cNT24rG4F.FCEAfQJwDA7e { panic("unexpected config request") }
func (*submitLoopPanel) GetAPIData() (*q5bqCgTdxaTM.D5gTIZ5aVg, error) {
	panic("unexpected data request")
}
func (p *submitLoopPanel) SubmitData(data *q5bqCgTdxaTM.FGmNba4MhioY) error { return p.report(data) }

type submitLoopInbound struct {
	data *q5bqCgTdxaTM.FGmNba4MhioY
}

func (i *submitLoopInbound) GetSubmitData() *q5bqCgTdxaTM.FGmNba4MhioY { return i.data }
func (*submitLoopInbound) Start() error                                { panic("unexpected start") }
func (*submitLoopInbound) Stop()                                       { panic("unexpected stop") }
func (*submitLoopInbound) UpdateAPIData(*q5bqCgTdxaTM.D5gTIZ5aVg) bool { panic("unexpected update") }
func (*submitLoopInbound) UserExist(int64) bool                        { panic("unexpected user lookup") }

type zeroIntervalReader struct{}

func (zeroIntervalReader) Read(p []byte) (int, error) { clear(p); return len(p), nil }
func zeroSubmitJitter(t *testing.T) {
	old := rand.Reader
	rand.Reader = zeroIntervalReader{}
	t.Cleanup(func() { rand.Reader = old })
}

func TestStarterSubmitLoopPreservesReportAndRetry(t *testing.T) {
	zeroSubmitJitter(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	data := &q5bqCgTdxaTM.FGmNba4MhioY{}
	i := &submitLoopInbound{data: data}
	var interval atomic.Int64
	var calls int
	p := &submitLoopPanel{report: func(got *q5bqCgTdxaTM.FGmNba4MhioY) error {
		if got != data {
			t.Fatal("report identity changed")
		}
		calls++
		if calls == 1 {
			return errors.New("first request failed")
		}
		interval.Store(int64(time.Hour))
		cancel()
		return nil
	}}
	starterSubmitLoop(ctx, p, func() MFFa9yT.Kz_8ajU7 { return i }, &interval)
	if calls != 2 {
		t.Fatalf("reports = %d, want retry after first error", calls)
	}
}

func TestStarterSubmitLoopTypedNilAndCancellation(t *testing.T) {
	zeroSubmitJitter(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var interval atomic.Int64
	var i *submitLoopInbound
	gets := 0
	p := &submitLoopPanel{report: func(*q5bqCgTdxaTM.FGmNba4MhioY) error { t.Fatal("typed nil submitted"); return nil }}
	starterSubmitLoop(ctx, p, func() MFFa9yT.Kz_8ajU7 {
		gets++
		interval.Store(int64(time.Hour))
		cancel()
		return i
	}, &interval)
	if gets != 1 {
		t.Fatalf("getter calls = %d", gets)
	}
}
