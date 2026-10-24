package config

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
	"time"
)

func TestOriginalRandomIntervalBounds(t *testing.T) {
	old := configAt50A2900.KS63SD
	t.Cleanup(func() { configAt50A2900.KS63SD = old })
	configAt50A2900.KS63SD = 0
	if got := EVzHyyNDe(-time.Second); got != -time.Second {
		t.Fatal(got)
	}
	configAt50A2900.KS63SD = 10
	var source bytes.Buffer
	for _, value := range []uint64{0, 19_999_999_999, 0, 14_999_999_999} {
		binary.Write(&source, binary.BigEndian, value)
	}
	oldReader := rand.Reader
	rand.Reader = &source
	t.Cleanup(func() { rand.Reader = oldReader })
	for _, tc := range []struct{ in, want time.Duration }{{60 * time.Second, 50 * time.Second}, {60 * time.Second, 70*time.Second - 1}, {5 * time.Second, 0}, {5 * time.Second, 15*time.Second - 1}} {
		if got := EVzHyyNDe(tc.in); got != tc.want {
			t.Fatalf("%v: %v != %v", tc.in, got, tc.want)
		}
	}
	configAt50A2900.KS63SD = -1
	defer func() {
		if recover() == nil {
			t.Fatal("invalid interval must preserve panic")
		}
	}()
	EVzHyyNDe(time.Second)
}
