package v2boardnode

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	L "recovered.local/soga/panel/types"
	C "recovered.local/soga/config"
	Q "recovered.local/soga/panel/registry"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestNativeProxypanelHTTPFlow(t *testing.T) {
	old := firstUsersAt378731C
	firstUsersAt378731C = true
	t.Cleanup(func() { firstUsersAt378731C = old })
	var mu sync.Mutex
	arrived := 0
	gate := make(chan struct{})
	posts := make(chan string, 8)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Key") != "panel-key" || r.Header.Get("Timestamp") == "" || r.URL.RawQuery != "" {
			t.Errorf("wrong auth %s %+v", r.URL, r.Header)
		}
		prefix := "/api/v2ray/v1/"
		if !strings.HasPrefix(r.URL.Path, prefix) {
			t.Error(r.URL.Path)
			w.WriteHeader(404)
			return
		}
		path := strings.TrimPrefix(r.URL.Path, prefix)
		if r.Method == "GET" {
			mu.Lock()
			arrived++
			if arrived == 3 {
				close(gate)
			}
			mu.Unlock()
			select {
			case <-gate:
			case <-time.After(3 * time.Second):
				t.Error("three original fetches did not overlap")
			}
			switch path {
			case "node/21":
				io.WriteString(w, `{"code":200,"data":{"v2_port":8443,"v2_alter_id":4,"v2_net":"http","v2_path":"/tunnel","v2_tls":true,"pem":"unused-cert","key":"unused-key","speed_limit":99,"client_limit":5}}`)
			case "userList/21":
				io.WriteString(w, `{"code":200,"data":[{"uid":17,"vmess_uid":"uuid-test","speed_limit":10}]}`)
			case "nodeRule/21":
				io.WriteString(w, `{"code":200,"data":{"mode":"black","rules":[{"id":31,"type":"domain","pattern":"blocked.test"}]}}`)
			default:
				t.Error(path)
				w.WriteHeader(404)
			}
		} else if r.Method == "POST" {
			var b any
			if e := json.NewDecoder(r.Body).Decode(&b); e != nil {
				t.Error(e)
			}
			posts <- path
			io.WriteString(w, `{"code":200}`)
		} else {
			t.Error(r.Method)
		}
	}))
	defer server.Close()
	cfg := &C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/", Bd08Li: "panel-key", Yf1KqSqJpNI7: 21, OS91X026uML: "proxypanel", ZyAhkq5th8bf: "v2ray", F3PCiaIra: "webapi"}
	factory, e := Q.SMkpVNb(cfg)
	if e != nil {
		t.Fatal(e)
	}
	p, ok := factory.(*UqRMl2W)
	if !ok || p.GetAPIConfig() != cfg {
		t.Fatalf("factory %T", factory)
	}
	result, e := p.GetAPIData()
	if e != nil {
		t.Fatal(e)
	}
	n := result.H47AzZsYHB1N
	if n == nil || n.EMlXEc != 21 || n.W3Dw_B4wgbM != 5 || n.TtWKyh != 99 || result.Xdo1D9 != 0 || result.DDu2Nm_P != 0 {
		t.Fatalf("node/result %+v %+v", n, result)
	}
	if len(result.HYnfYWa) != 1 || result.HYnfYWa[0].GetId() != 31 || !result.HYnfYWa[0].Test("x.blocked.test") {
		t.Fatal("detectors")
	}
	if len(result.VTme1SuHL6i_) != 1 {
		t.Fatal(result.VTme1SuHL6i_)
	}
	u := result.VTme1SuHL6i_[0]
	if u.WZ89pTAOrHK4 != 17 || u.Xmrna8 != 5 || u.LmOKwjha != 10 {
		t.Fatalf("user %+v", u)
	}
	if n.SwxJNgPpg.CKsOpayd != 8443 || n.SwxJNgPpg.F2GXyZZ != "h2" || n.SwxJNgPpg.ChJnngzb != "/tunnel" || !n.SwxJNgPpg.UVeahqDBT || n.LKMSis7zkX != "" || n.I3k4JnNp7nB != "" {
		t.Fatal(n.SwxJNgPpg)
	}
	credential, ok := u.JAfXWwzYrW.(*L.Yn8BJoL)
	if !ok || credential.SfiEGpJo != "uuid-test" || credential.FvQ2ltKOM7 != 4 {
		t.Fatal(u)
	}
	if !strings.Contains(n.Z1TIgHOQ, `"v2_net":"h2"`) {
		t.Fatal(n.Z1TIgHOQ)
	}

	if e = p.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{{Hy7Q97RROK: 17, XJkGTQf_qkAt: 100, I6LXObODav: 200}}, JNcHHwTZ7: L.CQhZrJ4dK{{Vy9ybu: 17, PNtn9KAV: []string{"127.0.0.1"}}}, Y0ys4ma4Tno4: L.QmiwDVjV00lZ{{GbzQUT_4f: 17, CAgUUdrv9: 31}}}); e != nil {
		t.Fatal(e)
	}
	want := map[string]bool{"nodeStatus/21": true, "userTraffic/21": true, "nodeOnline/21": true, "trigger/21": true}
	for range 4 {
		select {
		case path := <-posts:
			if !want[path] {
				t.Error(path)
			}
			delete(want, path)
		case <-time.After(3 * time.Second):
			t.Fatal("missing submit")
		}
	}
	if len(want) != 0 {
		t.Fatal(want)
	}
}
func TestNativeProxypanelNodeErrors(t *testing.T) {
	for _, tc := range []struct{ body, want string }{{`{"code":401,"message":"ignored"}`, "get node failed, code: 401"}, {`{`, "parse node failed: "}} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, tc.body) }))
		p := gHV0TDBY(&C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/"}).(*UqRMl2W)
		n, e := p.GetNode()
		server.Close()
		if n != nil || e == nil || !strings.HasPrefix(e.Error(), tc.want) {
			t.Fatalf("%+v %v", n, e)
		}
	}
}
