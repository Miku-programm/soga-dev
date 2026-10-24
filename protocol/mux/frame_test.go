package mux

import (
	"bytes"
	"encoding/binary"
	"io"
	"recovered.local/soga/buf/pool"
	"recovered.local/soga/buf/bufiface"
	"reflect"
	"testing"
)

var _ bufiface.Mf3phSOVqgf = (*pool.X_aXX5)(nil)

func frameBytes(meta, payload []byte, hasData bool) []byte {
	p := binary.BigEndian.AppendUint16(nil, uint16(len(meta)))
	p = append(p, meta...)
	if hasData {
		p = binary.BigEndian.AppendUint16(p, uint16(len(payload)))
		p = append(p, payload...)
	}
	return p
}
func TestOriginalMuxFrameLayoutAndFlags(t *testing.T) {
	typ := reflect.TypeFor[CUwik6VP2WU0]()
	if typ.Size() != 40 {
		t.Fatal(typ)
	}
	for i, want := range []uintptr{0, 2, 3, 4, 8, 16, 24} {
		if typ.Field(i).Offset != want {
			t.Fatal(typ.Field(i))
		}
	}
	for flags := 0; flags < 256; flags++ {
		f := CUwik6VP2WU0{m3vf0n8: byte(flags)}
		if f.HasData() != (flags&1 != 0) || f.HasError() != (flags&2 != 0) {
			t.Fatal(flags)
		}
	}
	for i, want := range []string{"Unknown", "New", "Keep", "End", "KeepAlive", "Unknown"} {
		if FEGjbz2y(i).String() != want {
			t.Fatal(i)
		}
	}
}
func TestOriginalMuxFramePayloadOwnershipAndAddress(t *testing.T) {
	// New: id 0x1234, data flag, network UDP(2), port 53, domain a.test.
	meta := []byte{0x12, 0x34, 1, 1, 2, 0, 53, 2, 6}
	meta = append(meta, []byte("a.test")...)
	wire := frameBytes(meta, []byte("payload"), true)
	r := bytes.NewReader(append(wire, 99))
	f, err := HlVrAnC111(r)
	if err != nil {
		t.Fatal(err)
	}
	if f.ljtkJw != 0x1234 || f.ePmcTrBC != 2 || f.gQBBY59xyosW || f.hUHpeGfV.String() != "a.test:53" || string(f.c8AtRj7hb.Bytes()) != "payload" || r.Len() != 1 {
		t.Fatal(f, r.Len())
	}
	f.FreeData()
	f.FreeData()
	if f.c8AtRj7hb != nil {
		t.Fatal("released buffer not cleared")
	}
	meta[2] = 2
	f, err = HlVrAnC111(bytes.NewReader(frameBytes(meta, nil, true)))
	if err != nil || !f.gQBBY59xyosW || f.c8AtRj7hb == nil || !f.c8AtRj7hb.IsEmpty() {
		t.Fatal(f, err)
	}
	f.FreeData()
	for _, cmd := range []byte{2, 3, 4} {
		f, err := HlVrAnC111(bytes.NewReader(frameBytes([]byte{0, 1, cmd, 0}, nil, false)))
		if err != nil || f.hUHpeGfV != nil || f.gQBBY59xyosW {
			t.Fatal(f, err)
		}
	}
}
func TestOriginalMuxFrameBoundariesAndConsumption(t *testing.T) {
	for _, tc := range []struct {
		n    uint16
		want string
	}{{0, "mux metadata size < 4: 0"}, {3, "mux metadata size < 4: 3"}, {513, "mux metadata size is too big: 513"}} {
		_, err := HlVrAnC111(bytes.NewReader(binary.BigEndian.AppendUint16(nil, tc.n)))
		if err == nil || err.Error() != tc.want {
			t.Fatal(tc, err)
		}
	}
	meta := bytes.Repeat([]byte{99}, 512)
	copy(meta, []byte{0, 2, 4, 0})
	r := bytes.NewReader(frameBytes(meta, nil, false))
	if _, err := HlVrAnC111(r); err != nil || r.Len() != 0 {
		t.Fatal(err, r.Len())
	}
	if _, err := HlVrAnC111(bytes.NewReader(frameBytes([]byte{0, 1, 0, 0}, nil, false))); err == nil || err.Error() != "unknown mux cmd: Unknown" {
		t.Fatal(err)
	}
	if _, err := HlVrAnC111(bytes.NewReader(frameBytes([]byte{0, 1, 1, 0}, nil, false))); err != io.ErrUnexpectedEOF {
		t.Fatal("new missing network", err)
	}
	full := frameBytes([]byte{0, 1, 2, 1}, []byte("abcdef"), true)
	for i := 0; i < len(full); i++ {
		if _, err := HlVrAnC111(bytes.NewReader(full[:i])); err == nil {
			t.Fatalf("truncated prefix %d accepted", i)
		}
	}
}
