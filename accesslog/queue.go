package accesslog

import (
	C "recovered.local/soga/config"
	L "recovered.local/soga/util/logx"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 地址命名由分析者指定；分别对应 0x5096948、0x50c7d00、0x50c79d8。
// 原 init@0x185a100 建立容量 8192 的记录队列。
var queueAt5096948 = make(chan cUwNbDJ6opWy, 8192)
var onceAt50C7D00 sync.Once
var droppedAt50C79D8 atomic.Int64

// IPblyWsS: 0x185a160，func1@0x1872300 懒启动后台写入协程。
// access 为客户端原始请求的目标主机（sniff 前），host 为 sniff 后的主机。
func IPblyWsS(network string, node int, user int64, source, access, host string, port int, target string) {
	if !C.AccessLogEnabled() {
		return
	}
	onceAt50C7D00.Do(func() { go q1GuXw0() })
	r := cUwNbDJ6opWy{time.Now(), network, node, user, source, host, target, port, access}
	select {
	case queueAt5096948 <- r:
	default:
		droppedAt50C79D8.Add(1)
	}
}

// lacwk7cXz: 0x185a680。构造时选择后端，原代码没有运行中切换后端。
func lacwk7cXz() c5PkGn {
	mode, _, _ := C.AccessLogRetentionConfiguration()
	if mode == "clickhouse" {
		return z12Z6AW7z()
	}
	return &cToP4Y8{}
}

// q1GuXw0: 0x185a3a0。func1@0x1872260 直接 recover；原循环无退出请求，
// 队列满时丢弃新记录，累计数量在每秒刷新时报告。
func q1GuXw0() {
	defer func() {
		if p := recover(); p != nil {
			L.V8GN5Stk(8, "access log writer panic: %v", p)
		}
	}()
	writer := lacwk7cXz()
	defer writer.t_DddcBjF7()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var b strings.Builder
	for {
		select {
		case r := <-queueAt5096948:
			writer.ajkzPuuQ(&b, &r)
		case <-ticker.C:
			writer.eyAgzSwv7d()
			if n := droppedAt50C79D8.Swap(0); n > 0 {
				L.V8GN5Stk(4, "access log queue full, dropped %v records", n)
			}
		}
	}
}
