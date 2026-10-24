package xx5fRyrFw

import (
	"reflect"
	"testing"
)

func TestOriginalMergeAliasingAndDuplicateBehavior(t *testing.T) {
	right := VHWq49iRat{"a", "a", ""}
	out := VHWq49iRat{}.Merge(right)
	if &out[0] != &right[0] || !reflect.DeepEqual(out, right) {
		t.Fatal(out)
	}
	left := make(VHWq49iRat, 2, 8)
	left[0], left[1] = "a", "a"
	out = left.Merge(VHWq49iRat{"b", "b", "", "A"})
	if !reflect.DeepEqual(out, VHWq49iRat{"a", "a", "b", "", "A"}) || &out[0] != &left[0] {
		t.Fatal(out)
	}
	if out := VHWq49iRat(nil).Merge(nil); out != nil {
		t.Fatal(out)
	}
}
