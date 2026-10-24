package CAodbhZf

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
	"testing"
)

func TestOriginalByteReaderCursorAndOwnership(t *testing.T) {
	typ := reflect.TypeFor[DCeuEa2b]()
	if typ.Size() != 32 || typ.Field(0).Name != "cTs3S5" || typ.Field(1).Name != "fAIkcTHidRIc" || typ.Field(1).Offset != 24 {
		t.Fatal(typ)
	}
	data := []byte{1, 2, 3, 4}
	r := NewDCeuEa2b(data)
	if n, e := r.Read(nil); n != 0 || e != nil {
		t.Fatal(n, e)
	}
	if _, e := r.ReadBytes(5); e != io.ErrUnexpectedEOF || r.Len() != 4 {
		t.Fatal(e, r.Len())
	}
	b, e := r.ReadBytes(2)
	if e != nil || cap(b) != 4 {
		t.Fatal(b, e)
	}
	b[0] = 9
	if data[0] != 9 {
		t.Fatal("read copied storage")
	}
	if e = r.Skip(-1); e != nil || r.Len() != 3 {
		t.Fatal(e, r.Len())
	}
	if v, e := r.ReadUInt16(); e != nil || v != 0x0203 {
		t.Fatal(v, e)
	}
	if v, e := r.ReadUInt8(); e != nil || v != 4 {
		t.Fatal(v, e)
	}
	if _, e = r.ReadByte(); e != io.ErrUnexpectedEOF {
		t.Fatal(e)
	}
	if _, e = r.Read(nil); e != io.EOF {
		t.Fatal(e)
	}
	if b = r.ReadLeft(); b == nil || len(b) != 0 {
		t.Fatal("empty remainder must be non-nil")
	}
	if !bytes.Equal(r.AllData(), data) {
		t.Fatal("AllData filtered consumed bytes")
	}
	r = NewDCeuEa2b([]byte("test"))
	s, e := r.ReadString(2)
	r.AllData()[0] = 'x'
	if e != nil || s != "te" {
		t.Fatal(s, e)
	}
	if string(r.ReadLeft()) != "st" || !r.IsEmpty() {
		t.Fatal("remainder cursor")
	}
}
func TestOriginalByteReaderNumbersAndIP(t *testing.T) {
	data := make([]byte, 12)
	binary.BigEndian.PutUint32(data, 0x10203040)
	binary.BigEndian.PutUint64(data[4:], 0xfedcba9876543210)
	r := NewDCeuEa2b(data)
	a, e := r.ReadUInt32()
	b, e2 := r.ReadUInt64()
	if a != 0x10203040 || b != 0xfedcba9876543210 || e != nil || e2 != nil {
		t.Fatal(a, b, e, e2)
	}
	if _, e = r.ReadUInt64(); e != io.ErrUnexpectedEOF {
		t.Fatal(e)
	}
	r = NewDCeuEa2b(append([]byte{127, 0, 0, 1}, make([]byte, 16)...))
	v, e := r.ReadIPv4()
	v6, e6 := r.ReadIPv6()
	if v != "127.0.0.1" || v6 != "::" || e != nil || e6 != nil {
		t.Fatal(v, v6, e, e6)
	}
}
