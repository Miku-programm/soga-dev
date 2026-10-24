package BaEWsPy

import (
	"testing"
	"time"
)

func TestDefaultResolverConfiguration(t *testing.T) {
	if defaultResolverAt5096320 == nil {
		t.Fatal("missing default resolver")
	}
	cfg := defaultResolverAt5096320.bebyxp
	if len(cfg.Zfr5Q8ybETV) != 20 || cfg.Zfr5Q8ybETV[0] != "tcp://8.8.8.8:53" || cfg.Zfr5Q8ybETV[19] != "https://208.67.220.220/dns-query" {
		t.Fatalf("upstream order: %#v", cfg.Zfr5Q8ybETV)
	}
	if cfg.X6T7XOmp != "ipv4_first" || cfg.C_TYygjEa0 != 30*time.Second || cfg.DMa5Vg24rVR || !cfg.ZBMqaLb {
		t.Fatalf("default resolver: %#v", cfg)
	}
	if dialerAt5029CB0.Timeout != 10*time.Second || dialerAt5029CB0.KeepAlive != 60*time.Second || dialerAt5029CB0.Control == nil {
		t.Fatal("static dialer fields")
	}
}

func TestAddressFamilyFailureThreshold(t *testing.T) {
	availabilityMutexAt50C7910.Lock()
	old4, old6 := ipv4FailuresAt50C7918, ipv6FailuresAt50C7920
	available4, available6 := ipv4AvailableAt50C6D60.Load(), ipv6AvailableAt50C6D64.Load()
	ipv4FailuresAt50C7918, ipv6FailuresAt50C7920 = 2, 2
	ipv4AvailableAt50C6D60.Store(1)
	ipv6AvailableAt50C6D64.Store(1)
	availabilityMutexAt50C7910.Unlock()
	t.Cleanup(func() {
		availabilityMutexAt50C7910.Lock()
		defer availabilityMutexAt50C7910.Unlock()
		ipv4FailuresAt50C7918, ipv6FailuresAt50C7920 = old4, old6
		ipv4AvailableAt50C6D60.Store(available4)
		ipv6AvailableAt50C6D64.Store(available6)
	})
	egKhwF(false, false)
	if ipv4FailuresAt50C7918 != 2 || ipv6FailuresAt50C7920 != 2 || ipv4AvailableAt50C6D60.Load() != 1 || ipv6AvailableAt50C6D64.Load() != 1 {
		t.Fatal("dual failure must preserve previous state")
	}
	egKhwF(false, true)
	if ipv4FailuresAt50C7918 != 3 || ipv6FailuresAt50C7920 != 0 || ipv4AvailableAt50C6D60.Load() != 0 || ipv6AvailableAt50C6D64.Load() != 1 {
		t.Fatal("threshold")
	}
	egKhwF(false, true)
	if ipv4FailuresAt50C7918 != 3 {
		t.Fatal("count not capped")
	}
	egKhwF(true, true)
	if ipv4AvailableAt50C6D60.Load() != 1 || ipv4FailuresAt50C7918 != 0 {
		t.Fatal("successful family not reset")
	}
	for i := 0; i < 2; i++ {
		egKhwF(true, false)
		if ipv6AvailableAt50C6D64.Load() != 1 {
			t.Fatal("disabled before third failure")
		}
	}
	egKhwF(true, false)
	if ipv6AvailableAt50C6D64.Load() != 0 {
		t.Fatal("third failure did not disable")
	}
}

func TestRouteProbeUsesDialAndClose(t *testing.T) {
	if qYgUF5Z("invalid-network", "127.0.0.1:9") {
		t.Fatal("invalid network succeeded")
	}
	if !qYgUF5Z("udp4", "127.0.0.1:9") {
		t.Fatal("loopback UDP route should succeed without remote service")
	}
}
