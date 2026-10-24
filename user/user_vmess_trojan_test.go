package user

import (
	"bytes"
	"recovered.local/soga/panel/types"
	"strings"
	"testing"
	"unsafe"
)

func TestTrojanAuthSnapshotAndShortRawValue(t *testing.T) {
	var layout DWaKIV1T
	if unsafe.Sizeof(layout) != 72 || unsafe.Offsetof(layout.IdeY3L) != 16 {
		t.Fatal("Trojan layout")
	}
	p := &types.Aga91y{WZ89pTAOrHK4: 27, JAfXWwzYrW: &types.JxTGUar{ARARP1z0yA: "secret", QiUSI27: "ignored"}}
	u, err := KUa7qViw(p, nil)
	if err != nil {
		t.Fatal(err)
	}
	want := "95c7fbca92ac5083afda62a564a3d014fc3b72c9140e3cb99ea6bf12"
	if string(u.IdeY3L[:]) != want || u.String() != "id: 27, password: secret" {
		t.Fatal("Trojan auth")
	}
	ctx := u.GetCtx()
	if old := u.UpdatePanelUser(p); old != nil || u.GetCtx() != ctx {
		t.Fatal("unchanged update")
	}
	c := p.JAfXWwzYrW.(*types.JxTGUar)
	c.ARARP1z0yA = ""
	c.QiUSI27 = "abc"
	old := u.UpdatePanelUser(p)
	if old == nil || string(old[:]) != want || string(u.IdeY3L[:]) != "abc"+want[3:] || u.GetCtx() == ctx {
		t.Fatal("short auth old-tail behavior")
	}
	c.QiUSI27 = strings.Repeat("z", 70)
	old = u.UpdatePanelUser(p)
	if string(u.IdeY3L[:]) != strings.Repeat("z", 56) {
		t.Fatal("auth truncation")
	}
	c.QiUSI27 = ""
	u.UpdatePanelUser(p)
	if string(u.IdeY3L[:]) != strings.Repeat("z", 56) {
		t.Fatal("empty raw auth must retain old bytes")
	}
	p.JAfXWwzYrW = &types.Yn8BJoL{}
	p.Xmrna8 = 19
	ctx = u.GetCtx()
	if u.UpdatePanelUser(p) != nil || u.GetCtx() != ctx || u.GetConnLimit() != 19 {
		t.Fatal("wrong protocol mutation order")
	}
	if wrong, err := KUa7qViw(p, nil); wrong != nil || err == nil {
		t.Fatal("wrong protocol constructor")
	}
}
func TestVMessUserUpdatesAndInvalidUUIDRetention(t *testing.T) {
	var layout QKaIWV2
	if unsafe.Sizeof(layout) != 40 || unsafe.Offsetof(layout.fCmumUdZX6m) != 16 {
		t.Fatal("VMess layout")
	}
	c := &types.Yn8BJoL{SfiEGpJo: "01234567-89ab-cdef-0123-456789abcdef", FvQ2ltKOM7: 2}
	p := &types.Aga91y{WZ89pTAOrHK4: 31, JAfXWwzYrW: c}
	u, err := EOkO09L(p, nil)
	if err != nil {
		t.Fatal(err)
	}
	ids := u.GetUserIds()
	if len(ids) != 3 || cap(ids) != 3 || u.GetBlock() != nil {
		t.Fatal("alter IDs constructor")
	}
	if u.String() != "id: 31, alterId: 2, uuid: 01234567-89ab-cdef-0123-456789abcdef" {
		t.Fatal(u.String())
	}
	ctx := u.GetCtx()
	if old, changed := u.UpdatePanelUser(p); old != nil || changed || u.GetCtx() != ctx {
		t.Fatal("unchanged update")
	}
	c.FvQ2ltKOM7 = 0
	old, changed := u.UpdatePanelUser(p)
	if !changed || len(old) != 3 || old[0] != ids[0] || len(u.GetUserIds()) != 1 || u.GetBlock() == nil || u.GetCtx() == ctx {
		t.Fatal("AEAD transition")
	}
	plaintext := make([]byte, 16)
	encrypted := make([]byte, 16)
	decrypted := make([]byte, 16)
	u.GetBlock().Encrypt(encrypted, plaintext)
	u.GetBlock().Decrypt(decrypted, encrypted)
	if !bytes.Equal(decrypted, plaintext) || bytes.Equal(encrypted, plaintext) {
		t.Fatal("AES block")
	}
	valid := u.GetUserIds()
	c.SfiEGpJo = "bad UUID"
	old, changed = u.UpdatePanelUser(p)
	if !changed || old[0] != valid[0] || u.GetUserIds()[0] != valid[0] || u.nmOHVD.SfiEGpJo != "bad UUID" {
		t.Fatal("invalid UUID must retain identity but update credential snapshot")
	}
	c.FvQ2ltKOM7 = -1
	u.UpdatePanelUser(p)
	if u.GetUserIds() == nil || len(u.GetUserIds()) != 0 || u.GetCmdKey() != nil || u.GetBlock() != nil {
		t.Fatal("negative alterId nonnil empty state")
	}
	p.JAfXWwzYrW = &types.JxTGUar{}
	p.Xmrna8 = 23
	ctx = u.GetCtx()
	if old, changed := u.UpdatePanelUser(p); old != nil || changed || u.GetCtx() != ctx || u.GetConnLimit() != 23 {
		t.Fatal("wrong protocol mutation order")
	}
	if wrong, err := EOkO09L(p, nil); wrong != nil || err == nil {
		t.Fatal("wrong protocol constructor")
	}
}
