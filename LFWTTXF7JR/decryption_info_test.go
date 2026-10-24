package LFWTTXF7JR

import (
	"encoding/base64"
	"strings"
	"testing"
	"unsafe"
)

func TestOriginalDecryptionInfo(t *testing.T) {
	key := base64.RawURLEncoding.EncodeToString(make([]byte, 32))
	for i, mode := range []string{"native", "xorpub", "random"} {
		raw := "mlkem768x25519plus." + mode + ".600-900s.10-20.5-10." + key
		got, ok := (&VkkFLH{DlDK5Y9AS: raw}).GetDecryptionInfo()
		if !ok || got.KAeFR260tbP != uint32(i) || got.E3a03_Hf != 600 || got.R3tya30 != 900 || got.LRGrsuCn != "10-20.5-10" || got.I_Wzoo_nf != key {
			t.Fatalf("%s: %#v %v", mode, got, ok)
		}
	}
	for _, raw := range []string{"", "none", "wrong.native.0." + key, "mlkem768x25519plus.other.0." + key, "mlkem768x25519plus.native.-1." + key, "mlkem768x25519plus.native.1sS." + key, "mlkem768x25519plus.native.0." + strings.Repeat("A", 20)} {
		if _, ok := (&VkkFLH{DlDK5Y9AS: raw}).GetDecryptionInfo(); ok {
			t.Fatalf("accepted %q", raw)
		}
	}
	// 原代码虽忽略DecodeString错误，但这个非法尾部仅解出30字节，仍拒绝。
	raw := "mlkem768x25519plus.native.0." + key + "!"
	got, ok := (&VkkFLH{DlDK5Y9AS: raw}).GetDecryptionInfo()
	if ok || got != nil {
		t.Fatalf("partial decode %#v %v", got, ok)
	}
	var x Bqm34MJ
	if unsafe.Sizeof(x) != 56 || unsafe.Offsetof(x.KAeFR260tbP) != 16 || unsafe.Offsetof(x.E3a03_Hf) != 24 || unsafe.Offsetof(x.R3tya30) != 32 || unsafe.Offsetof(x.LRGrsuCn) != 40 {
		t.Fatal("RTTI layout")
	}
}
func TestOriginalDecryptionInfoShortOnlyPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("original out-of-range slice must panic")
		}
	}()
	(&VkkFLH{DlDK5Y9AS: "mlkem768x25519plus.native.0.short"}).GetDecryptionInfo()
}
