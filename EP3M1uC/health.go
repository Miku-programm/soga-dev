package EP3M1uC

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"recovered.local/soga/muc279Mpa6S4"
	"recovered.local/soga/uLofn6j"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var healthMutexAt50C7130 sync.Mutex
var healthManagerAt5095C60 *Qi1aOz5oQ

// tCGadgsq restores 0x1b7b040, canceling the previous manager before checking
// whether the new configuration enables health checks. It does not wait for it.
func tCGadgsq(routes *UIfnUsim) {
	healthMutexAt50C7130.Lock()
	defer healthMutexAt50C7130.Unlock()
	if healthManagerAt5095C60 != nil {
		healthManagerAt5095C60.aE1FW1()
		healthManagerAt5095C60 = nil
	}
	if routes == nil || !routes.CnK0a4 || !routes.H839j9DYDD5k {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	manager := &Qi1aOz5oQ{b0Ftnk: routes, bXx0sIP: ctx, aE1FW1: cancel}
	healthManagerAt5095C60 = manager
	go manager.Run()
}

// Run restores 0x1b7b260. The first sweep waits for the five-second ticker.
func (manager *Qi1aOz5oQ) Run() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			manager.qPMo12()
		case <-manager.bXx0sIP.Done():
			return
		}
	}
}

// qPMo12 restores 0x1b7b360 and closure 0x1b8eca0. Only non-direct endpoints
// in rules with multiple outbounds are checked, with pointer deduplication.
func (manager *Qi1aOz5oQ) qPMo12() {
	seen := make(map[*GyyO51]struct{})
	var candidates J_enG0wE
	for _, rule := range manager.b0Ftnk.MMKbp60np {
		if len(rule.UwkgrlFf8) <= 1 {
			continue
		}
		for _, out := range rule.UwkgrlFf8 {
			if out.IsDirect() || out.AHTlIEynXL == "" || out.JO1YIpXk <= 0 {
				continue
			}
			if _, exists := seen[out]; exists {
				continue
			}
			seen[out] = struct{}{}
			candidates = append(candidates, out)
		}
	}
	var wg sync.WaitGroup
	var changed atomic.Bool
	for _, out := range candidates {
		jitter := time.Duration(uLofn6j.TfHpoOcv(int64(2 * time.Second)))
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-manager.bXx0sIP.Done():
				return
			case <-time.NewTimer(jitter).C:
			}
			if manager.nvI5aMbq(out) {
				changed.Store(true)
			}
		}()
	}
	wg.Wait()
	if changed.Load() {
		for _, rule := range manager.b0Ftnk.MMKbp60np {
			rule.eEYGK2LJf()
		}
	}
}

// nvI5aMbq restores 0x1b7b820. Logs are emitted only on transitions.
func (manager *Qi1aOz5oQ) nvI5aMbq(out *GyyO51) bool {
	disabled, enabled := out.ebDTwNX(out.xyiQDHyNr(manager.bXx0sIP))
	if disabled {
		muc279Mpa6S4.V8GN5Stk(slog.LevelInfo, "routes health check disable out: %v", out.l9fABxN())
	}
	if enabled {
		muc279Mpa6S4.V8GN5Stk(slog.LevelInfo, "routes health check enable out: %v", out.l9fABxN())
	}
	return disabled || enabled
}

// xyiQDHyNr restores 0x1b7b9c0: a three-second TCP connect probe. It uses a
// fresh net.Dialer and does not run the configured proxy protocol handshake.
func (out *GyyO51) xyiQDHyNr(ctx context.Context) bool {
	dialer := net.Dialer{Timeout: 3 * time.Second}
	conn, err := dialer.DialContext(ctx, "tcp", net.JoinHostPort(out.AHTlIEynXL, strconv.Itoa(out.JO1YIpXk)))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// l9fABxN restores 0x1b7bb80.
func (out *GyyO51) l9fABxN() string {
	if out.HG8A97UIPf1A != "" {
		return out.HG8A97UIPf1A
	}
	if out.AHTlIEynXL != "" {
		return net.JoinHostPort(out.AHTlIEynXL, strconv.Itoa(out.JO1YIpXk))
	}
	return fmt.Sprintf("%v", out.NyLbiXb4OrlL)
}

// eEYGK2LJf restores 0x1b869e0, atomically publishing a new enabled snapshot.
func (rule *DgAjL9UN) eEYGK2LJf() {
	enabled := rule.UwkgrlFf8.Enabled()
	rule.jf_lZ8wS2GNO.Store(&enabled)
}
