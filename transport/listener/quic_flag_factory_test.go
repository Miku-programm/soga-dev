package listener

import (
	"bytes"
	"encoding/hex"
	"io"
	"net"
	T "recovered.local/soga/transport/encmux"
	"testing"
	"time"
)

// 验证原样本唯一直接 HNb 调用者的实际链：先监听 TCP，再包隧道验证。
// quic-tunnel 配置名没有在此分支创建 UDP/QUIC 连接，不能据名称扩展协议。
func TestOriginalFactoryQUICFlagRemainsTCP(t *testing.T) {
	cfg, e := T.FZzfvtn(nil, &T.OutboundConfiguration{Protocol: "quic-tunnel", Method: "aes-128-gcm", Password: "test-tunnel-secret", Mux: true})
	if e != nil || !cfg.IsQuic() || cfg.IsMux() {
		t.Fatal(cfg, e)
	}
	options := &Ap9uD2{Host: "127.0.0.1", Network: "tcp", Tunnel: cfg}
	listener, e := options.NewListener()
	if e != nil {
		t.Fatal(e)
	}
	defer listener.Close()
	if _, ok := listener.(*T.Ppx7OU); !ok {
		t.Fatalf("listener %T", listener)
	}
	if _, ok := listener.Addr().(*net.TCPAddr); !ok {
		t.Fatalf("address %T", listener.Addr())
	}
	udp, e := net.ListenPacket("udp4", listener.Addr().String())
	if e != nil {
		t.Fatal("unexpected UDP binding", e)
	}
	defer udp.Close()
	payload := bytes.Repeat([]byte("factory keeps original TCP stream"), 64)
	done := make(chan error, 1)
	go func() {
		c, e := listener.Accept()
		if e != nil {
			done <- e
			return
		}
		defer c.Close()
		c.SetDeadline(time.Now().Add(2 * time.Second))
		b := make([]byte, len(payload))
		_, e = io.ReadFull(c, b)
		if e == nil && !bytes.Equal(b, payload) {
			e = io.ErrUnexpectedEOF
		}
		if e == nil {
			_, e = c.Write(b)
		}
		done <- e
	}()
	c, e := net.DialTimeout("tcp4", listener.Addr().String(), time.Second)
	if e != nil {
		t.Fatal(e)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(2 * time.Second))
	token, _ := hex.DecodeString("9db29d0cad3a9fc2aea9cacd1135071c")
	request := append(append(token, byte(1)), payload...)
	if _, e = c.Write(request); e != nil {
		t.Fatal(e)
	}
	b := make([]byte, len(payload))
	if _, e = io.ReadFull(c, b); e != nil || !bytes.Equal(b, payload) {
		t.Fatal(e)
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
}
