package relay

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestRequestQueryReplacement(t *testing.T) {
	for _, tc := range []struct {
		name  string
		query url.Values
		want  string
	}{
		{"nil保留", nil, "old=1"},
		{"空映射清除", url.Values{}, ""},
		{"编码并替换", url.Values{"z": {"汉 字", "+"}, "a": {"2"}}, "a=2&z=%E6%B1%89+%E5%AD%97&z=%2B"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req, err := dfK5RMSe4h9B("http://example.invalid/x?old=1#fragment", "GET", tc.query, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			if req.URL.RawQuery != tc.want || req.URL.Fragment != "fragment" {
				t.Fatalf("URL = %v", req.URL)
			}
			if req.Context() != context.Background() {
				t.Fatal("请求上下文不匹配")
			}
		})
	}
}
func TestRequestHeadersAppendAndBody(t *testing.T) {
	headers := http.Header{"user-agent": {"custom"}, "CACHE-CONTROL": {"private", "max-age=0"}, "x-extra": {"a", "b"}, "x-empty": nil}
	req, err := dfK5RMSe4h9B("http://example.invalid", "POST", nil, bytes.NewBufferString("payload"), headers)
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{qword_5097990, "custom"}; !reflect.DeepEqual(req.Header.Values("User-Agent"), want) {
		t.Fatalf("UA: %v", req.Header)
	}
	if want := []string{"no-cache", "private", "max-age=0"}; !reflect.DeepEqual(req.Header.Values("Cache-Control"), want) {
		t.Fatalf("Cache-Control: %v", req.Header)
	}
	if !reflect.DeepEqual(req.Header.Values("X-Extra"), []string{"a", "b"}) {
		t.Fatalf("headers: %v", req.Header)
	}
	if _, ok := req.Header["X-Empty"]; ok {
		t.Fatal("空切片不应生成头字段")
	}
	if req.ContentLength != 7 || req.GetBody == nil {
		t.Fatal("bytes.Buffer 的请求元数据不匹配")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil || string(data) != "payload" {
		t.Fatalf("body=%q err=%v", data, err)
	}
	req.Header.Add("X-Extra", "c")
	if len(headers["x-extra"]) != 2 {
		t.Fatal("修改了调用方头字段切片")
	}
}
func TestRequestRejectsInvalidURLAndMethod(t *testing.T) {
	for _, tc := range []struct{ address, method string }{{"://broken", "GET"}, {"http://example.invalid", "bad method"}} {
		req, err := dfK5RMSe4h9B(tc.address, tc.method, nil, nil, nil)
		if req != nil || err == nil {
			t.Fatalf("req=%v err=%v", req, err)
		}
	}
}
