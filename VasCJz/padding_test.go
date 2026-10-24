package VasCJz

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"unsafe"
)

func TestOriginalPaddingGenerationAndWirePlacements(t *testing.T) {
	if unsafe.Sizeof(AHn9pfr{}) != 64 || unsafe.Sizeof(LaPHbx6B50uZ{}) != 88 {
		t.Fatal("RTTI")
	}
	cfg := &R0J19auWof{}
	for _, n := range []int{1, 3, 100, 1000} {
		got := D940rWda18j4("repeat-x", n)
		if got != strings.Repeat("X", n) {
			t.Fatal("repeat")
		}
		got = D940rWda18j4("tokenish", n)
		if !cfg.IsPaddingValid(got, int32(n), int32(n), "tokenish") {
			t.Fatalf("encoded size %d: %q", n, got)
		}
		for _, c := range got {
			if !strings.ContainsRune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", c) {
				t.Fatal("alphabet")
			}
		}
	}
	for _, p := range []string{"header", "queryInHeader", "query", "cookie"} {
		t.Run(p, func(t *testing.T) {
			r := httptest.NewRequest("POST", "http://example.test/p?keep=1", nil)
			cfg := &R0J19auWof{H6PfvFysST: "pad", LjJbxC6vDL: "X-Padding", V0qykwP: p}
			placement := LaPHbx6B50uZ{TldsrRa: 12, C5zKbVfj3g_: AHn9pfr{AnBhAR: p, AVHkw5m: "pad", HxzvtHynCM: "X-Padding", SZY0xnkLskud: "https://example.test/cover?remove=1#part"}, Yxjd2mW: "repeat-x"}
			cfg.ApplyXPaddingToRequest(r, placement)
			value, label := cfg.ExtractXPaddingFromRequest(r, true)
			if value != strings.Repeat("X", 12) || label == "" {
				t.Fatalf("extract %q %q", value, label)
			}
			if p == "query" && r.URL.Query().Get("keep") != "1" {
				t.Fatal("query preserve")
			}
			if p == "queryInHeader" && strings.Contains(r.Header.Get("X-Padding"), "remove=") {
				t.Fatal("header URL query must be replaced")
			}
			w := httptest.NewRecorder()
			cfg.ApplyXPaddingToResponse(w, placement)
			switch p {
			case "header", "queryInHeader":
				if w.Header().Get("X-Padding") == "" {
					t.Fatal("missing response header")
				}
			case "cookie":
				if !strings.Contains(w.Header().Get("Set-Cookie"), "pad=XXXXXXXXXXXX; Path=/") {
					t.Fatal(w.Header())
				}
			case "query":
				if len(w.Header()) != 0 {
					t.Fatal("query response must remain untouched")
				}
			}
		})
	}
	if s, ok := gEb3WsxKBjs(0, "abc"); s != "" || ok {
		t.Fatal("empty generator")
	}
	if s, ok := gEb3WsxKBjs(2, ""); s != "" || ok {
		t.Fatal("empty alphabet")
	}
}
func TestPaddingOriginalExtractionPrecedence(t *testing.T) {
	c := &R0J19auWof{H6PfvFysST: "pad", LjJbxC6vDL: "X-Cover", V0qykwP: "queryInHeader"}
	r := httptest.NewRequest("GET", "http://example.test/?pad=query&x_padding=legacy", nil)
	r.Header.Set("Referer", "https://cover.test/?unused=1")
	if v, label := c.ExtractXPaddingFromRequest(r, false); v != "" || label != "queryInHeader=Referer, key=x_padding" {
		t.Fatal(v, label)
	}
	r.Header.Set("Referer", ":bad")
	r.Header.Set("X-Cover", "https://cover.test/?pad=header")
	r.AddCookie(&http.Cookie{Name: "pad", Value: "cookie"})
	if v, label := c.ExtractXPaddingFromRequest(r, false); v != "cookie" || label != "cookie, key=pad" {
		t.Fatal(v, label)
	}
	r.Header.Del("Cookie")
	if v, _ := c.ExtractXPaddingFromRequest(r, true); v != "header" {
		t.Fatal(v)
	}
	r.Header.Set("X-Cover", "https://cover.test/?absent=1")
	if v, label := c.ExtractXPaddingFromRequest(r, true); v != "" || label != "queryInHeader=X-Cover, key=pad" {
		t.Fatal(v, label)
	}
	r.Header.Set("X-Cover", ":bad")
	if v, _ := c.ExtractXPaddingFromRequest(r, true); v != "query" {
		t.Fatal(v)
	}
	if v, l := c.ExtractXPaddingFromRequest(nil, true); v != "" || l != "" {
		t.Fatal(v, l)
	}
}
func TestOriginalMetadataMixedPlacementsAndCORS(t *testing.T) {
	r := httptest.NewRequest("OPTIONS", "http://example.test/base/path-session/9?x_session=query-session&x_seq=10", nil)
	r.Header.Set("X-Session", "header-session")
	r.Header.Set("X-Seq", "11")
	r.AddCookie(&http.Cookie{Name: "x_session", Value: "cookie-session"})
	r.AddCookie(&http.Cookie{Name: "x_seq", Value: "12"})
	for _, tc := range []struct{ sp, qp, s, q string }{{"path", "path", "path-session", "9"}, {"header", "path", "header-session", "path-session"}, {"query", "header", "query-session", "11"}, {"cookie", "query", "cookie-session", "10"}, {"header", "cookie", "header-session", "12"}} {
		c := &R0J19auWof{XEKp9SbqKQ: tc.sp, HqtZJaRZNw: tc.qp}
		s, q := c.ExtractMetaFromRequest(r, "/base/")
		if s != tc.s || q != tc.q {
			t.Fatalf("%+v -> %q %q", tc, s, q)
		}
	}
	w := httptest.NewRecorder()
	r.Header.Set("Origin", "https://client.test")
	r.Header.Set("Access-Control-Request-Method", "PATCH")
	r.Header.Set("Access-Control-Request-Headers", "X-Test")
	c := &R0J19auWof{XEKp9SbqKQ: "cookie"}
	c.WriteResponseHeader(w, "OPTIONS", r.Header)
	if w.Header().Get("Access-Control-Allow-Origin") != "https://client.test" || w.Header().Get("Access-Control-Allow-Methods") != "PATCH" || w.Header().Get("Access-Control-Allow-Headers") != "X-Test" || w.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Fatal(w.Header())
	}
	u, _ := url.Parse("https://x/?keep=1")
	VzIFuzeyXw(u, "pad", "x x")
	if u.RawQuery != "keep=1&pad=x+x" {
		t.Fatal(u.RawQuery)
	}
}
