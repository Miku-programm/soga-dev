package xboard

import (
	"encoding/json"
	"fmt"
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

func TestSevenProtocolHTTPFetchCacheAndParallelSubmission(t *testing.T) {
	old := stateAt378733B
	t.Cleanup(func() { stateAt378733B = old })
	for _, protocol := range []C.Mpn2JltM{"ss", "v2ray", "trojan", "vless", "hysteria", "anytls", "mieru"} {
		t.Run(string(protocol), func(t *testing.T) {
			stateAt378733B = 1
			var mu sync.Mutex
			seen := map[string]int{}
			fetched := make(chan struct{})
			submitted := make(chan struct{})
			var posts int
			var traffic map[string][2]int64
			var alive map[string][]string
			var status map[string]json.RawMessage
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
				if action == "push" || action == "alive" || action == "status" {
					if r.Method != "POST" {
						t.Error(r.Method)
					}
					if action == "push" {
						if err := json.NewDecoder(r.Body).Decode(&traffic); err != nil {
							t.Error(err)
						}
					} else if action == "status" {
						if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
							t.Error(err)
						}
					} else {
						if err := json.NewDecoder(r.Body).Decode(&alive); err != nil {
							t.Error(err)
						}
					}
					mu.Lock()
					posts++
					if posts == 3 {
						close(submitted)
					}
					mu.Unlock()
					select {
					case <-submitted:
					case <-time.After(3 * time.Second):
						t.Error("POSTs serialized")
					}
					fmt.Fprint(w, "ok")
					return
				}
				t.Error(r.URL)
				w.WriteHeader(404)
			}))
			defer server.Close()
			cfg := &C.FCEAfQJwDA7e{OS91X026uML: "xboard", ZyAhkq5th8bf: protocol, F3PCiaIra: "webapi", W_ogGtF: server.URL + "/", Bd08Li: "test key", Yf1KqSqJpNI7: 71}
			panel := NiF7OSBK(cfg)
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
			if !reflect.DeepEqual(traffic, map[string][2]int64{"13": {7, 8}}) || !reflect.DeepEqual(alive, map[string][]string{"13": {"192.0.2.1"}}) {
				t.Fatal(traffic, alive)
			}
			if len(status) != 4 || status["cpu"] == nil || status["mem"] == nil || status["swap"] == nil || status["disk"] == nil {
				t.Fatalf("status: %v", status)
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
	p := NiF7OSBK(&C.FCEAfQJwDA7e{W_ogGtF: "https://example.test/base", ZyAhkq5th8bf: "UNKNOWN"}).(*KlO65BuwOt)
	if p.rTN7wPdD("config") != "https://example.test/baseapi/v1/server/UniProxy/config" || p.tgg_NuFN().Has("node_type") {
		t.Fatal("normalization")
	}
	for _, tc := range []struct {
		status     int
		body, kind string
	}{{304, "", "can not find eTag:"}, {500, "denied", "request failed, result: denied, status code: 500"}, {200, `{"message":"denied"}`, "parse node failed: <nil>"}, {200, "broken", "parse node failed:"}} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(tc.status); fmt.Fprint(w, tc.body) }))
		p := NiF7OSBK(&C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/"}).(*KlO65BuwOt)
		if _, err := p.iEok3SmDhK(); err == nil || !strings.Contains(err.Error(), tc.kind) {
			t.Fatal(tc, err)
		}
		server.Close()
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(201); fmt.Fprint(w, "created") }))
	defer server.Close()
	p = NiF7OSBK(&C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/"}).(*KlO65BuwOt)
	err := p.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{{Hy7Q97RROK: 1}}})
	if err == nil || !strings.Contains(err.Error(), "post status code: 201, result: created") {
		t.Fatal(err)
	}
	if err := p.SubmitData(&Q.FGmNba4MhioY{}); err == nil {
		t.Fatal(err)
	}
}

func TestNilDataStillReportsStatusAndAliveErrorsPropagate(t *testing.T) {
	var mu sync.Mutex
	var actions []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action := strings.TrimPrefix(r.URL.Path, "/api/v1/server/UniProxy/")
		mu.Lock()
		actions = append(actions, action)
		mu.Unlock()
		if action == "alive" {
			w.WriteHeader(503)
			fmt.Fprint(w, "alive error")
		} else {
			fmt.Fprint(w, "ok")
		}
	}))
	defer server.Close()
	p := NiF7OSBK(&C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/"}).(*KlO65BuwOt)
	if err := p.SubmitData(nil); err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	first := append([]string(nil), actions...)
	mu.Unlock()
	if !reflect.DeepEqual(first, []string{"status"}) {
		t.Fatal(first)
	}
	err := p.SubmitData(&Q.FGmNba4MhioY{JNcHHwTZ7: L.CQhZrJ4dK{{Vy9ybu: 2, PNtn9KAV: []string{"192.0.2.9"}}}})
	if err == nil || !strings.Contains(err.Error(), "post status code: 503, result: alive error") {
		t.Fatal(err)
	}
}

func TestFailedNodeFetchPreservesUsersAndSameETagKeepsCachedBody(t *testing.T) {
	stateAt378733B = 1
	var version int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "config") {
			w.WriteHeader(502)
			fmt.Fprint(w, "node failed")
			return
		}
		w.Header().Set("ETag", "constant")
		fmt.Fprintf(w, `{"users":[{"id":%d,"uuid":"user"}]}`, version)
	}))
	defer server.Close()
	p := NiF7OSBK(&C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/", ZyAhkq5th8bf: "mieru"}).(*KlO65BuwOt)
	version = 23
	out, err := p.GetAPIData()
	if err == nil || out.H47AzZsYHB1N != nil || len(out.VTme1SuHL6i_) != 1 || out.VTme1SuHL6i_[0].WZ89pTAOrHK4 != 23 {
		t.Fatal(out, err)
	}
	cred, ok := out.VTme1SuHL6i_[0].JAfXWwzYrW.(*L.DO9qqM9)
	if !ok || cred.Ejujyaj != "user" || cred.GltkJGvV != "user" {
		t.Fatal(cred)
	}
	version = 24
	users, err := p.kDrbLuahIFA6()
	if err != nil || len(users.Ta593Rq2lSa) != 1 || users.Ta593Rq2lSa[0].IDLJXX6F != 23 {
		t.Fatal(users, err)
	}
}
