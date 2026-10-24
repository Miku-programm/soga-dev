package bQxeORT

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
	"lukechampine.com/blake3"
	"net"
	"testing"
	"time"
	"unsafe"
)

type fixtureConn struct {
	input    *bytes.Reader
	written  [][]byte
	writeN   int
	writeErr error
	readErr  error
	maxRead  int
	reads    int
	writes   int
}

func (c *fixtureConn) Read(p []byte) (int, error) {
	c.reads++
	if c.maxRead > 0 && len(p) > c.maxRead {
		p = p[:c.maxRead]
	}
	if c.input == nil {
		return 0, io.EOF
	}
	n, e := c.input.Read(p)
	if c.readErr != nil {
		return n, c.readErr
	}
	return n, e
}
func (c *fixtureConn) Write(p []byte) (int, error) {
	c.writes++
	c.written = append(c.written, bytes.Clone(p))
	if c.writeErr != nil {
		return 1, c.writeErr
	}
	if c.writeN >= 0 {
		return c.writeN, nil
	}
	return len(p), nil
}
func (*fixtureConn) Close() error                     { return nil }
func (*fixtureConn) LocalAddr() net.Addr              { return nil }
func (*fixtureConn) RemoteAddr() net.Addr             { return nil }
func (*fixtureConn) SetDeadline(time.Time) error      { return nil }
func (*fixtureConn) SetReadDeadline(time.Time) error  { return nil }
func (*fixtureConn) SetWriteDeadline(time.Time) error { return nil }

func TestRecordLayouts(t *testing.T) {
	var r RfBfjVN4XE
	var l LUKUT5W0R6
	var s GPFxUP0_
	var a RaS2myUz
	if unsafe.Sizeof(r) != 200 || unsafe.Offsetof(r.y6fLiKc) != 120 || unsafe.Offsetof(r.ponyGFmB4cO) != 160 || unsafe.Sizeof(l) != 112 || unsafe.Sizeof(s) != 232 || unsafe.Sizeof(a) != 32 {
		t.Fatal("RTTI layout mismatch")
	}
	var _ net.Conn = &r
	var _ net.Conn = &l
}
func TestRecordHeaderBounds(t *testing.T) {
	for _, n := range []int{0, 16, 17, 8192, 16640, 16641, 65535} {
		h := []byte{23, 3, 3, 0, 0}
		binary.BigEndian.PutUint16(h[3:], uint16(n))
		got, err := QkDUgv(h)
		if got != n || (err != nil) != (n < 17 || n > 16640) {
			t.Fatalf("length %d: %d %v", n, got, err)
		}
	}
	if n, e := QkDUgv([]byte{22, 3, 3, 0, 17}); n != 0 || e.Error() != "invalid header: [22 3 3 0 17]" {
		t.Fatalf("invalid tag %d %v", n, e)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("short header did not panic")
		}
	}()
	QkDUgv(make([]byte, 4, 5))
}
func TestNonceAndKeyDerivation(t *testing.T) {
	material := []byte("secret material")
	context := []byte("salt")
	var key [32]byte
	blake3.DeriveKey(key[:], string(context), material)
	block, _ := aes.NewCipher(key[:])
	base, _ := cipher.NewGCM(block)
	a := WrnUSnC6(context, material, true)
	nonce := make([]byte, 12)
	nonce[11] = 1
	want := base.Seal(nil, nonce, []byte("payload"), []byte("ad"))
	if got := a.Seal(nil, nil, []byte("payload"), []byte("ad")); !bytes.Equal(got, want) {
		t.Fatal("key or nonce mismatch")
	}
	a.NQGg9x = [12]byte{}
	a.NQGg9x[11] = 255
	a.Seal(nil, nil, nil, nil)
	if a.NQGg9x[10] != 1 || a.NQGg9x[11] != 0 {
		t.Fatal("nonce endian/carry")
	}
	before := a.NQGg9x
	if _, err := a.Open(nil, nil, []byte{1}, nil); err == nil || a.NQGg9x == before {
		t.Fatal("failed open must increment")
	}
	before = a.NQGg9x
	a.Seal(nil, make([]byte, 12), nil, nil)
	if a.NQGg9x != before {
		t.Fatal("explicit nonce incremented")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("nonnull empty nonce accepted")
		}
	}()
	a.Seal(nil, make([]byte, 0), nil, nil)
}
func TestRecordRoundTripChunkingAndBufferedReads(t *testing.T) {
	for _, useAES := range []bool{false, true} {
		senderWire := &fixtureConn{writeN: -1}
		key, salt := []byte("material"), []byte("context")
		send := &RfBfjVN4XE{DrpXao: senderWire, HiUDG1Fl: useAES, ENEXKXNftX: key, AVDV5UPQ: WrnUSnC6(salt, key, useAES)}
		plain := bytes.Repeat([]byte("testdata"), 2500)
		if n, e := send.Write(plain); e != nil || n != len(plain) {
			t.Fatalf("write %d %v", n, e)
		}
		if len(senderWire.written) != 3 || len(senderWire.written[0]) != 8213 {
			t.Fatal("8192-byte framing")
		}
		wire := bytes.Join(senderWire.written, nil)
		recv := &RfBfjVN4XE{DrpXao: &fixtureConn{input: bytes.NewReader(wire), maxRead: 3}, HiUDG1Fl: useAES, ENEXKXNftX: key, EufjD7UrtI: WrnUSnC6(salt, key, useAES)}
		var result []byte
		buf := make([]byte, 37)
		for {
			n, e := recv.Read(buf)
			result = append(result, buf[:n]...)
			if e == io.EOF {
				break
			}
			if e != nil {
				t.Fatal(e)
			}
		}
		if !bytes.Equal(result, plain) {
			t.Fatal("reassembled plaintext mismatch")
		}
	}
}
func TestRecordLazySaltPaddingAndNonceWrap(t *testing.T) {
	for _, useAES := range []bool{false, true} {
		key := []byte("shared")
		salt := []byte("0123456789abcdef")
		sendCipher := WrnUSnC6(salt, key, useAES)
		padding := sendCipher.Seal(nil, nil, []byte("padding"), nil)
		wire := &fixtureConn{writeN: -1}
		send := &RfBfjVN4XE{DrpXao: wire, HiUDG1Fl: useAES, ENEXKXNftX: key, AVDV5UPQ: sendCipher, A9M03eHa: append(bytes.Clone(salt), padding...)}
		send.Write([]byte("first"))
		recv := &RfBfjVN4XE{DrpXao: &fixtureConn{input: bytes.NewReader(bytes.Join(wire.written, nil))}, HiUDG1Fl: useAES, ENEXKXNftX: key, WWjo2UGyM: make([]byte, len(padding))}
		buf := make([]byte, 99)
		if n, e := recv.Read(buf); e != nil || string(buf[:n]) != "first" || recv.WWjo2UGyM != nil {
			t.Fatalf("lazy init %d %v", n, e)
		}
		sendCipher = WrnUSnC6(salt, key, useAES)
		readCipher := WrnUSnC6(salt, key, useAES)
		for i := range sendCipher.NQGg9x {
			sendCipher.NQGg9x[i] = 255
			readCipher.NQGg9x[i] = 255
		}
		wire = &fixtureConn{writeN: -1}
		send = &RfBfjVN4XE{DrpXao: wire, HiUDG1Fl: useAES, ENEXKXNftX: key, AVDV5UPQ: sendCipher}
		send.Write([]byte("wrap"))
		send.Write([]byte("next"))
		recv = &RfBfjVN4XE{DrpXao: &fixtureConn{input: bytes.NewReader(bytes.Join(wire.written, nil))}, HiUDG1Fl: useAES, ENEXKXNftX: key, EufjD7UrtI: readCipher}
		for _, want := range []string{"wrap", "next"} {
			n, e := recv.Read(buf)
			if e != nil || string(buf[:n]) != want {
				t.Fatalf("rotation %q %v", buf[:n], e)
			}
		}
	}
}
func TestRecordErrorSemantics(t *testing.T) {
	key, salt := []byte("key"), []byte("salt")
	wire := &fixtureConn{writeN: 1}
	r := &RfBfjVN4XE{DrpXao: wire, AVDV5UPQ: WrnUSnC6(salt, key, true)}
	if n, e := r.Write([]byte("hello")); n != 5 || e != nil {
		t.Fatal("short nil write behavior changed")
	}
	sentinel := errors.New("wire failure")
	wire.writeErr = sentinel
	r.A9M03eHa = []byte{1, 2}
	if n, e := r.Write([]byte("hello")); n != 0 || e != sentinel || r.A9M03eHa != nil {
		t.Fatal("write error/prefix behavior")
	}
	r = &RfBfjVN4XE{}
	if n, e := r.Read(nil); n != 0 || e != nil {
		t.Fatal("empty read")
	}
	if n, e := r.Write(nil); n != 0 || e != nil {
		t.Fatal("empty write")
	}
	state := &GPFxUP0_{Wliqugg3J89: []byte("key")}
	r = &RfBfjVN4XE{DrpXao: &fixtureConn{input: bytes.NewReader([]byte{22, 3, 3, 0, 17})}, R5MQKRg26Z: state, ENEXKXNftX: []byte("key suffix"), EufjD7UrtI: WrnUSnC6(salt, key, true)}
	if n, e := r.Read(make([]byte, 5)); n != 0 || e.Error() != "new handshake needed" || state.EQa3Gaz.IsZero() {
		t.Fatal("session invalidation")
	}
	wire = &fixtureConn{writeN: -1}
	r = &RfBfjVN4XE{DrpXao: wire, HiUDG1Fl: true, ENEXKXNftX: key, AVDV5UPQ: WrnUSnC6(salt, key, true)}
	r.Write([]byte("tamper"))
	wire.written[0][5] ^= 1
	readCipher := WrnUSnC6(salt, key, true)
	r = &RfBfjVN4XE{DrpXao: &fixtureConn{input: bytes.NewReader(wire.written[0])}, EufjD7UrtI: readCipher, R5MQKRg26Z: state}
	if n, e := r.Read(make([]byte, 99)); n != 0 || e == nil || r.R5MQKRg26Z != nil || readCipher.NQGg9x[11] != 1 {
		t.Fatal("auth error state transition")
	}
}

func TestHeaderStreamOnlyTransformsHeadersAcrossFragments(t *testing.T) {
	key := []byte("material")
	iv := make([]byte, 16)
	plain := append([]byte{23, 3, 3, 0, 17}, bytes.Repeat([]byte{0xa7}, 17)...)
	plain = append(plain, 23, 3, 3, 0, 18)
	plain = append(plain, bytes.Repeat([]byte{0xb8}, 18)...)
	headerBytes := append(bytes.Clone(plain[:5]), plain[22:27]...)
	expected := bytes.Clone(plain)
	s := GdvNyLM(key, iv)
	s.XORKeyStream(headerBytes, headerBytes)
	copy(expected[:5], headerBytes[:5])
	copy(expected[22:27], headerBytes[5:])
	wire := &fixtureConn{writeN: -1}
	writer := &LUKUT5W0R6{DrpXao: wire, D1_SMs0WWKz: GdvNyLM(key, iv)}
	for start := 0; start < len(plain); {
		end := min(start+3, len(plain))
		p := bytes.Clone(plain[start:end])
		n, e := writer.Write(p)
		if n != len(p) || e != nil {
			t.Fatal(n, e)
		}
		start = end
	}
	encoded := bytes.Join(wire.written, nil)
	if !bytes.Equal(encoded, expected) {
		t.Fatalf("headers/record bodies differ\n%x\n%x", encoded, expected)
	}
	reader := &LUKUT5W0R6{DrpXao: &fixtureConn{input: bytes.NewReader(encoded), maxRead: 2}, EO5NXI6ex6: GdvNyLM(key, iv)}
	decoded, e := io.ReadAll(reader)
	if e != nil || !bytes.Equal(decoded, plain) {
		t.Fatal("header decoding failed", e)
	}
}
func TestHeaderStreamReadAndWriteErrorBehavior(t *testing.T) {
	iv := make([]byte, 16)
	key := []byte("key")
	wire := &fixtureConn{writeN: 0}
	c := &LUKUT5W0R6{DrpXao: wire, D1_SMs0WWKz: GdvNyLM(key, iv)}
	p := []byte{23, 3, 3, 0, 17}
	if n, e := c.Write(p); n != 5 || e != nil {
		t.Fatal("short write behavior")
	}
	sentinel := errors.New("fault")
	wire.writeErr = sentinel
	c.IdEXNQK5b1k = 20
	if n, e := c.Write([]byte{1, 2}); n != 0 || e != sentinel || c.IdEXNQK5b1k != 18 {
		t.Fatal("write state on error")
	}
	reader := &fixtureConn{input: bytes.NewReader(p), readErr: sentinel}
	c = &LUKUT5W0R6{DrpXao: reader, EO5NXI6ex6: GdvNyLM(key, iv)}
	out := make([]byte, 5)
	if n, e := c.Read(out); n != 5 || e != sentinel || !bytes.Equal(out, []byte{23, 3, 3, 0, 17}) || c.LMy55MYg_l != 17 {
		t.Fatal("bytes alongside error weren't processed")
	}
}
