package J1P4KNvfI_J

import (
	"context"
	"errors"
	"net"
	"recovered.local/soga/HxKViEp2a"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/nzATvajt"
	"testing"
	"time"
)

type proxyPacketFixture struct {
	net.PacketConn
	output   chan proxyPacket
	writeErr error
}
type proxyPacket struct {
	data    string
	address net.Addr
}

// 包测试不加载全局配置文件。记录原请求的deadline，让真实UDP socket使用
// 测试上限，避免尚未执行配置回调时默认零值立刻超时。
type observedDeadlinePacket struct {
	net.PacketConn
	readDeadline chan time.Time
}

func (p *observedDeadlinePacket) SetReadDeadline(v time.Time) error {
	p.readDeadline <- v
	return p.PacketConn.SetReadDeadline(time.Now().Add(3 * time.Second))
}

func (p *proxyPacketFixture) GetUser() nzATvajt.PgaNd3014U                  { return nil }
func (p *proxyPacketFixture) ReadFromProxy(b []byte) (int, net.Addr, error) { return p.ReadFrom(b) }
func (p *proxyPacketFixture) WriteToProxy(b []byte, a net.Addr) (int, error) {
	p.output <- proxyPacket{string(b), a}
	return len(b), p.writeErr
}

func TestOriginalUDPReverseRelayActualDatagram(t *testing.T) {
	remote, e := net.ListenPacket("udp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer remote.Close()
	local, e := net.ListenPacket("udp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer local.Close()
	peer, e := net.ListenPacket("udp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer peer.Close()
	sentinel := errors.New("proxy write stopped")
	p := &proxyPacketFixture{PacketConn: local, output: make(chan proxyPacket, 1), writeErr: sentinel}
	done := make(chan error, 1)
	observed := &observedDeadlinePacket{remote, make(chan time.Time, 1)}
	started := time.Now()
	go func() { done <- Nw8wZD(context.Background(), observed, p) }()
	if _, e = peer.WriteTo([]byte("real datagram"), remote.LocalAddr()); e != nil {
		t.Fatal(e)
	}
	select {
	case got := <-p.output:
		if got.data != "real datagram" || got.address.String() != peer.LocalAddr().String() {
			t.Fatal(got)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("relay stalled")
	}
	deadline := <-observed.readDeadline
	want := started.Add(time.Duration(cNT24rG4F.UDPTimeoutMinutes()) * time.Minute)
	if deadline.Sub(want) < 0 || deadline.Sub(want) > time.Second {
		t.Fatal("wrong UDP timeout", deadline, want)
	}
	select {
	case e = <-done:
		if e != sentinel {
			t.Fatal(e)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("relay did not return write error")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if e = Nw8wZD(ctx, nil, nil); e != HxKViEp2a.ErrorAt5097800 || errors.Is(e, context.Canceled) {
		t.Fatal("wrong cancellation identity", e)
	}
}
