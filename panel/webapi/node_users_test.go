package webapi

import (
	"encoding/json/jsontext"
	"math"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestNodeTransportsAndErrorPrecedence(t *testing.T) {
	for _, protocol := range []string{"vmess", "vless", "trojan"} {
		for _, transport := range []string{"tcp", "ws", "websocket", "h2", "grpc", "xhttp", ""} {
			config := &Yq4D8wXbIQyY{B0XVtyvn0k31: protocol, OWRl87iaah: jsontext.Value(`{"port":8443,"transport":"` + transport + `","security":"tls","transport_config":{"path":"/path","service_name":"rpc"}}`)}
			node, e := config.ToPanelNode(42)
			valid := transport == "tcp" || transport == "ws" || transport == "websocket" || transport == "grpc" || transport == "h2" && protocol != "trojan"
			if !valid {
				if e == nil || e.Error() != "unknown transport type: "+transport {
					t.Fatalf("%s/%s: %v", protocol, transport, e)
				}
				continue
			}
			if e != nil {
				t.Fatal(e)
			}
			mapped := transport
			if mapped == "websocket" {
				mapped = "ws"
			}
			switch protocol {
			case "vmess":
				if node.SwxJNgPpg != (L.JW0CeHo3O{CKsOpayd: 8443, F2GXyZZ: L.BoulTn31(mapped), UVeahqDBT: true, ChJnngzb: "/path", KWffzRDF: "rpc"}) {
					t.Fatal(node.SwxJNgPpg)
				}
			case "vless":
				if node.Ck6DxjB.GJtOdKOWBiFP != 8443 || node.Ck6DxjB.XmIaPj3CzOF != L.BoulTn31(mapped) || !reflect.DeepEqual(node.Ck6DxjB.BANVIZ72wh, []string{""}) || node.Ck6DxjB.IqSfNzea != ":0" {
					t.Fatal(node.Ck6DxjB)
				}
			case "trojan":
				if node.JhIYwEukjd.GbD_2PU != 8443 || node.JhIYwEukjd.T4cuhG7MTbG != L.BoulTn31(mapped) || node.JhIYwEukjd.AJjaqrBjj != "" || node.JhIYwEukjd.PBDb3Wc != "/path" {
					t.Fatal(node.JhIYwEukjd)
				}
			}
		}
	}
	for _, tc := range []struct{ protocol, config, want string }{{"trojan", `{"transport":"h2","security":"none"}`, "unknown transport type: h2"}, {"trojan", `{"transport":"tcp","security":"none"}`, "unknown security type: none"}, {"hysteria", `{}`, "unknown node protocol: hysteria"}, {"unknown", `broken`, "unknown node protocol: unknown"}} {
		_, e := (&Yq4D8wXbIQyY{B0XVtyvn0k31: tc.protocol, OWRl87iaah: jsontext.Value(tc.config)}).ToPanelNode(1)
		if e == nil || e.Error() != tc.want {
			t.Fatalf("%+v: %v", tc, e)
		}
	}
	for _, password := range []string{"", "mask"} {
		n, e := (&Yq4D8wXbIQyY{B0XVtyvn0k31: "hysteria2", OWRl87iaah: jsontext.Value(`{"port":443,"obfs_password":"` + password + `"}`)}).ToPanelNode(2)
		if e != nil {
			t.Fatal(e)
		}
		want := "plain"
		if password != "" {
			want = "salamander"
		}
		if n.AYf5kQSQW.Li6nvK1 != 2 || n.AYf5kQSQW.PIvpxN6YK != 443 || string(n.AYf5kQSQW.B30Jt2kO) != want {
			t.Fatal(n.AYf5kQSQW)
		}
	}
}
func TestVLESSRealityConfigPreservesOriginalFormatting(t *testing.T) {
	n, e := (&Yq4D8wXbIQyY{B0XVtyvn0k31: "vless", OWRl87iaah: jsontext.Value(`{"port":443,"flow":"xtls-rprx-vision","security":"reality","security_config":{"reality_private_key":"private","reality_server_addr":"::1","reality_server_port":443,"reality_short_id":"abcd","sni":"example.test"},"transport":"tcp"}`)}).ToPanelNode(8)
	if e != nil {
		t.Fatal(e)
	}
	v := n.Ck6DxjB
	if v.WAdJ4ftdZFk != "reality" || v.NMzMUQ2 != "xtls-rprx-vision" || v.IqSfNzea != "::1:443" || v.CrS6C1 != "private" || !reflect.DeepEqual(v.DxKJ7lSe3m8F, []string{"abcd"}) || !reflect.DeepEqual(v.BANVIZ72wh, []string{"example.test"}) {
		t.Fatal(v)
	}
}
func TestUserInitialBypassAndOriginalSliceLimit(t *testing.T) {
	old := stateAt3787334
	t.Cleanup(func() { stateAt3787334 = old })
	stateAt3787334 = 1
	users := make(T8lmI8Ic, 90)
	for i := range users {
		users[i] = &MACjYTsGk7G{AnbiaLml0Bs: int64(i), WtNHeT4: "credential"}
	}
	first := users.ToPanelUsers("v2ray")
	if len(first) != 90 || first[89].JAfXWwzYrW == nil || stateAt3787334 != 0 {
		t.Fatal("first pass")
	}
	second := users.ToPanelUsers("v2ray")
	if len(second) != 90 || len(users) != 90 {
		t.Fatal("user list must not be truncated")
	}
	if empty := T8lmI8Ic(nil).ToPanelUsers("unknown"); empty == nil {
		t.Fatal("empty output nil")
	}
	for _, tc := range []struct{ in, want float64 }{{-1, 0}, {10001, 10000}, {math.Inf(1), 10000}, {math.Inf(-1), 0}, {math.NaN(), 0}} {
		u := (&MACjYTsGk7G{YQbaJusq0g6N: tc.in, D7H3u6wL2q5Y: -1}).ToPanelUser("unknown")
		if u.LmOKwjha != tc.want || u.Xmrna8 != 0 || u.JAfXWwzYrW != nil {
			t.Fatalf("%v: %+v", tc, u)
		}
	}
	for _, p := range []C.Mpn2JltM{"ss", "v2ray", "trojan", "vless", "hysteria"} {
		u := (&MACjYTsGk7G{AnbiaLml0Bs: 73, WtNHeT4: "key", YQbaJusq0g6N: 18, D7H3u6wL2q5Y: 2}).ToPanelUser(p)
		if u.WZ89pTAOrHK4 != 73 || u.LmOKwjha != 18 || u.Xmrna8 != 2 || u.JAfXWwzYrW == nil {
			t.Fatal(u)
		}
	}
}
func TestRTTIAndIntervals(t *testing.T) {
	for _, v := range []struct {
		value any
		want  uintptr
	}{{Zkt5w_2o{}, 16}, {Yq4D8wXbIQyY{}, 56}, {MACjYTsGk7G{}, 40}, {LuuJNqNq1YTl{}, 160}, {Fp4V6S84pbeq{}, 72}, {PtTswARD{}, 72}, {XnaEdZLZi_Lz{}, 24}, {JV07s34JhN6z{}, 40}, {LC4dqy{}, 24}, {DjMeiycBVU{}, 24}} {
		if n := reflect.TypeOf(v.value).Size(); n != v.want {
			t.Fatalf("%T: %d", v.value, n)
		}
	}
	if unsafe.Offsetof(Yq4D8wXbIQyY{}.OWRl87iaah) != 32 {
		t.Fatal("raw config offset")
	}
	n := new(Yq4D8wXbIQyY)
	n.OWYIa8fDu.F2WCOEc5Eu6 = -2
	n.OWYIa8fDu.I8kigte = 0
	if n.GetPullInterval() != -2*time.Second || n.GetPushInterval() != 0 {
		t.Fatal("original unvalidated intervals")
	}
}
