package limit

import (
	"context"
	"github.com/jellydator/ttlcache/v3"
	"time"
)

// HRyRearcsgu: 16-byte RTTI 0x3631730. The ring values themselves have no lock
// in the original; retain that property rather than adding new synchronization.
type HRyRearcsgu struct {
	ctvhRd  int
	baRtduy *ttlcache.Cache[int64, *Mblp8L]
}

// g9s9GzLu9u: 0x183d4e0.
func (c *HRyRearcsgu) g9s9GzLu9u(id int64) (*Mblp8L, bool) {
	item := c.baRtduy.Get(id)
	if item == nil {
		return nil, false
	}
	return item.Value(), true
}

// a2bBbnhBw: 0x183d560.
func (c *HRyRearcsgu) a2bBbnhBw(id int64) *Mblp8L {
	w := &Mblp8L{d5H9TK: make([]int64, c.ctvhRd)}
	c.baRtduy.Set(id, w, ttlcache.DefaultTTL)
	return w
}

// AddUserTraffic: 0x183d640. An idle zero ring is deleted only on the following
// zero sample; the sample that reduces its sum to zero still refreshes its TTL.
func (c *HRyRearcsgu) AddUserTraffic(id, traffic int64) int64 {
	w, ok := c.g9s9GzLu9u(id)
	if traffic != 0 {
		if !ok {
			w = c.a2bBbnhBw(id)
		}
	} else {
		if !ok {
			return 0
		}
		if w.qzaQaE4 == 0 {
			c.baRtduy.Delete(id)
			return 0
		}
	}
	w.AddTraffic(traffic)
	c.baRtduy.Set(id, w, ttlcache.DefaultTTL)
	return w.qzaQaE4
}

// b3D2x6m1: 0x183d360; closure 0x183ef80. Cleanup uses the full window
// duration, not each five-second sample; it exits with the supplied context.
func b3D2x6m1(ctx context.Context, slots int) *HRyRearcsgu {
	ttl := time.Duration(slots) * 5 * time.Second
	cache := ttlcache.New[int64, *Mblp8L](ttlcache.WithTTL[int64, *Mblp8L](ttl))
	go func() {
		ticker := time.NewTicker(ttl)
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
	return &HRyRearcsgu{ctvhRd: slots, baRtduy: cache}
}
