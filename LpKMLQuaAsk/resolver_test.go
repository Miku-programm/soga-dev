package LpKMLQuaAsk

import (
	"context"
	"errors"
	"net"
	"testing"
)

type testResolver struct {
	call string
	ctx  context.Context
	ips  []net.IP
	err  error
}

func (r *testResolver) LookupAll(c context.Context, h string) ([]net.IP, error) {
	r.call = "all"
	r.ctx = c
	return r.ips, r.err
}
func (r *testResolver) LookupIPv4(c context.Context, h string) ([]net.IP, error) {
	r.call = "v4"
	r.ctx = c
	return r.ips, r.err
}
func (r *testResolver) LookupIPv6(c context.Context, h string) ([]net.IP, error) {
	r.call = "v6"
	r.ctx = c
	return r.ips, r.err
}
func TestOriginalDefaultsAndLiteralFamilies(t *testing.T) {
	if !ipv6DisabledAt3787310 || timeoutAt3787948 != 5e9 {
		t.Fatal("defaults")
	}
	ips, err := IYMzyr8f(context.Background(), "192.0.2.1")
	if err != nil || len(ips) != 1 || len(ips[0]) != 16 {
		t.Fatal(ips, err)
	}
	if _, err = IYMzyr8f(context.Background(), "::ffff:192.0.2.1"); err != versionErrorAt5097950 {
		t.Fatal(err)
	}
	if _, err = Jqw3ag(context.Background(), "::1"); err != ipv6DisabledErrorAt5097960 {
		t.Fatal(err)
	}
	ipv6DisabledAt3787310 = false
	defer func() { ipv6DisabledAt3787310 = true }()
	if _, err = Jqw3ag(context.Background(), "192.0.2.1"); err != versionErrorAt5097950 {
		t.Fatal(err)
	}
	ips, err = Jqw3ag(context.Background(), "::ffff:192.0.2.1")
	if err != nil || len(ips) != 1 {
		t.Fatal(ips, err)
	}
}
func TestCustomResolverAndEmptyWrapping(t *testing.T) {
	r := new(testResolver)
	customResolverAt5097930 = r
	defer func() { customResolverAt5097930 = nil }()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	ips, err := IYMzyr8f(ctx, "a.invalid")
	if ips != nil || err != nil || r.ctx != ctx || r.call != "v4" {
		t.Fatal(ips, err, r)
	}
	_, err = Adyk1_7N("a.invalid")
	if !errors.Is(err, noIPAt5097940) || err.Error() != "couldn't find ip: a.invalid" {
		t.Fatal(err)
	}
	sentinel := errors.New("custom")
	r.err = sentinel
	r.ips = []net.IP{net.IPv4(1, 2, 3, 4)}
	ips, err = IYMzyr8f(ctx, "a.invalid")
	if len(ips) != 1 || err != sentinel {
		t.Fatal("direct callback return", ips, err)
	}
	_, err = Adyk1_7N("a.invalid")
	if err != sentinel {
		t.Fatal(err)
	}
}
