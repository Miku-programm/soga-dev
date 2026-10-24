package proxyproto

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
	"strings"
	"testing"
)

func TestOriginalHeaderConstructionAndFormatting(t *testing.T) {
	a := &net.TCPAddr{IP: net.IP{192, 0, 2, 1}, Port: 123}
	b := &net.TCPAddr{IP: net.IP{198, 51, 100, 2}, Port: 443}
	h := X_WkJPwas4L(0, a, b)
	if h.TmZyJAm != 2 || h.UCYj4Sj65ui != 0x21 || h.Dafxa_BOdZdw != a {
		t.Fatal("original To16 precedence", h)
	}
	raw, err := h.Format()
	if err != nil {
		t.Fatal(err)
	}
	got, err := parseFixture(t, raw)
	if err != nil || !got.EqualsTo(h) {
		t.Fatal(got, err)
	}
	h.TmZyJAm = 1
	h.UCYj4Sj65ui = 0x11
	raw, err = h.Format()
	if err != nil || string(raw) != "PROXY TCP4 192.0.2.1 198.51.100.2 123 443\r\n" {
		t.Fatal(string(raw), err)
	}
	h.UCYj4Sj65ui = 0x12
	raw, err = h.Format()
	if err != nil || string(raw) != "PROXY UNKNOWN\r\n" {
		t.Fatal(string(raw), err)
	}
	h.UCYj4Sj65ui = 0x11
	h.TmZyJAm = 2
	a.Port = -1
	b.Port = 65536
	raw, err = h.Format()
	if err != nil || binary.BigEndian.Uint16(raw[24:26]) != 65535 || binary.BigEndian.Uint16(raw[26:28]) != 0 {
		t.Fatal("original truncates ports", raw, err)
	}
	h.TmZyJAm = 7
	if _, err = h.Format(); err != ErrorAt5097E90 {
		t.Fatal(err)
	}
	u := X_WkJPwas4L(2, &net.UnixAddr{Name: "a", Net: "unixgram"}, &net.UnixAddr{Name: "b", Net: "unixgram"})
	if u.UCYj4Sj65ui != 0x32 {
		t.Fatal(u)
	}
	u.SetTLVs([]M46tt8Zs{{Be8_QISO: 1, J8xIVW: []byte{7}}})
	raw, err = u.Format()
	if err != nil || binary.BigEndian.Uint16(raw[14:16]) != 216 || len(raw) != 16+216+4 {
		t.Fatal("original Unix length omits appended TLVs", len(raw), err)
	}
	h = &URdA97qPFc{TmZyJAm: 2, LgI8Cr1feIYb: 0x20, ttF0mSAdaYG: make([]byte, 65536)}
	if _, err = h.Format(); err != ErrorAt5097F30 {
		t.Fatal(err)
	}
}

type shortProxyWriter struct{}

func (shortProxyWriter) Write(p []byte) (int, error) { return len(p) - 1, nil }
func TestOriginalHeaderWriteAndTLVIsolation(t *testing.T) {
	h := &URdA97qPFc{TmZyJAm: 1}
	if _, err := h.WriteTo(shortProxyWriter{}); err != io.ErrShortWrite {
		t.Fatal(err)
	}
	raw := []byte{1, 0, 2, 8, 9, 4, 0, 3, 1, 2, 3}
	r, err := FUahHJe3a(raw)
	if err != nil || len(r) != 2 || r[1].J8xIVW != nil {
		t.Fatal(r, err)
	}
	raw[3] = 0
	if !bytes.Equal(r[0].J8xIVW, []byte{8, 9}) {
		t.Fatal("TLVs must copy values")
	}
	joined, err := I_BfofJ(r)
	if err != nil || !bytes.Equal(joined, []byte{1, 0, 2, 8, 9, 4, 0, 0}) {
		t.Fatal(joined, err)
	}
	h.ttF0mSAdaYG = []byte{9}
	err = h.SetTLVs([]M46tt8Zs{{Be8_QISO: 1, J8xIVW: make([]byte, 65536)}})
	if err == nil || !strings.Contains(err.Error(), "length 65536") || !bytes.Equal(h.ttF0mSAdaYG, []byte{9}) {
		t.Fatal(err)
	}
	for _, b := range [][]byte{{1}, {1, 0}, {1, 0, 2, 1}} {
		if r, err := FUahHJe3a(b); r != nil || err != ErrorAt5097F00 {
			t.Fatal(r, err)
		}
	}
	for i := 0; i < 256; i++ {
		p := NMEUUf9(i)
		want := i >= 1 && i <= 5 || i >= 0x20 && i <= 0x25 || i == 0x30 || i >= 0xe0
		if p.Spec() != want {
			t.Fatal(i)
		}
	}
}
