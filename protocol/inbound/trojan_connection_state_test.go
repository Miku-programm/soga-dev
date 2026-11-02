package inbound

import (
	"recovered.local/soga/util/rtx"
	"testing"
)

// Exercise both real methods only inside the original no-job interval.
func TestTrojanStateRecentCheckUpdatesSharedCount(t *testing.T) {
	originalCount := userCountAt50C6DA0
	t.Cleanup(func() { userCountAt50C6DA0 = originalCount })
	v := &EfaIFpp{JWN154qJ4: &JWN154qJ4{fUxO7UYAH: 81}}
	v.jfLH8mvq5e03.Store(uint32(rtx.C5Jff5vs0()))
	v.sbbdX3aT6E()
	if userCountAt50C6DA0 != 81 {
		t.Fatal(userCountAt50C6DA0)
	}
	v.fUxO7UYAH = 123
	v.bj10P7X()
	if userCountAt50C6DA0 != 123 {
		t.Fatal(userCountAt50C6DA0)
	}
}
