package G6bHD5JA

import (
	"context"
	"errors"
	"io"
	"net"
	"os"
	"path/filepath"
	H "recovered.local/soga/HxKViEp2a"
	C "recovered.local/soga/cNT24rG4F"
	"testing"
	"time"
)

func TestStreamDatagramForwardingActualLoopback(t *testing.T) {
	old := C.ConfigurationPath()
	defer C.MkwNWciPHS(old)
	path := filepath.Join(t.TempDir(), "soga.conf")
	if err := os.WriteFile(path, []byte("udp_timeout=1\n"), 0600); err != nil {
		t.Fatal(err)
	}
	C.MkwNWciPHS(path)
	if err := C.TvBECnD9fM6(); err != nil {
		t.Fatal(err)
	}
	server, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	remote, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer remote.Close()
	client, stream := net.Pipe()
	defer client.Close()
	defer stream.Close()
	client.SetDeadline(time.Now().Add(2 * time.Second))
	server.SetDeadline(time.Now().Add(2 * time.Second))
	forward, reverse := make(chan error, 1), make(chan error, 1)
	go func() {
		forward <- ZcvMMHXas8k(context.Background(), stream, remote, server.LocalAddr().(*net.UDPAddr), make([]byte, 4096))
	}()
	go func() { reverse <- DjO2Gp(context.Background(), remote, stream) }()
	if _, err = client.Write([]byte("datagram")); err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 32)
	n, addr, err := server.ReadFromUDP(buf)
	if err != nil || string(buf[:n]) != "datagram" {
		t.Fatal(n, err)
	}
	if _, err = server.WriteToUDP([]byte("response"), addr); err != nil {
		t.Fatal(err)
	}
	if _, err = io.ReadFull(client, buf[:8]); err != nil || string(buf[:8]) != "response" {
		t.Fatal(string(buf[:8]), err)
	}
	client.Close()
	remote.Close()
	for _, done := range []chan error{forward, reverse} {
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("copy did not stop")
		}
	}
}

type packetRelayFixture struct {
	net.PacketConn
	readErr, writeErr                        error
	readCalls, readDeadlines, writeDeadlines int
	written                                  []byte
	zero                                     bool
}

func (p *packetRelayFixture) ReadFrom(b []byte) (int, net.Addr, error) {
	p.readCalls++
	if p.readCalls > 1 {
		return 0, nil, io.EOF
	}
	if p.zero {
		return 0, nil, p.readErr
	}
	return copy(b, "packet"), nil, p.readErr
}
func (p *packetRelayFixture) WriteTo(b []byte, _ net.Addr) (int, error) {
	p.written = append(p.written, b...)
	return len(b), p.writeErr
}
func (p *packetRelayFixture) SetReadDeadline(time.Time) error {
	p.readDeadlines++
	return errors.New("ignored read deadline")
}
func (p *packetRelayFixture) SetWriteDeadline(time.Time) error {
	p.writeDeadlines++
	return errors.New("ignored write deadline")
}

func TestDatagramCopyOriginalErrorSemantics(t *testing.T) {
	src, dst := &relayErrorFixture{readErr: io.EOF}, &packetRelayFixture{writeErr: errors.New("write failure")}
	err := ZcvMMHXas8k(context.Background(), src, dst, &net.UDPAddr{}, make([]byte, 32))
	if err == nil || err.Error() != "EOF\nwrite failure" || string(dst.written) != "payload" || dst.readDeadlines != 1 || dst.writeDeadlines != 1 {
		t.Fatalf("forward: %v %+v", err, dst)
	}
	packet, stream := &packetRelayFixture{readErr: io.EOF}, &relayErrorFixture{}
	if err = DjO2Gp(context.Background(), packet, stream); err != io.EOF || len(stream.written) != 0 {
		t.Fatal("data accompanying ReadFrom error not dropped", err)
	}
	packet, stream = &packetRelayFixture{zero: true}, &relayErrorFixture{writeErr: errors.New("empty write failed")}
	err = DjO2Gp(context.Background(), packet, stream)
	if err == nil || err.Error() != "empty write failed" || stream.deadlines != 1 {
		t.Fatal("zero datagram was skipped", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if DjO2Gp(ctx, packet, stream) != H.ErrorAt5097800 || ZcvMMHXas8k(ctx, src, dst, nil, nil) != H.ErrorAt5097800 {
		t.Fatal("cancel error identity changed")
	}
}
