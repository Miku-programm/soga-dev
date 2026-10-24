package CpHCdepV

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"net"
	"net/http"
	ws "recovered.local/soga/FgcbiePIJLm"
	util "recovered.local/soga/UYoO4Hz"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNativeRawWebsocketListenerTCP(t *testing.T) {
	for _, tc := range []struct {
		name, path, uri, headers, remote string
		status                           int
	}{
		{"path_normalized", "native?ed=2048", "/native", "X-Real-Ip: 203.0.113.8\r\n", "203.0.113.8:0", 101},
		{"first_valid_header", "/native", "/native", "X-Forwarded-For: 198.51.100.2, 1.1.1.1\r\nX-Real-Ip: 203.0.113.8\r\n", "198.51.100.2:0", 101},
		{"invalid_first_ip", "/native", "/native", "X-Real-Ip: bad\r\nX-Forwarded-For: 198.51.100.2, 1.1.1.1\r\n", "198.51.100.2:0", 101},
		{"raw_uri_exact", "/native", "/native?q=1", "", "", 500},
		{"root_default", "", "/", "", "", 101},
	} {
		t.Run(tc.name, func(t *testing.T) {
			raw, e := net.Listen("tcp", "127.0.0.1:0")
			if e != nil {
				t.Fatal(e)
			}
			l := H7N2ayu(raw, tc.path)
			defer l.Close()
			c, e := net.Dial("tcp", raw.Addr().String())
			if e != nil {
				t.Fatal(e)
			}
			defer c.Close()
			c.SetDeadline(time.Now().Add(4 * time.Second))
			early := []byte{0xfb, 0xff, 0, 1, 2}
			encoded := base64.RawURLEncoding.EncodeToString(early)
			io.WriteString(c, "GET "+tc.uri+" HTTP/1.1\r\nHost: h\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Version: 13\r\nSec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==\r\nSec-WebSocket-Protocol: "+encoded+"\r\n"+tc.headers+"\r\n")
			res, e := http.ReadResponse(bufio.NewReader(c), nil)
			if e != nil {
				t.Fatal(e)
			}
			defer res.Body.Close()
			if res.StatusCode != tc.status {
				t.Fatal(res.Status)
			}
			if tc.status != 101 {
				p, _ := io.ReadAll(res.Body)
				if !strings.Contains(string(p), "path mismatch: expect: /native, actual: /native?q=1") {
					t.Fatal(string(p))
				}
				return
			}
			if res.Header.Get("Sec-WebSocket-Protocol") != "" {
				t.Fatal("raw native path does not echo early data")
			}
			s, e := l.Accept()
			if e != nil {
				t.Fatal(e)
			}
			defer s.Close()
			s.SetDeadline(time.Now().Add(4 * time.Second))
			if tc.remote != "" && s.RemoteAddr().String() != tc.remote {
				t.Fatal(s.RemoteAddr())
			}
			b := make([]byte, len(early))
			if _, e = io.ReadFull(s, b); e != nil || !bytes.Equal(b, early) {
				t.Fatal(b, e)
			}
			data := bytes.Repeat([]byte("native-payload"), 4000)
			done := make(chan error, 1)
			go func() { done <- util.WriteFrameRecovered(c, ws.StateClientSide, ws.OpBinary, true, data) }()
			b = make([]byte, len(data))
			if _, e = io.ReadFull(s, b); e != nil || !bytes.Equal(data, b) {
				t.Fatal(e)
			}
			if e = <-done; e != nil {
				t.Fatal(e)
			}
			go func() { _, e := s.Write(data); done <- e }()
			reader := &util.HLKO3Cdqw{GX04GFKh1Gz: c, WXUh5psqY7: ws.StateClientSide}
			if _, e = reader.NextFrame(); e != nil {
				t.Fatal(e)
			}
			b, e = io.ReadAll(reader)
			if e != nil || !bytes.Equal(b, data) {
				t.Fatal(e)
			}
			if e = <-done; e != nil {
				t.Fatal(e)
			}
		})
	}
}
func TestNativeWebsocketRTTIAndHTTPHeaderPriority(t *testing.T) {
	for _, tc := range []struct {
		v    any
		size uintptr
		off  []uintptr
	}{{RETaJc91{}, 48, []uintptr{0, 16, 24, 40}}, {E1sAyQm{}, 88, []uintptr{0, 16, 32, 40, 64, 80}}} {
		typ := reflect.TypeOf(tc.v)
		if typ.Size() != tc.size {
			t.Fatal(typ)
		}
		for i, x := range tc.off {
			if typ.Field(i).Offset != x {
				t.Fatal(typ.Field(i))
			}
		}
	}
	l := &RETaJc91{}
	if p := l.fa_PWrTv(http.Header{"X-Real-Ip": {"bad"}, "X-Forwarded-For": {"1.2.3.4"}}); p != nil {
		t.Fatal(p)
	}
	if p := l.fa_PWrTv(http.Header{"X-Forwarded-For": {" 1.2.3.4, 5.6.7.8"}}); p.String() != "1.2.3.4" {
		t.Fatal(p)
	}
}
