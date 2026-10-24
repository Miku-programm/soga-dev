package G6bHD5JA

import (
	"golang.org/x/time/rate"
	"recovered.local/soga/muc279Mpa6S4"
	"time"
)

// L337vk2G：原 RTTI 16 字节，Limiter 指针与 int64 速率。
type L337vk2G struct {
	liai0O8     *rate.Limiter
	hgkiIUSUPlQ int64
}

// UpdateLimit @ 0x15b2300。比较原始参数后才裁剪；正值同时作为速率与突发容量。
func (l *L337vk2G) UpdateLimit(limit int64) {
	if l.hgkiIUSUPlQ == limit {
		return
	}
	if limit < 0 {
		limit = 0
	} else if limit > 0x7fffffff {
		limit = 0x7fffffff
	}
	l.hgkiIUSUPlQ = limit
	if limit > 0 {
		if l.liai0O8 != nil {
			limiter := l.liai0O8
			limiter.SetLimit(rate.Limit(limit))
			limiter.SetBurst(int(limit))
		} else {
			l.liai0O8 = rate.NewLimiter(rate.Limit(limit), int(limit))
		}
	} else {
		l.liai0O8 = nil
	}
}

// Add @ 0x15b2460。超出 burst 时原行为为警告并返回，不拆分为多次等待。
func (l *L337vk2G) Add(n int) {
	if l.hgkiIUSUPlQ > 0 && l.liai0O8 != nil {
		limiter := l.liai0O8
		now := time.Now()
		reservation := limiter.ReserveN(now, n)
		if reservation.OK() {
			if delay := reservation.DelayFrom(now); delay != 0 {
				time.Sleep(delay)
			}
		} else {
			muc279Mpa6S4.V8GN5Stk(4, "speed limit <%v> is too small: %v", l.hgkiIUSUPlQ, n)
		}
	}
}
func (l *L337vk2G) AllowAndAdd(n int) bool { // 0x15b2660
	if l.hgkiIUSUPlQ <= 0 || l.liai0O8 == nil {
		return true
	}
	return l.liai0O8.AllowN(time.Now(), n)
}

// GetSpeedLimit @ 0x15b2720，原浮点运算顺序保留。
func (l *L337vk2G) GetSpeedLimit() float64 {
	return float64(l.hgkiIUSUPlQ) * 0.0009765625 * 8 * 0.0009765625
}
