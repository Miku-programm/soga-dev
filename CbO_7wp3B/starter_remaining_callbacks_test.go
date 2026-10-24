//go:build linux

package CbO_7wp3B

import (
	"crypto/rand"
	"testing"
	"time"

	"recovered.local/soga/cNT24rG4F"
)

func TestRemainingStarterAttemptsPreserveEntropyAndSignal(t *testing.T) {
	old := rand.Reader
	t.Cleanup(func() { rand.Reader = old; cNT24rG4F.T15Jifo(false) })
	for _, attempt := range []func(*int, chan int, chan struct{}){ppanelCallbackAttempt, sogaV1CallbackAttempt, sspanelCallbackAttempt} {
		for _, tc := range []struct{ value, bytes int }{{0, 296}, {2, 296}, {3, 552}, {100, 552}} {
			r := &thresholdEntropy{}
			rand.Reader = r
			queue, done := make(chan int, 1), make(chan struct{}, 1)
			value := tc.value
			attempt(&value, queue, done)
			select {
			case got := <-queue:
				if got != 0 {
					t.Fatalf("successor %d", got)
				}
			case <-time.After(time.Second):
				t.Fatal("missing successor")
			}
			if len(done) != 1 || r.read != tc.bytes {
				t.Fatalf("value %d: signal %d entropy %d, expected %d", value, len(done), r.read, tc.bytes)
			}
		}
	}
}
