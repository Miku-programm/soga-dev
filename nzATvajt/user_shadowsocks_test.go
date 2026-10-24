package nzATvajt

import (
	"bytes"
	"encoding/base64"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/w2YaTd9_"
	"strings"
	"testing"
	"unsafe"
)

func TestShadowsocksCredentialsAndBothAEADPaths(t *testing.T) {
	var layout WtK9HQSy66
	if unsafe.Sizeof(layout) != 80 || unsafe.Offsetof(layout.n5IKaq) != 40 || unsafe.Offsetof(layout.we1MNFya) != 64 {
		t.Fatal("user layout")
	}
	for _, method := range []w2YaTd9_.C6q3Upu{"aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "chacha20-ietf-poly1305", "2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm"} {
		password := "password"
		n, _ := method.KeyLen()
		if method.Is2022() {
			password = base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{3}, n))
		}
		p := &LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 48, JAfXWwzYrW: &LFWTTXF7JR.YlfOV0DBBnx{T1oDDIgYJHh: method, A_wNUTkMh6pH: password}}
		u, err := MbpJoYH(p, nil)
		if err != nil {
			t.Fatal(err)
		}
		if u.GetCipher() != method || len(u.PSK()) != n || (u.GetBlock() != nil) != method.Is2022() {
			t.Fatal("constructor", method)
		}
		salt := make([]byte, n)
		a, err := u.NewAEAD(salt)
		if err != nil {
			t.Fatal(err)
		}
		b, err := u.NewAEADWith(w2YaTd9_.SaYhx5xog9(), salt)
		if err != nil {
			t.Fatal(err)
		}
		sealed := a.Seal(nil, make([]byte, 12), []byte("authentication fixture"), []byte("ad"))
		opened, err := b.Open(nil, make([]byte, 12), sealed, []byte("ad"))
		if err != nil || string(opened) != "authentication fixture" {
			t.Fatal("AEAD path mismatch", method, err)
		}
		if u.String() != "id: 48, method: "+string(method)+", password: "+password {
			t.Fatal(u.String())
		}
	}
}
func TestShadowsocksFailedUpdateRetainsKeyButStoresPassword(t *testing.T) {
	credential := &LFWTTXF7JR.YlfOV0DBBnx{T1oDDIgYJHh: "2022-blake3-aes-128-gcm", A_wNUTkMh6pH: base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{1}, 16))}
	p := &LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 49, JAfXWwzYrW: credential}
	u, err := MbpJoYH(p, nil)
	if err != nil {
		t.Fatal(err)
	}
	key := u.PSK()
	block := u.GetBlock()
	ctx := u.GetCtx()
	credential.A_wNUTkMh6pH = "invalid base64"
	if changed, err := u.UpdatePanelUser(p); changed || err == nil {
		t.Fatal("invalid update accepted")
	}
	if u.bRxpsrh13ry != "invalid base64" || !bytes.Equal(key, u.PSK()) || u.GetBlock() != block || u.GetCtx() != ctx {
		t.Fatal("failure mutation order")
	}
	// Repeating the failed credential does not retry the derivation.
	if changed, err := u.UpdatePanelUser(p); changed || err != nil {
		t.Fatal("failed credential retried")
	}
	credential.A_wNUTkMh6pH = base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{2}, 16))
	credential.T1oDDIgYJHh = "aes-256-gcm"
	if changed, err := u.UpdatePanelUser(p); !changed || err != nil || u.GetCipher() != "2022-blake3-aes-128-gcm" || len(u.PSK()) != 16 || u.GetCtx() == ctx {
		t.Fatal("original cipher must stay fixed")
	}
	oldkey := u.PSK()
	oldkey[0] = 77
	if u.PSK()[0] != 77 {
		t.Fatal("PSK getter must return live slice")
	}
	p.JAfXWwzYrW = &LFWTTXF7JR.JxTGUar{}
	p.Xmrna8 = 15
	ctx = u.GetCtx()
	if changed, err := u.UpdatePanelUser(p); changed || err == nil || u.GetConnLimit() != 15 || u.GetCtx() != ctx {
		t.Fatal("wrong protocol mutation order")
	}
}
func TestShadowsocksInitialEmptyPasswordDoesNotDerive(t *testing.T) {
	p := &LFWTTXF7JR.Aga91y{WZ89pTAOrHK4: 50, JAfXWwzYrW: &LFWTTXF7JR.YlfOV0DBBnx{T1oDDIgYJHh: "unknown"}}
	u, err := MbpJoYH(p, nil)
	if err != nil || u.PSK() != nil {
		t.Fatal("initial empty branch")
	}
	if a, err := u.NewAEAD(nil); a != nil || err.Error() != "user 50 new aead err: unknown ss cipher: unknown" {
		t.Fatal(err)
	}
	if a, err := u.NewAEADWith(w2YaTd9_.SaYhx5xog9(), nil); a != nil || err.Error() != "user 50 new aead err: unknown ss cipher: unknown" {
		t.Fatal(err)
	}
	p.JAfXWwzYrW = &LFWTTXF7JR.YlfOV0DBBnx{T1oDDIgYJHh: "unknown", A_wNUTkMh6pH: "x"}
	if u, err := MbpJoYH(p, nil); u != nil || !strings.Contains(err.Error(), "unknown ss cipher <unknown> key len") {
		t.Fatal(err)
	}
}
