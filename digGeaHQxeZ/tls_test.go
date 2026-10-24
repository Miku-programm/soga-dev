package digGeaHQxeZ

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"testing"
	"time"
	"unsafe"
)

func prefixed16(p []byte) []byte { return append([]byte{byte(len(p) >> 8), byte(len(p))}, p...) }
func helloFixture(host string, ticket []byte) []byte {
	p := append([]byte{1, 0, 0, 0, 3, 3}, make([]byte, 32)...)
	p = append(p, 32)
	p = append(p, bytes.Repeat([]byte{0x55}, 32)...)
	p = append(p, 0, 2, 0xcc, 0xa8, 1, 0)
	names := prefixed16(append([]byte{0}, prefixed16([]byte(host))...))
	ext := append([]byte{0, 0}, prefixed16(names)...)
	ext = append(ext, 0, 35)
	ext = append(ext, prefixed16(ticket)...)
	p = append(p, prefixed16(ext)...)
	return p
}
func TestNativeTLSHelloParserBoundaries(t *testing.T) {
	p := helloFixture("example.org", []byte("first"))
	session, host, payload, e := vYzlYln7glEA(p)
	if e != nil || len(session) != 32 || host != "example.org" || string(payload) != "first" {
		t.Fatal(session, host, payload, e)
	}
	for i := 0; i < len(p); i++ {
		if _, _, _, e = vYzlYln7glEA(p[:i]); e != io.ErrUnexpectedEOF {
			t.Fatalf("truncation %d %v", i, e)
		}
	}
	bad := append([]byte(nil), p...)
	bad[0] = 2
	if _, _, _, e = vYzlYln7glEA(bad); e == nil || e.Error() != "not a tls client hello: 2" {
		t.Fatal(e)
	}
	// Invalid nested SNI is silently ignored; valid extension framing remains accepted.
	if got := jHOgG9GS([]byte{0, 3, 0, 0, 9}); got != "" {
		t.Fatal(got)
	}
	// Declared handshake length and protocol version are intentionally not checked.
	p[1] = 0xff
	p[4] = 0xff
	if _, _, _, e = vYzlYln7glEA(p); e != nil {
		t.Fatal(e)
	}
}
func TestNativeTLSWireReadsAndServerHello(t *testing.T) {
	var v IkPl_ual
	if unsafe.Sizeof(v) != 80 || unsafe.Offsetof(v.vGM3bgb8VDgL) != 56 {
		t.Fatal("RTTI")
	}
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()
	client.SetDeadline(time.Now().Add(3 * time.Second))
	server.SetDeadline(time.Now().Add(3 * time.Second))
	c := NewTLSForFactory(server, "example.org", 1)
	done := make(chan error, 1)
	response := bytes.Repeat([]byte{0x66}, 20000)
	go func() {
		p := make([]byte, 5)
		if _, e := io.ReadFull(c, p); e != nil || string(p) != "first" {
			done <- fmt.Errorf("ticket payload %v %q", e, p)
			return
		}
		p = make([]byte, 4)
		if _, e := io.ReadFull(c, p); e != nil || string(p) != "next" {
			done <- fmt.Errorf("record %v %q", e, p)
			return
		}
		if n, e := c.Write(response); e != nil || n != len(response) {
			done <- fmt.Errorf("response %d %v", n, e)
			return
		}
		done <- nil
	}()
	hello := helloFixture("example.org", []byte("first"))
	wire := append([]byte{22, 3, 1}, prefixed16(hello)...)
	wire = append(wire, 23, 3, 9, 0, 0, 23, 3, 9, 0, 4)
	wire = append(wire, []byte("next")...)
	if _, e := client.Write(wire); e != nil {
		t.Fatal(e)
	}
	h := make([]byte, 96)
	if _, e := io.ReadFull(client, h); e != nil {
		t.Fatal(e)
	}
	if !bytes.Equal(h[:9], []byte{22, 3, 3, 0, 91, 2, 0, 0, 87}) || !bytes.Equal(h[44:76], bytes.Repeat([]byte{0x55}, 32)) || !bytes.Equal(h[76:], helloTailAt509A750) {
		t.Fatal("server hello", h)
	}
	change := make([]byte, 6)
	io.ReadFull(client, change)
	if !bytes.Equal(change, changeAt503B8B0) {
		t.Fatal("change cipher", change)
	}
	var got []byte
	for _, kind := range []byte{22, 23} {
		header := make([]byte, 5)
		if _, e := io.ReadFull(client, header); e != nil {
			t.Fatal(e)
		}
		if header[0] != kind {
			t.Fatal("record type", header)
		}
		p := make([]byte, int(binary.BigEndian.Uint16(header[3:])))
		if _, e := io.ReadFull(client, p); e != nil {
			t.Fatal(e)
		}
		got = append(got, p...)
	}
	if !bytes.Equal(got, response) {
		t.Fatal("response body")
	}
	if e := <-done; e != nil {
		t.Fatal(e)
	}
}
func TestNativeTLSRejectsHostAndRecord(t *testing.T) {
	for _, tc := range []struct {
		host string
		wire []byte
		want string
	}{
		{"example.org", []byte{22, 3, 3, 0, 0}, "empty tls handshake record"},
		{"example.org", []byte{23, 3, 3, 0, 0}, "not a tls handshake record: 170303"},
		{"different.test", append([]byte{22, 3, 3}, prefixed16(helloFixture("example.org", nil))...), "client host (example.org) not match server host (different.test)"},
	} {
		a, b := net.Pipe()
		a.SetDeadline(time.Now().Add(time.Second))
		b.SetDeadline(time.Now().Add(time.Second))
		done := make(chan error, 1)
		go func() { c := NewTLSForFactory(b, tc.host, 3); _, e := c.Read(make([]byte, 1)); done <- e }()
		a.Write(tc.wire)
		e := <-done
		a.Close()
		b.Close()
		if e == nil || e.Error() != tc.want {
			t.Fatal(e, tc.want)
		}
	}
}
