package syncpool

import (
	"errors"
	"math/bits"
	"sync"
)

// L4upasOm preserves the 24-byte RTTI layout and 17 power-of-two classes.
type L4upasOm struct{ sOcwxhPNs []sync.Pool }

// SharedPoolAt5096978 exposes the original package global under an analyst name.
var SharedPoolAt5096978 = newSharedPool()

// init: 0x1941820; New closure: 0x19438c0.
func newSharedPool() *L4upasOm {
	p := &L4upasOm{sOcwxhPNs: make([]sync.Pool, 17)}
	for i := range p.sOcwxhPNs {
		n := 1 << i
		p.sOcwxhPNs[i].New = func() any { return make([]byte, n) }
	}
	return p
}

// Get: 0x1941960.
func (p *L4upasOm) Get(n int) []byte {
	if n < 0 {
		panic("alloc.Get: len out of range")
	}
	if n == 0 {
		return nil
	}
	if n > 65536 {
		return make([]byte, n)
	}
	return p.sOcwxhPNs[bits.Len(uint(n-1))].Get().([]byte)[:n]
}

// Put: 0x1941b20. The caller's length is retained; capacity selects the class.
func (p *L4upasOm) Put(b []byte) error {
	n := cap(b)
	if n == 0 || n > 65536 {
		return nil
	}
	if n&(n-1) != 0 {
		return errors.New("allocator Put() incorrect buffer size")
	}
	p.sOcwxhPNs[bits.Len(uint(n))-1].Put(b)
	return nil
}
