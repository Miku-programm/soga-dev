//go:build linux

package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

// This executes a freshly built cmd/soga, never the supplied ELF. The fixture
// lives inside user/network/mount namespaces. Its network contains only lo;
// /etc is a bind-mounted temporary tree because original main initializes geo
// data at hard-coded paths before reading the configurable file flags.
func TestRebuiltMainSogaV1PanelAndNativeTCPUDP(t *testing.T) {
	if os.Getenv("SOGA_MAIN_SOGA_INTEGRATION_ROLE") == "fixture" {
		runMainSogaV1NetworkFixture(t)
		return
	}
	if testing.Short() {
		t.Skip("builds real application and uses Linux namespaces")
	}
	for _, tool := range []string{"unshare", "mount", "ip"} {
		if _, e := exec.LookPath(tool); e != nil {
			t.Skipf("namespace integration unavailable: %s", tool)
		}
	}
	if out, e := exec.Command("unshare", "-Urnm", "true").CombinedOutput(); e != nil {
		t.Skipf("namespace permission unavailable: %v %s", e, out)
	}
	dir := t.TempDir()
	etc := filepath.Join(dir, "etc")
	if e := os.MkdirAll(filepath.Join(etc, "soga"), 0700); e != nil {
		t.Fatal(e)
	}
	for name, data := range map[string]string{"hosts": "127.0.0.1 localhost\n::1 localhost\n", "resolv.conf": "nameserver 127.0.0.1\noptions timeout:1 attempts:1\n", "nsswitch.conf": "hosts: files dns\n"} {
		if e := os.WriteFile(filepath.Join(etc, name), []byte(data), 0600); e != nil {
			t.Fatal(e)
		}
	}
	module, e := filepath.Abs("../..")
	if e != nil {
		t.Fatal(e)
	}
	binaryPath := filepath.Join(dir, "soga-rebuilt")
	goBinary := filepath.Join(runtime.GOROOT(), "bin", "go")
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	build := exec.CommandContext(ctx, goBinary, "run", "./buildhttp", "-go", goBinary, "build", "-o", binaryPath, "./cmd/soga")
	build.Dir = module
	if out, e := build.CombinedOutput(); e != nil {
		t.Fatalf("build source application: %v\n%s", e, out)
	}
	compiled, e := os.ReadFile(binaryPath)
	if e != nil {
		t.Fatal(e)
	}
	t.Logf("rebuilt cmd/soga SHA256: %x", sha256.Sum256(compiled))
	script := `set -eu
mount --make-rprivate /
mount --bind "$1" /etc
ip link set lo up
exec "$2" -test.run=^TestRebuiltMainSogaV1PanelAndNativeTCPUDP$ -test.v -test.timeout=90s`
	child := exec.CommandContext(ctx, "unshare", "-Urnm", "sh", "-c", script, "fixture", etc, os.Args[0])
	child.Env = append(os.Environ(), "SOGA_MAIN_SOGA_INTEGRATION_ROLE=fixture", "SOGA_MAIN_SOGA_INTEGRATION_BINARY="+binaryPath, "SOGA_MAIN_SOGA_INTEGRATION_DIR="+dir)
	out, e := child.CombinedOutput()
	t.Logf("isolated application fixture:\n%s", out)
	if e != nil {
		t.Fatalf("isolated main: %v", e)
	}
}

func runMainSogaV1NetworkFixture(t *testing.T) {
	interfaces, e := net.Interfaces()
	if e != nil {
		t.Fatal(e)
	}
	for _, i := range interfaces {
		if i.Name != "lo" {
			t.Fatalf("unexpected external interface %s", i.Name)
		}
	}
	routes, e := os.ReadFile("/proc/net/route")
	if e != nil || len(strings.Fields(string(routes))) != 11 {
		t.Fatalf("unexpected routes: %s (%v)", routes, e)
	}
	// Answer only in this namespace. NXDOMAIN makes original mandatory NTP
	// resolution fail promptly; no NTP clock or license check is replaced.
	dns, e := net.ListenPacket("udp4", "127.0.0.1:53")
	if e != nil {
		t.Fatal(e)
	}
	defer dns.Close()
	go func() {
		b := make([]byte, 2048)
		for {
			n, a, e := dns.ReadFrom(b)
			if e != nil {
				return
			}
			if n < 12 {
				continue
			}
			p := append([]byte{}, b[:n]...)
			p[2] = 0x81
			p[3] = 0x83
			for i := 6; i < 12; i++ {
				p[i] = 0
			}
			dns.WriteTo(p, a)
		}
	}()
	tcp, e := net.Listen("tcp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer tcp.Close()
	go func() {
		for {
			c, e := tcp.Accept()
			if e != nil {
				return
			}
			go func() { defer c.Close(); c.SetDeadline(time.Now().Add(5 * time.Second)); io.Copy(c, c) }()
		}
	}()
	udp, e := net.ListenPacket("udp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer udp.Close()
	go func() {
		b := make([]byte, 65535)
		for {
			n, a, e := udp.ReadFrom(b)
			if e != nil {
				return
			}
			udp.WriteTo(b[:n], a)
		}
	}()
	port1 := mainFixturePort(t)
	const password = "01234567-89ab-cdef-0123-456789abcdef"
	var userVersion atomic.Int64
	var configGets, userGets, auditGets, whiteGets, pushes, statusPushes, trafficUp, trafficDown atomic.Int64
	var panelError atomic.Pointer[string]
	panel := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("API-KEY") != "local-key-one" || r.Header.Get("NODE-ID") != "71" || r.Header.Get("NODE-TYPE") != "shadowsocks" || r.URL.RawQuery != "" {
			v := fmt.Sprintf("headers=%v url=%s", r.Header, r.URL.String())
			panelError.Store(&v)
			http.Error(w, "wrong panel authentication", 400)
			return
		}
		path := strings.TrimPrefix(r.URL.Path, "/api/v1/")
		if path == r.URL.Path {
			v := r.URL.Path
			panelError.Store(&v)
			http.NotFound(w, r)
			return
		}
		switch path {
		case "node":
			if r.Method != "GET" {
				v := "node method " + r.Method
				panelError.Store(&v)
			}
			configGets.Add(1)
			fmt.Fprintf(w, `{"basic":{"pull_interval":1,"push_interval":1},"config":{"port":%d,"cipher":"aes-128-gcm","stream_type":"tcp"}}`, port1)
		case "users":
			userGets.Add(1)
			credential := password
			if userVersion.Load() != 0 {
				credential = "replacement-user-password"
			}
			fmt.Fprintf(w, `[{"id":13,"password":%q,"speed_limit":0,"device_limit":0,"tcp_limit":0}]`, credential)
		case "audit_rules":
			auditGets.Add(1)
			fmt.Fprint(w, `[]`)
		case "white_list":
			whiteGets.Add(1)
			fmt.Fprint(w, `[]`)
		case "traffic":
			pushes.Add(1)
			var entries []struct {
				ID int64 `json:"id"`
				U  int64 `json:"u"`
				D  int64 `json:"d"`
			}
			if e := json.NewDecoder(r.Body).Decode(&entries); e != nil {
				v := e.Error()
				panelError.Store(&v)
			}
			for _, entry := range entries {
				if entry.ID != 13 {
					v := fmt.Sprintf("wrong traffic user %d", entry.ID)
					panelError.Store(&v)
				}
				trafficUp.Add(entry.U)
				trafficDown.Add(entry.D)
			}
			fmt.Fprint(w, `{"code":0,"message":"ok"}`)
		case "status":
			statusPushes.Add(1)
			var status map[string]json.RawMessage
			if e := json.NewDecoder(r.Body).Decode(&status); e != nil {
				v := e.Error()
				panelError.Store(&v)
			}
			for _, key := range []string{"cpu", "mem", "swap", "disk", "uptime"} {
				if _, ok := status[key]; !ok {
					v := "missing status " + key
					panelError.Store(&v)
				}
			}
			fmt.Fprint(w, `{"code":0,"message":"ok"}`)
		case "alive_ip", "audit_log":
			fmt.Fprint(w, `{"code":0,"message":"ok"}`)
		default:
			v := r.URL.Path
			panelError.Store(&v)
			http.NotFound(w, r)
		}
	}))
	defer panel.Close()
	dir := os.Getenv("SOGA_MAIN_SOGA_INTEGRATION_DIR")
	config := filepath.Join(dir, "soga.conf")
	writeConfig := func(id int, key string) {
		text := fmt.Sprintf("type=soga-v1\nserver_type=ss\napi=webapi\nwebapi_url=%s\nwebapi_key=%s\nnode_id=%d\nlisten=127.0.0.1\ncheck_interval=1\nsubmit_interval=1\nrandom_interval_range=0\nauto_update=false\ngeo_update_enable=false\naccess_log_enable=false\nlog_file_dir=\nlog_level=debug\nproxy_protocol=false\nudp_proxy_protocol=false\ndy_limit_enable=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=2\nudp_timeout=2\n", panel.URL, key, id)
		if e := os.WriteFile(config, []byte(text), 0600); e != nil {
			t.Fatal(e)
		}
	}
	writeConfig(71, "local-key-one")
	args := []string{"-c", config, "-soga_key", ""}
	for _, f := range []string{"b", "w", "d", "r"} {
		p := filepath.Join(dir, f)
		if e := os.WriteFile(p, nil, 0600); e != nil {
			t.Fatal(e)
		}
		args = append(args, "-"+f, p)
	}
	logPath := filepath.Join(dir, "application.log")
	log, e := os.Create(logPath)
	if e != nil {
		t.Fatal(e)
	}
	defer log.Close()
	app := exec.Command(os.Getenv("SOGA_MAIN_SOGA_INTEGRATION_BINARY"), args...)
	app.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGKILL}
	app.Stdout = log
	app.Stderr = log
	app.Env = append(os.Environ(), "HTTP_PROXY=", "HTTPS_PROXY=", "ALL_PROXY=", "http_proxy=", "https_proxy=", "all_proxy=", "NO_PROXY=*", "no_proxy=*")
	if e = app.Start(); e != nil {
		t.Fatal(e)
	}
	defer func() { _ = app.Process.Kill() }()
	exited := make(chan error, 1)
	go func() { exited <- app.Wait() }()
	defer func() {
		b, _ := os.ReadFile(logPath)
		if t.Failed() {
			t.Logf("real application log:\n%s", b)
		}
	}()
	wait := func(label string, condition func() bool) {
		t.Helper()
		deadline := time.Now().Add(20 * time.Second)
		for {
			if condition() {
				return
			}
			select {
			case e := <-exited:
				t.Fatalf("application exited during %s: %v", label, e)
			default:
			}
			if time.Now().After(deadline) {
				t.Fatalf("timeout: %s", label)
			}
			time.Sleep(50 * time.Millisecond)
		}
	}
	tcpTarget, udpTarget := tcp.Addr().(*net.TCPAddr).Port, udp.LocalAddr().(*net.UDPAddr).Port
	var last error
	wait("first real TCP relay", func() bool { last = mainFixtureRoundtrip("tcp", port1, tcpTarget, password); return last == nil })
	if e = mainFixtureRoundtrip("udp", port1, udpTarget, password); e != nil {
		t.Fatal("first UDP", e)
	}
	wait("traffic submission", func() bool { return trafficUp.Load() > 0 && trafficDown.Load() > 0 })
	userVersion.Store(1)
	wait("live panel user replacement", func() bool {
		last = mainFixtureRoundtrip("tcp", port1, tcpTarget, "replacement-user-password")
		return last == nil
	})
	if e = mainFixtureRoundtrip("udp", port1, udpTarget, "replacement-user-password"); e != nil {
		t.Fatal("updated UDP", e)
	}
	if e = mainFixtureRoundtrip("tcp", port1, tcpTarget, password); e == nil {
		t.Fatal("old credential remained accepted after panel replacement")
	}
	if e = mainFixtureRoundtrip("udp", port1, udpTarget, password); e == nil {
		t.Fatal("old UDP credential remained accepted after panel replacement")
	}
	if configGets.Load() == 0 || userGets.Load() < 2 || auditGets.Load() == 0 || whiteGets.Load() == 0 || pushes.Load() == 0 || statusPushes.Load() == 0 {
		t.Fatal("missing soga-v1 API lifecycle", configGets.Load(), userGets.Load(), auditGets.Load(), whiteGets.Load(), pushes.Load(), statusPushes.Load())
	}

	if p := panelError.Load(); p != nil {
		t.Fatal("panel", *p)
	}
	if e = app.Process.Signal(syscall.SIGTERM); e != nil {
		t.Fatal(e)
	}
	select {
	case e = <-exited:
		if e != nil {
			t.Fatal("TERM exit", e)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("TERM timeout")
	}
	if c, e := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(port1)), 100*time.Millisecond); e == nil {
		c.Close()
		t.Fatal("listener survived process exit")
	}
	appLog, e := os.ReadFile(logPath)
	if e != nil {
		t.Fatal(e)
	}
	if strings.Count(string(appLog), "load config <"+config+"> success") < 1 {
		t.Fatal("missing real initial configuration log", string(appLog))
	}
	t.Logf("real soga-v1 main verified: node=%d users=%d audit_rules=%d white_list=%d traffic=%d status=%d upload=%d download=%d; TCP+UDP before/after password update; old TCP+UDP rejected; TERM exit0; only lo", configGets.Load(), userGets.Load(), auditGets.Load(), whiteGets.Load(), pushes.Load(), statusPushes.Load(), trafficUp.Load(), trafficDown.Load())
}
