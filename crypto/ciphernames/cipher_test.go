package ciphernames

import (
	"bytes"
	"errors"
	"net"
	stream "recovered.local/soga/transport/wserr"
	aead "recovered.local/soga/crypto/aead"
	"testing"
	"unsafe"
)

func TestFactoryAliasesKeysAndWrappers(t *testing.T) {
	for _, name := range []string{"AEAD_AES_128_GCM", "AEAD_AES_192_GCM", "AEAD_AES_256_GCM", "AEAD_CHACHA20_POLY1305", "AEAD_XCHACHA20_POLY1305", "aes-128-gcm", "AES-192-GCM", "AES-256-GCM", "chacha20-ietf-poly1305", "xchacha20-ietf-poly1305", "RC4-MD5", "AES-128-CTR", "AES-192-CTR", "AES-256-CTR", "AES-128-CFB", "AES-192-CFB", "AES-256-CFB", "CHACHA20-IETF", "XCHACHA20"} {
		c, e := B93MVPtD(name, nil, "pass")
		if e != nil {
			t.Fatal(name, e)
		}
		switch v := c.(type) {
		case *Dc5Zt0:
			if !bytes.Equal(v.Ez2IY3fv, ArNrGS_EV("pass", v.KeySize())) {
				t.Fatal(name)
			}
			if v.StreamConn(nil).(*aead.AZVSCxZmZQ).HugVbq != v || v.PacketConn(nil).(*aead.Zh37aRKMK).HugVbq != v {
				t.Fatal("dynamic cipher")
			}
		case *E3HaL6A8:
			if v.StreamConn(nil).(*stream.AriwPpia).ZDHlm7vHo != v || v.PacketConn(nil).(*stream.HtnJMv4Lsng).ZDHlm7vHo != v {
				t.Fatal("dynamic cipher")
			}
		default:
			t.Fatalf("%T", c)
		}
	}
	for _, name := range []string{"none", "dummy", "plain"} {
		c, e := B93MVPtD(name, []byte{1}, "")
		if e != nil {
			t.Fatal(e)
		}
		var raw net.Conn
		if c.StreamConn(raw) != raw {
			t.Fatal("identity")
		}
	}
	if _, e := B93MVPtD(" AES-128-GCM", nil, ""); e != errorAt5098510 {
		t.Fatal(e)
	}
	if _, e := B93MVPtD("AES-128-GCM", []byte{1}, ""); e != aead.ZPe9opSsSr9k(16) {
		t.Fatal(e)
	}
	if _, e := B93MVPtD("AES-128-CTR", []byte{1}, ""); e != stream.XJzap3(16) {
		t.Fatal(e)
	}
	key := make([]byte, 16)
	c, _ := B93MVPtD("AES-128-GCM", key, "")
	key[0] = 99
	if c.(*Dc5Zt0).Ez2IY3fv[0] != 99 {
		t.Fatal("key copied")
	}
	if unsafe.Sizeof(Dc5Zt0{}) != 40 || unsafe.Sizeof(E3HaL6A8{}) != 40 {
		t.Fatal("RTTI")
	}
}
func TestConstructorErrorStillReturnsWrapper(t *testing.T) {
	entry := aeadAt5096498["AEAD_AES_128_GCM"]
	defer func() { aeadAt5096498["AEAD_AES_128_GCM"] = entry }()
	sentinel := errors.New("constructor")
	entry.NHaagjNZeaDf = func([]byte) (aead.HugVbq, error) { return nil, sentinel }
	aeadAt5096498["AEAD_AES_128_GCM"] = entry
	c, e := B93MVPtD("AES-128-GCM", nil, "")
	if e != sentinel || c == nil || c.(*Dc5Zt0).HugVbq != nil {
		t.Fatal(c, e)
	}
}
