package cMTGNnWy

import (
	"encoding/base64"
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

func TestFiveProtocolHTTPPullCacheAndSubmit(t *testing.T) {
	old := stateAt3787334
	t.Cleanup(func() { stateAt3787334 = old })
	for _, protocol := range []C.Mpn2JltM{"ss", "v2ray", "trojan", "vless", "hysteria"} {
		t.Run(string(protocol), func(t *testing.T) {
			stateAt3787334 = 1
			var mu sync.Mutex
			gets, posts := 0, 0
			getGate, postGate := make(chan struct{}), make(chan struct{})
			received := map[string]json.RawMessage{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				action := strings.TrimPrefix(r.URL.Path, "/v1/server/")
				wantProtocol := map[C.Mpn2JltM]string{"ss": "shadowsocks", "v2ray": "vmess", "trojan": "trojan", "vless": "vless", "hysteria": "hysteria"}[protocol]
				q := r.URL.Query()
				if q.Get("protocol") != wantProtocol || q.Get("server_id") != "19" || q.Get("secret_key") != "s +/&" {
					t.Error(q)
				}
				if action == "config" || action == "user" {
					if r.Method != "GET" {
						t.Error(r.Method)
					}
					mu.Lock()
					gets++
					if gets == 2 {
						close(getGate)
					}
					mu.Unlock()
					select {
					case <-getGate:
					case <-time.After(3 * time.Second):
						t.Error("pull requests not parallel")
					}
					if r.Header.Get("If-None-Match") == "fixture" {
						w.WriteHeader(304)
						return
					}
					w.Header().Set("Etag", "fixture")
					if action == "user" {
						fmt.Fprint(w, `{"users":[{"id":21,"uuid":"0123456789abcdef0123456789abcdef","speed_limit":12,"device_limit":3}]}`)
						return
					}
					nodeProtocol := wantProtocol
					if protocol == "hysteria" {
						nodeProtocol = "hysteria2"
					}
					fmt.Fprintf(w, `{"basic":{"pull_interval":13,"push_interval":17},"protocol":%q,"config":{"port":8443,"security":"tls","transport":"websocket","transport_config":{"path":"/stream","service_name":"svc"},"method":"2022-blake3-aes-128-gcm","server_key":"node-key","obfs_password":"mask"}}`, nodeProtocol)
					return
				}
				if action != "online" && action != "push" && action != "status" {
					t.Error(action)
					w.WriteHeader(404)
					return
				}
				if r.Method != "POST" {
					t.Error(r.Method)
				}
				var body json.RawMessage
				if e := json.NewDecoder(r.Body).Decode(&body); e != nil {
					t.Error(e)
				}
				mu.Lock()
				received[action] = body
				posts++
				if posts == 3 {
					close(postGate)
				}
				mu.Unlock()
				select {
				case <-postGate:
				case <-time.After(3 * time.Second):
					t.Error("submission requests not parallel")
				}
				fmt.Fprint(w, "ok")
			}))
			defer server.Close()
			cfg := &C.FCEAfQJwDA7e{OS91X026uML: "ppanel", ZyAhkq5th8bf: protocol, F3PCiaIra: "webapi", W_ogGtF: server.URL + "/", Bd08Li: "s +/&", Yf1KqSqJpNI7: 19}
			p := WF3vigUr(cfg).(*Zkt5w_2o)
			data, err := p.GetAPIData()
			if err != nil {
				t.Fatal(err)
			}
			if p.GetAPIConfig() != cfg || data.H47AzZsYHB1N.EMlXEc != 19 || data.Xdo1D9 != 13*time.Second || data.DDu2Nm_P != 17*time.Second {
				t.Fatalf("%+v", data)
			}
			if len(data.VTme1SuHL6i_) != 1 || data.VTme1SuHL6i_[0].WZ89pTAOrHK4 != 21 || data.VTme1SuHL6i_[0].LmOKwjha != 12 {
				t.Fatal(data.VTme1SuHL6i_)
			}
			if protocol == "ss" {
				cred := data.VTme1SuHL6i_[0].JAfXWwzYrW.(*L.YlfOV0DBBnx)
				if cred.T1oDDIgYJHh != "2022-blake3-aes-128-gcm" || cred.A_wNUTkMh6pH != base64.StdEncoding.EncodeToString([]byte("0123456789abcdef")) {
					t.Fatal(cred)
				}
			}
			cached, e := p.GetAPIData()
			if e != nil || cached.H47AzZsYHB1N.Z1TIgHOQ != data.H47AzZsYHB1N.Z1TIgHOQ {
				t.Fatal(e)
			}
			before := time.Now().UnixMilli()
			err = p.SubmitData(&Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{{Hy7Q97RROK: 21, XJkGTQf_qkAt: 31, I6LXObODav: 41}}, JNcHHwTZ7: L.CQhZrJ4dK{{Vy9ybu: 21, PNtn9KAV: []string{"1.2.3.4", "::1"}}}})
			if err != nil {
				t.Fatal(err)
			}
			if string(received["push"]) != `{"traffic":[{"uid":21,"upload":31,"download":41}]}` || string(received["online"]) != `{"users":[{"uid":21,"ip":"1.2.3.4"},{"uid":21,"ip":"::1"}]}` {
				t.Fatal(string(received["push"]), string(received["online"]))
			}
			var status map[string]float64
			if e := json.Unmarshal(received["status"], &status); e != nil {
				t.Fatal(e)
			}
			if len(status) != 4 || status["updated_at"] < float64(before) || status["updated_at"] > float64(time.Now().UnixMilli()) {
				t.Fatal(status)
			}
			if err = p.SubmitData(nil); err != nil {
				t.Fatal(err)
			}
		})
	}
}
func TestCacheSameTagResponseAndErrorBranches(t *testing.T) {
	status := 200
	body := "first"
	etag := "tag"
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Etag", etag)
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}))
	defer s.Close()
	p := WF3vigUr(&C.FCEAfQJwDA7e{W_ogGtF: s.URL + "/"}).(*Zkt5w_2o)
	first, e := p.aCLea6Z("config")
	if e != nil {
		t.Fatal(e)
	}
	body = "changed"
	same, e := p.aCLea6Z("config")
	if e != nil || string(same) != "first" || &same[0] != &first[0] {
		t.Fatal(string(same), e)
	}
	status = 304
	cached, e := p.aCLea6Z("config")
	if e != nil || &cached[0] != &first[0] {
		t.Fatal(e)
	}
	if _, e = p.aCLea6Z("user"); e == nil || e.Error() != "can not find eTag: "+s.URL+"/v1/server/user" {
		t.Fatal(e)
	}
	status = 500
	body = "bad"
	if _, e = p.aCLea6Z("user"); e == nil || e.Error() != "request failed, result: bad, status code: 500" {
		t.Fatal(e)
	}
	if e = p.wVyldOdHZ3gi("push", nil); e == nil || e.Error() != "statuc code: 500, result: bad" {
		t.Fatal(e)
	}
	status = 200
	etag = ""
	body = "fresh"
	fresh, e := p.aCLea6Z("config")
	if e != nil || string(fresh) != "fresh" {
		t.Fatal(string(fresh), e)
	}
	if tag := p.oDI12LC.GetETag(p.mTbgoUSgl("config")); tag != "tag" {
		t.Fatal("untagged result removed original cache")
	}
	p.oDI12LC.AddETag(p.mTbgoUSgl("user"), "tag", 42)
	status = 304
	func() {
		defer func() {
			if recover() == nil {
				t.Error("wrong cached type must panic")
			}
		}()
		p.aCLea6Z("user")
	}()
	if got := p.qOheTtVSGU("missing"); got != nil {
		t.Fatal(got)
	}
	if !reflect.DeepEqual(p.qOheTtVSGU(p.mTbgoUSgl("config")), http.Header{"IF-NONE-MATCH": []string{"tag"}}) {
		t.Fatal("header spelling")
	}
}
