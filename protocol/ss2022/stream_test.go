package ss2022

import (
	"bytes"
	"context"
	"crypto/aes"
	"encoding/binary"
	"io"
	"lukechampine.com/blake3"
	"net"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/user"
	"recovered.local/soga/crypto/ssmethod"
	"testing"
	"time"
)

func initialFixture(t *testing.T, m *N07NaA, u *user.WtK9HQSy66, kind byte, stamp uint64, payload []byte) ([]byte, []byte) {
	t.Helper()
	salt := bytes.Repeat([]byte{0x71}, len(u.PSK()))
	identity := make([]byte, len(salt))
	material := append(append([]byte{}, m.GetPSK()...), salt...)
	blake3.DeriveKey(identity, "shadowsocks 2022 identity subkey", material)
	block, err := aes.NewCipher(identity)
	if err != nil {
		t.Fatal(err)
	}
	hash := blake3.Sum512(u.PSK())
	encrypted := make([]byte, 16)
	block.Encrypt(encrypted, hash[:16])
	variable := append([]byte{3, 11}, []byte("example.org")...)
	variable = append(variable, 0, 80, 0, 3, 1, 2, 3)
	variable = append(variable, payload...)
	fixed := make([]byte, 11)
	fixed[0] = kind
	binary.BigEndian.PutUint64(fixed[1:], stamp)
	binary.BigEndian.PutUint16(fixed[9:], uint16(len(variable)))
	a, err := u.GetCipher().NewAEAD(u.GetCipher().SubKey(u.PSK(), salt))
	if err != nil {
		t.Fatal(err)
	}
	nonce := make([]byte, 12)
	wire := append(append([]byte{}, salt...), encrypted...)
	wire = append(wire, a.Seal(nil, nonce, fixed, nil)...)
	nonce[0]++
	wire = append(wire, a.Seal(nil, nonce, variable, nil)...)
	return wire, salt
}
func TestOriginalNativeTCPHandshakeFramesAndResponse(t *testing.T) {
	for _, method := range []ssmethod.C6q3Upu{"2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm"} {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		m := managerFixture(t, ctx, method)
		u := userFixture(t, method, 7, 9)
		m.UpdateUser(u)
		wire, salt := initialFixture(t, m, u, 0, uint64(rtx.C5Jff5vs0()), []byte("initial"))
		client, server := net.Pipe()
		defer client.Close()
		defer server.Close()
		client.SetDeadline(time.Now().Add(5 * time.Second))
		done := make(chan error, 1)
		go func() {
			c, err := YZbzDDrJX(server, m)
			if err != nil {
				done <- err
				return
			}
			stream := c.(*DjXTT7iwk)
			defer stream.Free()
			if c.GetDestAddr().GetHost() != "example.org" || c.GetDestAddr().GetPort() != 80 || c.GetUser() != u {
				done <- io.ErrUnexpectedEOF
				return
			}
			first, err := stream.ReadBuffer()
			if err != nil {
				done <- err
				return
			}
			if string(first.Bytes()) != "initial" {
				done <- io.ErrUnexpectedEOF
				return
			}
			if _, err = stream.Write([]byte("response")); err != nil {
				done <- err
				return
			}
			b := make([]byte, 3)
			n, err := stream.Read(b)
			if err != nil {
				done <- err
				return
			}
			if n != 3 || string(b) != "ini" {
				done <- io.ErrUnexpectedEOF
				return
			}
			b = make([]byte, 4)
			if _, err = io.ReadFull(stream, b); err != nil || string(b) != "tial" {
				done <- io.ErrUnexpectedEOF
				return
			}
			for _, size := range []int{32, 4, 2} {
				b = make([]byte, size)
				n, err = stream.Read(b)
				if err != nil {
					done <- err
					return
				}
				if string(b[:n]) != "next"[:min(size, 4)] {
					done <- io.ErrUnexpectedEOF
					return
				}
				if size == 2 {
					tail := make([]byte, 2)
					_, err = io.ReadFull(stream, tail)
					if err != nil || string(tail) != "xt" {
						done <- io.ErrUnexpectedEOF
						return
					}
				}
			}
			done <- nil
		}()
		if _, err := client.Write(wire); err != nil {
			t.Fatal(err)
		}
		responseSalt := make([]byte, len(salt))
		if _, err := io.ReadFull(client, responseSalt); err != nil {
			t.Fatal(err)
		}
		a, err := method.NewAEAD(method.SubKey(u.PSK(), responseSalt))
		if err != nil {
			t.Fatal(err)
		}
		sealed := make([]byte, 11+len(salt)+16)
		if _, err = io.ReadFull(client, sealed); err != nil {
			t.Fatal(err)
		}
		nonce := make([]byte, 12)
		header, err := a.Open(nil, nonce, sealed, nil)
		if err != nil {
			t.Fatal(err)
		}
		if header[0] != 1 || !bytes.Equal(header[9:9+len(salt)], salt) || binary.BigEndian.Uint16(header[9+len(salt):]) != 8 {
			t.Fatal("response header")
		}
		nonce[0]++
		sealed = make([]byte, 24)
		if _, err = io.ReadFull(client, sealed); err != nil {
			t.Fatal(err)
		}
		plain, err := a.Open(nil, nonce, sealed, nil)
		if err != nil || string(plain) != "response" {
			t.Fatal(err)
		}
		a, err = method.NewAEAD(method.SubKey(u.PSK(), salt))
		if err != nil {
			t.Fatal(err)
		}
		nonce[0] = 2
		for i := 0; i < 3; i++ {
			frame := a.Seal(nil, nonce, []byte{0, 4}, nil)
			incrementNonce(nonce)
			frame = append(frame, a.Seal(nil, nonce, []byte("next"), nil)...)
			incrementNonce(nonce)
			if _, err = client.Write(frame); err != nil {
				t.Fatal(err)
			}
		}
		if err = <-done; err != nil {
			t.Fatal(err)
		}
	}
}

// fixtureConn 保持每次 Read 尽可能填满，独立测试原首包错误和 deadline 清理。
type fixtureConn struct {
	*bytes.Reader
	writes    bytes.Buffer
	deadlines []time.Time
}

func (c *fixtureConn) Write(p []byte) (int, error)      { return c.writes.Write(p) }
func (c *fixtureConn) Close() error                     { return nil }
func (c *fixtureConn) LocalAddr() net.Addr              { return &net.TCPAddr{} }
func (c *fixtureConn) RemoteAddr() net.Addr             { return &net.TCPAddr{} }
func (c *fixtureConn) SetDeadline(time.Time) error      { return nil }
func (c *fixtureConn) SetWriteDeadline(time.Time) error { return nil }
func (c *fixtureConn) SetReadDeadline(v time.Time) error {
	c.deadlines = append(c.deadlines, v)
	return nil
}
func TestNativeTCPAuthenticationFailures(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	m := managerFixture(t, ctx, "2022-blake3-aes-128-gcm")
	u := userFixture(t, m.GetCipher(), 3, 4)
	m.UpdateUser(u)
	for _, tt := range []struct {
		name  string
		kind  byte
		stamp uint64
		want  string
	}{{"type", 1, uint64(rtx.C5Jff5vs0()), "ss2022 header type wrong: 1"}, {"old", 0, uint64(rtx.C5Jff5vs0() - 31), "ss2022 time error"}} {
		wire, _ := initialFixture(t, m, u, tt.kind, tt.stamp, nil)
		conn := &fixtureConn{Reader: bytes.NewReader(wire)}
		_, err := YZbzDDrJX(conn, m)
		if err == nil || err.Error() != tt.want {
			t.Fatalf("%s: %v", tt.name, err)
		}
		if len(conn.deadlines) != 2 || !conn.deadlines[1].IsZero() {
			t.Fatal("deadline not cleared")
		}
	}
	wire, salt := initialFixture(t, m, u, 0, uint64(rtx.C5Jff5vs0()), nil)
	m.AddSalt(salt)
	_, err := YZbzDDrJX(&fixtureConn{Reader: bytes.NewReader(wire)}, m)
	if err == nil || err.Error() != "repeat salt" {
		t.Fatal(err)
	}
	// 原首个 Read 没有 io.ReadFull 重试；不满 salt+43 立即拒绝。
	_, err = YZbzDDrJX(&fixtureConn{Reader: bytes.NewReader(wire[:20])}, m)
	if err == nil || err.Error() != "ss2022 bad header" {
		t.Fatal(err)
	}
}

var _ sess.Ki8urSHTx44 = (*DjXTT7iwk)(nil)
