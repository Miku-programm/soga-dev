package grpcproto

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"reflect"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

func TestOriginalDescriptor(t *testing.T) {
	r, err := gzip.NewReader(bytes.NewReader(file_protos_geo_proto_rawDescGZIP()))
	if err != nil {
		t.Fatal(err)
	}
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	r.Close()
	if len(b) != 653 || fmt.Sprintf("%x", sha256.Sum256(b)) != "2f48fac6d6429957690b783eef9facbe3c5e52116f582b684fff46025e784399" {
		t.Fatalf("descriptor differs from target: len=%d sha=%x", len(b), sha256.Sum256(b))
	}
}

func TestTargetStructLayouts(t *testing.T) {
	cases := []struct {
		v      any
		size   uintptr
		fields map[string]uintptr
	}{
		{HTBZ_woRreQ{}, 88, map[string]uintptr{"AfECsc": 40, "DSsRH80v92kO": 48, "D6LPupTVU": 64}},
		{P3tOWWIHzX{}, 80, map[string]uintptr{"I0dBJNoqAhc": 40, "HGtbzYnVdt0": 56}},
		{A_gvbs6{}, 64, map[string]uintptr{"HKyWw72": 40}},
		{ZlC9UehPDkp{}, 72, map[string]uintptr{"Agmm2h": 40, "HYmWBe": 64}},
		{E4umdYA{}, 80, map[string]uintptr{"OhCky4mjC": 40, "CnonHlPYKJ": 56}},
		{WDWzHppltkvi{}, 64, map[string]uintptr{"HKyWw72": 40}},
		{ZWMyJ4ZKrh{}, 72, map[string]uintptr{"RzqEir": 40, "RO_VnYbH2aZS": 56}},
		{JiHL3Xc{}, 1, map[string]uintptr{"Xu6lP1tsA": 0}},
		{Q27qQ6yjB9m{}, 8, map[string]uintptr{"TT_afC": 0}},
	}
	for _, c := range cases {
		typ := reflect.TypeOf(c.v)
		if typ.Size() != c.size {
			t.Errorf("%s size=%d want=%d", typ, typ.Size(), c.size)
		}
		for name, off := range c.fields {
			f, ok := typ.FieldByName(name)
			if !ok || f.Offset != off {
				t.Errorf("%s.%s offset=%d want=%d", typ, name, f.Offset, off)
			}
		}
	}
}

func TestGeoWireAndUnknownFieldRoundTrip(t *testing.T) {
	// 独立按 wire 编码构造 fixture，校验原字段号、嵌套消息和 oneof。
	attr := protowire.AppendTag(nil, 1, protowire.BytesType)
	attr = protowire.AppendString(attr, "flag")
	attr = protowire.AppendTag(attr, 2, protowire.VarintType)
	attr = protowire.AppendVarint(attr, 1)
	domain := []byte{8, 2}
	domain = protowire.AppendTag(domain, 2, protowire.BytesType)
	domain = protowire.AppendString(domain, "example.com")
	domain = protowire.AppendTag(domain, 3, protowire.BytesType)
	domain = protowire.AppendBytes(domain, attr)
	site := protowire.AppendTag(nil, 1, protowire.BytesType)
	site = protowire.AppendString(site, "TEST")
	site = protowire.AppendTag(site, 2, protowire.BytesType)
	site = protowire.AppendBytes(site, domain)
	unknown := protowire.AppendTag(nil, 123, protowire.VarintType)
	unknown = protowire.AppendVarint(unknown, 77)
	site = append(site, unknown...)
	var got P3tOWWIHzX
	if err := proto.Unmarshal(site, &got); err != nil {
		t.Fatal(err)
	}
	if got.I0dBJNoqAhc != "TEST" || len(got.HGtbzYnVdt0) != 1 {
		t.Fatal(&got)
	}
	d := got.HGtbzYnVdt0[0]
	if d.AfECsc != PCclujlSiD1l(2) || d.DSsRH80v92kO != "example.com" || len(d.D6LPupTVU) != 1 || !d.D6LPupTVU[0].GetBoolValue() {
		t.Fatal(d)
	}
	if !bytes.Equal(got.ProtoReflect().GetUnknown(), unknown) {
		t.Fatal("unknown field discarded")
	}
	encoded, err := proto.Marshal(&got)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(encoded, site) {
		t.Fatalf("wire changed: %x", encoded)
	}
	integer := &ZWMyJ4ZKrh{RzqEir: "n", RO_VnYbH2aZS: &Q27qQ6yjB9m{TT_afC: -19}}
	encoded, err = proto.Marshal(integer)
	if err != nil {
		t.Fatal(err)
	}
	var decoded ZWMyJ4ZKrh
	if err = proto.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.GetIntValue() != -19 || decoded.GetBoolValue() {
		t.Fatal(&decoded)
	}
	ip := &E4umdYA{OhCky4mjC: "TEST", CnonHlPYKJ: []*ZlC9UehPDkp{{Agmm2h: []byte{192, 0, 2, 0}, HYmWBe: 24}}}
	encoded, err = proto.Marshal(ip)
	if err != nil {
		t.Fatal(err)
	}
	var ip2 E4umdYA
	if err = proto.Unmarshal(encoded, &ip2); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(ip, &ip2) {
		t.Fatal(&ip2)
	}
}
