//go:build linux

package main

import (
	"bytes"
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

// This executes a freshly built cmd/soga, never the supplied ELF. The fixture
// lives inside user/network/mount namespaces. Its network contains only lo;
// /etc is a bind-mounted temporary tree because original main initializes geo
// data at hard-coded paths before reading the configurable file flags.
func TestRebuiltMainWHMCSDatabaseAndNativeTCPUDP(t *testing.T) {
	if os.Getenv("SOGA_MAIN_WHMCS_INTEGRATION_ROLE") == "fixture" {
		runMainWHMCSNetworkFixture(t)
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
exec "$2" -test.run=^TestRebuiltMainWHMCSDatabaseAndNativeTCPUDP$ -test.v -test.timeout=90s`
	child := exec.CommandContext(ctx, "unshare", "-Urnm", "sh", "-c", script, "fixture", etc, os.Args[0])
	child.Env = append(os.Environ(), "SOGA_MAIN_WHMCS_INTEGRATION_ROLE=fixture", "SOGA_MAIN_WHMCS_INTEGRATION_BINARY="+binaryPath, "SOGA_MAIN_WHMCS_INTEGRATION_DIR="+dir)
	out, e := child.CombinedOutput()
	t.Logf("isolated application fixture:\n%s", out)
	if e != nil {
		t.Fatalf("isolated main: %v", e)
	}
}

func runMainWHMCSNetworkFixture(t *testing.T) {
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
	peer := newMainWHMCSPeer(t, password)
	defer peer.close()
	dir := os.Getenv("SOGA_MAIN_WHMCS_INTEGRATION_DIR")
	config := filepath.Join(dir, "soga.conf")
	text := fmt.Sprintf("type=whmcs\nserver_type=ss\napi=db\ndb_host=127.0.0.1\ndb_port=%d\ndb_name=fixture\ndb_user=fixture-user\ndb_password=fixture-pass\ndb_max_conn=2\ndb_max_idle_conn=2\nwhmcs_port=%d\nwhmcs_method=aes-128-gcm\nwhmcs_rate=1\nnode_id=71\nlisten=127.0.0.1\ncheck_interval=1\nsubmit_interval=1\nrandom_interval_range=0\nauto_update=false\ngeo_update_enable=false\naccess_log_enable=false\nlog_file_dir=\nlog_level=debug\nproxy_protocol=false\nudp_proxy_protocol=false\ndy_limit_enable=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=2\nudp_timeout=2\n", peer.listener.Addr().(*net.TCPAddr).Port, port1)
	if e := os.WriteFile(config, []byte(text), 0600); e != nil {
		t.Fatal(e)
	}
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
	app := exec.Command(os.Getenv("SOGA_MAIN_WHMCS_INTEGRATION_BINARY"), args...)
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
	wait("traffic submission", func() bool { return peer.up.Load() > 0 && peer.down.Load() > 0 })
	peer.userVersion.Store(1)
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
	if peer.userGets.Load() < 2 || peer.trafficPosts.Load() == 0 {
		t.Fatal("missing actual SQL polling/report", peer.userGets.Load(), peer.trafficPosts.Load())
	}
	if p := peer.failure.Load(); p != nil {
		t.Fatal("mysql fixture", *p)
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
	t.Logf("real WHMCS DB main verified: user queries=%d traffic CASE writes=%d upload=%d download=%d; TCP+UDP before/after password update; old TCP+UDP rejected; TERM exit0; only lo", peer.userGets.Load(), peer.trafficPosts.Load(), peer.up.Load(), peer.down.Load())
}

// This small independent MySQL server implements the exact result protocol
// needed by the original WHMCS SS SELECT and CASE UPDATE, not a deployed DB.
// Unlike the adapter unit fixture, no driver dial hook exists in the child app.
type mainWHMCSPeer struct {
	listener                                      net.Listener
	password                                      string
	userVersion, userGets, trafficPosts, up, down atomic.Int64
	failure                                       atomic.Pointer[string]
	mutex                                         sync.Mutex
	conns                                         map[net.Conn]bool
	wg                                            sync.WaitGroup
}

func newMainWHMCSPeer(t *testing.T, password string) *mainWHMCSPeer {
	t.Helper()
	l, e := net.Listen("tcp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	p := &mainWHMCSPeer{listener: l, password: password, conns: make(map[net.Conn]bool)}
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		for {
			c, e := l.Accept()
			if e != nil {
				return
			}
			p.mutex.Lock()
			p.conns[c] = true
			p.mutex.Unlock()
			p.wg.Add(1)
			go func() {
				defer p.wg.Done()
				defer c.Close()
				defer func() { p.mutex.Lock(); delete(p.conns, c); p.mutex.Unlock() }()
				if e := p.serve(c); e != nil {
					if !strings.Contains(e.Error(), "closed network connection") && e != io.EOF {
						message := e.Error()
						p.failure.Store(&message)
					}
				}
			}()
		}
	}()
	return p
}
func (p *mainWHMCSPeer) close() {
	p.listener.Close()
	p.mutex.Lock()
	for c := range p.conns {
		c.Close()
	}
	p.mutex.Unlock()
	p.wg.Wait()
}
func (p *mainWHMCSPeer) read(c net.Conn) ([]byte, error) {
	var h [4]byte
	if _, e := io.ReadFull(c, h[:]); e != nil {
		return nil, e
	}
	n := int(h[0]) | int(h[1])<<8 | int(h[2])<<16
	if n > 1<<20 {
		return nil, fmt.Errorf("oversized MySQL packet %d", n)
	}
	b := make([]byte, n)
	_, e := io.ReadFull(c, b)
	return b, e
}
func (p *mainWHMCSPeer) write(c net.Conn, seq byte, b []byte) error {
	out := append([]byte{byte(len(b)), byte(len(b) >> 8), byte(len(b) >> 16), seq}, b...)
	_, e := c.Write(out)
	return e
}
func mainWHMCSString(s string) []byte {
	if len(s) >= 251 {
		panic("fixture short value")
	}
	return append([]byte{byte(len(s))}, []byte(s)...)
}
func (p *mainWHMCSPeer) serve(c net.Conn) error {
	_ = c.SetDeadline(time.Now().Add(45 * time.Second))
	const capabilities uint32 = 0x200 | 0x8000 | 0x80000 | 8 | 0x2000
	h := append([]byte{10}, []byte("5.7.0-fixture\x00")...)
	h = append(h, 1, 0, 0, 0)
	h = append(h, []byte("12345678\x00")...)
	h = append(h, byte(capabilities&255), byte((capabilities>>8)&255), 33, 2, 0, byte((capabilities>>16)&255), byte((capabilities>>24)&255), 21)
	h = append(h, make([]byte, 10)...)
	h = append(h, []byte("abcdefghijkl\x00mysql_native_password\x00")...)
	if e := p.write(c, 0, h); e != nil {
		return e
	}
	login, e := p.read(c)
	if e != nil {
		return e
	}
	if len(login) < 33 || !bytes.HasPrefix(login[32:], []byte("fixture-user\x00")) {
		return fmt.Errorf("unexpected MySQL username")
	}
	caps := binary.LittleEndian.Uint32(login[:4])
	if caps&0x8000 == 0 {
		return fmt.Errorf("expected secure auth")
	}
	start := 32 + len("fixture-user") + 1
	if len(login) < start+1 {
		return io.ErrUnexpectedEOF
	}
	n := int(login[start])
	start++
	if len(login) < start+n {
		return io.ErrUnexpectedEOF
	}
	first := sha1.Sum([]byte("fixture-pass"))
	second := sha1.Sum(first[:])
	combined := append([]byte("12345678abcdefghijkl"), second[:]...)
	scramble := sha1.Sum(combined)
	for i := range scramble {
		scramble[i] ^= first[i]
	}
	if !bytes.Equal(login[start:start+n], scramble[:]) {
		return fmt.Errorf("MySQL password challenge mismatch")
	}
	if e = p.write(c, 2, []byte{0, 0, 0, 2, 0, 0, 0}); e != nil {
		return e
	}
	for {
		request, e := p.read(c)
		if e != nil {
			return e
		}
		if len(request) == 0 {
			return io.ErrUnexpectedEOF
		}
		switch request[0] {
		case 1:
			return nil
		case 14:
			if e = p.write(c, 1, []byte{0, 0, 0, 2, 0, 0, 0}); e != nil {
				return e
			}
			continue
		case 3:
		default:
			return fmt.Errorf("unexpected MySQL command %d", request[0])
		}
		query := string(request[1:])
		if strings.HasPrefix(query, "update user set u = case id\n") {
			// The application must report this fixture's UID, and SQL must retain the
			// original three CASE blocks and trailing IN list.
			pattern := regexp.MustCompile(`(?s)^update user set u = case id\nwhen 13 then u \+ ([0-9]+)\nend,\nd = case id\nwhen 13 then d \+ ([0-9]+)\nend,\nt = case id\nwhen 13 then ([0-9]+)\nend\nwhere id in\(13\)$`)
			m := pattern.FindStringSubmatch(query)
			if m == nil {
				return fmt.Errorf("unexpected traffic SQL %q", query)
			}
			up, e := strconv.ParseInt(m[1], 10, 64)
			if e != nil {
				return e
			}
			down, e := strconv.ParseInt(m[2], 10, 64)
			if e != nil {
				return e
			}
			stamp, e := strconv.ParseUint(m[3], 10, 32)
			if e != nil || stamp == 0 {
				return fmt.Errorf("timestamp %v %v", stamp, e)
			}
			p.up.Add(up)
			p.down.Add(down)
			p.trafficPosts.Add(1)
			if e = p.write(c, 1, []byte{0, 1, 0, 2, 0, 0, 0}); e != nil {
				return e
			}
			continue
		}
		const expected = "select id, u, d, transfer_enable, uuid, speedlimitss from user where enable = 1 and u + d < transfer_enable"
		if query != expected {
			return fmt.Errorf("unexpected SELECT %q", query)
		}
		p.userGets.Add(1)
		password := p.password
		if p.userVersion.Load() != 0 {
			password = "replacement-user-password"
		}
		cols := []string{"id", "u", "d", "transfer_enable", "uuid", "speedlimitss"}
		values := []string{"13", "0", "0", "10000000000", password, "0"}
		seq := byte(1)
		if e = p.write(c, seq, []byte{byte(len(cols))}); e != nil {
			return e
		}
		seq++
		for _, col := range cols {
			var b []byte
			for _, s := range []string{"def", "fixture", "user", "user", col, col} {
				b = append(b, mainWHMCSString(s)...)
			}
			b = append(b, 12, 33, 0, 0, 4, 0, 0, 253, 0, 0, 0, 0, 0)
			if e = p.write(c, seq, b); e != nil {
				return e
			}
			seq++
		}
		if e = p.write(c, seq, []byte{254, 0, 0, 2, 0}); e != nil {
			return e
		}
		seq++
		var row []byte
		for _, s := range values {
			row = append(row, mainWHMCSString(s)...)
		}
		if e = p.write(c, seq, row); e != nil {
			return e
		}
		seq++
		if e = p.write(c, seq, []byte{254, 0, 0, 2, 0}); e != nil {
			return e
		}
	}
}
