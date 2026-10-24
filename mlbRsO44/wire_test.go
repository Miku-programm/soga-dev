package mlbRsO44

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/apernet/quic-go/quicvarint"
	"io"
	"math"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestOriginalMessageLayout(t *testing.T) {
	typ := reflect.TypeFor[ItLMumzBf]()
	if typ.Size() != 48 {
		t.Fatal(typ.Size())
	}
	names := []string{"D35ycLc5aB6r", "CBHlJenOD", "QmDvxK_8Q", "Kb9FLg7wuNE", "XrDcBXcD", "HmoFPk"}
	offsets := []uintptr{0, 4, 6, 7, 8, 24}
	types := []reflect.Type{reflect.TypeFor[uint32](), reflect.TypeFor[uint16](), reflect.TypeFor[uint8](), reflect.TypeFor[uint8](), reflect.TypeFor[string](), reflect.TypeFor[[]byte]()}
	for i, name := range names {
		f := typ.Field(i)
		if f.Name != name || f.Offset != offsets[i] || f.Type != types[i] || f.Tag != "" {
			t.Fatal(f)
		}
	}
}
func TestAuthHeadersAndOriginalParseUintErrors(t *testing.T) {
	for _, tc := range []struct {
		v    string
		want uint64
	}{{"123", 123}, {"-1", 0}, {"bad", 0}, {"18446744073709551616", math.MaxUint64}} {
		h := http.Header{"Hysteria-Auth": []string{"secret"}, "Hysteria-Cc-Rx": []string{tc.v}}
		auth, rx := GZDAjvxfC(h)
		if auth != "secret" || rx != tc.want {
			t.Fatal(auth, rx)
		}
	}
	for _, auto := range []bool{false, true} {
		h := make(http.Header)
		CIvpFjH270ca(h, true, 987, auto)
		want := "987"
		if auto {
			want = "auto"
		}
		if h.Get("Hysteria-UDP") != "true" || h.Get("Hysteria-CC-RX") != want {
			t.Fatal(h)
		}
		p := h.Get("Hysteria-Padding")
		if len(p) < 256 || len(p) >= 2048 || strings.Trim(p, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") != "" {
			t.Fatal(p)
		}
	}
}
func TestUDPIndependentWireAndAliases(t *testing.T) {
	wire, _ := hex.DecodeString("123456789abc020304686f7374010203")
	m, e := CY9jLhcY1f5(wire)
	if e != nil {
		t.Fatal(e)
	}
	if m.D35ycLc5aB6r != 0x12345678 || m.CBHlJenOD != 0x9abc || m.QmDvxK_8Q != 2 || m.Kb9FLg7wuNE != 3 || m.XrDcBXcD != "host" || !bytes.Equal(m.HmoFPk, []byte{1, 2, 3}) {
		t.Fatal(m)
	}
	wire[len(wire)-1] = 9
	if m.HmoFPk[2] != 9 {
		t.Fatal("payload must alias input")
	}
	wire[9] = 'X'
	if m.XrDcBXcD != "host" {
		t.Fatal("address must be copied")
	}
	out := make([]byte, m.Size())
	if m.Serialize(out) != len(out) || hex.EncodeToString(out) != "123456789abc020304686f7374010209" {
		t.Fatal(out)
	}
	short := bytes.Repeat([]byte{0xaa}, m.Size()-1)
	if m.Serialize(short) != -1 || short[0] != 0xaa {
		t.Fatal("short buffer mutated")
	}
	for n := 0; n < 9; n++ {
		_, e := CY9jLhcY1f5(out[:n])
		if e == nil {
			t.Fatalf("accepted truncation %d", n)
		}
	}
	for _, tc := range []struct {
		suffix []byte
		msg    string
	}{{[]byte{0}, "invalid address length"}, {[]byte{0x48, 1}, "invalid address length"}, {[]byte{1, 'x'}, "invalid message length"}} {
		_, e := CY9jLhcY1f5(append(make([]byte, 8), tc.suffix...))
		if e == nil || e.Error() != tc.msg {
			t.Fatal(e)
		}
	}
}
func TestTCPWireBoundaries(t *testing.T) {
	b := append([]byte{3, 'a', ':', '1', 2, 'X', 'Y'}, []byte("body")...)
	r := bytes.NewReader(b)
	addr, e := T7aXhP(r)
	if e != nil || addr != "a:1" || r.Len() != 4 {
		t.Fatal(addr, e, r.Len())
	}
	for _, tc := range []struct {
		b   []byte
		msg string
	}{{[]byte{0}, "invalid address length"}, {[]byte{0x48, 1}, "invalid address length"}, {[]byte{1, 'x', 0x50, 1}, "invalid padding length"}} {
		_, e := T7aXhP(bytes.NewReader(tc.b))
		if e == nil || e.Error() != tc.msg {
			t.Fatal(e)
		}
	}
	for _, b := range [][]byte{nil, {1}, {1, 'x'}, {1, 'x', 2, 'P'}} {
		if _, e := T7aXhP(bytes.NewReader(b)); e == nil {
			t.Fatal("truncation accepted")
		}
	}
	for _, ok := range []bool{true, false} {
		var out bytes.Buffer
		if e := C3FDNGT2zI(&out, ok, "test"); e != nil {
			t.Fatal(e)
		}
		status, _ := out.ReadByte()
		if (status == 0) != ok {
			t.Fatal(status)
		}
		n, _ := quicvarint.Read(&out)
		msg := make([]byte, n)
		io.ReadFull(&out, msg)
		if string(msg) != "test" {
			t.Fatal(msg)
		}
		n, _ = quicvarint.Read(&out)
		if n < 128 || n >= 1024 || out.Len() != int(n) {
			t.Fatal(n, out.Len())
		}
	}
	sentinel := errors.New("write failed")
	if e := C3FDNGT2zI(shortWriter{sentinel}, true, ""); e != sentinel {
		t.Fatal(e)
	}
	if e := C3FDNGT2zI(shortWriter{}, true, ""); e != nil {
		t.Fatal("original ignores short count", e)
	}
}

type shortWriter struct{ err error }

func (w shortWriter) Write([]byte) (int, error) { return 0, w.err }
func TestOriginalVarintBoundaries(t *testing.T) {
	for _, v := range []uint64{0, 63, 64, 16383, 16384, 1073741823, 1073741824, 4611686018427387903} {
		b := make([]byte, 8)
		n := vVaooWJ93Bp(b, v)
		want := quicvarint.Append(nil, v)
		if !bytes.Equal(b[:n], want) {
			t.Fatal(v, b, want)
		}
	}
	b := []byte{0}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("short write must panic")
			}
		}()
		vVaooWJ93Bp(b, 0x1234)
	}()
	if b[0] != 0x52 {
		t.Fatal("partial first byte", b)
	}
	func() {
		defer func() {
			if got := recover(); got != "0x4000000000000000 doesn't fit into 62 bits" {
				t.Fatal(got)
			}
		}()
		vVaooWJ93Bp(nil, 1<<62)
	}()
}
