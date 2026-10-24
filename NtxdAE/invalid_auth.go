package NtxdAE

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/juju/ratelimit"
	"recovered.local/soga/CAodbhZf"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/muc279Mpa6S4"
)

// M1Qn9BA preserves the original 64-byte invalid-authentication limiter.
type M1Qn9BA struct {
	iOAqGa9oV_  context.Context
	eDtCKI      *cNT24rG4F.Nl62Lnd_
	zeXkYKDmp9a sync.RWMutex
	q1T7LqXXEQ  *ttlcache.Cache[string, *ratelimit.Bucket]
	uiXaTdSBz   *ttlcache.Cache[string, struct{}]
}

// Azb3Zr recovers 0x1833020, retaining disabled caches as nil.
func Azb3Zr(ctx context.Context, config *cNT24rG4F.Nl62Lnd_) *M1Qn9BA {
	result := &M1Qn9BA{iOAqGa9oV_: ctx, eDtCKI: config}
	if config.SEU_Z687ugp {
		result.q1T7LqXXEQ = ttlcache.New[string, *ratelimit.Bucket](ttlcache.WithTTL[string, *ratelimit.Bucket](config.QS255wlN7m2W))
		result.uiXaTdSBz = ttlcache.New[string, struct{}](ttlcache.WithTTL[string, struct{}](config.NjiRkFvidOaA))
		result.dTRa2vHe5()
	}
	return result
}

// 0x1833200 and closure 0x183d9c0: one-minute expiry cleanup, same lock order.
func (limiter *M1Qn9BA) dTRa2vHe5() {
	if limiter.eDtCKI.SEU_Z687ugp {
		CAodbhZf.SwUrlM3M(limiter.iOAqGa9oV_, func() {
			limiter.zeXkYKDmp9a.Lock()
			limiter.q1T7LqXXEQ.DeleteExpired()
			limiter.uiXaTdSBz.DeleteExpired()
			limiter.zeXkYKDmp9a.Unlock()
		}, time.Minute)
	}
}

// Valid recovers 0x18332a0, clearing both counters and the ban.
func (limiter *M1Qn9BA) Valid(address string) {
	if !limiter.eDtCKI.SEU_Z687ugp {
		return
	}
	limiter.zeXkYKDmp9a.Lock()
	limiter.uiXaTdSBz.Delete(address)
	limiter.q1T7LqXXEQ.Delete(address)
	limiter.zeXkYKDmp9a.Unlock()
}

// Invalid recovers 0x1833360 and bucket closure 0x183da40. Refill is a
// complete capacity-sized quantum per interval, not one token per interval.
func (limiter *M1Qn9BA) Invalid(address string) bool {
	if !limiter.eDtCKI.SEU_Z687ugp {
		return false
	}
	limiter.zeXkYKDmp9a.Lock()
	defer limiter.zeXkYKDmp9a.Unlock()
	item, _ := limiter.q1T7LqXXEQ.GetOrSetFunc(address, func() *ratelimit.Bucket {
		config := limiter.eDtCKI
		return ratelimit.NewBucketWithQuantum(config.QS255wlN7m2W, int64(config.GxuEkr), int64(config.GxuEkr))
	})
	if item.Value().TakeAvailable(1) > 0 {
		return false
	}
	limiter.q1T7LqXXEQ.Delete(address)
	limiter.uiXaTdSBz.Set(address, struct{}{}, ttlcache.DefaultTTL)
	// Exact decoder@0x1833523: original log reports the counting duration.
	muc279Mpa6S4.V8GN5Stk(slog.LevelDebug, "ip %v invalid access limit time: %v", address, limiter.eDtCKI.QS255wlN7m2W)
	return true
}

// IsForbidden recovers 0x1833620. Has does not extend the ban's TTL.
func (limiter *M1Qn9BA) IsForbidden(address string) bool {
	if !limiter.eDtCKI.SEU_Z687ugp {
		return false
	}
	limiter.zeXkYKDmp9a.RLock()
	result := limiter.uiXaTdSBz.Has(address)
	limiter.zeXkYKDmp9a.RUnlock()
	return result
}
