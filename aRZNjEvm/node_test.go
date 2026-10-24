package aRZNjEvm

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http/httptest"
	"recovered.local/soga/CAodbhZf"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestOriginalLayoutsAndJSONFields(t *testing.T) {
	for _, tc := range []struct {
		v any
		n uintptr
	}{{GJOCQqa{}, 16}, {Xd0IG2Z{}, 16}, {LRlqOaYmO{}, 16}, {IMyQ5CxaikH{}, 40}, {FPZsKqad5Ox{}, 64}, {AOa7p4h{}, 40}, {DUnaqZqqbnWW{}, 296}, {BdT6k1HU8L{}, 616}} {
		if got := reflect.TypeOf(tc.v).Size(); got != tc.n {
			t.Fatalf("%T size %d want %d", tc.v, got, tc.n)
		}
	}
	var n BdT6k1HU8L
	if unsafe.Offsetof(n.edYRMY8) != 608 {
		t.Fatal("private state offset")
	}
	input := `{"server_port":9443,"tls":2,"network":"xhttp","networkSettings":{"host":"example.test","path":"/h","mode":"auto","extra":{"headers":{"X-One":"yes"},"xPaddingBytes":{"from":2,"to":4},"noSSEHeader":true,"scMaxEachPostBytes":{"from":5},"scMaxBufferedPosts":11,"scStreamUpServerSecs":{"to":9},"xPaddingObfsMode":true,"xPaddingKey":"pad","xPaddingHeader":"P","xPaddingPlacement":"header","xPaddingMethod":"repeat","uplinkHTTPMethod":"POST","sessionPlacement":"path","sessionKey":"session","seqPlacement":"query","seqKey":"seq","uplinkDataPlacement":"body","uplinkDataKey":"data","uplinkChunkSize":{"from":17},"serverMaxHeaderBytes":4096}},"flow":"xtls-rprx-vision","encryption":"mlkem768x25519plus","encryption_settings":{"mode":" native ","ticket":"","server_padding":" 20-40 ","private_key":"private"},"tls_settings":{"private_key":"reality-private","short_ids":["abc"],"server_names":["sni.test"],"dest":"remote.test","server_port":"443.9"},"base_config":{"pull_interval":"17.8","push_interval":23}}`
	if err := json.Unmarshal([]byte(input), &n); err != nil {
		t.Fatal(err)
	}
	out, err := n.ToPanelNode(19)
	if err != nil {
		t.Fatal(err)
	}
	if out.Ck6DxjB.WAdJ4ftdZFk != "reality" || out.JhIYwEukjd.AJjaqrBjj != "tls" || out.QaZjDzw95.J6bt8AOW8fL8 != "tls" {
		t.Fatal("legacy security", out)
	}
	if out.Ck6DxjB.DlDK5Y9AS != "mlkem768x25519plus.native.20-40.private" || out.Ck6DxjB.IqSfNzea != "remote.test:443" || out.Ck6DxjB.NMzMUQ2 != "xtls-rprx-vision" {
		t.Fatal(out.Ck6DxjB)
	}
	x := out.Ck6DxjB.FPhrzP645
	if x == nil || x.TzkkcVTusJD != "example.test" || x.HoBz4V != "/h" || x.MfWFUwd1 != "auto" || x.F0XSwqNYoaP["X-One"] != "yes" || x.N2C4XMihD.AS9zG362 != 2 || !x.A6NxppvA || x.LEaRo0 != 11 || x.ULnN_iDWLIXN.AS9zG362 != 5 || x.Iv950ouQ5AW.HHTPwX0N != 9 || !x.PnRF9ifkXKD || x.JvaHmm.AS9zG362 != 17 || x.Diwr8ey != 4096 {
		t.Fatalf("xhttp %+v", x)
	}
	if x.H6PfvFysST != "pad" || x.LjJbxC6vDL != "P" || x.V0qykwP != "header" || x.MLlGWkds != "repeat" || x.KJMkBJ != "POST" || x.XEKp9SbqKQ != "path" || x.G9oU9diBqr != "session" || x.HqtZJaRZNw != "query" || x.NPLvCM != "seq" || x.Aonc19Hp33E != "body" || x.HFEkUe4a9Sas != "data" {
		t.Fatal("string fields", x)
	}
	if n.Pa8W2AZ6Iw.F2WCOEc5Eu6.Value() != 17 || n.Pa8W2AZ6Iw.I8kigte.Value() != 23 {
		t.Fatal("interval tags")
	}
	n.PNpnXla.XK10gJ_O[0] = "shared"
	if out.Ck6DxjB.BANVIZ72wh[0] != "shared" || out.JhIYwEukjd.HIlblJq1[0] != "shared" {
		t.Fatal("original slice alias lost")
	}
}

func TestV2NodeNormalizationAndFailureOrder(t *testing.T) {
	for _, tc := range []struct {
		wire     string
		protocol C.Mpn2JltM
	}{{"vmess", "v2ray"}, {"vless", "vless"}, {"trojan", "trojan"}, {"anytls", "anytls"}, {"shadowsocks", "ss"}, {"hysteria2", "hysteria"}} {
		var n BdT6k1HU8L
		data := `{"protocol":"` + tc.wire + `","tls":1,"tls_settings":{"cert_mode":"none"},"network":"splithttp","network_settings":{"path":"/x","host":"host"},"obfs_password":"new","obfs-password":"old"}`
		if err := json.Unmarshal([]byte(data), &n); err != nil {
			t.Fatal(err)
		}
		if err := n.ocVy8aVcKm(tc.protocol); err != nil {
			t.Fatal(err)
		}
		if !n.edYRMY8 || n.FoQe96xgtPg != "xhttp" || n.Xg1j1ZA != 0 || n.F7emlj != "new" {
			t.Fatal(n)
		}
		if tc.protocol == "ss" && (n.YHnnp66mM != "http" || n.Enhxlh.UzanNits != "host") {
			t.Fatal(n)
		}
		if tc.protocol == "hysteria" && n.LQQHdt4QjrX != 2 {
			t.Fatal(n)
		}
	}
	n := BdT6k1HU8L{DWMzz2: "vmess"}
	if err := n.ocVy8aVcKm("ss"); err == nil || err.Error() != "v2node protocol vmess does not match server_type ss" || n.edYRMY8 {
		t.Fatal(n, err)
	}
	n.DWMzz2 = "VMESS"
	if err := n.ocVy8aVcKm("v2ray"); err == nil || err.Error() != "unsupported v2node protocol: VMESS" {
		t.Fatal(err)
	}
	n = BdT6k1HU8L{DWMzz2: "shadowsocks", LWnSaWm: []byte("broken"), YHnnp66mM: "tls"}
	if err := n.ocVy8aVcKm("ss"); err == nil || !strings.HasPrefix(err.Error(), "parse ss network settings failed:") || !n.edYRMY8 || n.YHnnp66mM != "" {
		t.Fatal(n, err)
	}
}

func TestCertificateBranchesAndErrorState(t *testing.T) {
	const id = 1971
	certificateErrorsAt509C8C0.Delete(id)
	t.Cleanup(func() { certificateErrorsAt509C8C0.Delete(id) })
	n := BdT6k1HU8L{edYRMY8: true, Xg1j1ZA: 1}
	n.PNpnXla.MtI3pIGhwl = "remote"
	if cert, key := n.wXQOQActA(id); cert != "" || key != "" {
		t.Fatal("empty cert")
	}
	if _, ok := certificateErrorsAt509C8C0.Load(id); !ok {
		t.Fatal("missing warning state")
	}
	n.PNpnXla.MtI3pIGhwl = "self"
	n.PNpnXla.VHVkiEr1NUE = "valid.example"
	n.PNpnXla.XK10gJ_O = []string{"bad/path"}
	if cert, key := n.wXQOQActA(id); cert != "" || key != "" {
		t.Fatal("invalid self domain")
	}
	if message, ok := certificateErrorsAt509C8C0.Load(id); !ok || !strings.Contains(message.(string), `invalid domain: "bad/path"`) {
		t.Fatal(message)
	}
	fixture := httptest.NewTLSServer(nil)
	defer fixture.Close()
	pair := fixture.TLS.Certificates[0]
	cert := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: pair.Certificate[0]}))
	keyDER, err := x509.MarshalPKCS8PrivateKey(pair.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	key := string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER}))
	if err := CAodbhZf.ZvmhWx(cert, key); err != nil {
		t.Fatal(err)
	}
	n.PNpnXla.MtI3pIGhwl = "remote"
	n.PNpnXla.Z6rZjvbHxvss = cert
	n.PNpnXla.UEzMLAOgTf = key
	gotC, gotK := n.wXQOQActA(id)
	if _, ok := certificateErrorsAt509C8C0.Load(id); ok {
		t.Fatal("success must clear state")
	}
	if gotC != cert || gotK != key {
		t.Fatal("remote changed")
	}
	n.PNpnXla.UEzMLAOgTf = "bad"
	n.wXQOQActA(id)
	if _, ok := certificateErrorsAt509C8C0.Load(id); !ok {
		t.Fatal("bad pair")
	}
	n.Xg1j1ZA = 2
	n.wXQOQActA(id)
	if _, ok := certificateErrorsAt509C8C0.Load(id); ok {
		t.Fatal("non-TLS must clear state")
	}
}

func TestSixUserCredentialsAndOriginalState(t *testing.T) {
	old := stateAt378733C
	t.Cleanup(func() { stateAt378733C = old })
	stateAt378733C = 1
	want := map[C.Mpn2JltM]any{"ss": &L.YlfOV0DBBnx{}, "v2ray": &L.Yn8BJoL{}, "trojan": &L.JxTGUar{}, "vless": &L.IImIwacp{}, "hysteria": &L.WQM2tKA{}, "anytls": &L.ZtMFlBTZXyD{}}
	for protocol, kind := range want {
		stateAt378733C = 1
		r := IMyQ5CxaikH{Ta593Rq2lSa: OJv5MKs{{IDLJXX6F: 2, UH0HXa5O: "credential", DwdlKBQpcgY: 20001, QeuUtZ8f: -2}}}
		out := r.ToPanelUsers(protocol)
		if len(out) != 1 || reflect.TypeOf(out[0].JAfXWwzYrW) != reflect.TypeOf(kind) || out[0].LmOKwjha != 10000 || stateAt378733C != 0 {
			t.Fatal(protocol, out)
		}
	}
	stateAt378733C = 1
	if got := (&AOa7p4h{}).dwkvm4ffDDTc("unknown", 0).JAfXWwzYrW; got != nil {
		t.Fatal(got)
	}
}

func TestNodeMalformedTransportAndDecryption(t *testing.T) {
	for _, network := range []L.BoulTn31{"ws", "grpc", "xhttp"} {
		n := BdT6k1HU8L{FoQe96xgtPg: network, PabrtEwnkJ: []byte("malformed")}
		if out, err := n.ToPanelNode(1); out != nil || err == nil || !strings.HasPrefix(err.Error(), "parse "+string(network)+" settings failed:") {
			t.Fatal(network, out, err)
		}
	}
	n := BdT6k1HU8L{HDMF3C: "unknown"}
	if out, err := n.ToPanelNode(1); out != nil || err == nil || err.Error() != "unknown vless decryption: unknown, err: unsupported vless decryption: unknown" {
		t.Fatal(out, err)
	}
	n = BdT6k1HU8L{}
	out, err := n.ToPanelNode(1)
	if err != nil || n.FoQe96xgtPg != "" || out.SwxJNgPpg.F2GXyZZ != "" || out.Ck6DxjB.IqSfNzea != ":0" || !reflect.DeepEqual(out.Ck6DxjB.BANVIZ72wh, []string{""}) || !reflect.DeepEqual(out.Ck6DxjB.DxKJ7lSe3m8F, []string{""}) {
		t.Fatal(out, err)
	}
	n.edYRMY8 = true
	out, err = n.ToPanelNode(1)
	if err != nil || out.JhIYwEukjd.AJjaqrBjj != "none" || out.QaZjDzw95.J6bt8AOW8fL8 != "none" {
		t.Fatal(out, err)
	}
}
