package v2raysocksnode

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	"reflect"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

func fixtureConfig(t *testing.T) {
	t.Helper()
	name := filepath.Join(t.TempDir(), "soga.conf")
	if e := os.WriteFile(name, []byte("type=v2raysocks\nserver_type=v2ray\nsoga_key=\nuser_speed_limit=12\nuser_conn_limit=4\nuser_tcp_limit=7\nnode_speed_limit=25\nforbidden_bit_torrent=true\n"), 0600); e != nil {
		t.Fatal(e)
	}
	C.MkwNWciPHS(name)
	if e := C.TvBECnD9fM6(); e != nil {
		t.Fatal(e)
	}
}
func testPanel(address string, kind C.Mpn2JltM) *W3tZSymoah {
	return JLydOCi(&C.FCEAfQJwDA7e{W_ogGtF: address, Bd08Li: "test-key", Yf1KqSqJpNI7: 71, ZyAhkq5th8bf: kind}).(*W3tZSymoah)
}

const nodeJSON = `{"inbounds":[{"tag":"ignore","port":1},{"tag":"proxy","port":8443,"settings":{"method":"aes-128-gcm","password":"node-secret"},"streamSettings":{"network":"grpc","security":"reality","wsSettings":{"path":"/raw"},"grpcSettings":{"serviceName":"soga"},"realitySettings":{"target":"example.invalid:443","serverNames":["example.invalid"],"privateKey":"private","shortIds":["ab"]}}},{"tag":"proxy","port":2}],"routing":{"rules":[{"outboundTag":"blocked"}]}}`

func TestRTTIAndNodeConversion(t *testing.T) {
	fixtureConfig(t)
	for _, v := range []struct {
		x any
		n uintptr
	}{{W3tZSymoah{}, 16}, {C7MFTE4gD{}, 32}, {QbI1asEYYq{}, 40}, {OBbQs9JkJPfq{}, 200}, {WsPqZpEZ{}, 72}, {MKWGBV{}, 72}, {FW_R_A{}, 24}, {Hjb4HGlua{}, 24}, {GFVixRJI2{}, 8}} {
		if reflect.TypeOf(v.x).Size() != v.n {
			t.Fatalf("RTTI %T", v.x)
		}
	}
	var r C7MFTE4gD
	if e := json.Unmarshal([]byte(nodeJSON), &r); e != nil {
		t.Fatal(e)
	}
	n, e := r.ToPanelNode(71)
	if e != nil {
		t.Fatal(e)
	}
	if n.EMlXEc != 71 || !n.HSA2b5V || n.TtWKyh != 25 || n.W3Dw_B4wgbM != 4 || n.FnMUny.Am3QoeF != 8443 || n.FnMUny.HMkc4BgR != "aes-128-gcm" || n.FnMUny.VuhWqou != "node-secret" {
		t.Fatalf("node base %+v", n)
	}
	if n.SwxJNgPpg.UVeahqDBT || n.SwxJNgPpg.ChJnngzb != "/raw" || n.JhIYwEukjd.QVlaSCZ4RCK != "soga" || n.Ck6DxjB.NMzMUQ2 != "xtls-rprx-vision" || n.Ck6DxjB.IqSfNzea != "example.invalid:443" {
		t.Fatal("transport conversion", n)
	}
	if n.Z1TIgHOQ != Q.T0j_Rco(r.Dps5ZWL[1]) {
		t.Fatal("hash scope")
	}
	r.Dps5ZWL[1].IHDmZMGzd.WmebQF3Gl7.DbrtlXUx[0] = "changed"
	if n.Ck6DxjB.BANVIZ72wh[0] != "changed" {
		t.Fatal("slice alias lost")
	}
	if d := r.GetDetectors(); d == nil || len(d) != 0 {
		t.Fatal("routing must produce nonnil empty detectors")
	}
	r.Dps5ZWL[1].IHDmZMGzd.FlU0GQuhpmy = "tls"
	n, e = r.ToPanelNode(1)
	if e != nil || !n.SwxJNgPpg.UVeahqDBT || n.Ck6DxjB.NMzMUQ2 != "" {
		t.Fatal("TLS exact mapping")
	}
	if _, e = (&C7MFTE4gD{}).ToPanelNode(1); e == nil || e.Error() != `invalid node config: can not find "proxy" tag` {
		t.Fatal(e)
	}
	if unsafe.Offsetof(OBbQs9JkJPfq{}.IouPnj9l) != 184 {
		t.Fatal("RTTI field offset")
	}
}
func TestUserConversionAndOriginalLimits(t *testing.T) {
	fixtureConfig(t)
	old := stateAt378733A
	defer func() { stateAt378733A = old; C.T15Jifo(false) }()
	u := &WsPqZpEZ{Gae2lLMaC4h: 42, PEtkkff: "ss-secret", CQ5RzKX: "uuid", JBz4bq: "trojan-secret", TjNsBPv: 8, U1cKBG6G62: 9}
	for kind, want := range map[C.Mpn2JltM]any{"ss": &L.YlfOV0DBBnx{A_wNUTkMh6pH: "ss-secret"}, "v2ray": &L.Yn8BJoL{SfiEGpJo: "uuid"}, "vless": &L.IImIwacp{Dh5iiakU: "uuid"}, "trojan": &L.JxTGUar{ARARP1z0yA: "trojan-secret"}, "unknown": nil} {
		stateAt378733A = 1
		got := u.ToPanelUser(kind, 0)
		if !reflect.DeepEqual(got.JAfXWwzYrW, want) || got.LmOKwjha != 8 || got.Xmrna8 != 4 || got.B4DYSxuA31za != 7 {
			t.Fatalf("%s %+v", kind, got)
		}
	}
	for _, v := range []struct{ speed, want float64 }{{-1, 12}, {0, 12}, {20000, 12}, {math.NaN(), math.NaN()}} {
		u.TjNsBPv = v.speed
		got := u.ToPanelUser("v2ray", 0).LmOKwjha
		if got != v.want && !(math.IsNaN(got) && math.IsNaN(v.want)) {
			t.Fatalf("speed %v -> %v", v.speed, got)
		}
	}
	// Real original restricted state, not a replacement authorization callback.
	C.T15Jifo(true)
	stateAt378733A = 1
	r := &QbI1asEYYq{DCbciSyJ_c: "not-ok"}
	for i := 0; i < 90; i++ {
		r.XH32wVY = append(r.XH32wVY, u)
	}
	if r.IsOk() {
		t.Fatal("IsOk exact literal")
	}
	r.DCbciSyJ_c = "ok"
	if !r.IsOk() {
		t.Fatal("IsOk")
	}
	first := r.ToPanelUsers("v2ray")
	if len(first) != 90 || first[89].JAfXWwzYrW == nil || stateAt378733A != 0 {
		t.Fatal("first-call state")
	}
	second := r.ToPanelUsers("v2ray")
	if len(second) != 90 || len(r.XH32wVY) != 90 {
		t.Fatal("user list must not be truncated")
	}
	if u.ToPanelUser("v2ray", 87).JAfXWwzYrW == nil || u.ToPanelUser("v2ray", 88).JAfXWwzYrW == nil {
		t.Fatal("per-user credentials must be present")
	}
}
func TestURLAndETagCache(t *testing.T) {
	call := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		call++
		if r.URL.Path != "/modules/addons/v2raysocks_nodes/web_api.php/" || r.URL.Query().Get("token") != "test-key" || r.URL.Query().Get("node_id") != "71" || r.URL.Query().Get("node_type") != "shadowsocks" || r.URL.Query().Get("act") != "config" {
			t.Errorf("request %v", r.URL)
		}
		if call > 1 && r.Header.Get("If-None-Match") != "v1" {
			t.Error("missing cache header")
		}
		if call == 3 {
			w.WriteHeader(304)
			return
		}
		w.Header().Set("ETag", "v1")
		if call == 1 {
			io.WriteString(w, "first")
		} else {
			io.WriteString(w, "ignored")
		}
	}))
	defer server.Close()
	p := testPanel(server.URL+"/", "ss")
	if p.GetAPIConfig() != p.iAFjt_tT9ZmG {
		t.Fatal("config identity")
	}
	for i := 0; i < 3; i++ {
		data, e := p.mSSnouymyTV1("config")
		if e != nil || string(data) != "first" {
			t.Fatal(i, string(data), e)
		}
	}
	// adwTxBv does not normalize a missing trailing slash or encode token.
	p = testPanel("http://127.0.0.1/base", "vless")
	p.iAFjt_tT9ZmG.Bd08Li = "a&b"
	got, e := p.adwTxBv("user")
	if e != nil || got != "http://127.0.0.1/basemodules/addons/v2raysocks_nodes/web_api.php/?token=a&b&node_id=71&node_type=v2ray&act=user" {
		t.Fatal(got, e)
	}
	for kind, want := range map[C.Mpn2JltM]string{"trojan": "trojan", "hysteria": "hysteria2", "v2ray": "v2ray"} {
		p.iAFjt_tT9ZmG.ZyAhkq5th8bf = kind
		s, e := p.adwTxBv("x")
		if e != nil || !strings.Contains(s, "node_type="+want+"&") {
			t.Fatal(s, e)
		}
	}
	p.iAFjt_tT9ZmG.ZyAhkq5th8bf = "invalid"
	if _, e = p.adwTxBv("x"); e == nil || e.Error() != "unsupported server type: invalid" {
		t.Fatal(e)
	}
}
func TestRequestErrorsAndCacheAssertions(t *testing.T) {
	status := 200
	tag := ""
	body := "body"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ETag", tag)
		w.WriteHeader(status)
		io.WriteString(w, body)
	}))
	defer server.Close()
	p := testPanel(server.URL+"/", "ss")
	address, _ := p.adwTxBv("config")
	if _, e := p.mSSnouymyTV1("config"); e != nil {
		t.Fatal(e)
	}
	if _, ok := p.oDI12LC.GetObj(address); ok {
		t.Fatal("untagged response cached")
	}
	status = 304
	if _, e := p.mSSnouymyTV1("config"); e == nil || e.Error() != "can not find eTag: "+address {
		t.Fatal(e)
	}
	status = 503
	if _, e := p.mSSnouymyTV1("config"); e == nil || e.Error() != "request failed, result: body, status code: 503" {
		t.Fatal(e)
	}
	if e := p.cpDvcLHkAY("submit", []int{1}); e != nil {
		t.Fatal("original POST ignores status", e)
	}
	if e := p.cpDvcLHkAY("submit", make(chan int)); e == nil || !strings.Contains(e.Error(), "status code: 0") {
		t.Fatal("marshal error", e)
	}
	status = 200
	tag = "v1"
	if _, e := p.mSSnouymyTV1("config"); e != nil {
		t.Fatal(e)
	}
	status = 503
	body = "error"
	p.mSSnouymyTV1("config")
	status = 304
	if data, e := p.mSSnouymyTV1("config"); e != nil || string(data) != "body" {
		t.Fatal("failure corrupted cache", string(data), e)
	}
	p.oDI12LC.AddETag(address, "v1", "wrong type")
	func() {
		defer func() {
			if recover() == nil {
				t.Error("original cached type assertion lost")
			}
		}()
		p.mSSnouymyTV1("config")
	}()
	p = testPanel("://bad/", "ss")
	if _, e := p.mSSnouymyTV1("config"); e == nil || !strings.HasSuffix(e.Error(), "status code: 0") {
		t.Fatal(e)
	}
}
func TestGetAPIDataActualRequestsAndFailureSemantics(t *testing.T) {
	fixtureConfig(t)
	old := stateAt378733A
	defer func() { stateAt378733A = old }()
	config, user := nodeJSON, `{"msg":"not-ok","data":[{"id":42,"secret":"ss-secret","uuid":"uuid","password":"trojan-secret","dt":2,"st":8}]}`
	var mu sync.Mutex
	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		act := r.URL.Query().Get("act")
		mu.Lock()
		calls[act]++
		mu.Unlock()
		if act == "config" {
			io.WriteString(w, config)
		} else {
			io.WriteString(w, user)
		}
	}))
	defer server.Close()
	for _, kind := range []C.Mpn2JltM{"ss", "v2ray", "vless", "trojan"} {
		p := testPanel(server.URL+"/", kind)
		data, e := p.GetAPIData()
		if e != nil {
			t.Fatal(kind, e)
		}
		if len(data.VTme1SuHL6i_) != 1 || data.HYnfYWa == nil || data.Xdo1D9 != 0 || data.DDu2Nm_P != 0 {
			t.Fatal("data shape", data)
		}
		if kind == "ss" {
			c := data.VTme1SuHL6i_[0].JAfXWwzYrW.(*L.YlfOV0DBBnx)
			if c.T1oDDIgYJHh != "aes-128-gcm" || c.A_wNUTkMh6pH != "ss-secret" {
				t.Fatal(c)
			}
		}
	}
	for _, tc := range []struct{ node, user, prefix string }{{"{", user, "parse node failed:"}, {`{"inbounds":[]}`, user, "convert node failed:"}, {nodeJSON, "{", "parse users failed:"}} {
		config, user = tc.node, tc.user
		p := testPanel(server.URL+"/", "ss")
		data, e := p.GetAPIData()
		if data != nil || e == nil || !strings.Contains(e.Error(), tc.prefix) {
			t.Fatalf("%+v %v", data, e)
		}
	}
	if calls["config"] != 7 || calls["user"] != 7 {
		t.Fatal("both concurrent branches must run", calls)
	}
}
func TestSubmitActualBodiesAndEmptyRequests(t *testing.T) {
	var mu sync.Mutex
	got := map[string]string{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, e := io.ReadAll(r.Body)
		if e != nil {
			t.Error(e)
		}
		if r.Method != "POST" || r.Header.Get("Content-Type") != "application/json" {
			t.Error("request shape")
		}
		mu.Lock()
		got[r.URL.Query().Get("act")] = string(b)
		mu.Unlock()
		w.WriteHeader(503)
		io.WriteString(w, `{"failed":true}`)
	}))
	defer server.Close()
	p := testPanel(server.URL+"/", "vless")
	data := &Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{{Hy7Q97RROK: 42, XJkGTQf_qkAt: 10, I6LXObODav: 20}, {Hy7Q97RROK: 42, XJkGTQf_qkAt: 30, I6LXObODav: 40}}, JNcHHwTZ7: L.CQhZrJ4dK{{Vy9ybu: 42, PNtn9KAV: []string{"127.0.0.1", "::1"}}}, Y0ys4ma4Tno4: L.QmiwDVjV00lZ{{GbzQUT_4f: 42, CAgUUdrv9: 123}}}
	if e := p.SubmitData(data); e != nil {
		t.Fatal(e)
	}
	for act, want := range map[string]string{"submit": `[{"uid":42,"u":10,"d":20},{"uid":42,"u":30,"d":40}]`, "onlineusers": `[{"uid":42,"ip":"127.0.0.1"},{"uid":42,"ip":"::1"}]`, "illegal": `[{"uid":42}]`} {
		if got[act] != want {
			t.Fatalf("%s=%s", act, got[act])
		}
	}
	var status MKWGBV
	if e := json.Unmarshal([]byte(got["nodestatus"]), &status); e != nil {
		t.Fatal(e)
	}
	if status.AUezFKXr != "" || status.FmIi0I4dOmz0 == 0 || !strings.HasSuffix(status.MwG7GAHIuy, "%") || !strings.HasSuffix(status.A6pOwZFaGw, "%") || !strings.HasSuffix(status.AazLqAvDDU, "%") {
		t.Fatalf("status %+v", status)
	}
	if e := p.SubmitData(&Q.FGmNba4MhioY{}); e != nil {
		t.Fatal(e)
	}
	for _, act := range []string{"submit", "onlineusers", "illegal"} {
		if got[act] != "[]" {
			t.Fatal("empty collection must still submit", act, got[act])
		}
	}
	if len(got) != 4 {
		t.Fatal(fmt.Sprint(got))
	}
}

func TestTruncatedHTTPAndAggregateFetchFailures(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, buffer, e := w.(http.Hijacker).Hijack()
		if e != nil {
			t.Error(e)
			return
		}
		defer conn.Close()
		buffer.WriteString("HTTP/1.1 200 OK\r\nContent-Length: 10\r\nETag: partial\r\nConnection: close\r\n\r\nx")
		buffer.Flush()
	}))
	defer server.Close()
	p := testPanel(server.URL+"/", "v2ray")
	if data, e := p.mSSnouymyTV1("config"); data != nil || e == nil || e.Error() != "result: , err: unexpected EOF, status code: 200" {
		t.Fatalf("GET: %q %v", data, e)
	}
	address, _ := p.adwTxBv("config")
	if _, ok := p.oDI12LC.GetObj(address); ok {
		t.Fatal("partial response cached")
	}
	if e := p.cpDvcLHkAY("submit", []int{}); e == nil || e.Error() != "result: , err: unexpected EOF, status code: 200" {
		t.Fatalf("POST: %v", e)
	}
	data, e := p.GetAPIData()
	if data != nil || e == nil || !strings.Contains(e.Error(), "get node failed:") || !strings.Contains(e.Error(), "get users failed:") {
		t.Fatalf("aggregate: %v %v", data, e)
	}
}
