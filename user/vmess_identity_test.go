package user

import (
	"bytes"
	"encoding/hex"
	"testing"
	"unsafe"
)

func TestVMessIdentityAndCacheAliasing(t *testing.T) {
	if unsafe.Sizeof(DALfaCvQ{}) != 33 {
		t.Fatal("identity layout")
	}
	id, err := w6MaVR("01234567-89ab-cdef-0123-456789abcdef")
	if err != nil {
		t.Fatal(err)
	}
	first := id.GetCmdKey()
	second := id.GetCmdKey()
	if hex.EncodeToString(first) != "a52ea939db2210f0e37ad1bea4b0d3a0" || !bytes.Equal(first, second) {
		t.Fatalf("command key: %x", first)
	}
	first[0] ^= 0xff
	if bytes.Equal(first, id.GetCmdKey()) {
		t.Fatal("first digest must be independent of cache")
	}
	second[0] ^= 0xff
	if !bytes.Equal(second, id.GetCmdKey()) {
		t.Fatal("later digest must expose cache")
	}
	if bad, err := w6MaVR("invalid"); bad != nil || err == nil {
		t.Fatal("invalid UUID accepted")
	}
}

func TestVMessAlterIDIndependentVectors(t *testing.T) {
	id, _ := w6MaVR("01234567-89ab-cdef-0123-456789abcdef")
	previous := id.Yfd33z6YYdJa[:]
	for _, want := range []string{"4b04e28f176fccd9ef7c5747035f2563", "22d00f50fb077ccdad124a11434ec24a", "60ea7c88bef8fbe663a0667d3557c061"} {
		next := make([]byte, 16)
		xshVaRu(previous, next)
		if hex.EncodeToString(next) != want {
			t.Fatalf("alter ID: %x want %s", next, want)
		}
		previous = next
	}
	// A too-small backing array is allocated internally by Sum, but the return
	// slice is discarded by the original helper; the caller's slice stays intact.
	short := []byte{7, 8}
	xshVaRu(previous, short)
	if !bytes.Equal(short, []byte{7, 8}) {
		t.Fatal("unexpected destination replacement")
	}
}
