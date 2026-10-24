package syncpool

import (
	"bytes"
	"sync"
)

// BytesBufferPoolAt5044220 exposes the original shared pool under its VA.
// Its static New funcval is init.func1 at 0x1943920.
var BytesBufferPoolAt5044220 = sync.Pool{New: func() any { return new(bytes.Buffer) }}

// LjZLCz9: 0x19437a0. Reset retains the buffer's backing allocation.
func LjZLCz9(b *bytes.Buffer) { b.Reset(); BytesBufferPoolAt5044220.Put(b) }
