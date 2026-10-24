package inbound

import (
	"context"
	D "recovered.local/soga/buf/pool"
	H "recovered.local/soga/util/errutil"
	J "recovered.local/soga/net/sess"
	S "recovered.local/soga/protocol/ssr"
	C "recovered.local/soga/config"
	Log "recovered.local/soga/util/logx"
	U "recovered.local/soga/user"
	"runtime"
)

// 1d7eea0、producer1e70c00/consumer1e70aa0。非 mode0 使用8192队列和
// max(4,numCPUStartup)消费者；原 mode0 只使用64队列和一个消费者。
func (v *Ze0fOaAko) e5tsZB2Ru_2v(ctx context.Context, pc *S.AXWGpkYq, u *U.LJJkaD) {
	capacity := 8192
	if u.GetSinglePortType() == 0 {
		capacity = 64
	}
	queue := make(chan J.AFui66gdB, capacity)
	go receiveSSRDatagrams(ctx, pc, queue, u)
	worker := func() {
		for p := range queue {
			v.dR8c1ufH6k(ctx, p)
		}
	}
	if u.GetSinglePortType() == 0 {
		go worker()
	} else {
		for range max(4, runtime.NumCPU()) {
			go worker()
		}
	}
}

// producer 1e70c00；新增源名称仅表达原闭包，不额外改变原错误重试行为。
func receiveSSRDatagrams(ctx context.Context, pc *S.AXWGpkYq, queue chan J.AFui66gdB, u *U.LJJkaD) {
	defer func() {
		H.SlG_Y_mg("handle ssr udp")
		_ = pc.Close()
		H.PpnNLr_pf7He(pc)
		Log.TRTbrquD(-4, "close ssr udp listener")
		close(queue)
	}()
	enabled, _ := C.TunnelSettings()
	tunnel := enabled && !C.Q0SDIN()
	buf := D.PoolAt50961D8.Get(4096)
	defer func() { D.PoolAt50961D8.Put(buf) }()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		n, user, client, target, err := pc.ReadFromProxy(buf)
		if err != nil {
			Log.TRTbrquD(-4, "read ssr udp err:", err)
			if tunnel {
				return
			}
			continue
		}
		if n == 0 {
			continue
		}
		if user == nil {
			user = u
		}
		packet := J.AFui66gdB{Y8YsJ7QnEI3h: pc, AIO3Qb: client, PNliwRWkrU: target, S5eGpRGjp5oo: user, AGHsUHa: buf[:n]}
		select {
		case queue <- packet:
			buf = D.PoolAt50961D8.Get(4096)
		default:
		}
	}
}
