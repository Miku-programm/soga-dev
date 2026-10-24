package aead

import (
	"bytes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"testing"
	"time"
	"unsafe"
)

func testAEAD(t *testing.T) cipher.AEAD {
	t.Helper()
	c, e := Rk8ppUT2aLh(make([]byte, 16))
	if e != nil {
		t.Fatal(e)
	}
	a, e := c.Encrypter(make([]byte, 16))
	if e != nil {
		t.Fatal(e)
	}
	return a
}

// independentFrame builds records directly from AEAD, independent of recovered writers.
func independentFrame(a cipher.AEAD, p []byte, flag uint16) []byte {
	nonce := make([]byte, a.NonceSize())
	h := []byte{0, 0}
	binary.BigEndian.PutUint16(h, uint16(len(p))|flag)
	out := a.Seal(nil, nonce, h, nil)
	nonce[0] = 1
	return a.Seal(out, nonce, p, nil)
}
func TestStreamWireAndBoundaries(t *testing.T) {
	for _, size := range []int{1, 16383, 16384, 40000} {
		p := bytes.Repeat([]byte{0x5a}, size)
		var wire bytes.Buffer
		a := testAEAD(t)
		w := &DqUAnLs6a{Writer: &wire, AEAD: a}
		n, e := w.Write(p)
		if n != size || e != nil {
			t.Fatalf("write %d %v", n, e)
		}
		raw := append([]byte(nil), wire.Bytes()...)
		nonce := make([]byte, a.NonceSize())
		var got []byte
		for len(raw) > 0 {
			h, err := a.Open(nil, nonce, raw[:18], nil)
			if err != nil {
				t.Fatal(err)
			}
			nonce[0]++
			k := int(binary.BigEndian.Uint16(h))
			raw = raw[18:]
			chunk, err := a.Open(nil, nonce, raw[:k+16], nil)
			if err != nil {
				t.Fatal(err)
			}
			nonce[0]++
			got = append(got, chunk...)
			raw = raw[k+16:]
		}
		if !bytes.Equal(got, p) {
			t.Fatal("wire mismatch")
		}
		r := &Gh18ibyrqaX{Reader: &wire, AEAD: testAEAD(t)}
		got, e = io.ReadAll(r)
		if e != nil || !bytes.Equal(got, p) {
			t.Fatalf("read size %d: %v", size, e)
		}
	}
}
func TestStreamZeroAndAuthentication(t *testing.T) {
	a := testAEAD(t)
	var b bytes.Buffer
	w := &DqUAnLs6a{Writer: &b, AEAD: a}
	if n, e := w.Write(nil); n != 0 || e != nil || b.Len() != 18 {
		t.Fatalf("zero write %d %v %d", n, e, b.Len())
	}
	r := &Gh18ibyrqaX{Reader: &b, AEAD: a}
	if _, e := r.Read(make([]byte, 10)); e != errorAt5098540 || r.fOP3Pl9EegxZ[0] != 1 {
		t.Fatal("zero chunk", e)
	}
	p := []byte("abc")
	wire := independentFrame(a, p, 0xc000)
	r = &Gh18ibyrqaX{Reader: bytes.NewReader(wire), AEAD: a}
	if n, e := r.Read(nil); n != 0 || e != nil || !bytes.Equal(r.sZSEFoeE, p) {
		t.Fatal("zero read must consume", n, e)
	}
	got := make([]byte, 3)
	if n, e := r.Read(got); n != 3 || e != nil || !bytes.Equal(got, p) || r.sZSEFoeE != nil {
		t.Fatal(n, e)
	}
	wire[0] ^= 1
	r = &Gh18ibyrqaX{Reader: bytes.NewReader(wire), AEAD: a}
	if _, e := r.Read(got); e == nil || r.fOP3Pl9EegxZ[0] != 1 {
		t.Fatal("auth failure nonce")
	}
}

type shortWriter struct {
	count int
	err   error
}

func (w shortWriter) Write(p []byte) (int, error) { return min(len(p), w.count), w.err }
func TestStreamWriteCountsAndReadFromEOF(t *testing.T) {
	sentinel := errors.New("writer")
	w := &DqUAnLs6a{Writer: shortWriter{1, sentinel}, AEAD: testAEAD(t)}
	if n, e := w.Write(make([]byte, 20000)); n != 16383 || e != sentinel {
		t.Fatal(n, e)
	}
	w = &DqUAnLs6a{Writer: shortWriter{}, AEAD: testAEAD(t)}
	if n, e := w.Write([]byte("abc")); n != 3 || e != nil {
		t.Fatal(n, e)
	}
	var b bytes.Buffer
	w = &DqUAnLs6a{Writer: &b, AEAD: testAEAD(t)}
	if n, e := w.ReadFrom(bytes.NewReader([]byte("abc"))); n != 3 || e != nil || b.Len() != 71 {
		t.Fatal(n, e, b.Len())
	}
	r := &Gh18ibyrqaX{Reader: &b, AEAD: testAEAD(t)}
	var out bytes.Buffer
	if n, e := r.WriteTo(&out); n != 3 || e != errorAt5098540 || out.String() != "abc" {
		t.Fatal(n, e, out.String())
	}
}
func TestAEADPacketWire(t *testing.T) {
	c, _ := Rk8ppUT2aLh(make([]byte, 16))
	p := []byte("independent UDP")
	wire, e := SJwXfg414(make([]byte, 200), p, c)
	if e != nil {
		t.Fatal(e)
	}
	a, _ := c.Decrypter(wire[:16])
	got, e := a.Open(nil, make([]byte, a.NonceSize()), wire[16:], nil)
	if e != nil || !bytes.Equal(got, p) {
		t.Fatal(e)
	}
	got, e = ED88t50zm(make([]byte, 100), wire, c)
	if e != nil || !bytes.Equal(got, p) {
		t.Fatal(e)
	}
	if _, e = ED88t50zm(make([]byte, 1), wire, c); e != io.ErrShortBuffer {
		t.Fatal(e)
	}
	if _, e = ED88t50zm(nil, wire[:31], c); e != errorAt5098530 {
		t.Fatal(e)
	}
	if _, e = SJwXfg414(make([]byte, 0, 64), p, c); e != io.ErrShortBuffer {
		t.Fatal(e)
	}
	wire[len(wire)-1] ^= 1
	if _, e = ED88t50zm(make([]byte, 100), wire, c); e == nil {
		t.Fatal("tamper accepted")
	}
}
func TestAEADRealStream(t *testing.T) {
	for _, ctor := range []func([]byte) (HugVbq, error){Rk8ppUT2aLh, IEFKpuUGsJ, BeqNJFSF} {
		c, e := ctor(make([]byte, 32))
		if e != nil {
			t.Fatal(e)
		}
		left, right := net.Pipe()
		left.SetDeadline(time.Now().Add(3 * time.Second))
		right.SetDeadline(time.Now().Add(3 * time.Second))
		lc := &AZVSCxZmZQ{Conn: left, HugVbq: c}
		rc := &AZVSCxZmZQ{Conn: right, HugVbq: c}
		p := bytes.Repeat([]byte("x"), 20000)
		done := make(chan error, 1)
		go func() {
			defer rc.Close()
			b := make([]byte, len(p))
			_, e := io.ReadFull(rc, b)
			if e == nil && !bytes.Equal(b, p) {
				e = errors.New("server plaintext")
			}
			if e == nil {
				_, e = rc.Write(b)
			}
			done <- e
		}()
		if _, e = lc.Write(p); e != nil {
			t.Fatal(e)
		}
		b := make([]byte, len(p))
		if _, e = io.ReadFull(lc, b); e != nil || !bytes.Equal(b, p) {
			t.Fatal(e)
		}
		lc.Close()
		if e = <-done; e != nil {
			t.Fatal(e)
		}
	}
}
func TestTransportLayouts(t *testing.T) {
	if unsafe.Sizeof(DqUAnLs6a{}) != 64 || unsafe.Sizeof(Gh18ibyrqaX{}) != 96 || unsafe.Sizeof(Zh37aRKMK{}) != 32 || unsafe.Sizeof(AZVSCxZmZQ{}) != 48 {
		t.Fatal("RTTI layout")
	}
}
