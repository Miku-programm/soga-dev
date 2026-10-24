package BOaPYrtdvWIg

import (
	"github.com/apernet/quic-go/congestion"
	"math"
	"testing"
	"time"
)

type machineRTT struct {
	congestion.RTTStatsProvider
	rtt time.Duration
}

func (r machineRTT) SmoothedRTT() time.Duration { return r.rtt }
func TestOriginalWindowMachineGoldens(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: -1}, _aNjasGyoX: 0, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("1", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 0}, _aNjasGyoX: 0, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("2", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1}, _aNjasGyoX: 0, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("3", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 123456789}, _aNjasGyoX: 0, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3fe999999999999a)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("4", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 2345678912}, _aNjasGyoX: 0, t9AYEa6hD: 1452, qB0G_BBq1s6u: math.Float64frombits(0x3fee666666666666)}
		if got := b.GetCongestionWindow(); got != 1452 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("5", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 9223372036854775807}, _aNjasGyoX: 0, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("6", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: 0, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x0)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("7", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: 0, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x7ff8000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("8", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: -1}, _aNjasGyoX: 10000000, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("9", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 0}, _aNjasGyoX: 10000000, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("10", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1}, _aNjasGyoX: 10000000, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("11", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 123456789}, _aNjasGyoX: 10000000, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3fe999999999999a)}
		if got := b.GetCongestionWindow(); got != 3086419 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("12", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 2345678912}, _aNjasGyoX: 10000000, t9AYEa6hD: 1452, qB0G_BBq1s6u: math.Float64frombits(0x3fee666666666666)}
		if got := b.GetCongestionWindow(); got != 49382713 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("13", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 9223372036854775807}, _aNjasGyoX: 10000000, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 184467440737095520 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("14", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: 10000000, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x0)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("15", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: 10000000, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x7ff8000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("16", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: -1}, _aNjasGyoX: 9223372036854775807, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("17", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 0}, _aNjasGyoX: 9223372036854775807, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("18", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1}, _aNjasGyoX: 9223372036854775807, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 18446744073 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("19", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 123456789}, _aNjasGyoX: 9223372036854775807, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3fe999999999999a)}
		if got := b.GetCongestionWindow(); got != 2846719738556200448 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("20", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 2345678912}, _aNjasGyoX: 9223372036854775807, t9AYEa6hD: 1452, qB0G_BBq1s6u: math.Float64frombits(0x3fee666666666666)}
		if got := b.GetCongestionWindow(); got != 1452 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("21", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 9223372036854775807}, _aNjasGyoX: 9223372036854775807, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("22", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: 9223372036854775807, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x0)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("23", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: 9223372036854775807, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x7ff8000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("24", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: -1}, _aNjasGyoX: -9223372036854775808, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("25", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 0}, _aNjasGyoX: -9223372036854775808, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("26", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1}, _aNjasGyoX: -9223372036854775808, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("27", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 123456789}, _aNjasGyoX: -9223372036854775808, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3fe999999999999a)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("28", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 2345678912}, _aNjasGyoX: -9223372036854775808, t9AYEa6hD: 1452, qB0G_BBq1s6u: math.Float64frombits(0x3fee666666666666)}
		if got := b.GetCongestionWindow(); got != 1452 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("29", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 9223372036854775807}, _aNjasGyoX: -9223372036854775808, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("30", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: -9223372036854775808, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x0)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("31", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: -9223372036854775808, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x7ff8000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("32", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: -1}, _aNjasGyoX: -1, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("33", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 0}, _aNjasGyoX: -1, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 10240 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("34", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1}, _aNjasGyoX: -1, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("35", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 123456789}, _aNjasGyoX: -1, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3fe999999999999a)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("36", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 2345678912}, _aNjasGyoX: -1, t9AYEa6hD: 1452, qB0G_BBq1s6u: math.Float64frombits(0x3fee666666666666)}
		if got := b.GetCongestionWindow(); got != 1452 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("37", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 9223372036854775807}, _aNjasGyoX: -1, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x3ff0000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("38", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: -1, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x0)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
	t.Run("39", func(t *testing.T) {
		b := &ON7Lun{vbyI7ILUmS1: machineRTT{rtt: 1000000000}, _aNjasGyoX: -1, t9AYEa6hD: 1280, qB0G_BBq1s6u: math.Float64frombits(0x7ff8000000000000)}
		if got := b.GetCongestionWindow(); got != 1280 {
			t.Fatalf("got %d", got)
		}
	})
}
