package update

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"
)

// 仅测试连接到真实共享 HTTP 客户端变量，以拒绝所有非本地请求；生产没有可替换 hooks。
//
//go:linkname originalHTTPClient recovered.local/soga/net/relay.clientAt5029428
var originalHTTPClient *http.Client

//go:linkname originalHTTPOnce recovered.local/soga/net/relay.clientOnceAt50C7B70
var originalHTTPOnce sync.Once

type localTransport func(*http.Request) (*http.Response, error)

func (f localTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func isolatedHTTP(t *testing.T, round localTransport) {
	t.Helper()
	old := originalHTTPClient
	originalHTTPOnce.Do(func() {})
	originalHTTPClient = &http.Client{Transport: round}
	t.Cleanup(func() { originalHTTPClient = old })
}
func TestReleaseRTTIAndTime(t *testing.T) {
	var r e5s7yjcoNuB6
	typ := reflect.TypeOf(r)
	if unsafe.Sizeof(r) != 32 || unsafe.Offsetof(r.KecFp709H) != 16 || typ.Field(0).Tag.Get("json") != "tag_name" || typ.Field(1).Tag.Get("json") != "published_at" {
		t.Fatal("original RTTI")
	}
	for _, v := range []string{"", "2026-02-30T00:00:00Z", "2026-01-02", "garbage"} {
		r.KecFp709H = v
		if !r.PublishTime().IsZero() {
			t.Fatal(v)
		}
	}
	r.KecFp709H = "2026-10-24T03:20:00+08:00"
	if r.PublishTime().UTC().Format(time.RFC3339) != "2026-10-23T19:20:00Z" {
		t.Fatal(r.PublishTime())
	}
}
func TestReleaseRetryAndActualLocalHTTP(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/repos/vaxilu/soga/releases/latest" {
			t.Errorf("request %v", r)
		}
		w.WriteHeader(503) // 原实现不按 HTTP 状态判错。
		io.WriteString(w, `{"tag_name":"2.18.0","published_at":"2026-10-20T01:02:03Z"}`)
	}))
	defer server.Close()
	endpoint, _ := url.Parse(server.URL)
	isolatedHTTP(t, func(r *http.Request) (*http.Response, error) {
		calls++
		if r.URL.String() != "https://api.github.com/repos/vaxilu/soga/releases/latest" {
			t.Fatalf("unexpected URL %s", r.URL)
		}
		if calls < 3 {
			return nil, errors.New("local fixture connection refused")
		}
		copy := r.Clone(r.Context())
		copy.URL.Scheme = endpoint.Scheme
		copy.URL.Host = endpoint.Host
		return server.Client().Transport.RoundTrip(copy)
	})
	version, published, err := m_33LUULr2()
	if err != nil || calls != 3 || version != "2.18.0" || published.IsZero() {
		t.Fatalf("%s %s %d %v", version, published, calls, err)
	}
}
func TestReleaseMalformedDoesNotRetryAndEmptyBody(t *testing.T) {
	for _, body := range []string{"", "{", `{"tag_name":"2.18.0","published_at":"invalid"}`} {
		t.Run(body, func(t *testing.T) {
			calls := 0
			isolatedHTTP(t, func(r *http.Request) (*http.Response, error) {
				calls++
				return &http.Response{StatusCode: 200, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}, nil
			})
			version, published, err := m_33LUULr2()
			if calls != 1 {
				t.Fatal(calls)
			}
			if body == "" || body == "{" {
				if err == nil {
					t.Fatal("invalid JSON accepted")
				}
			} else if version != "2.18.0" || !published.IsZero() || err != nil {
				t.Fatal(version, published, err)
			}
		})
	}
	calls := 0
	isolatedHTTP(t, func(*http.Request) (*http.Response, error) { calls++; return nil, errors.New("denied fixture") })
	if v, tm, e := m_33LUULr2(); v != "" || !tm.IsZero() || e == nil || calls != 3 {
		t.Fatal(v, tm, e, calls)
	}
}

type downloadBody struct {
	reads, closed int
	firstError    error
	fail          bool
}

func (b *downloadBody) Read(p []byte) (int, error) {
	b.reads++
	if b.reads == 1 {
		return copy(p, "retained payload"), b.firstError
	}
	if b.fail {
		return 0, errors.New("final read failed")
	}
	return 0, io.EOF
}
func (b *downloadBody) Close() error { b.closed++; return errors.New("close ignored") }
func TestDownloadExactURLReadErrorOverwriteAndCleanup(t *testing.T) {
	t.Chdir(t.TempDir())
	for _, fail := range []bool{false, true} {
		t.Run(map[bool]string{false: "write_overwrites_read_error", true: "read_failure_removes_file"}[fail], func(t *testing.T) {
			body := &downloadBody{firstError: errors.New("partial read failure intentionally overwritten"), fail: fail}
			isolatedHTTP(t, func(r *http.Request) (*http.Response, error) {
				if r.Method != "GET" || r.URL.String() != "https://github.com/vaxilu/soga/releases/download/2.17.0/soga-linux-amd64.tar.gz" {
					t.Fatalf("request %s %s", r.Method, r.URL)
				}
				return &http.Response{StatusCode: 404, Header: make(http.Header), Body: body}, nil
			})
			name, e := bYg0ELm("2.17.0")
			if body.closed != 1 || body.reads != 2 {
				t.Fatal(body)
			}
			if fail {
				if name != "" || e == nil {
					t.Fatal(name, e)
				}
				if _, e = os.Stat("soga.tar.gz"); !os.IsNotExist(e) {
					t.Fatal(e)
				}
			} else {
				if name != "soga.tar.gz" || e != nil {
					t.Fatal(name, e)
				}
				data, e := os.ReadFile(name)
				if e != nil || string(data) != "retained payload" {
					t.Fatal(string(data), e)
				}
			}
		})
	}
}
func archiveFixture(t *testing.T, headers []*tar.Header, bodies []string) string {
	t.Helper()
	name := filepath.Join(t.TempDir(), "fixture.tar.gz")
	f, e := os.Create(name)
	if e != nil {
		t.Fatal(e)
	}
	gz := gzip.NewWriter(f)
	tw := tar.NewWriter(gz)
	for i, h := range headers {
		if e = tw.WriteHeader(h); e != nil {
			t.Fatal(e)
		}
		if _, e = io.WriteString(tw, bodies[i]); e != nil {
			t.Fatal(e)
		}
	}
	if e = tw.Close(); e != nil {
		t.Fatal(e)
	}
	if e = gz.Close(); e != nil {
		t.Fatal(e)
	}
	if e = f.Close(); e != nil {
		t.Fatal(e)
	}
	return name
}
func TestNativeArchiveExtractionAndErrorBoundaries(t *testing.T) {
	dir := t.TempDir()
	name := archiveFixture(t, []*tar.Header{{Name: "soga/", Typeflag: tar.TypeDir, Mode: 0700}, {Name: "soga/soga", Mode: 0700, Size: 7}, {Name: "soga/nested/config", Size: 3}, {Name: "soga/link", Typeflag: tar.TypeSymlink, Linkname: "soga"}}, []string{"", "program", "cfg", ""})
	if e := pUbOA95m8Gea(name, dir); e != nil {
		t.Fatal(e)
	}
	for file, want := range map[string]string{"soga/soga": "program", "soga/nested/config": "cfg", "soga/link": ""} {
		data, e := os.ReadFile(filepath.Join(dir, file))
		if e != nil || string(data) != want {
			t.Fatal(file, string(data), e)
		}
	}
	info, _ := os.Lstat(filepath.Join(dir, "soga/link"))
	if !info.Mode().IsRegular() {
		t.Fatal("original ignores symlink type")
	}
	bad := filepath.Join(t.TempDir(), "bad.gz")
	os.WriteFile(bad, []byte("not gzip"), 0600)
	if pUbOA95m8Gea(bad, dir) == nil || pUbOA95m8Gea(bad+"missing", dir) == nil {
		t.Fatal("missing/gzip error")
	}
	// 所有路径都在临时根内。原算法没有清理 ../，保留同一目标路径行为。
	escape := archiveFixture(t, []*tar.Header{{Name: "../sibling", Size: 1}}, []string{"x"})
	if e := pUbOA95m8Gea(escape, filepath.Join(dir, "child")); e != nil {
		t.Fatal(e)
	}
	if b, e := os.ReadFile(filepath.Join(dir, "sibling")); e != nil || string(b) != "x" {
		t.Fatal(string(b), e)
	}
}
func TestFilenameNativeRuneByteIndexBoundary(t *testing.T) {
	t.Chdir(t.TempDir())
	func() {
		defer func() {
			if recover() == nil {
				t.Error("original slashless path must panic")
			}
		}()
		_, _ = kT0645CUQ9PS("plain")
	}()
	// 字节下标3作用到rune序列 [甲 / x]，会创建包含NUL或错误尾部的父目录。
	// 这里只断言原路径不能像 filepath.Dir 一样成功，避免依赖编译器切片容量。
	if f, e := kT0645CUQ9PS("甲/x"); e == nil {
		f.Close()
		t.Fatal("native rune/byte offset behavior was silently fixed")
	}
}
func TestUpdateFailureAndCancelledStartNeverTouchRemoteOrInstallation(t *testing.T) {
	calls := 0
	isolatedHTTP(t, func(*http.Request) (*http.Response, error) {
		calls++
		return nil, errors.New("fixture rejects every request")
	})
	g0x1jn("2.17.0") // 下载之前失败，不能走到删除安装目录/执行 restart。
	if calls != 1 {
		t.Fatal(calls)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	ZAWxLZe5zfNN(ctx)
	time.Sleep(20 * time.Millisecond)
	if calls != 1 {
		t.Fatal("unexpected immediate update request")
	}
}
