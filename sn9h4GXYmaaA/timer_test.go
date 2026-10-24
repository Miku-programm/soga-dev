package sn9h4GXYmaaA

import (
	"testing"
	"time"
)

func TestOriginalTimerFireAndIdempotentCancel(t *testing.T) {
	fired := make(chan struct{}, 1)
	cancel := LwNqe0FPzq(time.Millisecond, func() { fired <- struct{}{} })
	select {
	case <-fired:
	case <-time.After(time.Second):
		t.Fatal("timer did not fire")
	}
	cancel()
	cancel()
	fired = make(chan struct{}, 1)
	cancel = LwNqe0FPzq(50*time.Millisecond, func() { fired <- struct{}{} })
	cancel()
	cancel()
	select {
	case <-fired:
		t.Fatal("cancelled timer fired")
	case <-time.After(70 * time.Millisecond):
	}
	if ClientVersionAt5098D50 != "anytls/0.0.11" {
		t.Fatal(ClientVersionAt5098D50)
	}
}
