package ad5KEbw_EzWk

import (
	"encoding/hex"
	"fmt"
	"testing"
	"unsafe"
)

// Golden values are from bounded execution of original 0x1c844e0..0x1c84840,
// never from a second Go implementation of the recurrence.
func TestNativeXorshiftOriginalTraces(t *testing.T) {
	for _, row := range originalXorshiftTraces {
		t.Run(fmt.Sprintf("%s/%d/%d/%d", row.Function, row.Length, row.Capacity, row.Extra), func(t *testing.T) {
			var data [32]byte
			for i := range data {
				data[i] = byte(i + 1)
			}
			r := HelGdo2ETK{0x1122334455667788, 0x8877665544332211}
			switch row.Function {
			case "next":
				if got := r.Next(); got != row.Result {
					t.Fatalf("next=%x want%x", got, row.Result)
				}
			case "seed":
				r.InitFromBytes(data[:row.Length:row.Capacity])
			case "length":
				r.InitFromBytesLen(data[:row.Length:row.Capacity], row.Extra)
			default:
				t.Fatal(row.Function)
			}
			if [2]uint64{r.mFYAwci, r.bWvBk2x4oPae} != row.State {
				t.Fatalf("state=%x,%x want%x", r.mFYAwci, r.bWvBk2x4oPae, row.State)
			}
			if hex.EncodeToString(data[:]) != row.InputHex {
				t.Fatalf("input mutation=%x want%s", data, row.InputHex)
			}
		})
	}
}
func TestNativeSSRInfoLayout(t *testing.T) {
	var s L7E7H2aW8KVs
	if unsafe.Sizeof(s) != 224 || unsafe.Offsetof(s.NxYA16xMfF) != 96 || unsafe.Offsetof(s.N3pMXuwX85c) != 208 || unsafe.Offsetof(s.RLXBf_luGj) != 216 {
		t.Fatal("native server info RTTI mismatch")
	}
	if unsafe.Sizeof(HelGdo2ETK{}) != 16 {
		t.Fatal("native xorshift RTTI mismatch")
	}
}
