package accesslog

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	C "recovered.local/soga/config"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestOriginalAccessLayoutsAndEscaping(t *testing.T) {
	for _, tc := range []struct {
		typ     reflect.Type
		size    uintptr
		offsets []uintptr
	}{
		{reflect.TypeFor[cUwNbDJ6opWy](), 128, []uintptr{0, 24, 40, 48, 56, 72, 88, 104, 112}},
		{reflect.TypeFor[cToP4Y8](), 48, []uintptr{0, 8, 16, 32}},
		{reflect.TypeFor[aGuDZ8NsNT](), 112, []uintptr{0, 8, 16, 32, 72, 80, 88}},
	} {
		if tc.typ.Size() != tc.size {
			t.Fatal(tc.typ.Size())
		}
		for i, o := range tc.offsets {
			if tc.typ.Field(i).Offset != o {
				t.Fatalf("%v field %d", tc.typ, i)
			}
		}
	}
	raw := string([]byte{'"', '\\', '\n', '\r', '\t', 0, 8, 12, 31, 0xff}) + "<>&\u2028"
	var b strings.Builder
	oRjSReB(&b, raw)
	want := `"\"\\\n\r\t\u0000\u0008\u000c\u001f` + string([]byte{255}) + "<>&\u2028\""
	if b.String() != want {
		t.Fatalf("%q != %q", b.String(), want)
	}
	if v4V_xxf("a|b\r\nc\td") != "a_bc\td" {
		t.Fatal("local escaping")
	}
	if cap(queueAt5096948) != 8192 {
		t.Fatal("queue capacity")
	}
}

func TestClickhouseRealHTTPBatchAndDDL(t *testing.T) {
	var mu sync.Mutex
	var bodies []string
	var requests []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		r.Body.Close()
		if r.Method != "POST" || r.Header.Get("X-ClickHouse-User") != "alice" || r.Header.Get("X-ClickHouse-Key") != "secret" {
			t.Errorf("request %s %v", r.Method, r.Header)
		}
		mu.Lock()
		bodies = append(bodies, string(data))
		requests = append(requests, r.URL.RequestURI())
		mu.Unlock()
		w.WriteHeader(200)
		_, _ = w.Write([]byte("drain response"))
	}))
	defer srv.Close()
	s := &aGuDZ8NsNT{jVe51FV: &C.ZNxaPWNlDqa{ASvYULtJ7: srv.URL, ZH4VO_H4JZ: "db", IJFBj1lIX: "logs", H0x74Y: "alice", SNWlSe: "secret", BaLySDum3K: 2, EVFQh90cvJK: true}, yw1h7z1aNajs: srv.Client(), bV1jN0ia: srv.URL + "/?query=insert"}
	defer s.t_DddcBjF7()
	r := cUwNbDJ6opWy{time.Unix(123, 999), "tcp", 7, 42, "1.2.3.4", "a\"b", "::1", 443, "31.13.92.37"}
	var b strings.Builder
	s.ajkzPuuQ(&b, &r)
	if s.peQJTFa != 1 || len(bodies) != 0 {
		t.Fatal("premature flush")
	}
	row := "{\"time\":123,\"network\":\"tcp\",\"node_id\":7,\"user_id\":42,\"user_ip\":\"1.2.3.4\",\"access_host\":\"31.13.92.37\",\"host\":\"a\\\"b\",\"dest_ip\":\"::1\",\"dest_port\":443}\n"
	if s.egrmDIRRx5.String() != row {
		t.Fatalf("row %q", s.egrmDIRRx5.String())
	}
	s.ajkzPuuQ(&b, &r)
	if !s.q3mZAt || s.peQJTFa != 0 || s.egrmDIRRx5.Len() != 0 {
		t.Fatal("flush state")
	}
	if len(bodies) != 4 || bodies[0] != "CREATE DATABASE IF NOT EXISTS `db`" || bodies[1] != s.fJMhCf() || bodies[2] != "ALTER TABLE `db`.`logs` ADD COLUMN IF NOT EXISTS access_host String DEFAULT '' AFTER user_ip" || bodies[3] != row+row || requests[3] != "/?query=insert" {
		t.Fatalf("requests: %q %q", bodies, requests)
	}
	s.ajkzPuuQ(&b, &r)
	s.eyAgzSwv7d()
	if len(bodies) != 5 {
		t.Fatal("DDL repeated after success")
	}
}

type observingBody struct {
	*bytes.Reader
	closed bool
	read   int
}

func (b *observingBody) Read(p []byte) (int, error) {
	n, e := b.Reader.Read(p)
	b.read += n
	return n, e
}
func (b *observingBody) Close() error { b.closed = true; return nil }

type accessTransport func(*http.Request) (*http.Response, error)

func (f accessTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestClickhouseFailureDrainAndRetryDelay(t *testing.T) {
	body := &observingBody{Reader: bytes.NewReader([]byte(" " + strings.Repeat("x", 700)))}
	calls := 0
	s := &aGuDZ8NsNT{jVe51FV: &C.ZNxaPWNlDqa{ASvYULtJ7: "http://fixture", ZH4VO_H4JZ: "db", IJFBj1lIX: "logs", BaLySDum3K: 1}, bV1jN0ia: "http://fixture/insert"}
	s.yw1h7z1aNajs = &http.Client{Transport: accessTransport(func(r *http.Request) (*http.Response, error) {
		calls++
		return &http.Response{StatusCode: 201, Body: body, Header: make(http.Header)}, nil
	})}
	err := s.kByXwtYgYQ("http://fixture", nil)
	if err == nil || err.Error() != "clickhouse response 201: "+strings.Repeat("x", 511) || !body.closed || body.Reader.Len() != 0 {
		t.Fatalf("status/drain: %v %d %v", err, body.read, body.closed)
	}
	body.Reader = bytes.NewReader([]byte("failed"))
	var b strings.Builder
	r := cUwNbDJ6opWy{pv0C8wVKLJh: time.Now()}
	s.ajkzPuuQ(&b, &r)
	first := calls
	if s.q3mZAt || s.ixZxkij70Gar.IsZero() || s.peQJTFa != 0 || s.egrmDIRRx5.Len() != 0 {
		t.Fatal("failed batch not dropped")
	}
	s.ajkzPuuQ(&b, &r)
	if calls != first {
		t.Fatal("retry before one minute")
	}
	s.ixZxkij70Gar = time.Now().Add(-time.Minute)
	s.ajkzPuuQ(&b, &r)
	if calls != first+1 {
		t.Fatal("retry after one minute missing")
	}
	s.q3mZAt = true
	s.ajkzPuuQ(&b, &r)
	if calls != first+2 || s.peQJTFa != 0 {
		t.Fatal("failed insert must also drop batch")
	}
}
