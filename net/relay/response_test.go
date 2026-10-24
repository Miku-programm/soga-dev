package relay

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// 测试替换客户端只存在于测试文件；先执行原初始化，再在结束时恢复指针。
func useHTTPClient(t *testing.T, client *http.Client) {
	t.Helper()
	clientOnceAt50C7B70.Do(WuvamjL5)
	old := clientAt5029428
	clientAt5029428 = client
	t.Cleanup(func() { clientAt5029428 = old })
}
func TestHTTPWrappersWireBehavior(t *testing.T) {
	type captured struct {
		method, query, body, contentType string
		ua, cache                        []string
	}
	requests := make(chan captured, 3)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		requests <- captured{r.Method, r.URL.RawQuery, string(b), r.Header.Get("Content-Type"), r.Header.Values("User-Agent"), r.Header.Values("Cache-Control")}
		w.Header().Set("ETag", "version-2")
		w.WriteHeader(418)
		_, _ = io.WriteString(w, "response")
	}))
	defer server.Close()
	useHTTPClient(t, server.Client())
	code, data, h, err := IlWBbp1xdPw(server.URL+"?old=1", url.Values{"a": {"x y"}}, http.Header{"X-Test": {"one", "two"}})
	if err != nil || code != 418 || string(data) != "response" || h.Get("ETag") != "version-2" {
		t.Fatalf("GET: %d %q %v %v", code, data, h, err)
	}
	got := <-requests
	if got.method != "GET" || got.query != "a=x+y" || got.body != "" || got.contentType != "" {
		t.Fatalf("GET请求: %+v", got)
	}
	if !reflect.DeepEqual(got.ua, []string{qword_5097990}) || !reflect.DeepEqual(got.cache, []string{"no-cache"}) {
		t.Fatalf("默认头: %+v", got)
	}
	data, code, err = ZBovtj1xky(server.URL+"?old=1", url.Values{}, map[string]string{"z": "<tag>", "a": "value"}, http.Header{"Content-Type": {"wrong", "also-wrong"}})
	if err != nil || code != 418 || string(data) != "response" {
		t.Fatalf("JSON: %d %q %v", code, data, err)
	}
	got = <-requests
	if got.method != "POST" || got.query != "" || got.contentType != "application/json" || got.body != `{"a":"value","z":"\u003ctag\u003e"}` {
		t.Fatalf("JSON请求: %+v", got)
	}
	data, code, err = J1liFniHg(server.URL+"?old=1", url.Values{"z": {"x y", "+"}, "a": {"v"}})
	if err != nil || code != 418 || string(data) != "response" {
		t.Fatalf("form: %d %q %v", code, data, err)
	}
	got = <-requests
	if got.method != "POST" || got.query != "old=1" || got.body != "a=v&z=x+y&z=%2B" || got.contentType != "application/x-www-form-urlencoded" {
		t.Fatalf("表单请求: %+v", got)
	}
}

type testRoundTripper func(*http.Request) (*http.Response, error)

func (f testRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type failingResponseBody struct {
	closed bool
	err    error
}

func (b *failingResponseBody) Read(p []byte) (int, error) { return copy(p, "partial"), b.err }
func (b *failingResponseBody) Close() error {
	b.closed = true
	return errors.New("ignored close error")
}
func TestHTTPReadErrorDiscardsPartialBodyAndCloses(t *testing.T) {
	readErr := errors.New("response read failed")
	for _, kind := range []string{"GET", "JSON", "form"} {
		t.Run(kind, func(t *testing.T) {
			body := &failingResponseBody{err: readErr}
			useHTTPClient(t, &http.Client{Transport: testRoundTripper(func(*http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 206, Header: http.Header{"Etag": {"kept"}}, Body: body}, nil
			})})
			var data []byte
			var code int
			var err error
			switch kind {
			case "GET":
				var h http.Header
				code, data, h, err = IlWBbp1xdPw("http://example.invalid", nil, nil)
				if h.Get("ETag") != "kept" {
					t.Fatal("丢失响应头")
				}
			case "JSON":
				data, code, err = ZBovtj1xky("http://example.invalid", nil, nil, nil)
			case "form":
				data, code, err = J1liFniHg("http://example.invalid", nil)
			}
			if data != nil || code != 206 || err != readErr || !body.closed {
				t.Fatalf("data=%q code=%d err=%v closed=%v", data, code, err, body.closed)
			}
		})
	}
}
func TestHTTPConstructionAndMarshalErrors(t *testing.T) {
	calls := 0
	useHTTPClient(t, &http.Client{Transport: testRoundTripper(func(*http.Request) (*http.Response, error) { calls++; return nil, errors.New("transport failed") })})
	code, data, h, err := IlWBbp1xdPw("://bad", nil, nil)
	if code != 0 || data != nil || h != nil || err == nil || calls != 0 {
		t.Fatal("错误URL应在发请求前失败")
	}
	data, code, err = ZBovtj1xky("http://example.invalid", nil, make(chan int), nil)
	if code != 0 || data != nil || err == nil || calls != 0 {
		t.Fatal("JSON编码失败不应发请求")
	}
	code, data, h, err = IlWBbp1xdPw("http://example.invalid", nil, nil)
	if code != 0 || data != nil || h != nil || err == nil || calls != 1 || !strings.Contains(err.Error(), "transport failed") {
		t.Fatal("传输错误未保留")
	}
}
