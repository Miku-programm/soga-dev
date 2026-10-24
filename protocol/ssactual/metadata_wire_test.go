package ssactual

import (
	"encoding/hex"
	"net"
	E "recovered.local/soga/net/outbound"
	"strings"
	"testing"
)

func TestMetadataAddressEncoding(t *testing.T) {
	tests := []struct {
		m    E.D0mLgQF
		want string
	}{
		{E.D0mLgQF{ZouEpmjw8: net.ParseIP("192.0.2.1"), JRXMVvg: 443}, "01c000020101bb"},
		{E.D0mLgQF{WJ95LyOL: "a", JRXMVvg: 53}, "0301610035"},
		{E.D0mLgQF{}, "03000000"},
		{E.D0mLgQF{ZouEpmjw8: net.IP{}}, "040000"},
	}
	for _, test := range tests {
		if got := hex.EncodeToString(o2HGjDHIVO(&test.m)); got != test.want {
			t.Fatal(got, test.want)
		}
	}
	m := E.D0mLgQF{WJ95LyOL: strings.Repeat("a", 256)}
	got := o2HGjDHIVO(&m)
	if len(got) != 260 || got[1] != 0 {
		t.Fatal(len(got), got[1])
	}
}
