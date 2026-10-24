package xboard

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http/httptest"
	L "recovered.local/soga/panel/types"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestOriginalLayoutAndDetailedNodeMapping(t *testing.T) {
	for _, tc := range []struct {
		v any
		n uintptr
	}{{OFybGF5iIp{}, 16}, {YMnnZhxYFt{}, 16}, {KlO65BuwOt{}, 16}, {Sf0Wjail3Z{}, 40}, {CYKo3knGCrtI{}, 64}, {PForJoFExa{}, 40}, {ISasilPp{}, 56}, {SNlDfbf3j{}, 296}, {OGz0Bb0{}, 512}} {
		if got := reflect.TypeOf(tc.v).Size(); got != tc.n {
			t.Fatalf("%T size=%d", tc.v, got)
		}
	}
	var n OGz0Bb0
	if unsafe.Offsetof(n.DiDLlInAz) != 464 {
		t.Fatal("certificate layout")
	}
	input := `{"server_port":9443,"tls":2,"network":"xhttp","networkSettings":{"host":"example.test","path":"/h","mode":"auto","extra":{"headers":{"X-One":"yes"},"xPaddingBytes":{"from":2,"to":4},"noSSEHeader":true,"scMaxEachPostBytes":{"from":5},"scMaxBufferedPosts":11,"scStreamUpServerSecs":{"to":9},"xPaddingObfsMode":true,"xPaddingKey":"pad","xPaddingHeader":"P","xPaddingPlacement":"header","xPaddingMethod":"repeat","uplinkHTTPMethod":"POST","sessionPlacement":"path","sessionKey":"session","seqPlacement":"query","seqKey":"seq","uplinkDataPlacement":"body","uplinkDataKey":"data","uplinkChunkSize":{"from":17},"serverMaxHeaderBytes":4096}},"flow":"xtls-rprx-vision","decryption":"raw-opaque-value","tls_settings":{"private_key":"reality-private","short_id":"abc","server_name":"sni.test","server_port":"443.9","ech":{"enabled":"TRUE","key":"AAAAAA=="}},"base_config":{"pull_interval":"17.8","push_interval":23},"padding_scheme":["one","two"],"transport":"udp","traffic_pattern":"profile","plugin":"obfs","plugin_opts":"obfs=http;path=/ss;obfs-host=ss.test","obfs_settings":{"path":"ignored","host":"ignored"}}`
	if err := json.Unmarshal([]byte(input), &n); err != nil {
		t.Fatal(err)
	}
	out, err := n.ToPanelNode(19)
	if err != nil {
		t.Fatal(err)
	}
	if out.Ck6DxjB.WAdJ4ftdZFk != "reality" || out.JhIYwEukjd.AJjaqrBjj != "reality" || out.QaZjDzw95.J6bt8AOW8fL8 != "tls" {
		t.Fatal(out)
	}
	if out.LKMSis7zkX != "AAAAAA==" || out.Ck6DxjB.DlDK5Y9AS != "raw-opaque-value" || out.Ck6DxjB.IqSfNzea != "sni.test:443" || out.Ck6DxjB.NMzMUQ2 != "xtls-rprx-vision" || out.QaZjDzw95.UPnToXrM != "one\ntwo" || out.ABA4hq.GGSLzD != "UDP" || out.ABA4hq.CLWr9bQsb != "profile" {
		t.Fatal(out)
	}
	if out.FnMUny.ErKysFf2I != "http" || out.FnMUny.PzwWMqEpPguJ != "/ss" || out.FnMUny.DURh5VYO != "ss.test" {
		t.Fatal(out.FnMUny)
	}
	if len(out.QaZjDzw95.Jq7YxRdb18Xx) != 0 || out.QaZjDzw95.CTa7WN != "" || out.QaZjDzw95.F9P9aeOu23U != "" {
		t.Fatal("unexpected anytls reality fields")
	}
	out.Ck6DxjB.BANVIZ72wh[0] = "changed"
	out.Ck6DxjB.DxKJ7lSe3m8F[0] = "changed"
	if out.JhIYwEukjd.HIlblJq1[0] != "sni.test" || out.JhIYwEukjd.TJ8rFSEH[0] != "abc" {
		t.Fatal("slices must be independent")
	}
	x := out.Ck6DxjB.FPhrzP645
	if x == nil || x.TzkkcVTusJD != "example.test" || x.HoBz4V != "/h" || x.MfWFUwd1 != "auto" || x.F0XSwqNYoaP["X-One"] != "yes" || x.N2C4XMihD.AS9zG362 != 2 || !x.A6NxppvA || x.LEaRo0 != 11 || x.ULnN_iDWLIXN.AS9zG362 != 5 || x.Iv950ouQ5AW.HHTPwX0N != 9 || !x.PnRF9ifkXKD || x.JvaHmm.AS9zG362 != 17 || x.Diwr8ey != 4096 {
		t.Fatalf("%+v", x)
	}
	if x.H6PfvFysST != "pad" || x.LjJbxC6vDL != "P" || x.V0qykwP != "header" || x.MLlGWkds != "repeat" || x.KJMkBJ != "POST" || x.XEKp9SbqKQ != "path" || x.G9oU9diBqr != "session" || x.HqtZJaRZNw != "query" || x.NPLvCM != "seq" || x.Aonc19Hp33E != "body" || x.HFEkUe4a9Sas != "data" {
		t.Fatal(x)
	}
	if n.PnVsaJAWK.F2WCOEc5Eu6.Value() != 17 || n.PnVsaJAWK.I8kigte.Value() != 23 {
		t.Fatal("interval tags")
	}
}

func TestOptionsAndConversionFailures(t *testing.T) {
	n := OGz0Bb0{XXzsbzH3U7: "obfs=http; ignored ; path= /a=b ; obfs=tls;empty=; =ignored"}
	if got := n.oN0fey(); !reflect.DeepEqual(got, map[string]string{"obfs": "tls", "path": "/a=b"}) {
		t.Fatal(got)
	}
	for _, network := range []L.BoulTn31{"ws", "grpc", "xhttp"} {
		n := OGz0Bb0{T7xGio: network, MiIwSabwwMec: []byte("malformed")}
		if out, err := n.ToPanelNode(1); out != nil || err == nil || !strings.HasPrefix(err.Error(), "parse "+string(network)+" settings failed:") {
			t.Fatal(network, out, err)
		}
	}
	n = OGz0Bb0{}
	out, err := n.ToPanelNode(1)
	if err != nil || out.SwxJNgPpg.F2GXyZZ != "" || out.JhIYwEukjd.AJjaqrBjj != "none" || out.QaZjDzw95.J6bt8AOW8fL8 != "tls" || out.Ck6DxjB.IqSfNzea != ":0" || out.AYf5kQSQW.B30Jt2kO != "plain" {
		t.Fatal(out, err)
	}
	for _, tc := range []struct{ obfs, password, want string }{{"salamander", "secret", "secret"}, {"legacy", "ignored", "legacy"}} {
		n.Zd6n_x4ye, n.FBsOqCUmDeVc = tc.obfs, tc.password
		out, err = n.ToPanelNode(1)
		if err != nil || out.AYf5kQSQW.B30Jt2kO != "salamander" || out.AYf5kQSQW.UjgtNuouuKVl != tc.want {
			t.Fatal(out, err)
		}
	}
	n.KVZ5bT = "unsupported"
	if out, err = n.ToPanelNode(1); out != nil || err == nil || err.Error() != "unsupported ss plugin: unsupported" {
		t.Fatal(out, err)
	}
	n.Biccak.IMskp1_c.Orib8hhz.UnmarshalJSON([]byte("true"))
	if out, err = n.ToPanelNode(1); out != nil || err == nil || err.Error() != "ech enabled but key is empty" {
		t.Fatal("ECH must precede plugin", out, err)
	}
	n.Biccak.IMskp1_c.HQ5AFB2w4c = "!!"
	if out, err = n.ToPanelNode(1); out != nil || err == nil || !strings.HasPrefix(err.Error(), "ech key 格式错误:") {
		t.Fatal(out, err)
	}
	n.Biccak.IMskp1_c.Orib8hhz.UnmarshalJSON([]byte("false"))
	n.KVZ5bT = ""
	if out, err = n.ToPanelNode(1); err != nil || out.LKMSis7zkX != "" {
		t.Fatal(out, err)
	}
}

func TestContentCertificateNoTLSGateAndWarningState(t *testing.T) {
	const id = 811
	certificateErrorsAt509C880.Delete(id)
	t.Cleanup(func() { certificateErrorsAt509C880.Delete(id) })
	n := OGz0Bb0{}
	n.DiDLlInAz.H_HT1L = "content"
	if cert, key := n.yU6F1oN6a(id); cert != "" || key != "" {
		t.Fatal(cert, key)
	}
	if message, ok := certificateErrorsAt509C880.Load(id); !ok || message != "cert_mode is content but cert_content or key_content is empty, ignore panel cert" {
		t.Fatal(message)
	}
	fixture := httptest.NewTLSServer(nil)
	defer fixture.Close()
	pair := fixture.TLS.Certificates[0]
	cert := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: pair.Certificate[0]}))
	der, err := x509.MarshalPKCS8PrivateKey(pair.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	key := string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}))
	n.DiDLlInAz.HrieyL3JKL5, n.DiDLlInAz.FjIridNPs19h = cert, key
	out, err := n.ToPanelNode(id)
	if err != nil || out.I3k4JnNp7nB != cert || out.XATiVun != key {
		t.Fatal(out, err)
	}
	if _, ok := certificateErrorsAt509C880.Load(id); ok {
		t.Fatal("success must clear warning")
	}
	n.DiDLlInAz.FjIridNPs19h = "bad"
	if cert, key := n.yU6F1oN6a(id); cert != "" || key != "" {
		t.Fatal(cert, key)
	}
	if message, ok := certificateErrorsAt509C880.Load(id); !ok || !strings.HasPrefix(message.(string), "panel cert content invalid, ignore panel cert:") {
		t.Fatal(message)
	}
	n.DiDLlInAz.H_HT1L = "other"
	n.yU6F1oN6a(id)
	if _, ok := certificateErrorsAt509C880.Load(id); ok {
		t.Fatal("other mode must clear warning")
	}
}
