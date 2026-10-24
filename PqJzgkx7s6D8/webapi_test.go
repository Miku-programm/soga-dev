package PqJzgkx7s6D8

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/load"
	"io"
	"net/http"
	"net/http/httptest"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	Q "recovered.local/soga/q5bqCgTdxaTM"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestCommonETagAndStatusSemantics(t *testing.T) {
	status, body, etag := 500, "first", "tag"
	hits := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		if r.URL.Path != "/mod_mu/test" || r.URL.Query().Get("key") != "secret" || r.URL.Query().Get("node_id") != "7" {
			t.Error(r.URL)
		}
		if hits > 1 && r.Header.Get("If-None-Match") != "tag" {
			t.Error(r.Header)
		}
		if etag != "" {
			w.Header().Set("ETag", etag)
		}
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}))
	defer srv.Close()
	p := &JYkPXk4ZdW{EA_5pkz: &C.FCEAfQJwDA7e{W_ogGtF: srv.URL + "/", Bd08Li: "secret", Yf1KqSqJpNI7: 7}}
	b, e := p.GetCommon("test")
	if e != nil || string(b) != "first" {
		t.Fatal(string(b), e)
	}
	status, body = 200, "second"
	b, e = p.GetCommon("test")
	if e != nil || string(b) != "first" {
		t.Fatal(string(b), e)
	}
	status = 304
	b, e = p.GetCommon("test")
	if e != nil || string(b) != "first" {
		t.Fatal(string(b), e)
	}
	fresh := &JYkPXk4ZdW{EA_5pkz: p.EA_5pkz}
	hits = 0
	_, e = fresh.GetCommon("test")
	if e == nil || !strings.Contains(e.Error(), "can not find eTag:") {
		t.Fatal(e)
	}
}
func TestPostValidationAndStatusSuppression(t *testing.T) {
	status, body := 200, `{"ret":1}`
	var payload string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		payload = string(b)
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}))
	defer srv.Close()
	p := &JYkPXk4ZdW{EA_5pkz: &C.FCEAfQJwDA7e{W_ogGtF: srv.URL + "/"}}
	if e := p.PostCommon("x", map[string]int{"x": 2}); e != nil || payload != `{"x":2}` {
		t.Fatal(e, payload)
	}
	body = `{"ret":0}`
	if e := p.PostCommon("x", nil); e == nil || e.Error() != `result: {"ret":0}` {
		t.Fatal(e)
	}
	body = `[bad`
	if e := p.PostCommon("x", nil); e == nil {
		t.Fatal("invalid JSON accepted")
	}
	status = 503
	body = "unavailable"
	if e := p.PostCommon("x", nil); e == nil || e.Error() != "request failed, result: unavailable, status code: 503" {
		t.Fatal(e)
	}
	oldLoad, oldUp := hostLoad, hostUptime
	defer func() { hostLoad, hostUptime = oldLoad, oldUp }()
	hostLoad = func() (*load.AvgStat, error) { return nil, errors.New("load unavailable") }
	hostUptime = func() (uint64, error) { return 37, errors.New("uptime partial") }
	if e := p.SubmitNodeInfo(); e != nil {
		t.Fatal("original discards POST error", e)
	}
	if payload != `{"load":"0.00 0.00 0.00","uptime":37}` {
		t.Fatal(payload)
	}
}
func TestSubmitFourTasksOverlapAndJSON(t *testing.T) {
	var mu sync.Mutex
	got := map[string]string{}
	arrived := make(chan string, 4)
	release := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		mu.Lock()
		got[r.URL.Path] = string(b)
		mu.Unlock()
		arrived <- r.URL.Path
		<-release
		fmt.Fprint(w, `{"ret":1}`)
	}))
	defer srv.Close()
	p := &JYkPXk4ZdW{EA_5pkz: &C.FCEAfQJwDA7e{W_ogGtF: srv.URL + "/", Yf1KqSqJpNI7: 4}}
	data := &Q.FGmNba4MhioY{HTgPyE3VytZd: L.CWfcUv{&L.DLxslMwYJXQw{Hy7Q97RROK: 3, XJkGTQf_qkAt: 8, I6LXObODav: 9}}, JNcHHwTZ7: L.CQhZrJ4dK{&L.LmNKoY94{Vy9ybu: 3, PNtn9KAV: []string{"1.2.3.4", "::1"}}}, Y0ys4ma4Tno4: L.QmiwDVjV00lZ{&L.B9FS114R4F9h{GbzQUT_4f: 3, CAgUUdrv9: 6}}}
	done := make(chan error, 1)
	go func() { done <- p.SubmitData(data) }()
	for i := 0; i < 4; i++ {
		select {
		case <-arrived:
		case <-time.After(3 * time.Second):
			close(release)
			t.Fatal("tasks did not overlap")
		}
	}
	close(release)
	if e := <-done; e != nil {
		t.Fatal(e)
	}
	want := map[string]string{"/mod_mu/users/traffic": `{"data":[{"user_id":3,"u":8,"d":9}]}`, "/mod_mu/users/aliveip": `{"data":[{"ip":"1.2.3.4","user_id":3},{"ip":"::1","user_id":3}]}`, "/mod_mu/users/detectlog": `{"data":[{"list_id":6,"user_id":3}]}`}
	for path, b := range want {
		if got[path] != b {
			t.Errorf("%s: %s", path, got[path])
		}
	}
}
