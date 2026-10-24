package user

import (
	"crypto/sha256"
	"recovered.local/soga/panel/types"
	"strings"
	"testing"
	"unsafe"
)

func TestAuthUserLayouts(t *testing.T) {
	var a HaB1NEnqcph
	var h V0wEAoUPFk
	var m WIZQ28
	if unsafe.Sizeof(a) != 56 || unsafe.Offsetof(a.w2oeIPBSz) != 24 || unsafe.Sizeof(h) != 24 || unsafe.Sizeof(m) != 40 || unsafe.Offsetof(m.birIwyB_DM) != 24 {
		t.Fatal("original layout changed")
	}
	var _ PgaNd3014U = &a
	var _ PgaNd3014U = &h
	var _ PgaNd3014U = &m
}

func TestAnyTLSAuthUpdateAndContext(t *testing.T) {
	p := &types.Aga91y{WZ89pTAOrHK4: 73, JAfXWwzYrW: &types.ZtMFlBTZXyD{CBWeRfrO1ahs: "secret"}}
	u, err := NZl_wC1qYi(p, nil)
	if err != nil {
		t.Fatal(err)
	}
	want := sha256.Sum256([]byte("secret"))
	if u.GetAuth() != want || u.String() != "id: 73, password: secret" {
		t.Fatal("constructor credential mismatch")
	}
	old := u.GetCtx()
	prior, changed, err := u.UpdatePanelUser(p)
	if err != nil || changed || prior != ([32]byte{}) || u.GetCtx() != old {
		t.Fatal("unchanged credential altered context/auth")
	}
	p.JAfXWwzYrW = &types.ZtMFlBTZXyD{CBWeRfrO1ahs: ""}
	prior, changed, err = u.UpdatePanelUser(p)
	if err != nil || !changed || prior != want || u.GetAuth() != sha256.Sum256(nil) || u.GetCtx() == old {
		t.Fatal("password transition lost old digest/context reset")
	}
	select {
	case <-old.Done():
	default:
		t.Fatal("old context not canceled")
	}
	if u.GetCtx().Err() != nil {
		t.Fatal("new context canceled")
	}
	empty, err := NZl_wC1qYi(p, nil)
	if err != nil || empty.GetAuth() != ([32]byte{}) {
		t.Fatal("initial empty password must keep all-zero auth")
	}
}

func TestHysteriaPasswordTransitions(t *testing.T) {
	p := &types.Aga91y{WZ89pTAOrHK4: 19, JAfXWwzYrW: &types.WQM2tKA{CBWeRfrO1ahs: "abc"}}
	u, err := PB4jVTrQHa(p, nil)
	if err != nil {
		t.Fatal(err)
	}
	if u.GetPassword() != "abc" || u.String() != "id: 19, password: abc" {
		t.Fatal("constructor mismatch")
	}
	old := u.GetCtx()
	if changed, err := u.UpdatePanelUser(p); changed || err != nil || u.GetCtx() != old {
		t.Fatal("unchanged password reset")
	}
	p.JAfXWwzYrW = &types.WQM2tKA{CBWeRfrO1ahs: "xyz"}
	if changed, err := u.UpdatePanelUser(p); !changed || err != nil || u.GetPassword() != "xyz" || u.GetCtx() == old {
		t.Fatal("changed password not applied")
	}
	select {
	case <-old.Done():
	default:
		t.Fatal("old context not canceled")
	}
}

func TestMieruEitherCredentialChange(t *testing.T) {
	p := &types.Aga91y{WZ89pTAOrHK4: 42, JAfXWwzYrW: &types.DO9qqM9{Ejujyaj: "alice", GltkJGvV: "first"}}
	u, err := U1KSvhrlJFSB(p, nil)
	if err != nil {
		t.Fatal(err)
	}
	if u.GetUsername() != "alice" || u.GetPassword() != "first" || u.String() != "id: 42, username: alice, password: first" {
		t.Fatal("constructor mismatch")
	}
	for _, next := range []struct {
		user, password string
		changed        bool
	}{{"alice", "first", false}, {"bob", "first", true}, {"bob", "second", true}, {"bob", "second", false}, {"", "", true}} {
		old := u.GetCtx()
		p.JAfXWwzYrW = &types.DO9qqM9{Ejujyaj: next.user, GltkJGvV: next.password}
		changed, err := u.UpdatePanelUser(p)
		if err != nil || changed != next.changed || u.GetUsername() != next.user || u.GetPassword() != next.password {
			t.Fatalf("transition %+v: %v %v", next, changed, err)
		}
		if (old != u.GetCtx()) != next.changed {
			t.Fatal("context change mismatch")
		}
		if next.changed {
			select {
			case <-old.Done():
			default:
				t.Fatal("old context not canceled")
			}
		}
	}
}

func TestAuthWrongProtocolUpdatesBaseBeforeError(t *testing.T) {
	for _, kind := range []string{"anytls", "hysteria", "mieru"} {
		t.Run(kind, func(t *testing.T) {
			p := &types.Aga91y{WZ89pTAOrHK4: 11}
			var base *AhSFeSua
			var update func(*types.Aga91y) (bool, error)
			switch kind {
			case "anytls":
				p.JAfXWwzYrW = &types.ZtMFlBTZXyD{CBWeRfrO1ahs: "pw"}
				u, err := NZl_wC1qYi(p, nil)
				if err != nil {
					t.Fatal(err)
				}
				base = u.AhSFeSua
				update = func(p *types.Aga91y) (bool, error) {
					prior, changed, err := u.UpdatePanelUser(p)
					if prior != ([32]byte{}) {
						t.Fatal("error returned nonzero previous auth")
					}
					return changed, err
				}
			case "hysteria":
				p.JAfXWwzYrW = &types.WQM2tKA{CBWeRfrO1ahs: "pw"}
				u, err := PB4jVTrQHa(p, nil)
				if err != nil {
					t.Fatal(err)
				}
				base = u.AhSFeSua
				update = u.UpdatePanelUser
			case "mieru":
				p.JAfXWwzYrW = &types.DO9qqM9{Ejujyaj: "u", GltkJGvV: "pw"}
				u, err := U1KSvhrlJFSB(p, nil)
				if err != nil {
					t.Fatal(err)
				}
				base = u.AhSFeSua
				update = u.UpdatePanelUser
			}
			old := base.GetCtx()
			p.WZ89pTAOrHK4 = 12
			p.Xmrna8 = 7
			p.JAfXWwzYrW = nil
			changed, err := update(p)
			if changed || err == nil || !strings.Contains(err.Error(), kind) || !strings.Contains(err.Error(), "12") {
				t.Fatalf("protocol error %v %v", changed, err)
			}
			if base.GetId() != 11 || base.GetConnLimit() != 7 || base.GetCtx() != old {
				t.Fatal("base update/error order changed")
			}
		})
	}
}

func TestAuthConstructorProtocolErrors(t *testing.T) {
	p := &types.Aga91y{WZ89pTAOrHK4: 91}
	a, ae := NZl_wC1qYi(p, nil)
	h, he := PB4jVTrQHa(p, nil)
	m, me := U1KSvhrlJFSB(p, nil)
	if a != nil || h != nil || m != nil || ae == nil || he == nil || me == nil {
		t.Fatal("mismatched payload accepted")
	}
	if !strings.HasPrefix(ae.Error(), "update user 91 failed: ") || !strings.HasPrefix(me.Error(), "update user 91 failed: ") || strings.HasPrefix(he.Error(), "update user") {
		t.Fatalf("constructor error formatting %v / %v / %v", ae, he, me)
	}
}

func TestAuthTypedNilRetainsPanic(t *testing.T) {
	for _, kind := range []string{"anytls", "hysteria", "mieru"} {
		t.Run(kind, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("typed nil payload must panic")
				}
			}()
			p := &types.Aga91y{}
			switch kind {
			case "anytls":
				p.JAfXWwzYrW = (*types.ZtMFlBTZXyD)(nil)
				NZl_wC1qYi(p, nil)
			case "hysteria":
				p.JAfXWwzYrW = (*types.WQM2tKA)(nil)
				PB4jVTrQHa(p, nil)
			case "mieru":
				p.JAfXWwzYrW = (*types.DO9qqM9)(nil)
				U1KSvhrlJFSB(p, nil)
			}
		})
	}
}
