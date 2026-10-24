package userhash

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
	"testing"
	"time"
	"unsafe"
)

func TestDJHTTPConfigurationAndTraffic(t *testing.T) {
	previous := emailByIDAt5095E48
	t.Cleanup(func() { emailByIDAt5095E48 = previous })
	var traffic map[string]map[string]int64
	var paths []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		if r.Method != "POST" {
			t.Error(r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Error(err)
		}
		if r.PostForm.Get("secret") != "a +/ key" {
			t.Error(r.PostForm)
		}
		switch r.URL.Path {
		case "/sub_server/get_config":
			w.WriteHeader(503)
			fmt.Fprint(w, `{"inbounds":[{"port":1,"settings":{"clients":[{"email":"ignored"}]}},{"port":8843,"settings":{"clients":[{"id":"uuid-a","email":"a@example.com","alterId":9},{"id":"uuid-b","email":"b@example.com","alterId":0}]}}],"routing":{"rules":[{"domain":["domain:blocked.example","keyword:bad"]},{"domain":["full:exact.example"]}]}}`)
		case "/sub_server/submit_traffic":
			if err := json.Unmarshal([]byte(r.PostForm.Get("email_traffic")), &traffic); err != nil {
				t.Error(err)
			}
			w.WriteHeader(500)
		default:
			t.Error(r.URL.Path)
		}
	}))
	defer server.Close()
	cfg := &C.FCEAfQJwDA7e{OS91X026uML: "dj", ZyAhkq5th8bf: "v2ray", F3PCiaIra: "webapi", W_ogGtF: server.URL + "/", Bd08Li: "a +/ key"}
	panel, err := Q.SMkpVNb(cfg)
	if err != nil {
		t.Fatal(err)
	}
	p := panel.(*Iw4id8kfTTT)
	data, err := p.GetAPIData()
	if err != nil {
		t.Fatal(err)
	}
	if p.GetAPIConfig() != cfg || p.GetSubmitInterval() != 3*time.Second {
		t.Fatal("configuration/interval")
	}
	if data.H47AzZsYHB1N.SwxJNgPpg.CKsOpayd != 8843 || data.H47AzZsYHB1N.SwxJNgPpg.ChJnngzb != "/ws" || !data.H47AzZsYHB1N.HSA2b5V {
		t.Fatal(data.H47AzZsYHB1N)
	}
	// Independent published MD5 fixture: a@example.com -> b418773a2c51fb9777a1648346fa7394.
	if len(data.VTme1SuHL6i_) != 2 || data.VTme1SuHL6i_[0].WZ89pTAOrHK4 != 0xb418773a {
		t.Fatal(data.VTme1SuHL6i_)
	}
	if credential := data.VTme1SuHL6i_[0].JAfXWwzYrW.(*L.Yn8BJoL); credential.SfiEGpJo != "uuid-a" || credential.FvQ2ltKOM7 != 9 {
		t.Fatal(credential)
	}
	if len(data.HYnfYWa) != 3 || data.HYnfYWa[2].GetRule() != "full:exact.example" || data.HYnfYWa[0].GetId() != 0 {
		t.Fatal(data.HYnfYWa)
	}
	err = p.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{
		{Hy7Q97RROK: 0xb418773a, XJkGTQf_qkAt: 1, I6LXObODav: 2},
		{Hy7Q97RROK: 0x1b418773a, XJkGTQf_qkAt: 3, I6LXObODav: 4},
		{Hy7Q97RROK: 42, XJkGTQf_qkAt: 5, I6LXObODav: 6},
	}})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]map[string]int64{"a@example.com": {"upload": 3, "download": 4}}
	if !reflect.DeepEqual(traffic, want) {
		t.Fatalf("traffic %#v", traffic)
	}
	if len(paths) != 2 {
		t.Fatal(paths)
	}
}
func TestDJStatusPayloadAndHTTPFailure(t *testing.T) {
	var status map[string]json.RawMessage
	code := 200
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.URL.Path != "/sub_server/submit_info" || r.PostForm.Get("secret") != "secret" {
			t.Error(r.URL, r.PostForm)
		}
		if e := json.Unmarshal([]byte(r.PostForm.Get("status")), &status); e != nil {
			t.Error(e)
		}
		w.WriteHeader(code)
	}))
	defer s.Close()
	p := &Iw4id8kfTTT{jamJA0Qs: &C.FCEAfQJwDA7e{W_ogGtF: s.URL + "/", Bd08Li: "secret"}}
	info := &DpeSn3T{Ext49lvSw: 99, CUcRQ2: 12.5, Z9XlEaJL: 1, Iqf1PJb9gAuJ: 2, ND2QesObp: 3, ZoKBTLsJbb: 4, MyXxXp12C: 5, BMQhQ6It6ocB: 6, SOaNHKQg3a: 7, ZiFazv7: 8, XR83N6vHA: 9, DrY4mmyOP: 10, IMkIJFzv: 11, Rl5I6SPhu50u: 12, XixSqS4zM6: 13, OEcQalt: 14, PaNuEBV208: 15}
	if e := p.SubmitInfo(info); e != nil {
		t.Fatal(e)
	}
	want := map[string]string{"v2": `{"code":0}`, "cpu": `{"percent":12.5}`, "uptime": "99", "memory": `{"total":2,"used":1}`, "swap": `{"total":4,"used":3}`, "disk": `{"total":6,"used":5}`, "loads": "[7,8,9]", "net_io": `{"down":11,"up":10}`, "net_traffic": `{"recv":13,"sent":12}`, "tcp_count": "14", "udp_count": "15"}
	for k, v := range want {
		if string(status[k]) != v {
			t.Errorf("%s = %s want %s", k, status[k], v)
		}
	}
	if len(status) != len(want) {
		t.Fatal(status)
	}
	code = 204
	if e := p.SubmitInfo(info); e == nil || e.Error() != "submit status failed: <nil>, status code: 204" {
		t.Fatal(e)
	}
	info.CUcRQ2 = math.NaN()
	if e := p.SubmitInfo(info); e == nil || !strings.Contains(e.Error(), "unsupported value: NaN") {
		t.Fatal(e)
	}
}
func TestDJOriginalBoundsAndMapLifetime(t *testing.T) {
	old := emailByIDAt5095E48
	t.Cleanup(func() { emailByIDAt5095E48 = old })
	emailByIDAt5095E48 = nil
	mustPanic := func(f func()) {
		t.Helper()
		defer func() {
			if recover() == nil {
				t.Error("expected original panic")
			}
		}()
		f()
	}
	mustPanic(func() { (&A_JUhZ{}).GetUser() })
	mustPanic(func() { (&E1rI_uHRR1{}).ToPanelNode() })
	mustPanic(func() { (&Iw4id8kfTTT{}).SubmitData(nil) })
	emailByIDAt5095E48 = map[uint32]string{7: "prior"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "broken") }))
	defer server.Close()
	p := &Iw4id8kfTTT{jamJA0Qs: &C.FCEAfQJwDA7e{W_ogGtF: server.URL + "/"}}
	if _, e := p.GetAPIData(); e == nil || !strings.HasPrefix(e.Error(), "parse node failed: ") {
		t.Fatal(e)
	}
	if emailByIDAt5095E48[7] != "prior" {
		t.Fatal("parse failure reset map")
	}
	if unsafe.Sizeof(DpeSn3T{}) != 160 || unsafe.Sizeof(Iw4id8kfTTT{}) != 16 || unsafe.Sizeof(A_JUhZ{}) != 40 {
		t.Fatal("RTTI sizes")
	}
}
