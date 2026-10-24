package g0MlUOhvO

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	P "recovered.local/soga/JcmlVIK0T0i"
	"testing"
	"time"
)

// 原 quic-tunnel 标志的可达行为：WrapListener/WrapClientConn 原样返回
// 已有流，HNb 仍增加令牌验证和可选 PROXY。这个测试不声称建立了 QUIC。
func TestOriginalQUICFlagRawTCPTokenProxyAndFraming(t *testing.T) {
	for _, proxy := range []bool{false, true} {
		for _, command := range []byte{1, 2} {
			t.Run(fmt.Sprintf("proxy=%t/command=%d", proxy, command), func(t *testing.T) {
				receive := 0
				if proxy {
					receive = 1
				}
				cfg, e := FZzfvtn(&InboundConfiguration{Protocol: "quic-tunnel", Method: "aes-128-gcm", Password: "test-tunnel-secret", Certificate: tunnelCertificate(t), ReceiveProxy: receive, Mux: true}, nil)
				if e != nil {
					t.Fatal(e)
				}
				if !cfg.IsQuic() || cfg.IsMux() || !cfg.rj1DhSBHZ7O {
					t.Fatal("configuration flags")
				}
				raw, e := net.Listen("tcp4", "127.0.0.1:0")
				if e != nil {
					t.Fatal(e)
				}
				defer raw.Close()
				unchanged, e := cfg.WrapListener(raw)
				if unchanged != raw || e != nil {
					t.Fatal("QUIC wrapper must return existing listener")
				}
				listener, e := HNbZee047(raw, cfg)
				if e != nil {
					t.Fatal(e)
				}
				defer listener.Close()
				var verify *Ppx7OU
				if proxy {
					outer, ok := listener.(*P.KprDQw4aZEOg)
					if !ok {
						t.Fatalf("outer %T", listener)
					}
					verify = outer.I9Hb4SAu.(*Ppx7OU)
				} else {
					verify = listener.(*Ppx7OU)
				}
				if verify.O6B2qZ0p0 != raw {
					t.Fatal("unexpected TLS/encryption/mux under verification")
				}
				payload := bytes.Repeat([]byte("original QUIC flag keeps existing stream"), 29)
				done := make(chan error, 1)
				go func() {
					c, e := listener.Accept()
					if e != nil {
						done <- e
						return
					}
					defer c.Close()
					c.SetDeadline(time.Now().Add(3 * time.Second))
					b := make([]byte, len(payload))
					_, e = io.ReadFull(c, b)
					if e == nil && !bytes.Equal(b, payload) {
						e = fmt.Errorf("payload mismatch")
					}
					if e == nil && proxy && c.RemoteAddr().String() != "198.51.100.23:12345" {
						e = fmt.Errorf("proxy remote %s", c.RemoteAddr())
					}
					if e == nil {
						_, e = c.Write(b)
					}
					done <- e
				}()
				conn, e := net.DialTimeout("tcp", raw.Addr().String(), time.Second)
				if e != nil {
					t.Fatal(e)
				}
				defer conn.Close()
				conn.SetDeadline(time.Now().Add(3 * time.Second))
				client, e := cfg.WrapClientConn(context.Background(), conn)
				if client != conn || e != nil {
					t.Fatal("must return original TCP connection")
				}
				// 独立 SHA3/HKDF fixture 已核固定向量，不用原验证函数生成令牌。
				token, _ := hex.DecodeString("9db29d0cad3a9fc2aea9cacd1135071c")
				data := append([]byte(nil), payload...)
				if proxy {
					data = append([]byte("PROXY TCP4 198.51.100.23 203.0.113.9 12345 443\r\n"), data...)
				}
				wire := append(token, command)
				if command == 2 {
					wire = binary.BigEndian.AppendUint16(wire, uint16(len(data)))
				}
				wire = append(wire, data...)
				if _, e = client.Write(wire); e != nil {
					t.Fatal(e)
				}
				if command == 2 {
					var n [2]byte
					if _, e = io.ReadFull(client, n[:]); e != nil {
						t.Fatal(e)
					}
					if int(binary.BigEndian.Uint16(n[:])) != len(payload) {
						t.Fatal("datagram framing")
					}
				}
				b := make([]byte, len(payload))
				if _, e = io.ReadFull(client, b); e != nil || !bytes.Equal(b, payload) {
					t.Fatal(e)
				}
				if e = <-done; e != nil {
					t.Fatal(e)
				}
			})
		}
	}
}
func TestOriginalQUICFlagStillRejectsWrongToken(t *testing.T) {
	cfg := testTunnelConfig(t)
	cfg.hZORuZstvS07 = true
	raw, e := net.Listen("tcp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer raw.Close()
	listener, e := HNbZee047(raw, cfg)
	if e != nil {
		t.Fatal(e)
	}
	defer listener.Close()
	done := make(chan error, 1)
	go func() {
		c, e := listener.Accept()
		if e != nil {
			done <- e
			return
		}
		defer c.Close()
		c.SetDeadline(time.Now().Add(time.Second))
		_, e = c.Read(make([]byte, 1))
		done <- e
	}()
	c, e := net.DialTimeout("tcp4", raw.Addr().String(), time.Second)
	if e != nil {
		t.Fatal(e)
	}
	defer c.Close()
	wrong := append(append([]byte(nil), cfg.x5Lor8hDyn[:]...), byte(1))
	wrong[0] ^= 0xff
	if _, e = c.Write(wrong); e != nil {
		t.Fatal(e)
	}
	if e = <-done; e == nil || e.Error() != "tunnel verify failed" {
		t.Fatal(e)
	}
}
