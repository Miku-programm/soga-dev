package gk1z9v

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"hash/fnv"
	"reflect"
	"strings"
	"testing"
)

func headerFixture(command, kind, options byte, address []byte) []byte {
	b := make([]byte, 38)
	b[0] = 1
	for i := 0; i < 32; i++ {
		b[i+1] = byte(i)
	}
	b[33] = 0x42
	b[34] = options
	b[35] = 0x23
	b[37] = command
	if command != 3 {
		b = append(b, 1, 187, kind)
		b = append(b, address...)
	}
	b = append(b, 0xaa, 0xbb)
	return finishHeaderChecksum(b)
}
func finishHeaderChecksum(b []byte) []byte {
	h := fnv.New32a()
	h.Write(b)
	return binary.BigEndian.AppendUint32(b, h.Sum32())
}
func TestNativeVMessHeaderAddressesAndOptions(t *testing.T) {
	if reflect.TypeFor[Wpqfspy6zZ_B]().Size() != 72 || reflect.TypeFor[URTFuf2wVWf]().Size() != 56 {
		t.Fatal("RTTI layout")
	}
	for _, tc := range []struct {
		name                   string
		command, kind, options byte
		address                []byte
		host                   string
		literal, s, m, pad     bool
	}{
		{"ipv4", 1, 1, 13, []byte{192, 0, 2, 7}, "192.0.2.7", true, true, true, true},
		{"ipv6", 2, 3, 4, []byte{0x20, 1, 0x0d, 0xb8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7}, "2001:db8::7", true, false, false, false},
		{"domain", 1, 2, 5, append([]byte{11}, []byte("example.com")...), "example.com", false, true, true, false},
		{"mux", 3, 0, 0, nil, "", false, false, false, false},
		{"long-domain", 1, 2, 1, append([]byte{255}, bytes.Repeat([]byte{'z'}, 255)...), strings.Repeat("z", 255), false, true, false, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			raw := headerFixture(tc.command, tc.kind, tc.options, tc.address)
			b := bytes.NewBuffer(append(append([]byte(nil), raw...), []byte("next-frame")...))
			h, err := Gk1af7wDMnD(b)
			if err != nil {
				t.Fatal(err)
			}
			if b.String() != "next-frame" || h.daehRKryOT != 0x42 || h.mTRhOfKViN != 3 || h.uaHWu6c != 2 || h.OptS() != tc.s || h.OptM() != tc.m || h.OptPadding() != tc.pad {
				t.Fatalf("header %#v tail %q", h, b.String())
			}
			if h.NetworkType() != 0 && tc.command == 1 || h.NetworkType() != 1 && tc.command == 2 || h.NetworkType() != 2 && tc.command == 3 {
				t.Fatal("network")
			}
			if tc.command == 3 {
				if h.GetDestAddr() != nil {
					t.Fatal("mux address")
				}
			} else if a := h.GetDestAddr(); a.GetHost() != tc.host || a.GetPort() != 443 || a.IsIP() != tc.literal {
				t.Fatal(a)
			}
			for n := 0; n < len(raw); n++ {
				if out, err := Gk1af7wDMnD(bytes.NewReader(raw[:n])); out != nil || err == nil {
					t.Fatalf("short %d/%d: %#v %v", n, len(raw), out, err)
				}
			}
			raw[len(raw)-1] ^= 1
			if _, err := Gk1af7wDMnD(bytes.NewReader(raw)); err == nil || err.Error() != "VMess Header fnv invalid" {
				t.Fatal(err)
			}
		})
	}
}
func TestNativeVMessHeaderKnownBytesAndErrors(t *testing.T) {
	raw, _ := hex.DecodeString("01000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f420d23000101bb01c0000207aabb23920a14")
	// 固定向量的校验独立检查，避免仅由测试构造器自证。
	h, err := Gk1af7wDMnD(bytes.NewReader(raw))
	if err != nil {
		t.Fatal(err)
	}
	if h.GetDestAddr().HostPort() != "192.0.2.7:443" {
		t.Fatal(h.GetDestAddr())
	}
	raw[0] = 2
	if _, err := Gk1af7wDMnD(bytes.NewReader(raw)); err == nil || err.Error() != "VMess version is not 1: 2" {
		t.Fatal(err)
	}
	raw[0] = 1
	raw[40] = 4
	if _, err := Gk1af7wDMnD(bytes.NewReader(raw)); err == nil || err.Error() != "unknown addr type: 4" {
		t.Fatal(err)
	}
	unknown := headerFixture(9, 1, 0, []byte{127, 0, 0, 1})
	h, err = Gk1af7wDMnD(bytes.NewReader(unknown))
	if err != nil || h.NetworkType() != 255 {
		t.Fatal(h, err)
	}
	response := &URTFuf2wVWf{iQlcjP47sl: 7, mFdVdC: 8, j3mthjsl: 9, maCTf0DLoP: []byte{1, 2}}
	if response.HeaderLen() != 4 || !bytes.Equal(response.ToBytes(), []byte{7, 8, 9, 0}) {
		t.Fatal(response.ToBytes())
	}
}
