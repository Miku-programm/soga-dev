package aRZNjEvm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestSixProtocolHTTPFetchCacheAndParallelSubmission(t *testing.T) {
	old := stateAt378733C
	t.Cleanup(func() { stateAt378733C = old })
	for _, protocol := range []C.Mpn2JltM{"ss", "v2ray", "trojan", "vless", "hysteria", "anytls"} {
		t.Run(string(protocol), func(t *testing.T) {
			stateAt378733C = 1
			var mu sync.Mutex
			seen := map[string]int{}
			fetched := make(chan struct{})
			submitted := make(chan struct{})
			var posts int
			var traffic map[string][2]int64
			var alive map[int64][]string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nodeType := string(protocol)
				if protocol == "ss" {
					nodeType = "shadowsocks"
				}
				if q := r.URL.Query(); q.Get("token") != "test key" || q.Get("node_id") != "71" || q.Get("node_type") != nodeType {
					t.Errorf("query: %v", q)
				}
				action := strings.TrimPrefix(r.URL.Path, "/api/v1/server/UniProxy/")
				if action == "config" || action == "user" {
					mu.Lock()
					seen[action]++
					if seen["config"]+seen["user"] == 2 {
						close(fetched)
					}
					mu.Unlock()
					select {
					case <-fetched:
					case <-time.After(3 * time.Second):
						t.Error("GETs serialized")
					}
					if r.Header.Get("If-None-Match") == "fixture-v1" {
						w.WriteHeader(304)
						return
					}
					w.Header().Set("ETag", "fixture-v1")
					if action == "config" {
						fmt.Fprint(w, `{"server_port":8443,"network":"ws","networkSettings":{"path":"/stream"},"tls":1,"cipher":"aes-128-gcm","base_config":{"pull_interval":"13.9","push_interval":27},"routes":[{"id":14,"action":"block","match":["domain:deny.example"]},{"action":"dns","match":["domain:local.example"],"action_value":"1.1.1.1"}]}`)
					} else {
						fmt.Fprint(w, `{"users":[{"id":13,"uuid":"01234567-89ab-cdef-0123-456789abcdef","speed_limit":8,"device_limit":3}]}`)
					}
					return
				}
				if action == "push" || action == "alive" {
					if r.Method != "POST" {
						t.Error(r.Method)
					}
					if action == "push" {
						if err := json.NewDecoder(r.Body).Decode(&traffic); err != nil {
							t.Error(err)
						}
					} else {
						if err := json.NewDecoder(r.Body).Decode(&alive); err != nil {
							t.Error(err)
						}
					}
					mu.Lock()
					posts++
					if posts == 2 {
						close(submitted)
					}
					mu.Unlock()
					select {
					case <-submitted:
					case <-time.After(3 * time.Second):
						t.Error("POSTs serialized")
					}
					if action == "alive" {
						w.WriteHeader(503)
						fmt.Fprint(w, "alive failure intentionally ignored")
					} else {
						fmt.Fprint(w, "ok")
					}
					return
				}
				t.Error(r.URL)
				w.WriteHeader(404)
			}))
			defer server.Close()
			cfg := &C.FCEAfQJwDA7e{OS91X026uML: "xiaov2board", ZyAhkq5th8bf: protocol, F3PCiaIra: "webapi", W_ogGtF: server.URL + "/", Bd08Li: "test key", Yf1KqSqJpNI7: 71}
			panel := AnxiJP(cfg)
			if panel.GetAPIConfig() != cfg {
				t.Fatal("pointer copy")
			}
			for i := 0; i < 2; i++ {
				out, err := panel.GetAPIData()
				if err != nil {
					t.Fatal(err)
				}
				if out.H47AzZsYHB1N.EMlXEc != 71 || out.H47AzZsYHB1N.SwxJNgPpg.ChJnngzb != "/stream" || out.Xdo1D9 != 13*time.Second || out.DDu2Nm_P != 27*time.Second || len(out.VTme1SuHL6i_) != 1 || len(out.HYnfYWa) != 1 || !out.HYnfYWa[0].Test("a.deny.example") || len(out.H47AzZsYHB1N.AjXcBfRcAXR) != 1 {
					t.Fatalf("%+v", out)
				}
			}
			if err := panel.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{{Hy7Q97RROK: 13, XJkGTQf_qkAt: 5, I6LXObODav: 6}, {Hy7Q97RROK: 13, XJkGTQf_qkAt: 7, I6LXObODav: 8}}, JNcHHwTZ7: L.CQhZrJ4dK{{Vy9ybu: 13, PNtn9KAV: []string{"old"}}, {Vy9ybu: 13, PNtn9KAV: []string{"192.0.2.1"}}}}); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(traffic, map[string][2]int64{"13": {7, 8}}) || !reflect.DeepEqual(alive, map[int64][]string{13: {"192.0.2.1"}}) {
				t.Fatal(traffic, alive)
			}
			mu.Lock()
			defer mu.Unlock()
			if seen["config"] != 2 || seen["user"] != 2 {
				t.Fatal(seen)
			}
		})
	}
}

func TestHTTPFailuresAndNoURLNormalization(t *testing.T) {
	p := AnxiJP(&C.FCEAfQJwDA7e{W_ogGtF: "https://example.test/base", ZyAhkq5th8bf: "UNKNOWN"}).(*LRlqOaYmO)
	if p.GetUniUrl("config") != "https://example.test/baseapi/v1/server/UniProxy/config" || p.fUwRFGkp().Has("node_type") {
		t.Fatal("normalization")
	}
	for _, tc := range []struct {
		status     int
		body, kind string
	}{{304, "", "can not find eTag:"}, {500, "denied", "request failed, result: denied, status code: 500"}, {200, `{"message":"denied"}`, "parse node failed: <nil>"}, {200, "broken", "parse node failed:"}} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(tc.status); fmt.Fprint(w, tc.body) }))
		p := AnxiJP(&C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/"}).(*LRlqOaYmO)
		if _, err := p.GetUniNode(); err == nil || !strings.Contains(err.Error(), tc.kind) {
			t.Fatal(tc, err)
		}
		server.Close()
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(201); fmt.Fprint(w, "created") }))
	defer server.Close()
	p = AnxiJP(&C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/"}).(*LRlqOaYmO)
	err := p.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{{Hy7Q97RROK: 1}}})
	if err == nil || !strings.Contains(err.Error(), "post status code: 201, result: created") {
		t.Fatal(err)
	}
	if err := p.SubmitData(&Q.FGmNba4MhioY{}); err != nil {
		t.Fatal(err)
	}
}
