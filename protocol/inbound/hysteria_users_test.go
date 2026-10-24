package inbound

import (
	"context"
	L "recovered.local/soga/panel/types"
	U "recovered.local/soga/user"
	"testing"
)

func TestHysteriaPasswordIndexUpdateRemovalAndStop(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	base, e := n2s6Gf9ey(&DV_XPJ4z0{Dmpl48sXM: context.Background()})
	if e != nil {
		t.Fatal(e)
	}
	defer base.yPUbKO()
	v := &Aq5i0Bwez8{JWN154qJ4: base}
	p := &L.Aga91y{WZ89pTAOrHK4: 91, JAfXWwzYrW: &L.WQM2tKA{CBWeRfrO1ahs: "first"}}
	v.spcNlUL(L.Sxcf3wu3L{p})
	u, ok := v.jc_QPN6.Get(91)
	if !ok {
		t.Fatal("missing user")
	}
	ctx := u.GetCtx()
	if got, ok := v.Authenticate(nil, "first", 999); !ok || got != u {
		t.Fatal("password lookup")
	}
	if _, ok := v.Authenticate(nil, "wrong", 0); ok {
		t.Fatal("wrong password accepted")
	}
	v.spcNlUL(L.Sxcf3wu3L{p})
	if u.GetCtx() != ctx {
		t.Fatal("unchanged update reset context")
	}
	changed := *p
	changed.JAfXWwzYrW = &L.WQM2tKA{CBWeRfrO1ahs: "second"}
	v.spcNlUL(L.Sxcf3wu3L{&changed})
	now, _ := v.jc_QPN6.Get(91)
	if now != u || now.(*U.V0wEAoUPFk).GetPassword() != "second" {
		t.Fatal("user object update")
	}
	select {
	case <-ctx.Done():
	default:
		t.Fatal("former context not cancelled")
	}
	if _, ok := v.Authenticate(nil, "first", 0); ok {
		t.Fatal("old password remains")
	}
	if got, ok := v.Authenticate(nil, "second", 0); !ok || got != u {
		t.Fatal("new password missing")
	}
	ctx = u.GetCtx()
	invalid := *p
	invalid.JAfXWwzYrW = nil
	v.spcNlUL(L.Sxcf3wu3L{&invalid})
	if _, ok := v.jc_QPN6.Get(91); ok {
		t.Fatal("invalid updated user retained")
	}
	if _, ok := v.Authenticate(nil, "second", 0); ok {
		t.Fatal("invalid user's password retained")
	}
	select {
	case <-ctx.Done():
	default:
		t.Fatal("removed context not cancelled")
	}
	if v.fUxO7UYAH != 1 {
		t.Fatal("original count must use panel rows")
	}
	v.spcNlUL(L.Sxcf3wu3L{p})
	u, _ = v.jc_QPN6.Get(91)
	ctx = u.GetCtx()
	v.Stop()
	select {
	case <-ctx.Done():
	default:
		t.Fatal("Stop user context")
	}
	if got, ok := v.Authenticate(nil, "first", 0); !ok || got != u {
		t.Fatal("Stop must retain password index")
	}
	v.spcNlUL(nil)
	if _, ok := v.Authenticate(nil, "first", 0); ok {
		t.Fatal("nil panel did not remove password")
	}
	if v.fUxO7UYAH != 0 {
		t.Fatal("empty count")
	}
}
func TestHysteriaUnsupportedVersionBeforeTLSAndListen(t *testing.T) {
	if !vlessFixtureProcess(t) {
		return
	}
	v := &Aq5i0Bwez8{JWN154qJ4: &JWN154qJ4{}}
	v.wiy6P3.Store(&L.Wr5SVP{AYf5kQSQW: L.ZMYytVS{Li6nvK1: 1}})
	e := v.Start()
	if e == nil || e.Error() != "not support hysteria version 1" {
		t.Fatal(e)
	}
	if len(v.uqbfp_xRjia) != 0 {
		t.Fatal("listened")
	}
}
