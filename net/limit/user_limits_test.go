package limit

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	R "recovered.local/soga/net/redis"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"
)

// A real TCP RESP fixture verifies wire requests independently of redis.Client.
// It is intentionally not presented as integration with a Redis 8 server.
type redisWireFixture struct {
	mu          sync.Mutex
	ln          net.Listener
	connections map[net.Conn]bool
	commands    [][]string
	hashes      map[string]map[string]string
	failures    map[string]string
	wg          sync.WaitGroup
}

func redisFixture(t *testing.T) *redisWireFixture {
	t.Helper()
	ln, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	f := &redisWireFixture{ln: ln, connections: map[net.Conn]bool{}, hashes: map[string]map[string]string{}, failures: map[string]string{}}
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		for {
			c, e := ln.Accept()
			if e != nil {
				return
			}
			f.mu.Lock()
			f.connections[c] = true
			f.mu.Unlock()
			f.wg.Add(1)
			go f.serve(c)
		}
	}()
	t.Cleanup(func() {
		ln.Close()
		f.mu.Lock()
		for c := range f.connections {
			c.Close()
		}
		f.mu.Unlock()
		f.wg.Wait()
	})
	return f
}
func (f *redisWireFixture) serve(c net.Conn) {
	defer f.wg.Done()
	defer c.Close()
	r := bufio.NewReader(c)
	for {
		line, e := r.ReadString('\n')
		if e != nil {
			return
		}
		if len(line) < 3 || line[0] != '*' {
			return
		}
		n, e := strconv.Atoi(strings.TrimSpace(line[1:]))
		if e != nil {
			return
		}
		a := make([]string, n)
		for i := range a {
			line, e = r.ReadString('\n')
			if e != nil || len(line) < 3 || line[0] != '$' {
				return
			}
			sz, e := strconv.Atoi(strings.TrimSpace(line[1:]))
			if e != nil || sz < 0 {
				return
			}
			b := make([]byte, sz+2)
			if _, e = io.ReadFull(r, b); e != nil {
				return
			}
			a[i] = string(b[:sz])
		}
		f.mu.Lock()
		f.commands = append(f.commands, a)
		command := strings.ToUpper(a[0])
		var out string
		if failure := f.failures[command]; failure != "" {
			out = "-ERR " + failure + "\r\n"
		} else {
			switch command {
			case "PING":
				out = "+PONG\r\n"
			case "HKEYS":
				h := f.hashes[a[1]]
				out = fmt.Sprintf("*%d\r\n", len(h))
				for k := range h {
					out += fmt.Sprintf("$%d\r\n%s\r\n", len(k), k)
				}
			case "HSETEX":
				if len(a) < 8 || a[2] != "EX" || a[4] != "FIELDS" {
					out = "-ERR bad HSETEX\r\n"
					break
				}
				count, e := strconv.Atoi(a[5])
				if e != nil || len(a) != 6+count*2 {
					out = "-ERR field count\r\n"
					break
				}
				if f.hashes[a[1]] == nil {
					f.hashes[a[1]] = map[string]string{}
				}
				for i := 6; i < len(a); i += 2 {
					f.hashes[a[1]][a[i]] = a[i+1]
				}
				out = fmt.Sprintf(":%d\r\n", count)
			default:
				out = "-ERR unexpected command\r\n"
			}
		}
		f.mu.Unlock()
		if _, e = io.WriteString(c, out); e != nil {
			return
		}
	}
}
func (f *redisWireFixture) manager(t *testing.T) *YM8R9i {
	t.Helper()
	client := R.ZHNMa9Nz(&R.Csx1an{Addr: f.ln.Addr().String(), MaxRetries: -1})
	t.Cleanup(func() { client.Close() })
	return &YM8R9i{eHz2FQHNc: context.Background(), zAD6a7zB: client, p7KF5sbP: 60 * time.Second, pCsxt7t: map[uuMFf5B]*e36g5GkEd{}, xLikNInZ: map[int64]*fj057PaXtu{}}
}
func (f *redisWireFixture) snapshot() [][]string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([][]string(nil), f.commands...)
}

type connectionTestUser struct {
	dynamicTestUser
	limit int
}

func (u *connectionTestUser) GetConnLimit() int { return u.limit }
func waitCall(t *testing.T, ch <-chan struct{}) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Fatal("callback timeout")
	}
}
func TestRedisProbeAndConnectionWire(t *testing.T) {
	f := redisFixture(t)
	m := f.manager(t)
	if unsafe.Sizeof(*m) != 104 || unsafe.Sizeof(uuMFf5B{}) != 24 || unsafe.Sizeof(fj057PaXtu{}) != 48 {
		t.Fatal("RTTI layout")
	}
	m.vih7o_()
	want := [][]string{{"ping"}, {"HSETEX", "soga_conn_hsetex_probe", "EX", "1", "FIELDS", "1", "probe", "0"}}
	if got := f.snapshot(); !reflect.DeepEqual(got, want) {
		t.Fatalf("probe: %q", got)
	}
	u := &connectionTestUser{dynamicTestUser: dynamicTestUser{id: 42}, limit: 1}
	ran := make(chan struct{}, 1)
	if !m.StartUserConn(context.Background(), func() { t.Error("cancelled valid connection") }, func() { ran <- struct{}{} }, u, "192.0.2.1") {
		t.Fatal("rejected")
	}
	waitCall(t, ran)
	got := f.snapshot()
	want = append(want, []string{"hkeys", "soga_conn_42"}, []string{"HSETEX", "soga_conn_42", "EX", "60", "FIELDS", "1", "192.0.2.1", "0"})
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("connection: %q", got)
	}
	synchronous := false
	if !m.StartUserConn(context.Background(), func() {}, func() { synchronous = true }, u, "192.0.2.1") || !synchronous {
		t.Fatal("approved fast path")
	}
	if m.StartUserConn(context.Background(), func() {}, func() { t.Error("excess callback") }, u, "192.0.2.2") {
		t.Fatal("cached over-limit accepted")
	}
	m.EndUserConn(u, "192.0.2.1")
	if m.pCsxt7t[uuMFf5B{42, "192.0.2.1"}].a7fxke != 1 {
		t.Fatal("count")
	}
	m.EndUserConn(u, "192.0.2.1")
	if len(m.pCsxt7t) != 0 || len(f.snapshot()) != 4 {
		t.Fatal("end must only decrement locally")
	}
}
func TestRedisUnavailableAndRejectionSemantics(t *testing.T) {
	for _, command := range []string{"PING", "HSETEX"} {
		t.Run(command, func(t *testing.T) {
			f := redisFixture(t)
			m := f.manager(t)
			f.failures[command] = "unsupported"
			m.vih7o_()
			if m.cmA3rUr.Load() != (command == "PING") {
				t.Fatal("original pause flag")
			}
		})
	}
	t.Run("query failure permits callback", func(t *testing.T) {
		f := redisFixture(t)
		m := f.manager(t)
		f.failures["HKEYS"] = "query failed"
		u := &connectionTestUser{dynamicTestUser: dynamicTestUser{id: 8}, limit: 1}
		ran := make(chan struct{})
		m.StartUserConn(context.Background(), func() { t.Error("cancel") }, func() { close(ran) }, u, "ip")
		waitCall(t, ran)
		if len(f.snapshot()) != 2 || !m.pCsxt7t[uuMFf5B{8, "ip"}].aXnE_3N_gdt {
			t.Fatal("error branch")
		}
	})
	t.Run("fresh query rejects", func(t *testing.T) {
		f := redisFixture(t)
		m := f.manager(t)
		f.hashes["soga_conn_9"] = map[string]string{"old": "0"}
		u := &connectionTestUser{dynamicTestUser: dynamicTestUser{id: 9}, limit: 1}
		cancelled := make(chan struct{})
		if !m.StartUserConn(context.Background(), func() { close(cancelled) }, func() { t.Error("ran") }, u, "new") {
			t.Fatal("async admission must return true")
		}
		waitCall(t, cancelled)
		if len(f.snapshot()) != 1 || m.pCsxt7t[uuMFf5B{9, "new"}].aXnE_3N_gdt {
			t.Fatal("rejection")
		}
		m.EndUserConn(u, "new")
	})
	t.Run("paused is synchronous but respects cache", func(t *testing.T) {
		f := redisFixture(t)
		m := f.manager(t)
		m.cmA3rUr.Store(true)
		u := &connectionTestUser{dynamicTestUser: dynamicTestUser{id: 7}, limit: 1}
		ran := false
		m.StartUserConn(context.Background(), func() {}, func() { ran = true }, u, "one")
		if !ran || len(f.snapshot()) != 0 {
			t.Fatal("paused")
		}
		m.nkCoQC(7, []string{"one", "two"})
		if m.StartUserConn(context.Background(), func() {}, func() {}, u, "two") {
			t.Fatal("overfull cache must reject matching IP too")
		}
	})
}
func TestRedisCacheSnapshotAndPipeline(t *testing.T) {
	f := redisFixture(t)
	m := f.manager(t)
	start := time.Now()
	original := []string{"a"}
	m.nkCoQC(1, original)
	deadline := m.xLikNInZ[1].eaCklg8jch
	if d := deadline.Sub(start); d < 4500*time.Millisecond || d > 5*time.Second+time.Millisecond {
		t.Fatal("jitter", d)
	}
	m.ba5WPJg(1, "b")
	if len(original) != 1 || len(m.xLikNInZ[1].k4kEqO3) != 2 || m.xLikNInZ[1].eaCklg8jch != deadline {
		t.Fatal("append/expiry")
	}
	m.ba5WPJg(1, "b")
	if len(m.xLikNInZ[1].k4kEqO3) != 2 {
		t.Fatal("duplicate")
	}
	m.ba5WPJg(99, "a")
	if _, ok := m.xLikNInZ[99]; ok {
		t.Fatal("missing must stay missing")
	}
	m.xLikNInZ[1].eaCklg8jch = time.Now().Add(-time.Second)
	m.ba5WPJg(1, "c")
	if _, ok := m.n_augufdnS(1); ok {
		t.Fatal("expired")
	}
	for i := int64(0); i < 501; i++ {
		m.pCsxt7t[uuMFf5B{i, "a"}] = &e36g5GkEd{1, true}
	}
	m.pCsxt7t[uuMFf5B{0, "b"}] = &e36g5GkEd{2, true}
	m.pCsxt7t[uuMFf5B{999, "pending"}] = &e36g5GkEd{1, false}
	m.pCsxt7t[uuMFf5B{999, "zero"}] = &e36g5GkEd{0, true}
	m.nZcet5T6783()
	commands := f.snapshot()
	if len(commands) != 501 {
		t.Fatal("renew command count", len(commands))
	}
	for _, a := range commands {
		if a[0] != "HSETEX" || a[2] != "EX" || a[3] != "60" {
			t.Fatal("unexpected framing", a)
		}
	}
	f.mu.Lock()
	if len(f.hashes["soga_conn_0"]) != 2 || len(f.hashes) != 501 {
		t.Error("renew snapshot")
	}
	f.mu.Unlock()
	m.p7KF5sbP = time.Nanosecond
	if e := m.uPDNclE9W_k(-7, "v6"); e != nil {
		t.Fatal(e)
	}
	commands = f.snapshot()
	if !reflect.DeepEqual(commands[len(commands)-1], []string{"HSETEX", "soga_conn_-7", "EX", "1", "FIELDS", "1", "v6", "0"}) {
		t.Fatal("minimum TTL")
	}
}

func TestRedisCallbackPanicAndCancelledContext(t *testing.T) {
	f := redisFixture(t)
	m := f.manager(t)
	u := &connectionTestUser{dynamicTestUser: dynamicTestUser{id: 31}, limit: 0}
	recovered := false
	func() {
		defer func() { recovered = recover() == "sync" }()
		m.StartUserConn(context.Background(), func() {}, func() { panic("sync") }, u, "ip")
	}()
	if !recovered || len(m.pCsxt7t) != 0 {
		t.Fatal("unlimited callback must propagate panic without bookkeeping")
	}
	u.limit = 1
	ran := make(chan struct{})
	m.StartUserConn(context.Background(), func() {}, func() { close(ran); panic("async") }, u, "ip")
	waitCall(t, ran)
	// A subsequent operation exercises the manager after its async callback panic.
	m.EndUserConn(u, "ip")
	u = &connectionTestUser{dynamicTestUser: dynamicTestUser{id: 32}, limit: 1}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	m.StartUserConn(ctx, func() { t.Error("cancel called twice") }, func() { t.Error("cancelled request ran") }, u, "cancelled")
	// Original code reserves the count synchronously, and does not remove it when
	// the goroutine observes an already-cancelled context; EndUserConn owns that.
	m.p81VkaYJai.Lock()
	entry := m.pCsxt7t[uuMFf5B{32, "cancelled"}]
	count, approved := entry.a7fxke, entry.aXnE_3N_gdt
	m.p81VkaYJai.Unlock()
	if count != 1 || approved {
		t.Fatal("cancelled bookkeeping")
	}
	m.EndUserConn(u, "cancelled")
}
