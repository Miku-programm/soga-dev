package limit

import (
	"context"
	H "recovered.local/soga/util/errutil"
	L "recovered.local/soga/util/logx"
	N "recovered.local/soga/user"
)

// StartUserConn:0x1836760; asynchronous validation closure0x183e3c0.
func (m *YM8R9i) StartUserConn(ctx context.Context, cancel context.CancelFunc, run func(), user N.PgaNd3014U, ip string) bool {
	limit := user.GetConnLimit()
	if limit <= 0 {
		run()
		return true
	}
	id := user.GetId()
	key := uuMFf5B{id, ip}
	m.p81VkaYJai.Lock()
	item, ok := m.pCsxt7t[key]
	if ok && item.aXnE_3N_gdt {
		item.a7fxke++
		m.p81VkaYJai.Unlock()
		run()
		return true
	}
	if m.a4bLiQWM(id, ip, limit) {
		m.p81VkaYJai.Unlock()
		return false
	}
	if !ok {
		item = &e36g5GkEd{}
		m.pCsxt7t[key] = item
	}
	item.a7fxke++
	paused := m.cmA3rUr.Load()
	if paused {
		item.aXnE_3N_gdt = true
	}
	m.p81VkaYJai.Unlock()
	if paused {
		run()
		return true
	}
	go func() {
		defer H.SlG_Y_mg("redis ip limit")
		select {
		case <-ctx.Done():
			return
		default:
		}
		exceeded, err := m.lKMA3xU_eZa(id, ip, limit)
		if err != nil {
			L.V8GN5Stk(-4, "get user <%v> alive ip err: %v", id, err)
		}
		if exceeded {
			cancel()
			return
		}
		m.p81VkaYJai.Lock()
		item, ok := m.pCsxt7t[key]
		if !ok {
			m.p81VkaYJai.Unlock()
			return
		}
		registered := item.aXnE_3N_gdt
		item.aXnE_3N_gdt = true
		m.p81VkaYJai.Unlock()
		if !registered {
			if err := m.uPDNclE9W_k(id, ip); err != nil {
				L.V8GN5Stk(4, "register user <%v> alive ip err: %v", id, err)
			}
			m.ba5WPJg(id, ip)
		}
		run()
	}()
	return true
}

// xtMqCa5:0x18380a0. Snapshot only approved entries with positive refcounts.
func (m *YM8R9i) xtMqCa5() map[int64][]string {
	m.p81VkaYJai.Lock()
	defer m.p81VkaYJai.Unlock()
	out := make(map[int64][]string, len(m.pCsxt7t))
	for key, item := range m.pCsxt7t {
		if item.aXnE_3N_gdt && item.a7fxke > 0 {
			out[key.dCzJC4or4Xl] = append(out[key.dCzJC4or4Xl], key.gDlwyqlt)
		}
	}
	return out
}
