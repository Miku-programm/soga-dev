package webapi_test

import (
	_ "recovered.local/soga/panel/reg/ppaneltrojanwebapi"
	_ "recovered.local/soga/panel/reg/ppanelhysteriawebapi"
	P "recovered.local/soga/panel/webapi"
	C "recovered.local/soga/config"
	_ "recovered.local/soga/panel/reg/ppanelvlesswebapi"
	Q "recovered.local/soga/panel/registry"
	_ "recovered.local/soga/panel/reg/ppanelsswebapi"
	_ "recovered.local/soga/panel/reg/ppanelv2raywebapi"
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
