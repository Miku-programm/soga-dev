package relay

import (
	"net"
	"net/http"
	"testing"
)

func TestOriginalForwardedAddressPrecedence(t *testing.T) {
	for _, tc := range []struct{ real, forwarded, expected string }{
		{"192.0.2.1", "198.51.100.1", "192.0.2.1"},
		{"", " 2001:db8::1 , 198.51.100.1", "2001:db8::1"},
		{"bad", "192.0.2.1", ""},
		{" 192.0.2.1 ", "192.0.2.2", ""},
		{"", "bad,192.0.2.1", ""},
		{"", ",192.0.2.1", ""},
		{"", "", ""},
	} {
		h := http.Header{}
		h.Set("X-Real-IP", tc.real)
		h.Set("X-Forwarded-For", tc.forwarded)
		got := SOmphOG1Ty(h)
		if tc.expected == "" {
			if got != nil {
				t.Fatal("unexpected address", tc, got)
			}
		} else if got == nil || !got.(*net.TCPAddr).IP.Equal(net.ParseIP(tc.expected)) || got.(*net.TCPAddr).Port != 0 {
			t.Fatal("wrong address", tc, got)
		}
	}
}
