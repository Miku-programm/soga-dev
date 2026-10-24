package NtxdAE

import (
	"context"
	"fmt"
	"math/rand"
	L "recovered.local/soga/muc279Mpa6S4"
	"strconv"
	"time"
)

// n_augufdnS:0x1837000. Expiration uses strict After; no deletion or copy here.
func (m *YM8R9i) n_augufdnS(id int64) ([]string, bool) {
	m.nttEJ6.RLock()
	v, ok := m.xLikNInZ[id]
	m.nttEJ6.RUnlock()
	if ok && !time.Now().After(v.eaCklg8jch) {
		return v.k4kEqO3, true
	}
	return nil, false
}

// nkCoQC:0x1837100. Cache TTL is min(expiry,5s) minus up to ten percent jitter.
func (m *YM8R9i) nkCoQC(id int64, ips []string) {
	now := time.Now()
	ttl := min(m.p7KF5sbP, 5*time.Second)
	expires := now.Add(ttl - time.Duration(rand.Int63n(int64(ttl/10+1))))
	v := &fj057PaXtu{ips, expires}
	m.nttEJ6.Lock()
	m.xLikNInZ[id] = v
	m.nttEJ6.Unlock()
}

// ba5WPJg:0x18372e0. Copy-on-append leaves already-returned slices unchanged;
// missing/expired entries are left untouched, and expiry is not extended.
func (m *YM8R9i) ba5WPJg(id int64, ip string) {
	m.nttEJ6.Lock()
	defer m.nttEJ6.Unlock()
	v, ok := m.xLikNInZ[id]
	if !ok || time.Now().After(v.eaCklg8jch) {
		return
	}
	for _, old := range v.k4kEqO3 {
		if old == ip {
			return
		}
	}
	ips := make([]string, len(v.k4kEqO3), len(v.k4kEqO3)+1)
	copy(ips, v.k4kEqO3)
	ips = append(ips, ip)
	m.xLikNInZ[id] = &fj057PaXtu{ips, v.eaCklg8jch}
}

// a4bLiQWM:0x1836d00; returns whether the cached IP list excludes this request.
func (m *YM8R9i) a4bLiQWM(id int64, ip string, limit int) bool {
	ips, ok := m.n_augufdnS(id)
	if !ok || len(ips) < limit {
		return false
	}
	if len(ips) > limit {
		return true
	}
	for _, old := range ips {
		if old == ip {
			return false
		}
	}
	return true
}

// lKMA3xU_eZa:0x1836e00. Redis failures return false plus the original error.
func (m *YM8R9i) lKMA3xU_eZa(id int64, ip string, limit int) (bool, error) {
	ips, err := m.xP_MCQ(id)
	if err != nil {
		return false, err
	}
	if len(ips) < limit {
		return false, nil
	}
	if len(ips) > limit {
		return true, nil
	}
	for _, old := range ips {
		if old == ip {
			return false, nil
		}
	}
	return true, nil
}

// xP_MCQ:0x1836f00, singleflight closure0x183e8e0.
func (m *YM8R9i) xP_MCQ(id int64) ([]string, error) {
	if ips, ok := m.n_augufdnS(id); ok {
		return ips, nil
	}
	result, err, _ := m.zHP8cN6.Do(strconv.FormatInt(id, 10), func() (any, error) {
		if ips, ok := m.n_augufdnS(id); ok {
			return ips, nil
		}
		ctx, cancel := context.WithTimeout(m.eHz2FQHNc, 5*time.Second)
		defer cancel()
		ips, err := m.zAD6a7zB.HKeys(ctx, m2pV1HIFtW7(id)).Result()
		if err != nil {
			return nil, err
		}
		L.V8GN5Stk(-4, "user <%v> redis alive ip: %v", id, ips)
		m.nkCoQC(id, ips)
		return ips, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]string), nil
}

// m2pV1HIFtW7:0x1838700.
func m2pV1HIFtW7(id int64) string { return fmt.Sprintf("%s_%d", "soga_conn", id) }
