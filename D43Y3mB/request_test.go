package D43Y3mB

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	L "recovered.local/soga/LFWTTXF7JR"
	C "recovered.local/soga/cNT24rG4F"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestOriginalHTTPStatusAndResultBranches(t *testing.T) {
	for _, tc := range []struct {
		name, body, want string
		status           int
	}{
		{"success", `{"code":200,"message":"ok"}`, "", 200},
		{"application-error", `{"code":403,"message":"denied"}`, `result: {"code":403,"message":"denied"}`, 200},
		{"http-error", `down`, `request failed, result: down, status code: 503`, 503},
		{"bad-json", `{`, `result: {, status code: 200`, 200},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" || r.Header.Get("Key") != "test-secret" || r.Header.Get("Content-Type") != "application/json" || r.URL.RawQuery != "" {
					t.Errorf("request %s %s %+v", r.Method, r.URL, r.Header)
				}
				stamp, e := strconv.ParseUint(r.Header.Get("Timestamp"), 10, 32)
				if e != nil || stamp == 0 {
					t.Error("timestamp", stamp, e)
				}
				raw, _ := io.ReadAll(r.Body)
				if string(raw) != `{"n":7}` {
					t.Error(string(raw))
				}
				w.WriteHeader(tc.status)
				io.WriteString(w, tc.body)
			}))
			defer s.Close()
			p := GMLWI3{TMfNUn: &C.FCEAfQJwDA7e{Bd08Li: "test-secret"}}
			err := p.PostCommon(s.URL, map[string]int{"n": 7})
			if tc.want == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("%v want %s", err, tc.want)
			}
		})
	}
	p := GMLWI3{TMfNUn: &C.FCEAfQJwDA7e{}}
	if err := p.PostCommon(":bad", map[string]int{}); err == nil || !strings.Contains(err.Error(), "result: , status code: 0") {
		t.Fatal(err)
	}
	if err := p.PostCommon("http://127.0.0.1:1", make(chan int)); err == nil || !strings.Contains(err.Error(), "unsupported type") {
		t.Fatal(err)
	}
}

func TestOriginalGetIgnoresHTTPStatus(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(503); io.WriteString(w, "original-body") }))
	defer s.Close()
	p := GMLWI3{TMfNUn: &C.FCEAfQJwDA7e{}}
	b, e := p.GetCommon(s.URL)
	if e != nil || string(b) != "original-body" {
		t.Fatalf("%s %v", b, e)
	}
}
func TestOriginalSubmissionWire(t *testing.T) {
	type capture struct {
		path string
		raw  string
	}
	requests := make(chan capture, 8)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		requests <- capture{r.URL.Path, string(b)}
		io.WriteString(w, `{"code":200}`)
	}))
	defer s.Close()
	p := GMLWI3{TMfNUn: &C.FCEAfQJwDA7e{}}
	if err := p.SubmitTraffic(s.URL+"/traffic", L.CWfcUv{{Hy7Q97RROK: 7, XJkGTQf_qkAt: 123, I6LXObODav: 456}}); err != nil {
		t.Fatal(err)
	}
	if err := p.SubmitAliveIP(s.URL+"/online", L.CQhZrJ4dK{{Vy9ybu: 7, PNtn9KAV: []string{"::1", "127.0.0.1"}}}); err != nil {
		t.Fatal(err)
	}
	if err := p.SubmitDetectLog(s.URL+"/trigger", L.QmiwDVjV00lZ{{GbzQUT_4f: 7, CAgUUdrv9: 19}, {GbzQUT_4f: 8, CAgUUdrv9: 20}}); err != nil {
		t.Fatal(err)
	}
	expect := map[string]int{`/traffic [{"uid":7,"upload":123,"download":456}]`: 1, `/online [{"uid":7,"ip":"::1,127.0.0.1"}]`: 1, `/trigger {"uid":7,"rule_id":19,"reason":"无"}`: 1, `/trigger {"uid":8,"rule_id":20,"reason":"无"}`: 1}
	for range 4 {
		select {
		case r := <-requests:
			key := r.path + " " + r.raw
			if expect[key] != 1 {
				t.Errorf("unexpected %s", key)
			}
			delete(expect, key)
		case <-time.After(3 * time.Second):
			t.Fatal("missing asynchronous request")
		}
	}
	if len(expect) != 0 {
		t.Fatal(expect)
	}
	p.SubmitTraffic(s.URL, nil)
	p.SubmitAliveIP(s.URL, nil)
	p.SubmitDetectLog(s.URL, nil)
	select {
	case r := <-requests:
		t.Fatalf("empty submission made request %+v", r)
	case <-time.After(20 * time.Millisecond):
	}
}
func TestOriginalDetectorBranches(t *testing.T) {
	var p Y4Pw3Ojd6HO
	json.Unmarshal([]byte(`{"code":200,"data":{"mode":"ignored","rules":[{"id":1,"type":"reg","pattern":"^abc$"},{"id":2,"type":"domain","pattern":"example.test"},{"id":3,"type":"unknown","pattern":"x"}]}}`), &p)
	d := p.ToDetectors()
	if len(d) != 2 || d[0].GetId() != 1 || d[0].GetRule() != "regexp:^abc$" || !d[0].Test("abc") || !d[1].Test("www.example.test") {
		t.Fatal(d)
	}
	if e := func() error { _, e := (&Lv22mM8eYq{USQvAO4WtdOX: "unknown"}).ToDetector(); return e }(); e == nil || e.Error() != "unknown detect rule type: unknown" {
		t.Fatal(e)
	}
	for _, body := range []string{`{"code":401}`, `{`} {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, body) }))
		v, e := (&GMLWI3{TMfNUn: &C.FCEAfQJwDA7e{}}).GetDetectors(s.URL)
		s.Close()
		if v != nil || e == nil {
			t.Fatal(v, e)
		}
	}
}
func TestNodeStatusRealLocalMetricsWire(t *testing.T) {
	var got map[string]any
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if e := json.NewDecoder(r.Body).Decode(&got); e != nil {
			t.Error(e)
		}
		io.WriteString(w, `{"code":200}`)
	}))
	defer s.Close()
	p := GMLWI3{TMfNUn: &C.FCEAfQJwDA7e{}}
	if e := p.SubmitNodeInfo(s.URL); e != nil {
		t.Fatal(e)
	}
	if len(got) != 5 {
		t.Fatal(got)
	}
	for _, key := range []string{"cpu", "mem", "disk"} {
		v, ok := got[key].(string)
		if !ok || !strings.HasSuffix(v, "%") {
			t.Fatal(key, got)
		}
	}
	if v, ok := got["net"].(string); !ok || !strings.Contains(v, "GB↑ ") || !strings.HasSuffix(v, "GB↓") {
		t.Fatal(got)
	}
	if _, ok := got["uptime"].(float64); !ok {
		t.Fatal(got)
	}
}
