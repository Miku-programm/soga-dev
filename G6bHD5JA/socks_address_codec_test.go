package G6bHD5JA

import (
	"bytes"
	"io"
	"net"
	"recovered.local/soga/D_zCag1qTV"
	"strings"
	"testing"
)

func TestOriginalSocksAddressCodec(t *testing.T) {
	for _, tt := range []struct {
		wire []byte
		host string
		port int
		ip   bool
	}{
		{[]byte{1, 192, 0, 2, 1, 0, 80}, "192.0.2.1", 80, true},
		{append(append([]byte{4}, net.ParseIP("2001:db8::2").To16()...), 1, 187), "2001:db8::2", 443, true},
		{append(append([]byte{3, 9}, []byte("127.0.0.1")...), 0, 80), "127.0.0.1", 80, false},
		{[]byte{3, 0, 255, 255}, "", 65535, false},
	} {
		a, err := SzhW4tOzUZg(bytes.NewReader(tt.wire))
		if err != nil || a.GetHost() != tt.host || a.GetPort() != tt.port || a.IsIP() != tt.ip {
			t.Fatal(a, err)
		}
		b := D_zCag1qTV.LdHEmSs9(32)
		n, err := DdYroT1n(&b, a)
		if err != nil || n != len(tt.wire) || !bytes.Equal(b.Bytes(), tt.wire) {
			t.Fatal("wire", n, err, b.Bytes())
		}
		b.Free()
		for i := 0; i < len(tt.wire); i++ {
			_, err = SzhW4tOzUZg(bytes.NewReader(tt.wire[:i]))
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				t.Fatal("truncation", i, err)
			}
		}
	}
	_, err := SzhW4tOzUZg(bytes.NewReader([]byte{2}))
	if err == nil || err.Error() != "unknown addr type: 2" {
		t.Fatal(err)
	}
	b := D_zCag1qTV.LdHEmSs9(32)
	defer b.Free()
	if n, err := DdYroT1n(&b, &AHU2MRmjzQu{dlcy_1: strings.Repeat("x", 256)}); err == nil || n != 0 || b.Len() != 0 {
		t.Fatal("domain bound")
	}
	if n, err := DdYroT1n(&b, &AHU2MRmjzQu{oKcMaQDwbH: []net.IP{{1, 2}}, pDnAz_: 65536}); err != nil || n != 19 || !bytes.Equal(b.Bytes(), []byte{4, 0, 0}) {
		t.Fatal("original abnormal IP/port", n, err, b.Bytes())
	}
}
