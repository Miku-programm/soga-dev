package FgcbiePIJLm

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type rawHandshakeWire struct {
	*bytes.Reader
	output bytes.Buffer
}

func (w *rawHandshakeWire) Write(p []byte) (int, error) { return w.output.Write(p) }
func TestOriginalRawUpgradeHeaderConditions(t *testing.T) {
	base := "GET /native HTTP/1.1\r\nHost: \r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Version: 13\r\nSec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==\r\n\r\n"
	for _, tc := range []struct {
		name, old, new string
		status         int
		want           error
	}{
		{"empty_host_allowed", "", "", 101, nil},
		{"missing_host", "Host: \r\n", "", 400, ErrHandshakeBadHost},
		{"method", "GET ", "POST ", 405, ErrHandshakeBadMethod},
		{"major", "HTTP/1.1", "HTTP/2.0", 505, ErrHandshakeBadProtocol},
		{"raw_minor_nibble_quirk", "HTTP/1.1", "HTTP/1.:", 101, nil},
		{"missing_version", "Sec-WebSocket-Version: 13\r\n", "", 400, ErrHandshakeBadSecVersion},
		{"bad_version", "Sec-WebSocket-Version: 13", "Sec-WebSocket-Version: 14", 426, ErrHandshakeUpgradeRequired},
		{"bad_connection", "Connection: Upgrade", "Connection: none", 400, ErrHandshakeBadConnection},
	} {
		t.Run(tc.name, func(t *testing.T) {
			request := base
			if tc.old != "" {
				request = strings.Replace(request, tc.old, tc.new, 1)
			}
			w := &rawHandshakeWire{Reader: bytes.NewReader([]byte(request))}
			_, e := (L4gwp9B{}).Upgrade(w)
			if e != tc.want {
				t.Fatalf("%v != %v", e, tc.want)
			}
			r, e := http.ReadResponse(bufio.NewReader(&w.output), nil)
			if e != nil {
				t.Fatal(e)
			}
			defer r.Body.Close()
			if r.StatusCode != tc.status {
				t.Fatal(r.Status)
			}
		})
	}
	if n, e := fm9r93qPp([]byte("1:")); n != 20 || e != nil {
		t.Fatal(n, e)
	}
	if _, e := fm9r93qPp(nil); e == nil || e.Error() != "converting empty bytes to int" {
		t.Fatal(e)
	}
	if _, e := fm9r93qPp([]byte("a")); e == nil || e.Error() != "a is not a numeric character" {
		t.Fatal(e)
	}
}
func TestOriginalRawUpgradeCallbacksAndInputErrors(t *testing.T) {
	request := "GET /x?q=1 HTTP/1.1\r\nHost: h\r\nUpgrade: WebSocket\r\nConnection: other, upgrade\r\nSec-WebSocket-Version: 13\r\nSec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==\r\nSec-WebSocket-Protocol: data\r\nX-Foo: value\r\n\r\n"
	var calls []string
	u := L4gwp9B{OnRequest: func(b []byte) error { calls = append(calls, string(b)); return nil }, OnHost: func(b []byte) error { calls = append(calls, string(b)); return nil }, ProtocolCustom: func(b []byte) (string, bool) { calls = append(calls, string(b)); return "native", true }, OnHeader: func(k, v []byte) error { calls = append(calls, string(k)+"="+string(v)); return nil }, OnBeforeUpgrade: func() (HDpTJacC, error) {
		calls = append(calls, "before")
		return BuqwQfCAXz("X-Native: true\r\n"), nil
	}}
	w := &rawHandshakeWire{Reader: bytes.NewReader([]byte(request))}
	hs, e := u.Upgrade(w)
	if e != nil || hs.Protocol != "native" {
		t.Fatal(hs, e)
	}
	if !reflect.DeepEqual(calls, []string{"/x?q=1", "h", "data", "X-Foo=value", "before"}) {
		t.Fatal(calls)
	}
	if !strings.Contains(w.output.String(), "X-Native: true\r\n") {
		t.Fatal(w.output.String())
	}
	w = &rawHandshakeWire{Reader: bytes.NewReader([]byte("GET / HTTP/1.1"))}
	if _, e = u.Upgrade(w); !errors.Is(e, io.EOF) || w.output.Len() != 0 {
		t.Fatal(e, w.output.String())
	}
	if typ := reflect.TypeOf(u); typ.Size() != 104 || typ.FieldByIndex([]int{8}).Offset != 72 {
		t.Fatal(typ)
	}
}
