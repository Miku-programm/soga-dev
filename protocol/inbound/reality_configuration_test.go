package inbound

import (
	"context"
	"encoding/base64"
	"io"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestRealityConfigurationOriginal(t *testing.T) {
	key := base64.RawURLEncoding.EncodeToString(make([]byte, 32))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := fwKvAvIIe(ctx, []string{"", "a", "a"}, key, "127.0.0.1:1", []string{"", "abcd", "abcd000000000000"})
	if err != nil || cfg.Type != "tcp" || cfg.Dest != "127.0.0.1:1" || !cfg.SessionTicketsDisabled || len(cfg.ShortIds) != 2 || len(cfg.ServerNames) != 2 || cfg.DialContext == nil {
		t.Fatalf("%+v %v", cfg, err)
	}
	cases := []struct {
		names     []string
		key, dest string
		ids       []string
		want      string
	}{
		{nil, "", "", nil, "reality short id 不能为空"},
		{nil, "", "", []string{"0"}, "reality short id 格式错误: 0:"},
		{nil, "", "", []string{strings.Repeat("ff", 9)}, "reality short id 超过 8 字节:"},
		{nil, "!", "", []string{""}, "reality private key 格式错误:"},
		{nil, "", "", []string{""}, "实际 0 字节"},
		{nil, key, "", []string{""}, "reality sni (serverNames) 不能为空"},
		{[]string{"a"}, key, "", []string{""}, "reality dest 不能为空"},
	}
	for _, c := range cases {
		got, e := fwKvAvIIe(ctx, c.names, c.key, c.dest, c.ids)
		if got != nil || e == nil || !strings.Contains(e.Error(), c.want) {
			t.Fatalf("%+v: %v", c, e)
		}
	}
	var x rWW6Rmakjzl
	var r fsqq14xvr
	if unsafe.Sizeof(x) != 184 || unsafe.Offsetof(x.rqmdTVGqu) != 32 || unsafe.Offsetof(x.sDhm_cyQ1) != 176 || unsafe.Sizeof(r) != 32 {
		t.Fatal("RTTI layout")
	}
}
func TestRealityDestStablePartition(t *testing.T) {
	a, b, c := net.ParseIP("::1"), net.ParseIP("127.0.0.1"), net.ParseIP("::2")
	x, y := eF7oFHXp9Xe([]net.IP{a, b, c})
	if !reflect.DeepEqual(x, []net.IP{a, c}) || !reflect.DeepEqual(y, []net.IP{b}) {
		t.Fatal(x, y)
	}
	x, y = eF7oFHXp9Xe([]net.IP{b, a, net.IPv4(1, 2, 3, 4)})
	if len(x) != 2 || len(y) != 1 {
		t.Fatal(x, y)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("empty original input must panic")
		}
	}()
	eF7oFHXp9Xe(nil)
}
func TestRealityCachedDialerActualTCP(t *testing.T) {
	l, e := net.Listen("tcp4", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer l.Close()
	_, port, _ := net.SplitHostPort(l.Addr().String())
	go func() {
		for {
			c, e := l.Accept()
			if e != nil {
				return
			}
			go func() { defer c.Close(); io.Copy(c, c) }()
		}
	}()
	for _, ips := range [][]net.IP{nil, {net.ParseIP("127.0.0.2"), net.ParseIP("127.0.0.1")}, {net.ParseIP("::1"), net.ParseIP("127.0.0.1")}, {net.ParseIP("127.0.0.1"), net.ParseIP("::1")}} {
		d := &rWW6Rmakjzl{af0cE6r5P: "ignored", tkRa6u: port, rqmdTVGqu: net.Dialer{Timeout: time.Second}}
		if ips != nil {
			d.sDhm_cyQ1.Store(&ips)
		}
		c, e := d.DialContext(context.Background(), "tcp", l.Addr().String())
		if e != nil {
			t.Fatal(e)
		}
		c.SetDeadline(time.Now().Add(2 * time.Second))
		c.Write([]byte("real TCP"))
		b := make([]byte, 8)
		_, e = io.ReadFull(c, b)
		c.Close()
		if e != nil || string(b) != "real TCP" {
			t.Fatalf("%q %v", b, e)
		}
	}
}
