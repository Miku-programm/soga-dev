package v2raysockssswebapi_test

import (
	_ "recovered.local/soga/panel/reg/v2raysockssswebapi"
	_ "recovered.local/soga/panel/reg/v2raysocksv2raywebapi"
	C "recovered.local/soga/config"
	_ "recovered.local/soga/panel/reg/v2raysocksvlesswebapi"
	Q "recovered.local/soga/panel/registry"
	U "recovered.local/soga/panel/v2raysocksnode"
	_ "recovered.local/soga/panel/reg/v2raysockstrojanwebapi"
	"testing"
)

func TestOriginalFourFactoryRegistrations(t *testing.T) {
	for _, kind := range []C.Mpn2JltM{"ss", "trojan", "vless", "v2ray"} {
		c := &C.FCEAfQJwDA7e{OS91X026uML: "v2raysocks", F3PCiaIra: "webapi", ZyAhkq5th8bf: kind}
		p, e := Q.SMkpVNb(c)
		if e != nil {
			t.Fatal(e)
		}
		if _, ok := p.(*U.W3tZSymoah); !ok || p.GetAPIConfig() != c {
			t.Fatal(kind, p)
		}
	}
	c := &C.FCEAfQJwDA7e{OS91X026uML: "v2raysocks", F3PCiaIra: "webapi", ZyAhkq5th8bf: "hysteria"}
	if _, e := Q.SMkpVNb(c); e == nil {
		t.Fatal("unregistered hysteria must not be fabricated")
	}
}
