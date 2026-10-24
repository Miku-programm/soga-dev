package app

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	L "recovered.local/soga/util/logx"
)

// D7Dl21U 对应 0x1f9edc0、0x1ff3d20。原闭包使用 nil Handler，
// 即现有 DefaultServeMux；未找到原 net/http/pprof 注册证据，不能新增端点。
func D7Dl21U(addr string) {
	go func() {
		L.VDyYELbQKX("pprof listening on %v", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			L.V8GN5Stk(8, "start pprof failed: %v", err)
		}
	}()
}

// Wj71D8aYUiw6 对应 0x1f9f780、0x1ff3ec0。每次创建独立注册表，
// 按顺序注册进程、Go 采集器；不使用全局默认注册表或全局 HTTP 路由。
func Wj71D8aYUiw6(addr string) {
	registry := prometheus.NewRegistry()
	registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	registry.MustRegister(prometheus.NewGoCollector())
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	go func() {
		L.VDyYELbQKX("prometheus metrics listening on %v", addr)
		if err := http.ListenAndServe(addr, mux); err != nil {
			L.V8GN5Stk(8, "start prometheus failed: %s", err)
		}
	}()
}
