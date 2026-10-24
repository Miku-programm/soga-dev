package _ZwYXGIdVc

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func fixtureConfig(t *testing.T) {
	t.Helper()
	p := filepath.Join(t.TempDir(), "soga.conf")
	if e := os.WriteFile(p, []byte("type=whmcs\nserver_type=ss\nsoga_key=\nuser_speed_limit=12\nuser_conn_limit=4\nuser_tcp_limit=7\nnode_speed_limit=25\nforbidden_bit_torrent=false\n"), 0600); e != nil {
		t.Fatal(e)
	}
	C.MkwNWciPHS(p)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
}

const nodeJSON = `{"inbounds":[{"port":8443,"protocol":"shadowsocks","settings":{"method":"aes-128-gcm","password":"node-secret"},"streamingSettings":{"security":"ignored-for-ss"}},{"port":1,"protocol":"invalid"}],"routing":{"rules":[{"outboundTag":"block","domain":["domain:blocked.invalid",""],"protocol":["bittorrent"]},{"outboundTag":"direct","domain":["ignored"]}]}}`

func TestOriginalNodeValidationAndRTTI(t *testing.T) {
	fixtureConfig(t)
	for _, v := range []struct {
		x any
		n uintptr
	}{{DrEoaRCt2{}, 8}, {PNf9Em87{}, 120}, {AnlFZbNq{}, 48}, {AenyiikKLY67{}, 72}, {Oh1oNuTp{}, 40}, {Vg6Nm38GB9{}, 72}} {
		if reflect.TypeOf(v.x).Size() != v.n {
			t.Fatalf("RTTI %T", v.x)
		}
	}
	var r AnlFZbNq
	if e := json.Unmarshal([]byte(nodeJSON), &r); e != nil {
		t.Fatal(e)
	}
	n, e := r.ToPanelNode()
	if e != nil {
		t.Fatal(e)
	}
	if n.EMlXEc != 0 || n.FnMUny.Am3QoeF != 8443 || n.FnMUny.HMkc4BgR != "aes-128-gcm" || n.FnMUny.VuhWqou != "node-secret" || !n.HSA2b5V || n.TtWKyh != 25 || n.W3Dw_B4wgbM != 4 {
		t.Fatalf("node %+v", n)
	}
	if r.Dps5ZWL[0].Jedod2gPZ4UT.WZVYFl != "tcp" || n.Z1TIgHOQ != Q.T0j_Rco(r.Dps5ZWL[0]) {
		t.Fatal("default mutation/hash scope")
	}
	d := r.GetDetectors()
	if len(d) != 2 {
		t.Fatal(d)
	}
	if _, ok := d.Test("www.blocked.invalid"); !ok {
		t.Fatal("domain rule")
	}
	if _, ok := d.Test("other.invalid"); ok {
		t.Fatal("empty panel rule must not become DNS match-all")
	}
	if got := (&AnlFZbNq{}).GetDetectors(); got == nil {
		t.Fatal("nil empty detectors")
	}
	if _, e = (&AnlFZbNq{}).ToPanelNode(); e == nil || e.Error() != "inbounds is empty" {
		t.Fatal(e)
	}
	for _, protocol := range []string{"vmess", "vless", "trojan", "shadowsocks"} {
		for _, network := range []string{"", "tcp", "ws", "h2", "grpc"} {
			in := &PNf9Em87{QxNm4UvLN: protocol}
			in.Jedod2gPZ4UT.WZVYFl = network
			if e := in.mCZQ7N3(); e != nil {
				t.Fatal(e)
			}
		}
	}
	for _, v := range []struct{ p, s, n, w string }{{"invalid", "", "", "unknown protocol: invalid"}, {"vmess", "reality", "", "unknown security: reality"}, {"vmess", "tls", "quic", "unknown network: quic"}} {
		in := &PNf9Em87{QxNm4UvLN: v.p}
		in.Jedod2gPZ4UT.Id1HxAg5Oy = v.s
		in.Jedod2gPZ4UT.WZVYFl = v.n
		if e := in.mCZQ7N3(); e == nil || e.Error() != v.w {
			t.Fatal(e)
		}
	}
}
func TestOriginalUsersAndNaN(t *testing.T) {
	fixtureConfig(t)
	u := &AenyiikKLY67{PqTXyjlHNDD_: 42, SDB5DM3E68Y: "uuid", HZO6BjGa: "ss", NRLcM3Xm2zgj: "trojan", C2TUmBU9gT: 9, WEkiWDf5si: 8}
	r := &Oh1oNuTp{DCbciSyJ_c: "bad", XH32wVY: []*AenyiikKLY67{u}}
	if got, e := r.ToPanelUsers("ss"); got != nil || e == nil || e.Error() != "bad" {
		t.Fatal(got, e)
	}
	r.DCbciSyJ_c = "ok"
	for kind, want := range map[C.Mpn2JltM]any{"ss": &L.YlfOV0DBBnx{A_wNUTkMh6pH: "ss"}, "v2ray": &L.Yn8BJoL{SfiEGpJo: "uuid"}, "trojan": &L.JxTGUar{ARARP1z0yA: "trojan"}, "vless": nil} {
		a, e := r.ToPanelUsers(kind)
		if e != nil || !reflect.DeepEqual(a[0].JAfXWwzYrW, want) || a[0].LmOKwjha != 8 || a[0].Xmrna8 != 4 || a[0].B4DYSxuA31za != 7 {
			t.Fatal(kind, a, e)
		}
	}
	for _, s := range []float64{-1, 0, math.NaN(), 20000} {
		u.WEkiWDf5si = s
		a, _ := r.ToPanelUsers("ss")
		if a[0].LmOKwjha != 12 {
			t.Fatalf("original operand order %v -> %v", s, a[0].LmOKwjha)
		}
	}
	C.T15Jifo(true)
	defer C.T15Jifo(false)
	r.XH32wVY = make([]*AenyiikKLY67, 90)
	for i := range r.XH32wVY {
		r.XH32wVY[i] = u
	}
	a, e := r.ToPanelUsers("ss")
	if e != nil || len(a) != 90 || a[89].JAfXWwzYrW == nil {
		t.Fatal("this original adapter has no user-count license branch", e)
	}
}
func TestWebAPIWireAndConcurrentSubmit(t *testing.T) {
	fixtureConfig(t)
	var mu sync.Mutex
	requests := map[string]int{}
	bodies := map[string][]byte{}
	badUsers := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/base/web_api/" || r.URL.Query().Get("token") != "key" || r.URL.Query().Get("node_id") != "71" || r.URL.Query().Get("node_type") != "shadowsocks" {
			t.Errorf("request %s", r.URL)
		}
		if r.Header.Get("If-None-Match") != "" {
			t.Error("unexpected cache header")
		}
		action := r.URL.Query().Get("act")
		mu.Lock()
		requests[action]++
		bad := badUsers
		mu.Unlock()
		w.Header().Set("ETag", "same")
		if r.Method == "POST" {
			b, _ := io.ReadAll(r.Body)
			mu.Lock()
			bodies[action] = b
			mu.Unlock()
			w.WriteHeader(503)
			io.WriteString(w, "ignored")
			return
		}
		w.WriteHeader(503)
		switch action {
		case "config":
			io.WriteString(w, nodeJSON)
		case "user":
			if bad {
				io.WriteString(w, `{"msg":"denied"}`)
			} else {
				io.WriteString(w, `{"msg":"ok","data":[{"id":42,"secret":"pass","st":8,"dt":9}]}`)
			}
		}
	}))
	defer srv.Close()
	conf := &C.FCEAfQJwDA7e{W_ogGtF: srv.URL + "/base/", Bd08Li: "key", Yf1KqSqJpNI7: 71, ZyAhkq5th8bf: "ss"}
	p := D_5dOciQK_A(conf).(*DrEoaRCt2)
	if p.GetAPIConfig() != conf {
		t.Fatal("config identity")
	}
	for i := 0; i < 2; i++ {
		out, e := p.GetAPIData()
		if e != nil {
			t.Fatal(e)
		}
		if out.VTme1SuHL6i_[0].JAfXWwzYrW.(*L.YlfOV0DBBnx).T1oDDIgYJHh != "aes-128-gcm" || len(out.HYnfYWa) != 2 {
			t.Fatal(out)
		}
	}
	mu.Lock()
	badUsers = true
	mu.Unlock()
	if out, e := p.GetAPIData(); out != nil || e == nil || !strings.Contains(e.Error(), "convert users failed: denied") {
		t.Fatal(out, e)
	}
	if e := p.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{{Hy7Q97RROK: 42, XJkGTQf_qkAt: 3, I6LXObODav: 4}}, JNcHHwTZ7: L.CQhZrJ4dK{{Vy9ybu: 42, PNtn9KAV: []string{"1.2.3.4", "::1"}}}, Y0ys4ma4Tno4: L.QmiwDVjV00lZ{{GbzQUT_4f: 42, CAgUUdrv9: 99}}}); e != nil {
		t.Fatal(e)
	}
	mu.Lock()
	defer mu.Unlock()
	for _, a := range []string{"submit", "onlineusers", "illegal", "nodestatus"} {
		if requests[a] != 1 {
			t.Fatal(a, requests)
		}
	}
	if string(bodies["submit"]) != `[{"uid":42,"u":3,"d":4}]` || string(bodies["illegal"]) != `[{"uid":42}]` {
		t.Fatal(bodies)
	}
	var online []map[string]any
	if e := json.Unmarshal(bodies["onlineusers"], &online); e != nil || len(online) != 2 {
		t.Fatal(online, e)
	}
	for kind, want := range map[C.Mpn2JltM]string{"ss": "shadowsocks", "v2ray": "v2ray", "trojan": "trojan"} {
		conf.ZyAhkq5th8bf = kind
		u, e := p.d5nKoaGxYfB("x")
		if e != nil || !strings.Contains(u, "node_type="+want) {
			t.Fatal(u, e)
		}
	}
	conf.ZyAhkq5th8bf = "vless"
	if _, e := p.d5nKoaGxYfB("x"); e == nil {
		t.Fatal("unsupported")
	}
}
