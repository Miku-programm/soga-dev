package w2YaTd9_

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"lukechampine.com/blake3"
	"strings"
	"testing"
	"unsafe"
)

func TestSubkeyStateLayoutAndOriginalCipherBranches(t *testing.T) {
	var s MXXLkaPB
	if unsafe.Sizeof(s) != 352 || unsafe.Offsetof(s.gQWRfV0dgL) != 164 || unsafe.Offsetof(s.lGjGYnir7l_L) != 328 {
		t.Fatal("state layout")
	}
	for _, name := range []C6q3Upu{"aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "chacha20-ietf-poly1305", "2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm"} {
		n, err := name.KeyLen()
		if err != nil {
			t.Fatal(err)
		}
		if salt, _ := name.SaltLen(); salt != n || name.NonceLen() != 12 || name.Overhead() != 16 {
			t.Fatal("cipher dimensions")
		}
		a, err := name.NewAEAD(make([]byte, n))
		if err != nil {
			t.Fatal(err)
		}
		sealed := a.Seal(nil, make([]byte, 12), []byte("payload"), []byte("header"))
		plain, err := a.Open(nil, make([]byte, 12), sealed, []byte("header"))
		if err != nil || string(plain) != "payload" {
			t.Fatal("AEAD round trip")
		}
	}
	// AES constructor accepts any AES key length, regardless of the name's KeyLen.
	if _, err := C6q3Upu("aes-128-gcm").NewAEAD(make([]byte, 32)); err != nil {
		t.Fatal("extra key validation")
	}
	if _, err := C6q3Upu("unknown").NewBlock(make([]byte, 16)); err != nil {
		t.Fatal("NewBlock must ignore cipher name")
	}
	if a, err := C6q3Upu("unknown").NewAEAD(nil); a != nil || err.Error() != "unknown ss cipher: unknown" {
		t.Fatal(err)
	}
	a, _ := C6q3Upu("aes-128-gcm").NewAEAD(make([]byte, 16))
	if got := hex.EncodeToString(a.Seal(nil, make([]byte, 12), make([]byte, 16), nil)); got != "0388dace60b6a392f328c2b971b2fe78ab6e47d42cec13bdf53a67b21257bddf" {
		t.Fatal(got)
	}
}
func TestPasswordDerivationAndErrors(t *testing.T) {
	c := C6q3Upu("aes-256-gcm")
	key, err := c.PSK("password")
	if err != nil || hex.EncodeToString(key) != "5f4dcc3b5aa765d61d8327deb882cf992b95990a9151374abd8ff8c5a7a0fe08" {
		t.Fatalf("PSK %x %v", key, err)
	}
	if k, err := C6q3Upu("unknown").PSK(""); k != nil || err == nil {
		t.Fatal("unknown cipher")
	}
	for _, c := range []C6q3Upu{"2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm"} {
		n, _ := c.KeyLen()
		raw := bytes.Repeat([]byte{7}, n)
		encoded := base64.StdEncoding.EncodeToString(raw)
		key, err := c.PSK(encoded)
		if err != nil || !bytes.Equal(key, raw) {
			t.Fatal("base64 PSK")
		}
		raw = append(raw, 9)
		key, err = c.PSK(base64.StdEncoding.EncodeToString(raw))
		digest := sha256.Sum256(raw)
		if err != nil || !bytes.Equal(key, digest[:n]) || cap(key) != 32 {
			t.Fatal("overlength PSK hashing")
		}
		if k, err := c.PSK(""); k != nil || err.Error() != "empty password" {
			t.Fatal(err)
		}
		if k, err := c.PSK("!"); k != nil || !strings.HasPrefix(err.Error(), "decode ss2022 password <!> err: ") {
			t.Fatal(err)
		}
		if k, err := c.PSK("AA=="); k != nil || !strings.HasPrefix(err.Error(), "ss2022 password key length too short: 1 < ") {
			t.Fatal(err)
		}
	}
}
func TestSubkeyIndependentHKDFVectorsAndReuse(t *testing.T) {
	key, _ := hex.DecodeString("5f4dcc3b5aa765d61d8327deb882cf992b95990a9151374abd8ff8c5a7a0fe08")
	salt := make([]byte, 16)
	for i := range salt {
		salt[i] = byte(i)
	}
	cases := []struct {
		salt        []byte
		short, long string
	}{
		{salt, "21b989020f67b371530ef9a3818f30db2dc789e804d948098dc5a91ec6fe382f", "93fa25dff664ff9bec3a78cf18ba97794cd8864b208a53350d31cabb9970c163"},
		{bytes.Repeat([]byte("A"), 80), "2767aefa2022872081158a7942323623c457984f1bc6c74e3d6998d38dca6ca8", "59f8369b296f01a0155f15382db40d8551ea3831fb6919b2b159ff926a5fe626"},
	}
	c := C6q3Upu("aes-256-gcm")
	state := SaYhx5xog9()
	for _, fixture := range cases {
		out := make([]byte, 5121)
		c.SubKeyTo(state, out, key, fixture.salt)
		sum := sha256.Sum256(out)
		if hex.EncodeToString(out[:32]) != fixture.short || hex.EncodeToString(sum[:]) != fixture.long {
			t.Fatalf("HKDF %x %x", out[:32], sum)
		}
		if got := c.SubKey(key, fixture.salt); !bytes.Equal(got, out[:32]) {
			t.Fatal("allocated subkey")
		}
		c.SubKeyTo(state, state.KeyBuf(32), key, fixture.salt)
		if !bytes.Equal(state.KeyBuf(32), out[:32]) {
			t.Fatal("same salt cache reuse")
		}
		// Cached pads were overwritten by expansion: the serialized extraction states
		// must still recover the correct result for a different key with the same salt.
		changed := []byte("different key")
		c.SubKeyTo(state, state.KeyBuf(20), changed, fixture.salt)
		extract := hmac.New(sha1.New, fixture.salt)
		extract.Write(changed)
		expand := hmac.New(sha1.New, extract.Sum(nil))
		expand.Write([]byte("ss-subkey\x01"))
		if !bytes.Equal(state.KeyBuf(20), expand.Sum(nil)) {
			t.Fatal("cached extraction state")
		}
	}
	if state.KeyBuf(0) == nil {
		t.Fatal("empty KeyBuf must alias fixed array")
	}
}
func TestSubkeyBlake3AndBounds(t *testing.T) {
	key := bytes.Repeat([]byte{5}, 32)
	salt := bytes.Repeat([]byte{7}, 16)
	c := C6q3Upu("2022-blake3-aes-256-gcm")
	want := make([]byte, 32)
	blake3.DeriveKey(want, "shadowsocks 2022 session subkey", append(append([]byte(nil), key...), salt...))
	if got := c.SubKey(key, salt); !bytes.Equal(got, want) {
		t.Fatal("blake3 material")
	}
	s := SaYhx5xog9()
	c.SubKeyTo(s, s.KeyBuf(32), key, salt)
	if !bytes.Equal(s.KeyBuf(32), want) || !bytes.Equal(s.lGjGYnir7l_L, append(key, salt...)) {
		t.Fatal("blake3 state")
	}
	for _, n := range []int{-1, 33} {
		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("KeyBuf(%d) accepted", n)
				}
			}()
			s.KeyBuf(n)
		}()
	}
}
