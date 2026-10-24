package mieru

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/bits"
	mathrand "math/rand"
	"os"
	"sync"
)

var paddingCacheAt509C2E0 sync.Map
var paddingFloorAt50C7928 = fsAE63SXO(17) + 24

// jyDI5g @0x1c4bb40. The salt was decoded from the original callback at
// 0x1c4bbe7. This is a protocol fingerprint, not the reconstruction version.
func jyDI5g(n int, key string) int {
	if n <= 0 {
		return 0
	}
	value, ok := paddingCacheAt509C2E0.Load(key)
	if !ok {
		sum := sha256.Sum256([]byte(key + " " + "3.32.0-soga"))
		v := int(binary.BigEndian.Uint32(sum[:4]) & 0x7fffffff)
		paddingCacheAt509C2E0.Store(key, v)
		value = v
	}
	return value.(int) % n
}
func fsAE63SXO(n int) int { name, _ := os.Hostname(); return jyDI5g(n, name) } // 0x1c4bd60

// ofUaiB3MIo @0x1c4bac0. Both random draws and truncation are preserved.
func ofUaiB3MIo(n int) int {
	if n <= 0 {
		return 0
	}
	amount := float64(mathrand.Intn(n + 1))
	r := mathrand.Float64()
	return int(amount * math.Sqrt(r*r*r))
}

func a2dHepfMMQ(payload int, key string) []byte { // 0x1c4b5a0
	if jyDI5g(2, key) != 0 {
		return yR3UD3cAea(255, payload+64)
	}
	return rRHwjt(255, min(paddingFloorAt50C7928, 255))
}

func rRHwjt(maximum, minimum int) []byte { // 0x1c4b640
	extra := ofUaiB3MIo(maximum - minimum + 1)
	out := make([]byte, minimum+extra)
	_, _ = rand.Read(out)
	if minimum > 0 {
		start := 0
		if minimum < len(out) {
			start = mathrand.Intn(extra)
		}
		goAOlbfD50OV(out, start, start+minimum)
	}
	return out
}

func goAOlbfD50OV(b []byte, start, end int) { // 0x1c4b720
	for i := start; i < end; i++ {
		c := b[i]
		if c >= 32 && c <= 126 {
			continue
		}
		if c&128 != 0 && c&127 >= 32 && c&127 <= 126 {
			b[i] = c & 127
		} else {
			b[i] = byte(mathrand.Intn(95) + 32)
		}
	}
}

// yR3UD3cAea @0x1c4b800 chooses padding with a bit-density target. Repeated
// random bit positions are allowed, so the final density need not hit target.
func yR3UD3cAea(maximum, existing int) []byte {
	sample := make([]byte, existing)
	_, _ = rand.Read(sample)
	ones := 0
	for _, b := range sample {
		ones += bits.OnesCount8(b)
	}
	zeros := 8*existing - ones
	minority := min(ones, zeros)
	needed := max(int(float64(minority)/0.325)-8*existing, 0)
	ceil := (needed + 7) / 8
	size := maximum
	if maximum > ceil {
		size = ceil + ofUaiB3MIo(maximum-ceil+1)
	}
	out := make([]byte, size)
	count := max(int(float64(8*(size+existing))*0.325)-minority, 0)
	count = min(count, 8*size)
	if ones < zeros {
		for i := 0; i < count; i++ {
			pos := mathrand.Intn(8 * size)
			out[pos/8] |= 1 << uint(pos%8)
		}
	} else {
		for i := range out {
			out[i] = 255
		}
		for i := 0; i < count; i++ {
			pos := mathrand.Intn(8 * size)
			out[pos/8] &^= 1 << uint(pos%8)
		}
	}
	return out
}
