package limit

import (
	"context"
	"crypto/tls"
	"golang.org/x/sync/singleflight"
	C "recovered.local/soga/config"
	N "recovered.local/soga/user"
	R "recovered.local/soga/net/redis"
	"sync"
	"sync/atomic"
	"time"
)

// Original ZpnJ03Vh46F interface; concrete methods retain their contracts.
type ZpnJ03Vh46F interface {
	EndUserConn(N.PgaNd3014U, string)
	StartUserConn(context.Context, context.CancelFunc, func(), N.PgaNd3014U, string) bool
}
type uuMFf5B struct {
	dCzJC4or4Xl int64
	gDlwyqlt    string
}
type e36g5GkEd struct {
	a7fxke      int
	aXnE_3N_gdt bool
}
type fj057PaXtu struct {
	k4kEqO3    []string
	eaCklg8jch time.Time
}

// YM8R9i: 104-byte RTTI0x36fc9d8. The Redis pointer maps to its verified library.
type YM8R9i struct {
	eHz2FQHNc  context.Context
	zAD6a7zB   *R.HXzAD33aVBDR
	p7KF5sbP   time.Duration
	p81VkaYJai sync.Mutex
	pCsxt7t    map[uuMFf5B]*e36g5GkEd
	nttEJ6     sync.RWMutex
	xLikNInZ   map[int64]*fj057PaXtu
	zHP8cN6    singleflight.Group
	cmA3rUr    atomic.Bool
}

// Cder0rQ:0x1835e20. The enable option is checked by its caller, not here.
func Cder0rQ(ctx context.Context) ZpnJ03Vh46F {
	_, address, password, db, useTLS, seconds := C.RedisIPLimitConfiguration()
	expiry := time.Duration(seconds) * time.Second
	if expiry < 2*time.Second {
		expiry = 2 * time.Second
	}
	o := &R.Csx1an{Addr: address, Password: password, DB: db}
	if useTLS {
		o.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	m := &YM8R9i{eHz2FQHNc: ctx, zAD6a7zB: R.ZHNMa9Nz(o), p7KF5sbP: expiry, pCsxt7t: make(map[uuMFf5B]*e36g5GkEd), xLikNInZ: make(map[int64]*fj057PaXtu)}
	m.vih7o_()
	m.fV5Dht3hwUlM()
	return m
}

// EndUserConn:0x1836b60. It always looks up and decrements the local entry,
// independent of the current user limit or Redis health.
func (m *YM8R9i) EndUserConn(user N.PgaNd3014U, ip string) {
	key := uuMFf5B{user.GetId(), ip}
	m.p81VkaYJai.Lock()
	defer m.p81VkaYJai.Unlock()
	if item, ok := m.pCsxt7t[key]; ok {
		item.a7fxke--
		if item.a7fxke <= 0 {
			delete(m.pCsxt7t, key)
		}
	}
}
