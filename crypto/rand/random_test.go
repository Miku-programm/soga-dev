package rand

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func withRandomWords(t *testing.T, words ...uint64) {
	t.Helper()
	var b bytes.Buffer
	for _, v := range words {
		binary.Write(&b, binary.BigEndian, v)
	}
	old := rand.Reader
	rand.Reader = &b
	t.Cleanup(func() { rand.Reader = old })
}
func TestRandomBigEndianAndRejection(t *testing.T) {
	withRandomWords(t, 0xfedcba9876543210, 0x7fffffff00000000, 0x0000000700000000, 0x7fffffffffffffff, 17)
	if got := Ctna67jV(); got != 0x7edcba9876543210 {
		t.Fatalf("got %x", got)
	}
	if got := NMT2SOqf(3); got != 1 {
		t.Fatalf("Int31n rejection: %d", got)
	}
	if got := TfHpoOcv(3); got != 2 {
		t.Fatalf("Int63n rejection: %d", got)
	}
}
func TestRandomIntnBoundary(t *testing.T) {
	withRandomWords(t, 0x0000000600000000, 0x0000000000000007)
	if Co6D_zfls(8) != 6 {
		t.Fatal("small range must use high31bits")
	}
	if Co6D_zfls(1<<32) != 7 {
		t.Fatal("large range must use low bits of63bit word")
	}
}
func TestRandomInvalidArguments(t *testing.T) {
	tests := []struct {
		f   func()
		msg string
	}{
		{func() { Co6D_zfls(0) }, "invalid argument to Intn"},
		{func() { NMT2SOqf(-1) }, "invalid argument to Int31n"},
		{func() { TfHpoOcv(0) }, "invalid argument to Int63n"},
	}
	for _, tc := range tests {
		func() {
			defer func() {
				if got := recover(); got != tc.msg {
					t.Errorf("panic=%v want %q", got, tc.msg)
				}
			}()
			tc.f()
		}()
	}
}
