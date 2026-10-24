package ssudp

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"io"
	"net"
	"os"
	"path/filepath"
	"recovered.local/soga/util/genmap"
	"recovered.local/soga/protocol/ss2022"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/net/limit"
	"recovered.local/soga/transport/muxerr"
	"recovered.local/soga/config"
	"recovered.local/soga/user"
	"recovered.local/soga/crypto/ssmethod"
	"testing"
	"time"
	"unsafe"
)

type wireConn struct{ *bytes.Reader }

func (c *wireConn) Write(p []byte) (int, error)      { return len(p), nil }
func (c *wireConn) Close() error                     { return nil }
func (c *wireConn) LocalAddr() net.Addr              { return &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1} }
func (c *wireConn) RemoteAddr() net.Addr             { return &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 2} }
func (c *wireConn) SetDeadline(time.Time) error      { return nil }
func (c *wireConn) SetReadDeadline(time.Time) error  { return nil }
func (c *wireConn) SetWriteDeadline(time.Time) error { return nil }
func len16(p []byte) []byte                          { return append([]byte{byte(len(p) >> 8), byte(len(p))}, p...) }
func tlsHello(payload []byte) []byte {
	h := append([]byte{1, 0, 0, 0, 3, 3}, make([]byte, 32)...)
	h = append(h, 0, 0, 2, 0xcc, 0xa8, 1, 0)
	names := len16(append([]byte{0}, len16([]byte("example.org"))...))
	ext := append([]byte{0, 0}, len16(names)...)
	ext = append(ext, 0, 35)
	ext = append(ext, len16(payload)...)
	h = append(h, len16(ext)...)
	return append([]byte{22, 3, 3}, len16(h)...)
}
func TestNativeSSFactoryObfsAndTypedFailure(t *testing.T) {
	var stream StreamOptions
	var packet PacketOptions
	if unsafe.Sizeof(stream) != 120 || unsafe.Offsetof(stream.Manager) != 112 || unsafe.Sizeof(packet) != 64 {
		t.Fatal("original ABI aggregate")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	u, err := user.MbpJoYH(&types.Aga91y{WZ89pTAOrHK4: 7, JAfXWwzYrW: &types.YlfOV0DBBnx{T1oDDIgYJHh: "aes-128-gcm", A_wNUTkMh6pH: "factory-test"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	users := &user.ZpNFqUVZp{GTFAWF: new(genmap.GTFAWF[int64, user.PgaNd3014U])}
	users.Add(u)
	o := StreamOptions{Cipher: "aes-128-gcm", Mode: 1, Path: "abc", Host: "example.org", Candidates: sess.YCcacra(ctx), Invalid: limit.Azb3Zr(ctx, &config.Nl62Lnd_{}), Users: users}
	for _, obfs := range []string{"", "HTTP", "http", "tls"} {
		salt := make([]byte, 16)
		rand.Read(salt)
		a, err := u.NewAEAD(salt)
		if err != nil {
			t.Fatal(err)
		}
		plain := append([]byte{3, 11}, []byte("example.org")...)
		plain = append(plain, 1, 187)
		plain = append(plain, []byte("data")...)
		var l [2]byte
		binary.BigEndian.PutUint16(l[:], uint16(len(plain)))
		nonce := make([]byte, 12)
		wire := append(salt, a.Seal(nil, nonce, l[:], nil)...)
		nonce[0] = 1
		wire = append(wire, a.Seal(nil, nonce, plain, nil)...)
		switch obfs {
		case "http":
			wire = append([]byte("GET /abc HTTP/1.1\r\nHost: example.org\r\n\r\n"), wire...)
		case "tls":
			wire = tlsHello(wire)
		}
		o.Obfs = obfs
		o.Conn = &wireConn{bytes.NewReader(wire)}
		c, err := FciVp4yFHL(o)
		if err != nil {
			t.Fatal(obfs, err)
		}
		if c.GetUser() != u || c.GetDestAddr().GetHost() != "example.org" || c.GetDestAddr().GetPort() != 443 {
			t.Fatal("factory target/user", obfs)
		}
		p := make([]byte, 4)
		if _, err := io.ReadFull(c, p); err != nil || string(p) != "data" {
			t.Fatal(obfs, string(p), err)
		}
		c.(*muxerr.BgeR1vhZZG).Free()
	}
	o.Obfs = ""
	o.Conn = &wireConn{bytes.NewReader(nil)}
	c, err := FciVp4yFHL(o)
	if err == nil || c == nil || c.(*muxerr.BgeR1vhZZG) != nil {
		t.Fatal("legacy failure loses original typed nil", c, err)
	}
	for _, method := range []ssmethod.C6q3Upu{"2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm"} {
		o.Cipher = method
		n, _ := method.KeyLen()
		o.Manager, err = ss2022.DkAuKfGU(ctx, &types.AAMPrdcYTL6{HMkc4BgR: method, VuhWqou: base64.StdEncoding.EncodeToString(make([]byte, n))})
		if err != nil {
			t.Fatal(err)
		}
		c, err = FciVp4yFHL(o)
		if c != nil || err == nil {
			t.Fatal("2022 failure interface", c, err)
		}
		if _, ok := CXui61CuSLa(PacketOptions{Cipher: method}).(*ss2022.KNAU41ESnn); !ok {
			t.Fatal("2022 packet factory")
		}
	}
	if _, ok := CXui61CuSLa(PacketOptions{Cipher: "aes-128-gcm"}).(*muxerr.SUfEmBl); !ok {
		t.Fatal("legacy packet factory")
	}
}

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "soga-native-factory-test-")
	if err != nil {
		panic(err)
	}
	path := filepath.Join(dir, "soga.conf")
	if err = os.WriteFile(path, []byte("udp_timeout=2\nlog_file=false\naccess_log_enable=false\n"), 0600); err != nil {
		panic(err)
	}
	config.MkwNWciPHS(path)
	if err = config.TvBECnD9fM6(); err != nil {
		panic(err)
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}
