package CbO_7wp3B

import (
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"testing"
)

// 固定表来自原 ELF 的 52 个注册调用，验证主入口真正导入全部注册包。
func TestAllOriginalPanelRegistrations(t *testing.T) {
	expected := []struct{ key, panel, protocol, api string }{
		{"dj-v2ray-webapi", "dj", "v2ray", "webapi"},                         // _U2caaPPsm.init.0 0x169efa0
		{"ppanel-hysteria-webapi", "ppanel", "hysteria", "webapi"},           // HULJhSRw0LH.init.0 0x16b8400
		{"ppanel-ss-webapi", "ppanel", "ss", "webapi"},                       // tNCtf5NWXas2.init.0 0x16b84c0
		{"ppanel-trojan-webapi", "ppanel", "trojan", "webapi"},               // CsBw7EWKCp.init.0 0x16b8580
		{"ppanel-v2ray-webapi", "ppanel", "v2ray", "webapi"},                 // tWaD46z.init.0 0x16b8700
		{"ppanel-vless-webapi", "ppanel", "vless", "webapi"},                 // mQFF7Rd7.init.0 0x16b8640
		{"proxypanel-ssr-webapi", "proxypanel", "ssr", "webapi"},             // OieaM9.init.0 0x16c2520
		{"proxypanel-trojan-webapi", "proxypanel", "trojan", "webapi"},       // sTqn2l9.init.0 0x16c92a0
		{"proxypanel-v2ray-webapi", "proxypanel", "v2ray", "webapi"},         // DMAoms9O.init.0 0x16cf400
		{"soga-v1-anytls-webapi", "soga-v1", "anytls", "webapi"},             // ss9YuZy.init.0 0x16e7760
		{"soga-v1-hysteria-webapi", "soga-v1", "hysteria", "webapi"},         // W4PUtj.init.0 0x16e7820
		{"soga-v1-mieru-webapi", "soga-v1", "mieru", "webapi"},               // SaoPT0jEL.init.0 0x16e78e0
		{"soga-v1-ss-webapi", "soga-v1", "ss", "webapi"},                     // CtXV8P.init.0 0x16e79a0
		{"soga-v1-ssr-webapi", "soga-v1", "ssr", "webapi"},                   // Fvby6xg.init.0 0x16e7a60
		{"soga-v1-trojan-webapi", "soga-v1", "trojan", "webapi"},             // F5pRgr3hi.init.0 0x16e7b20
		{"soga-v1-v2ray-webapi", "soga-v1", "v2ray", "webapi"},               // Iu4AOrdRh.init.0 0x16e7ca0
		{"soga-v1-vless-webapi", "soga-v1", "vless", "webapi"},               // F9kjkaHzdn.init.0 0x16e7be0
		{"sspanel-uim-ss-db", "sspanel-uim", "ss", "db"},                     // GgDwT38SMQf5.init.0 0x1737480
		{"sspanel-uim-ss-webapi", "sspanel-uim", "ss", "webapi"},             // GgDwT38SMQf5.init.1 0x17397a0
		{"sspanel-uim-ssr-db", "sspanel-uim", "ssr", "db"},                   // mIfcoDkc6.init.0 0x173b600
		{"sspanel-uim-ssr-webapi", "sspanel-uim", "ssr", "webapi"},           // mIfcoDkc6.init.1 0x173ccc0
		{"sspanel-uim-trojan-db", "sspanel-uim", "trojan", "db"},             // Se1QrdYD7Aa.init.0 0x173fe20
		{"sspanel-uim-trojan-webapi", "sspanel-uim", "trojan", "webapi"},     // Se1QrdYD7Aa.init.1 0x1740cc0
		{"sspanel-uim-v2ray-db", "sspanel-uim", "v2ray", "db"},               // uavUP6a.init.0 0x17418a0
		{"sspanel-uim-v2ray-webapi", "sspanel-uim", "v2ray", "webapi"},       // uavUP6a.init.1 0x1742780
		{"v2board-ss-webapi", "v2board", "ss", "webapi"},                     // t3wvrixO2e7V.init.0 0x174e240
		{"v2board-trojan-webapi", "v2board", "trojan", "webapi"},             // _lslMIG7j.init.0 0x174e300
		{"v2board-v2ray-webapi", "v2board", "v2ray", "webapi"},               // MDpSu9EFoD.init.0 0x174e3c0
		{"v2raysocks-ss-webapi", "v2raysocks", "ss", "webapi"},               // HakjBGjG.init.0 0x175b800
		{"v2raysocks-trojan-webapi", "v2raysocks", "trojan", "webapi"},       // ygzRkd5EN.init.0 0x175b8c0
		{"v2raysocks-v2ray-webapi", "v2raysocks", "v2ray", "webapi"},         // HhkySK.init.0 0x175ba40
		{"v2raysocks-vless-webapi", "v2raysocks", "vless", "webapi"},         // mOR0KqC.init.0 0x175b980
		{"whmcs-ss-db", "whmcs", "ss", "db"},                                 // LNVPAaJc.init.0 0x1769be0
		{"whmcs-ss-webapi", "whmcs", "ss", "webapi"},                         // LNVPAaJc.init.1 0x1771420
		{"whmcs-ssr-db", "whmcs", "ssr", "db"},                               // ckaMRyvJxQR.init.0 0x17716e0
		{"whmcs-trojan-db", "whmcs", "trojan", "db"},                         // JlGfHYMcZ.init.0 0x177a800
		{"whmcs-trojan-webapi", "whmcs", "trojan", "webapi"},                 // JlGfHYMcZ.init.1 0x1781480
		{"whmcs-v2ray-db", "whmcs", "v2ray", "db"},                           // R95yk_mwAr.init.0 0x1781600
		{"whmcs-v2ray-webapi", "whmcs", "v2ray", "webapi"},                   // R95yk_mwAr.init.1 0x1788440
		{"xboard-anytls-webapi", "xboard", "anytls", "webapi"},               // z0244yEhS6W.init.0 0x179b3e0
		{"xboard-hysteria-webapi", "xboard", "hysteria", "webapi"},           // k9EI4fqL9.init.0 0x179b4a0
		{"xboard-mieru-webapi", "xboard", "mieru", "webapi"},                 // HbH5rTkHd7zj.init.0 0x179b560
		{"xboard-ss-webapi", "xboard", "ss", "webapi"},                       // rusgvYjAAU5.init.0 0x179b620
		{"xboard-trojan-webapi", "xboard", "trojan", "webapi"},               // iiKUPoplen.init.0 0x179b6e0
		{"xboard-v2ray-webapi", "xboard", "v2ray", "webapi"},                 // W_EMft.init.0 0x179b860
		{"xboard-vless-webapi", "xboard", "vless", "webapi"},                 // aJU33bUzb.init.0 0x179b7a0
		{"xiaov2board-anytls-webapi", "xiaov2board", "anytls", "webapi"},     // qx2LmdvLIV.init.0 0x17b0040
		{"xiaov2board-hysteria-webapi", "xiaov2board", "hysteria", "webapi"}, // V9HXHRH7Fig.init.0 0x17b0100
		{"xiaov2board-ss-webapi", "xiaov2board", "ss", "webapi"},             // pcrbn5.init.0 0x17b01c0
		{"xiaov2board-trojan-webapi", "xiaov2board", "trojan", "webapi"},     // bjHvRkRuyWRt.init.0 0x17b0280
		{"xiaov2board-v2ray-webapi", "xiaov2board", "v2ray", "webapi"},       // MkeN6dA737wC.init.0 0x17b0400
		{"xiaov2board-vless-webapi", "xiaov2board", "vless", "webapi"},       // FkyU22TRjL_B.init.0 0x17b0340
	}
	if len(Q.PanelFactoriesAt5095E50) != len(expected) {
		t.Errorf("main registered %d combinations; original has %d", len(Q.PanelFactoriesAt5095E50), len(expected))
	}
	for _, tc := range expected {
		t.Run(tc.key, func(t *testing.T) {
			if Q.PanelFactoriesAt5095E50[tc.key] == nil {
				t.Fatal("original panel registration missing from actual main imports")
			}
			cfg := &C.FCEAfQJwDA7e{OS91X026uML: C.FSOZSH7g(tc.panel), ZyAhkq5th8bf: C.Mpn2JltM(tc.protocol), F3PCiaIra: C.HWLVDK2(tc.api), Yf1KqSqJpNI7: 17}
			panel, err := Q.SMkpVNb(cfg)
			if err != nil || panel == nil {
				t.Fatalf("actual constructor: %T %v", panel, err)
			}
			if panel.GetAPIConfig() != cfg {
				t.Fatal("constructor did not preserve original configuration pointer")
			}
			second, err := Q.SMkpVNb(cfg)
			if err != nil || second == panel {
				t.Fatalf("independent allocation: %T %v", second, err)
			}
		})
	}
}
