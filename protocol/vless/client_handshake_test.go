package vless

import (
	"bytes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/mlkem"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"lukechampine.com/blake3"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestClientInitAndPadding(t *testing.T) {
	var s GPFxUP0_
	if e := s.Init(nil, 0, 0, ""); e == nil || e.Error() != "empty nfsPKeysBytes" || s.QUnOciyca6eb != nil {
		t.Fatal(e)
	}
	x, _ := ecdh.X25519().GenerateKey(rand.Reader)
	kem, _ := mlkem.GenerateKey768()
	keys := [][]byte{x.PublicKey().Bytes(), kem.EncapsulationKey().Bytes()}
	if e := s.Init(keys, 2, 600, "100-35-35.100-0-0.100-7-9-extra"); e != nil {
		t.Fatal(e)
	}
	if s.JhDzI4ehWF != 1152 || s.S799IXI != 2 || s.BsCV7uFJFW != 600 || !reflect.DeepEqual(s.Q8W_Oz, [][3]int{{100, 35, 35}, {100, 7, 9}}) || !reflect.DeepEqual(s.PVif2PW, [][3]int{{100, 0, 0}}) {
		t.Fatal("initialization mismatch")
	}
	for i, k := range keys {
		sum := blake3.Sum512(k)
		if !bytes.Equal(s.J2battR3S[i][:], sum[:32]) {
			t.Fatal("fingerprint mismatch")
		}
	}
	if e := s.Init(keys, 0, 0, ""); e == nil || e.Error() != "already initialized" {
		t.Fatal(e)
	}
	var bad GPFxUP0_
	if e := bad.Init([][]byte{{1, 2}}, 0, 0, ""); e == nil || bad.QUnOciyca6eb == nil {
		t.Fatal("partial init lost")
	}
	if e := bad.Init(keys, 0, 0, ""); e == nil || e.Error() != "already initialized" {
		t.Fatal("failed init allowed reuse")
	}
	for _, tc := range []struct{ s, w string }{{"no", "invalid padding lenth/gap parameter: no"}, {"99-35-35", "first padding length must not be smaller than 35"}, {"100-35-65554", "total padding length must not be larger than 65553"}} {
		var sizes, delays [][3]int
		e := TQNZfKT(tc.s, &sizes, &delays)
		if e == nil || e.Error() != tc.w {
			t.Fatal(tc, e)
		}
	}
	var sizes, delays [][3]int
	if e := TQNZfKT("100-35-35.100-0-0.bad", &sizes, &delays); e == nil || len(sizes) != 1 || len(delays) != 1 {
		t.Fatal("parse rollback changed")
	}
	total, l, d := OODk_yfADN([][3]int{{100, 35, 35}, {-1, 7, 7}}, [][3]int{{100, 2, 2}})
	if total != 35 || !reflect.DeepEqual(l, []int{35, 0}) || !reflect.DeepEqual(d, []time.Duration{2 * time.Millisecond}) {
		t.Fatal("padding sampling", total, l, d)
	}
	for i := 0; i < 20; i++ {
		n := IaLb4jaWtI(9, 4)
		if n < 4 || n >= 9 {
			t.Fatal(n)
		}
	}
}

type handshakeFixtureKey struct {
	public []byte
	x      *ecdh.PrivateKey
	kem    *mlkem.DecapsulationKey768
}

func newHandshakeFixtureKey(t *testing.T, kem bool) handshakeFixtureKey {
	t.Helper()
	if kem {
		k, e := mlkem.GenerateKey768()
		if e != nil {
			t.Fatal(e)
		}
		return handshakeFixtureKey{public: k.EncapsulationKey().Bytes(), kem: k}
	}
	k, e := ecdh.X25519().GenerateKey(rand.Reader)
	if e != nil {
		t.Fatal(e)
	}
	return handshakeFixtureKey{public: k.PublicKey().Bytes(), x: k}
}

// This peer parses the wire separately from the production handshake and checks
// both derived public-key fingerprints and the encrypted padding blocks.
func runHandshakeFixture(conn net.Conn, keys []handshakeFixtureKey, mode uint32, nfs int) error {
	defer conn.Close()
	request := make([]byte, nfs+1266+35)
	if _, e := io.ReadFull(conn, request); e != nil {
		return e
	}
	var shared []byte
	var previous cipher.Stream
	offset := 16
	for i, k := range keys {
		size := 32
		if k.kem != nil {
			size = 1088
		}
		encoded := bytes.Clone(request[offset : offset+size])
		offset += size
		if previous != nil {
			previous.XORKeyStream(encoded[:32], encoded[:32])
		}
		if mode != 0 {
			GdvNyLM(k.public, request[:16]).XORKeyStream(encoded, encoded)
		}
		var e error
		if k.kem != nil {
			shared, e = k.kem.Decapsulate(encoded)
		} else {
			var pub *ecdh.PublicKey
			pub, e = ecdh.X25519().NewPublicKey(encoded)
			if e == nil {
				shared, e = k.x.ECDH(pub)
			}
		}
		if e != nil {
			return e
		}
		if i+1 < len(keys) {
			fingerprint := bytes.Clone(request[offset : offset+32])
			offset += 32
			previous = GdvNyLM(shared, request[:16])
			previous.XORKeyStream(fingerprint, fingerprint)
			want := blake3.Sum512(keys[i+1].public)
			if !bytes.Equal(fingerprint, want[:32]) {
				return fmt.Errorf("next fingerprint mismatch")
			}
		}
	}
	if offset != nfs+16 {
		return fmt.Errorf("incorrect initial exchange size")
	}
	handshake := WrnUSnC6(request[:16], shared, true)
	sizePlain, e := handshake.Open(nil, nil, request[offset:offset+18], nil)
	if e != nil {
		return e
	}
	offset += 18
	if binary.BigEndian.Uint16(sizePlain) != 1232 {
		return fmt.Errorf("ephemeral exchange length mismatch")
	}
	publicPair, e := handshake.Open(nil, nil, request[offset:offset+1232], nil)
	if e != nil {
		return e
	}
	offset += 1232
	padLen, e := handshake.Open(nil, nil, request[offset:offset+18], nil)
	if e != nil {
		return e
	}
	offset += 18
	if binary.BigEndian.Uint16(padLen) != 17 {
		return fmt.Errorf("padding length mismatch")
	}
	pad, e := handshake.Open(nil, nil, request[offset:], nil)
	if e != nil || !bytes.Equal(pad, []byte{0}) {
		return fmt.Errorf("padding body mismatch: %v", e)
	}
	clientKEM, e := mlkem.NewEncapsulationKey768(publicPair[:1184])
	if e != nil {
		return e
	}
	kemShared, kemCipher := clientKEM.Encapsulate()
	clientX, e := ecdh.X25519().NewPublicKey(publicPair[1184:])
	if e != nil {
		return e
	}
	serverX, e := ecdh.X25519().GenerateKey(rand.Reader)
	if e != nil {
		return e
	}
	xShared, e := serverX.ECDH(clientX)
	if e != nil {
		return e
	}
	responsePlain := append(kemCipher, serverX.PublicKey().Bytes()...)
	response := handshake.Seal(nil, nonceMaximumAt509A1F0, responsePlain, nil)
	session := append(append(bytes.Clone(kemShared), xShared...), shared...)
	send := WrnUSnC6(responsePlain, session, true)
	token := make([]byte, 16)
	binary.BigEndian.PutUint16(token, 600)
	copy(token[2:], "session token!")
	response = append(response, send.Seal(nil, nil, token, nil)...)
	response = append(response, send.Seal(nil, nil, []byte{0, 17}, nil)...)
	paddingCipher := send.Seal(nil, nil, []byte{0}, nil)
	header := []byte{23, 3, 3, 0, 21}
	record := send.Seal(bytes.Clone(header), nil, []byte("reply"), header)
	if mode == 2 {
		GdvNyLM(session, token).XORKeyStream(record[:5], record[:5])
	}
	response = append(response, paddingCipher...)
	response = append(response, record...)
	if _, e = conn.Write(response); e != nil {
		return e
	}
	h := make([]byte, 5)
	if _, e = io.ReadFull(conn, h); e != nil {
		return e
	}
	if mode == 2 {
		GdvNyLM(session, request[:16]).XORKeyStream(h, h)
	}
	if !bytes.Equal(h, []byte{23, 3, 3, 0, 21}) {
		return fmt.Errorf("client record header mismatch: %v", h)
	}
	ct := make([]byte, 21)
	if _, e = io.ReadFull(conn, ct); e != nil {
		return e
	}
	decoded, e := WrnUSnC6(publicPair, session, true).Open(nil, nil, ct, h)
	if e != nil || string(decoded) != "query" {
		return fmt.Errorf("client payload mismatch %v", e)
	}
	return nil
}

func TestClientHandshakeAgainstIndependentPeer(t *testing.T) {
	for _, mode := range []uint32{0, 1, 2} {
		for _, mixed := range []bool{false, true} {
			t.Run(fmt.Sprintf("mode%d-mixed%v", mode, mixed), func(t *testing.T) {
				keys := []handshakeFixtureKey{newHandshakeFixtureKey(t, false)}
				if mixed {
					keys = append(keys, newHandshakeFixtureKey(t, true), newHandshakeFixtureKey(t, false))
				}
				pubs := make([][]byte, len(keys))
				for i, k := range keys {
					pubs[i] = k.public
				}
				var state GPFxUP0_
				if e := state.Init(pubs, mode, 600, "100-35-35"); e != nil {
					t.Fatal(e)
				}
				client, server := net.Pipe()
				defer client.Close()
				client.SetDeadline(time.Now().Add(10 * time.Second))
				server.SetDeadline(time.Now().Add(10 * time.Second))
				result := make(chan error, 1)
				go func() { result <- runHandshakeFixture(server, keys, mode, state.JhDzI4ehWF) }()
				conn, e := state.Handshake(client)
				if e != nil {
					t.Fatal(e)
				}
				buf := make([]byte, 5)
				if _, e = io.ReadFull(conn, buf); e != nil || string(buf) != "reply" {
					t.Fatalf("response %q %v", buf, e)
				}
				if n, e := conn.Write([]byte("query")); n != 5 || e != nil {
					t.Fatal(n, e)
				}
				if e := <-result; e != nil {
					t.Fatal(e)
				}
				if len(state.Wliqugg3J89) != 64 || cap(state.Wliqugg3J89) != 64 || len(state.HMnsiUUhd) != 16 || !state.EQa3Gaz.After(time.Now()) {
					t.Fatal("session not saved")
				}
			})
		}
	}
}

func TestClientCachedHandshake(t *testing.T) {
	for _, mode := range []uint32{0, 2} {
		key := newHandshakeFixtureKey(t, false)
		var state GPFxUP0_
		if e := state.Init([][]byte{key.public}, mode, 600, "100-35-35"); e != nil {
			t.Fatal(e)
		}
		state.Wliqugg3J89 = bytes.Repeat([]byte{7}, 64)
		state.HMnsiUUhd = bytes.Repeat([]byte{9}, 16)
		state.EQa3Gaz = time.Now().Add(time.Minute)
		wire := &fixtureConn{writeN: -1}
		conn, e := state.Handshake(wire)
		if e != nil {
			t.Fatal(e)
		}
		if wire.reads != 0 || wire.writes != 0 || len(conn.A9M03eHa) != state.JhDzI4ehWF+66 || conn.R5MQKRg26Z != &state || conn.EufjD7UrtI != nil {
			t.Fatal("cached path performed premature I/O")
		}
		prefix := bytes.Clone(conn.A9M03eHa)
		pub := bytes.Clone(prefix[16:48])
		if mode != 0 {
			GdvNyLM(key.public, prefix[:16]).XORKeyStream(pub, pub)
		}
		peerPub, e := ecdh.X25519().NewPublicKey(pub)
		if e != nil {
			t.Fatal(e)
		}
		shared, e := key.x.ECDH(peerPub)
		if e != nil {
			t.Fatal(e)
		}
		h := WrnUSnC6(prefix[:16], shared, true)
		size, e := h.Open(nil, nil, prefix[48:66], nil)
		if e != nil || !bytes.Equal(size, []byte{0, 32}) {
			t.Fatal("resume length", e)
		}
		token, e := h.Open(nil, nil, prefix[66:98], nil)
		if e != nil || !bytes.Equal(token, state.HMnsiUUhd) {
			t.Fatal("resume token", e)
		}
		if !bytes.Equal(conn.ENEXKXNftX, append(bytes.Clone(state.Wliqugg3J89), shared...)) {
			t.Fatal("resume key mismatch")
		}
		if _, e = conn.Write([]byte("query")); e != nil {
			t.Fatal(e)
		}
		if !bytes.Equal(wire.written[0][:len(prefix)], prefix) {
			t.Fatal("prefix altered by stream cipher")
		}
	}
}
func TestClientHandshakeErrors(t *testing.T) {
	var state GPFxUP0_
	if c, e := state.Handshake(nil); c != nil || e == nil || e.Error() != "uninitialized" {
		t.Fatal(c, e)
	}
	key := newHandshakeFixtureKey(t, false)
	state.Init([][]byte{key.public}, 0, 0, "100-35-35")
	sentinel := fmt.Errorf("write failed")
	if c, e := state.Handshake(&fixtureConn{writeErr: sentinel}); c != nil || e != sentinel {
		t.Fatal(c, e)
	}
	if c, e := state.Handshake(&fixtureConn{writeN: -1}); c != nil || e != io.EOF {
		t.Fatal(c, e)
	}
	var bad GPFxUP0_
	bad.Init([][]byte{make([]byte, 32)}, 0, 0, "100-35-35")
	if c, e := bad.Handshake(&fixtureConn{}); c != nil || e == nil || !strings.Contains(e.Error(), "low order") {
		t.Fatal(c, e)
	}
}
