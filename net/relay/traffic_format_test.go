package relay

import "testing"

func TestOriginalTrafficUnits(t *testing.T) {
	for _, tt := range []struct {
		n uint64
		s string
	}{{0, "0B"}, {1023, "1023B"}, {1024, "1.00KB"}, {(1 << 20) - 1, "1024.00KB"}, {1 << 20, "1.00MB"}, {1 << 30, "1.00GB"}, {1 << 40, "1.00TB"}, {1 << 63, "8388608.00TB"}, {^uint64(0), "16777216.00TB"}} {
		if s := LTaVAEf(tt.n); s != tt.s {
			t.Errorf("%d: %q != %q", tt.n, s, tt.s)
		}
	}
}
