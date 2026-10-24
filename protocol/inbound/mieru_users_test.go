package inbound

import (
	"context"
	L "recovered.local/soga/panel/types"
	N "recovered.local/soga/net/limit"
	U "recovered.local/soga/user"
	"testing"
)

func TestMieruInvalidUpdateRemovesUserAndCountsPanelRows(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	base, e := n2s6Gf9ey(&DV_XPJ4z0{Dmpl48sXM: context.Background()})
	if e != nil {
		t.Fatal(e)
	}
	defer base.yPUbKO()
	v := &ONGT2EHM40r{JWN154qJ4: base, xVMggjOZQW: &N.O0UldNL4Eq{RCaEGDlm: base.zb8f3cHWW8v, RacA3he0TX: base.jc_QPN6}}
	p := &L.Aga91y{WZ89pTAOrHK4: 1, JAfXWwzYrW: &L.DO9qqM9{Ejujyaj: "u", GltkJGvV: "p"}}
	v.spcNlUL(L.Sxcf3wu3L{p})
	u, ok := v.jc_QPN6.Get(1)
	if !ok {
		t.Fatal("user missing")
	}
	ctx := u.GetCtx()
	// Successful changed credentials keep the actual user object and cancel its
	// former context; an invalid update is omitted from seen and then removed.
	changed := *p
	changed.JAfXWwzYrW = &L.DO9qqM9{Ejujyaj: "u2", GltkJGvV: "p2"}
	v.spcNlUL(L.Sxcf3wu3L{&changed})
	now, _ := v.jc_QPN6.Get(1)
	if now != u || now.(*U.WIZQ28).GetPassword() != "p2" {
		t.Fatal("replacement/update")
	}
	select {
	case <-ctx.Done():
	default:
		t.Fatal("old context active")
	}
	ctx = u.GetCtx()
	invalid := *p
	invalid.JAfXWwzYrW = nil
	v.spcNlUL(L.Sxcf3wu3L{&invalid})
	if _, ok = v.jc_QPN6.Get(1); ok {
		t.Fatal("invalid retained")
	}
	if v.fUxO7UYAH != 1 {
		t.Fatal("count should use incoming rows")
	}
	select {
	case <-ctx.Done():
	default:
		t.Fatal("removed user context active")
	}
	v.spcNlUL(L.Sxcf3wu3L{p})
	u, _ = v.jc_QPN6.Get(1)
	ctx = u.GetCtx()
	v.Stop()
	if _, ok = v.jc_QPN6.Get(1); !ok {
		t.Fatal("Stop must retain map")
	}
	select {
	case <-ctx.Done():
	default:
		t.Fatal("Stop context")
	}
}
func TestMieruStartRejectsUDPTransportBeforeListen(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := &ONGT2EHM40r{JWN154qJ4: &JWN154qJ4{}}
	v.wiy6P3.Store(&L.Wr5SVP{ABA4hq: L.Xh4IiOa0{GGSLzD: "uDp"}})
	e := v.Start()
	if e == nil || e.Error() != "unsupported mieru transport: uDp (only TCP is supported)" {
		t.Fatal(e)
	}
	if len(v.d7Fav9U0ykgg) != 0 {
		t.Fatal("listened")
	}
}
