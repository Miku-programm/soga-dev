package NxHajL

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"io"
	"net"
	"recovered.local/soga/GYzZ5v"
	"reflect"
	"testing"
	"time"
)

type testStream struct {
	ctx              context.Context
	frames           [][]byte
	recvErr, sendErr error
	sent             *GYzZ5v.Z27tYEZ9hy1o
	calls            int
}

func (s *testStream) Context() context.Context { return s.ctx }
func (s *testStream) Recv() (*GYzZ5v.Z27tYEZ9hy1o, error) {
	s.calls++
	if len(s.frames) == 0 {
		return nil, s.recvErr
	}
	p := s.frames[0]
	s.frames = s.frames[1:]
	return &GYzZ5v.Z27tYEZ9hy1o{Q2bCNbJ12i: p}, nil
}
func (s *testStream) Send(m *GYzZ5v.Z27tYEZ9hy1o) error { s.sent = m; return s.sendErr }

func TestOriginalFrameAndErrorSemantics(t *testing.T) {
	sentinel := errors.New("original cause")
	stream := &testStream{ctx: context.Background(), frames: [][]byte{[]byte("abc"), nil, []byte("d")}, recvErr: sentinel}
	canceled := 0
	conn := HpSkqj(stream, &net.TCPAddr{Port: 99}, func() { canceled++ })
	p := make([]byte, 2)
	for _, c := range []struct {
		n    int
		text string
	}{{2, "ab"}, {1, "c"}, {0, ""}, {1, "d"}} {
		n, err := conn.Read(p)
		if n != c.n || err != nil || string(p[:n]) != c.text {
			t.Fatal(n, err, string(p[:n]))
		}
	}
	if _, err := conn.Read(p); err == nil || err.Error() != "unable to read from gun tunnel" || errors.Is(err, sentinel) {
		t.Fatal(err)
	}
	data := []byte("payload")
	if n, err := conn.Write(data); n != len(data) || err != nil {
		t.Fatal(n, err)
	}
	data[0] = 'P'
	if stream.sent.Q2bCNbJ12i[0] != 'P' {
		t.Fatal("Write copied original slice")
	}
	stream.sendErr = sentinel
	if n, err := conn.Write(nil); n != 0 || err == nil || err.Error() != "unable to send data over gun" || errors.Is(err, sentinel) {
		t.Fatal(n, err)
	}
	if conn.RemoteAddr().String() != "0.0.0.0:0" || conn.LocalAddr().String() != ":99" {
		t.Fatal(conn.RemoteAddr(), conn.LocalAddr())
	}
	if conn.SetWriteDeadline(time.Now().Add(-time.Hour)) != nil {
		t.Fatal("write deadline original no-op changed")
	}
	conn.Close()
	conn.Close()
	if canceled != 2 {
		t.Fatal("Close unexpectedly became once-only")
	}
	if reflect.TypeOf(*conn).Size() != 64 || reflect.TypeOf(wtSXw52c8Wt4{}).Size() != 24 || reflect.TypeOf(Q1SOcdpqfGG{}).Size() != 56 {
		t.Fatal("layout")
	}
}

func TestOriginalForwardedAddress(t *testing.T) {
	for _, c := range []struct {
		headers metadata.MD
		want    string
	}{
		{metadata.Pairs("x-real-ip", "192.0.2.8", "x-forwarded-for", "198.51.100.9"), "192.0.2.8:321"},
		{metadata.Pairs("x-real-ip", "unknown", "x-forwarded-for", "bad, [2001:db8::7]:junk, 192.0.2.8"), "[2001:db8::7]:321"},
		{metadata.MD{}, "192.0.2.1:321"},
	} {
		ctx := peer.NewContext(context.Background(), &peer.Peer{Addr: &net.UDPAddr{IP: net.ParseIP("192.0.2.1"), Port: 321}})
		ctx = metadata.NewIncomingContext(ctx, c.headers)
		conn := HpSkqj(&testStream{ctx: ctx}, nil, nil)
		if conn.RemoteAddr().String() != c.want {
			t.Fatal(conn.RemoteAddr(), c.want)
		}
	}
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("x-real-ip", "192.0.2.8"))
	conn := HpSkqj(&testStream{ctx: ctx}, nil, nil)
	if conn.RemoteAddr().String() != "0.0.0.0:0" {
		t.Fatal("original requires peer before using metadata")
	}
}

func TestOriginalListenerTCPGRPCRoundTrip(t *testing.T) {
	raw, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	listener := SNiVo8(raw, "custom.Tunnel", nil)
	defer listener.Close()
	client, err := grpc.NewClient(raw.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-real-ip", "192.0.2.77")
	stream, err := client.NewStream(ctx, &grpc.StreamDesc{ServerStreams: true, ClientStreams: true}, "/custom.Tunnel/Tun")
	if err != nil {
		t.Fatal(err)
	}
	if err = stream.SendMsg(&GYzZ5v.Z27tYEZ9hy1o{Q2bCNbJ12i: []byte("hello")}); err != nil {
		t.Fatal(err)
	}
	accepted := make(chan net.Conn, 1)
	go func() { conn, _ := listener.Accept(); accepted <- conn }()
	var conn net.Conn
	select {
	case conn = <-accepted:
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	}
	if conn == nil {
		t.Fatal("nil connection")
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(time.Second))
	p := make([]byte, 5)
	if _, err = io.ReadFull(conn, p); err != nil || string(p) != "hello" {
		t.Fatal(string(p), err)
	}
	if !conn.RemoteAddr().(*net.TCPAddr).IP.Equal(net.ParseIP("192.0.2.77")) {
		t.Fatal(conn.RemoteAddr())
	}
	if n, err := conn.Write([]byte("world")); n != 5 || err != nil {
		t.Fatal(n, err)
	}
	var got GYzZ5v.Z27tYEZ9hy1o
	if err = stream.RecvMsg(&got); err != nil || string(got.Q2bCNbJ12i) != "world" {
		t.Fatal(&got, err)
	}
	conn.Close()
	if err = stream.RecvMsg(&got); err != io.EOF {
		t.Fatal(err)
	}
	listener.Close()
	if conn, err := listener.Accept(); conn != nil || err == nil || err.Error() != "grpc listener closed" {
		t.Fatal(conn, err)
	}
}
