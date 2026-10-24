package low0s_10x

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rc4"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"io"
	"net"
	H "recovered.local/soga/H8_In_Vuy"
	K "recovered.local/soga/m3Cw8P3_Le"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

type wireConn struct {
	input  *bytes.Reader
	output bytes.Buffer
	err    error
	calls  int
}

func (c *wireConn) Read(p []byte) (int, error) {
	c.calls++
	if c.err != nil {
		return 0, c.err
	}
	if c.input == nil {
		return 0, io.EOF
	}
	return c.input.Read(p)
}
func (c *wireConn) Write(p []byte) (int, error) {
	if c.err != nil {
		return 0, c.err
	}
	return c.output.Write(p)
}
func (*wireConn) Close() error                     { return nil }
func (*wireConn) LocalAddr() net.Addr              { return nil }
func (*wireConn) RemoteAddr() net.Addr             { return nil }
func (*wireConn) SetDeadline(time.Time) error      { return nil }
func (*wireConn) SetReadDeadline(time.Time) error  { return nil }
func (*wireConn) SetWriteDeadline(time.Time) error { return nil }

func TestLayoutsAndRegistry(t *testing.T) {
	for _, x := range []struct {
		v any
		n uintptr
	}{{PZfVxx4UGRb{}, 48}, {bhRsfpsWw{}, 32}, {t1vL_y55t{}, 16}, {oNvVPHch3WB{}, 32}, {sdA1Al25W{}, 72}, {vcMVzsjn5n{}, 200}, {cQb8tRiyK{}, 56}, {mh5mOR8HqIcG{}, 48}, {A6KE0r{}, 112}, {S_N0Ixh{}, 32}} {
		if reflect.TypeOf(x.v).Size() != x.n {
			t.Fatalf("%T size %d", x.v, reflect.TypeOf(x.v).Size())
		}
	}
	if unsafe.Offsetof(sdA1Al25W{}.kGLv0DUCiIh) != 60 || unsafe.Offsetof(vcMVzsjn5n{}.j4NIsaaxx) != 192 {
		t.Fatal("field offsets")
	}
	for n, w := range map[string]int{"origin": 0, "auth_aes128_md5": 9, "auth_aes128_sha1": 9, "auth_sha1_v4": 7, "auth_chain_a": 4, "auth_chain_b": 4} {
		b := &PZfVxx4UGRb{Z2UHwhi2h4: bytes.Repeat([]byte{1}, 16), UqLugVVWCKJ: 5}
		p, e := HBLvlfQr0d(n, b)
		if e != nil || p == nil || b.UqLugVVWCKJ != 5+w {
			t.Fatal(n, e, b)
		}
	}
	if p, e := HBLvlfQr0d("ORIGIN", nil); p != nil || e == nil || e.Error() != "protocol ORIGIN not supported" {
		t.Fatal(p, e)
	}
}
func newAES(t *testing.T, name string) (*sdA1Al25W, *wireConn) {
	t.Helper()
	p, e := HBLvlfQr0d(name, &PZfVxx4UGRb{Z2UHwhi2h4: []byte("0123456789abcdef"), BOJR3xVshAb: "42:secret:ignored"})
	if e != nil {
		t.Fatal(e)
	}
	raw := &wireConn{}
	c := p.StreamConn(raw, []byte("fedcba9876543210")).(*A6KE0r)
	return c.Shdj3DH1.(*sdA1Al25W), raw
}
func mac(f func() hash.Hash, key, b []byte) []byte {
	h := hmac.New(f, key)
	h.Write(b)
	return h.Sum(nil)
}
func TestAESIndependentHandshakeAndRecord(t *testing.T) {
	for _, name := range []string{"auth_aes128_md5", "auth_aes128_sha1"} {
		t.Run(name, func(t *testing.T) {
			a, _ := newAES(t, name)
			hf := md5.New
			if name == "auth_aes128_sha1" {
				hf = sha1.New
			}
			h := hf()
			h.Write([]byte("secret"))
			uk := h.Sum(nil)
			if !bytes.Equal(uk, a.oYNnYN0vN) || binary.LittleEndian.Uint32(a.qC81k5Mk[:]) != 42 {
				t.Fatal("user parameter")
			}
			var buf bytes.Buffer
			data := []byte("hello")
			a.Encode(&buf, data)
			w := buf.Bytes()
			key := append(append([]byte(nil), a.fq3rseqxscn...), a.Z2UHwhi2h4...)
			if len(w) < 40 || !hmac.Equal(w[1:7], mac(hf, key, w[:1])[:6]) || !hmac.Equal(w[27:31], mac(hf, key, w[7:27])[:4]) || !hmac.Equal(w[len(w)-4:], mac(hf, uk, w[:len(w)-4])[:4]) {
				t.Fatal("authentication MAC")
			}
			block, _ := aes.NewCipher(K.ArNrGS_EV(base64.StdEncoding.EncodeToString(uk)+name, 16))
			plain := make([]byte, 16)
			cipher.NewCBCDecrypter(block, make([]byte, 16)).CryptBlocks(plain, w[11:27])
			if int(binary.LittleEndian.Uint16(plain[12:])) != len(w) || int(binary.LittleEndian.Uint16(plain[14:])) != len(w)-35-len(data) || binary.LittleEndian.Uint32(plain[8:]) != a.xOB1UU || !bytes.Equal(plain[4:8], a.kh9UY106p3Y6[:]) {
				t.Fatalf("CBC fields %x", plain)
			}
			if !bytes.Equal(w[len(w)-4-len(data):len(w)-4], data) {
				t.Fatal("auth payload")
			}
			buf.Reset()
			large := bytes.Repeat([]byte{0x5a}, 17000)
			if e := a.Encode(&buf, large); e != nil {
				t.Fatal(e)
			}
			wire := append([]byte(nil), buf.Bytes()...)
			got := []byte{}
			id := uint32(1)
			for len(wire) > 0 {
				n := int(binary.LittleEndian.Uint16(wire))
				if n < 9 || n >= 8192 || n > len(wire) {
					t.Fatal(n)
				}
				key := append(append([]byte(nil), uk...), make([]byte, 4)...)
				binary.LittleEndian.PutUint32(key[len(uk):], id)
				if !hmac.Equal(wire[2:4], mac(hf, key, wire[:2])[:2]) || !hmac.Equal(wire[n-4:n], mac(hf, key, wire[:n-4])[:4]) {
					t.Fatal("record MAC")
				}
				pos := int(wire[4])
				if pos == 255 {
					pos = int(binary.LittleEndian.Uint16(wire[5:7]))
				}
				got = append(got, wire[4+pos:n-4]...)
				wire = wire[n:]
				id++
			}
			if !bytes.Equal(got, large) || id != 4 {
				t.Fatal("chunked payload", len(got), id)
			}
			var decoded bytes.Buffer
			full := append([]byte(nil), buf.Bytes()...)
			src := bytes.NewBuffer(full[:3])
			if e := a.Decode(&decoded, src); e != nil || src.Len() != 3 {
				t.Fatal(e)
			}
			src.Write(full[3:])
			if e := a.Decode(&decoded, src); e != nil || !bytes.Equal(decoded.Bytes(), large) || src.Len() != 0 {
				t.Fatal(e, "decode")
			}
		})
	}
}
func TestSHA1V4IndependentWireAndErrors(t *testing.T) {
	p, _ := HBLvlfQr0d("auth_sha1_v4", &PZfVxx4UGRb{Z2UHwhi2h4: []byte("0123456789abcdef")})
	a := p.StreamConn(&wireConn{}, []byte("0123456789abcdef")).(*A6KE0r).Shdj3DH1.(*mh5mOR8HqIcG)
	var out bytes.Buffer
	a.Encode(&out, []byte("payload"))
	w := out.Bytes()
	n := int(binary.BigEndian.Uint16(w))
	crcinput := append(append(append([]byte{}, w[:2]...), []byte("auth_sha1_v4")...), a.Z2UHwhi2h4...)
	key := append(append([]byte{}, a.j4sXbBI9LwU...), a.Z2UHwhi2h4...)
	if n != len(w) || binary.LittleEndian.Uint32(w[2:6]) != crc32.ChecksumIEEE(crcinput) || !hmac.Equal(w[n-10:], mac(sha1.New, key, w[:n-10])[:10]) {
		t.Fatal("auth wire")
	}
	out.Reset()
	data := bytes.Repeat([]byte("packet"), 2500)
	a.Encode(&out, data)
	src := bytes.NewBuffer(append([]byte{}, out.Bytes()...))
	var decoded bytes.Buffer
	if e := a.Decode(&decoded, src); e != nil || !bytes.Equal(decoded.Bytes(), data) {
		t.Fatal(e)
	}
	// Construct an independent server record with one padding-length byte.
	server := []byte{0, 12, 0, 0, 1, 'a', 'b', 'c', 0, 0, 0, 0}
	binary.LittleEndian.PutUint16(server[2:], uint16(crc32.ChecksumIEEE(server[:2])))
	binary.LittleEndian.PutUint32(server[8:], adler32.Checksum(server[:8]))
	decoded.Reset()
	if e := a.Decode(&decoded, bytes.NewBuffer(server)); e != nil || decoded.String() != "abc" {
		t.Fatal(e, decoded.String())
	}
	bad := append([]byte(nil), server...)
	bad[2] ^= 1
	src = bytes.NewBuffer(bad)
	if e := a.Decode(&decoded, src); e != errorAt5098590 || a.enGr6iXKNAp || src.Len() != 0 {
		t.Fatal(e)
	}
	bad = append([]byte(nil), server...)
	bad[len(bad)-1] ^= 1
	src = bytes.NewBuffer(bad)
	if e := a.Decode(&decoded, src); e != errorAt50985B0 || !a.enGr6iXKNAp || src.Len() != 0 {
		t.Fatal(e)
	}
	decoded.Reset()
	if e := a.Decode(&decoded, bytes.NewBufferString("raw")); e != nil || decoded.String() != "raw" {
		t.Fatal(e)
	}
}
func TestAESUDPDirectionAndFailures(t *testing.T) {
	for _, name := range []string{"auth_aes128_md5", "auth_aes128_sha1"} {
		a, _ := newAES(t, name)
		var b bytes.Buffer
		a.EncodePacket(&b, []byte("udp"))
		w := b.Bytes()
		if len(w) != 11 || string(w[:3]) != "udp" || binary.LittleEndian.Uint32(w[3:7]) != 42 || !bytes.Equal(w[7:], a.efwls6Jy(a.oYNnYN0vN, w[:7])[:4]) {
			t.Fatal(name, w)
		}
		inbound := append([]byte("server"), a.efwls6Jy(a.Z2UHwhi2h4, []byte("server"))[:4]...)
		got, e := a.DecodePacket(inbound)
		if e != nil || string(got) != "server" {
			t.Fatal(name, e)
		}
		inbound[len(inbound)-1] ^= 1
		if _, e = a.DecodePacket(inbound); e != errorAt50985E0 {
			t.Fatal(e)
		}
		if _, e = a.DecodePacket(nil); e != errorAt50985D0 {
			t.Fatal(e)
		}
	}
}
func chainProtocol(t *testing.T, name string) (Shdj3DH1, *vcMVzsjn5n) {
	t.Helper()
	p, e := HBLvlfQr0d(name, &PZfVxx4UGRb{Z2UHwhi2h4: []byte("0123456789abcdef"), BOJR3xVshAb: "42:chain secret"})
	if e != nil {
		t.Fatal(e)
	}
	c := p.StreamConn(&wireConn{}, []byte("fedcba9876543210")).(*A6KE0r)
	if b, ok := c.Shdj3DH1.(*cQb8tRiyK); ok {
		return b, b.vcMVzsjn5n
	}
	a := c.Shdj3DH1.(*vcMVzsjn5n)
	return a, a
}
func TestAuthChainHandshakeAndInboundRecord(t *testing.T) {
	for _, name := range []string{"auth_chain_a", "auth_chain_b"} {
		t.Run(name, func(t *testing.T) {
			p, a := chainProtocol(t, name)
			var out bytes.Buffer
			p.Encode(&out, []byte("first payload"))
			w := out.Bytes()
			mk := append(append([]byte(nil), a.fTzASkj...), a.Z2UHwhi2h4...)
			clientHash := mac(md5.New, mk, w[:4])
			if !hmac.Equal(w[4:12], clientHash[:8]) || binary.LittleEndian.Uint32(w[12:16])^binary.LittleEndian.Uint32(clientHash[8:12]) != 42 {
				t.Fatal("initial MAC/UID")
			}
			serverHash := mac(md5.New, a.oYNnYN0vN, w[12:32])
			if !hmac.Equal(w[32:36], serverHash[:4]) {
				t.Fatal("server MAC")
			}
			block, _ := aes.NewCipher(K.ArNrGS_EV(base64.StdEncoding.EncodeToString(a.oYNnYN0vN)+name, 16))
			plain := make([]byte, 16)
			cipher.NewCBCDecrypter(block, make([]byte, 16)).CryptBlocks(plain, w[16:32])
			if binary.LittleEndian.Uint16(plain[12:14]) != 4 || binary.LittleEndian.Uint16(plain[14:16]) != 0 {
				t.Fatal("auth overhead", plain)
			}
			// Independent server RC4 stream. The first payload includes two MSS bytes.
			rc, _ := rc4.NewCipher(K.ArNrGS_EV(base64.StdEncoding.EncodeToString(a.oYNnYN0vN)+base64.StdEncoding.EncodeToString(clientHash), 16))
			payload := append([]byte{0xb4, 5}, []byte("reply payload")...)
			rc.XORKeyStream(payload, payload)
			var rng H.YE9p7Ja6vRfW
			pad := a.eXaIwIUeDfSi(len(payload), serverHash, &rng)
			start := 0
			if pad > 0 {
				start = int(int64(rng.Next()%8589934609) % int64(pad))
			}
			record := make([]byte, 2+pad+len(payload)+2)
			binary.LittleEndian.PutUint16(record, uint16(len(payload))^binary.LittleEndian.Uint16(serverHash[14:]))
			copy(record[2+start:], payload)
			mkey := append(append([]byte{}, a.oYNnYN0vN...), 1, 0, 0, 0)
			copy(record[len(record)-2:], mac(md5.New, mkey, record[:len(record)-2])[:2])
			var decoded bytes.Buffer
			if e := p.Decode(&decoded, bytes.NewBuffer(record)); e != nil || decoded.String() != "reply payload" || a.nKJS2vXY6 != 2 {
				t.Fatal(e, decoded.String(), a.nKJS2vXY6)
			}
		})
	}
}
func TestAuthChainUDPMutationAndInbound(t *testing.T) {
	for _, name := range []string{"auth_chain_a", "auth_chain_b"} {
		p, a := chainProtocol(t, name)
		data := []byte("client UDP payload")
		original := append([]byte(nil), data...)
		var out bytes.Buffer
		if e := p.EncodePacket(&out, data); e != nil {
			t.Fatal(e)
		}
		w := out.Bytes()
		if bytes.Equal(data, original) || !bytes.Equal(w[:len(data)], data) || !hmac.Equal(w[len(w)-1:], mac(md5.New, a.oYNnYN0vN, w[:len(w)-1])[:1]) {
			t.Fatal("encode mutation/MAC")
		}
		auth := w[len(w)-8 : len(w)-5]
		digest := mac(md5.New, a.Z2UHwhi2h4, auth)
		if binary.LittleEndian.Uint32(w[len(w)-5:])^binary.LittleEndian.Uint32(digest) != 42 {
			t.Fatal("uid")
		}
		rc, _ := rc4.NewCipher(K.ArNrGS_EV(base64.StdEncoding.EncodeToString(a.oYNnYN0vN)+base64.StdEncoding.EncodeToString(digest), 16))
		got := append([]byte{}, data...)
		rc.XORKeyStream(got, got)
		if !bytes.Equal(got, original) {
			t.Fatal("client cipher")
		}
		serverAuth := []byte{1, 2, 3, 4, 5, 6, 7}
		digest = mac(md5.New, a.Z2UHwhi2h4, serverAuth)
		var rng H.YE9p7Ja6vRfW
		rng.InitFromBin(digest)
		padding := int(rng.Next() % 127)
		in := []byte("server UDP payload")
		rc, _ = rc4.NewCipher(K.ArNrGS_EV(base64.StdEncoding.EncodeToString(a.oYNnYN0vN)+base64.StdEncoding.EncodeToString(digest), 16))
		rc.XORKeyStream(in, in)
		in = append(in, make([]byte, padding)...)
		in = append(in, serverAuth...)
		in = append(in, mac(md5.New, a.oYNnYN0vN, in)[0])
		plain, e := p.DecodePacket(in)
		if e != nil || string(plain) != "server UDP payload" {
			t.Fatal(name, e, string(plain))
		}
	}
}
func TestAuthCounterAndTransportErrors(t *testing.T) {
	a := &t1vL_y55t{kh9UY106p3Y6: [4]byte{1, 2, 3, 4}, xOB1UU: 0xff000000}
	n := a.uU1wpTCSA()
	if n.xOB1UU != 0xff000001 || n.kh9UY106p3Y6 != a.kh9UY106p3Y6 {
		t.Fatal(n)
	}
	n = a.uU1wpTCSA()
	if n.xOB1UU > 0x1000000 || n.xOB1UU == 0 {
		t.Fatal(n)
	}
	sentinel := errors.New("wire failure")
	r := &wireConn{err: sentinel}
	p := &fy1trpRM{}
	c := &A6KE0r{Conn: r, Shdj3DH1: p}
	if n, e := c.Write([]byte("x")); n != 0 || e != sentinel {
		t.Fatal(n, e)
	}
	if n, e := c.Read(make([]byte, 8)); n != 0 || e != sentinel {
		t.Fatal(n, e)
	}
	var b bytes.Buffer
	p.EncodePacket(&b, []byte("a"))
	v, e := p.DecodePacket(b.Bytes())
	if e != nil || string(v) != "a" {
		t.Fatal(e)
	}
}
