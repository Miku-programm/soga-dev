package app

import (
	"context"
	"crypto/rand"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"recovered.local/soga/protocol/inbound"
	"recovered.local/soga/config"
	"recovered.local/soga/panel/registry"
)

type submitLoopPanel struct {
	report func(*registry.FGmNba4MhioY) error
}

func (*submitLoopPanel) GetAPIConfig() *config.FCEAfQJwDA7e { panic("unexpected config request") }
func (*submitLoopPanel) GetAPIData() (*registry.D5gTIZ5aVg, error) {
	panic("unexpected data request")
}
func (p *submitLoopPanel) SubmitData(data *registry.FGmNba4MhioY) error { return p.report(data) }

type submitLoopInbound struct {
	data *registry.FGmNba4MhioY
}

func (i *submitLoopInbound) GetSubmitData() *registry.FGmNba4MhioY { return i.data }
func (*submitLoopInbound) Start() error                                { panic("unexpected start") }
func (*submitLoopInbound) Stop()                                       { panic("unexpected stop") }
func (*submitLoopInbound) UpdateAPIData(*registry.D5gTIZ5aVg) bool { panic("unexpected update") }
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
	data := &registry.FGmNba4MhioY{}
	i := &submitLoopInbound{data: data}
	var interval atomic.Int64
	var calls int
	p := &submitLoopPanel{report: func(got *registry.FGmNba4MhioY) error {
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
	starterSubmitLoop(ctx, p, func() inbound.Kz_8ajU7 { return i }, &interval)
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
	p := &submitLoopPanel{report: func(*registry.FGmNba4MhioY) error { t.Fatal("typed nil submitted"); return nil }}
	starterSubmitLoop(ctx, p, func() inbound.Kz_8ajU7 {
		gets++
		interval.Store(int64(time.Hour))
		cancel()
		return i
	}, &interval)
	if gets != 1 {
		t.Fatalf("getter calls = %d", gets)
	}
}
