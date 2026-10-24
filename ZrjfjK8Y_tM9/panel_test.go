package ZrjfjK8Y_tM9

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"recovered.local/soga/LFWTTXF7JR"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/q5bqCgTdxaTM"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestPanelFullHTTPFlow(t *testing.T) {
	preservePanelByte(t)
	stateAt3787339 = 1
	var mu sync.Mutex
	var arrived int
	gate := make(chan struct{})
	counts := map[string]int{}
	var pushed map[string][2]int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("token") != "secret" || q.Get("node_id") != "14" || q.Get("node_type") != "shadowsocks" {
			t.Errorf("query %v", q)
		}
		switch r.URL.Path {
		case "/api/v1/server/UniProxy/config", "/api/v1/server/UniProxy/user":
			mu.Lock()
			counts[r.URL.Path]++
			arrived++
			if arrived == 2 {
				close(gate)
			}
			mu.Unlock()
			select {
			case <-gate:
			case <-time.After(3 * time.Second):
				t.Error("node/user requests were not concurrent")
			}
			if r.Header.Get("If-None-Match") == "v1" {
				w.WriteHeader(304)
				return
			}
			w.Header().Set("ETag", "v1")
			if strings.HasSuffix(r.URL.Path, "/config") {
				_, _ = w.Write([]byte(`{"server_port":8388,"cipher":"2022-blake3-aes-128-gcm","server_key":"node-secret","base_config":{"pull_interval":"12.9","push_interval":-2},"routes":[{"id":29,"action":"block","match":["domain:blocked.test"]}]}`))
			} else {
				_, _ = w.Write([]byte(`{"users":[{"id":3,"uuid":"0123456789abcdef-remaining","speed_limit":7,"device_limit":2},{"id":4,"uuid":"short"}]}`))
			}
		case "/api/v1/server/UniProxy/push":
			if r.Method != "POST" {
				t.Error(r.Method)
			}
			if err := json.NewDecoder(r.Body).Decode(&pushed); err != nil {
				t.Error(err)
			}
			_, _ = w.Write([]byte(`{}`))
		default:
			t.Error(r.URL.Path)
			w.WriteHeader(404)
		}
	}))
	defer server.Close()
	cfg := &cNT24rG4F.FCEAfQJwDA7e{W_ogGtF: server.URL + "/", Bd08Li: "secret", Yf1KqSqJpNI7: 14, ZyAhkq5th8bf: "ss"}
	panel := C0p2qbSZQL(cfg)
	if panel.GetAPIConfig() != cfg {
		t.Fatal("configuration pointer changed")
	}
	for attempt := 0; attempt < 2; attempt++ {
		result, err := panel.GetAPIData()
		if err != nil {
			t.Fatal(err)
		}
		if result.H47AzZsYHB1N.EMlXEc != 14 || result.H47AzZsYHB1N.FnMUny.Am3QoeF != 8388 || result.Xdo1D9 != 12*time.Second || result.DDu2Nm_P != -2*time.Second {
			t.Fatalf("result %+v", result)
		}
		if len(result.HYnfYWa) != 1 || result.HYnfYWa[0].GetId() != 29 || !result.HYnfYWa[0].Test("a.blocked.test") {
			t.Fatal("detectors")
		}
		if len(result.VTme1SuHL6i_) != 2 {
			t.Fatal("users")
		}
		long := result.VTme1SuHL6i_[0].JAfXWwzYrW.(*LFWTTXF7JR.YlfOV0DBBnx)
		short := result.VTme1SuHL6i_[1].JAfXWwzYrW.(*LFWTTXF7JR.YlfOV0DBBnx)
		if long.T1oDDIgYJHh != "2022-blake3-aes-128-gcm" || long.A_wNUTkMh6pH != base64.StdEncoding.EncodeToString([]byte("0123456789abcdef")) || short.A_wNUTkMh6pH != "short" {
			t.Fatalf("%+v %+v", long, short)
		}
	}
	if err := panel.SubmitData(&q5bqCgTdxaTM.FGmNba4MhioY{HTgPyE3VytZd: LFWTTXF7JR.CWfcUv{{Hy7Q97RROK: 3, XJkGTQf_qkAt: 100, I6LXObODav: 200}}}); err != nil {
		t.Fatal(err)
	}
	if pushed["3"] != [2]int64{100, 200} {
		t.Fatal(pushed)
	}
	mu.Lock()
	defer mu.Unlock()
	if counts["/api/v1/server/UniProxy/config"] != 2 || counts["/api/v1/server/UniProxy/user"] != 2 {
		t.Fatal(counts)
	}
}

func TestAPIDataPartialAndCipherFailures(t *testing.T) {
	preservePanelByte(t)
	for _, tc := range []struct {
		name, node, users            string
		wantNode, wantUsers, wantNil bool
		errorText                    string
	}{
		{"node_error", `{"message":"bad-node"}`, `{"users":[{"id":1,"uuid":"u"}]}`, false, true, false, "parse node failed"},
		{"user_error", `{"server_port":1,"cipher":"aes-128-gcm"}`, `{"message":"bad-users"}`, true, false, false, "parse users failed"},
		{"cipher_error", `{"cipher":"unsupported"}`, `{"users":[]}`, false, false, true, "ss node err: unknown ss cipher <unsupported> key len"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			stateAt3787339 = 1
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasSuffix(r.URL.Path, "/config") {
					_, _ = w.Write([]byte(tc.node))
				} else {
					_, _ = w.Write([]byte(tc.users))
				}
			}))
			defer server.Close()
			p := C0p2qbSZQL(&cNT24rG4F.FCEAfQJwDA7e{W_ogGtF: server.URL + "/", ZyAhkq5th8bf: "ss"})
			out, err := p.GetAPIData()
			if err == nil || !strings.Contains(err.Error(), tc.errorText) {
				t.Fatalf("%+v %v", out, err)
			}
			if tc.wantNil {
				if out != nil {
					t.Fatal("cipher failure must return nil")
				}
				return
			}
			if out == nil || (out.H47AzZsYHB1N != nil) != tc.wantNode || (out.VTme1SuHL6i_ != nil) != tc.wantUsers {
				t.Fatalf("partial output %+v", out)
			}
		})
	}
}
