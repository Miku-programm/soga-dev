package E3F12urwswL

import "testing"

func TestAdapterTypeStrings(t *testing.T) {
	want := []string{"Direct", "Reject", "Shadowsocks", "ShadowsocksR", "Snell", "Socks5", "Http", "Vmess", "Trojan", "Relay", "Selector", "Fallback", "URLTest", "LoadBalance"}
	for i, s := range want {
		if ASFyBDZD(i).String() != s {
			t.Fatalf("%d", i)
		}
	}
	for _, i := range []int{-1, 14, 256, 1 << 32} {
		if ASFyBDZD(i).String() != "Unknown" {
			t.Fatalf("%d", i)
		}
	}
}
