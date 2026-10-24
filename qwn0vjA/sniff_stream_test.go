package qwn0vjA

import (
	"crypto/tls"
	"encoding/binary"
	"io"
	"net"
	"recovered.local/soga/HxKViEp2a"
	"testing"
)

func TestOriginalHTTPHeaderOrderingAndErrorIdentity(t *testing.T) {
	h := &q7GP8uwJQpnC{}
	for _, tc := range []struct {
		data, want string
		err        error
	}{
		{"GET / HTTP/1.1\r\nHost: EXAMPLE.COM:443\r\n\r\n", "example.com", nil},
		{"getaway\nHost: First.test\r\nHost: Last.test\n", "last.test", nil},
		{"GET /\r\nHost: Before.test\r\n\r\nHost: After.test\r\n", "after.test", nil},
		{"GET /\nHost: Before.test\n\nHost: After.test\n", "before.test", nil},
		{"G", "", HxKViEp2a.ErrorAt5097810},
		{"PATCH / HTTP/1.1\nHost: a\n", "", errorAt5098370},
		{"GET /\nHost : a\n", "", HxKViEp2a.ErrorAt5097820},
	} {
		got, err := h.Sniff([]byte(tc.data))
		if err != tc.err {
			t.Fatalf("%q: error %v want %v", tc.data, err, tc.err)
		}
		if err == nil && (got.Type() != "http" || got.Value() != tc.want) {
			t.Fatal(got)
		}
	}
	if _, err := h.ParseHost("a:b:c"); err == nil {
		t.Fatal("bad colon accepted")
	}
	if got, err := h.ParseHost("[::1]:443"); err != nil || got != "::1" {
		t.Fatal(got, err)
	}
}
func TestOriginalBTAndDispatcher(t *testing.T) {
	s := KZgBMVcpL{&ezgFiQP115F{}, &q7GP8uwJQpnC{}, &aln4LlVj3z{}}
	for _, data := range [][]byte{append([]byte{19}, []byte("BitTorrent protocol")...), append(append([]byte{19}, []byte("BitTorrent protocol")...), 1, 2, 3)} {
		got, err := s.Sniff(data, 0)
		if err != nil || got.Type() != "bt" || got.Value() != "bt" {
			t.Fatal(got, err)
		}
		if _, err = s.Sniff(data, 1); err != errorAt5098350 {
			t.Fatal(err)
		}
	}
	if _, err := s.Sniff([]byte{1, 2}, 0); err != errorAt5098350 {
		t.Fatal(err)
	}
}
func captureHello(t *testing.T) []byte {
	t.Helper()
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()
	done := make(chan error, 1)
	go func() {
		c := tls.Client(a, &tls.Config{ServerName: "Mixed.Example.test", InsecureSkipVerify: true})
		done <- c.Handshake()
	}()
	head := make([]byte, 5)
	if _, err := io.ReadFull(b, head); err != nil {
		t.Fatal(err)
	}
	body := make([]byte, int(binary.BigEndian.Uint16(head[3:])))
	if _, err := io.ReadFull(b, body); err != nil {
		t.Fatal(err)
	}
	b.Close()
	<-done
	return append(head, body...)
}
func syntheticHello(name string) []byte {
	body := make([]byte, 42)
	body[0] = 1
	body[4] = 3
	body[5] = 3 // 原解析器本身不验证这些字段。
	body[38] = 0
	body[39] = 0
	body[40] = 0
	body[41] = 0
	names := []byte{0, byte((3 + len(name)) >> 8), byte(3 + len(name)), 0, byte(len(name) >> 8), byte(len(name))}
	names = names[1:]
	names = append(names, []byte(name)...)
	ext := []byte{0, 0, byte(len(names) >> 8), byte(len(names))}
	ext = append(ext, names...)
	body = append(body, byte(len(ext)>>8), byte(len(ext)))
	body = append(body, ext...)
	record := []byte{22, 3, 0, byte(len(body) >> 8), byte(len(body))}
	return append(record, body...)
}
func TestOriginalTLSClientHelloAndBoundaryClassifications(t *testing.T) {
	s := &aln4LlVj3z{}
	actual := captureHello(t)
	got, err := s.Sniff(actual)
	if err != nil || got.Value() != "Mixed.Example.test" {
		t.Fatal(got, err)
	}
	for i := 0; i < len(actual); i++ {
		if _, err := s.Sniff(actual[:i]); err != errorAt50983A0 {
			t.Fatalf("truncation %d: %v", i, err)
		}
	}
	for _, name := range []string{"", "Upper.Example"} {
		got, err := s.Sniff(syntheticHello(name))
		if err != nil || got.Value() != name {
			t.Fatal(name, got, err)
		}
	}
	if _, err := s.Sniff(syntheticHello("bad.test.")); err != errorAt50983B0 {
		t.Fatal(err)
	}
	malformed := syntheticHello("good.test")
	malformed[5] = 2 // 原函数不检查 ClientHello 类型。
	if _, err := s.Sniff(malformed); err != nil {
		t.Fatal("added unsupported handshake-type check", err)
	}
	malformed[0] = 23
	if _, err := s.Sniff(malformed); err != errorAt50983C0 {
		t.Fatal(err)
	}
}
