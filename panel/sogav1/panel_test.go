package sogav1

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func panelForTest(address string, kind C.Mpn2JltM) *CsKk5VzGb {
	return DDky8V(&C.FCEAfQJwDA7e{W_ogGtF: address + "/", Bd08Li: "fixture-key", Yf1KqSqJpNI7: 17, ZyAhkq5th8bf: kind}).(*CsKk5VzGb)
}

func TestNativeSogaPanelHTTPRoundTrip(t *testing.T) {
	saved := stateAt3787335
	t.Cleanup(func() { stateAt3787335 = saved })
	stateAt3787335 = 1
	var mu sync.Mutex
	counts := map[string]int{}
	pushed := map[string]json.RawMessage{}
	ready := make(chan struct{})
	arrived := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("API-KEY") != "fixture-key" || r.Header.Get("NODE-ID") != "17" || r.Header.Get("NODE-TYPE") != "shadowsocksr" || r.URL.RawQuery != "" {
			t.Errorf("request %s %v", r.URL, r.Header)
		}
		name := strings.TrimPrefix(r.URL.Path, "/api/v1/")
		if r.Method == "POST" {
			var body json.RawMessage
			if e := json.NewDecoder(r.Body).Decode(&body); e != nil {
				t.Error(e)
			}
			mu.Lock()
			pushed[name] = body
			mu.Unlock()
			fmt.Fprint(w, `{"code":0,"message":"ok"}`)
			return
		}
		mu.Lock()
		counts[name]++
		n := counts[name]
		if n == 1 {
			arrived++
			if arrived == 4 {
				close(ready)
			}
		}
		mu.Unlock()
		select {
		case <-ready:
		case <-time.After(3 * time.Second):
			t.Error("four fetches did not run concurrently")
			w.WriteHeader(500)
			return
		}
		if n == 2 {
			if r.Header.Get("If-None-Match") != `"original"` {
				t.Error("missing ETag")
			}
			w.WriteHeader(304)
			return
		}
		w.Header().Set("ETag", `"original"`)
		switch name {
		case "node":
			fmt.Fprint(w, `{"basic":{"pull_interval":11,"push_interval":13,"speed_limit":25},"config":{"port":31234,"stream_type":"grpc","tls_type":"tls","path":"/stream","service_name":"srv","cipher":"aes-128-gcm","password":"carrier","host":"host.test","obfs":"plain","method":"aes-128-cfb","protocol":"auth_aes128_md5","single_port_type":"obfs","flow":"xtls-rprx-vision","server_names":["server.test"],"private_key":"private","short_ids":["abcd"],"dest":"dest.test:443","decryption":"none","xhttp_settings":{"host":"h","path":"/x","mode":"stream-one","extra":{"headers":{"X-F":"yes"},"scMaxBufferedPosts":7}},"up_mbps":70,"down_mbps":80,"padding_scheme":["one","two"],"transport":"tCp","traffic_pattern":"pattern"}}`)
		case "users":
			fmt.Fprint(w, `[{"id":101,"uuid":"uuid","password":"secret","username":"name","speed_limit":12,"device_limit":2,"tcp_limit":3}]`)
		case "audit_rules":
			fmt.Fprint(w, `[{"id":4,"rule":"blocked.test"}]`)
		case "white_list":
			fmt.Fprint(w, `["allowed.test"]`)
		default:
			t.Error("unexpected path", name)
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()
	p := panelForTest(srv.URL, "ssr")
	for i := 0; i < 2; i++ {
		result, e := p.GetAPIData()
		if e != nil {
			t.Fatal(e)
		}
		n := result.H47AzZsYHB1N
		if n.EMlXEc != 17 || n.SwxJNgPpg.CKsOpayd != 31234 || !n.SwxJNgPpg.UVeahqDBT || n.JhIYwEukjd.QVlaSCZ4RCK != "srv" || n.Ck6DxjB.FPhrzP645.HoBz4V != "/x" || n.QaZjDzw95.UPnToXrM != "one\ntwo" || n.ABA4hq.GGSLzD != "TCP" || n.AYf5kQSQW.Li6nvK1 != 2 {
			t.Fatalf("node conversion: %+v", n)
		}
		if result.Xdo1D9 != 0 || result.DDu2Nm_P != 0 || len(result.VTme1SuHL6i_) != 1 || len(result.HYnfYWa) != 1 || !reflect.DeepEqual([]string(result.NjjWbjH), []string{"allowed.test"}) {
			t.Fatalf("aggregate %+v", result)
		}
		u := result.VTme1SuHL6i_[0]
		c := u.JAfXWwzYrW.(*L.QcqpSL3i6ri)
		if u.WZ89pTAOrHK4 != 101 || c.YyAgnMatmj != "secret" || c.YrjVl_2g_S != 31234 || c.Z_taAb != "aes-128-cfb" || c.H_y2xZht7S8 != "auth_aes128_md5" || c.BskzckkE7qI != 0 {
			t.Fatalf("user %+v/%+v", u, c)
		}
		carrier := n.ASJxBbFS.SinglePortUsers()
		if len(carrier) != 1 || carrier[0].WZ89pTAOrHK4 != 0 || carrier[0].JAfXWwzYrW.(*L.QcqpSL3i6ri).BskzckkE7qI != 1 {
			t.Fatal("SSR carrier", carrier)
		}
	}
	e := p.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{&L.DLxslMwYJXQw{Hy7Q97RROK: 101, XJkGTQf_qkAt: 5, I6LXObODav: 8}}, JNcHHwTZ7: L.CQhZrJ4dK{&L.LmNKoY94{Vy9ybu: 101, PNtn9KAV: []string{"192.0.2.8"}}}, Y0ys4ma4Tno4: L.QmiwDVjV00lZ{&L.B9FS114R4F9h{GbzQUT_4f: 101, CAgUUdrv9: 4}}})
	if e != nil {
		t.Fatal(e)
	}
	mu.Lock()
	defer mu.Unlock()
	for path, want := range map[string]string{"traffic": `[{"id":101,"u":5,"d":8}]`, "alive_ip": `[{"id":101,"ips":["192.0.2.8"]}]`, "audit_log": `[{"user_id":101,"audit_id":4}]`} {
		if string(pushed[path]) != want {
			t.Errorf("%s: %s", path, pushed[path])
		}
	}
	var state map[string]any
	if e := json.Unmarshal(pushed["status"], &state); e != nil {
		t.Fatal(e)
	}
	if len(state) != 5 || state["uptime"] == nil || state["mem"] == nil {
		t.Fatal(state)
	}
}

func TestNativeSogaCacheAndErrors(t *testing.T) {
	status, body, etag := 200, "first", `"x"`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ETag", etag)
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}))
	defer srv.Close()
	p := panelForTest(srv.URL, "ss")
	first, e := p.sXWeuE("node")
	if e != nil {
		t.Fatal(e)
	}
	first[0] = 'F'
	body = "second"
	cached, e := p.sXWeuE("node")
	if e != nil || string(cached) != "First" {
		t.Fatal(string(cached), e)
	}
	status = 304
	cached, e = p.sXWeuE("node")
	if e != nil || string(cached) != "First" {
		t.Fatal(e)
	}
	if _, e = p.sXWeuE("missing"); e == nil || !strings.HasPrefix(e.Error(), "can not find eTag:") {
		t.Fatal(e)
	}
	status = 500
	body = "bad"
	if _, e = p.GetNode(); e == nil || e.Error() != "get node failed: request failed, result: bad, status code: 500" {
		t.Fatal(e)
	}
	for _, tc := range []struct {
		body   string
		prefix string
	}{{`{"code":1,"message":"denied"}`, "post failed: denied, data:"}, {"bad", "parse post result failed:"}, {`{}`, ""}, {`null`, ""}} {
		status = 200
		body = tc.body
		e = p.olDPHF3j("traffic", []int{1})
		if tc.prefix == "" {
			if e != nil {
				t.Fatal(e)
			}
		} else if e == nil || !strings.HasPrefix(e.Error(), tc.prefix) {
			t.Fatal(e)
		}
	}
	if e = p.olDPHF3j("traffic", make(chan int)); e == nil || !strings.HasPrefix(e.Error(), "post ") {
		t.Fatal(e)
	}
}

func TestOriginalCredentialSelectionAndLimits(t *testing.T) {
	u := &D6uXyl88ab7{IMizGzpEPYfj: 8, S7USjTWE8o: "uuid", DHknGz: "password", WUNMcwtiDw: "username", GZJoKAlg8K9: -1, LNDBhsiH: -1, DGYgUabBB6Jx: -1}
	for _, kind := range []C.Mpn2JltM{"ss", "ssr", "v2ray", "vless", "trojan", "hysteria", "anytls", "mieru", "unknown"} {
		got := u.ToPanelUser(kind)
		var value string
		switch c := got.JAfXWwzYrW.(type) {
		case *L.YlfOV0DBBnx:
			value = c.A_wNUTkMh6pH
		case *L.QcqpSL3i6ri:
			value = c.YyAgnMatmj
		case *L.Yn8BJoL:
			value = c.SfiEGpJo
		case *L.IImIwacp:
			value = c.Dh5iiakU
		case *L.JxTGUar:
			value = c.ARARP1z0yA
		case *L.WQM2tKA:
			value = c.CBWeRfrO1ahs
		case *L.ZtMFlBTZXyD:
			value = c.CBWeRfrO1ahs
		case *L.DO9qqM9:
			value = c.GltkJGvV
			if c.Ejujyaj != "username" {
				t.Fatal(c)
			}
		}
		want := "password"
		if kind == "v2ray" || kind == "vless" {
			want = "uuid"
		}
		if kind == "unknown" {
			want = ""
		}
		if value != want {
			t.Fatal(kind, value)
		}
	}
	for _, tc := range []struct{ a, b, w float64 }{{-1, 20, 20}, {30, 20, 20}, {0, 5, 5}, {20000, 0, 10000}, {math.NaN(), 7, 7}, {7, math.NaN(), math.NaN()}} {
		v := mergeSpeed(tc.a, tc.b)
		if math.IsNaN(tc.w) {
			if !math.IsNaN(v) {
				t.Fatal(v)
			}
		} else if v != tc.w {
			t.Fatal(v, tc.w)
		}
	}
	n := new(GX5_vLs)
	n.XI5uZkt.H5wL4HQdqP = -3
	n.XI5uZkt.ULhXjJaRde = 5
	if n.GetPullInterval() != -3*time.Second || n.GetPushInterval() != 5*time.Second {
		t.Fatal("interval")
	}
}
