//go:build linux

package LCe4Wkk2

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	E "recovered.local/soga/E3F12urwswL"
	cipher "recovered.local/soga/m3Cw8P3_Le"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestShadowsocksOptions(t *testing.T) {
	base := ShadowsocksOptions{Name: "ss", Server: "::1", Port: 443, Cipher: "aes-128-gcm", Password: "pw", UDP: true}
	a, e := PQ9jOjtX(base)
	if e != nil || a.Addr() != "[::1]:443" || a.Type() != 2 || !a.SupportUDP() {
		t.Fatal(a, e)
	}
	base.Plugin = "unknown"
	if a, e = PQ9jOjtX(base); e != nil || a.pChyaQS6 != "" {
		t.Fatal(a, e)
	}
	base.Plugin = "obfs"
	base.PluginOptions = map[string]any{"mode": "tls"}
	a, e = PQ9jOjtX(base)
	if e != nil || a.pChyaQS6 != "tls" || a.ar09vYxbfPZ.OJdz_5FmLCKj != "bing.com" {
		t.Fatal(a, e)
	}
	base.PluginOptions["mode"] = "bad"
	if _, e = PQ9jOjtX(base); e == nil || !strings.Contains(e.Error(), "obfs mode error: bad") {
		t.Fatal(e)
	}
	base.Plugin = "v2ray-plugin"
	base.PluginOptions = map[string]any{"mode": "websocket", "skip-cert-verify": true}
	a, e = PQ9jOjtX(base)
	if e != nil || !a.x93TNqC4dB.AUavUP || a.x93TNqC4dB.HjmCr8 || a.x93TNqC4dB.D6eC4LPWJm0 != "" {
		t.Fatal(a, e)
	}
	base.PluginOptions["tls"] = true
	base.PluginOptions["mux"] = false
	a, e = PQ9jOjtX(base)
	if e != nil || a.x93TNqC4dB.AUavUP || !a.x93TNqC4dB.HjmCr8 {
		t.Fatal(a, e)
	}
	base.Cipher = "unsupported"
	if _, e = PQ9jOjtX(base); e == nil || e.Error() != "ss [::1]:443 initialize error: cipher not supported" {
		t.Fatal(e)
	}
	if unsafe.Sizeof(ShadowsocksOptions{}) != 128 || unsafe.Sizeof(TKaQisW{}) != 56 || unsafe.Sizeof(cJlUN29eP{}) != 32 || unsafe.Sizeof(dXR4UapBiBa{}) != 72 {
		t.Fatal("RTTI")
	}
}
func TestShadowsocksActualTCPAndUDP(t *testing.T) {
	for _, algorithm := range []string{"aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "chacha20-ietf-poly1305", "xchacha20-ietf-poly1305", "RC4-MD5", "AES-128-CTR", "AES-192-CTR", "AES-256-CTR", "AES-128-CFB", "AES-192-CFB", "AES-256-CFB", "CHACHA20-IETF", "XCHACHA20", "none"} {
		t.Run(algorithm, func(t *testing.T) {
			serverCipher, e := cipher.B93MVPtD(algorithm, nil, "password")
			if e != nil {
				t.Fatal(e)
			}
			listener, e := net.Listen("tcp", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer listener.Close()
			host, port, _ := net.SplitHostPort(listener.Addr().String())
			portNum, _ := strconv.Atoi(port)
			a, e := PQ9jOjtX(ShadowsocksOptions{Name: "ss-test", Server: host, Port: portNum, Cipher: algorithm, Password: "password", UDP: true})
			if e != nil {
				t.Fatal(e)
			}
			m := &E.D0mLgQF{WJ95LyOL: "test.example", JRXMVvg: 443}
			wantHeader := []byte{3, 12, 't', 'e', 's', 't', '.', 'e', 'x', 'a', 'm', 'p', 'l', 'e', 1, 187}
			payload := bytes.Repeat([]byte("payload"), 3000)
			done := make(chan error, 1)
			go func() {
				raw, e := listener.Accept()
				if e != nil {
					done <- e
					return
				}
				defer raw.Close()
				raw.SetDeadline(time.Now().Add(4 * time.Second))
				conn := serverCipher.StreamConn(raw)
				header := make([]byte, len(wantHeader))
				if _, e = io.ReadFull(conn, header); e == nil && !bytes.Equal(header, wantHeader) {
					e = fmt.Errorf("header %x", header)
				}
				b := make([]byte, len(payload))
				if e == nil {
					_, e = io.ReadFull(conn, b)
				}
				if e == nil && !bytes.Equal(b, payload) {
					e = fmt.Errorf("payload")
				}
				if e == nil {
					_, e = conn.Write(b)
				}
				done <- e
			}()
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			defer cancel()
			conn, e := a.DialContext(ctx, m)
			if e != nil {
				t.Fatal(e)
			}
			conn.SetDeadline(time.Now().Add(4 * time.Second))
			if _, e = conn.Write(payload); e != nil {
				t.Fatal(e)
			}
			received := make([]byte, len(payload))
			if _, e = io.ReadFull(conn, received); e != nil || !bytes.Equal(received, payload) {
				t.Fatal(e)
			}
			if len(conn.Chains()) != 1 || conn.Chains()[0] != "ss-test" {
				t.Fatal(conn.Chains())
			}
			conn.Close()
			if e = <-done; e != nil {
				t.Fatal(e)
			}
			rawPacket, e := net.ListenPacket("udp", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer rawPacket.Close()
			rawPacket.SetDeadline(time.Now().Add(4 * time.Second))
			serverPacket := serverCipher.PacketConn(rawPacket)
			_, port, _ = net.SplitHostPort(rawPacket.LocalAddr().String())
			a.B8EHsVY7a.x3HUaU8Ln1S1 = net.JoinHostPort(host, port)
			packet, e := a.ListenPacketContext(ctx, nil)
			if e != nil {
				t.Fatal(e)
			}
			defer packet.Close()
			packet.SetDeadline(time.Now().Add(4 * time.Second))
			target := &net.UDPAddr{IP: net.IPv4(1, 2, 3, 4), Port: 53}
			if n, e := packet.WriteTo([]byte("abc"), target); n != 10 || e != nil {
				t.Fatal(n, e)
			}
			wire := make([]byte, 65536)
			n, peer, e := serverPacket.ReadFrom(wire)
			if e != nil || !bytes.Equal(wire[:n], []byte{1, 1, 2, 3, 4, 0, 53, 'a', 'b', 'c'}) {
				t.Fatal(n, e)
			}
			if _, e = serverPacket.WriteTo(wire[:n], peer); e != nil {
				t.Fatal(e)
			}
			n, from, e := packet.ReadFrom(wire)
			if e != nil || n != 3 || string(wire[:n]) != "abc" || from.String() != target.String() {
				t.Fatal(n, from, e)
			}
		})
	}
}
