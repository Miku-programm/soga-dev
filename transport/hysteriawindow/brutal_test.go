package hysteriawindow

import (
	"github.com/apernet/quic-go/congestion"
	"testing"
	"time"
	"unsafe"
)

func TestOriginalBrutalLayoutAndLifecycle(t *testing.T) {
	var b ON7Lun
	if unsafe.Sizeof(b) != 184 || unsafe.Offsetof(b._aNjasGyoX) != 16 || unsafe.Offsetof(b.iVGkY7) != 32 || unsafe.Offsetof(b.bzbrOl) != 40 || unsafe.Offsetof(b.qB0G_BBq1s6u) != 160 || unsafe.Offsetof(b.hxjEhjlUXU) != 168 || unsafe.Offsetof(b.mP5_CmSWJZaP) != 176 || unsafe.Sizeof(gHgCJl{}) != 24 {
		t.Fatal("RTTI")
	}
	for _, env := range []string{"", "false", "invalid", "true", "1", "TRUE"} {
		t.Setenv("HYSTERIA_BRUTAL_DEBUG", env)
		v := GPauZ708C6(1000000)
		want := env == "true" || env == "1" || env == "TRUE"
		if v.hxjEhjlUXU != want || v._aNjasGyoX != 1000000 || v.t9AYEa6hD != 1280 || v.qB0G_BBq1s6u != 1 || v.iVGkY7 == nil {
			t.Fatal("constructor", env)
		}
	}
	t.Setenv("HYSTERIA_BRUTAL_DEBUG", "false")
	v := GPauZ708C6(1000000)
	v.SetRTTStatsProvider(machineRTT{rtt: 100 * time.Millisecond})
	if v.GetCongestionWindow() != 200000 || !v.CanSend(200000) || v.CanSend(200001) {
		t.Fatal("congestion window")
	}
	now := time.Unix(100, 0)
	v.OnCongestionEventEx(0, now, make([]congestion.AckedPacketInfo, 39), make([]congestion.LostPacketInfo, 11))
	if v.qB0G_BBq1s6u != .8 || v.bzbrOl[0] != (gHgCJl{100, 39, 11}) {
		t.Fatal("event")
	}
	v.OnCongestionEventEx(0, now, make([]congestion.AckedPacketInfo, 61), nil)
	if v.qB0G_BBq1s6u != 100./111. || v.bzbrOl[0].EYG4t7Mh_5zN != 100 {
		t.Fatal("accumulation")
	}
	v.OnCongestionEventEx(0, time.Unix(110, 0), nil, nil)
	if v.qB0G_BBq1s6u != 1 || v.bzbrOl[0] != (gHgCJl{110, 0, 0}) {
		t.Fatal("expired slot")
	}
	if !v.HasPacingBudget(now) || !v.TimeUntilSend(0).IsZero() {
		t.Fatal("initial pacing")
	}
	v.OnPacketSent(now, 0, 1, 12800, true)
	if v.HasPacingBudget(now) || v.TimeUntilSend(0).Before(now.Add(time.Millisecond)) {
		t.Fatal("pacing after burst")
	}
	v.SetMaxDatagramSize(1452)
	if v.t9AYEa6hD != 1452 {
		t.Fatal("datagram")
	}
	v.OnPacketAcked(0, 0, 0, now)
	v.OnCongestionEvent(0, 0, 0)
	v.MaybeExitSlowStart()
	v.OnRetransmissionTimeout(true)
	if v.InSlowStart() || v.InRecovery() {
		t.Fatal("original constant false")
	}
}
