package ZrjfjK8Y_tM9

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"recovered.local/soga/cNT24rG4F"
	"recovered.local/soga/wd42zQ0FT"
	"strings"
	"testing"
)

func panelForRequestTests() *JZA3kKksUe {
	return &JZA3kKksUe{EA_5pkz: &cNT24rG4F.FCEAfQJwDA7e{Bd08Li: "secret", Yf1KqSqJpNI7: 7, ZyAhkq5th8bf: "ss"}, eyq32s: &wd42zQ0FT.NG0hCZyRHK{}}
}
func TestGetCommonETagAndParameters(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if got := r.URL.Query()["token"]; len(got) != 2 || got[0] != "secret" || got[1] != "extra" {
			t.Errorf("token = %v", got)
		}
		if r.URL.Query().Get("node_id") != "7" || r.URL.Query().Get("node_type") != "shadowsocks" || r.URL.Query().Has("old") {
			t.Errorf("query=%v", r.URL.Query())
		}
		if calls > 1 && r.Header.Get("If-None-Match") != "v1" {
			t.Errorf("ETag头=%v", r.Header)
		}
		if calls == 3 {
			w.WriteHeader(304)
			return
		}
		w.Header().Set("ETag", "v1")
		if calls == 1 {
			_, _ = w.Write([]byte("first"))
		} else {
			_, _ = w.Write([]byte("changed"))
		}
	}))
	defer server.Close()
	p := panelForRequestTests()
	address := server.URL + "?old=1"
	extra := url.Values{"token": {"extra"}}
	for i := 0; i < 3; i++ {
		data, err := p.GetCommon(address, extra)
		if err != nil || string(data) != "first" {
			t.Fatalf("call %d: %q %v", i, data, err)
		}
	}
	if len(extra["token"]) != 1 {
		t.Fatal("修改了额外参数")
	}
}
func TestGetCommonFailureAndNoTag(t *testing.T) {
	status := 200
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(status); _, _ = w.Write([]byte("body")) }))
	defer server.Close()
	p := panelForRequestTests()
	data, err := p.GetCommon(server.URL, nil)
	if err != nil || string(data) != "body" {
		t.Fatalf("%q %v", data, err)
	}
	if _, ok := p.eyq32s.GetObj(server.URL); ok {
		t.Fatal("无ETag不缓存")
	}
	status = 304
	data, err = p.GetCommon(server.URL, nil)
	if data != nil || err == nil || err.Error() != "can not find eTag: "+server.URL {
		t.Fatalf("304: %q %v", data, err)
	}
	status = 503
	data, err = p.GetCommon(server.URL, nil)
	if data != nil || err == nil || err.Error() != "request failed, result: body, status code: 503" {
		t.Fatalf("503: %q %v", data, err)
	}
	data, err = p.GetCommon("://bad", nil)
	if data != nil || err == nil || !strings.HasSuffix(err.Error(), ", result: , status code: 0") {
		t.Fatalf("传输错误: %q %v", data, err)
	}
}
func TestGetCommonRetainsCacheTypeAssertion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(304) }))
	defer server.Close()
	p := panelForRequestTests()
	p.eyq32s.AddETag(server.URL, "v", "wrong type")
	defer func() {
		if recover() == nil {
			t.Fatal("缓存类型错误应panic")
		}
	}()
	_, _ = p.GetCommon(server.URL, nil)
}
func TestPostCommonStatusAndEncodingErrors(t *testing.T) {
	status := 200
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Query().Get("token") != "secret" {
			t.Errorf("request=%v", r)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte("body"))
	}))
	defer server.Close()
	p := panelForRequestTests()
	if err := p.PostCommon(server.URL, nil); err != nil {
		t.Fatal(err)
	}
	status = 201
	if err := p.PostCommon(server.URL, nil); err == nil || err.Error() != "post status code: 201, result: body" {
		t.Fatalf("201: %v", err)
	}
	if err := p.PostCommon(server.URL, make(chan int)); err == nil || !strings.HasPrefix(err.Error(), "post "+server.URL+" err: json: unsupported type: chan int") || !strings.HasSuffix(err.Error(), ", status code: 0, result: ") {
		t.Fatalf("marshal: %v", err)
	}
}
