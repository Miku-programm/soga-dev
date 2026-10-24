package inbound

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	CA "recovered.local/soga/util/rtx"
	L "recovered.local/soga/panel/types"
	Ntx "recovered.local/soga/net/limit"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
)

// Explicit opt-in to a verified local Redis 8 executable. This test never uses
// an externally configured Redis: it owns a loopback-only, nonpersistent child.
func TestNativeRedis8GlobalIPLifecycle(t *testing.T) {
	binaryPath := os.Getenv("SOGA_REDIS8_TEST_BINARY")
	if binaryPath == "" {
		t.Skip("set SOGA_REDIS8_TEST_BINARY to a local Redis 8 executable")
	}
	if os.Getenv("SOGA_REDIS8_NATIVE_CHILD") != "1" {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestNativeRedis8GlobalIPLifecycle$", "-test.v", "-test.timeout=55s")
		cmd.Env = append(os.Environ(), "SOGA_REDIS8_NATIVE_CHILD=1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("native Redis fixture: %v\n%s", err, out)
		}
		t.Logf("%s", out)
		return
	}
	version, err := exec.Command(binaryPath, "--version").CombinedOutput()
	if err != nil || !strings.Contains(string(version), "v=8.") {
		t.Fatalf("Redis 8 required: %v %s", err, version)
	}
	t.Logf("%s", bytes.TrimSpace(version))
	reserve, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := reserve.Addr().String()
	port := reserve.Addr().(*net.TCPAddr).Port
	reserve.Close()
	logFile, err := os.Create(filepath.Join(t.TempDir(), "redis.log"))
	if err != nil {
		t.Fatal(err)
	}
	defer logFile.Close()
	process := exec.Command(binaryPath, "--bind", "127.0.0.1", "--port", fmt.Sprint(port), "--protected-mode", "yes", "--save", "", "--appendonly", "no", "--dir", t.TempDir())
	process.Stdout = logFile
	process.Stderr = logFile
	if err = process.Start(); err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- process.Wait() }()
	t.Cleanup(func() {
		_ = process.Process.Signal(syscall.SIGTERM)
		select {
		case <-done:
		case <-time.After(3 * time.Second):
			_ = process.Process.Kill()
			<-done
		}
	})
	client := redis.NewClient(&redis.Options{Addr: addr, MaxRetries: -1, ReadTimeout: time.Second, WriteTimeout: time.Second})
	defer client.Close()
	ctx := context.Background()
	wait := func(label string, fn func() bool) {
		t.Helper()
		until := time.Now().Add(6 * time.Second)
		for !fn() {
			if time.Now().After(until) {
				t.Fatal("timeout:", label)
			}
			time.Sleep(20 * time.Millisecond)
		}
	}
	wait("Redis ready", func() bool { return client.Ping(ctx).Err() == nil })
	path := filepath.Join(t.TempDir(), "soga.conf")
	conf := "listen=127.0.0.1\ndy_limit_enable=false\naccess_log_enable=false\nlog_file=false\nban_private_ip=false\nforbidden_bit_torrent=false\ntcp_timeout=30\nproxy_protocol=false\nudp_proxy_protocol=false\nredis_ip_limit_enable=true\nredis_ip_limit_expiry=2\nredis_ip_limit_addr=" + addr + "\n"
	if err = os.WriteFile(path, []byte(conf), 0600); err != nil {
		t.Fatal(err)
	}
	C.MkwNWciPHS(path)
	if err = C.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
	const userID int64 = 80173
	const credential = "11111111-1111-4111-8111-111111111111"
	create := func(nodeID int) *KHonGaqbIy_6 {
		t.Helper()
		user := panelVless(userID, credential)
		user.Xmrna8 = 1
		panel := &Q.D5gTIZ5aVg{H47AzZsYHB1N: &L.Wr5SVP{EMlXEc: nodeID, Ck6DxjB: L.VkkFLH{XmIaPj3CzOF: "tcp"}}, VTme1SuHL6i_: L.Sxcf3wu3L{user}}
		core, e := oHr1LeATmS(&DV_XPJ4z0{Dmpl48sXM: ctx, HG5k0365: panel})
		if e != nil {
			t.Fatal(e)
		}
		v := core.(*KHonGaqbIy_6)
		t.Cleanup(v.Stop)
		if _, ok := v.af8f_J.(*Ntx.YM8R9i); !ok {
			t.Fatalf("not actual Redis manager: %T", v.af8f_J)
		}
		v.jfLH8mvq5e03.Store(uint32(CA.C5Jff5vs0()))
		if e = v.Start(); e != nil {
			t.Fatal(e)
		}
		return v
	}
	a, b := create(601), create(602)
	echo, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer echo.Close()
	go func() {
		for {
			c, e := echo.Accept()
			if e != nil {
				return
			}
			go func() { defer c.Close(); c.SetDeadline(time.Now().Add(40 * time.Second)); io.Copy(c, c) }()
		}
	}()
	targetPort := echo.Addr().(*net.TCPAddr).Port
	payload := []byte("native-redis-authenticated-echo")
	connect := func(core *KHonGaqbIy_6, ip string) net.Conn {
		t.Helper()
		d := net.Dialer{Timeout: time.Second, LocalAddr: &net.TCPAddr{IP: net.ParseIP(ip)}}
		c, e := d.Dial("tcp4", core.kX5i8U[0].Addr().String())
		if e != nil {
			t.Fatal(e)
		}
		t.Cleanup(func() { c.Close() })
		c.SetDeadline(time.Now().Add(6 * time.Second))
		id := uuid.FromStringOrNil(credential)
		wire := append([]byte{0}, id[:]...)
		wire = append(wire, 0, 1)
		wire = binary.BigEndian.AppendUint16(wire, uint16(targetPort))
		wire = append(wire, 1, 127, 0, 0, 1)
		wire = append(wire, payload...)
		if _, e = c.Write(wire); e != nil {
			t.Fatal(e)
		}
		return c
	}
	accepted := func(c net.Conn) {
		t.Helper()
		want := append([]byte{0, 0}, payload...)
		got := make([]byte, len(want))
		if _, e := io.ReadFull(c, got); e != nil || !bytes.Equal(got, want) {
			t.Fatalf("authenticated target echo: %v %x", e, got)
		}
		c.SetDeadline(time.Time{})
	}
	closed := func(c net.Conn) {
		t.Helper()
		c.SetReadDeadline(time.Now().Add(3 * time.Second))
		buf := make([]byte, 512)
		for {
			_, e := c.Read(buf)
			if e != nil {
				if ne, ok := e.(net.Error); ok && ne.Timeout() {
					t.Fatalf("expected session cancellation, got timeout: %v", e)
				}
				return
			}
		}
	}
	key := fmt.Sprintf("soga_conn_%d", userID)
	first := connect(a, "127.0.0.2")
	accepted(first)
	wait("first Redis field", func() bool { return client.HExists(ctx, key, "127.0.0.2").Val() })
	denied := connect(b, "127.0.0.3")
	closed(denied)
	if client.HExists(ctx, key, "127.0.0.3").Val() {
		t.Fatal("denied IP registered")
	}
	sameIP := connect(b, "127.0.0.2")
	accepted(sameIP)
	if n := client.HLen(ctx, key).Val(); n != 1 {
		t.Fatalf("same IP on two nodes counted twice: %d", n)
	}
	// Wait beyond the real two-second field expiry while both real sessions stay
	// active. Only the original renewal goroutine can keep this hash field alive.
	time.Sleep(3200 * time.Millisecond)
	if !client.HExists(ctx, key, "127.0.0.2").Val() {
		t.Fatal("active field was not renewed")
	}
	t.Log("two native nodes: global deny, same-IP sharing, and periodic renewal passed")
	a.spcNlUL(nil)
	closed(first)
	sameIP.SetDeadline(time.Now().Add(2 * time.Second))
	sameIP.Write(payload)
	got := make([]byte, len(payload))
	if _, e := io.ReadFull(sameIP, got); e != nil || !bytes.Equal(got, payload) {
		t.Fatalf("other node lost its active session: %v", e)
	}
	sameIP.Close()
	wait("last active IP expires", func() bool { return !client.HExists(ctx, key, "127.0.0.2").Val() })
	// The original per-node cache also expires within two seconds. No direct
	// mutation of the Redis manager/cache is used to admit the new client.
	time.Sleep(2100 * time.Millisecond)
	reconnected := connect(b, "127.0.0.3")
	accepted(reconnected)
	wait("new IP after reconnect", func() bool { return client.HExists(ctx, key, "127.0.0.3").Val() })
	b.spcNlUL(nil)
	closed(reconnected)
	wait("canceled session stops renewal", func() bool { return !client.HExists(ctx, key, "127.0.0.3").Val() })
	t.Log("native user cancellation, field expiry, and different-IP reconnect passed")
}
