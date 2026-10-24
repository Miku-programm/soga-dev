package GYzZ5v

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
)

func TestOtherOriginalDescriptors(t *testing.T) {
	for _, c := range []struct {
		data []byte
		hash string
	}{
		{file_protos_addons_proto_rawDescGZIP(), "a83786c84c6b4eaf3276935ce2fb573cfbefe40edc81bd88343064b28434f84e"},
		{file_protos_grpc_proto_rawDescGZIP(), "936b738eed69404f8c429c1846b0016fb6971d6a2aadec848ad3b386f0324d14"},
	} {
		r, err := gzip.NewReader(bytes.NewReader(c.data))
		if err != nil {
			t.Fatal(err)
		}
		b, err := io.ReadAll(r)
		r.Close()
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprintf("%x", sha256.Sum256(b)) != c.hash {
			t.Fatal("original descriptor changed")
		}
	}
	addon := reflect.TypeOf(N_SQ5a{})
	hunk := reflect.TypeOf(Z27tYEZ9hy1o{})
	if addon.Size() != 80 || hunk.Size() != 64 || reflect.TypeOf(utpvEgZY{}).Size() != 16 || reflect.TypeOf(EmfCw_9j{}).Size() != 0 {
		t.Fatal("RTTI sizes differ")
	}
	for _, c := range []struct {
		typ    reflect.Type
		field  string
		offset uintptr
	}{{addon, "AfJaYuHm3H", 40}, {addon, "S3EXSAVeAUeF", 56}, {hunk, "Q2bCNbJ12i", 40}} {
		f, ok := c.typ.FieldByName(c.field)
		if !ok || f.Offset != c.offset {
			t.Fatal(c)
		}
	}
	var a N_SQ5a
	if err := proto.Unmarshal([]byte{10, 4, 'f', 'l', 'o', 'w', 18, 2, 1, 2}, &a); err != nil {
		t.Fatal(err)
	}
	if a.AfJaYuHm3H != "flow" || !bytes.Equal(a.S3EXSAVeAUeF, []byte{1, 2}) {
		t.Fatal(&a)
	}
}

type recoveredEchoService struct{}

func (recoveredEchoService) Tun(s XZFdvGb2) error {
	for {
		m, err := s.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err = s.Send(m); err != nil {
			return err
		}
	}
}

func TestOriginalGRPCStreamRealRoundTrip(t *testing.T) {
	listener := bufconn.Listen(1 << 16)
	server := grpc.NewServer()
	server.RegisterService(&serviceDescriptionAt5057580, recoveredEchoService{})
	go func() { _ = server.Serve(listener) }()
	defer server.Stop()
	defer listener.Close()
	conn, err := grpc.NewClient("passthrough:///target", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) { return listener.DialContext(ctx) }))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stream, err := conn.NewStream(ctx, &serviceDescriptionAt5057580.Streams[0], "/protos.GunService/Tun")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range [][]byte{nil, []byte("round trip"), {0, 1, 255, 0}} {
		if err = stream.SendMsg(&Z27tYEZ9hy1o{Q2bCNbJ12i: data}); err != nil {
			t.Fatal(err)
		}
		var got Z27tYEZ9hy1o
		if err = stream.RecvMsg(&got); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got.Q2bCNbJ12i, data) {
			t.Fatal(got.Q2bCNbJ12i)
		}
	}
	if err = stream.CloseSend(); err != nil {
		t.Fatal(err)
	}
	if err = stream.RecvMsg(new(Z27tYEZ9hy1o)); err != io.EOF {
		t.Fatalf("want EOF, got %v", err)
	}
	if serviceDescriptionAt5057580.Methods == nil || serviceDescriptionAt5057580.Metadata != "protos/grpc.proto" {
		t.Fatal("static descriptor changed")
	}
}

type failingStream struct {
	grpc.ServerStream
	err  error
	sent any
}

func (s *failingStream) RecvMsg(v any) error {
	v.(*Z27tYEZ9hy1o).Q2bCNbJ12i = []byte("partial")
	return s.err
}
func (s *failingStream) SendMsg(v any) error { s.sent = v; return s.err }
func TestOriginalGRPCErrorPaths(t *testing.T) {
	sentinel := errors.New("stream failed")
	raw := &failingStream{err: sentinel}
	s := &utpvEgZY{Y5md9Ywqjd4m: raw}
	got, err := s.Recv()
	if got != nil || err != sentinel {
		t.Fatalf("%v %v", got, err)
	}
	msg := new(Z27tYEZ9hy1o)
	if s.Send(msg) != sentinel || raw.sent != msg {
		t.Fatal("Send changed message or error")
	}
	err = (&EmfCw_9j{}).Tun(nil)
	if status.Code(err) != codes.Unimplemented || status.Convert(err).Message() != "method Tun not implemented" {
		t.Fatal(err)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("invalid server did not panic")
		}
	}()
	_ = eehVxKUB(struct{}{}, raw)
}
