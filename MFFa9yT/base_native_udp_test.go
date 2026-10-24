package MFFa9yT

import (
	"context"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	D "recovered.local/soga/D_zCag1qTV"
	G "recovered.local/soga/G6bHD5JA"
	J "recovered.local/soga/J1P4KNvfI_J"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	N "recovered.local/soga/nzATvajt"
	"testing"
	"time"
)

type nativeRelayResponse struct {
	data           string
	source, remote net.Addr
	user           N.PgaNd3014U
}
type nativeRelayFixture struct {
	J.KsKg3yls
	local     net.Addr
	responses chan nativeRelayResponse
}

func (p *nativeRelayFixture) LocalAddr() net.Addr              { return p.local }
func (p *nativeRelayFixture) SetWriteDeadline(time.Time) error { return nil }
func (p *nativeRelayFixture) WriteToProxy(data []byte, u N.PgaNd3014U, source, remote net.Addr) (int, error) {
	p.responses <- nativeRelayResponse{string(data), source, remote, u}
	return len(data), nil
}

func TestNativeUDPActualRelayCacheAndUserCancellation(t *testing.T) {
	if os.Getenv("RECOVERED_NATIVE_UDP_FIXTURE") != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^TestNativeUDPActualRelayCacheAndUserCancellation$", "-test.timeout=20s")
		cmd.Env = append(os.Environ(), "RECOVERED_NATIVE_UDP_FIXTURE=1")
		if out, e := cmd.CombinedOutput(); e != nil {
			t.Fatalf("%v\n%s", e, out)
		}
		return
	}
	conf := filepath.Join(t.TempDir(), "soga.conf")
	if e := os.WriteFile(conf, []byte("udp_timeout=1\naccess_log_enable=false\nlog_file=false\nforbidden_bit_torrent=false\nban_private_ip=false\n"), 0600); e != nil {
		t.Fatal(e)
	}
	C.MkwNWciPHS(conf)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
	server, e := net.ListenPacket("udp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer server.Close()
	target, e := G.U80FDt(server.LocalAddr().String())
	if e != nil {
		t.Fatal(e)
	}
	userCtx, userCancel := context.WithCancel(context.Background())
	defer userCancel()
	user := &sessionUserFixture{alive: true, tcp: true, ctx: userCtx}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	b := &JWN154qJ4{zb8f3cHWW8v: ctx, hS8xXFfcBr: new(J.TaQ7Prswa)}
	b.wiy6P3.Store(&L.Wr5SVP{EMlXEc: 31})
	if e = b.init(); e != nil {
		t.Fatal(e)
	}
	local := &nativeRelayFixture{local: &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 20000}, responses: make(chan nativeRelayResponse, 3)}
	source := &net.UDPAddr{IP: net.IPv4(192, 0, 2, 50), Port: 30000}
	ignored, ignoreCancel := context.WithCancel(context.Background())
	ignoreCancel()
	var outbound string
	for _, payload := range []string{"first authenticated datagram", "cached second datagram"} {
		data := D.PoolAt50961D8.Get(4096)
		copy(data, payload)
		b.dR8c1ufH6k(ignored, J.AFui66gdB{Y8YsJ7QnEI3h: local, AIO3Qb: source, PNliwRWkrU: target, S5eGpRGjp5oo: user, AGHsUHa: data[:len(payload)]})
		_ = server.SetDeadline(time.Now().Add(3 * time.Second))
		buf := make([]byte, 4096)
		n, from, err := server.ReadFrom(buf)
		if err != nil {
			t.Fatal(err)
		}
		if string(buf[:n]) != payload {
			t.Fatal(string(buf[:n]))
		}
		if outbound != "" && outbound != from.String() {
			t.Fatal("NAT entry not reused")
		}
		outbound = from.String()
		if _, err = server.WriteTo(buf[:n], from); err != nil {
			t.Fatal(err)
		}
		select {
		case got := <-local.responses:
			if got.data != payload || got.user != user || got.source != source || got.remote.String() != server.LocalAddr().String() {
				t.Fatal(got)
			}
		case <-time.After(3 * time.Second):
			t.Fatal("no native reply")
		}
	}
	entry, exists := b.hS8xXFfcBr.Get(source.String())
	if !exists {
		t.Fatal("missing NAT entry")
	}
	<-entry.Ready()
	if entry.Conn() == nil {
		t.Fatal("missing remote")
	}
	userCancel()
	limit := time.Now().Add(3 * time.Second)
	for {
		_, ok := b.hS8xXFfcBr.Get(source.String())
		if !ok {
			break
		}
		if time.Now().After(limit) {
			t.Fatal("user cancellation did not remove NAT entry")
		}
		time.Sleep(time.Millisecond)
	}
}
