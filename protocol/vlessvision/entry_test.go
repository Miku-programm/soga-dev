package vlessvision

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"errors"
	"io"
	"math/big"
	"net"
	"reflect"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/transport/grpcproto"
	"recovered.local/soga/net/sess"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/protocol/vless"
	"recovered.local/soga/user"
	"recovered.local/soga/reality"
)

type entryMemoryConn struct {
	net.Conn
	r      io.Reader
	wire   bytes.Buffer
	closed int
}

func (c *entryMemoryConn) Read(p []byte) (int, error)  { return c.r.Read(p) }
func (c *entryMemoryConn) Write(p []byte) (int, error) { return c.wire.Write(p) }
func (c *entryMemoryConn) Close() error                { c.closed++; return nil }
func originalUser(t *testing.T) *user.A6CvRbtca {
	t.Helper()
	u, e := user.Sa_NiNyuJ_ax(&types.Aga91y{WZ89pTAOrHK4: 41, JAfXWwzYrW: &types.IImIwacp{Dh5iiakU: "00112233-4455-6677-8899-aabbccddeeff"}}, nil)
	if e != nil {
		t.Fatal(e)
	}
	return u
}
func entryWire(t *testing.T, flow string, cmd byte) []byte {
	t.Helper()
	addon, e := proto.Marshal(&grpcproto.N_SQ5a{AfJaYuHm3H: flow})
	if e != nil {
		t.Fatal(e)
	}
	out := append([]byte{byte(len(addon))}, addon...)
	out = append(out, cmd)
	if cmd != 3 {
		out = append(out, 1, 187, 1, 127, 0, 0, 1)
	}
	return out
}
func countedConn(c net.Conn, u *user.A6CvRbtca) *sess.Mdn09F6aMrpa {
	x := &sess.Mdn09F6aMrpa{DrpXao: c}
	x.SetUser(u)
	return x
}

func TestEntryOriginalLayouts(t *testing.T) {
	for _, x := range []struct {
		typ     reflect.Type
		size    uintptr
		offsets []uintptr
	}{{reflect.TypeFor[J3LOqQMzo](), 88, []uintptr{0, 16, 32, 48, 64, 72, 80}}, {reflect.TypeFor[LZ7AGw7](), 80, []uintptr{0, 16, 24, 32, 40, 48, 49, 56, 72}}, {reflect.TypeFor[Qx58rV](), 72, []uintptr{0, 16, 24, 32, 56}}} {
		if x.typ.Size() != x.size {
			t.Fatal(x.typ, x.typ.Size())
		}
		for i, off := range x.offsets {
			if x.typ.Field(i).Offset != off {
				t.Fatal(x.typ, x.typ.Field(i))
			}
		}
	}
}

func TestEntryPlainTCPAndUDP(t *testing.T) {
	for _, cmd := range []byte{1, 2, 3} {
		u := originalUser(t)
		header := entryWire(t, "", cmd)
		payload := []byte("hello")
		if cmd == 2 {
			payload = append([]byte{0, 5}, payload...)
		}
		raw := &entryMemoryConn{r: bytes.NewReader(append(header, payload...))}
		c, e := V7ZRKYAK(countedConn(raw, u), "", 0)
		if e != nil {
			t.Fatal(e)
		}
		if c.GetUser() != u || c.NetworkType() != sess.ZES6F5(cmd-1) || c.GetRawConn() == nil || c.GetHeader() == nil {
			t.Fatal("entry accessors")
		}
		if cmd != 3 && c.GetDestAddr().Join() != "127.0.0.1:443" {
			t.Fatal("destination")
		}
		buf := make([]byte, 5)
		if _, e = io.ReadFull(c, buf); e != nil || string(buf) != "hello" {
			t.Fatal(string(buf), e)
		}
		if n, e := c.Write([]byte("reply")); n != 5 || e != nil {
			t.Fatal(n, e)
		}
		want := []byte{0, 0}
		if cmd == 2 {
			want = append(want, 0, 5)
		}
		want = append(want, []byte("reply")...)
		if !bytes.Equal(raw.wire.Bytes(), want) {
			t.Fatalf("wire %x", raw.wire.Bytes())
		}
		up, down := u.GetAndClearTraffic()
		if up != int64(len(header)+len(payload)) || down != int64(len(want)) {
			t.Fatalf("traffic %d %d", up, down)
		}
		if e = c.Close(); e != nil || raw.closed != 1 {
			t.Fatal("close", e)
		}
	}
}

func TestEntryFlowRejectionsAndMissingUUID(t *testing.T) {
	for _, tc := range []struct {
		flow, allowed string
		cmd           byte
		want          string
	}{{"", "xtls-rprx-vision", 1, "can not use empty flow"}, {"xtls-rprx-vision", "", 1, "can not able to use flow: xtls-rprx-vision"}, {"unsupported", "", 1, "unknown vless flow unsupported"}, {"xtls-rprx-vision", "xtls-rprx-vision", 1, "XTLS only supports TLS and REALITY directly for now"}} {
		raw := &entryMemoryConn{r: bytes.NewReader(entryWire(t, tc.flow, tc.cmd))}
		if c, e := V7ZRKYAK(countedConn(raw, originalUser(t)), tc.allowed, 0); c != nil || e == nil || e.Error() != tc.want {
			t.Fatalf("%v %v", c, e)
		}
	}
	u := originalUser(t)
	empty := &user.A6CvRbtca{AhSFeSua: u.AhSFeSua}
	raw := &entryMemoryConn{r: bytes.NewReader(entryWire(t, "", 3))}
	if _, e := V7ZRKYAK(countedConn(raw, empty), "", 0); e == nil || e.Error() != "failed to get user 41 uuid" {
		t.Fatal(e)
	}
}

func TestExactUnwrapOrder(t *testing.T) {
	raw := &entryMemoryConn{r: bytes.NewReader(nil)}
	tlsConn := tls.Client(raw, &tls.Config{})
	realityConn := reality.Client(raw, &reality.Config{})
	for _, v := range []net.Conn{tlsConn, realityConn, &vless.RfBfjVN4XE{DrpXao: raw}} {
		if got := CocBnud6(&sess.Mdn09F6aMrpa{DrpXao: v}); got != raw {
			t.Fatalf("unwrap %T = %T", v, got)
		}
	}
	outer := &vless.RfBfjVN4XE{DrpXao: tlsConn}
	if got := CocBnud6(&sess.Mdn09F6aMrpa{DrpXao: outer}); got != tlsConn {
		t.Fatal("record wrapper branch must not also unwrap TLS")
	}
	ctr := &vless.LUKUT5W0R6{DrpXao: raw}
	if got := CocBnud6(&sess.Mdn09F6aMrpa{DrpXao: ctr}); got != ctr {
		t.Fatal("CTR must remain")
	}
}

func TestVisionRawSwitchDrainsBuffersAndCountsOnce(t *testing.T) {
	u := originalUser(t)
	raw := &entryMemoryConn{r: bytes.NewReader([]byte("network"))}
	counted := countedConn(raw, u)
	state := &MxS0kZ{SDw6IIM1D51U: true, Pd9MU7P6IX: true}
	r := &LZ7AGw7{fB4afneh7k: counted, woJdOhbWO: state, sztBlp: bytes.NewReader([]byte("plain")), k9rmTO: bytes.NewBufferString("pending"), oqHPkr: pool.S4OQXWIZ(0)}
	p := make([]byte, 32)
	for _, want := range []string{"plain", "pending", "network"} {
		n, e := r.Read(p)
		if e != nil || string(p[:n]) != want {
			t.Fatalf("%q %v", p[:n], e)
		}
	}
	w := &Qx58rV{uYYP2p: counted, taFXeH: state}
	if n, e := w.Write([]byte("out")); n != 3 || e != nil {
		t.Fatal(n, e)
	}
	up, down := u.GetAndClearTraffic()
	if up != 7 || down != 3 {
		t.Fatal(up, down)
	}
	if r.oqHPkr != nil {
		t.Fatal("pool buffer not freed on raw switch")
	}
	c := &J3LOqQMzo{DrpXao: counted, eJeaO7LW: state}
	if e := c.Close(); e != nil || raw.closed != 1 {
		t.Fatal(e)
	}
}

func tlsFixture(t *testing.T, version uint16) (*tls.Conn, *tls.Conn) {
	t.Helper()
	key, e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if e != nil {
		t.Fatal(e)
	}
	cert := &x509.Certificate{SerialNumber: big.NewInt(1), NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature}
	der, e := x509.CreateCertificate(rand.Reader, cert, cert, &key.PublicKey, key)
	if e != nil {
		t.Fatal(e)
	}
	a, b := net.Pipe()
	a.SetDeadline(time.Now().Add(3 * time.Second))
	b.SetDeadline(time.Now().Add(3 * time.Second))
	t.Cleanup(func() { a.Close(); b.Close() })
	s := tls.Server(a, &tls.Config{MinVersion: version, MaxVersion: version, Certificates: []tls.Certificate{{Certificate: [][]byte{der}, PrivateKey: key}}})
	c := tls.Client(b, &tls.Config{InsecureSkipVerify: true, MinVersion: version, MaxVersion: version})
	return s, c
}

func TestVisionTLS13EntryRoundTrip(t *testing.T) {
	server, client := tlsFixture(t, tls.VersionTLS13)
	u := originalUser(t)
	id, _ := u.GetUuid()
	frame := pool.S4OQXWIZ(0)
	prefix := bytes.Clone(id[:])
	IFq7RHaxl(frame, []byte("query"), 0, &prefix, false)
	wire := append(entryWire(t, "xtls-rprx-vision", 1), frame.Bytes()...)
	frame.Free()
	done := make(chan error, 1)
	go func() {
		_, e := client.Write(wire)
		if e != nil {
			done <- e
			return
		}
		head := make([]byte, 23)
		if _, e = io.ReadFull(client, head); e != nil {
			done <- e
			return
		}
		if !bytes.Equal(head[:2], []byte{0, 0}) || !bytes.Equal(head[2:18], id[:]) {
			done <- errors.New("VLESS response or Vision UUID differs")
			return
		}
		length := int(binary.BigEndian.Uint16(head[19:21]))
		padding := int(binary.BigEndian.Uint16(head[21:23]))
		body := make([]byte, length+padding)
		if _, e = io.ReadFull(client, body); e == nil && string(body[:length]) != "reply" {
			e = errors.New("Vision reply differs")
		}
		done <- e
	}()
	c, e := V7ZRKYAK(countedConn(server, u), "xtls-rprx-vision", 0)
	if e != nil {
		t.Fatal(e)
	}
	defer c.Free()
	p := make([]byte, 5)
	if _, e = io.ReadFull(c, p); e != nil || string(p) != "query" {
		t.Fatal(string(p), e)
	}
	if n, e := c.Write([]byte("reply")); n != 5 || e != nil {
		t.Fatal(n, e)
	}
	if e := <-done; e != nil {
		t.Fatal(e)
	}
}

func TestVisionTLS12Rejected(t *testing.T) {
	server, client := tlsFixture(t, tls.VersionTLS12)
	done := make(chan error, 1)
	go func() { _, e := client.Write(entryWire(t, "xtls-rprx-vision", 1)); done <- e }()
	_, e := V7ZRKYAK(countedConn(server, originalUser(t)), "xtls-rprx-vision", 0)
	if e == nil || e.Error() != "failed to use xtls-rprx-vision, found outer tls version 771" {
		t.Fatal(e)
	}
	if e = <-done; e != nil {
		t.Fatal(e)
	}
}
