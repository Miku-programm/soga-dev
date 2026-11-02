package limit

import (
	"context"
	A "recovered.local/soga/util/rtx"
	C "recovered.local/soga/config"
	L "recovered.local/soga/util/logx"
	N "recovered.local/soga/user"
	"sync/atomic"
	"time"
)

// XKYW8YhsbqJ: original RTTI 0x36df068, 56 bytes.
type XKYW8YhsbqJ struct {
	a119HAhPh7oZ context.Context
	a9ebqRa0Lr   *N.ZpNFqUVZp
	rbjjjB5m6T   *C.HwIE6xOrX1OH
	nxmaanKl     *HRyRearcsgu
	xBTUM1U2bEY  *ZWHtxl
	f9fwna9vs    atomic.Bool
}

// DMIy6PEVB: 0x1830f80. Snapshot the configuration pointer for window behavior;
// the enabled flag is nevertheless read from the shared pointer on later ticks.
func DMIy6PEVB(ctx context.Context, users *N.ZpNFqUVZp) *XKYW8YhsbqJ {
	config := C.DynamicConfiguration()
	traffic := b3D2x6m1(ctx, int(config.Vs46JphL/(5*time.Second)))
	speeds := uFaljJ(ctx, users)
	m := &XKYW8YhsbqJ{a119HAhPh7oZ: ctx, a9ebqRa0Lr: users, rbjjjB5m6T: config, nxmaanKl: traffic, xBTUM1U2bEY: speeds}
	m.init()
	return m
}

// init: 0x18310e0. No worker is started if disabled at construction.
func (m *XKYW8YhsbqJ) init() {
	if C.DynamicConfiguration().DiN7MxyB7 {
		go m.dTRa2vHe5()
	}
}

// dTRa2vHe5: 0x1831180.
func (m *XKYW8YhsbqJ) dTRa2vHe5() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if C.DynamicConfiguration().DiN7MxyB7 {
				m._jEz0Fre9w()
			}
		case <-m.a119HAhPh7oZ.Done():
			return
		}
	}
}

// _jEz0Fre9w: 0x18312a0. Out of the configured time ranges only clears active;
// it does not drain counters or remove outstanding cached speed limits.
func (m *XKYW8YhsbqJ) _jEz0Fre9w() {
	if m.rbjjjB5m6T.BgrTZk.IsInTime(time.Now().In(A.TimeLocationAt5095CB0)) {
		m.f9fwna9vs.Store(true)
		m.a9ebqRa0Lr.Range(func(_ int64, user N.PgaNd3014U) bool { m._5fobl7OClw4(user); return true })
	} else {
		m.f9fwna9vs.Store(false)
	}
}

// _5fobl7OClw4: 0x18313a0. Counter draining precedes the whitelist and speed
// checks. The inverted comparison also preserves the original unordered-NaN
// branches (ucomisd/jbe and ucomisd/jb), unlike a positive disjunction in Go.
func (m *XKYW8YhsbqJ) _5fobl7OClw4(user N.PgaNd3014U) {
	up, down := user.GetAndClearDyTraffic()
	id := user.GetId()
	config := m.rbjjjB5m6T
	if config.GK1kRuhBCq.Has(id) {
		return
	}
	speed := user.GetSpeedLimit()
	if speed > 0 && config.M6lj3M4NlG >= speed {
		return
	}
	total := m.nxmaanKl.AddUserTraffic(id, up+down)
	if total != 0 && !m.xBTUM1U2bEY.HasUser(id) && total >= int64(config.Vs46JphL/time.Second)*int64(1024*(0.125*(config.Yt1TMC2*1024))) {
		limit := config.M6lj3M4NlG
		L.V8GN5Stk(-4, "user %v dynamic speed limit: %.2f Mbps -> %.2f Mbps", id, speed, limit)
		user.SetSpeedLimit(limit)
		m.xBTUM1U2bEY.SetUser(user, limit, config.ZxdryEoK)
	}
}

// GetUserSpeedLimit: 0x1831700.
func (m *XKYW8YhsbqJ) GetUserSpeedLimit(id int64) (float64, bool) {
	if C.DynamicConfiguration().DiN7MxyB7 && m.f9fwna9vs.Load() && !m.rbjjjB5m6T.GK1kRuhBCq.Has(id) {
		return m.xBTUM1U2bEY.GetUserSpeedLimit(id)
	}
	return 0, false
}
