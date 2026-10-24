package dns

import (
	"context"
	"github.com/jellydator/ttlcache/v3"
	"net"
	"recovered.local/soga/util/rtx"
	"time"
)

// u838o9bPnv: 0x1592740；清理闭包 0x1595d00。
func u838o9bPnv(ctx context.Context, ttl time.Duration) *yMnO8y {
	var cache *ttlcache.Cache[string, []net.IP]
	if ttl > 0 {
		cache = ttlcache.New[string, []net.IP](ttlcache.WithTTL[string, []net.IP](ttl), ttlcache.WithDisableTouchOnHit[string, []net.IP]())
		rtx.SwUrlM3M(ctx, cache.DeleteExpired, ttl)
	}
	return &yMnO8y{cache, ttl > 0}
}

// AddIP: 0x1592920。
func (c *yMnO8y) AddIP(host string, ips []net.IP) {
	if c.iG4qMhe8R {
		c.nSAwGB1m2.Set(host, ips, 0)
	}
}

// Get: 0x15929c0。
func (c *yMnO8y) Get(host string) ([]net.IP, bool) {
	if c.iG4qMhe8R {
		if item := c.nSAwGB1m2.Get(host); item != nil {
			return item.Value(), true
		}
	}
	return nil, false
}
