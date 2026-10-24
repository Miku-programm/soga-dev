package config

import (
	"math"
	"testing"
)

func TestRecoveredTimeoutCallbacks(t *testing.T) {
	dns, transport, redis := configAt509CB40, configAt509BEA0, configAt509C940
	defer func() { configAt509CB40, configAt509BEA0, configAt509C940 = dns, transport, redis }()
	for _, value := range []int{-1, 0, 1} {
		configAt509CB40.PBvoBS2NX5 = value
		configAt509BEA0.Hni9vNgOp4g = value
		configAt509BEA0.JtFIOYSoCNi = value
		configAt509C940.Eb9vzdoE = value
		init6func1()
		qZO6G3A()
		ekZtUL9pJS()
		wantDNS, wantTCP, wantUDP, wantRedis := value, value, value, value
		if value <= 0 {
			wantDNS, wantUDP, wantRedis = 10, 2, 60
		}
		if value < 0 {
			wantTCP = 120
		}
		if configAt509CB40.PBvoBS2NX5 != wantDNS || configAt509BEA0.Hni9vNgOp4g != wantTCP || configAt509BEA0.JtFIOYSoCNi != wantUDP || configAt509C940.Eb9vzdoE != wantRedis {
			t.Fatalf("timeout boundary %d was changed incorrectly", value)
		}
	}
}

func TestRecoveredPortCallbacks(t *testing.T) {
	pprof, prometheus := configAt509B3A0, configAt509B3C0
	defer func() { configAt509B3A0, configAt509B3C0 = pprof, prometheus }()
	for _, port := range []int{math.MinInt, -1, 0, 1, 65535, 65536, math.MaxInt} {
		configAt509B3A0.DXN_loYa, configAt509B3C0.DXN_loYa = port, port
		init14func1()
		init15func1()
		wantP, wantM := port, port
		if port < 1 || port > 65535 {
			wantP, wantM = 21010, 19090
		}
		if configAt509B3A0.DXN_loYa != wantP || configAt509B3C0.DXN_loYa != wantM {
			t.Fatalf("port boundary %d", port)
		}
	}
}

func TestRecoveredMiscellaneousCallback(t *testing.T) {
	old := configAt50A2900
	defer func() { configAt50A2900 = old }()
	configAt50A2900 = watWxZSBeVv{Qxaw2_S: -1, Fak_7LgVV: 0, KS63SD: 0, GozCogFPm: 0, ItKGl9keQrlY: -1, P6KMmld: 7, JQC3mlZzpS69: -1, BkelkzTWy: 129, BIdU1vYs: 0}
	tzfa9ae()
	c := configAt50A2900
	if c.Qxaw2_S != 60 || c.Fak_7LgVV != 60 || c.KS63SD != 0 || c.GozCogFPm != 2048 || c.ItKGl9keQrlY != 0 || c.P6KMmld != 7 || c.JQC3mlZzpS69 != 32 || c.BkelkzTWy != 128 || c.BIdU1vYs != 24 {
		t.Fatalf("unexpected normalized settings: %+v", c)
	}
	configAt50A2900.KS63SD = -1
	configAt50A2900.JQC3mlZzpS69 = 0
	configAt50A2900.BkelkzTWy = 0
	tzfa9ae()
	if configAt50A2900.KS63SD != 10 || configAt50A2900.JQC3mlZzpS69 != 0 || configAt50A2900.BkelkzTWy != 0 {
		t.Fatal("zero CIDR prefixes and negative random interval")
	}
	for _, value := range []float64{math.Inf(-1), -1, math.Copysign(0, -1), 0, 1, math.Inf(1), math.Float64frombits(0x7ff8000000000123)} {
		configAt50A2900.QzlEP1r3t = value
		configAt50A2900.YyLGOC = value
		configAt50A2900.KI4eIM8J = value
		configAt50A2900.DhXdMvN = value
		tzfa9ae()
		for _, got := range []float64{configAt50A2900.QzlEP1r3t, configAt50A2900.YyLGOC, configAt50A2900.KI4eIM8J, configAt50A2900.DhXdMvN} {
			if math.IsNaN(value) {
				if !math.IsNaN(got) {
					t.Fatal("NaN lost")
				}
				continue
			}
			want := value
			if value <= 0 {
				want = 0
			}
			if math.Float64bits(got) != math.Float64bits(want) {
				t.Fatalf("float %g: bits %x, want %x", value, math.Float64bits(got), math.Float64bits(want))
			}
		}
	}
}

func TestRecoveredLogDirectoryCallback(t *testing.T) {
	old := configAt509B6A0
	defer func() { configAt509B6A0 = old }()
	for _, tc := range []struct{ input, want string }{
		{"false", ""}, {"FALSE", "FALSE"}, {" false ", " false "}, {"", ""},
		{"/", ""}, {"///", ""}, {"/etc/soga///", "/etc/soga"}, {"path/child", "path/child"},
	} {
		configAt509B6A0.HSBWD7 = tc.input
		init10func1()
		if got := configAt509B6A0.HSBWD7; got != tc.want {
			t.Errorf("%q: got %q, want %q", tc.input, got, tc.want)
		}
	}
}
