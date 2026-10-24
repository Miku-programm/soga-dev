package rtx

import (
	"context"
	"testing"
	"time"
)

func TestUXNvag9ExpirationAndCancellation(t *testing.T) {
	if !UXNvag9(context.Background(), time.NewTimer(0)) {
		t.Fatal("expiration returned cancellation")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	timer := time.NewTimer(time.Hour)
	if UXNvag9(ctx, timer) {
		t.Fatal("cancellation returned expiration")
	}
	if timer.Stop() {
		t.Fatal("cancellation did not stop timer")
	}
}
