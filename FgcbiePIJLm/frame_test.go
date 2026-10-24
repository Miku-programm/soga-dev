package FgcbiePIJLm

import (
	"bytes"
	"errors"
	"testing"
)

func TestOriginalFrameWriterByteVectors(t *testing.T) {
	for _, tc := range []struct {
		h    Wmyq5K0sQ0w
		want []byte
	}{
		{Wmyq5K0sQ0w{HcIhwPLvD: true, IsHe8C: 2, L4ixlgT8VpTo: 125}, []byte{0x82, 125}},
		{Wmyq5K0sQ0w{HcIhwPLvD: true, IsHe8C: 2, L4ixlgT8VpTo: 126}, []byte{0x82, 126, 0, 126}},
		{Wmyq5K0sQ0w{HcIhwPLvD: true, IsHe8C: 2, X4IS_ER04TVv: true, BiCmXtyECQl: [4]byte{1, 2, 3, 4}, L4ixlgT8VpTo: 65536}, []byte{0x82, 255, 0, 0, 0, 0, 0, 1, 0, 0, 1, 2, 3, 4}},
	} {
		var b bytes.Buffer
		if e := KzLwHzhWB90h(&b, tc.h); e != nil || !bytes.Equal(b.Bytes(), tc.want) {
			t.Fatal(b.Bytes(), tc.want, e)
		}
	}
	p := []byte("offsets cross chunk boundaries")
	whole := append([]byte(nil), p...)
	mask := [4]byte{1, 2, 3, 4}
	W2gws0VQo(whole, mask, 0)
	W2gws0VQo(p[:7], mask, 0)
	W2gws0VQo(p[7:], mask, 7)
	if !bytes.Equal(p, whole) {
		t.Fatal(p, whole)
	}
}
func TestOriginalProtocolErrors(t *testing.T) {
	if e := IYI3RQDfj4(Wmyq5K0sQ0w{IsHe8C: 3}, 0); !errors.Is(e, ErrProtocolOpCodeReserved) {
		t.Fatal(e)
	}
	if e := IYI3RQDfj4(Wmyq5K0sQ0w{IsHe8C: 9, L4ixlgT8VpTo: 126}, 0); !errors.Is(e, ErrProtocolControlPayloadOverflow) {
		t.Fatal(e)
	}
	if e := QXKFKa(1005, ""); !errors.Is(e, ErrProtocolStatusCodeApplicationLevel) {
		t.Fatal(e)
	}
	if e := QXKFKa(1000, string([]byte{255})); !errors.Is(e, ErrProtocolInvalidUTF8) {
		t.Fatal(e)
	}
}
