package limit

// Mblp8L: original RTTI 0x36b6800, 48 bytes. The last field is retained even
// though the recovered AddTraffic/Sum methods do not access it.
type Mblp8L struct {
	qzaQaE4    int64
	d5H9TK     []int64
	mgvGqsp    int
	cHQ7V3cDa0 int
}

// AddTraffic: 0x183d2c0. A zero sum plus zero sample does not advance the ring.
// Bounds and zero-length panics are left to Go, matching original operations.
func (w *Mblp8L) AddTraffic(value int64) {
	if w.qzaQaE4 != 0 || value != 0 {
		w.qzaQaE4 -= w.d5H9TK[w.mgvGqsp]
		w.d5H9TK[w.mgvGqsp] = value
		w.qzaQaE4 += value
		w.mgvGqsp = (w.mgvGqsp + 1) % len(w.d5H9TK)
	}
}

// Sum: 0x183d340.
func (w *Mblp8L) Sum() int64 { return w.qzaQaE4 }
