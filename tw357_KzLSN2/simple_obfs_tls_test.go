package tw357_KzLSN2

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	S "recovered.local/soga/ad5KEbw_EzWk"
	"testing"
	"time"
	"unsafe"
)

func clientTLSFixture(early []byte, host string) []byte {
	record := []byte{1, 0, 0, 0, 3, 3}
	random := make([]byte, 32)
	rand.Read(random)
	binary.BigEndian.PutUint32(random, uint32(time.Now().Unix()))
	record = append(record, random...)
	record = append(record, 32)
	record = append(record, bytes.Repeat([]byte{0x59}, 32)...)
	record = append(record, make([]byte, 62)...)
	record = append(record, 0, 35, byte(len(early)>>8), byte(len(early)))
	record = append(record, early...)
	record = append(record, make([]byte, 7)...)
	record = append(record, byte(len(host)>>8), byte(len(host)))
	record = append(record, host...)
	binary.BigEndian.PutUint16(record[2:], uint16(len(record)-4))
	return append([]byte{22, 3, 1, byte(len(record) >> 8), byte(len(record))}, record...)
}
func TestSimpleTLSHandshakeAndRecords(t *testing.T) {
	if unsafe.Sizeof(cqVDRTx{}) != 104 {
		t.Fatal("RTTI")
	}
	hello := clientTLSFixture([]byte("early"), "example.org")
	wire := append(append([]byte{}, hello...), 23, 3, 3, 0, 6)
	wire = append(wire, []byte("record")...)
	c := crpBB1_mZ(&wireFixture{read: bytes.NewReader(wire)}).(*cqVDRTx)
	info := &S.L7E7H2aW8KVs{}
	c.SetServerInfo(info)
	defer c.Free()
	var got []byte
	for len(got) < 11 {
		b := make([]byte, 2)
		n, e := c.Read(b)
		if e != nil {
			t.Fatal(e)
		}
		got = append(got, b[:n]...)
	}
	if string(got) != "earlyrecord" || info.EFLMTH != "example.org" || c.jCBNufN != 1 || c.bg3qmAv != 0 || c.pDziAaUrdOb != nil {
		t.Fatal(string(got), info, c)
	}
	duplicate := crpBB1_mZ(&wireFixture{read: bytes.NewReader(hello)}).(*cqVDRTx)
	duplicate.SetServerInfo(&S.L7E7H2aW8KVs{})
	defer duplicate.Free()
	if e := duplicate.a8xWy1CIZ3n(); e == nil || e.Error() != "tls auth repeat attack detect" {
		t.Fatal(e)
	}
}
func TestSimpleTLSErrors(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func([]byte)
		want   string
	}{
		{"prefix", func(p []byte) { p[0] = 23 }, "tls handshake failed by prefix: \\x16\\x03\\x01"},
		{"hello", func(p []byte) { p[5] = 2 }, "tls_auth not client hello message"},
		{"length", func(p []byte) { p[8]++ }, "tls_auth wrong message size"},
		{"version", func(p []byte) { p[10] = 1 }, "tls_auth wrong tls version"},
		{"session", func(p []byte) { p[43] = 31 }, "tls_auth wrong session id length"},
		{"time", func(p []byte) { binary.BigEndian.PutUint32(p[11:], 1) }, "tls auth wrong time"},
		{"extension", func(p []byte) { p[138] = 1 }, "ext header error"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := clientTLSFixture(nil, "x")
			tc.mutate(h)
			c := crpBB1_mZ(&wireFixture{read: bytes.NewReader(h)}).(*cqVDRTx)
			defer c.Free()
			c.SetServerInfo(&S.L7E7H2aW8KVs{})
			e := c.a8xWy1CIZ3n()
			if e == nil || e.Error() != tc.want {
				t.Fatal(e)
			}
		})
	}
	c := crpBB1_mZ(&wireFixture{read: bytes.NewReader([]byte{22, 3, 3, 0, 0})}).(*cqVDRTx)
	defer c.Free()
	c.jCBNufN = 1
	if n, e := c.Read(nil); n != 0 || e == nil || e.Error() != "data prefix wrong:" {
		t.Fatal(n, e)
	}
}
func TestSimpleTLSWriteFraming(t *testing.T) {
	w := &wireFixture{}
	c := crpBB1_mZ(w).(*cqVDRTx)
	defer c.Free()
	c.a1yhtxpPH = bytes.Repeat([]byte{0x53}, 32)
	payload := bytes.Repeat([]byte("data"), 20000)
	n, e := c.Write(payload)
	if n != len(payload) || e != nil {
		t.Fatal(n, e)
	}
	wire := w.output.Bytes()
	if !bytes.Equal(wire[:11], []byte{22, 3, 1, 0, 91, 2, 0, 0, 87, 3, 3}) || !bytes.Equal(wire[44:76], c.a1yhtxpPH) {
		t.Fatal("hello header")
	}
	if !bytes.Equal(wire[96:102], []byte{20, 3, 3, 0, 1, 1}) {
		t.Fatal("change cipher", wire[96:102])
	}
	wire = wire[102:]
	var got []byte
	first := true
	for len(wire) > 0 {
		typ := byte(23)
		if first {
			typ = 22
			first = false
		}
		if !bytes.Equal(wire[:3], []byte{typ, 3, 3}) {
			t.Fatal("record")
		}
		size := int(binary.BigEndian.Uint16(wire[3:5]))
		got = append(got, wire[5:5+size]...)
		wire = wire[5+size:]
	}
	if !bytes.Equal(got, payload) {
		t.Fatal("payload")
	}
	broken := &wireFixture{writeErr: io.ErrClosedPipe}
	d := crpBB1_mZ(broken).(*cqVDRTx)
	defer d.Free()
	if n, e := d.Write([]byte("x")); n != 0 || !errors.Is(e, io.ErrClosedPipe) || d.sBRQD6eVrE != 1 {
		t.Fatal(n, e, d.sBRQD6eVrE)
	}
}
