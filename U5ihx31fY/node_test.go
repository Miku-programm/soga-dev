package U5ihx31fY

import (
	"encoding/json"
	"math"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestRTTIAndCustomJSON(t *testing.T) {
	for _, c := range []struct{ got, want uintptr }{{unsafe.Sizeof(NNrxxXeBt{}), 344}, {unsafe.Sizeof(QhOI5ti{}), 200}, {unsafe.Sizeof(CEJsaaVtomQg{}), 288}, {unsafe.Sizeof(H0izB2agvG{}), 352}} {
		if c.got != c.want {
			t.Fatalf("layout %v != %v", c.got, c.want)
		}
	}
	n := NNrxxXeBt{HIgDb6THe1A9: "old", PzPwREkQp: CEJsaaVtomQg{ZeH9JLmUs: "500"}}
	if err := json.Unmarshal([]byte(`{"server":"new","custom_config":{"v2_port":"900","security":"tls"}}`), &n); err != nil {
		t.Fatal(err)
	}
	if n.HIgDb6THe1A9 != "new" || n.PzPwREkQp.ZeH9JLmUs != "500" || n.PzPwREkQp.UPnr_Bjaq != "900" {
		t.Fatalf("%+v", n)
	}
	if err := json.Unmarshal([]byte(`{"server":"next","custom_config":false}`), &n); err != nil || n.HIgDb6THe1A9 != "next" {
		t.Fatalf("ignored custom error: %v %+v", err, n)
	}
	if err := json.Unmarshal([]byte(`{"server":7}`), &n); err == nil || n.HIgDb6THe1A9 != "next" {
		t.Fatalf("outer mutation: %v %+v", err, n)
	}
	b, _ := json.Marshal(n)
	if strings.Contains(string(b), "custom_config") {
		t.Fatal(string(b))
	}
}
func TestServerParsersOriginalBranches(t *testing.T) {
	for _, tc := range []struct {
		s                  string
		port, alter        int
		net, path, service string
		tls                bool
		err                string
	}{
		{"", 80, 2, "tcp", "/", "", false, ""}, {"host;", 0, 2, "tcp", "/", "", false, ""}, {"h;83;", 83, 0, "tcp", "/", "", false, ""},
		{"h;83;2;ws", 83, 2, "tcp", "/", "", false, "parse server failed: h;83;2;ws"},
		{" h ;83;2;tls;grpc;path=/x=a|servicename=svc|inside_port=99 ", 99, 2, "grpc", "/x", "svc", true, ""},
		{"h;1;2;ws;plain;", 1, 2, "ws", "", "", false, ""},
		{"h;x", 0, 2, "tcp", "/", "", false, "parse server failed: h;x, reason: port invalid: x"},
		{"h;80;x", 80, 0, "tcp", "/", "", false, "parse server failed: h;80;x, reason: alterId invalid: x"},
	} {
		var n L.JW0CeHo3O
		err := mAndhwjjX(tc.s, &n)
		got := ""
		if err != nil {
			got = err.Error()
		}
		if got != tc.err || n.CKsOpayd != tc.port || n.VUwf1n9g != tc.alter || string(n.F2GXyZZ) != tc.net || n.ChJnngzb != tc.path || n.KWffzRDF != tc.service || n.UVeahqDBT != tc.tls {
			t.Errorf("%q: %+v err %v", tc.s, n, err)
		}
	}
	n := L.L9AiYC{PBDb3Wc: "preserved"}
	if err := bmG8Nfaaa(" h ;port=100#200#300", &n); err != nil || n.GbD_2PU != 200 || n.PBDb3Wc != "preserved" {
		t.Fatal(n, err)
	}
	n2 := L.Dv6QJyFW6bU{QawbKS2x: true}
	if err := caT62H0U8n("h;soga_port=600|port=500#700", &n2); err != nil || n2.PU1Di18 != 500 || n2.C33jbIzI != 600 || !n2.QawbKS2x {
		t.Fatal(n2, err)
	}
	if err := caT62H0U8n("h;soga_port=no|port=500", &n2); err == nil || !strings.Contains(err.Error(), "parse soga_port err:") || n2.PU1Di18 != 0 {
		t.Fatal(n2, err)
	}
}
func TestCustomPrecedenceAndSSRValidation(t *testing.T) {
	c := CEJsaaVtomQg{ZeH9JLmUs: "900", ITA1ZyGMvV2: "800", Z3x174Dd4P3: Z3x174Dd4P3{UPnr_Bjaq: "invalid", SFDGfYC: "-1"}}
	n := L.JW0CeHo3O{VUwf1n9g: 12, UVeahqDBT: true}
	c.UpdateVMessNode(&n)
	if n.CKsOpayd != 900 || n.VUwf1n9g != 12 || !n.UVeahqDBT {
		t.Fatal(n)
	}
	c.D2u1w_fniHh = D2u1w_fniHh{QUUtCN7gjjra: "unknown-preserved", IBgOeC: "key"}
	ss := L.AAMPrdcYTL6{}
	c.UpdateSSNode(&ss)
	if ss.HMkc4BgR != "unknown-preserved" || ss.VuhWqou != "key" {
		t.Fatal(ss)
	}
	c.HS63ppHsdS = HS63ppHsdS{RNq55eY: "secret", PyqeODiH: "aes-128-cfb", BaCpMmr7: "auth_chain_f", PDnKxsV8fttw: "tls1.2_ticket_auth"}
	node := L.Dv6QJyFW6bU{QawbKS2x: true}
	c.UpdateSSRNode(&node)
	if reflect.ValueOf(node).FieldByName("kJbZdr").Len() != 1 {
		t.Fatal("SSR missing credential")
	}
	c.PyqeODiH = "bad"
	c.UpdateSSRNode(&node)
	if reflect.ValueOf(node).FieldByName("kJbZdr").Len() != 1 {
		t.Fatal("invalid cipher accepted")
	}
}
func TestUserCredentialLimitsAndDetector(t *testing.T) {
	saved := stateAt3787337
	stateAt3787337 = 1
	defer func() { stateAt3787337 = saved }()
	u := QhOI5ti{BD16miI0: 7, ESm7aAEddE: "p", Tiybx6: "explicit", TnjwdH: math.NaN(), BAyDRja4o: "aes-128-gcm", TdPoGXzK: 11, RiD2Bnl: 2}
	x := u.GetUser("v2ray", 100)
	if x.JAfXWwzYrW.(*L.Yn8BJoL).SfiEGpJo != "explicit" {
		t.Fatal(x)
	}
	u.Tiybx6 = ""
	x = u.GetUser("v2ray", 0)
	if got := x.JAfXWwzYrW.(*L.Yn8BJoL).SfiEGpJo; got != "3df3326d-414a-3299-ae8f-9d8cefc630ed" {
		t.Fatalf("original UUIDv3 namespace DNS fixture: %s", got)
	}
	if !math.IsNaN(mergeSpeed(0, math.NaN())) || mergeSpeed(math.NaN(), 0) != 0 {
		t.Fatal("unordered merge")
	}
	all := &O8T6nD3uPzb{L8xOe2JrpGDq: make([]*QhOI5ti, 90)}
	for i := range all.L8xOe2JrpGDq {
		all.L8xOe2JrpGDq[i] = &QhOI5ti{Tiybx6: "x"}
	}
	stateAt3787337 = 1
	if got := len(all.GetUsers(C.Mpn2JltM("v2ray"))); got != 90 {
		t.Fatal(got)
	}
	d := (&E3s7VeiKCij{Uyj0tudIZN: 9, XEmS9xCt: "example\\.com", Bk26FTPOD: 999}).ToDetector()
	if d.GetId() != 9 || !d.Test("example.com") {
		t.Fatal("rule conversion")
	}
}
