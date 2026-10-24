package dialctx

import (
	"sync"
	"time"
)

// X7WcH4QYYyc6 restores the 56-byte layout read at 0x19204c0.
// Its field names are not retained in RTTI; the descriptive names record offsets.
type X7WcH4QYYyc6 struct {
	mu       sync.Mutex    // +0
	started  time.Time     // +8
	duration time.Duration // +32
	call     *tdfGa3B_     // +40
	result   *SRClJmN_     // +48
}
type SRClJmN_ struct {
	WI24xSJ   any
	OCGlGkALT error
}
type tdfGa3B_ struct {
	mWqKaiStC56 sync.WaitGroup
	ep8BE0      any
	iebH9pdHU2  error
}

// NewCache exposes original static/inlined construction across restored packages.
// It is a source organization helper, not a retained original function symbol.
func NewCache(duration time.Duration) *X7WcH4QYYyc6 { return &X7WcH4QYYyc6{duration: duration} }

// Do: 0x19204c0. Both errors and successful results are cached. Expiry starts
// before the work begins. The original has no panic recovery for the callback.
func (g *X7WcH4QYYyc6) Do(fn func() (any, error)) (any, error, bool) {
	g.mu.Lock()
	now := time.Now()
	if now.Before(g.started.Add(g.duration)) {
		g.mu.Unlock()
		return g.result.WI24xSJ, g.result.OCGlGkALT, true
	}
	if c := g.call; c != nil {
		g.mu.Unlock()
		c.mWqKaiStC56.Wait()
		return c.ep8BE0, c.iebH9pdHU2, true
	}
	c := new(tdfGa3B_)
	c.mWqKaiStC56.Add(1)
	g.call = c
	g.mu.Unlock()
	c.ep8BE0, c.iebH9pdHU2 = fn()
	c.mWqKaiStC56.Done()
	g.mu.Lock()
	g.call = nil
	g.result = &SRClJmN_{c.ep8BE0, c.iebH9pdHU2}
	g.started = now
	g.mu.Unlock()
	return c.ep8BE0, c.iebH9pdHU2, false
}
