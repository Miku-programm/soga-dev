package sstypes

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net"
	S "recovered.local/soga/util/auth"
	"testing"
	"unsafe"
)

type captureWriter struct {
	writes [][]byte
	failAt int
	count  int
	err    error
}

func (w *captureWriter) Write(p []byte) (int, error) {
	w.writes = append(w.writes, append([]byte(nil), p...))
	if w.failAt == len(w.writes) {
		return w.count, w.err
	}
	return len(p), nil
}

type readConn struct {
	net.Conn
	io.Reader
}

func (c *readConn) Read(b []byte) (int, error) { return c.Reader.Read(b) }

func TestClientLayoutAndHeader(t *testing.T) {
	if unsafe.Sizeof(LckgbWP0Vg4{}) != 64 || unsafe.Sizeof(MrjKETeSu{}) != 32 || unsafe.Sizeof(KTJbqZyG_iM{}) != 56 || unsafe.Sizeof(Mah1Kab{}) != 48 {
		t.Fatal("RTTI layout")
	}
	config := &LckgbWP0Vg4{HQxCy9slSH: "password"}
	c := NewClient(config)
	digest := sha256.Sum224([]byte("password"))
	want := hex.EncodeToString(digest[:])
	if string(c._mr1q5jad7H) != want || cap(c._mr1q5jad7H) != 56 || c.rTtnkhQSU != config {
		t.Fatal(c)
	}
	addr := S.L4JRaRodkw("192.0.2.1:80")
	var out bytes.Buffer
	if err := c.WriteHeader(&out, 1, addr); err != nil {
		t.Fatal(err)
	}
	expected := append(append([]byte(want+"\r\n\x01"), addr...), []byte("\r\n")...)
	if !bytes.Equal(out.Bytes(), expected) {
		t.Fatal(out.Bytes())
	}
	w := &captureWriter{failAt: 1, count: 0}
	if err := c.WriteHeader(w, 1, addr); err != nil {
		t.Fatal("short write was checked", err)
	}
}

func TestPacketFramingCountsAndErrors(t *testing.T) {
	addr := S.L4JRaRodkw("192.0.2.1:53")
	w := &captureWriter{}
	if n, err := AwkQMyI(w, addr, []byte("x")); err != nil || n != 12 {
		t.Fatal(n, err)
	}
	if !bytes.Equal(w.writes[0], append(append(append([]byte{}, addr...), 0, 1, '\r', '\n'), 'x')) {
		t.Fatal(w.writes)
	}
	w = &captureWriter{}
	if n, err := AwkQMyI(w, addr, make([]byte, 8193)); err != nil || n != 8193 || len(w.writes) != 2 || len(w.writes[0]) != 8203 || len(w.writes[1]) != 12 {
		t.Fatal(n, err, len(w.writes))
	}
	failure := errors.New("write")
	w = &captureWriter{failAt: 2, count: 3, err: failure}
	if n, err := AwkQMyI(w, addr, make([]byte, 8193)); err != failure || n != 8195 {
		t.Fatal(n, err)
	}
}

func TestPacketReadStagesAndRemainder(t *testing.T) {
	addr := S.L4JRaRodkw("192.0.2.1:53")
	makeFrame := func(length int, delimiter, payload []byte) []byte {
		b := append([]byte{}, addr...)
		b = append(b, byte(length>>8), byte(length))
		b = append(b, delimiter...)
		return append(b, payload...)
	}
	for _, test := range []struct {
		frame []byte
		want  string
	}{
		{nil, "read addr error"},
		{S.L4JRaRodkw("host:9"), "parse addr error"},
		{addr, "read length error"},
		{makeFrame(8193, nil, nil), "packet invalid"},
		{makeFrame(1, nil, nil), "read crlf error"},
		{makeFrame(1, []byte("\r\n"), nil), "read packet error"},
	} {
		_, n, left, err := WroRKupgG3k7(bytes.NewReader(test.frame), make([]byte, 259))
		if err == nil || err.Error() != test.want || n != 0 || left != 0 {
			t.Fatal(n, left, err, test.want)
		}
	}
	// The delimiter's bytes are consumed but never validated.
	frame := makeFrame(500, []byte("XY"), bytes.Repeat([]byte{'a'}, 500))
	c := &Mah1Kab{Conn: &readConn{Reader: bytes.NewReader(frame)}}
	b := make([]byte, 259)
	n, peer, err := c.ReadFrom(b)
	if err != nil || n != 259 || peer.String() != "192.0.2.1:53" || c.iFfzJtf3G != 241 || b[0] != 'a' {
		t.Fatal(n, peer, err, c.iFfzJtf3G)
	}
	n, peer, err = c.ReadFrom(b)
	if err != nil || n != 241 || peer.String() != "192.0.2.1:53" || c.iFfzJtf3G != 0 || c.bsHi8N0rMT != nil {
		t.Fatal(n, peer, err, c.iFfzJtf3G)
	}
	if _, _, _, err := WroRKupgG3k7(bytes.NewReader(frame), make([]byte, 258)); err == nil || err.Error() != "read addr error" {
		t.Fatal(err)
	}
}
