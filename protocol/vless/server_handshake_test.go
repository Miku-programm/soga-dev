package vless

import (
	"bytes"
	"crypto/ecdh"
	"crypto/mlkem"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestServerStateLayoutAndInitialization(t *testing.T) {
	typ := reflect.TypeOf(LgF3OB1QX{})
	offsets := []uintptr{0, 24, 48, 72, 80, 88, 96, 104, 128, 152, 176, 184, 192, 216}
	if typ.Size() != 224 || typ.NumField() != len(offsets) {
		t.Fatal("server RTTI layout", typ)
	}
	for i, want := range offsets {
		if typ.Field(i).Offset != want {
			t.Fatalf("%s offset=%d want%d", typ.Field(i).Name, typ.Field(i).Offset, want)
		}
	}
	if reflect.TypeOf(MZKKvEA81HV{}).Size() != 72 {
		t.Fatal("session RTTI size")
	}
	var server LgF3OB1QX
	if e := server.Init(nil, 0, 0, 0, ""); e == nil || e.Error() != "empty nfsSKeysBytes" {
		t.Fatal(e)
	}
	key, _ := ecdh.X25519().GenerateKey(rand.Reader)
	kem, _ := mlkem.GenerateKey768()
	if e := server.Init([][]byte{key.Bytes(), kem.Bytes()}, 2, 0, 0, "100-35-35"); e != nil {
		t.Fatal(e)
	}
	if server.DptoPRgY != 1152 || !bytes.Equal(server.SpB4gP6aCP[0], key.PublicKey().Bytes()) || !bytes.Equal(server.SpB4gP6aCP[1], kem.EncapsulationKey().Bytes()) {
		t.Fatal("server key initialization")
	}
	if e := server.Init(nil, 0, 0, 0, ""); e == nil || e.Error() != "already initialized" {
		t.Fatal(e)
	}
	var invalid LgF3OB1QX
	if e := invalid.Init([][]byte{{1}}, 0, 0, 0, ""); e == nil || invalid.EiKprMhvVvc == nil {
		t.Fatal("partial state rollback")
	}
	server.Close()
	if !server.JNEGgibSOG {
		t.Fatal("Close flag")
	}
}

// 实际回环TCP同时运行独立恢复的客户端与服务端。全双工握手先写响应
// 再读对端padding，采用真实TCP缓冲语义，不用net.Pipe改变协议顺序。
func TestServerClientFullAndResumedConnections(t *testing.T) {
	for _, mode := range []uint32{0, 1, 2} {
		for _, mixed := range []bool{false, true} {
			t.Run(fmt.Sprintf("mode%d_mixed%v", mode, mixed), func(t *testing.T) {
				x, _ := ecdh.X25519().GenerateKey(rand.Reader)
				secrets := [][]byte{x.Bytes()}
				publics := [][]byte{x.PublicKey().Bytes()}
				if mixed {
					k, _ := mlkem.GenerateKey768()
					x2, _ := ecdh.X25519().GenerateKey(rand.Reader)
					secrets = append(secrets, k.Bytes(), x2.Bytes())
					publics = append(publics, k.EncapsulationKey().Bytes(), x2.PublicKey().Bytes())
				}
				var server LgF3OB1QX
				var client GPFxUP0_
				if e := server.Init(secrets, mode, 600, 600, "100-35-35"); e != nil {
					t.Fatal(e)
				}
				defer server.Close()
				if e := client.Init(publics, mode, 600, "100-35-35"); e != nil {
					t.Fatal(e)
				}
				for pass := 0; pass < 2; pass++ {
					listener, e := net.Listen("tcp", "127.0.0.1:0")
					if e != nil {
						t.Fatal(e)
					}
					done := make(chan error, 1)
					go func() {
						raw, e := listener.Accept()
						if e != nil {
							done <- e
							return
						}
						defer raw.Close()
						raw.SetDeadline(time.Now().Add(5 * time.Second))
						capture := []byte{9}
						conn, e := server.Handshake(raw, &capture)
						if e != nil {
							done <- fmt.Errorf("server handshake: %w", e)
							return
						}
						if capture != nil {
							done <- fmt.Errorf("authenticated capture retained")
							return
						}
						payload := make([]byte, 9000)
						if _, e = io.ReadFull(conn, payload); e != nil {
							done <- fmt.Errorf("server read: %w", e)
							return
						}
						if !bytes.Equal(payload, bytes.Repeat([]byte{42}, 9000)) {
							done <- fmt.Errorf("query mismatch")
							return
						}
						if _, e = conn.Write(bytes.Repeat([]byte{77}, 8300)); e != nil {
							done <- fmt.Errorf("server write: %w", e)
							return
						}
						done <- nil
					}()
					raw, e := net.Dial("tcp", listener.Addr().String())
					if e != nil {
						listener.Close()
						t.Fatal(e)
					}
					raw.SetDeadline(time.Now().Add(5 * time.Second))
					conn, e := client.Handshake(raw)
					if e == nil {
						_, e = conn.Write(bytes.Repeat([]byte{42}, 9000))
					}
					if e == nil {
						reply := make([]byte, 8300)
						_, e = io.ReadFull(conn, reply)
						if e == nil && !bytes.Equal(reply, bytes.Repeat([]byte{77}, 8300)) {
							e = fmt.Errorf("reply mismatch")
						}
					}
					raw.Close()
					listener.Close()
					serverErr := <-done
					if e != nil || serverErr != nil {
						t.Fatalf("pass%d client=%v server=%v", pass, e, serverErr)
					}
					if len(client.Wliqugg3J89) != 64 || len(client.HMnsiUUhd) != 16 {
						t.Fatal("session not stored")
					}
				}
			})
		}
	}
}

func makeServerResumeWire(t *testing.T, server *LgF3OB1QX, useAES bool, ticket []byte) ([]byte, []byte) {
	t.Helper()
	peer, _ := ecdh.X25519().GenerateKey(rand.Reader)
	secret, e := peer.ECDH(server.EiKprMhvVvc[0].(*ecdh.PrivateKey).PublicKey())
	if e != nil {
		t.Fatal(e)
	}
	wire := append(bytes.Repeat([]byte{17}, 16), peer.PublicKey().Bytes()...)
	h := WrnUSnC6(wire[:16], secret, useAES)
	wire = h.Seal(wire, nil, []byte{0, 32}, nil)
	wire = h.Seal(wire, nil, ticket, nil)
	return wire, secret
}

func TestServerAuthenticationCaptureAndReplay(t *testing.T) {
	key, _ := ecdh.X25519().GenerateKey(rand.Reader)
	var s LgF3OB1QX
	if e := s.Init([][]byte{key.Bytes()}, 0, 0, 0, "100-35-35"); e != nil {
		t.Fatal(e)
	}
	if _, e := new(LgF3OB1QX).Handshake(nil, nil); e == nil || e.Error() != "uninitialized" {
		t.Fatal(e)
	}
	ticket := bytes.Repeat([]byte{2}, 16)
	for _, aes := range []bool{true, false} {
		wire, secret := makeServerResumeWire(t, &s, aes, ticket)
		s.MOy_iE1vq = 1
		s.GIHdzaFGhi8 = map[[16]byte]*MZKKvEA81HV{[16]byte(ticket): {MwUsryYRwlpa: bytes.Repeat([]byte{5}, 64)}}
		capture := []byte{1}
		c, e := s.Handshake(&fixtureConn{input: bytes.NewReader(wire), writeN: -1}, &capture)
		if e != nil || c.HiUDG1Fl != aes || capture != nil || len(c.ENEXKXNftX) != 96 || !bytes.Equal(c.ENEXKXNftX[64:], secret) {
			t.Fatal("resume cipher/capture", e)
		}
		if _, e = s.Handshake(&fixtureConn{input: bytes.NewReader(wire), writeN: -1}, nil); e == nil || e.Error() != "replay detected" {
			t.Fatal("replay", e)
		}
		delete(s.GIHdzaFGhi8, [16]byte(ticket))
		unknown := &fixtureConn{input: bytes.NewReader(wire), writeN: -1}
		if _, e = s.Handshake(unknown, nil); e == nil || e.Error() != "expired ticket" || len(unknown.written) != 1 || len(unknown.written[0]) < 1279 || len(unknown.written[0]) >= 2279 {
			t.Fatal("expired ticket", e)
		}
		if _, e = QkDUgv(unknown.written[0]); e == nil {
			t.Fatal("expired response must not look like valid record")
		}
		s.MOy_iE1vq = 0
		if _, e = s.Handshake(&fixtureConn{input: bytes.NewReader(wire)}, nil); e == nil || e.Error() != "0-RTT is not allowed" {
			t.Fatal(e)
		}
		corrupt := bytes.Clone(wire)
		corrupt[65] ^= 1
		capture = nil
		if _, e = s.Handshake(&fixtureConn{input: bytes.NewReader(corrupt)}, &capture); e == nil || !bytes.Equal(capture, corrupt[:66]) {
			t.Fatal("failed auth capture", e)
		}
	}
	peer, _ := ecdh.X25519().GenerateKey(rand.Reader)
	shared, _ := peer.ECDH(key.PublicKey())
	prefix := append(bytes.Repeat([]byte{33}, 16), peer.PublicKey().Bytes()...)
	var length [2]byte
	binary.BigEndian.PutUint16(length[:], 100)
	wire := WrnUSnC6(prefix[:16], shared, true).Seal(prefix, nil, length[:], nil)
	if _, e := s.Handshake(&fixtureConn{input: bytes.NewReader(wire)}, nil); e == nil || e.Error() != "too short length" {
		t.Fatal(e)
	}
	wire = bytes.Clone(prefix)
	wire[47] |= 128
	if _, e := s.Handshake(&fixtureConn{input: bytes.NewReader(wire)}, nil); e == nil || e.Error() != "the highest bit of the last byte of the peer-sent X25519 public key is not 0" {
		t.Fatal(e)
	}
}
