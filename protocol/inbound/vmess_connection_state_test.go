package inbound

import (
	"testing"

	"recovered.local/soga/util/rtx"
)

func TestVMessStateMethodsWithinInterval(t *testing.T) {
	old := userCountAt50C6DA0
	defer func() { userCountAt50C6DA0 = old }()
	v := &K4avFiYOPi{JWN154qJ4: &JWN154qJ4{fUxO7UYAH: 83, mKB02yhVfTfP: 4}}
	// This is the original proved no-goroutine branch, not an authorization bypass.
	v.jfLH8mvq5e03.Store(uint32(rtx.C5Jff5vs0()))
	for _, method := range []func(){v.sbbdX3aT6E, v.bj10P7X} {
		method()
		if userCountAt50C6DA0 != 83 || v.mKB02yhVfTfP != 4 {
			t.Fatal("state changed within interval")
		}
	}
}
