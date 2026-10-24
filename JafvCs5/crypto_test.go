package JafvCs5

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"testing"
	"unsafe"
)

type finalRead struct{ p []byte }

func (r finalRead) Read(p []byte) (int, error) { return copy(p, r.p), io.EOF }
func TestOriginalStreamReaderDataAndError(t *testing.T) {
	if unsafe.Sizeof(FIo_wn2vRGa{}) != 32 {
		t.Fatal("RTTI")
	}
	block, _ := aes.NewCipher([]byte("0123456789abcdef"))
	iv := make([]byte, 16)
	raw := []byte("data returned with EOF")
	encrypted := make([]byte, len(raw))
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(encrypted, raw)
	r := NewStreamReader(cipher.NewCFBDecrypter(block, iv), finalRead{encrypted})
	p := make([]byte, len(raw)+4)
	n, e := r.Read(p)
	if n != len(raw) || e != io.EOF || !bytes.Equal(p[:n], raw) {
		t.Fatal(n, e, p)
	}
}
func TestOriginalAESGCMOpenAppendAndFailures(t *testing.T) {
	key := []byte("0123456789abcdef")
	nonce := make([]byte, 12)
	plain := []byte("header")
	aad := []byte("auth ID")
	block, _ := aes.NewCipher(key)
	a, _ := cipher.NewGCM(block)
	encrypted := a.Seal(nil, nonce, plain, aad)
	dst := make([]byte, 3, 3+len(plain))
	copy(dst, "pre")
	if err := E_IV2PKO(key, nonce, dst, encrypted, aad); err != nil {
		t.Fatal(err)
	}
	if string(dst[:cap(dst)]) != "preheader" {
		t.Fatal(dst[:cap(dst)])
	}
	if err := E_IV2PKO(key, nonce, nil, encrypted, []byte("wrong")); err == nil {
		t.Fatal("bad AAD accepted")
	}
	if err := E_IV2PKO([]byte{1}, nonce, nil, encrypted, aad); err == nil {
		t.Fatal("bad key accepted")
	}
}
