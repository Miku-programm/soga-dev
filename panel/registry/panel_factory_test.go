package registry_test

import (
	B "recovered.local/soga/panel/xboard"
	_ "recovered.local/soga/panel/reg/xboardmieruwebapi"
	_ "recovered.local/soga/panel/reg/xboardv2raywebapi"
	_ "recovered.local/soga/panel/reg/xboardvlesswebapi"
	_ "recovered.local/soga/panel/reg/xboardtrojanwebapi"
	_ "recovered.local/soga/panel/reg/xboardhysteriawebapi"
	_ "recovered.local/soga/panel/reg/xboardsswebapi"
	_ "recovered.local/soga/panel/reg/xboardanytlswebapi"

	_ "recovered.local/soga/panel/reg/xiaov2boardvlesswebapi"
	_ "recovered.local/soga/panel/reg/xiaov2boardv2raywebapi"
	_ "recovered.local/soga/panel/reg/xiaov2boardhysteriawebapi"
	A "recovered.local/soga/panel/xiaov2board"
	_ "recovered.local/soga/panel/reg/xiaov2boardtrojanwebapi"
	_ "recovered.local/soga/panel/reg/xiaov2boardsswebapi"
	_ "recovered.local/soga/panel/reg/xiaov2boardanytlswebapi"

	_ "recovered.local/soga/panel/reg/v2boardv2raywebapi"
	Z "recovered.local/soga/panel/webapibase"
	_ "recovered.local/soga/panel/reg/v2boardtrojanwebapi"
	C "recovered.local/soga/config"
	P "recovered.local/soga/panel/registry"
	_ "recovered.local/soga/panel/reg/v2boardsswebapi"
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
