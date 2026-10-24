package MFFa9yT

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	L "recovered.local/soga/LFWTTXF7JR"
	U "recovered.local/soga/nzATvajt"
	"testing"
	"unsafe"
)

func trojanPanelFixture(id int64, password string) *L.Aga91y {
	return &L.Aga91y{WZ89pTAOrHK4: id, JAfXWwzYrW: &L.JxTGUar{ARARP1z0yA: password}}
}
func TestTrojanNativeUsersAuthAndStop(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	var layout EfaIFpp
	if unsafe.Sizeof(layout) != 80 || unsafe.Offsetof(layout.d9Cxrbh1g7) != 32 {
		t.Fatal("RTTI")
	}
	b, err := n2s6Gf9ey(&DV_XPJ4z0{Dmpl48sXM: context.Background()})
	if err != nil {
		t.Fatal(err)
	}
	defer b.yPUbKO()
	v := &EfaIFpp{JWN154qJ4: b}
	v.spcNlUL(L.Sxcf3wu3L{trojanPanelFixture(7, "first")})
	base, ok := v.jc_QPN6.Get(7)
	if !ok {
		t.Fatal("create")
	}
	u := base.(*U.DWaKIV1T)
	digest := sha256.Sum224([]byte("first"))
	auth := []byte(hex.EncodeToString(digest[:]))
	c := &vmessReadFixture{reads: [][]byte{nil, auth}}
	if got, fallback := v.s3mqrqd_(c); got != u || fallback != nil {
		t.Fatal("valid auth")
	}
	if len(c.deadlines) != 2 || !c.deadlines[1].IsZero() {
		t.Fatal("deadline reset")
	}
	before := u.GetCtx()
	oldHash := u.IdeY3L
	v.spcNlUL(L.Sxcf3wu3L{trojanPanelFixture(7, "second")})
	if before.Err() != context.Canceled {
		t.Fatal("password change session cancellation")
	}
	if _, ok := v.d9Cxrbh1g7.Load(oldHash); ok {
		t.Fatal("old hash retained")
	}
	if got, _ := v.d9Cxrbh1g7.Load(u.IdeY3L); got != u {
		t.Fatal("new hash")
	}
	// 原existing分支即使收到错误协议类型仍保留认证条目，并更新共同用户字段。
	wrong := panelVless(7, "ignored")
	wrong.LmOKwjha = 12
	v.spcNlUL(L.Sxcf3wu3L{wrong})
	if v.jc_QPN6.Len() != 1 || u.GetSpeedLimit() != 12 {
		t.Fatal("invalid protocol update behavior")
	}
	before = u.GetCtx()
	v.spcNlUL(nil)
	if before.Err() != context.Canceled || v.jc_QPN6.Len() != 0 {
		t.Fatal("removal")
	}
	if _, ok := v.d9Cxrbh1g7.Load(u.IdeY3L); ok {
		t.Fatal("removed auth index")
	}
	v.spcNlUL(L.Sxcf3wu3L{trojanPanelFixture(7, "third")})
	base, _ = v.jc_QPN6.Get(7)
	u = base.(*U.DWaKIV1T)
	before = u.GetCtx()
	v.Stop()
	if v.jc_QPN6.Len() != 1 || before.Err() != context.Canceled {
		t.Fatal("stop retains table")
	}
	if _, ok := v.d9Cxrbh1g7.Load(u.IdeY3L); !ok {
		t.Fatal("stop retains index")
	}
}
func TestTrojanNativeSingleReadFallback(t *testing.T) {
	var v EfaIFpp
	auth := []byte("12345678901234567890123456789012345678901234567890123456")
	for _, n := range []int{0, 1, 55, 56} {
		c := &vmessReadFixture{reads: [][]byte{auth[:n]}}
		if n == 0 {
			c.readErr = io.EOF
		}
		got, fallback := v.s3mqrqd_(c)
		if got != nil || string(fallback) != string(auth[:n]) {
			t.Fatal("fallback", n)
		}
	}
	c := &vmessReadFixture{reads: [][]byte{auth}, readErr: io.EOF}
	u := &U.DWaKIV1T{}
	v.d9Cxrbh1g7.Store([56]byte(auth), u)
	if got, fallback := v.s3mqrqd_(c); got != nil || string(fallback) != string(auth) {
		t.Fatal("data with EOF must not authenticate")
	}
}
