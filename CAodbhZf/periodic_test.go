package CAodbhZf

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestSwUrlM3MCancellationAndSerialCallbacks(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	entered := make(chan struct{}, 1)
	release := make(chan struct{})
	var active, calls atomic.Int32
	SwUrlM3M(ctx, func() {
		if active.Add(1) != 1 {
			t.Error("parallel callback")
		}
		calls.Add(1)
		select {
		case entered <- struct{}{}:
		default:
		}
		<-release
		active.Add(-1)
	}, 10*time.Millisecond)
	select {
	case <-entered:
	case <-time.After(time.Second):
		t.Fatal("ticker never invoked callback")
	}
	// The callback is running while multiple ticks elapse; it must not be reentered.
	time.Sleep(30 * time.Millisecond)
	if calls.Load() != 1 {
		t.Fatal("callback overlapped")
	}
	cancel()
	close(release)
	// One already-ready tick is permitted to win select after cancellation.
	time.Sleep(30 * time.Millisecond)
	n := calls.Load()
	time.Sleep(30 * time.Millisecond)
	if calls.Load() != n {
		t.Fatal("ticker continues after cancellation")
	}
}
