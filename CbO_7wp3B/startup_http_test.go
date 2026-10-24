package CbO_7wp3B

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// 原启动入口不返回服务句柄，也不注册取消回调。用隔离测试进程保证
// 实际监听行为可验证，且不会给测试主进程遗留不可关闭的服务。
func TestOriginalAuxiliaryHTTPStartup(t *testing.T) {
	if os.Getenv("SOGA_HTTP_STARTUP_FIXTURE") != "1" {
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestOriginalAuxiliaryHTTPStartup$", "-test.count=1")
		cmd.Env = append(os.Environ(), "SOGA_HTTP_STARTUP_FIXTURE=1")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("HTTP startup: %v\n%s", err, out)
		}
		return
	}
	reserve := func() string {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		a := l.Addr().String()
		l.Close()
		return a
	}
	metrics, diagnostic := reserve(), reserve()
	prometheus.MustRegister(prometheus.NewGauge(prometheus.GaugeOpts{Name: "must_not_be_in_native_registry", Help: "isolated fixture"}))
	http.DefaultServeMux.HandleFunc("/fixture-default", func(w http.ResponseWriter, _ *http.Request) { io.WriteString(w, "existing-default-mux") })
	Wj71D8aYUiw6(metrics)
	D7Dl21U(diagnostic)
	client := &http.Client{Timeout: time.Second}
	fetch := func(addr, path string) (int, string) {
		deadline := time.Now().Add(3 * time.Second)
		for {
			r, err := client.Get("http://" + addr + path)
			if err == nil {
				b, err := io.ReadAll(r.Body)
				r.Body.Close()
				if err != nil {
					t.Fatal(err)
				}
				return r.StatusCode, string(b)
			}
			if time.Now().After(deadline) {
				t.Fatal(err)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}
	code, body := fetch(metrics, "/metrics")
	if code != 200 || !strings.Contains(body, "go_goroutines ") || !strings.Contains(body, "process_cpu_seconds_total ") || strings.Contains(body, "must_not_be_in_native_registry") {
		t.Fatalf("metrics: %d %s", code, body)
	}
	if code, _ := fetch(metrics, "/fixture-default"); code != 404 {
		t.Fatalf("metrics leaked default mux: %d", code)
	}
	if code, body := fetch(diagnostic, "/fixture-default"); code != 200 || body != "existing-default-mux" {
		t.Fatalf("nil handler: %d %s", code, body)
	}
	if code, _ := fetch(diagnostic, "/debug/pprof/"); code != 404 {
		t.Fatalf("unproved pprof handler added: %d", code)
	}
}
