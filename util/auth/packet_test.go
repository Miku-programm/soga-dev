package auth

import (
	"bytes"
	"net"
	"net/netip"
	"strings"
	"testing"
)

func TestWireAddressesAndMalformedInputs(t *testing.T) {
	for _, address := range []string{"bad", "host:-1", "host:65536", strings.Repeat("a", 256) + ":1"} {
		if got := L4JRaRodkw(address); got != nil {
			t.Fatal(address, got)
		}
	}
	for _, test := range []struct {
		in  string
		out []byte
	}{
		{"192.0.2.1:53", []byte{1, 192, 0, 2, 1, 0, 53}},
		{"host:443", []byte{3, 4, 'h', 'o', 's', 't', 1, 187}},
		{":0", []byte{3, 0, 0, 0}},
	} {
		if got := L4JRaRodkw(test.in); !bytes.Equal(got, test.out) {
			t.Fatal(test.in, got)
		}
	}
	// The net.Addr path deliberately preserves the original DefaultMask test.
	if got := KvyKMZ(&net.UDPAddr{IP: net.IP{224, 0, 0, 1}, Port: -1}); !bytes.Equal(got, []byte{1, 224, 0, 0, 1, 255, 255}) {
		t.Fatal(got)
	}
	if got := KvyKMZ(&net.TCPAddr{IP: net.IP{}, Port: 9}); !bytes.Equal(got, []byte{4, 0, 9}) {
		t.Fatal(got)
	}
	if got := Va09RL21(netip.Addr{}, 7); len(got) != 19 || got[0] != 4 || got[18] != 7 {
		t.Fatal(got)
	}
	if got := Va09RL21(netip.MustParseAddr("::ffff:192.0.2.1"), 7); len(got) != 19 || got[0] != 4 {
		t.Fatal(got)
	}
}

func TestUDPEncodingDecodingAndAliases(t *testing.T) {
	address := L4JRaRodkw("host:80")
	packet, err := XFq85XFER(address, []byte("abc"))
	if err != nil {
		t.Fatal(err)
	}
	a, p, err := CkUmXXHwOEH(packet)
	if err != nil || !bytes.Equal(a, address) || string(p) != "abc" {
		t.Fatal(a, p, err)
	}
	a[0] = 9
	p[0] = 'x'
	if packet[3] != 9 || packet[len(packet)-3] != 'x' {
		t.Fatal("lost alias")
	}
	for _, test := range []struct {
		packet []byte
		want   string
		keep   bool
	}{
		{[]byte{0, 0, 0, 1}, "insufficient length of packet", false},
		{[]byte{1, 0, 0, 1, 0}, "reserved fields should be zero", false},
		{[]byte{0, 0, 1, 1, 0}, "discarding fragmented payload", false},
		{[]byte{0, 0, 0, 1, 0}, "failed to read UDP header", true},
		{[]byte{0, 0, 0, 3, 9}, "failed to read UDP header", true},
		{[]byte{0, 0, 0, 9, 0}, "failed to read UDP header", true},
	} {
		a, p, err := CkUmXXHwOEH(test.packet)
		if a != nil || err == nil || err.Error() != test.want {
			t.Fatal(a, p, err)
		}
		if test.keep {
			if &p[0] != &test.packet[3] {
				t.Fatal("lost error payload")
			}
		} else if p != nil {
			t.Fatal(p)
		}
	}
	if _, err := XFq85XFER(nil, nil); err == nil || err.Error() != "address is invalid" {
		t.Fatal(err)
	}
	if b, err := XFq85XFER([]byte{}, nil); err != nil || !bytes.Equal(b, []byte{0, 0, 0}) {
		t.Fatal(b, err)
	}
}
