package q5bqCgTdxaTM_test

import (
	B "recovered.local/soga/BP3PTlx"
	_ "recovered.local/soga/HbH5rTkHd7zj"
	_ "recovered.local/soga/W_EMft"
	_ "recovered.local/soga/aJU33bUzb"
	_ "recovered.local/soga/iiKUPoplen"
	_ "recovered.local/soga/k9EI4fqL9"
	_ "recovered.local/soga/rusgvYjAAU5"
	_ "recovered.local/soga/z0244yEhS6W"

	_ "recovered.local/soga/FkyU22TRjL_B"
	_ "recovered.local/soga/MkeN6dA737wC"
	_ "recovered.local/soga/V9HXHRH7Fig"
	A "recovered.local/soga/aRZNjEvm"
	_ "recovered.local/soga/bjHvRkRuyWRt"
	_ "recovered.local/soga/pcrbn5"
	_ "recovered.local/soga/qx2LmdvLIV"

	_ "recovered.local/soga/MDpSu9EFoD"
	Z "recovered.local/soga/ZrjfjK8Y_tM9"
	_ "recovered.local/soga/_lslMIG7j"
	C "recovered.local/soga/cNT24rG4F"
	P "recovered.local/soga/q5bqCgTdxaTM"
	_ "recovered.local/soga/t3wvrixO2e7V"
	"testing"
)

func TestPanelFactoryOriginalV2BoardRegistrations(t *testing.T) {
	for _, protocol := range []C.Mpn2JltM{"ss", "trojan", "v2ray"} {
		cfg := &C.FCEAfQJwDA7e{OS91X026uML: "v2board", ZyAhkq5th8bf: protocol, F3PCiaIra: "webapi"}
		p, err := P.SMkpVNb(cfg)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := p.(*Z.JZA3kKksUe); !ok || p.GetAPIConfig() != cfg {
			t.Fatalf("constructor/pointer %T", p)
		}
		second, err := P.SMkpVNb(cfg)
		if err != nil || second == p {
			t.Fatal("constructor must allocate independent instance")
		}
	}
	for _, cfg := range []*C.FCEAfQJwDA7e{
		{OS91X026uML: "V2Board", ZyAhkq5th8bf: "ss", F3PCiaIra: "webapi"},
		{OS91X026uML: "v2board", ZyAhkq5th8bf: "shadowsocks", F3PCiaIra: "webapi"},
		{OS91X026uML: "v2board", ZyAhkq5th8bf: "ss", F3PCiaIra: "db"},
	} {
		p, err := P.SMkpVNb(cfg)
		want := "not support server type <" + string(cfg.ZyAhkq5th8bf) + "> for panel <" + string(cfg.OS91X026uML) + "> using <" + string(cfg.F3PCiaIra) + ">"
		if p != nil || err == nil || err.Error() != want {
			t.Fatal(p, err, want)
		}
	}
}

func TestPanelFactoryOriginalXiaoV2BoardRegistrations(t *testing.T) {
	for _, protocol := range []C.Mpn2JltM{"anytls", "hysteria", "ss", "trojan", "vless", "v2ray"} {
		cfg := &C.FCEAfQJwDA7e{OS91X026uML: "xiaov2board", ZyAhkq5th8bf: protocol, F3PCiaIra: "webapi"}
		p, err := P.SMkpVNb(cfg)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := p.(*A.LRlqOaYmO); !ok || p.GetAPIConfig() != cfg {
			t.Fatalf("%T", p)
		}
	}
}

func TestPanelFactoryOriginalXBoardRegistrations(t *testing.T) {
	for _, protocol := range []C.Mpn2JltM{"anytls", "hysteria", "mieru", "ss", "trojan", "vless", "v2ray"} {
		cfg := &C.FCEAfQJwDA7e{OS91X026uML: "xboard", ZyAhkq5th8bf: protocol, F3PCiaIra: "webapi"}
		p, err := P.SMkpVNb(cfg)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := p.(*B.KlO65BuwOt); !ok || p.GetAPIConfig() != cfg {
			t.Fatalf("%T", p)
		}
	}
}
