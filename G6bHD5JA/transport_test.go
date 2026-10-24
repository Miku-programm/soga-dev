package G6bHD5JA

import (
	"context"
	"errors"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestRecoveredHTTPTransportFieldsAndLocalDial(t *testing.T) {
	clientOnceAt50C7B70.Do(WuvamjL5)
	transport, ok := clientAt5029428.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("transport %T", clientAt5029428.Transport)
	}
	if clientAt5029428.Timeout != 60*time.Second || !transport.DisableKeepAlives || transport.MaxIdleConns != 5 || transport.IdleConnTimeout != 10*time.Second || transport.ResponseHeaderTimeout != 30*time.Second || transport.TLSHandshakeTimeout != 10*time.Second || transport.ExpectContinueTimeout != time.Second || transport.ForceAttemptHTTP2 || transport.Proxy == nil {
		t.Fatalf("transport differs: %#v", transport)
	}
	if conn, err := transport.DialContext(context.Background(), "tcp", "bad address"); err == nil || conn != nil {
		t.Fatal("invalid host:port accepted")
	}
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	accepted := make(chan net.Conn, 1)
	go func() { conn, _ := listener.Accept(); accepted <- conn }()
	conn, err := transport.DialContext(context.Background(), "tcp4", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
	serverConn := <-accepted
	if serverConn == nil {
		t.Fatal("missing accepted connection")
	}
	serverConn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	conn, err = transport.DialContext(ctx, "tcp4", listener.Addr().String())
	if err == nil || conn != nil {
		t.Fatal("canceled context dial unexpectedly succeeded")
	}
	// 原聚合错误不提供 Unwrap，不应替换成 errors.Join。
	if errors.Is(err, context.Canceled) {
		t.Fatal("error unwrap behavior changed")
	}
}
