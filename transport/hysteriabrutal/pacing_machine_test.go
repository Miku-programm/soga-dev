package hysteriabrutal

import (
	"github.com/apernet/quic-go/congestion"
	"math"
	"testing"
	"time"
)

type pacingMachineRTT struct {
	congestion.RTTStatsProvider
	values []time.Duration
	calls  int
}

func (r *pacingMachineRTT) MinRTT() time.Duration {
	i := r.calls
	r.calls++
	if i >= len(r.values) {
		i = len(r.values) - 1
	}
	return r.values[i]
}
func TestOriginalPacingRateMachineGoldens(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != -1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("1", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != true || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 2 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("2", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("3", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != -9223372036854775808 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("4", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 150000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("5", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("6", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("7", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("8", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 100000 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("9", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("10", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("11", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("12", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("13", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != -1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("14", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("15", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("16", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("17", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("18", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 40960000000000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("19", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("20", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 3276800 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 2 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("21", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 99999 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("22", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != -1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("23", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("24", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("25", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("26", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 13835058055282163712 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("27", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != -1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("28", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("29", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 327680000000000 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != false || r.calls != 2 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("30", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != -1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("31", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("32", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("33", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 100000 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("34", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 13835058055282163712 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("35", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 100000 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("36", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("37", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != -1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("38", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 100000 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("39", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != -1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("40", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("41", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("42", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 40960000000000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("43", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("44", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 6917529027641081856 || b.wnghShWZg != 1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("45", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 2 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("46", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("47", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("48", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("49", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("50", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 100000 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("51", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 6917529027641081856 || b.wnghShWZg != -1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("52", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("53", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 2 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("54", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("55", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 40960000000000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("56", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 327680000000000 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 2 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("57", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("58", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != -1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("59", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 750000 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("60", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("61", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("62", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("63", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("64", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 750000 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("65", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("66", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("67", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("68", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != -1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("69", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("70", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 750000 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("71", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("72", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("73", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("74", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("75", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 2 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("76", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("77", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("78", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != -1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("79", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("80", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 40960000000000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("81", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("82", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 2 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("83", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("84", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x7ff8000000000000), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("85", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != -2 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("86", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("87", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("88", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 150000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("89", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073708551616 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("90", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 100000 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("91", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 13835058055282163712 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("92", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("93", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("94", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("95", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 0, dv_p1GSp: true, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("96", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != -1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("97", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 750000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("98", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 2885000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("99", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("100", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: -1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != -1 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("101", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: true, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("102", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0x4007147ae147ae14), rbHB4J1ksIV: 1, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("103", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 9223372036854775808}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != -1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("104", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 100000, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 100000 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("105", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x3ff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 255}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 1000000 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("106", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 0}}}, rDtgaAw: math.Float64frombits(0x3fe8000000000000), rbHB4J1ksIV: 1, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 9223372036854775807, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 1 || b.wnghShWZg != 9223372036854775807 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("107", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x0), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("108", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{-1, -1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 18446744073709551615}}}, rDtgaAw: math.Float64frombits(0xbff0000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: true, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 0}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(50000) }()
		if panicked != false || b.rbHB4J1ksIV != 9223372036854775808 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("109", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{1, 1}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x7ff0000000000000), rbHB4J1ksIV: 18446744073709551615, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 1}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(-1) }()
		if panicked != false || b.rbHB4J1ksIV != 18446744073709551615 || b.wnghShWZg != 0 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("110", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 200000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x4000000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 1638400 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 2 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("111", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000, 0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x4000000000000000), rbHB4J1ksIV: 0, dv_p1GSp: false, yL5qU8a: false, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != true || b.rbHB4J1ksIV != 0 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 2 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("112", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x4000000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != false || b.rbHB4J1ksIV != 2000000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("113", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x4000000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 40960, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 2000000 || b.wnghShWZg != 0 || b.yL5qU8a != false || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("114", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{0}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x4000000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: true, wnghShWZg: 1, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(0) }()
		if panicked != true || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 1 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
	t.Run("115", func(t *testing.T) {
		r := &pacingMachineRTT{values: []time.Duration{100000000}}
		b := &tGYu23kout{jzNuIRaj5L4f: r, clm34LfHOfH: &Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{fPLAVE6X: []f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{{nvK8flg3g: 1000000}}}, rDtgaAw: math.Float64frombits(0x4000000000000000), rbHB4J1ksIV: 100000000, dv_p1GSp: false, yL5qU8a: true, hcFC7avZ: false, wnghShWZg: 0, kPUkZLy: 40960, zaXoVmSmYwgs: 5120, gX1RE3uW: 2}
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); b.n0_hVN2CE0Lj(1) }()
		if panicked != false || b.rbHB4J1ksIV != 100000000 || b.wnghShWZg != 1 || b.yL5qU8a != true || r.calls != 0 {
			t.Fatalf("panic=%v rate=%d loss=%d over=%v calls=%d", panicked, b.rbHB4J1ksIV, b.wnghShWZg, b.yL5qU8a, r.calls)
		}
	})
}
