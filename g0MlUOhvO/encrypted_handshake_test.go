package g0MlUOhvO

import (
	"bytes"
	"crypto/rand"
	"io"
	"net"
	D "recovered.local/soga/DDYEn9Z_b"
	"testing"
	"time"
)

func testTunnelConfig(t *testing.T) *VTl94pzz2IkB {
	t.Helper()
	m, e := D.MakT9YbIpyNJ("aes-128-gcm")
	if e != nil {
		t.Fatal(e)
	}
	return &VTl94pzz2IkB{zkn35IN: m, oWvGY4I9cA: []byte("test-derived-secret")}
}
func TestTunnelSaltReadAndReplay(t *testing.T) {
	cfg := testTunnelConfig(t)
	salt := make([]byte, 12)
	rand.Read(salt)
	for _, tc := range []struct {
		max  int
		err  error
		want string
	}{{6, nil, "read salt len wrong"}, {0, io.EOF, "EOF"}, {0, nil, ""}, {0, nil, "repeat salt"}} {
		raw := &wireFixture{Reader: bytes.NewReader(salt), max: tc.max, readErr: tc.err}
		c := &Cb7vj3KQA8{DrpXao: raw, hp3cHCehD2pe: cfg}
		e := c.dlEStwP4xJ()
		if tc.want == "" {
			if e != nil || c.e8aahi_R3S == nil {
				t.Fatal(e)
			}
		} else if e == nil || e.Error() != tc.want || c.e8aahi_R3S != nil {
			t.Fatal(tc, e)
		}
	}
}
func TestEncryptedTunnelActualTCPBidirectional(t *testing.T) {
	cfg := testTunnelConfig(t)
	listener, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer listener.Close()
	wrapped := &T4k2R7cJ{O6B2qZ0p0: listener, tS5gsM72s: cfg}
	payload := bytes.Repeat([]byte("native tunnel "), 3500)
	reply := bytes.Repeat([]byte("reply "), 7000)
	done := make(chan error, 1)
	go func() {
		conn, err := wrapped.Accept()
		if err != nil {
			done <- err
			return
		}
		defer conn.Close()
		c := conn.(*Cb7vj3KQA8)
		defer c.Free()
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		incoming := make([]byte, len(payload))
		_, err = io.ReadFull(conn, incoming)
		if err == nil && !bytes.Equal(incoming, payload) {
			err = io.ErrNoProgress
		}
		if err == nil {
			_, err = conn.Write(reply)
		}
		done <- err
	}()
	raw, e := net.Dial("tcp", listener.Addr().String())
	if e != nil {
		t.Fatal(e)
	}
	defer raw.Close()
	raw.SetDeadline(time.Now().Add(10 * time.Second))
	client := &Cb7vj3KQA8{DrpXao: raw, hp3cHCehD2pe: cfg}
	defer client.Free()
	if n, e := client.Write(payload); n != len(payload) || e != nil {
		t.Fatal(n, e)
	}
	received := make([]byte, len(reply))
	if _, e := io.ReadFull(client, received); e != nil || !bytes.Equal(received, reply) {
		t.Fatal(e)
	}
	if e := <-done; e != nil {
		t.Fatal(e)
	}
}
