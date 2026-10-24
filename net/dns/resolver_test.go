package dns

import (
	"context"
	"errors"
	"github.com/jellydator/ttlcache/v3"
	"net"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type lookupFunc func(context.Context, string, uint16) ([]net.IP, error)

func (f lookupFunc) LookupDNSType(ctx context.Context, h string, typ uint16) ([]net.IP, error) {
	return f(ctx, h, typ)
}
func testResolver(t *testing.T, ttl time.Duration) *D23y2888a_ {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return DAlA5L(ctx, &PQL1aLrtju{C_TYygjEa0: ttl})
}
func TestRecoveredTypeLayouts(t *testing.T) {
	if unsafe.Sizeof(PQL1aLrtju{}) != 104 || unsafe.Sizeof(D23y2888a_{}) != 176 || unsafe.Sizeof(I1JijVatatst{}) != 72 {
		t.Fatal("DNS RTTI structure size mismatch")
	}
	var cache ttlcache.Cache[string, []net.IP]
	typ := reflect.TypeOf(cache)
	want := map[string]uintptr{"items": 0, "cost": 72, "metricsMu": 80, "metrics": 104, "events": 144, "stopMu": 264, "stopCh": 272, "stopped": 280, "options": 288}
	if typ.Size() != 360 {
		t.Fatalf("ttlcache selected version size %d != original 360", typ.Size())
	}
	for name, offset := range want {
		field, ok := typ.FieldByName(name)
		if !ok || field.Offset != offset {
			t.Errorf("ttlcache %s offset mismatch", name)
		}
	}
}
func TestForceResolveSelectsFirstPerTypeInRequestedOrder(t *testing.T) {
	r := testResolver(t, time.Minute)
	old4, old6 := ipv4AvailableAt50C6D60.Load(), ipv6AvailableAt50C6D64.Load()
	defer ipv4AvailableAt50C6D60.Store(old4)
	defer ipv6AvailableAt50C6D64.Store(old6)
	ipv4AvailableAt50C6D60.Store(1)
	ipv6AvailableAt50C6D64.Store(1)
	resolver := lookupFunc(func(ctx context.Context, _ string, typ uint16) ([]net.IP, error) {
		deadline, ok := ctx.Deadline()
		if !ok || time.Until(deadline) > 3*time.Second {
			t.Error("missing 3-second timeout")
		}
		if typ == 28 {
			return []net.IP{net.ParseIP("2001:db8::1"), net.ParseIP("2001:db8::2")}, nil
		}
		return []net.IP{net.ParseIP("192.0.2.1"), net.ParseIP("192.0.2.2")}, nil
	})
	got, err := r.ForceResolve(resolver, []uint16{28, 1}, "example.test")
	if err != nil || len(got) != 2 || got[0].String() != "2001:db8::1" || got[1].String() != "192.0.2.1" {
		t.Fatalf("got %v %v", got, err)
	}
	cached, ok := r.p_zkWSwY4Y.Get("example.test")
	if !ok || !reflect.DeepEqual(cached, got) {
		t.Fatal("successful result not cached")
	}
}
func TestForceResolveAvailabilityAndErrors(t *testing.T) {
	r := testResolver(t, 0)
	old4, old6 := ipv4AvailableAt50C6D60.Load(), ipv6AvailableAt50C6D64.Load()
	defer ipv4AvailableAt50C6D60.Store(old4)
	defer ipv6AvailableAt50C6D64.Store(old6)
	var mu sync.Mutex
	var calls []uint16
	resolver := lookupFunc(func(_ context.Context, _ string, typ uint16) ([]net.IP, error) {
		mu.Lock()
		calls = append(calls, typ)
		mu.Unlock()
		return nil, errors.New("lookup failed")
	})
	ipv4AvailableAt50C6D60.Store(1)
	ipv6AvailableAt50C6D64.Store(0)
	_, err := r.ForceResolve(resolver, []uint16{28, 1}, "a")
	if err == nil || !reflect.DeepEqual(calls, []uint16{1}) {
		t.Fatalf("filter %v %v", calls, err)
	}
	calls = nil
	_, _ = r.ForceResolve(resolver, []uint16{28}, "a")
	if !reflect.DeepEqual(calls, []uint16{28}) {
		t.Fatal("single explicit family must bypass availability")
	}
	calls = nil
	ipv4AvailableAt50C6D60.Store(0)
	_, _ = r.ForceResolve(resolver, []uint16{28, 1}, "a")
	if len(calls) != 2 {
		t.Fatal("empty availability filter must preserve requested list")
	}
	got, err := r.ForceResolve(resolver, nil, "a")
	if got != nil || err != nil {
		t.Fatal("zero types must return nil,nil")
	}
}
func TestForceResolvePartialAndEmptySuccess(t *testing.T) {
	r := testResolver(t, 0)
	got, err := r.ForceResolve(lookupFunc(func(context.Context, string, uint16) ([]net.IP, error) { return nil, nil }), []uint16{1}, "a")
	if err != nil || got == nil || len(got) != 0 {
		t.Fatalf("empty success %v %v", got, err)
	}
	old4, old6 := ipv4AvailableAt50C6D60.Load(), ipv6AvailableAt50C6D64.Load()
	defer ipv4AvailableAt50C6D60.Store(old4)
	defer ipv6AvailableAt50C6D64.Store(old6)
	ipv4AvailableAt50C6D60.Store(1)
	ipv6AvailableAt50C6D64.Store(1)
	got, err = r.ForceResolve(lookupFunc(func(_ context.Context, _ string, typ uint16) ([]net.IP, error) {
		if typ == 28 {
			return []net.IP{net.ParseIP("2001:db8::99")}, errors.New("discard addresses returned alongside an error")
		}
		return []net.IP{net.IPv4(192, 0, 2, 99)}, nil
	}), []uint16{28, 1}, "partial.test")
	if err != nil || len(got) != 1 || got[0].String() != "192.0.2.99" {
		t.Fatalf("partial success %v %v", got, err)
	}
}

func TestResolveLiteralCacheAndSingleflight(t *testing.T) {
	r := testResolver(t, time.Minute)
	fakeIP := net.ParseIP("192.0.2.8")
	r.p_zkWSwY4Y.AddIP("127.0.0.1", []net.IP{fakeIP})
	got, err := r.Resolve("127.0.0.1")
	if err != nil || len(got) != 1 || !got[0].Equal(fakeIP) {
		t.Fatal("cache must take precedence over literal parsing")
	}
	got, err = r.Resolve("::1")
	if err != nil || len(got) != 1 || got[0].String() != "::1" {
		t.Fatal("literal ipv6")
	}
	var calls atomic.Int32
	release := make(chan struct{})
	entered := make(chan struct{}, 20)
	r.bkaoo_vYuCup = []uint16{1}
	r.yrk3eTuNjOam = lookupFunc(func(context.Context, string, uint16) ([]net.IP, error) {
		calls.Add(1)
		entered <- struct{}{}
		<-release
		return []net.IP{fakeIP}, nil
	})
	var wg sync.WaitGroup
	for range 20 {
		wg.Go(func() {
			got, err := r.Resolve("same.test")
			if err != nil || len(got) != 1 || !got[0].Equal(fakeIP) {
				t.Error("coalesced result")
			}
		})
	}
	<-entered
	close(release)
	wg.Wait()
	if calls.Load() != 1 {
		t.Fatalf("lookup count %d", calls.Load())
	}
}
func TestParallelResolversAndPreferredFallback(t *testing.T) {
	if _, err := (GG2P4UAyQqI{}).LookupDNSType(context.Background(), "a", 1); err == nil || err.Error() != "no resolver" {
		t.Fatal("empty server error")
	}
	fast := lookupFunc(func(context.Context, string, uint16) ([]net.IP, error) { return []net.IP{net.IPv4(192, 0, 2, 1)}, nil })
	cancelled := make(chan struct{})
	slow := lookupFunc(func(ctx context.Context, _ string, _ uint16) ([]net.IP, error) {
		<-ctx.Done()
		close(cancelled)
		return nil, ctx.Err()
	})
	got, err := (GG2P4UAyQqI{slow, fast}).LookupDNSType(context.Background(), "a", 1)
	if err != nil || got[0].String() != "192.0.2.1" {
		t.Fatal("first success")
	}
	select {
	case <-cancelled:
	case <-time.After(time.Second):
		t.Fatal("loser not cancelled")
	}
	fallbackReady := make(chan struct{})
	primary := lookupFunc(func(context.Context, string, uint16) ([]net.IP, error) {
		<-fallbackReady
		return []net.IP{net.IPv4(192, 0, 2, 2)}, nil
	})
	backup := lookupFunc(func(context.Context, string, uint16) ([]net.IP, error) {
		close(fallbackReady)
		return []net.IP{net.IPv4(192, 0, 2, 3)}, nil
	})
	got, err = (&by9pe4XG{primary, backup}).LookupDNSType(context.Background(), "a", 1)
	if err != nil || got[0].String() != "192.0.2.2" {
		t.Fatal("primary must win despite faster backup")
	}
}
func TestCacheDisabledAndNoTouchOnHit(t *testing.T) {
	disabled := testResolver(t, 0).p_zkWSwY4Y
	disabled.AddIP("x", nil)
	if _, ok := disabled.Get("x"); ok {
		t.Fatal("disabled cache hit")
	}
	cache := testResolver(t, time.Second).p_zkWSwY4Y
	cache.AddIP("x", nil)
	item := cache.nSAwGB1m2.Get("x")
	expiration := item.ExpiresAt()
	for range 3 {
		if ips, ok := cache.Get("x"); !ok || ips != nil {
			t.Fatal("nil value still a hit")
		}
	}
	if !item.ExpiresAt().Equal(expiration) {
		t.Fatal("hit extended TTL")
	}
	cache.nSAwGB1m2.Set("expired", []net.IP{}, time.Nanosecond)
	time.Sleep(time.Millisecond)
	if _, ok := cache.Get("expired"); ok {
		t.Fatal("expired item returned")
	}
}
