package CAodbhZf

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/synctest"
	"time"
)

type configRoundTrip func(*http.Request) (*http.Response, error)

func (f configRoundTrip) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestRemoteFetchReadsErrorStatusAndFollowsRedirect(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/redirect" {
			http.Redirect(w, r, "/body", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = io.WriteString(w, "configuration body")
	}))
	defer srv.Close()
	data, err := oDPndU2e(srv.URL + "/redirect")
	if err != nil || string(data) != "configuration body" {
		t.Fatalf("%q, %v", data, err)
	}
	if remoteClientAt5029248.Timeout != time.Minute {
		t.Fatal("wrong original timeout")
	}
}

func TestRemotePollingChangesFailuresAndCancellation(t *testing.T) {
	old := remoteClientAt5029248
	t.Cleanup(func() { remoteClientAt5029248 = old })
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		requests, calls := 0, 0
		payload := ""
		fetchError := false
		remoteClientAt5029248 = &http.Client{Timeout: time.Minute, Transport: configRoundTrip(func(*http.Request) (*http.Response, error) {
			requests++
			if fetchError {
				return nil, errors.New("transport failure")
			}
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(payload)), Header: make(http.Header)}, nil
		})}
		failLoad := false
		done := make(chan struct{})
		go func() {
			defer close(done)
			KJa7rJUaLPYe(ctx, "http://test.invalid/config", func(data []byte) error {
				calls++
				if failLoad {
					return errors.New("invalid configuration")
				}
				return nil
			})
		}()
		synctest.Wait()
		if requests != 1 || calls != 0 {
			t.Fatal("initial empty response must not call loader")
		}
		payload = "first"
		time.Sleep(time.Minute)
		synctest.Wait()
		if requests != 2 || calls != 1 {
			t.Fatal(requests, calls)
		}
		time.Sleep(time.Minute)
		synctest.Wait()
		if calls != 1 {
			t.Fatal("unchanged response reloaded")
		}
		payload = "second"
		failLoad = true
		time.Sleep(time.Minute)
		synctest.Wait()
		if calls != 2 {
			t.Fatal(calls)
		}
		failLoad = false
		time.Sleep(time.Minute)
		synctest.Wait()
		if calls != 3 {
			t.Fatal("failed parse was incorrectly cached")
		}
		fetchError = true
		time.Sleep(time.Minute)
		synctest.Wait()
		if calls != 3 {
			t.Fatal("transport error passed to parser")
		}
		fetchError = false
		time.Sleep(time.Minute)
		synctest.Wait()
		if calls != 3 {
			t.Fatal("transport failure destroyed cached body")
		}
		cancel()
		synctest.Wait()
		select {
		case <-done:
		default:
			t.Fatal("cancellation did not end polling")
		}
	})
}
