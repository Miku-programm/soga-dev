package geo

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestGeoDownloadParallelWritesAndTimestamp(t *testing.T) {
	previous := updateTimeAt5095C48.Load()
	t.Cleanup(func() { updateTimeAt5095C48.Store(previous) })
	updateTimeAt5095C48.Store(nil)
	if at, ok := MCOm37H(time.Time{}); !at.IsZero() || ok {
		t.Fatal("uninitialized timestamp")
	}
	arrived := make(chan struct{}, 2)
	release := make(chan struct{})
	var once sync.Once
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		arrived <- struct{}{}
		<-release
		if r.URL.Path == "/ip" {
			w.WriteHeader(http.StatusNotFound)
		}
		io.WriteString(w, r.URL.Path+" payload")
	}))
	defer server.Close()
	go func() { <-arrived; <-arrived; once.Do(func() { close(release) }) }()
	dir := t.TempDir()
	site, ip := filepath.Join(dir, "site.dat"), filepath.Join(dir, "ip.dat")
	writer := func(path string) func(io.Reader) (bool, error) {
		return func(r io.Reader) (bool, error) { err := kBh3sTD66VJD(r, path); return err == nil, err }
	}
	before := time.Now()
	done := make(chan struct{})
	go func() {
		updateGeoFiles(server.Client(), server.URL+"/site", server.URL+"/ip", writer(site), writer(ip))
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		once.Do(func() { close(release) })
		t.Fatal("downloads did not run concurrently")
	}
	for path, want := range map[string]string{site: "/site payload", ip: "/ip payload"} {
		b, e := os.ReadFile(path)
		if e != nil || string(b) != want {
			t.Fatalf("%s: %q, %v", path, b, e)
		}
	}
	at, ok := MCOm37H(before)
	if !ok || at.After(time.Now()) {
		t.Fatal("successful update timestamp", at, ok)
	}
	if got, ok := MCOm37H(at); ok || got != at {
		t.Fatal("equal timestamp")
	}
	if got, ok := MCOm37H(at.Add(time.Hour)); ok || got != at {
		t.Fatal("newer caller snapshot")
	}
	// Both HTTP requests fail before writing; the last successful timestamp stays.
	updateGeoFiles(server.Client(), ":invalid", ":invalid", writer(site), writer(ip))
	if got, ok := MCOm37H(at); ok || got != at {
		t.Fatal("error changed timestamp")
	}
}

type geoRoundTripper func(*http.Request) (*http.Response, error)

func (f geoRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type geoTrackedBody struct{ closed chan struct{} }

func (b *geoTrackedBody) Read([]byte) (int, error) { return 0, io.EOF }
func (b *geoTrackedBody) Close() error             { close(b.closed); return nil }
func TestGeoDownloadBodyCleanupAndPartialSuccess(t *testing.T) {
	old := updateTimeAt5095C48.Load()
	defer updateTimeAt5095C48.Store(old)
	for _, success := range []bool{false, true} {
		stamp := time.Now().Add(-time.Hour)
		updateTimeAt5095C48.Store(&stamp)
		siteBody, ipBody := &geoTrackedBody{make(chan struct{})}, &geoTrackedBody{make(chan struct{})}
		client := &http.Client{Transport: geoRoundTripper(func(r *http.Request) (*http.Response, error) {
			var body io.ReadCloser = siteBody
			if r.URL.Path == "/ip" {
				body = ipBody
			}
			return &http.Response{StatusCode: 200, Body: body, Header: make(http.Header)}, nil
		})}
		updateGeoFiles(client, "http://fixture/site", "http://fixture/ip", func(io.Reader) (bool, error) { return success, nil }, func(io.Reader) (bool, error) { return false, errors.New("write fixture") })
		for _, body := range []*geoTrackedBody{siteBody, ipBody} {
			select {
			case <-body.closed:
			default:
				t.Fatal("response body not closed")
			}
		}
		got, changed := MCOm37H(stamp)
		if changed != success || (!success && got != stamp) {
			t.Fatal("partial success timestamp", changed)
		}
	}
}

func TestGeoScheduleWaitsAndKeepsIndependentHTTPTimeout(t *testing.T) {
	original := http.DefaultTransport
	calls := make(chan time.Duration, 8)
	http.DefaultTransport = geoRoundTripper(func(r *http.Request) (*http.Response, error) {
		deadline, ok := r.Context().Deadline()
		if !ok {
			return nil, errors.New("missing timeout")
		}
		calls <- time.Until(deadline)
		return nil, errors.New("fixture download failure")
	})
	defer func() { http.DefaultTransport = original }()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	Yk2PlrnXo(ctx, 100*time.Millisecond, "http://fixture/site", "http://fixture/ip")
	select {
	case <-calls:
		t.Fatal("unexpected immediate download")
	case <-time.After(20 * time.Millisecond):
	}
	for range 2 {
		select {
		case left := <-calls:
			if left < 599*time.Second || left > 600*time.Second {
				t.Fatal("HTTP timeout", left)
			}
		case <-time.After(time.Second):
			t.Fatal("scheduled request missing")
		}
	}
	cancel()
	// Wait for the current pair to complete and the canceled ticker worker to exit.
	select {
	case <-calls:
		t.Fatal("extra download after cancellation")
	case <-time.After(150 * time.Millisecond):
	}
}
