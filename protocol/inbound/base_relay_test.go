package inbound

import (
	"bytes"
	"context"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	G "recovered.local/soga/net/relay"
	J "recovered.local/soga/net/sess"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	"testing"
	"time"
)

// 配置是原全局，使用独立测试进程执行真实文件加载，避免改变其它包内用例的初始状态。
func TestCentralSessionActualTCPAndUDP(t *testing.T) {
	if os.Getenv("RECOVERED_INBOUND_RELAY_FIXTURE") != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^TestCentralSessionActualTCPAndUDP$", "-test.timeout=20s")
		cmd.Env = append(os.Environ(), "RECOVERED_INBOUND_RELAY_FIXTURE=1")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("fixture %v\n%s", err, out)
		}
		return
	}
	conf := filepath.Join(t.TempDir(), "soga.conf")
	if err := os.WriteFile(conf, []byte("tcp_timeout=1\nudp_timeout=1\naccess_log_enable=false\nlog_file=false\nforbidden_bit_torrent=false\nban_private_ip=false\n"), 0600); err != nil {
		t.Fatal(err)
	}
	C.MkwNWciPHS(conf)
	if err := C.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
	for _, kind := range []string{"tcp", "udp", "proxy-udp"} {
		t.Run(kind, func(t *testing.T) { testActualInboundRelay(t, kind) })
	}
}

type proxyRelayFixture struct {
	*sessionConnFixture
	address net.Addr
}

func (c *proxyRelayFixture) ReadFromProxy(p []byte) (int, net.Addr, error) {
	n, e := c.Conn.Read(p)
	return n, c.address, e
}
func (c *proxyRelayFixture) WriteToProxy(p []byte, a net.Addr) (int, error) { return c.Conn.Write(p) }

func testActualInboundRelay(t *testing.T, kind string) {
	t.Helper()
	var addr net.Addr
	serverDone := make(chan error, 1)
	if kind == "tcp" {
		ln, err := net.Listen("tcp4", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()
		addr = ln.Addr()
		go func() {
			c, e := ln.Accept()
			if e != nil {
				serverDone <- e
				return
			}
			defer c.Close()
			_ = c.SetDeadline(time.Now().Add(5 * time.Second))
			_, e = io.Copy(c, c)
			serverDone <- e
		}()
	} else {
		pc, err := net.ListenPacket("udp4", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer pc.Close()
		addr = pc.LocalAddr()
		go func() {
			p := make([]byte, 4096)
			_ = pc.SetDeadline(time.Now().Add(5 * time.Second))
			n, a, e := pc.ReadFrom(p)
			if e == nil {
				_, e = pc.WriteTo(p[:n], a)
			}
			serverDone <- e
		}()
	}
	target, err := G.U80FDt(addr.String())
	if err != nil {
		t.Fatal(err)
	}
	client, inbound := net.Pipe()
	defer client.Close()
	defer inbound.Close()
	_ = client.SetDeadline(time.Now().Add(5 * time.Second))
	user := &sessionUserFixture{alive: true, tcp: true, ctx: context.Background()}
	k := J.ZES6F5(0)
	if kind != "tcp" {
		k = 1
	}
	raw := &sessionConnFixture{Conn: inbound, user: user, dest: target, kind: k}
	var conn J.Ki8urSHTx44 = raw
	if kind == "proxy-udp" {
		conn = &proxyRelayFixture{raw, addr}
	}
	s := &J.ZzkV7_5l{KeHzUwuXWh0: net.IPv4(192, 0, 2, 1), Uqzvq94e: "masked", C66t8X: "masked", Iayb4V: target, PayAlbdv6: user, YlxwftjCvz_: k}
	b := &JWN154qJ4{zb8f3cHWW8v: context.Background()}
	b.wiy6P3.Store(&L.Wr5SVP{EMlXEc: 31})
	if err = b.init(); err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	go func() { b.xR518G_hBbKv(s, conn); close(done) }()
	data := []byte("original UDP payload")
	if kind == "tcp" {
		data = bytes.Repeat([]byte("original TCP bytes\x00"), 4096)
	}
	writeDone := make(chan error, 1)
	go func() { _, e := client.Write(data); writeDone <- e }()
	received := make([]byte, len(data))
	if _, err = io.ReadFull(client, received); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(received, data) {
		t.Fatal("relay payload mismatch")
	}
	if err = <-writeDone; err != nil {
		t.Fatal(err)
	}
	client.Close()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("session did not terminate")
	}
	select {
	case err = <-serverDone:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("server stalled")
	}
	if s.IcoyFqRBU != 31 || s.R61M7ZTr == nil {
		t.Fatalf("route metadata not set: %+v", s)
	}
	if kind == "tcp" && s.IksfOg == nil {
		t.Fatal("outbound remote IP missing")
	}
	if kind != "tcp" && s.IksfOg != nil {
		t.Fatal("UDP must not set outbound remote")
	}
	if len(user.calls) < 2 || user.calls[len(user.calls)-2] != "deltcp" || user.calls[len(user.calls)-1] != "delalive:masked:masked" {
		t.Fatalf("cleanup %v", user.calls)
	}
}
