package cMTGNnWy_test

import (
	_ "recovered.local/soga/CsBw7EWKCp"
	_ "recovered.local/soga/HULJhSRw0LH"
	P "recovered.local/soga/cMTGNnWy"
	C "recovered.local/soga/cNT24rG4F"
	_ "recovered.local/soga/mQFF7Rd7"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	_ "recovered.local/soga/tNCtf5NWXas2"
	_ "recovered.local/soga/tWaD46z"
	"testing"
)

func TestFiveOriginalRegistrations(t *testing.T) {
	for _, protocol := range []C.Mpn2JltM{"hysteria", "ss", "trojan", "vless", "v2ray"} {
		config := &C.FCEAfQJwDA7e{OS91X026uML: "ppanel", ZyAhkq5th8bf: protocol, F3PCiaIra: "webapi"}
		panel, e := Q.SMkpVNb(config)
		if e != nil {
			t.Fatal(e)
		}
		if _, ok := panel.(*P.Zkt5w_2o); !ok || panel.GetAPIConfig() != config {
			t.Fatal(panel)
		}
	}
}
