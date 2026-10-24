package JafvCs5

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestOriginalTruncatedHMACSHA1(t *testing.T) {
	got := L01ck7znS(bytes.Repeat([]byte{0x0b}, 20), []byte("Hi "), []byte("There"))
	if hex.EncodeToString(got) != "b617318655057264e28b" || len(got) != 10 || cap(got) < 20 {
		t.Fatal(hex.EncodeToString(got), len(got), cap(got))
	}
}
