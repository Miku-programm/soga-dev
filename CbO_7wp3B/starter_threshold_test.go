package CbO_7wp3B

import (
	"crypto/rand"
	"testing"
)

type thresholdEntropy struct {
	pattern [8]byte
	read    int
}

func (r *thresholdEntropy) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = r.pattern[r.read%8]
		r.read++
	}
	return len(p), nil
}

// 预期值来自受限原机码执行，且验证不可删去的每字节八字节熵读取。
func TestStarterThresholdOriginalMachineFixtures(t *testing.T) {
	old := rand.Reader
	t.Cleanup(func() { rand.Reader = old })
	var input [32]byte
	for i := range input {
		input[i] = byte(i)
	}
	for _, pattern := range [][8]byte{{}, {255, 255, 255, 255, 255, 255, 255, 255}, {0, 1, 2, 3, 4, 5, 6, 7}} {
		r := &thresholdEntropy{pattern: pattern}
		rand.Reader = r
		if got := h68JK8g(input[:]); got != 3 {
			t.Fatalf("result %d", got)
		}
		if r.read != 256 {
			t.Fatalf("entropy bytes = %d, want original 256", r.read)
		}
	}
	r := &thresholdEntropy{}
	rand.Reader = r
	if got := h68JK8g(nil); got != 3 || r.read != 0 {
		t.Fatalf("empty result=%d entropy=%d", got, r.read)
	}
}
