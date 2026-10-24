package VasCJz

import "testing"

// 固定期望值来自原 ELF 指令受限执行，脚本及证据保存在源码目录之外。
func TestConfigOriginalMachineGoldens(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		c := &R0J19auWof{ULnN_iDWLIXN: nil, Aonc19Hp33E: ""}
		got := c.GetNormalizedScMaxEachPostBytes()
		want := MUVoWl0{1000000, 1000000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("1", func(t *testing.T) {
		c := &R0J19auWof{ULnN_iDWLIXN: &MUVoWl0{13, 0}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScMaxEachPostBytes()
		want := MUVoWl0{1000000, 1000000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		c := &R0J19auWof{ULnN_iDWLIXN: &MUVoWl0{-4, -2}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScMaxEachPostBytes()
		want := MUVoWl0{-4, -2}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		c := &R0J19auWof{ULnN_iDWLIXN: &MUVoWl0{0, 1}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScMaxEachPostBytes()
		want := MUVoWl0{0, 1}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("4", func(t *testing.T) {
		c := &R0J19auWof{ULnN_iDWLIXN: &MUVoWl0{63, 65}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScMaxEachPostBytes()
		want := MUVoWl0{63, 65}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("5", func(t *testing.T) {
		c := &R0J19auWof{ULnN_iDWLIXN: &MUVoWl0{64, 1}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScMaxEachPostBytes()
		want := MUVoWl0{64, 1}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("6", func(t *testing.T) {
		c := &R0J19auWof{ULnN_iDWLIXN: &MUVoWl0{100, 99}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScMaxEachPostBytes()
		want := MUVoWl0{100, 99}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("7", func(t *testing.T) {
		c := &R0J19auWof{ULnN_iDWLIXN: &MUVoWl0{2147483647, -2147483648}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScMaxEachPostBytes()
		want := MUVoWl0{2147483647, -2147483648}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("8", func(t *testing.T) {
		c := &R0J19auWof{Iv950ouQ5AW: nil, Aonc19Hp33E: ""}
		got := c.GetNormalizedScStreamUpServerSecs()
		want := MUVoWl0{20, 80}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("9", func(t *testing.T) {
		c := &R0J19auWof{Iv950ouQ5AW: &MUVoWl0{13, 0}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScStreamUpServerSecs()
		want := MUVoWl0{20, 80}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("10", func(t *testing.T) {
		c := &R0J19auWof{Iv950ouQ5AW: &MUVoWl0{-4, -2}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScStreamUpServerSecs()
		want := MUVoWl0{-4, -2}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("11", func(t *testing.T) {
		c := &R0J19auWof{Iv950ouQ5AW: &MUVoWl0{0, 1}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScStreamUpServerSecs()
		want := MUVoWl0{0, 1}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("12", func(t *testing.T) {
		c := &R0J19auWof{Iv950ouQ5AW: &MUVoWl0{63, 65}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScStreamUpServerSecs()
		want := MUVoWl0{63, 65}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("13", func(t *testing.T) {
		c := &R0J19auWof{Iv950ouQ5AW: &MUVoWl0{64, 1}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScStreamUpServerSecs()
		want := MUVoWl0{64, 1}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("14", func(t *testing.T) {
		c := &R0J19auWof{Iv950ouQ5AW: &MUVoWl0{100, 99}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScStreamUpServerSecs()
		want := MUVoWl0{100, 99}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("15", func(t *testing.T) {
		c := &R0J19auWof{Iv950ouQ5AW: &MUVoWl0{2147483647, -2147483648}, Aonc19Hp33E: ""}
		got := c.GetNormalizedScStreamUpServerSecs()
		want := MUVoWl0{2147483647, -2147483648}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("16", func(t *testing.T) {
		c := &R0J19auWof{N2C4XMihD: nil, Aonc19Hp33E: ""}
		got := c.GetNormalizedXPaddingBytes()
		want := MUVoWl0{100, 1000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("17", func(t *testing.T) {
		c := &R0J19auWof{N2C4XMihD: &MUVoWl0{13, 0}, Aonc19Hp33E: ""}
		got := c.GetNormalizedXPaddingBytes()
		want := MUVoWl0{100, 1000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("18", func(t *testing.T) {
		c := &R0J19auWof{N2C4XMihD: &MUVoWl0{-4, -2}, Aonc19Hp33E: ""}
		got := c.GetNormalizedXPaddingBytes()
		want := MUVoWl0{-4, -2}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("19", func(t *testing.T) {
		c := &R0J19auWof{N2C4XMihD: &MUVoWl0{0, 1}, Aonc19Hp33E: ""}
		got := c.GetNormalizedXPaddingBytes()
		want := MUVoWl0{0, 1}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("20", func(t *testing.T) {
		c := &R0J19auWof{N2C4XMihD: &MUVoWl0{63, 65}, Aonc19Hp33E: ""}
		got := c.GetNormalizedXPaddingBytes()
		want := MUVoWl0{63, 65}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("21", func(t *testing.T) {
		c := &R0J19auWof{N2C4XMihD: &MUVoWl0{64, 1}, Aonc19Hp33E: ""}
		got := c.GetNormalizedXPaddingBytes()
		want := MUVoWl0{64, 1}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("22", func(t *testing.T) {
		c := &R0J19auWof{N2C4XMihD: &MUVoWl0{100, 99}, Aonc19Hp33E: ""}
		got := c.GetNormalizedXPaddingBytes()
		want := MUVoWl0{100, 99}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("23", func(t *testing.T) {
		c := &R0J19auWof{N2C4XMihD: &MUVoWl0{2147483647, -2147483648}, Aonc19Hp33E: ""}
		got := c.GetNormalizedXPaddingBytes()
		want := MUVoWl0{2147483647, -2147483648}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("24", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: nil, Aonc19Hp33E: ""}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{1000000, 1000000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("25", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: &MUVoWl0{13, 0}, Aonc19Hp33E: ""}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{1000000, 1000000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("26", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: &MUVoWl0{-4, -2}, Aonc19Hp33E: ""}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{64, 64}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("27", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: &MUVoWl0{0, 1}, Aonc19Hp33E: ""}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{64, 64}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("28", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: &MUVoWl0{63, 65}, Aonc19Hp33E: ""}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{64, 65}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("29", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: &MUVoWl0{64, 1}, Aonc19Hp33E: ""}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{64, 1}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("30", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: &MUVoWl0{100, 99}, Aonc19Hp33E: ""}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{100, 99}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("31", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: &MUVoWl0{2147483647, -2147483648}, Aonc19Hp33E: ""}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{2147483647, -2147483648}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("32", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: nil, Aonc19Hp33E: "cookie"}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{2048, 3072}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("33", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: nil, Aonc19Hp33E: "header"}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{3000, 4000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("34", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: nil, Aonc19Hp33E: "Cookie"}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{1000000, 1000000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("35", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: nil, Aonc19Hp33E: "body"}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{1000000, 1000000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
	t.Run("36", func(t *testing.T) {
		c := &R0J19auWof{JvaHmm: nil, Aonc19Hp33E: "query"}
		got := c.GetNormalizedUplinkChunkSize()
		want := MUVoWl0{1000000, 1000000}
		if got != want {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
}
