package G6bHD5JA

import (
	"io"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestAddressOriginalSemantics(t *testing.T) {
	for _, c := range []struct {
		input, host, join, display string
		port, size                 int
		literal                    bool
	}{
		{"example.com:80", "example.com", "example.com:80", "example.com:80", 80, 15, false},
		{"[2001:0db8::1]:443", "2001:db8::1", "[2001:db8::1]:443", "2001:db8::1:443", 443, 19, true},
		{"[::ffff:192.0.2.1]:7", "192.0.2.1", "192.0.2.1:7", "192.0.2.1:7", 7, 7, true},
		{"host:-1", "host", "host:-1", "host:-1", -1, 8, false},
		{"host:65536", "host", "host:65536", "host:65536", 65536, 8, false},
		{"[fe80::1%eth0]:9", "fe80::1%eth0", "[fe80::1%eth0]:9", "fe80::1%eth0:9", 9, 16, false},
	} {
		a, err := U80FDt(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if a.GetHost() != c.host || a.GetPort() != c.port || a.Join() != c.join || a.HostPort() != c.join || a.String() != c.display || a.Size() != c.size || a.IsIP() != c.literal {
			t.Fatalf("%s: %+v size=%d", c.input, a, a.Size())
		}
	}
	for _, s := range []string{"missing-port", "example.com:https", "x:"} {
		if _, err := U80FDt(s); err == nil {
			t.Fatal(s)
		}
	}
	ip := net.IPv4(192, 0, 2, 1)
	a, err := RNFJeaRY(&net.TCPAddr{IP: ip, Port: 8, Zone: "discarded"})
	if err != nil {
		t.Fatal(err)
	}
	a.SetHost("new.example")
	if a.IsIP() || !a.HasIP() || a.GetHost() != "new.example" {
		t.Fatal(a)
	}
	if a.String() != "new.example([192.0.2.1]):8" {
		t.Fatal(a.String())
	}
	if err = a.ResolveHost(nil); err != nil {
		t.Fatal(err)
	} // 原缓存分支不访问 resolver。
	tcp, err := a.ToTCPAddrs()
	if err != nil || len(tcp) != 1 || tcp[0].Zone != "" {
		t.Fatalf("%v %v", tcp, err)
	}
	tcp[0].IP[15] = 9
	if a.GetFirstIp()[15] != 9 || ip[15] != 9 {
		t.Fatal("IP backing slice copied")
	}
	udp, err := a.ToUDPAddr()
	if err != nil || udp.Port != 8 {
		t.Fatal(udp, err)
	}
	b, _ := U80FDt("new.example:8")
	if !a.EqualTo(b) {
		t.Fatal("equality unexpectedly compares IP cache")
	}
	if _, err = b.ToTCPAddrs(); err == nil || err.Error() != "addr is not ip" {
		t.Fatal(err)
	}
	if _, err = b.ToUDPAddr(); err == nil || err.Error() != "addr is not ip" {
		t.Fatal(err)
	}
	nilIP, err := RNFJeaRY(&net.UDPAddr{Port: 3})
	if err != nil || !nilIP.HasIP() || nilIP.GetFirstIp() != nil || nilIP.Size() != 19 {
		t.Fatal(nilIP, err)
	}
	if _, err = RNFJeaRY(nil); err == nil || err.Error() != "unknown addr: <nil>" {
		t.Fatal(err)
	}
	typ := reflect.TypeOf(AHU2MRmjzQu{})
	if typ.Size() != 56 {
		t.Fatal(typ.Size())
	}
	for name, offset := range map[string]uintptr{"dlcy_1": 0, "oKcMaQDwbH": 16, "pDnAz_": 40, "r_VdNaPA": 48} {
		field, _ := typ.FieldByName(name)
		if field.Offset != offset {
			t.Fatal(name, field.Offset)
		}
	}
}

func TestSourceAddressFamily(t *testing.T) {
	v4 := net.ParseIP("192.0.2.1")
	v6 := net.ParseIP("2001:db8::1")
	for _, c := range []struct {
		local, remote net.IP
		want          bool
	}{{nil, nil, true}, {net.IPv4zero, v6, true}, {net.IPv6zero, v4, true}, {v4, v6, false}, {v6, v4, false}, {v4, v4, true}, {v6, nil, false}, {net.IP{}, net.IP{}, true}, {net.IP{}, nil, false}} {
		if got := Y0B53iB(c.local, c.remote); got != c.want {
			t.Fatalf("%v %v got %v", c.local, c.remote, got)
		}
	}
}

func TestOriginalDirectDialLoopback(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	result := make(chan error, 1)
	go func() {
		conn, e := listener.Accept()
		if e != nil {
			result <- e
			return
		}
		defer conn.Close()
		_, e = conn.Write([]byte("ok"))
		result <- e
	}()
	conn, err := DE9CHPNLbK(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1)}, listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Second))
	b := make([]byte, 2)
	if _, err = io.ReadFull(conn, b); err != nil || string(b) != "ok" {
		t.Fatal(string(b), err)
	}
	if err = <-result; err != nil {
		t.Fatal(err)
	}
}
