package wserr

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rc4"
	"errors"
	"golang.org/x/crypto/chacha20"
	"io"
	"net"
	"testing"
	"time"
	"unsafe"
)

func allCiphers(t *testing.T) []ZDHlm7vHo {
	t.Helper()
	var out []ZDHlm7vHo
	for _, ctor := range []func([]byte) (ZDHlm7vHo, error){E6vQRRYI, HQZtzHEUA3, E9kDmJ55P, AfbgTLVHrju, F3Yl5rl} {
		c, e := ctor(make([]byte, 32))
		if e != nil {
			t.Fatal(e)
		}
		out = append(out, c)
	}
	return out
}
func TestCipherIndependentWire(t *testing.T) {
	for i, c := range allCiphers(t) {
		iv := bytes.Repeat([]byte{3}, c.IVSize())
		p := bytes.Repeat([]byte("abc"), 1000)
		got := make([]byte, len(p))
		c.Encrypter(iv).XORKeyStream(got, p)
		var independent cipher.Stream
		block, _ := aes.NewCipher(make([]byte, 32))
		switch i {
		case 0:
			independent = cipher.NewCTR(block, iv)
		case 1:
			independent = cipher.NewCFBEncrypter(block, iv)
		case 2, 3:
			independent, _ = chacha20.NewUnauthenticatedCipher(make([]byte, 32), iv)
		case 4:
			h := md5.Sum(append(make([]byte, 32), iv...))
			independent, _ = rc4.NewCipher(h[:])
		}
		want := make([]byte, len(p))
		independent.XORKeyStream(want, p)
		if !bytes.Equal(got, want) {
			t.Fatalf("cipher %d", i)
		}
		c.Decrypter(iv).XORKeyStream(got, got)
		if !bytes.Equal(got, p) {
			t.Fatal("decrypt")
		}
		wire, e := LTkaLW3KDj(make([]byte, 4000), p, c)
		if e != nil {
			t.Fatal(e)
		}
		dec := make([]byte, len(p))
		c.Decrypter(wire[:c.IVSize()]).XORKeyStream(dec, wire[c.IVSize():])
		if !bytes.Equal(dec, p) {
			t.Fatal("packet wire")
		}
		dec, e = SP6JM3LOTFRv(dec, wire, c)
		if e != nil || !bytes.Equal(dec, p) {
			t.Fatal(e)
		}
		if _, e = SP6JM3LOTFRv(nil, wire[:c.IVSize()-1], c); e != errorAt5098550 {
			t.Fatal(e)
		}
		if _, e = LTkaLW3KDj(nil, p, c); e != io.ErrShortBuffer {
			t.Fatal(e)
		}
	}
}

type dataAndEOF struct{ p []byte }

func (r *dataAndEOF) Read(p []byte) (int, error) { n := copy(p, r.p); r.p = r.p[n:]; return n, io.EOF }

type shortOutput struct {
	writes int
	err    error
}

func (w *shortOutput) Write(p []byte) (int, error) { w.writes++; return min(1, len(p)), w.err }
func TestReadAndTransferErrorSemantics(t *testing.T) {
	c, _ := E6vQRRYI(make([]byte, 16))
	iv := make([]byte, 16)
	raw := []byte{1, 2, 3}
	r := &Z0mJsCE{Reader: &dataAndEOF{append([]byte(nil), raw...)}, Stream: c.Decrypter(iv)}
	p := make([]byte, 3)
	if n, e := r.Read(p); n != 0 || e != io.EOF || !bytes.Equal(p, raw) {
		t.Fatal(n, e, p)
	}
	r = &Z0mJsCE{Reader: &dataAndEOF{raw}, Stream: c.Decrypter(iv)}
	out := &shortOutput{}
	if n, e := r.WriteTo(out); n != 1 || e != nil {
		t.Fatal(n, e)
	}
	out = &shortOutput{}
	w := &SBrL2pUbjeTr{Writer: out, Stream: c.Encrypter(iv)}
	if n, e := w.ReadFrom(bytes.NewReader(raw)); n != 3 || e != nil || out.writes != 2 {
		t.Fatal(n, e, out.writes)
	}
	err := errors.New("write")
	w = &SBrL2pUbjeTr{Writer: &shortOutput{err: err}, Stream: c.Encrypter(iv)}
	if n, e := w.Write(make([]byte, 3000)); n != 1 || e != err {
		t.Fatal(n, e)
	}
}
func TestRealConnectionsAndIVCache(t *testing.T) {
	for _, c := range allCiphers(t) {
		left, right := net.Pipe()
		left.SetDeadline(time.Now().Add(3 * time.Second))
		right.SetDeadline(time.Now().Add(3 * time.Second))
		lc := &AriwPpia{Conn: left, ZDHlm7vHo: c}
		rc := &AriwPpia{Conn: right, ZDHlm7vHo: c}
		iv, e := lc.ObtainWriteIV()
		if e != nil {
			t.Fatal(e)
		}
		again, e := lc.ObtainWriteIV()
		if e != nil || &again[0] != &iv[0] {
			t.Fatal("IV not cached")
		}
		p := bytes.Repeat([]byte("abc"), 3000)
		done := make(chan error, 1)
		go func() {
			defer rc.Close()
			b := make([]byte, len(p))
			_, e := io.ReadFull(rc, b)
			if e == nil && !bytes.Equal(b, p) {
				e = errors.New("request")
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
func TestConstructorsAndLayouts(t *testing.T) {
	if _, e := E6vQRRYI([]byte{1}); e != aes.KeySizeError(1) {
		t.Fatal(e)
	}
	for _, ctor := range []func([]byte) (ZDHlm7vHo, error){E9kDmJ55P, AfbgTLVHrju} {
		if _, e := ctor(nil); e != XJzap3(32) {
			t.Fatal(e)
		}
	}
	if _, e := F3Yl5rl(nil); e != nil {
		t.Fatal(e)
	}
	if XJzap3(32).Error() != "key size error: need 32 bytes" {
		t.Fatal("literal")
	}
	if unsafe.Sizeof(SBrL2pUbjeTr{}) != 2080 || unsafe.Sizeof(Z0mJsCE{}) != 2080 || unsafe.Sizeof(AriwPpia{}) != 96 || unsafe.Sizeof(HtnJMv4Lsng{}) != 32 {
		t.Fatal("RTTI")
	}
}
