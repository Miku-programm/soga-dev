package vlessvision

import (
	"bytes"
	"encoding/binary"
	"recovered.local/soga/buf/pool"
	"reflect"
	"testing"
)

func TestVisionStateRTTI(t *testing.T) {
	typ := reflect.TypeFor[MxS0kZ]()
	if typ.Size() != 64 {
		t.Fatal(typ.Size())
	}
	for i, want := range []uintptr{0, 24, 32, 33, 34, 36, 40, 44, 45, 48, 52, 56, 60, 61, 62} {
		if typ.Field(i).Offset != want {
			t.Fatal(typ.Field(i))
		}
	}
}
func TestVisionSplitBoundaryAndAliasing(t *testing.T) {
	for _, length := range []int{0, 8170, 8171, 16342, 40000} {
		data := make([]byte, length)
		chunks := AaBEhR(nil, data)
		if len(chunks) == 0 {
			t.Fatal("empty input still appends")
		}
		if !bytes.Equal(bytes.Join(chunks, nil), data) {
			t.Fatal("reassembly")
		}
		for _, chunk := range chunks {
			if len(chunk) >= 8171 {
				t.Fatal(len(chunk))
			}
		}
	}
	data := make([]byte, 9000)
	copy(data[7000:], []byte{0x17, 3, 3})
	chunks := AaBEhR(nil, data)
	if len(chunks[0]) != 7000 {
		t.Fatal(len(chunks[0]))
	}
	chunks[0][0] = 4
	if data[0] != 4 {
		t.Fatal("original chunks alias data")
	}
}
func TestVisionPaddingRoundTripAndPrefixConsumed(t *testing.T) {
	for _, size := range []int{0, 100, 899, 900, 8170} {
		for _, long := range []bool{false, true} {
			t.Run(string(rune(size+65))+map[bool]string{true: "long", false: "short"}[long], func(t *testing.T) {
				out := pool.S4OQXWIZ(0)
				defer out.Free()
				uuid := []byte("0123456789abcdef")
				prefix := bytes.Clone(uuid)
				data := bytes.Repeat([]byte{'p'}, size)
				IFq7RHaxl(out, data, 1, &prefix, long)
				wire := out.Bytes()
				if prefix != nil || !bytes.Equal(wire[:16], uuid) || wire[16] != 1 || int(binary.BigEndian.Uint16(wire[17:])) != size {
					t.Fatal("header")
				}
				padding := int(binary.BigEndian.Uint16(wire[19:]))
				if len(wire) != 21+size+padding || padding > 8171-size {
					t.Fatal("padding bounds")
				}
				if long && size < 900 {
					if padding < 900-size || padding >= 1400-size {
						t.Fatal(padding)
					}
				} else if padding >= 256 {
					t.Fatal(padding)
				}
				state := &MxS0kZ{RW3eHGaIH0ya: uuid, Z43zq7: -1, ZOEKPa7JPx5: -1, PIsBMe: -1}
				decoded := DtcbafIY7SUd(out, state)
				if decoded == out {
					t.Fatal("no allocation")
				}
				defer decoded.Free()
				if !bytes.Equal(decoded.Bytes(), data) || state.Z43zq7 != -1 || state.EXbj1L5 != 1 {
					t.Fatalf("decode %d %+v", decoded.Len(), state)
				}
			})
		}
	}
}
func TestVisionUnpaddingIncrementalAndTrailingBytes(t *testing.T) {
	uuid := []byte("0123456789abcdef")
	state := &MxS0kZ{RW3eHGaIH0ya: uuid, Z43zq7: -1, ZOEKPa7JPx5: -1, PIsBMe: -1}
	first := pool.S4OQXWIZ(0)
	defer first.Free()
	_, _ = first.Write(append(bytes.Clone(uuid), 0, 0, 3, 0, 2, 'a'))
	out := DtcbafIY7SUd(first, state)
	defer out.Free()
	if string(out.Bytes()) != "a" || state.ZOEKPa7JPx5 != 2 {
		t.Fatalf("%q %+v", out.Bytes(), state)
	}
	second := pool.S4OQXWIZ(0)
	defer second.Free()
	_, _ = second.Write([]byte{'b', 'c', 9, 9, 2, 0, 1, 0, 1, 'd', 8, 'r', 'a', 'w'})
	out2 := DtcbafIY7SUd(second, state)
	defer out2.Free()
	if string(out2.Bytes()) != "bcdraw" || string(second.Bytes()) != "raw" || state.EXbj1L5 != 2 || state.Z43zq7 != -1 {
		t.Fatalf("%q %q %+v", out2.Bytes(), second.Bytes(), state)
	}
	unmatched := pool.S4OQXWIZ(0)
	defer unmatched.Free()
	_, _ = unmatched.Write(bytes.Repeat([]byte{'z'}, 25))
	if DtcbafIY7SUd(unmatched, state) != unmatched {
		t.Fatal("unmatched should preserve pointer")
	}
}
