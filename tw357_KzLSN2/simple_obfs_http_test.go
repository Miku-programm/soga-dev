package tw357_KzLSN2

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	S "recovered.local/soga/ad5KEbw_EzWk"
	"strings"
	"testing"
	"unsafe"
)

func TestSimpleObfsHTTPBodyAndRawTail(t *testing.T) {
	for _, tc := range []struct{ name, headers, body, want string }{
		{"length", "Content-Length: 3\r\n", "abcTAIL", "abcTAIL"},
		{"no-body", "", "TAIL", "TAIL"},
		{"chunked", "Transfer-Encoding: chunked\r\n", "3\r\nabc\r\n2\r\nde\r\n0\r\n\r\nTAIL", "abcdeTAIL"},
		{"trailer", "Transfer-Encoding: chunked\r\nTrailer: X-End\r\n", "3\r\nabc\r\n0\r\nX-End: value\r\n\r\nTAIL", "abcTAIL"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := &wireFixture{read: bytes.NewReader([]byte("POST / HTTP/1.1\r\nHost: [::1]:443\r\n" + tc.headers + "\r\n" + tc.body))}
			c := i0Z5v8x(w).(*k7yarZ7PgV6)
			info := &S.L7E7H2aW8KVs{}
			c.SetServerInfo(info)
			got, e := io.ReadAll(c)
			if e != nil || string(got) != tc.want {
				t.Fatalf("%q %v", got, e)
			}
			if info.EFLMTH != "::1" || c.dNeNOip != nil {
				t.Fatal("state", info.EFLMTH)
			}
			if tc.name == "trailer" && c.d3aah5M.Trailer.Get("X-End") != "value" {
				t.Fatal("trailer")
			}
		})
	}
}
func TestSimpleObfsHTTPErrorsAndWrite(t *testing.T) {
	if unsafe.Sizeof(k7yarZ7PgV6{}) != 72 {
		t.Fatal("RTTI")
	}
	for _, body := range []string{"G\r\nabc\r\n", "3\nabc\r\n", "3\r\nabcXX", "3\r\na"} {
		c := i0Z5v8x(&wireFixture{read: bytes.NewReader([]byte("POST / HTTP/1.1\r\nHost: x\r\nTransfer-Encoding: chunked\r\n\r\n" + body))}).(*k7yarZ7PgV6)
		c.SetServerInfo(&S.L7E7H2aW8KVs{})
		if n, e := c.Read(make([]byte, 100)); n != 0 || e == nil || !c.xJOdz9bQ9L8 {
			t.Fatal(n, e, c.xJOdz9bQ9L8)
		}
	}
	w := &wireFixture{short: true}
	c := i0Z5v8x(w).(*k7yarZ7PgV6)
	if n, e := c.Write([]byte("payload")); n != 7 || e != nil {
		t.Fatal(n, e)
	}
	r, e := http.ReadResponse(bufio.NewReader(bytes.NewReader(w.output.Bytes())), nil)
	if e != nil {
		t.Fatal(e)
	}
	random, e := base64.StdEncoding.DecodeString(r.Header.Get("Sec-WebSocket-Accept"))
	if e != nil || len(random) != 20 || r.StatusCode != 101 {
		t.Fatal(r, e)
	}
	if !strings.HasSuffix(w.output.String(), "\r\n\r\npayload") {
		t.Fatal(w.output.String())
	}
	if n, e := c.Write([]byte("raw")); n != 0 || e != nil {
		t.Fatal("second short write", n, e)
	}
}
