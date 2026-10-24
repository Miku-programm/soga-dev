package gk1z9v

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"testing"
	"time"

	"golang.org/x/crypto/chacha20poly1305"
	M "recovered.local/soga/G7A7csgrs62p"
	L "recovered.local/soga/LFWTTXF7JR"
	U "recovered.local/soga/nzATvajt"
)

func TestNativeVMessAuthLayoutAndKnownWireValues(t *testing.T) {
	for _, tc := range []struct {
		v any
		n uintptr
	}{{CWAyuvwV9{}, 16}, {E3b13Z3FO{}, 88}, {WR4lMNfab{}, 56}, {M3rebiEbLfQx{}, 48}} {
		if n := reflect.TypeOf(tc.v).Size(); n != tc.n {
			t.Fatalf("%T: %d", tc.v, n)
		}
	}
	user, err := U.EOkO09L(&L.Aga91y{WZ89pTAOrHK4: 0x10000001f, JAfXWwzYrW: &L.Yn8BJoL{SfiEGpJo: "01234567-89ab-cdef-0123-456789abcdef"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	users := &U.ZpNFqUVZp{GTFAWF: new(M.GTFAWF[int64, U.PgaNd3014U])}
	users.Add(user)
	m := NewMD5AuthManager(context.Background(), users, false)
	bucket := m.zKRpQupzoQE([]*U.QKaIWV2{user}, 1700000000, 1700000001)
	raw, _ := hex.DecodeString("cc1eaa49d86aa2901342030b7d3d4a30")
	var key [16]byte
	copy(key[:], raw)
	if v, ok := bucket.iDoMs2[key]; !ok || v != [2]uint32{31, 1700000000} {
		t.Fatal(v, ok)
	}
	m.jxyBGY(bucket)
	id, stamp, ok := m.Get(key)
	if !ok || id != 31 || stamp != 1700000000 {
		t.Fatal(id, stamp, ok)
	}
	auth := &WR4lMNfab{Xwwpafq: user, QAkNeaNePGH: 1700000000}
	if got := hex.EncodeToString(auth.GetIV()); got != "7a29d7bdc972aaefb1f54e5205589e78" {
		t.Fatal(got)
	}
	if !bytes.Equal(auth.GetCmdKey(), user.GetCmdKey()) {
		t.Fatal("command key")
	}
	if m.gHO0Rwg != 70 || NewMD5AuthManager(context.Background(), users, true).gHO0Rwg != 30 {
		t.Fatal("auth windows")
	}
	m.GenAuthForNewUsers([]*U.QKaIWV2{user})
	m.GenAuthForNewUsers(nil)
	if p := m.nlX_TBEa.Swap(nil); p == nil || *p != nil {
		t.Fatal("pending update must overwrite")
	}
}
func TestNativeVMessAuthSearchPriorityAndPrune(t *testing.T) {
	m := &E3b13Z3FO{}
	key := [16]byte{1}
	for _, start := range []uint32{30, 0, 20, 10, 40} {
		m.jxyBGY(&CWAyuvwV9{map[[16]byte][2]uint32{key: {start, start}}, start})
	}
	id, _, ok := m.Get(key)
	if !ok || id != 10 {
		t.Fatal("original center-left priority", id, ok)
	}
	backing := m.hz3qQl7H5
	m.ajZZhr(20)
	if len(m.hz3qQl7H5) != 3 || backing[0] != nil || backing[1] != nil {
		t.Fatal("prune")
	}
	m.hzdGla95()
	if m.hz3qQl7H5 == nil || len(m.hz3qQl7H5) != 0 {
		t.Fatal("non-nil empty reset")
	}
}
func TestNativeVMessTaskCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	users := &U.ZpNFqUVZp{GTFAWF: new(M.GTFAWF[int64, U.PgaNd3014U])}
	m := NewMD5AuthManager(ctx, users, false)
	done := make(chan struct{})
	go func() { m.ofF3NP4P(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("canceled worker did not exit")
	}
}
func TestNativeVMessAEADWireCounterFailureAndRollover(t *testing.T) {
	for _, tc := range []struct {
		name string
		new  func([]byte, []byte) (*M3rebiEbLfQx, error)
	}{{"aes-gcm", eG_qzDpw}, {"chacha", vLtpACzP6C}} {
		t.Run(tc.name, func(t *testing.T) {
			key := []byte("0123456789abcdef")
			iv := []byte("abcdefghijklmnop")
			plain := []byte("original vmess frame")
			enc, err := tc.new(key, iv)
			if err != nil {
				t.Fatal(err)
			}
			dec, err := tc.new(key, iv)
			if err != nil {
				t.Fatal(err)
			}
			var reference cipher.AEAD
			if tc.name == "aes-gcm" {
				block, _ := aes.NewCipher(key)
				reference, _ = cipher.NewGCM(block)
			} else {
				a := md5.Sum(key)
				b := md5.Sum(a[:])
				reference, _ = chacha20poly1305.New(append(a[:], b[:]...))
			}
			nonce := append([]byte(nil), iv[:12]...)
			nonce[0] = 0
			nonce[1] = 0
			dst := make([]byte, len(plain)+16)
			n, err := enc.Encrypt(dst, plain)
			if err != nil || n != len(dst) || !bytes.Equal(dst, reference.Seal(nil, nonce, plain, nil)) {
				t.Fatal("wire", n, err)
			}
			got := make([]byte, len(plain))
			n, err = dec.Decrypt(got, dst)
			if err != nil || n != len(plain) || !bytes.Equal(got, plain) {
				t.Fatal(n, err)
			}
			if n, err := dec.Decrypt(got, dst); n != 0 || err == nil || dec.pufJF2 != 2 {
				t.Fatal("failure must consume nonce", n, err)
			}
			enc.pufJF2 = 65535
			enc.Encrypt(dst, plain)
			if enc.pufJF2 != 0 || enc.fKs57NhW6o[0] != 255 || enc.fKs57NhW6o[1] != 255 {
				t.Fatal("16-bit wrap")
			}
			enc.Encrypt(dst, plain)
			if enc.pufJF2 != 1 || enc.fKs57NhW6o[0] != 0 || enc.fKs57NhW6o[1] != 0 {
				t.Fatal("wrapped nonce")
			}
			if dec.GetDecryptLen(nil) != -16 {
				t.Fatal("length helper must not clamp")
			}
		})
	}
	var identity P8gDoqFWbS
	if n, e := identity.Encrypt(make([]byte, 2), []byte("abcd")); n != 2 || e != nil {
		t.Fatal(n, e)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("original empty-destination panic")
		}
	}()
	identity.Decrypt(nil, []byte{1})
}
