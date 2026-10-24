package mieru

import (
	"sync"
	"time"
)

// fA9rSx, original 64-byte RTTI 0x36df338: two generations of 64-bit
// fingerprints. The first sixteen input bytes alone determine identity.
type fA9rSx struct {
	oCqr1HwQivr sync.Mutex
	zKMbevD     int
	g0gLsjzW    time.Duration
	ljZIb5e9OR  time.Time
	cgoZbeQ     map[uint64]struct{}
	vksDKih     map[uint64]struct{}
}

var replayOnceAt50C7C30 sync.Once
var replayAt5096328 *fA9rSx

func yCMyVg() *fA9rSx { //0x1c4bfa0 and 0x1c5d340
	replayOnceAt50C7C30.Do(func() { replayAt5096328 = b3L9YozU4h(0x400000, 360*time.Second) })
	return replayAt5096328
}
func b3L9YozU4h(size int, duration time.Duration) *fA9rSx { //0x1c4c000
	return &fA9rSx{zKMbevD: size, g0gLsjzW: duration, ljZIb5e9OR: time.Now().Add(duration), cgoZbeQ: make(map[uint64]struct{}), vksDKih: make(map[uint64]struct{})}
}
func ijqr7F(b []byte) uint64 { //0x1c4c440, FNV-1a
	h := uint64(0xcbf29ce484222325)
	for _, v := range b {
		h ^= uint64(v)
		h *= 0x100000001b3
	}
	return h
}
func (r *fA9rSx) n7hmYhVTVA(b []byte) bool { //0x1c4c100
	if len(b) < 16 {
		return false
	}
	h := ijqr7F(b[:16])
	r.oCqr1HwQivr.Lock()
	defer r.oCqr1HwQivr.Unlock()
	if time.Since(r.ljZIb5e9OR) > r.g0gLsjzW {
		r.cgoZbeQ = make(map[uint64]struct{})
		r.vksDKih = make(map[uint64]struct{})
		r.ljZIb5e9OR = time.Now().Add(r.g0gLsjzW)
	}
	if len(r.cgoZbeQ) >= r.zKMbevD || time.Now().After(r.ljZIb5e9OR) {
		r.vksDKih = r.cgoZbeQ
		r.cgoZbeQ = make(map[uint64]struct{})
		r.ljZIb5e9OR = time.Now().Add(r.g0gLsjzW)
	}
	if _, ok := r.cgoZbeQ[h]; ok {
		return true
	}
	if _, ok := r.vksDKih[h]; ok {
		return true
	}
	r.cgoZbeQ[h] = struct{}{}
	return false
}
