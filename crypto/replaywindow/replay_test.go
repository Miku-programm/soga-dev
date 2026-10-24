package replaywindow

import (
	"crypto/sha256"
	"encoding/binary"
	"hash/maphash"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestReplayLayoutAndHash(t *testing.T) {
	var c AiTSgb7OJeoK
	if unsafe.Sizeof(c) != 80 || unsafe.Offsetof(c.d6aQwf8rDu) != 32 || unsafe.Offsetof(c.oKYp41ywPPbH) != 56 {
		t.Fatal("RTTI")
	}
	for _, p := range [][]byte{nil, []byte("native salt"), {0, 255, 123}} {
		a, b := g1TDyioZI(p)
		if a != maphash.Bytes(seedAt50C79F8, p) || b != maphash.Bytes(seedAt50C7A00, p) {
			t.Fatal(a, b)
		}
	}
}
func TestReplayStrictRotationAndAtomicDuplicate(t *testing.T) {
	c := EJaHm2W(2, 4, 1e-20)
	data := func(n uint64) []byte {
		p := make([]byte, 8)
		binary.BigEndian.PutUint64(p, n)
		h := sha256.Sum256(p)
		return h[:]
	}
	for i := uint64(1); i <= 3; i++ {
		if !c.TestAndAdd(data(i)) {
			t.Fatal(i)
		}
	}
	if c.ii_apO != 0 || c.uSo7A4DU != 3 {
		t.Fatal("threshold is strictly greater", c.ii_apO, c.uSo7A4DU)
	}
	if !c.TestAndAdd(data(4)) || c.ii_apO != 1 || c.uSo7A4DU != 1 {
		t.Fatal("rotation")
	}
	if c.TestAndAdd(data(1)) {
		t.Fatal("old layer lost early")
	}
	c.TestAndAdd(data(5))
	c.TestAndAdd(data(6))
	c.TestAndAdd(data(7))
	if c.ii_apO != 0 || !c.TestAndAdd(data(1)) {
		t.Fatal("cleared layer retains old entry")
	}
	c = EJaHm2W(2, 100, 1e-12)
	var accepted atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if c.TestAndAdd([]byte("same")) {
				accepted.Add(1)
			}
		}()
	}
	wg.Wait()
	if accepted.Load() != 1 {
		t.Fatal(accepted.Load())
	}
	if !(*AiTSgb7OJeoK)(nil).TestAndAdd(nil) {
		t.Fatal("nil cache")
	}
}
