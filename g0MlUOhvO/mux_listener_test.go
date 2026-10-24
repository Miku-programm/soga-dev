package g0MlUOhvO

import (
	"bytes"
	"io"
	"net"
	"recovered.local/soga/DDYEn9Z_b"
	"recovered.local/soga/I4AWsAkBLRP"
	"reflect"
	"testing"
	"time"
)

func TestOriginalEncryptedMuxListenerTCP(t *testing.T) {
	typ := reflect.TypeOf(PRSLFMd{})
	if typ.Size() != 56 {
		t.Fatal(typ.Size())
	}
	for i, want := range []uintptr{0, 16, 32, 40, 48} {
		if typ.Field(i).Offset != want {
			t.Fatal(typ.Field(i))
		}
	}
	method, err := DDYEn9Z_b.MakT9YbIpyNJ("aes-128-gcm")
	if err != nil {
		t.Fatal(err)
	}
	cfg := &VTl94pzz2IkB{zkn35IN: method, oWvGY4I9cA: bytes.Repeat([]byte{7}, 16)}
	raw, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	mux := Bn7A4E9vXH(&T4k2R7cJ{O6B2qZ0p0: raw, tS5gsM72s: cfg}, cfg)
	defer mux.Close()
	conn, err := net.Dial("tcp", raw.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	encrypted := &Cb7vj3KQA8{DrpXao: conn, hp3cHCehD2pe: cfg}
	client, err := I4AWsAkBLRP.CXFDjl_TYx8(encrypted, &I4AWsAkBLRP.BWpFHH3x{RTNOROJC5: 1, DU0iTVAg: true, XsqVEjw4: 32768, K3nkPQuOI: 4194304, CUuHLYfWnC8q: 65536})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	for i := range 4 {
		c, e := client.OpenStream()
		if e != nil {
			t.Fatal(e)
		}
		s, e := mux.Accept()
		if e != nil {
			t.Fatal(e)
		}
		data := bytes.Repeat([]byte{byte(i), 0, 255}, 21000)
		done := make(chan error, 1)
		go func() { _, e := c.Write(data); done <- e }()
		s.SetReadDeadline(time.Now().Add(3 * time.Second))
		b := make([]byte, len(data))
		if _, e = io.ReadFull(s, b); e != nil || !bytes.Equal(b, data) {
			t.Fatalf("stream %d: %v", i, e)
		}
		if e = <-done; e != nil {
			t.Fatal(e)
		}
		go func() { _, e := s.Write(data[:31000]); done <- e }()
		c.SetReadDeadline(time.Now().Add(3 * time.Second))
		b = b[:31000]
		if _, e = io.ReadFull(c, b); e != nil || !bytes.Equal(b, data[:31000]) {
			t.Fatal("reply", e)
		}
		if e = <-done; e != nil {
			t.Fatal(e)
		}
		c.Close()
		s.Close()
	}
	mux.Close()
	if c, e := mux.Accept(); c != nil || e == nil || e.Error() != "closed" {
		t.Fatalf("closed accept %v %v", c, e)
	}
}
