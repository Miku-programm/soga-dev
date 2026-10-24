package muxerr

import (
	"bytes"
	"context"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/net/limit"
	"recovered.local/soga/config"
	"recovered.local/soga/user"
	"recovered.local/soga/crypto/ssmethod"
	"testing"
	"time"
	"unsafe"
)

var _ sess.Ki8urSHTx44 = (*BgeR1vhZZG)(nil)

func legacyUser(t *testing.T, method ssmethod.C6q3Upu) *user.WtK9HQSy66 {
	t.Helper()
	u, e := user.MbpJoYH(&types.Aga91y{WZ89pTAOrHK4: 7, JAfXWwzYrW: &types.YlfOV0DBBnx{T1oDDIgYJHh: method, A_wNUTkMh6pH: "original-wire-fixture"}}, nil)
	if e != nil {
		t.Fatal(e)
	}
	return u
}
func frame(a cipher.AEAD, nonce, p []byte) []byte {
	var l [2]byte
	binary.BigEndian.PutUint16(l[:], uint16(len(p)))
	wire := a.Seal(nil, nonce, l[:], nil)
	incrementNonce(nonce)
	wire = append(wire, a.Seal(nil, nonce, p, nil)...)
	incrementNonce(nonce)
	return wire
}
func readFrame(r io.Reader, a cipher.AEAD, nonce []byte) ([]byte, error) {
	l := make([]byte, 2+a.Overhead())
	if _, e := io.ReadFull(r, l); e != nil {
		return nil, e
	}
	l, e := a.Open(l[:0], nonce, l, nil)
	incrementNonce(nonce)
	if e != nil {
		return nil, e
	}
	p := make([]byte, int(binary.BigEndian.Uint16(l))+a.Overhead())
	if _, e = io.ReadFull(r, p); e != nil {
		return nil, e
	}
	p, e = a.Open(p[:0], nonce, p, nil)
	incrementNonce(nonce)
	return p, e
}
func TestNativeLegacyHandshakeAndSplitFrames(t *testing.T) {
	for _, method := range []ssmethod.C6q3Upu{"aes-128-gcm", "aes-256-gcm", "chacha20-ietf-poly1305"} {
		t.Run(string(method), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			u := legacyUser(t, method)
			users := &user.ZpNFqUVZp{GTFAWF: new(genmap.GTFAWF[int64, user.PgaNd3014U])}
			users.Add(u)
			candidates := sess.YCcacra(ctx)
			invalid := limit.Azb3Zr(ctx, &config.Nl62Lnd_{})
			listener, e := net.Listen("tcp", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			defer listener.Close()
			done := make(chan error, 1)
			response := bytes.Repeat([]byte{0x6d}, 40000)
			go func() {
				raw, e := listener.Accept()
				if e != nil {
					done <- e
					return
				}
				defer raw.Close()
				raw.SetDeadline(time.Now().Add(5 * time.Second))
				c, e := VLnoIxcGjx(raw, method, candidates, invalid, users)
				if e != nil {
					done <- e
					return
				}
				if c.GetUser() != u || c.GetDestAddr().GetHost() != "example.org" || c.GetDestAddr().GetPort() != 443 {
					done <- fmt.Errorf("target/user")
					return
				}
				b, e := c.ReadBuffer()
				if e != nil || string(b.Bytes()) != "first" {
					done <- fmt.Errorf("buffer %v %q", e, b.Bytes())
					return
				}
				first := make([]byte, 5)
				if _, e = io.ReadFull(c, first); e != nil || string(first) != "first" {
					done <- fmt.Errorf("read retained buffer %v", e)
					return
				}
				p := make([]byte, 4)
				if _, e = io.ReadFull(c, p); e != nil || string(p) != "next" {
					done <- fmt.Errorf("next %v", e)
					return
				}
				n, e := c.Write(response)
				if e != nil || n != len(response) {
					done <- fmt.Errorf("response %d %v", n, e)
					return
				}
				done <- nil
			}()
			client, e := net.Dial("tcp", listener.Addr().String())
			if e != nil {
				t.Fatal(e)
			}
			defer client.Close()
			client.SetDeadline(time.Now().Add(5 * time.Second))
			salt := make([]byte, len(u.PSK()))
			rand.Read(salt)
			a, e := u.NewAEAD(salt)
			if e != nil {
				t.Fatal(e)
			}
			nonce := make([]byte, 12)
			payload := append([]byte{3, 11}, []byte("example.org")...)
			payload = append(payload, 1, 187)
			payload = append(payload, []byte("first")...)
			wire := append(salt, frame(a, nonce, payload)...)
			wire = append(wire, frame(a, nonce, []byte("next"))...)
			if _, e = client.Write(wire); e != nil {
				t.Fatal(e)
			}
			serverSalt := make([]byte, len(salt))
			if _, e = io.ReadFull(client, serverSalt); e != nil {
				t.Fatal(e)
			}
			a, e = u.NewAEAD(serverSalt)
			if e != nil {
				t.Fatal(e)
			}
			clear(nonce)
			var got []byte
			for len(got) < len(response) {
				p, e := readFrame(client, a, nonce)
				if e != nil {
					t.Fatal(e)
				}
				if len(p) != 10000 {
					t.Fatalf("original recursive splitting size %d", len(p))
				}
				got = append(got, p...)
			}
			if !bytes.Equal(got, response) {
				t.Fatal("response mismatch")
			}
			if e = <-done; e != nil {
				t.Fatal(e)
			}
		})
	}
}
func TestNativeLegacyLayoutAndCorruptHeader(t *testing.T) {
	var c BgeR1vhZZG
	if unsafe.Sizeof(c) != 200 || unsafe.Offsetof(c.f4ApgxW4hCCA) != 168 {
		t.Fatal("RTTI")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	u := legacyUser(t, "aes-128-gcm")
	users := &user.ZpNFqUVZp{GTFAWF: new(genmap.GTFAWF[int64, user.PgaNd3014U])}
	users.Add(u)
	l, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer l.Close()
	done := make(chan error, 1)
	go func() {
		raw, e := l.Accept()
		if e != nil {
			done <- e
			return
		}
		defer raw.Close()
		raw.SetDeadline(time.Now().Add(time.Second))
		_, e = VLnoIxcGjx(raw, "aes-128-gcm", sess.YCcacra(ctx), limit.Azb3Zr(ctx, &config.Nl62Lnd_{}), users)
		done <- e
	}()
	conn, e := net.Dial("tcp", l.Addr().String())
	if e != nil {
		t.Fatal(e)
	}
	defer conn.Close()
	b := make([]byte, 34)
	rand.Read(b)
	conn.Write(b)
	if e = <-done; e != errorAt5098410 {
		t.Fatalf("bad authentication: %v", e)
	}
}
