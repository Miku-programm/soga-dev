package limit

import (
	"context"
	A "recovered.local/soga/util/rtx"
	L "recovered.local/soga/util/logx"
	"time"
)

// vih7o_:0x1836040. A ping error pauses enforcement; unsupported HSETEX only
// logs the original error and does not change the pause flag.
func (m *YM8R9i) vih7o_() {
	ctx, cancel := context.WithTimeout(m.eHz2FQHNc, 5*time.Second)
	defer cancel()
	if err := m.zAD6a7zB.Ping(ctx).Err(); err != nil {
		m.cmA3rUr.Store(true)
		L.V8GN5Stk(4, "redis ip limit connect err: %v, pause ip limit", err)
		return
	}
	if err := m.zAD6a7zB.Do(ctx, "HSETEX", "soga_conn_hsetex_probe", "EX", int64(1), "FIELDS", 1, "probe", "0").Err(); err != nil {
		L.V8GN5Stk(8, "redis ip limit needs redis 8.0+: %v", err)
	}
}

// fV5Dht3hwUlM:0x1836600. All jobs wait for their first tick.
func (m *YM8R9i) fV5Dht3hwUlM() {
	A.SwUrlM3M(m.eHz2FQHNc, func() {
		ctx, cancel := context.WithTimeout(m.eHz2FQHNc, 3*time.Second)
		defer cancel()
		if err := m.zAD6a7zB.Ping(ctx).Err(); err != nil {
			if m.cmA3rUr.CompareAndSwap(false, true) {
				L.V8GN5Stk(4, "redis ip limit connect err: %v, pause ip limit", err)
			}
		} else if m.cmA3rUr.CompareAndSwap(true, false) {
			L.TRTbrquD(0, "redis ip limit connect success, resume ip limit")
		}
	}, 5*time.Second)
	A.SwUrlM3M(m.eHz2FQHNc, func() {
		if !m.cmA3rUr.Load() {
			m.nZcet5T6783()
		}
	}, m.p7KF5sbP/2)
	A.SwUrlM3M(m.eHz2FQHNc, func() {
		now := time.Now()
		m.nttEJ6.Lock()
		for id, v := range m.xLikNInZ {
			if now.After(v.eaCklg8jch) {
				delete(m.xLikNInZ, id)
			}
		}
		m.nttEJ6.Unlock()
	}, m.p7KF5sbP)
}

// uPDNclE9W_k:0x1837680. Each IP is a separately expiring Redis hash field.
func (m *YM8R9i) uPDNclE9W_k(id int64, ip string) error {
	ctx, cancel := context.WithTimeout(m.eHz2FQHNc, 5*time.Second)
	defer cancel()
	seconds := int64(m.p7KF5sbP / time.Second)
	if seconds <= 0 {
		seconds = 1
	}
	return m.zAD6a7zB.Do(ctx, "HSETEX", m2pV1HIFtW7(id), "EX", seconds, "FIELDS", 1, ip, "0").Err()
}

// nZcet5T6783:0x1837a60. Batch at most500 users per pipeline, reuse slice storage.
func (m *YM8R9i) nZcet5T6783() {
	active := m.xtMqCa5()
	if len(active) == 0 {
		return
	}
	batch := make([][]any, 0, min(len(active), 500))
	total := 0
	for id, ips := range active {
		seconds := int64(m.p7KF5sbP / time.Second)
		if seconds <= 0 {
			seconds = 1
		}
		args := make([]any, 0, 6+2*len(ips))
		args = append(args, "HSETEX", m2pV1HIFtW7(id), "EX", seconds, "FIELDS", len(ips))
		for _, ip := range ips {
			args = append(args, ip, "0")
		}
		batch = append(batch, args)
		total += len(ips)
		if len(batch) >= 500 {
			m.fZyGGLn(batch)
			batch = batch[:0]
		}
	}
	m.fZyGGLn(batch)
	L.V8GN5Stk(-4, "redis ip limit renewed %v ips for %v users", total, len(active))
}

// fZyGGLn:0x18382e0. The original execution callback reaches the ordinary
// pipeline writer (0xf7b140), without MULTI/EXEC transaction framing.
func (m *YM8R9i) fZyGGLn(batch [][]any) {
	if len(batch) == 0 {
		return
	}
	ctx, cancel := context.WithTimeout(m.eHz2FQHNc, 5*time.Second)
	defer cancel()
	pipe := m.zAD6a7zB.Pipeline()
	defer pipe.Close()
	for _, args := range batch {
		pipe.Do(ctx, args...)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		L.TRTbrquD(4, "redis ip limit renew err:", err)
	}
}
