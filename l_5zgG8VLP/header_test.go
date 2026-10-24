package l_5zgG8VLP

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"recovered.local/soga/D_zCag1qTV"
	"recovered.local/soga/GYzZ5v"
	"reflect"
	"testing"
)

// 测试只给原 net.Conn 签名提供可读内存；生产函数无替换接口。
type readConn struct {
	net.Conn
	r *bytes.Reader
}

func (c *readConn) Read(p []byte) (int, error) { return c.r.Read(p) }
func connBytes(data []byte) *readConn          { return &readConn{r: bytes.NewReader(data)} }
func TestHeaderRTTI(t *testing.T) {
	typ := reflect.TypeFor[ZfmpaHtmu]()
	if typ.Size() != 24 {
		t.Fatal(typ.Size())
	}
	for i, offset := range []uintptr{0, 1, 8, 16} {
		if typ.Field(i).Offset != offset {
			t.Fatal(typ.Field(i))
		}
	}
	if reflect.TypeFor[KYlGQuja]().Kind() != reflect.Uint8 {
		t.Fatal("command type")
	}
}
func TestAddonsFrameAndUnknownFields(t *testing.T) {
	want := &GYzZ5v.N_SQ5a{AfJaYuHm3H: "xtls-rprx-vision", S3EXSAVeAUeF: []byte{1, 2, 3}}
	encoded, err := proto.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	encoded = append(encoded, 0x18, 0x07)
	conn := connBytes(append(append([]byte{byte(len(encoded))}, encoded...), 0xfe))
	got, err := a8D0ppAPaqu(conn)
	if err != nil || got.GetFlow() != want.GetFlow() || !bytes.Equal(got.GetSeed(), want.GetSeed()) || conn.r.Len() != 1 {
		t.Fatalf("%+v %v", got, err)
	}
	if !bytes.Equal(got.ProtoReflect().GetUnknown(), []byte{0x18, 7}) {
		t.Fatal("protobuf unknown bytes lost")
	}
	empty, err := a8D0ppAPaqu(connBytes([]byte{0}))
	if empty == nil || err != nil {
		t.Fatalf("%+v %v", empty, err)
	}
	for _, bad := range [][]byte{nil, {2, 1}, {1, 0xff}} {
		if got, err := a8D0ppAPaqu(connBytes(bad)); got != nil || err == nil {
			t.Fatalf("%x: %+v %v", bad, got, err)
		}
	}
}
func TestHeaderCommandsAndVersion(t *testing.T) {
	for _, command := range []byte{1, 2, 4, 255} {
		conn := connBytes([]byte{0, command, 1, 187, 1, 127, 0, 0, 1, 0xfe})
		h, err := vF01yD4y(conn, 0)
		if err != nil || h.aj8eMgsHQRY != KYlGQuja(command) || h.if2alzq2rt6.Join() != "127.0.0.1:443" || h.qssMeIUTR == nil || conn.r.Len() != 1 {
			t.Fatalf("%+v %v", h, err)
		}
	}
	conn := connBytes([]byte{0, 3, 0xfe})
	h, err := vF01yD4y(conn, 0)
	if err != nil || h.if2alzq2rt6 != nil || h.aj8eMgsHQRY != 3 || conn.r.Len() != 1 || h.ResponseHeaderLen() != 2 {
		t.Fatalf("%+v %v", h, err)
	}
	conn = connBytes([]byte{0, 3})
	h, err = vF01yD4y(conn, 255)
	if h != nil || err == nil || err.Error() != "unknown vless version: 255" || conn.r.Len() != 2 {
		t.Fatalf("%+v %v", h, err)
	}
	if h, err = vF01yD4y(connBytes([]byte{0}), 0); h != nil || err != io.EOF {
		t.Fatalf("%+v %v", h, err)
	}
}

func TestResponseHeaderWritesVersionAndZeroAddons(t *testing.T) {
	b := D_zCag1qTV.S4OQXWIZ(1)
	defer b.Free()
	b.WriteByte(0xab)
	h := &ZfmpaHtmu{h8FCLanKLDuj: 9, qssMeIUTR: &GYzZ5v.N_SQ5a{AfJaYuHm3H: "ignored"}}
	h.WriteResponseHeader(b)
	if !bytes.Equal(b.Bytes(), []byte{0xab, 9, 0}) {
		t.Fatal(b.Bytes())
	}
}
