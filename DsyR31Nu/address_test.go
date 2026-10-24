package DsyR31Nu

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
	"testing"
	"time"
	"unsafe"
)

func TestAddressOriginalWireAndText(t *testing.T) {
	if unsafe.Sizeof(FGMp1YT7{}) != 48 {
		t.Fatal("RTTI")
	}
	for _, tc := range []struct {
		wire      []byte
		host      string
		size, cap int
	}{{[]byte{1, 127, 0, 0, 1, 1, 187}, "127.0.0.1", 7, 6}, {append(append([]byte{4}, net.ParseIP("2001:db8::1")...), 1, 187), "2001:db8::1", 19, 18}, {[]byte{3, 4, 'h', 'o', 's', 't', 1, 187}, "host", 8, 0}, {[]byte{3, 0, 1, 187}, "", 4, 0}} {
		a, e := UXgN4RrAvUx(bytes.NewReader(tc.wire))
		if e != nil || a.GetHost() != tc.host || a.GetPort() != 443 || a.AddrDataSize() != tc.size || cap(a.GetIP()) != tc.cap {
			t.Fatal(a, e, tc)
		}
		if tc.cap != 0 {
			tcp, e := a.ToTCPAddr()
			if e != nil || &tcp.IP[0] != &a.b6HasR_[0] {
				t.Fatal("IP alias", e)
			}
		} else {
			if _, e := a.ToUDPAddr(); e == nil || e.Error() != "addr is not ip" {
				t.Fatal(e)
			}
		}
	}
	a := Fqn4mE5("2001:0db8::1", -7)
	b := Fqn4mE5("2001:db8::1", -7)
	if a.String() != "2001:0db8::1:-7" || a.Join() != "[2001:0db8::1]:-7" || a.EqualTo(b) || !a.GetIP().Equal(b.GetIP()) {
		t.Fatal("text identity")
	}
	for _, wire := range [][]byte{{}, {1, 1}, {3}, {3, 2, 'a'}, {4, 1}} {
		if _, e := UXgN4RrAvUx(bytes.NewReader(wire)); e == nil {
			t.Fatal("truncation", wire)
		}
	}
	if _, e := UXgN4RrAvUx(bytes.NewReader([]byte{9})); e == nil || e.Error() != "unknown addr type: 9" {
		t.Fatal(e)
	}
	wire := append([]byte{3, 9}, []byte("127.0.0.1")...)
	wire = binary.BigEndian.AppendUint16(wire, 80)
	a, e := UXgN4RrAvUx(bytes.NewReader(wire))
	if e != nil || !a.HasIP() || len(a.GetIP()) != 16 {
		t.Fatal("numeric domain", a, e)
	}
}
func TestFiniteAndUnlimitedDrainActualPipe(t *testing.T) {
	for _, size := range []int{5, 0, -1} {
		t.Run(time.Duration(size).String(), func(t *testing.T) {
			a, b := net.Pipe()
			defer a.Close()
			defer b.Close()
			done := make(chan error, 1)
			go func() { _, e := b.Write([]byte("abcde")); b.Close(); done <- e }()
			e := Q6wDan(a, size, time.Second)
			if size > 0 && e != nil || size <= 0 && e != io.EOF {
				t.Fatal(e)
			}
			if e = <-done; e != nil {
				t.Fatal(e)
			}
		})
	}
}
