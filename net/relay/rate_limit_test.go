package relay

import (
	"math"
	"testing"
	"time"
	"unsafe"
)

func TestOriginalLimiterBoundaries(t *testing.T) {
	var l L337vk2G
	if unsafe.Sizeof(l) != 16 || unsafe.Offsetof(l.hgkiIUSUPlQ) != 8 {
		t.Fatal("layout")
	}
	if !l.AllowAndAdd(math.MaxInt) || l.GetSpeedLimit() != 0 {
		t.Fatal("unlimited zero")
	}
	l.UpdateLimit(1 << 20)
	if l.GetSpeedLimit() != 8 || l.liai0O8.Burst() != 1<<20 {
		t.Fatal("units")
	}
	ptr := l.liai0O8
	l.UpdateLimit(math.MaxInt64)
	if ptr != l.liai0O8 || l.hgkiIUSUPlQ != math.MaxInt32 || l.liai0O8.Burst() != math.MaxInt32 {
		t.Fatal("clamp and reuse")
	}
	l.UpdateLimit(-1)
	if l.liai0O8 != nil || l.hgkiIUSUPlQ != 0 {
		t.Fatal("disable")
	}
	l.UpdateLimit(10)
	if !l.AllowAndAdd(10) || l.AllowAndAdd(11) {
		t.Fatal("burst")
	}
	// Oversize Add logs and returns; it must not reserve or sleep for this payload.
	before := time.Now()
	l.Add(11)
	if time.Since(before) > time.Second {
		t.Fatal("unexpected oversize delay")
	}
	l.UpdateLimit(0)
	l.Add(1)
}

func TestOriginalLimiterWaitsForReservation(t *testing.T) {
	var l L337vk2G
	l.UpdateLimit(100)
	if !l.AllowAndAdd(100) {
		t.Fatal("initial burst")
	}
	before := time.Now()
	l.Add(10)
	if time.Since(before) < 70*time.Millisecond {
		t.Fatal("reservation did not wait")
	}
}
