package mXe7cPU8Mzi

import (
	"bytes"
	"context"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"net"
	D "recovered.local/soga/D_zCag1qTV"
	Map "recovered.local/soga/G7A7csgrs62p"
	N "recovered.local/soga/NtxdAE"
	U "recovered.local/soga/nzATvajt"
	"testing"
	"time"
	"unsafe"
)

// This client fixture builds raw metadata, key material and AEAD frames without
// calling any recovered writer/reader, parser, padding or nonce helper.
type mieruFixtureClient struct {
	c             net.Conn
	a             cipher.AEAD
	send, receive []byte
	first         bool
}

func newMieruFixtureClient(t *testing.T, c net.Conn, u *U.WIZQ28) *mieruFixtureClient {
	t.Helper()
	secret := sha256.Sum256([]byte(u.GetPassword() + "\x00" + u.GetUsername()))
	var stamp [8]byte
	binary.BigEndian.PutUint64(stamp[:], uint64(time.Now().Round(2*time.Minute).Unix()))
	salt := sha256.Sum256(stamp[:])
	key := pbkdf2.Key(secret[:], salt[:], 64, 32, sha256.New)
	a, _ := chacha20poly1305.NewX(key)
	nonce := make([]byte, 24)
	rand.Read(nonce)
	sum := sha256.Sum256(append([]byte(u.GetUsername()), nonce[:16]...))
	copy(nonce[20:], sum[:4])
	return &mieruFixtureClient{c: c, a: a, send: nonce, first: true}
}
func fixtureNext(n []byte) {
	for i := len(n) - 1; i >= 0; i-- {
		n[i]++
		if n[i] != 0 {
			break
		}
	}
}
func (f *mieruFixtureClient) write(kind byte, id, seq uint32, payload []byte) error {
	h := make([]byte, 32)
	h[0] = kind
	binary.BigEndian.PutUint32(h[2:6], uint32(time.Now().Unix()/60))
	binary.BigEndian.PutUint32(h[6:10], id)
	binary.BigEndian.PutUint32(h[10:14], seq)
	if kind < 6 {
		binary.BigEndian.PutUint16(h[15:17], uint16(len(payload)))
	} else {
		binary.BigEndian.PutUint16(h[22:24], uint16(len(payload)))
	}
	var b []byte
	if f.first {
		b = append(b, f.send...)
		f.first = false
	}
	b = append(b, f.a.Seal(nil, f.send, h, nil)...)
	fixtureNext(f.send)
	if len(payload) > 0 {
		b = append(b, f.a.Seal(nil, f.send, payload, nil)...)
		fixtureNext(f.send)
	}
	_, err := f.c.Write(b)
	return err
}
func (f *mieruFixtureClient) read() (byte, uint32, []byte, error) {
	if f.receive == nil {
		f.receive = make([]byte, 24)
		if _, e := io.ReadFull(f.c, f.receive); e != nil {
			return 0, 0, nil, e
		}
	}
	b := make([]byte, 48)
	if _, e := io.ReadFull(f.c, b); e != nil {
		return 0, 0, nil, e
	}
	h, e := f.a.Open(nil, f.receive, b, nil)
	if e != nil {
		return 0, 0, nil, e
	}
	fixtureNext(f.receive)
	n, suffix, prefix := int(binary.BigEndian.Uint16(h[15:17])), int(h[17]), 0
	if h[0] >= 6 {
		n, suffix, prefix = int(binary.BigEndian.Uint16(h[22:24])), int(h[24]), int(h[21])
	}
	if _, e = io.CopyN(io.Discard, f.c, int64(prefix)); e != nil {
		return 0, 0, nil, e
	}
	var p []byte
	if n > 0 {
		b = make([]byte, n+16)
		if _, e = io.ReadFull(f.c, b); e != nil {
			return 0, 0, nil, e
		}
		p, e = f.a.Open(nil, f.receive, b, nil)
		if e != nil {
			return 0, 0, nil, e
		}
		fixtureNext(f.receive)
	}
	if _, e = io.CopyN(io.Discard, f.c, int64(suffix)); e != nil {
		return 0, 0, nil, e
	}
	return h[0], binary.BigEndian.Uint32(h[6:10]), p, nil
}
func TestNativeMultiplexedConnectionWire(t *testing.T) {
	if unsafe.Sizeof(WMK1ddta3da{}) != 96 || unsafe.Sizeof(xjTVP8{}) != 184 || unsafe.Sizeof(zXJm01kG7F4E{}) != 32 {
		t.Fatal("RTTI")
	}
	u := testMieruUser(t, 510, "native-user", "secret")
	users := &U.ZpNFqUVZp{GTFAWF: &Map.GTFAWF[int64, U.PgaNd3014U]{}}
	users.Add(u)
	server, client := net.Pipe()
	defer client.Close()
	defer server.Close()
	client.SetDeadline(time.Now().Add(5 * time.Second))
	done := make(chan error, 1)
	go func() {
		s, err := M_03miDCRcqa(context.Background(), server, &N.O0UldNL4Eq{RacA3he0TX: users})
		if err != nil {
			done <- err
			return
		}
		defer s.Close()
		for i := 0; i < 2; i++ {
			accepted, err := s.Next()
			if err != nil {
				done <- err
				return
			}
			c, err := accepted.Accept()
			if err != nil {
				done <- err
				return
			}
			if c.GetDestAddr().Join() != "127.0.0.1:8080" {
				done <- errors.New("destination")
				return
			}
			b := make([]byte, 4)
			if _, err = io.ReadFull(c, b); err != nil {
				done <- err
				return
			}
			if _, err = c.Write(append([]byte("echo:"), b...)); err != nil {
				done <- err
				return
			}
			if err = c.Close(); err != nil {
				done <- err
				return
			}
			accepted.j_Ra2Blepbh.Free()
		}
		done <- nil
	}()
	f := newMieruFixtureClient(t, client, u)
	for _, id := range []uint32{101, 202} {
		request := []byte{5, 1, 0, 1, 127, 0, 0, 1, 0x1f, 0x90}
		if err := f.write(2, id, 0, request); err != nil {
			t.Fatal(err)
		}
		kind, sid, p, err := f.read()
		if err != nil || kind != 3 || sid != id || len(p) != 0 {
			t.Fatalf("open %d %d %x %v", kind, sid, p, err)
		}
		kind, sid, p, err = f.read()
		if err != nil || kind != 7 || sid != id || !bytes.Equal(p, []byte{5, 0, 0, 1, 0, 0, 0, 0, 0, 0}) {
			t.Fatalf("accept %d %d %x %v", kind, sid, p, err)
		}
		if err = f.write(6, id, 1, []byte("data")); err != nil {
			t.Fatal(err)
		}
		kind, sid, p, err = f.read()
		if err != nil || kind != 7 || sid != id || string(p) != "echo:data" {
			t.Fatalf("echo %d %d %x %v", kind, sid, p, err)
		}
		kind, sid, _, err = f.read()
		if err != nil || kind != 4 || sid != id {
			t.Fatalf("close %d %d %v", kind, sid, err)
		}
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}
func TestStreamResidualDeadlineAndClose(t *testing.T) {
	s := vnqzugKQF1jX(1, nil, new(frameMemoryConn), nil)
	b := D.PoolAt50961D8.Get(5)
	copy(b, "hello")
	s.xfdnp9xA <- b
	out := make([]byte, 2)
	if n, e := s.Read(out); n != 2 || e != nil || string(out) != "he" {
		t.Fatal(n, e)
	}
	s.SetReadDeadline(time.Now().Add(-time.Second))
	s.Close()
	out = make([]byte, 9)
	if n, e := s.Read(out); n != 3 || e != nil || string(out[:n]) != "llo" {
		t.Fatal(n, e)
	}
	if _, e := s.Read(out); e == nil {
		t.Fatal("timeout")
	} else if timeout, ok := e.(net.Error); !ok || !timeout.Timeout() {
		t.Fatal(e)
	}
	s.SetReadDeadline(time.Time{})
	if _, e := s.Read(out); e != io.EOF {
		t.Fatal(e)
	}
	s.Free()
}
