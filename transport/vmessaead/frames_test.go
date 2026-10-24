package vmessaead

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
	"testing"
	"unsafe"
)

type writeResult struct {
	n     int
	err   error
	calls int
}

func (w *writeResult) Write(p []byte) (int, error) { w.calls++; return w.n, w.err }
func testAEAD(t *testing.T) cipher.AEAD {
	t.Helper()
	b, e := aes.NewCipher(make([]byte, 16))
	if e != nil {
		t.Fatal(e)
	}
	a, e := cipher.NewGCM(b)
	if e != nil {
		t.Fatal(e)
	}
	return a
}
func TestFrameRTTILayout(t *testing.T) {
	if unsafe.Sizeof(tojW0z34{}) != 16 || unsafe.Sizeof(vV27Q3k{}) != 72 || unsafe.Sizeof(sBT2tDxX3LZ{}) != 104 || unsafe.Sizeof(se83rHogAI7{}) != 152 {
		t.Fatal("RTTI sizes")
	}
}
func TestPlainFrameWireAndPartialRead(t *testing.T) {
	var wire bytes.Buffer
	p := bytes.Repeat([]byte{0x57}, 16387)
	w := tojW0z34{&wire}
	if n, e := w.Write(p); n != len(p) || e != nil {
		t.Fatal(n, e)
	}
	b := wire.Bytes()
	if len(b) != 16391 || binary.BigEndian.Uint16(b) != 16384 || !bytes.Equal(b[16386:], []byte{0, 3, 0x57, 0x57, 0x57}) {
		t.Fatal("wire")
	}
	r := vV27Q3k{Reader: &wire, gOoPyboUld: make([]byte, 2)}
	out, e := io.ReadAll(&r)
	if e != nil || !bytes.Equal(out, p) {
		t.Fatal(len(out), e)
	}
	// An empty caller buffer still consumes a header and retains the frame.
	r = vV27Q3k{Reader: bytes.NewReader([]byte{0, 2, 7, 8}), gOoPyboUld: make([]byte, 2)}
	if n, e := r.Read(nil); n != 0 || e != nil || len(r.dphuptdYtW) != 2 {
		t.Fatal(n, e)
	}
	one := make([]byte, 1)
	r.Read(one)
	if one[0] != 7 {
		t.Fatal(one)
	}
	r.Read(one)
	if one[0] != 8 || r.dphuptdYtW != nil {
		t.Fatal(one)
	}
}
func TestPlainErrorAndShortWriteSemantics(t *testing.T) {
	r := vV27Q3k{Reader: bytes.NewReader([]byte{0, 3, 1}), gOoPyboUld: make([]byte, 2)}
	if n, e := r.Read(make([]byte, 8)); n != 0 || e != io.ErrUnexpectedEOF {
		t.Fatal(n, e)
	}
	r = vV27Q3k{Reader: bytes.NewReader([]byte{0x44, 1}), gOoPyboUld: make([]byte, 2)}
	if n, e := r.Read(nil); n != 0 || e == nil || e.Error() != "buffer is larger than standard" {
		t.Fatal(n, e)
	}
	sink := &writeResult{}
	w := tojW0z34{sink}
	if n, e := w.Write([]byte{1, 2}); n != 2 || e != nil {
		t.Fatal(n, e)
	}
	sink.err = io.ErrClosedPipe
	if n, e := w.Write([]byte{3}); n != 0 || e != io.ErrClosedPipe {
		t.Fatal(n, e)
	}
}
func TestAEADFrameKnownWireNonceAndAuthentication(t *testing.T) {
	var wire bytes.Buffer
	a := testAEAD(t)
	iv := []byte{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 9, 8}
	w := sBT2tDxX3LZ{Writer: &wire, AEAD: a, xNF1Gb47: iv, hEJ1_Tabe1: 65535}
	p := []byte("frame")
	w.Write(p)
	w.Write(p)
	nonce := append([]byte{255, 255}, iv[2:]...)
	want := a.Seal(nil, nonce, p, nil)
	if binary.BigEndian.Uint16(wire.Bytes()) != 21 || !bytes.Equal(wire.Bytes()[2:23], want) || w.hEJ1_Tabe1 != 1 {
		t.Fatal("wire/counter")
	}
	r := se83rHogAI7{Reader: &wire, AEAD: a, frnZkWRc: iv, jYxJbABejRtF: make([]byte, 2), zIlPJ70bRN: 65535}
	out, e := io.ReadAll(&r)
	if e != nil || string(out) != "frameframe" || r.zIlPJ70bRN != 1 {
		t.Fatal(string(out), e)
	}
	bad := append([]byte{0, 21}, want...)
	bad[len(bad)-1] ^= 1
	r = se83rHogAI7{Reader: bytes.NewReader(bad), AEAD: a, frnZkWRc: iv, jYxJbABejRtF: make([]byte, 2), zIlPJ70bRN: 65535}
	if n, e := r.Read(make([]byte, 10)); n != 0 || e == nil || r.zIlPJ70bRN != 0 || r.aiE3W5V != nil {
		t.Fatal(n, e)
	}
}
func TestAEADWriterFailureAdvancesNonce(t *testing.T) {
	err := errors.New("fail")
	sink := &writeResult{err: err}
	w := sBT2tDxX3LZ{Writer: sink, AEAD: testAEAD(t), xNF1Gb47: make([]byte, 12)}
	if n, e := w.Write([]byte{1}); n != 0 || e != err || w.hEJ1_Tabe1 != 1 {
		t.Fatal(n, e)
	}
	sink.err = nil
	if n, e := w.Write([]byte{2}); n != 1 || e != nil || w.hEJ1_Tabe1 != 2 {
		t.Fatal(n, e)
	}
}
