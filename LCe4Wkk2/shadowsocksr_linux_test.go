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

func TestShadowsocksROptions(t *testing.T) {
	o := ShadowsocksROptions{Name: "ssr", Server: "::1", Port: 443, Password: "password", Cipher: "none", Obfs: "plain", Protocol: "origin", UDP: true}
	a, e := CR6bC5(o)
	if e != nil || a.Addr() != "[::1]:443" || a.Type() != 3 || !a.SupportUDP() {
		t.Fatal(a, e)
	}
	if unsafe.Sizeof(o) != 168 || unsafe.Sizeof(RpNX5wlEsF{}) != 56 || unsafe.Offsetof(o.UDP) != 160 {
		t.Fatal("RTTI")
	}
	for _, ob := range []string{"plain", "http_simple", "http_post", "random_head", "tls1.2_ticket_auth", "tls1.2_ticket_fastauth"} {
		for _, pr := range []string{"origin", "auth_sha1_v4", "auth_aes128_md5", "auth_aes128_sha1", "auth_chain_a", "auth_chain_b"} {
			o.Obfs = ob
			o.Protocol = pr
			if _, e := CR6bC5(o); e != nil {
				t.Fatal(ob, pr, e)
			}
		}
	}
	o.Obfs = "plain"
	o.Protocol = "origin"
	for _, bad := range []string{"aes-128-gcm", "NONE", "plain", "DUMMY"} {
		o.Cipher = bad
		if _, e := CR6bC5(o); e == nil || !strings.Contains(e.Error(), "is not none or a supported stream cipher in ssr") {
			t.Fatal(bad, e)
		}
	}
	o.Cipher = "dummy"
	o.Obfs = "bad"
	if _, e := CR6bC5(o); e == nil || e.Error() != "ssr [::1]:443 initialize obfs error: Obfs bad not supported" {
		t.Fatal(e)
	}
	o.Obfs = "plain"
	o.Protocol = "bad"
	if _, e := CR6bC5(o); e == nil || e.Error() != "ssr [::1]:443 initialize protocol error: protocol bad not supported" {
		t.Fatal(e)
	}
	o.Cipher = "unsupported"
	if _, e := CR6bC5(o); e == nil || e.Error() != "ssr [::1]:443 initialize error: cipher not supported" {
		t.Fatal(e)
	}
}
func TestShadowsocksRActualTCPAndUDP(t *testing.T) {
	for _, algorithm := range []string{"RC4-MD5", "AES-128-CTR", "AES-192-CTR", "AES-256-CTR", "AES-128-CFB", "AES-192-CFB", "AES-256-CFB", "CHACHA20-IETF", "XCHACHA20", "none"} {
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
			a, e := CR6bC5(ShadowsocksROptions{Name: "ss-test", Server: host, Port: portNum, Cipher: algorithm, Password: "password", UDP: true, Obfs: "plain", Protocol: "origin"})
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
