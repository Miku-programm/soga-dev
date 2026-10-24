package limit

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
	"unsafe"

	"recovered.local/soga/config"
)

func TestInvalidAuthDisabledAndRTTI(t *testing.T) {
	limiter := Azb3Zr(context.Background(), &config.Nl62Lnd_{})
	if unsafe.Sizeof(*limiter) != 64 || unsafe.Offsetof(limiter.q1T7LqXXEQ) != 48 || unsafe.Sizeof(config.Nl62Lnd_{}) != 32 {
		t.Fatal("RTTI mismatch")
	}
	if limiter.q1T7LqXXEQ != nil || limiter.uiXaTdSBz != nil {
		t.Fatal("disabled allocated caches")
	}
	limiter.Valid("ip")
	if limiter.Invalid("ip") || limiter.IsForbidden("ip") {
		t.Fatal("disabled limiter active")
	}
}

func TestInvalidAuthThresholdResetExpiryAndRefill(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		config := &config.Nl62Lnd_{SEU_Z687ugp: true, GxuEkr: 2, QS255wlN7m2W: 10 * time.Second, NjiRkFvidOaA: 20 * time.Second}
		limiter := Azb3Zr(ctx, config)
		if limiter.Invalid("a") || limiter.Invalid("a") {
			t.Fatal("banned before capacity exhausted")
		}
		if limiter.IsForbidden("a") || !limiter.Invalid("a") || !limiter.IsForbidden("a") {
			t.Fatal("third failure did not ban")
		}
		if limiter.q1T7LqXXEQ.Has("a") {
			t.Fatal("counter retained on ban")
		}
		if limiter.Invalid("b") {
			t.Fatal("other IP shared count")
		}
		time.Sleep(15 * time.Second)
		if !limiter.IsForbidden("a") {
			t.Fatal("ban expired early")
		}
		time.Sleep(6 * time.Second)
		if limiter.IsForbidden("a") {
			t.Fatal("Has extended ban expiry")
		}
		if limiter.Invalid("a") || limiter.Invalid("a") {
			t.Fatal("new bucket did not refill")
		}
		limiter.Valid("a")
		if limiter.Invalid("a") || limiter.Invalid("a") {
			t.Fatal("Valid did not clear counter")
		}
		if !limiter.Invalid("a") {
			t.Fatal("ban missing")
		}
		limiter.Valid("a")
		if limiter.IsForbidden("a") {
			t.Fatal("Valid did not clear ban")
		}
		// Both tokens replenish together after one fill interval.
		if limiter.Invalid("c") || limiter.Invalid("c") {
			t.Fatal("initial tokens")
		}
		time.Sleep(11 * time.Second)
		if limiter.Invalid("c") || limiter.Invalid("c") {
			t.Fatal("full-quantum refill lost")
		}
		time.Sleep(time.Minute)
		if limiter.q1T7LqXXEQ.Len() != 0 || limiter.uiXaTdSBz.Len() != 0 {
			t.Fatal("periodic cleanup failed")
		}
	})
}
