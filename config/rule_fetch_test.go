package config

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"recovered.local/soga/audit/matcher"
)

func saveRuleFetchState(t *testing.T) {
	blockURL, whiteURL := configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi
	block, white, client := blockListAt5095A90.Load(), whiteListAt5095AD8.Load(), ruleHTTPClientAt5029350
	t.Cleanup(func() {
		configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi = blockURL, whiteURL
		blockListAt5095A90.Store(block)
		whiteListAt5095AD8.Store(white)
		ruleHTTPClientAt5029350 = client
	})
}

func TestRuleFetchHTTPStatusAndLineSemantics(t *testing.T) {
	saveRuleFetchState(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method %s", r.Method)
		}
		if r.URL.Path == "/redirect" {
			http.Redirect(w, r, "/rules", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		io.WriteString(w, "\n # comment\r\n domain:example.test # suffix \n\nport:80\nport:80\n")
	}))
	defer server.Close()
	configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi = server.URL+"/redirect", server.URL+"/rules"
	want := matcher.VHWq49iRat{"# comment", "domain:example.test # suffix", "port:80", "port:80"}
	for _, get := range []func() (matcher.VHWq49iRat, error){Klz7sa_01, Cf33JCugNk1} {
		got, err := get()
		if err != nil || !reflect.DeepEqual(got, want) {
			t.Fatalf("%v, %v", got, err)
		}
	}
	configAt509B660.BooNOdXcSa3U, configAt5096E00.M4qDq3KsUi = "", ""
	local := matcher.VHWq49iRat{"original"}
	blockListAt5095A90.Store(&local)
	whiteListAt5095AD8.Store(&local)
	got, _ := Klz7sa_01()
	got[0] = "mutated"
	white, _ := Cf33JCugNk1()
	if white[0] != "mutated" {
		t.Fatal("snapshot was copied")
	}
}

type failingRuleBody struct {
	closed bool
	err    error
}

func (b *failingRuleBody) Read([]byte) (int, error) { return 0, b.err }
func (b *failingRuleBody) Close() error             { b.closed = true; return errors.New("close ignored") }

type ruleRoundTripper func(*http.Request) (*http.Response, error)

func (f ruleRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestRuleFetchReadFailureClosesBody(t *testing.T) {
	saveRuleFetchState(t)
	want := errors.New("body failure")
	body := &failingRuleBody{err: want}
	ruleHTTPClientAt5029350 = &http.Client{Transport: ruleRoundTripper(func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: body, Header: make(http.Header)}, nil
	})}
	configAt509B660.BooNOdXcSa3U = "http://rules.invalid/"
	got, err := Klz7sa_01()
	if got != nil || err != want || !body.closed {
		t.Fatalf("rules=%v error=%v closed=%v", got, err, body.closed)
	}
}
