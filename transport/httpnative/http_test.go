package httpnative

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestNativeHTTPObfsWire(t *testing.T) {
	var layout Foi2jG
	if unsafe.Sizeof(layout) != 80 || unsafe.Offsetof(layout.ewxQcQsYiLi) != 72 {
		t.Fatal("native RTTI layout")
	}
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	a.SetDeadline(time.Now().Add(3 * time.Second))
	b.SetDeadline(time.Now().Add(3 * time.Second))
	payload := bytes.Repeat([]byte{0x67}, 6000)
	done := make(chan error, 1)
	go func() {
		c := NewHTTPForFactory(b, "/a b", "example.org", 1)
		defer c.Free()
		p := make([]byte, len(payload))
		if _, err := io.ReadFull(c, p); err != nil || !bytes.Equal(p, payload) {
			done <- fmt.Errorf("buffered payload %v", err)
			return
		}
		if n, err := c.Write([]byte("reply")); n != 5 || err != nil {
			done <- fmt.Errorf("response %d %v", n, err)
			return
		}
		p = make([]byte, 4)
		if _, err := io.ReadFull(c, p); err != nil || string(p) != "next" {
			done <- fmt.Errorf("raw subsequent read %q %v", p, err)
			return
		}
		_, err := c.Write([]byte("raw"))
		done <- err
	}()
	// Query is ignored and Path is unescaped. No Content-Length is required.
	header := "GET /a%20b?x=1 HTTP/1.1\r\nHost: example.org:8080\r\n\r\n"
	if _, err := a.Write(append([]byte(header), payload...)); err != nil {
		t.Fatal(err)
	}
	r := bufio.NewReader(a)
	var response strings.Builder
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			t.Fatal(err)
		}
		response.WriteString(line)
		if strings.HasPrefix(line, "Sec-WebSocket-Accept: ") {
			p, err := base64.StdEncoding.DecodeString(strings.TrimSpace(strings.TrimPrefix(line, "Sec-WebSocket-Accept: ")))
			if err != nil || len(p) != 20 {
				t.Fatal("random accept", p, err)
			}
		}
		if line == "\r\n" {
			break
		}
	}
	if !strings.HasPrefix(response.String(), "HTTP/1.1 101 Switching Protocols\r\nServer: nginx/1.27.0\r\n") {
		t.Fatal(response.String())
	}
	p := make([]byte, 5)
	if _, err := io.ReadFull(r, p); err != nil || string(p) != "reply" {
		t.Fatal(string(p), err)
	}
	if _, err := a.Write([]byte("next")); err != nil {
		t.Fatal(err)
	}
	p = make([]byte, 3)
	if _, err := io.ReadFull(r, p); err != nil || string(p) != "raw" {
		t.Fatal(string(p), err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

// Only the first Read is used to collect the HTTP header in the original.
type httpScriptConn struct {
	net.Conn
	parts   [][]byte
	written bytes.Buffer
	short   bool
	reads   int
}

func (c *httpScriptConn) Read(p []byte) (int, error) {
	c.reads++
	if len(c.parts) == 0 {
		return 0, io.EOF
	}
	n := copy(p, c.parts[0])
	c.parts[0] = c.parts[0][n:]
	if len(c.parts[0]) == 0 {
		c.parts = c.parts[1:]
	}
	return n, nil
}
func (c *httpScriptConn) Write(p []byte) (int, error) {
	c.written.Write(p)
	if c.short {
		return 1, nil
	}
	return len(p), nil
}
func TestNativeHTTPObfsParserBranches(t *testing.T) {
	t.Setenv("GODEBUG", "httplaxcontentlength=0")
	for _, tc := range []struct{ name, request, errorText string }{
		{"conflicting-content-length", "GET / HTTP/1.1\r\nHost: x\r\nContent-Length: 1\r\nContent-Length: 2\r\n\r\nx", "http: message cannot contain multiple Content-Length headers; got [\"1\" \"2\"]"},
		{"empty-content-length", "GET / HTTP/1.1\r\nContent-Length:\r\n\r\nx", "invalid empty Content-Length \"\""},
		{"negative-content-length", "GET / HTTP/1.1\r\nContent-Length: -1\r\n\r\nx", "bad Content-Length \"-1\""},
		{"unsupported-transfer-encoding", "GET / HTTP/1.1\r\nTransfer-Encoding: gzip\r\n\r\nx", "unsupported transfer encoding: \"gzip\""},
		{"multiple-transfer-encoding", "GET / HTTP/1.1\r\nTransfer-Encoding: chunked\r\nTransfer-Encoding: chunked\r\n\r\nx", "too many transfer encodings: [\"chunked\" \"chunked\"]"},
		{"bad-trailer", "GET / HTTP/1.1\r\nTransfer-Encoding: chunked\r\nTrailer: Content-Length\r\n\r\nx", "bad trailer key \"Content-Length\""},
		{"duplicate-host", "GET / HTTP/1.1\r\nHost: x\r\nHost: y\r\n\r\nx", "too many Host headers"},
		{"initial-space", "GET / HTTP/1.1\r\n Host: x\r\n\r\nx", "malformed MIME header initial line: \" Host: x\""},
		{"invalid-version", "GET / HTTP/1.11\r\n\r\nx", "malformed HTTP version \"HTTP/1.11\""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			raw := &httpScriptConn{parts: [][]byte{[]byte(tc.request)}}
			c := NewHTTPForFactory(raw, "", "", 0)
			_, err := c.Read(make([]byte, 16))
			if err == nil || err.Error() != tc.errorText {
				t.Fatalf("%v; want %s", err, tc.errorText)
			}
			if raw.reads != 1 {
				t.Fatal("read escaped bounded buffer", raw.reads)
			}
		})
	}
	for _, header := range []string{
		"GET / HTTP/1.0\r\nTransfer-Encoding: gzip\r\n\r\n",
		"GET / HTTP/1.1\r\nContent-Length: 1\r\nContent-Length: 1\r\n\r\n",
		"GET / HTTP/1.1\r\nTransfer-Encoding: ChUnKeD\r\nTrailer: X-Foo\r\n\r\n",
		"CONNECT example.org:443 HTTP/1.1\r\n\r\n",
		"PRI * HTTP/2.0\r\n\r\n",
	} {
		c := NewHTTPForFactory(&httpScriptConn{parts: [][]byte{[]byte(header + "wire")}}, "", "", 0)
		p := make([]byte, 4)
		if _, err := io.ReadFull(c, p); err != nil || string(p) != "wire" {
			t.Fatal(header, string(p), err)
		}
	}
}
func TestNativeHTTPObfsReadFailureAndShortWriteState(t *testing.T) {
	raw := &httpScriptConn{parts: [][]byte{[]byte("GET / HTTP/1.1\r\nHost: x\r\n"), []byte("\r\nrest")}, short: true}
	c := NewHTTPForFactory(raw, "", "", 0)
	if _, err := c.Read(make([]byte, 20)); !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatal(err)
	}
	p := make([]byte, 20)
	if n, err := c.Read(p); err != nil || string(p[:n]) != "\r\nrest" {
		t.Fatal(n, err, string(p[:n]))
	}
	if n, err := c.Write([]byte("abcd")); n != 4 || err != nil {
		t.Fatal(n, err)
	}
	if n, err := c.Write([]byte("efgh")); n != 1 || err != nil {
		t.Fatal(n, err)
	}
	if strings.Count(raw.written.String(), "101 Switching Protocols") != 1 {
		t.Fatal("response repeated")
	}
}
