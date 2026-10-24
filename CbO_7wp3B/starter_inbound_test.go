//go:build linux

package CbO_7wp3B

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/MFFa9yT"
	"recovered.local/soga/q5bqCgTdxaTM"
)

type starterTestInbound struct {
	submitLoopInbound
	start func() error
	stop  func()
}

func (i *starterTestInbound) Start() error { return i.start() }
func (i *starterTestInbound) Stop()        { i.stop() }

func TestOriginalInboundCancellationFailureAndOnce(t *testing.T) {
	starterOnceAt50C7AA0 = sync.Once{}
	t.Cleanup(func() { starterOnceAt50C7AA0 = sync.Once{} })
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	callback := func(int) { t.Error("unexpected callback") }
	construct := func(*MFFa9yT.DV_XPJ4z0) (MFFa9yT.Kz_8ajU7, error) {
		t.Error("unexpected construction")
		return nil, nil
	}
	got, err := startOriginalInbound(ctx, nil, callback, construct)
	if got != nil || err != context.Canceled {
		t.Fatal(got, err)
	}
	data := &q5bqCgTdxaTM.D5gTIZ5aVg{H47AzZsYHB1N: &LFWTTXF7JR.Wr5SVP{}}
	want := errors.New("factory failure")
	got, err = startOriginalInbound(context.Background(), data, callback, func(*MFFa9yT.DV_XPJ4z0) (MFFa9yT.Kz_8ajU7, error) { return nil, want })
	if got != nil || err != want {
		t.Fatal(got, err)
	}
	called := make(chan int, 2)
	starts, stops := 0, 0
	startErr := errors.New("start failure")
	i := &starterTestInbound{start: func() error { starts++; return startErr }, stop: func() { stops++ }}
	construct = func(arg *MFFa9yT.DV_XPJ4z0) (MFFa9yT.Kz_8ajU7, error) {
		if arg.HG5k0365 != data {
			t.Error("data identity changed")
		}
		return i, nil
	}
	got, err = startOriginalInbound(context.Background(), data, func(n int) { called <- n }, construct)
	if got != nil || err != startErr || starts != 1 || stops != 1 {
		t.Fatal(got, err, starts, stops)
	}
	select {
	case n := <-called:
		if n != 0 {
			t.Fatal(n)
		}
	case <-time.After(time.Second):
		t.Fatal("missing callback")
	}
	i.start = func() error { starts++; return nil }
	got, err = startOriginalInbound(context.Background(), data, func(n int) { called <- n }, construct)
	if got != i || err != nil || starts != 2 || stops != 1 {
		t.Fatal(got, err, starts, stops)
	}
	select {
	case <-called:
		t.Fatal("Once callback repeated after failed Start")
	default:
	}
}
