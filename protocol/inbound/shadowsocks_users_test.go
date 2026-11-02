package inbound

import (
	"context"
	"net"
	"recovered.local/soga/panel/types"
	"recovered.local/soga/config"
	"testing"
	"unsafe"
)

func ssPanelFixture(id int64, password string) *types.Aga91y {
	return &types.Aga91y{WZ89pTAOrHK4: id, JAfXWwzYrW: &types.YlfOV0DBBnx{T1oDDIgYJHh: "aes-128-gcm", A_wNUTkMh6pH: password}}
}
func TestShadowsocksOriginalUsersAndStop(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	var layout Lp07FzBLx9V
	if unsafe.Sizeof(layout) != 80 || unsafe.Offsetof(layout.mCYkzsUd) != 72 {
		t.Fatal("RTTI")
	}
	b, e := n2s6Gf9ey(&DV_XPJ4z0{Dmpl48sXM: context.Background()})
	if e != nil {
		t.Fatal(e)
	}
	defer b.yPUbKO()
	v := &Lp07FzBLx9V{JWN154qJ4: b}
	v.spcNlUL(types.Sxcf3wu3L{ssPanelFixture(1, "first")})
	u, ok := v.jc_QPN6.Get(1)
	if !ok {
		t.Fatal("create")
	}
	old := u.GetCtx()
	v.spcNlUL(types.Sxcf3wu3L{ssPanelFixture(1, "second")})
	if old.Err() != context.Canceled {
		t.Fatal("changed password retained old context")
	}
	// The native SS branch removes an existing user after a bad panel type.
	old = u.GetCtx()
	v.spcNlUL(types.Sxcf3wu3L{panelVless(1, "invalid")})
	if _, ok = v.jc_QPN6.Get(1); ok || old.Err() != context.Canceled || v.fUxO7UYAH != 1 {
		t.Fatal("failed update removal/count")
	}
	v.spcNlUL(types.Sxcf3wu3L{ssPanelFixture(2, "second")})
	u, _ = v.jc_QPN6.Get(2)
	old = u.GetCtx()
	l, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer l.Close()
	v.vLRjbS7NDeZZ = []net.Listener{l}
	v.Stop()
	if _, e = l.Accept(); e == nil || old.Err() != context.Canceled || v.jc_QPN6.Len() != 1 || len(v.vLRjbS7NDeZZ) != 1 {
		t.Fatal("stop")
	}
}
func TestShadowsocksOriginalRestrictedCeiling(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	before := config.Q0SDIN()
	config.T15Jifo(true)
	defer config.T15Jifo(before)
	b, e := n2s6Gf9ey(&DV_XPJ4z0{Dmpl48sXM: context.Background()})
	if e != nil {
		t.Fatal(e)
	}
	defer b.yPUbKO()
	v := &Lp07FzBLx9V{JWN154qJ4: b}
	panels := make(types.Sxcf3wu3L, 90)
	for i := range panels {
		panels[i] = ssPanelFixture(int64(i), "password")
	}
	v.spcNlUL(panels)
	if v.fUxO7UYAH != 90 || v.jc_QPN6.Len() != 90 || userCountAt50C6DA0 != 90 {
		t.Fatal("user limit must be removed")
	}
}
