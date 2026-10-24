package NtxdAE

import (
	"context"
	"github.com/jellydator/ttlcache/v3"
	L "recovered.local/soga/muc279Mpa6S4"
	N "recovered.local/soga/nzATvajt"
	"time"
)

// ZWHtxl: original 8-byte RTTI 0x35efcf0.
type ZWHtxl struct {
	jnsGsOq2H2ov *ttlcache.Cache[int64, float64]
}

// SetUser: 0x1834b00. This only inserts the cache value; it does not change the
// user's limiter. The caller performs that separate original action.
func (c *ZWHtxl) SetUser(u N.PgaNd3014U, speed float64, ttl time.Duration) {
	c.jnsGsOq2H2ov.Set(u.GetId(), speed, ttl)
}

// HasUser: 0x1834bc0.
func (c *ZWHtxl) HasUser(id int64) bool { return c.jnsGsOq2H2ov.Get(id) != nil }

// GetUserSpeedLimit: 0x1834c20.
func (c *ZWHtxl) GetUserSpeedLimit(id int64) (float64, bool) {
	item := c.jnsGsOq2H2ov.Get(id)
	if item == nil {
		return 0, false
	}
	return item.Value(), true
}

// DelUser: 0x1834ca0.
func (c *ZWHtxl) DelUser(id int64) { c.jnsGsOq2H2ov.Delete(id) }

// uFaljJ: 0x1834900, eviction closure 0x183dc20 and sweeper 0x183dac0.
// Eviction of every reason resets the currently registered user, not a captured
// old user object, and does not restore the numeric cache value as its speed.
func uFaljJ(ctx context.Context, users *N.ZpNFqUVZp) *ZWHtxl {
	cache := ttlcache.New[int64, float64](ttlcache.WithTTL[int64, float64](5*time.Second), ttlcache.WithDisableTouchOnHit[int64, float64]())
	cache.OnEviction(func(_ context.Context, _ ttlcache.EvictionReason, item *ttlcache.Item[int64, float64]) {
		if user, ok := users.Get(item.Key()); ok {
			old := user.GetSpeedLimit()
			user.SetSpeedLimit(0)
			current := user.GetSpeedLimit()
			L.V8GN5Stk(-4, "user %v recover speed limit: %.2f Mbps -> %.2f Mbps", item.Key(), old, current)
		}
	})
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				cache.DeleteExpired()
			case <-ctx.Done():
				return
			}
		}
	}()
	return &ZWHtxl{jnsGsOq2H2ov: cache}
}
