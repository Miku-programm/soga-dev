package HakjBGjG_test

import (
	_ "recovered.local/soga/HakjBGjG"
	_ "recovered.local/soga/HhkySK"
	C "recovered.local/soga/cNT24rG4F"
	_ "recovered.local/soga/mOR0KqC"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	U "recovered.local/soga/uazGa0Ap"
	_ "recovered.local/soga/ygzRkd5EN"
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
