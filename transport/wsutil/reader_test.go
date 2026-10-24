package wsutil

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	ws "recovered.local/soga/transport/websocket"
	"testing"
)

func wsWire(op byte, fin bool, p []byte) []byte {
	h := []byte{op, 128}
	if fin {
		h[0] |= 128
	}
	switch {
	case len(p) < 126:
		h[1] |= byte(len(p))
	case len(p) <= 65535:
		h[1] |= 126
		h = binary.BigEndian.AppendUint16(h, uint16(len(p)))
	default:
		h[1] |= 127
		h = binary.BigEndian.AppendUint64(h, uint64(len(p)))
	}
	mask := [4]byte{1, 5, 9, 13}
	h = append(h, mask[:]...)
	for i, b := range p {
		h = append(h, b^mask[i%4])
	}
	return h
}
func TestOriginalReaderFragmentationAndIntermediate(t *testing.T) {
	wire := wsWire(2, false, []byte("first"))
	wire = append(wire, wsWire(9, true, []byte("ping"))...)
	wire = append(wire, wsWire(0, true, []byte("last"))...)
	var ping string
	r := &HLKO3Cdqw{GX04GFKh1Gz: bytes.NewReader(wire), WXUh5psqY7: ws.StateServerSide, EbVEHBvY4lh: func(h ws.Wmyq5K0sQ0w, src io.Reader) error {
		b, e := io.ReadAll(src)
		ping = string(b)
		if h.IsHe8C != 9 {
			t.Fatal(h)
		}
		return e
	}}
	if _, e := r.Read(make([]byte, 1)); e != ErrNoFrameAdvance {
		t.Fatal(e)
	}
	h, e := r.NextFrame()
	if e != nil || h.IsHe8C != 2 || h.HcIhwPLvD {
		t.Fatal(h, e)
	}
	b, e := io.ReadAll(r)
	if e != nil || string(b) != "firstlast" || ping != "ping" {
		t.Fatal(string(b), ping, e)
	}
	if _, e = r.NextFrame(); e != io.EOF {
		t.Fatal(e)
	}
}
func TestOriginalReaderWireLengthsMaskAndLimits(t *testing.T) {
	for _, n := range []int{0, 1, 125, 126, 65535, 65536} {
		p := bytes.Repeat([]byte{0xa5}, n)
		r := &HLKO3Cdqw{GX04GFKh1Gz: bytes.NewReader(wsWire(2, true, p)), WXUh5psqY7: ws.StateServerSide}
		h, e := r.NextFrame()
		if e != nil || h.L4ixlgT8VpTo != int64(n) {
			t.Fatal(n, h, e)
		}
		b, e := io.ReadAll(r)
		if e != nil || !bytes.Equal(b, p) {
			t.Fatal(n, len(b), e)
		}
	}
	for _, tc := range []struct {
		wire []byte
		max  int64
		want error
	}{
		{[]byte{0x82, 0}, 0, ws.ErrProtocolMaskRequired},
		{wsWire(2, true, []byte("long")), 3, ErrFrameTooLarge},
		{[]byte{0x82, 0xff, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, ws.ErrHeaderLengthMSB},
	} {
		r := &HLKO3Cdqw{GX04GFKh1Gz: bytes.NewReader(tc.wire), WXUh5psqY7: ws.StateServerSide, ZbPIcS: tc.max}
		if _, e := r.NextFrame(); !errors.Is(e, tc.want) {
			t.Fatal(e, tc.want)
		}
	}
	r := &HLKO3Cdqw{GX04GFKh1Gz: bytes.NewReader(wsWire(2, true, []byte("abc"))[:8]), WXUh5psqY7: ws.StateServerSide}
	if _, e := r.NextFrame(); e != nil {
		t.Fatal(e)
	}
	if _, e := io.ReadAll(r); e != io.ErrUnexpectedEOF {
		t.Fatal(e)
	}
}
func TestOriginalUTF8AcrossReadBoundaries(t *testing.T) {
	r := &Wvd84IIKB{RPKyxC: bytes.NewReader([]byte{0xe4, 0xb8, 0xad})}
	b := make([]byte, 1)
	for i := 0; i < 3; i++ {
		if n, e := r.Read(b); n != 1 || e != nil {
			t.Fatal(n, e)
		}
		if r.Valid() != (i == 2) {
			t.Fatal(i, r.Valid())
		}
	}
	r.Reset(bytes.NewReader([]byte{'a', 0xff}))
	if n, e := r.Read(make([]byte, 2)); n != 1 || e != ErrInvalidUTF8 {
		t.Fatal(n, e)
	}
}
