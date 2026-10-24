package inbound

import (
	"context"
	D "recovered.local/soga/buf/pool"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	Log "recovered.local/soga/util/logx"
	"runtime"
)

// wXPmSs2 @0x1d6e660 and worker closure @0x1e6f0a0. The original
// 50C6D18 is runtime.numCPUStartup, written by runtime.osinit @0x4511cf;
// runtime.NumCPU expresses the original inlined read, not GOMAXPROCS.
func (v *Lp07FzBLx9V) wXPmSs2(conn J.KsKg3yls) {
	ctx := v.zb8f3cHWW8v
	queue := make(chan J.AFui66gdB, 8192)
	go receiveShadowsocksDatagrams(ctx, conn, queue)
	worker := func() {
		for packet := range queue {
			v.dR8c1ufH6k(ctx, packet)
		}
	}
	for range max(4, runtime.NumCPU()) {
		go worker()
	}
}

// receiveShadowsocksDatagrams is the source name for reader closure
// 0x1e6f200, cleanup 0x1e6f4a0 and buffer cleanup 0x1e6f460.
func receiveShadowsocksDatagrams(ctx context.Context, conn J.KsKg3yls, queue chan J.AFui66gdB) {
	defer func() {
		_ = conn.Close()
		H.U9tk4Xosh9(conn)
		close(queue)
		Log.V8GN5Stk(-4, "ss udp listener %v close", conn.LocalAddr())
	}()
	buf := D.PoolAt50961D8.Get(4096)
	defer func() { D.PoolAt50961D8.Put(buf) }()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		n, user, source, target, err := conn.ReadFromProxy(buf)
		if err != nil {
			return
		}
		if n == 0 {
			continue
		}
		packet := J.AFui66gdB{Y8YsJ7QnEI3h: conn, AIO3Qb: source, PNliwRWkrU: target, S5eGpRGjp5oo: user, AGHsUHa: buf[:n]}
		select {
		case queue <- packet:
			buf = D.PoolAt50961D8.Get(4096)
		default: // Queue full: drop this datagram and reuse the same buffer.
		}
	}
}
