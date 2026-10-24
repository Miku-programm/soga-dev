package hysteriabrutal

import (
	"fmt"
	"github.com/apernet/quic-go/congestion"
	"net"
	"testing"
	"time"
)

type fixedClock struct{ at time.Time }

func (c *fixedClock) Now() time.Time { return c.at }
func mustPanicText(t *testing.T, want string, f func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil || fmt.Sprint(r) != want {
			t.Errorf("panic=%v want=%q", r, want)
		}
	}()
	f()
}

func TestRingGrowthWrapClearAndOriginalPanics(t *testing.T) {
	var r YXV6CyZ5Ux[kO8aHEaU9]
	for _, op := range []struct {
		name string
		f    func()
	}{{"front", func() { r.Front() }}, {"back", func() { r.Back() }}, {"pop", func() { r.PopFront() }}} {
		mustPanicText(t, "github.com/quic-go/quic-go/internal/utils/ringbuffer: "+op.name+" from an empty queue", op.f)
	}
	mustPanicText(t, "github.com/quic-go/quic-go/internal/utils/ringbuffer: offset from invalid index", func() { r.Offset(0) })
	for i := 0; i < 300; i++ {
		r.PushBack(kO8aHEaU9{joNWs9: congestion.ByteCount(i)})
	}
	for i := 0; i < 200; i++ {
		if got := r.PopFront().joNWs9; got != congestion.ByteCount(i) {
			t.Fatalf("pop %d=%d", i, got)
		}
	}
	for i := 300; i < 600; i++ {
		r.PushBack(kO8aHEaU9{joNWs9: congestion.ByteCount(i)})
	}
	for i := 200; i < 600; i++ {
		if got := r.Offset(i - 200).joNWs9; got != congestion.ByteCount(i) {
			t.Fatalf("offset %d=%d", i, got)
		}
	}
	if r.Len() != 400 || r.Front().joNWs9 != 200 || r.Back().joNWs9 != 599 {
		t.Fatal("ring order")
	}
	r.Clear()
	if !r.Empty() || r.Len() != 0 {
		t.Fatal("clear")
	}
	for _, v := range r.ckRjtgECJ {
		if v != (kO8aHEaU9{}) {
			t.Fatal("clear retains data")
		}
	}
}

func TestSparsePacketQueueCleanup(t *testing.T) {
	q := newPacketNumberIndexedQueue[wzBRgOvr8BC](2)
	v := wzBRgOvr8BC{jNhPqIq: 1250}
	if q.Emplace(-1, &v) || q.Emplace(1, nil) || !q.Emplace(10, &v) || !q.Emplace(13, &v) || q.Emplace(12, &v) {
		t.Fatal("insert policy")
	}
	if q.NumberOfPresentEntries() != 2 || q.EntrySlotsUsed() != 4 || q.GetEntry(11) != nil {
		t.Fatal("sparse slots")
	}
	called := 0
	if !q.Remove(10, func(x wzBRgOvr8BC) {
		called++
		if x.jNhPqIq != 1250 {
			t.Fatal("callback")
		}
	}) || called != 1 || q.FirstPacket() != 13 {
		t.Fatal("remove cleanup")
	}
	q.RemoveUpTo(13)
	if q.IsEmpty() {
		t.Fatal("exclusive cutoff")
	}
	q.RemoveUpTo(14)
	if !q.IsEmpty() || q.FirstPacket() != -1 || q.EntrySlotsUsed() != 0 {
		t.Fatal("final cleanup")
	}
}

func TestAckCandidateOriginalShrinkingLoop(t *testing.T) {
	b := iTmwEHfnE0Fe(10)
	for i := 0; i < 6; i++ {
		b.vFlrYPDp4.PushBack(kO8aHEaU9{joNWs9: congestion.ByteCount(i * 100)})
	}
	var a kO8aHEaU9
	if !b.bxaXSIOnnib(999, &a) || a.joNWs9 != 500 {
		t.Fatal("last point")
	}
	// 原函数每次迭代重读 Len，6 项只移除 3 项，不能改成清到只剩尾项。
	if b.vFlrYPDp4.Len() != 3 || b.vFlrYPDp4.Front().joNWs9 != 300 {
		t.Fatal("original shrinking-loop behavior changed")
	}
}

func TestWindowAging(t *testing.T) {
	f := NewWindowedFilter[JaadYOR_Vaf6, awuEkJg9p](10, Ez2Vlj[JaadYOR_Vaf6])
	f.Update(100, 0)
	f.Update(80, 3)
	f.Update(60, 9)
	if f.GetBest() != 100 || f.GetSecondBest() != 80 || f.GetThirdBest() != 60 {
		t.Fatal("quarters")
	}
	f.Update(40, 11)
	if f.GetBest() != 80 || f.GetSecondBest() != 60 || f.GetThirdBest() != 40 {
		t.Fatal("promote expired")
	}
	f.Update(20, 30)
	if f.GetBest() != 20 || f.GetThirdBest() != 20 {
		t.Fatal("all expired")
	}
}

func TestPacketLifecycleAndOriginalBoundaries(t *testing.T) {
	t.Setenv("HYSTERIA_BBR_DEBUG", "")
	clock := &fixedClock{time.Unix(1000, 0)}
	b := NewSenderFromInline(clock, 1280)
	rtt := &pacingMachineRTT{values: []time.Duration{100 * time.Millisecond}}
	b.SetRTTStatsProvider(rtt)
	if !b.InSlowStart() || b.GetCongestionWindow() != 40960 || b.irwaYXRFoja != 5120 || b.qkdjbFU0l != 25600000 || b.qAZ13bY != 2 {
		t.Fatal("constructor")
	}
	b.OnPacketSent(clock.at, 0, 1, 1280, true)
	b.OnPacketSent(clock.at.Add(time.Millisecond), 1280, 2, 1280, true)
	b.OnCongestionEventEx(2560, clock.at.Add(100*time.Millisecond), []congestion.AckedPacketInfo{{PacketNumber: 1, BytesAcked: 1280}}, []congestion.LostPacketInfo{{PacketNumber: 2, BytesLost: 1280}})
	if b.eiff4Dk6J.TotalBytesSent() != 2560 || b.eiff4Dk6J.TotalBytesAcked() != 1280 || b.eiff4Dk6J.TotalBytesLost() != 1280 || b.ehQ4ZmGJ4u != 0 || b.es2Rg_vDy != 100*time.Millisecond {
		t.Fatal("packet accounting")
	}
	if b.bandwidthEstimate() == 0 || b.PacingRate() == 0 || b.GetCongestionWindow() < 5120 {
		t.Fatal("model did not advance")
	}
	b.ndBPgZ = 5120
	b.SetMaxDatagramSize(1400)
	if b.ndBPgZ != 5120 || b.irwaYXRFoja != 5120 {
		t.Fatal("original no-rescale policy")
	}
	mustPanicText(t, "congestion BUG: decreased max datagram size from 1400 to 1280", func() { b.SetMaxDatagramSize(1280) })
	stamp := b.p4pnrwiNi
	if b.a4jZwIN(stamp.Add(10*time.Second), 200*time.Millisecond) {
		t.Fatal("expiry exact boundary")
	}
	if !b.a4jZwIN(stamp.Add(10*time.Second+time.Nanosecond), 200*time.Millisecond) || b.es2Rg_vDy != 200*time.Millisecond {
		t.Fatal("expired min RTT")
	}
	if CMryyxB(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)}) != 1280 || CMryyxB(&net.UDPAddr{IP: net.ParseIP("::1")}) != 1280 {
		t.Fatal("initial datagram size")
	}
}
