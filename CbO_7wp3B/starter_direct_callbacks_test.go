//go:build linux

package CbO_7wp3B

import (
	"errors"
	"testing"
	"time"

	"recovered.local/soga/cNT24rG4F"
)

func TestDirectStarterOriginalRetrySchedule(t *testing.T) {
	t.Cleanup(func() { cNT24rG4F.T15Jifo(false) })
	for _, failures := range []int{0, 2, 3} {
		calls := 0
		start := time.Now()
		directStarterCallback(123, func() error {
			calls++
			if calls <= failures {
				return errors.New("fixture failure")
			}
			return nil
		})
		want := failures + 1
		if want > 3 {
			want = 3
		}
		if calls != want || time.Since(start) < time.Duration(want-1)*time.Second {
			t.Fatalf("failures=%d calls=%d elapsed=%s", failures, calls, time.Since(start))
		}
	}
}
