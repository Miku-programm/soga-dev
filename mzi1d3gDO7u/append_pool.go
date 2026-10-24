package mzi1d3gDO7u

import (
	"recovered.local/soga/Hx04IbKk"
	"sync"
)

// AppendPoolAt5044260 is the original shared sync.Pool. The exported name
// identifies its VA because the original Get operation was inlined at callers.
// New: 0x1943960 (init.func2), a non-nil empty named slice.
var AppendPoolAt5044260 = sync.Pool{New: func() any {
	b := Hx04IbKk.BaDekw7qe{}
	return &b
}}

// Wr41CoUmtL0: 0x1943800.
func Wr41CoUmtL0(b *Hx04IbKk.BaDekw7qe) {
	b.Reset()
	AppendPoolAt5044260.Put(b)
}
