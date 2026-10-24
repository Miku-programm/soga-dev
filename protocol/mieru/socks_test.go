package mieru

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"strings"
	"testing"
)

func TestSocksRequestOriginalBranches(t *testing.T) {
	for _, tc := range []struct {
		wire []byte
		host string
		cmd  byte
		err  string
	}{
		{[]byte{5, 1, 99, 1, 127, 0, 0, 1, 0, 80}, "127.0.0.1", 1, ""},
		{append([]byte{5, 3, 0, 3, 9}, append([]byte("127.0.0.1"), 0, 53)...), "127.0.0.1", 3, ""},
		{append([]byte{5, 1, 0, 3, 3}, append([]byte("dns"), 0, 53)...), "dns", 1, ""},
		{[]byte{5, 1, 0, 3, 0, 0, 0}, "<nil>", 1, ""},
		{[]byte{4, 1, 0, 9}, "", 0, "unsupported socks version: 4"},
		{[]byte{5, 3, 0, 9}, "", 3, "unsupported atyp: 9"},
		{[]byte{5, 1, 0, 3}, "", 0, "read fqdn len: EOF"},
		{[]byte{5, 1, 0, 1, 1, 2}, "", 0, "read ipv4: unexpected EOF"},
		{[]byte{5, 1, 0, 1, 1, 2, 3, 4}, "", 0, "read port: EOF"},
	} {
		cmd, a, e := dOUoamsBibCa(bytes.NewReader(tc.wire))
		if cmd != tc.cmd {
			t.Fatal(cmd, tc)
		}
		if tc.err != "" {
			if e == nil || e.Error() != tc.err {
				t.Fatal(e, tc)
			}
			continue
		}
		if e != nil || a.GetHost() != tc.host {
			t.Fatal(a, e, tc)
		}
	}
}
func TestUDPParserErrorPrecedence(t *testing.T) {
	cases := []struct {
		b    []byte
		want string
	}{
		{[]byte{1, 2, 3}, "socks5 udp datagram too short: 3"},
		{[]byte{1, 2, 3, 9, 0, 0}, "socks5 udp datagram RSV not zero: 1 2"},
		{[]byte{0, 0, 3, 9, 0, 0}, "socks5 udp fragmentation not supported (FRAG=3)"},
		{[]byte{0, 0, 0, 1, 0, 0}, "socks5 udp ipv4 short: 6"},
		{[]byte{0, 0, 0, 4, 0, 0}, "socks5 udp ipv6 short: 6"},
		{[]byte{0, 0, 0, 3, 2, 0}, "socks5 udp fqdn body short: need 9, have 6"},
		{[]byte{0, 0, 0, 9, 0, 0}, "socks5 udp unsupported ATYP 9"},
	}
	for _, tc := range cases {
		_, _, e := rNsudCbv(tc.b)
		if e == nil || e.Error() != tc.want {
			t.Fatal(e, tc.want)
		}
	}
}
func TestUDPOverStreamWire(t *testing.T) {
	for _, ip := range []net.IP{net.IP{127, 0, 0, 1}, net.ParseIP("2001:db8::1")} {
		a := testFrameAEAD(t)
		conn := new(frameMemoryConn)
		writer, e := vfhZtLu(conn, a, "fixture")
		if e != nil {
			t.Fatal(e)
		}
		h := &h7SOFexGyw{vnqzugKQF1jX(42, nil, conn, writer)}
		payload := []byte("udp payload")
		n, e := h.WriteToProxy(payload, &net.UDPAddr{IP: ip, Port: 5353})
		if e != nil || n != len(payload) {
			t.Fatal(n, e)
		}
		wire := conn.Bytes()
		nonce := append([]byte(nil), wire[:24]...)
		raw, e := a.Open(nil, nonce, wire[24:72], nil)
		if e != nil {
			t.Fatal(e)
		}
		fixtureNext(nonce)
		prefix := int(raw[21])
		length := int(binary.BigEndian.Uint16(raw[22:24]))
		body, e := a.Open(nil, nonce, wire[72+prefix:72+prefix+length+16], nil)
		if e != nil {
			t.Fatal(e)
		}
		wantHeader := 10
		if ip.To4() == nil {
			wantHeader = 22
		}
		if body[0] != 0 || int(binary.BigEndian.Uint16(body[1:3])) != wantHeader+len(payload) || body[len(body)-1] != 255 {
			t.Fatalf("wrapper %x", body)
		}
		if !bytes.Equal(body[3+wantHeader:len(body)-1], payload) {
			t.Fatalf("payload %x", body)
		}
		receiving := &h7SOFexGyw{vnqzugKQF1jX(43, nil, new(frameMemoryConn), nil)}
		receiving.xfdnp9xA <- append([]byte(nil), body[:4]...)
		receiving.xfdnp9xA <- append([]byte(nil), body[4:]...)
		b := make([]byte, 256)
		n, addr, e := receiving.ReadFromProxy(b)
		if e != nil || !bytes.Equal(b[:n], payload) || !addr.(*net.UDPAddr).IP.Equal(ip) || addr.(*net.UDPAddr).Port != 5353 {
			t.Fatal(n, addr, e)
		}
		receiving.Free()
	}
}
func TestUDPShortAndBadWrapper(t *testing.T) {
	for _, tc := range []struct {
		b    []byte
		size int
		want string
	}{
		{[]byte{1, 0, 0}, 20, "packet prefix 0x1 is not 0x00"},
		{[]byte{0, 0, 25}, 20, "packet length 25 exceeds buffer 20"},
		{[]byte{0, 0, 0, 0}, 20, "packet suffix 0x0 is not 0xFF"},
		{[]byte{0, 0, 0, 255}, 20, "socks5 udp datagram too short: 0"},
	} {
		h := &h7SOFexGyw{vnqzugKQF1jX(1, nil, new(frameMemoryConn), nil)}
		h.xfdnp9xA <- tc.b
		_, _, e := h.ReadFromProxy(make([]byte, tc.size))
		if e == nil || e.Error() != tc.want {
			t.Fatal(e, tc.want)
		}
		h.Free()
	}
	_, _, e := dOUoamsBibCa(strings.NewReader("\x05\x01"))
	if !errors.Is(e, io.ErrUnexpectedEOF) {
		t.Fatal(e)
	}
}
